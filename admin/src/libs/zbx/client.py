import logging
from typing import Any, Generic, Literal, TypeAlias, TypeVar, overload

from httpx import AsyncClient, HTTPError
from pydantic import BaseModel

from src.core.config import settings
from src.libs.zbx import schemas

logger = logging.getLogger(__name__)

_ZbxCreateSchema = TypeVar("_ZbxCreateSchema", bound=BaseModel)
_ZbxQuerySchema = TypeVar("_ZbxQuerySchema", bound=BaseModel)
_ZbxUpdateSchema = TypeVar("_ZbxUpdateSchema", bound=BaseModel)
_ZbxObjects: TypeAlias = Literal[
    "host",
    "hostgroup",
    "template",
    "templategroup",
    "proxy",
    "item",
    "history",
    "mediatype",
    "action",
    "configuration",
    "hostinterface",
    "usermacro",
    "event",
]

ParamsT = TypeVar("ParamsT", str, dict, list)


class ZbxError(Exception):
    pass


class Client:
    def __init__(self, url: str, token: str) -> None:
        self.url = url
        self.token = token
        self.headers = {"Content-Type": "application/json-rpc", "Authorization": "Bearer " + self.token}
        self.host = _Host(self.url, self.token, self.headers, "host")
        self.hostgroup = _HostGroup(self.url, self.token, self.headers, "hostgroup")
        self.hostinterface = _HostInterface(self.url, self.token, self.headers, "hostinterface")
        self.template = _Template(self.url, self.token, self.headers, "template")
        self.templategroup = _TemplateGroup(self.url, self.token, self.headers, "templategroup")
        self.proxy = _Proxy(self.url, self.token, self.headers, "proxy")
        self.item = _Item(self.url, self.token, self.headers, "item")
        self.history = _History(self.url, self.token, self.headers, "history")
        self.mediatype = _MediaType(self.url, self.token, self.headers, "mediatype")
        self.action = _Action(self.url, self.token, self.headers, "action")
        self.configuration = _Configuration(self.url, self.token, self.headers, "configuration")
        self.usermacro = _UserMacro(self.url, self.token, self.headers, "usermacro")
        self.event = _Event(self.url, self.token, self.headers, "event")


