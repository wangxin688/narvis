from pathlib import Path

from anyio import open_file
from httpx import AsyncClient

from src.core.config import settings,  PROJECT_DIR
from src.core.security import token_encrypt
from src.libs.zbx.client import get_zbx_api

async def init_monitor_system() -> None:
    async with AsyncClient() as client:
        url = settings.ZBX_URL + "/api_jsonrpc.php"
        login = await client.post(
            url=url,
            headers={"Content-Type": "application/json-rpc"},
            json={
                "jsonrpc": "2.0",
                "method": "user.login",
                "params": {"username": "Admin", "password": "zabbix"},
                "id": 1,
            },
        )
        assert login.status_code == 200
        if "error" not in login.json():
            token = login.json()["result"]
            headers = {"Content-Type": "application/json-rpc", "Authorization": "Bearer " + token}
            basic_params = [
                {"name": "alertName", "value": "{TRIGGER.NAME}"},
                {"name": "HostId", "value": "{HOSTNAME}"},
                {"name": "itemName", "value": "{ITEM.NAME}"},
                {"name": "opData", "value": "{EVENT.OPDATA}"},
                {"name": "status", "value": "{TRIGGER.STATUS}"},
                {"name": "triggerId", "value": "{TRIGGER.ID}"},
                {"name": "eventId", "value": "{EVENT.ID}"},
                {"name": "token", "value": "{$XAUTHTOKEN}"},
                {"name": "tags", "value": "{EVENT.TAGSJSON}"},
            ]
            if settings.CURRENT_ENV in ("TEST", "DEV"):
                _url = "http://host.docker.internal:8000"
            else:
                _url = settings.BASE_URL
            alert_media_params = basic_params + [{"name": "url", "value": f"{_url}/api/assurance/alert/alerts"}]
            event_media_params = basic_params + [{"name": "url", "value": f"{_url}/api/assurance/alert/events"}]
            alert_media_type = {
                "type": "4",
                "name": "Narvis Alerts",
                "status": "0",
                "script": (
                    "try {\r\n    var result = {},"
                    "\r\n    params = JSON.parse(value),\r\n    req = new HttpRequest(),\r\n    fields = {},"
                    '\r\n    resp;\r\n\r\n    req.addHeader("Content-Type: application/json");\r\n'
                    '  req.addHeader("Authorization: Bearer "+params.token);  \r\n    '
                    "fields.alertName = params.alertName;\r\n    fields.HostId = params.HostId;\r\n    "
                    "fields.itemName = params.itemName;\r\n    fields.opData = params.opData;\r\n    "
                    "fields.status = params.status;\r\n    fields.triggerId = params.triggerId;\r\n    fields.url "
                    "= params.url;\r\n      fields.eventId = "
                    "params.eventId;\r\n    fields.tags = params.tags;\r\n    for (var i in fields) {\r\n  "
                    "      if (fields[i].startsWith('{') )\r\n        {\r\n            delete fields[i]\r\n        "
                    "}\r\n    }\r\n    if (JSON.stringify(fields) === '{}') {\r\n        result = {}\r\n    } else "
                    "{\r\n        resp = req.post(fields.url, JSON.stringify(fields));\r\n        Zabbix.log("
                    "'fields: '+JSON.stringify(fields))\r\n        if (req.getStatus() != 200) {\r\n            "
                    "throw 'Response Code: ' + req.getStatus();\r\n        }\r\n        resp = JSON.parse("
                    "resp);\r\n        result.code = resp.code;\r\n        result.data = resp.data;\r\n        "
                    "result.msg = resp.msg;\r\n    }\r\n    \r\n}   catch (error) {\r\n    Zabbix.log(4, "
                    "'send alert to narvis alert failed: eventId: '+JSON.stringify(fields.eventId));\r\n    "
                    "result = {};\r\n}\r\nreturn JSON.stringify(result);"
                ),
                "parameters": alert_media_params,
            }
            event_media_type = {
                "type": "4",
                "name": "Narvis Events",
                "status": "0",
                "script": (
                    "try {\r\n    var result = {},"
                    "\r\n    params = JSON.parse(value),\r\n    req = new HttpRequest(),\r\n    fields = {},"
                    '\r\n    resp;\r\n\r\n    req.addHeader("Content-Type: application/json");\r\n '
                    'req.addHeader("Authorization: Bearer "+params.token);   \r\n'
                    "fields.alertName = params.alertName;\r\n    fields.HostId = params.HostId;\r\n    "
                    "fields.itemName = params.itemName;\r\n    fields.opData = params.opData;\r\n    "
                    "fields.status = params.status;\r\n    fields.triggerId = params.triggerId;\r\n    fields.url "
                    "= params.url;\r\n      fields.eventId = "
                    "params.eventId;\r\n    fields.tags = params.tags;\r\n    for (var i in fields) {\r\n  "
                    "      if (fields[i].startsWith('{') )\r\n        {\r\n            delete fields[i]\r\n        "
                    "}\r\n    }\r\n    if (JSON.stringify(fields) === '{}') {\r\n        result = {}\r\n    } else "
                    "{\r\n        resp = req.post(fields.url, JSON.stringify(fields));\r\n        Zabbix.log("
                    "'fields: '+JSON.stringify(fields))\r\n        if (req.getStatus() != 200) {\r\n            "
                    "throw 'Response Code: ' + req.getStatus();\r\n        }\r\n        resp = JSON.parse("
                    "resp);\r\n        result.code = resp.code;\r\n        result.data = resp.data;\r\n        "
                    "result.msg = resp.msg;\r\n    }\r\n    \r\n}   catch (error) {\r\n    Zabbix.log(4, "
                    "'send alert to narvis alert failed: eventId: '+JSON.stringify(fields.eventId));\r\n    "
                    "result = {};\r\n}\r\nreturn JSON.stringify(result);"
                ),
                "parameters": event_media_params,
            }
            alert_media_type_result = await client.post(
                url=url,
                headers=headers,
                json={"jsonrpc": "2.0", "method": "mediatype.create", "params": alert_media_type, "id": 1},
            )
            assert alert_media_type_result.status_code == 200
            new_global_macro = await client.post(
                url=url,
                headers=headers,
                json={
                    "jsonrpc": "2.0",
                    "method": "usermacro.createglobal",
                    "params": {
                        "macro": "{$XAUTHTOKEN}",
                        "value": token_encrypt(settings.PUBLIC_AUTH_KEY),
                        "type": 1,
                    },
                    "id": 1,
                },
            )
            assert new_global_macro.status_code == 200
            if "error" not in new_global_macro.json():
                print("Init app: Create GlobalMacro: {$XAUTHTOKEN} successfully!")
            if "error" not in alert_media_type_result.json():
                alert_media_type_id = alert_media_type_result.json()["result"]["mediatypeids"][0]
                print("Init app: Create Monitoring MeditType: Narvis Alerts successfully!")
            else:
                print("Init app: Create Monitoring MediaType: Narvis Alerts already created!")
            event_media_type_result = await client.post(
                url=url,
                headers=headers,
                json={"jsonrpc": "2.0", "method": "mediatype.create", "params": event_media_type, "id": 1},
            )
            assert event_media_type_result.status_code == 200
            if "error" not in event_media_type_result.json():
                event_media_type_id: int = event_media_type_result.json()["result"]["mediatypeids"][0]
                print("Create Monitoring MeditType: Narvis Events successfully!")
                new_superuser = {
                    "username": settings.ZBX_USERNAME,
                    "passwd": settings.ZBX_PASSWORD,
                    "roleid": "3",
                    "usrgrps": [{"usrgrpid": "7"}],
                    "medias": [
                        {
                            "mediatypeid": alert_media_type_id,
                            "sendto": ["narvis"],
                            "active": 0,
                            "severity": 63,
                            "period": "1-7,00:00-24:00",
                        },
                        {
                            "mediatypeid": event_media_type_id,
                            "sendto": ["narvis"],
                            "active": 0,
                            "severity": 63,
                            "period": "1-7,00:00-24:00",
                        },
                    ],
                }
                new_superuser_result = await client.post(
                    url=url,
                    headers=headers,
                    json={"jsonrpc": "2.0", "method": "user.create", "params": new_superuser, "id": 1},
                )
                assert new_superuser_result.status_code == 200
                if "error" not in new_superuser_result.json():
                    new_super_user_id = new_superuser_result.json()["result"]["userids"][0]
                    print("Init app: Create Monitoring superuser: narvis successfully!")
                    new_token = {"name": "narvis token", "userid": new_super_user_id}
                    new_token_result = await client.post(
                        url=url,
                        headers=headers,
                        json={"jsonrpc": "2.0", "method": "token.create", "params": new_token, "id": 1},
                    )
                    assert new_token_result.status_code == 200
                    if "error" not in new_token_result.json():
                        new_action = {
                            "name": "NetCare Alert",
                            "eventsource": "0",
                            "status": "0",
                            "esc_period": "1h",
                            "pause_suppressed": "1",
                            "notify_if_canceled": "1",
                            "pause_symptoms": "1",
                            "filter": {
                                "evaltype": "0",
                                "conditions": [
                                    {
                                        "conditiontype": "16",
                                        "operator": "11",
                                    },
                                    {"conditiontype": "0", "operator": "1", "value": "4"},
                                ],
                            },
                            "operations": [
                                {
                                    "operationtype": "0",
                                    "esc_period": "20m",
                                    "esc_step_from": "1",
                                    "esc_step_to": "0",
                                    "evaltype": "0",
                                    "opconditions": [{"conditiontype": "14", "operator": "0", "value": "0"}],
                                    "opmessage": {
                                        "default_msg": "0",
                                        "subject": "Problem: {EVENT.NAME}",
                                        "message": "",
                                        "mediatypeid": alert_media_type_id,
                                    },
                                    "opmessage_grp": [],
                                    "opmessage_usr": [
                                        {
                                            "userid": new_super_user_id,
                                        }
                                    ],
                                }
                            ],
                            "recovery_operations": [
                                {
                                    "operationtype": "0",
                                    "opmessage": {
                                        "default_msg": "0",
                                        "subject": "Resolved: {EVENT.NAME}",
                                        "message": "",
                                        "mediatypeid": alert_media_type_id,
                                    },
                                    "opmessage_grp": [],
                                    "opmessage_usr": [{"userid": new_super_user_id}],
                                }
                            ],
                        }
                        new_action_result = await client.post(
                            url=url,
                            headers=headers,
                            json={"jsonrpc": "2.0", "method": "action.create", "params": new_action, "id": 1},
                        )
                        assert new_action_result.status_code == 200
                        print(new_action_result.json())
                        new_token_id = new_token_result.json()["result"]["tokenids"][0]
                        new_token_generate = await client.post(
                            url=url,
                            headers=headers,
                            json={"jsonrpc": "2.0", "method": "token.generate", "params": [new_token_id], "id": 1},
                        )
                        assert new_token_generate.status_code == 200
                        if "error" not in new_token_generate.json():
                            new_user_token = new_token_generate.json()["result"][0]["token"]
                            print(f"init monitoring server: token: {new_user_token}")
                            async with await open_file(file=f"{PROJECT_DIR}/.env", mode="r+", encoding="utf-8") as f:
                                contents = await f.readlines()
                                for line_num, line in enumerate(contents):
                                    if "ZBX_TOKEN" in line:
                                        prefix = line.split("=")[0]
                                        contents[line_num] = prefix + "=" + new_user_token + "\n"
                                        await f.seek(0)
                                        await f.writelines(contents)
                                        break
                            print("New Token: " + new_user_token)
                            disable_admin = await client.post(
                                url=url,
                                headers={
                                    "Content-Type": "application/json-rpc",
                                    "Authorization": "Bearer " + new_user_token,
                                },
                                json={
                                    "jsonrpc": "2.0",
                                    "method": "user.update",
                                    "params": {
                                        "userid": "1",
                                        "usrgrps": [{"usrgrpid": "9"}],
                                    },
                                    "id": 1,
                                },
                            )
                            assert disable_admin.status_code == 200
                            if "error" not in disable_admin.json():
                                print("Init app: Disable Admin user successfully!")
                            await import_monitor_template(new_user_token)
                        print("Init app: Create Monitoring token successfully!")
                else:
                    print("Init app: Create Monitoring superuser: narvis already created!")
            else:
                print("Init app: Create Monitoring MediaType: Narvis Events already created!")
        else:
            token = settings.ZBX_TOKEN
            await import_monitor_template(token)
            print("Init app: Monitoring template update successfully!")




async def import_monitor_template(token: str) -> None:
    dir_path = Path(f"{PROJECT_DIR}/src/bootstrap/templates")
    async with AsyncClient() as client:
        for path in dir_path.iterdir():
            if not path.is_file():
                continue
            filename = path.name
            if not filename.endswith(".yaml"):
                continue
            content = path.read_text()
            zbx_api = get_zbx_api(token)
            await zbx_api.configuration.import_(client, content)
            print(f"Init app: import monitor template {filename} successfully!")