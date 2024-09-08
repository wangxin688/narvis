import json

from httpx import AsyncClient


class RabbitMQ:
    def __init__(self, rabbit_url: str, rabbit_admin_user: str, rabbit_admin_password: str) -> None:
        self._rabbit_url = rabbit_url
        self._rabbit_admin_user = rabbit_admin_user
        self._rabbit_admin_password = rabbit_admin_password

    async def get_vhost(self, client: AsyncClient) -> dict:
        result = await client.get(
            url=f"{self._rabbit_url}/api/vhosts",
            headers={"Content-Type": "application/json"},
            auth=(self._rabbit_admin_user, self._rabbit_admin_password),
        )
        return result.json()

    async def create_vhost(self, client: AsyncClient, vhost: str) -> int:
        result = await client.put(
            url=f"{self._rabbit_url}/api/vhosts/{vhost}",
            headers={"Content-Type": "application/json"},
            auth=(self._rabbit_admin_user, self._rabbit_admin_password),
        )
        return result.status_code

    async def create_user(self, client: AsyncClient, username: str, password: str, tags: str = "") -> int:
        result = await client.put(
            url=f"{self._rabbit_url}/api/users/{username}",
            headers={"Content-Type": "application/json"},
            auth=(self._rabbit_admin_user, self._rabbit_admin_password),
            data=json.dumps({"username": username, "password": password, "tags": tags}),
        )
        return result.status_code

    async def get_users(self, client: AsyncClient) -> dict:
        result = await client.get(
            url=f"{self._rabbit_url}/api/users/",
            headers={"Content-Type": "application/json"},
            auth=(self._rabbit_admin_user, self._rabbit_admin_password),
        )
        return result.json()

    async def create_vhost_permission(self, client: AsyncClient, vhost: str, username: str) -> int:
        result = await client.put(
            url=f"{self._rabbit_url}/api/permissions/{vhost}/{username}",
            headers={"Content-Type": "application/json"},
            auth=(self._rabbit_admin_user, self._rabbit_admin_password),
            data=json.dumps({"username": username, "vhost": vhost, "configure": ".*", "write": ".*", "read": ".*"}),
        )
        return result.status_code