class ZbxBase(Generic[_ZbxCreateSchema, _ZbxQuerySchema, _ZbxUpdateSchema]):
    __slots__ = ["url", "token", "headers", "obj_name"]

    def __init__(self, url: str, token: str, headers: dict[str, Any], obj_name: _ZbxObjects) -> None:
        self.url = f"{url}/api_jsonrpc.php" if url[-1] != "/" else url[:-1]
        self.token = token
        self.headers = headers
        self.obj_name = obj_name

    async def create(self, client: AsyncClient, data: _ZbxCreateSchema) -> int | None:
        """create zbx object

        Args:
            client (AsyncClient): httpx async client
            data (ZbxCreateSchema): Generic zbx object create schema defined by pydantic

        Returns:
            str | None: return zabbix object id
        """
        response = await self.rpc(self.obj_name + ".create", data.model_dump(exclude_none=True), client)
        if response:
            return self.get_id(response, self.obj_name)  # type: ignore

    @overload
    async def get(
        self,
        client: AsyncClient,
        data: _ZbxQuerySchema,
        handler: Literal["return_ids"] = "return_ids",
    ) -> list[int] | None:
        ...

    @overload
    async def get(
        self,
        client: AsyncClient,
        data: _ZbxQuerySchema,
        handler: Literal["return_id"] = "return_id",
    ) -> int | None:
        ...

    @overload
    async def get(
        self,
        client: AsyncClient,
        data: _ZbxQuerySchema,
        handler: Literal["return_name"] = "return_name",
    ) -> str | None:
        ...

    @overload
    async def get(
        self,
        client: AsyncClient,
        data: _ZbxQuerySchema,
        handler: Literal[None] = None,
    ) -> Any:
        ...

    async def get(
        self,
        client: AsyncClient,
        data: _ZbxQuerySchema,
        handler: Literal["return_ids", "return_id", "return_name", None] = None,
    ) -> int | list[int] | str | Any | None:
        """get zbx object with filters

        Args:
            client (AsyncClient): httpx async client
            data (ZbxQuerySchema): Generic query schema defined by pydantic
            handler (Literal[&quot;return_ids&quot;, &quot;return_id&quot;, &quot;return_name&quot;] | None): \n
            Defaults to None, return the raw data depends on the `output` options in query schema. \n
            `return_ids`: return list of object ids \n
            `return_id`: return single object id \n
            `return_name`: not Implemented right now\n
        """
        response = await self.rpc(self.obj_name + ".get", data.model_dump(exclude_none=True), client)
        if not handler:
            return response
        else:
            response = getattr(self, handler)(response, self.obj_name)
            return response

    async def update(self, client: AsyncClient, data: _ZbxUpdateSchema) -> int | None:
        """update zbx object

        Args:
            client (AsyncClient): httpx async client
            data (ZbxUpdateSchema): Generic zbx object update schema defined by pydantic

        Returns:
            str | None: return zbx object id
        """
        response = await self.rpc(self.obj_name + ".update", data.model_dump(exclude_unset=True), client)
        if response:
            return self.get_id(response, self.obj_name)  # type: ignore

    async def delete(self, client: AsyncClient, ids: list[int] | int) -> dict[str, Any] | list[Any] | None:
        """delete zbx object by id or ids

        Args:
            client (AsyncClient): httpx async client
            ids (list[str] | str): object id or ids

        Returns:
            _type_: _description_
        """
        ids = self.id_trans_to_list(ids)
        if ids:
            response = await self.rpc(self.obj_name + ".delete", ids, client=client)
            return response

    async def rpc(self, method: str, params: ParamsT, client: AsyncClient) -> dict[str, Any] | list[Any] | None:
        """reusable method for zbx api request

        Args:
            method (str): rpc method
            params (ParamsT): rpc params
            client (AsyncClient): httpx async client

        Returns:
            dict[str, Any] | None: _description_
        """
        request_id = 1
        payload = {"jsonrpc": "2.0", "method": method, "params": params, "id": request_id}
        try:
            response = await client.post(url=self.url, headers=self.headers, json=payload)
            assert response.status_code == 200, response.json
            result = response.json()
            if "error" not in result.keys():
                return result["result"]
            else:
                logger.error(f"{method}-{params}: {result['error']}")
                raise ZbxError(result["error"])
        except HTTPError as e:
            logger.exception("Send request to monitor server failed")
            raise ZbxError from e

    @staticmethod
    def get_id(data: dict[str, Any], name: str) -> int | None:
        id_name_mapping = {
            "host": "hostids",
            "hostgroup": "groupids",
            "hostinterface": "interfaceids",
            "template": "templateids",
            "templategroup": "groupids",
            "proxy": "proxyids",
            "item": "itemids",
            "mediatype": "mediatypeids",
            "action": "actionids",
            "usermacro": "hostmacroids",
        }
        if data:
            return int(data[id_name_mapping[name]][0])

    @staticmethod
    def return_ids(data: list[dict[str, Any]], name: str) -> list[int] | None:
        id_name_mapping = {
            "host": "hostid",
            "hostgroup": "groupid",
            "hostinterface": "interfaceid",
            "template": "templateid",
            "templategroup": "groupid",
            "proxy": "proxyid",
            "item": "itemid",
            "mediatype": "mediatypeid",
            "action": "actionid",
            "usermacro": "hostmacroids",
        }
        if data:
            return [int(item[id_name_mapping[name]]) for item in data]

    @staticmethod
    def return_id(data: list[dict[str, Any]], name: str) -> int | None:
        id_name_mapping = {
            "host": "hostid",
            "hostgroup": "groupid",
            "template": "templateid",
            "templategroup": "groupid",
            "proxy": "proxyid",
            "item": "itemid",
            "mediatype": "mediatypeid",
            "action": "actionid",
            "usermacro": "hostmacroids",
        }
        if data and len(data) > 0:
            return next(int(item[id_name_mapping[name]]) for item in data)

    @staticmethod
    def id_trans_to_list(ids: list[int] | int):
        return ids if isinstance(ids, list) else [ids]


