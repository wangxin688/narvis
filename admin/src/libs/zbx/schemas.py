from enum import IntEnum
from typing import Any, Literal

from pydantic import BaseModel, Field, field_validator, model_validator


class Search(BaseModel):
    key_: str


class Details(BaseModel):
    version: int = Field(default=2, description="current only support snmp v2c")
    bulk: int = 1
    community: str = Field(description="SNMP community")


class Interface(BaseModel):
    type: int | None = Field(default=2, description="1 - agent, 2 - SNMP, 3 - IPMI, 4 - JMX")
    main: int | None = Field(default=1, description="0 - not default, 1 - default")
    useip: int | None = Field(default=1, description="0 - use DNS, 1 - use IP")
    ip: str | None = Field(default="", description="IP address")
    dns: str | None = Field(default="")
    port: int | None = Field(
        default=161, description="port number used by interface, default snmp is 161, default agent is 10050"
    )
    details: Details | None = Field(default=None, description="only need when type is 2")


class HostInterfaceGet(BaseModel):
    hostids: int | list[int] | None = Field(
        default=None, description="Return only hosts that belong to the given host IDs."
    )


class HostInterfaceCreate(BaseModel): ...


class HostInterfaceUpdate(BaseModel):
    interfaceid: int
    type: int | None = Field(default=2, description="1 - agent, 2 - SNMP, 3 - IPMI, 4 - JMX")
    ip: str | None = None
    port: int | None = Field(
        default=161, description="port number used by interface, default snmp is 161, default agent is 10050"
    )
    details: Details | None = None


class HostId(BaseModel):
    hostid: int


class GroupId(BaseModel):
    groupid: int


class TemplateID(BaseModel):
    templateid: int


class Tag(BaseModel):
    tag: str
    value: Any = None


class TagQuery(Tag):
    operator: int = Field(
        default=0,
        description="""Possible operator values:0 - (default) Contains;
                    1 - Equals;
                    2 - Not like;
                    3 - Not equal;
                    4 - Exists;
                    5 - Not exists.""",
    )


class Macro(BaseModel):
    macro: str
    value: Any = None
    description: str | None = None

    @field_validator("macro")
    @classmethod
    def validate_macro(cls, field_value: str) -> str:
        if field_value.startswith("{$") and field_value.endswith("}"):
            return field_value
        raise ValueError("macro must in {$}")


class HostCreate(BaseModel):
    host: str
    interfaces: list[Interface]
    groups: list[GroupId]
    tags: list[Tag] | None = None
    macros: list[Macro] | None = None
    proxy_hostid: int | None = None
    status: int = Field(default=0, description="0: enable, 1: disable")
    templates: list[TemplateID] | None
    inventory_mode: int = Field(default=1, description="disable zabbix inventory function")


class HostGet(BaseModel):
    groupids: int | list[int] | None = Field(
        default=None, description="Return only hosts that belong to the given groups."
    )
    hostids: int | list[int] | None = Field(
        default=None, description="Return only hosts that belong to the given host IDs."
    )
    proxyids: int | list[int] | None = Field(
        default=None, description="Return only hosts that are linked to the given proxies."
    )
    templateids: int | list[int] | None = Field(
        default=None, description="Return only hosts that are linked to the given templates."
    )
    tags: TagQuery | None = None
    filter: dict | None = Field(default=None, description="eg: {'filter': 'host': ['zabbix server', 'linux server']}")
    output: Literal["extend", "shorten"] | list[str] | None = None


class HostUpdate(BaseModel):
    hostid: int
    host: str | None = None
    groups: list[GroupId] | None = Field(
        default=None, description="replace the current host groups the host belongs to"
    )
    interfaces: list[HostInterfaceUpdate] | None = Field(
        default=None, description=" replace the current host interfaces."
    )
    tags: list[Tag] | None = None
    macros: list[Macro] | None = None
    templates: list[TemplateID] | None = Field(default=None, description="templates replace but not clear")
    templates_clear: list[TemplateID] | None = Field(default=None, description="unlink and clear")
    status: int | None = Field(default=None, description="0: enable, 1: disable")
    proxy_hostid: int | None = None


class HostMassUpdate(BaseModel):
    hosts: list[HostId]
    groups: list[GroupId] | None = None
    status: int | None = None
    proxy_hostid: int | None = None


class HostGroupCreate(BaseModel):
    name: str


class HostGroupGet(BaseModel):
    groupids: int | list[int] | None = Field(
        default=None, description="Return only host groups with the given host group IDs."
    )
    hostids: int | list[int] | None = Field(default=None, description="Return only host groups with the given hosts.")
    filter: dict | None = Field(default=None, description="eg: {'filter': 'name': ['zabbix servers', 'linux servers']}")
    output: Literal["extend", "shorten"] | list[int] | None = None


class HostGroupUpdate(BaseModel):
    groupid: int
    name: str


class TemplateCreate(BaseModel):
    host: str
    groups: list[GroupId] = Field(description="template groups to add the template to")
    templates: list[TemplateID] | None = Field(default=None, description="templates to be linked to the template")
    tags: list[Tag] | None = None
    macros: list[Macro] | None = None


class TemplateGet(BaseModel):
    templateids: list[int] | int | None = Field(
        default=None, description="Return only templates with the given template IDs."
    )
    groupids: list[int] | int | None = Field(
        default=None, description="Return only templates that belong to the given template groups."
    )
    hostids: list[int] | int | None = Field(
        default=None, description="Return only templates that are linked to the given hosts/templates."
    )
    filter: dict | None = None
    output: Literal["extend", "shorten"] | list[str] | None = None


class TemplateUpdate(BaseModel):
    templateid: int
    host: str | None = None
    templates: list[TemplateID] | None = Field(default=None, description="replace templates linked to")
    groups: list[GroupId] | None = None


class TemplateGroupCreate(BaseModel):
    name: str


class TemplateGroupGet(BaseModel):
    groupids: list[int] | int | None = None
    templateids: list[int] | int | None = None
    filter: dict | None = None
    output: Literal["extend", "shorten"] | list[str] | None = None


class TemplateGroupUpdate(BaseModel):
    groupid: int
    name: str | None = None


class ProxyCreate(BaseModel):
    host: str
    status: Literal[5, 6] = Field(default=5, description="5: active; 6: passive")
    tls_connect: int = Field(default=1, description="1: No encryption.2: PSK.3: certificate")
    tls_psk_identity: str = Field(description="proxy config, can keep it as host")
    tls_accept: int = Field(default=2, description="1: No encryption.2: PSK.3: certificate")
    tls_psk: str = Field(description="pre share key")


class ProxyGet(BaseModel):
    proxyids: list[int] | int | None = Field(default=None, description="Return only proxies with the given proxy IDs.")
    selectHosts: bool = Field(  # noqa: N815
        default=False, description="Return a hosts property with the hosts monitored by the proxy."
    )

    @model_validator(mode="after")
    def validate_selecthosts(self):
        self.selectHosts = "extend" if self.selectHosts else None  # type: ignore # ! "selectHosts": value must be "extend".
        return self


class ProxyUpdate(BaseModel):
    proxyid: int
    host: str | None = None
    hosts: list[HostId] | None = Field(
        default=None,
        description="Hosts to be monitored by the proxy. If a host is already monitored by a different proxy, it will be reassigned to the current proxy.",
    )
    tls_psk_identity: str | None = Field(default=None, description="proxy config, can keep it as host")
    tls_psk: str | None = Field(default=None, description="pre share key")


class ItemCreate(BaseModel): ...


class ItemGet(BaseModel):
    templateids: list[int] | int | None = Field(
        default=None, description="Return only items that belong to the given templates."
    )
    hostids: list[int] | int | None = Field(
        default=None, description="Return only items that belong to the given hosts."
    )
    triggerids: list[int] | int | None = Field(
        default=None, description="Return only items that belong to the given triggers."
    )
    search: Search | None = None


class ItemUpdate(BaseModel): ...


class HistoryGet(BaseModel):
    history: int = Field(
        default=3,
        ge=0,
        le=4,
        description="0 - numeric float;, 1 - character; 2 - log;3 - (default) numeric unsigned; 4 - text",
    )
    itemids: list[int] | int | None = None
    time_from: int
    time_till: int
    limit: int | None = None
    output: Literal["extend", "shorten"] | list[str] | None = None


class MediaTypeCreate(BaseModel):
    pass


class MediaTypeGet(BaseModel):
    mediatypeids: list[int] | int | None = None
    mediaids: list[int] | int | None = None


class MediaTypeUpdate(BaseModel):
    pass


class ActionCreate(BaseModel): ...


class ActionGet(BaseModel):
    pass


class ActionUpdate(BaseModel):
    pass


class InterfaceGet(BaseModel):
    hostids: list[int] | int


class InterfaceUpdate(Interface): ...


class UserMacroCreate(BaseModel):
    hostid: int
    macro: str
    value: str


class UserMacroUpdate(BaseModel):
    hostmacroid: int
    value: str


class UserMacroGet(BaseModel):
    hostids: list[int] | int


class EventCreate(BaseModel): ...


class EventUpdate(BaseModel): ...


class EventGet(BaseModel):
    eventids: list[int] | int


class ZbxStatus(IntEnum):
    enable = 0
    disable = 1