class _Host(ZbxBase[schemas.HostCreate, schemas.HostGet, schemas.HostUpdate]):
    async def massupdate(self, client: AsyncClient, data: schemas.HostMassUpdate) -> dict[str, Any] | list[int] | None:
        params = data.model_dump(exclude_unset=True)
        response = await self.rpc(method=self.obj_name + ".massupdate", params=params, client=client)
        if response:
            return response


class _HostGroup(ZbxBase[schemas.HostGroupCreate, schemas.HostGroupGet, schemas.HostGroupUpdate]):
    ...


class _HostInterface(ZbxBase[schemas.HostInterfaceCreate, schemas.HostInterfaceGet, schemas.HostInterfaceUpdate]):
    ...


class _Template(ZbxBase[schemas.TemplateCreate, schemas.TemplateGet, schemas.TemplateUpdate]):
    ...


class _TemplateGroup(ZbxBase[schemas.TemplateGroupCreate, schemas.TemplateGroupGet, schemas.TemplateGroupUpdate]):
    ...


class _Proxy(ZbxBase[schemas.ProxyCreate, schemas.ProxyGet, schemas.ProxyUpdate]):
    ...


class _Item(ZbxBase[schemas.ItemCreate, schemas.ItemGet, schemas.ItemUpdate]):
    ...


class _UserMacro(ZbxBase[schemas.UserMacroCreate, schemas.UserMacroGet, schemas.UserMacroUpdate]):
    ...


class _History(ZbxBase):
    async def create(self, client: AsyncClient, data: Any):
        raise NotImplementedError

    async def update(self, client: AsyncClient, data: Any):
        raise NotImplementedError

    async def delete(self, client: AsyncClient, ids: list[int] | int):
        raise NotImplementedError


class _MediaType(ZbxBase[schemas.MediaTypeCreate, schemas.MediaTypeGet, schemas.MediaTypeUpdate]):
    ...


class _Action(ZbxBase[schemas.ActionCreate, schemas.ActionGet, schemas.ActionUpdate]):
    ...


class _Configuration(ZbxBase):
    async def create(self, client: AsyncClient, data: Any):
        raise NotImplementedError

    async def get(self, client: AsyncClient, data: Any):
        raise NotImplementedError

    async def update(self, client: AsyncClient, data: Any):
        raise NotImplementedError

    async def delete(self, client: AsyncClient, ids: list[int] | int):
        raise NotImplementedError

    async def import_(self, client: AsyncClient, source: str):
        params = {
            "format": "yaml",
            "rules": {
                "templates": {"createMissing": True, "updateExisting": True},
                "items": {"createMissing": True, "updateExisting": True, "deleteMissing": True},
                "triggers": {"createMissing": True, "updateExisting": True, "deleteMissing": True},
                "valueMaps": {"createMissing": True, "updateExisting": True},
                "discoveryRules": {"createMissing": True, "updateExisting": True, "deleteMissing": True},
                "graphs": {"createMissing": True, "updateExisting": True, "deleteMissing": True},
                "template_groups": {"createMissing": True, "updateExisting": True},
            },
            "source": source,
        }

        response = await self.rpc(method=self.obj_name + ".import", params=params, client=client)
        return response


class _Event(ZbxBase[schemas.EventCreate, schemas.EventGet, schemas.EventUpdate]):
    async def create(self, client: AsyncClient, data: Any) -> None:
        raise NotImplementedError

    async def update(self, client: AsyncClient, data: Any) -> int | None:
        raise NotImplementedError


def get_zbx_api(token: str) -> "Client":
    return Client(settings.ZBX_URL, token)


zbx_api = Client(settings.ZBX_URL, settings.ZBX_TOKEN)
