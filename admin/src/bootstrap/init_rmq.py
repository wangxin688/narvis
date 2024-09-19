from httpx import AsyncClient

from src.core.config import settings
from src.libs.rabbitmq.rmq import RabbitMQ

async def init_rabbit_mq() -> None:
    async with AsyncClient() as client:
        rq = RabbitMQ(
            rabbit_url=settings.RABBIT_MQ_URL,
            rabbit_admin_user=settings.RABBIT_MQ_SERVER_USER,
            rabbit_admin_password=settings.RABBIT_MQ_SERVER_PASSWORD,
        )
        vhosts = await rq.get_vhost(client)
        if len(vhosts) >= 2:
            print("Init app: Vhost already exists, will not be created anymore.")
        else:
            vhost_server = await rq.create_vhost(client, settings.RABBIT_MQ_SERVER_VHOST)
            assert vhost_server == 201
            vhost_client = await rq.create_vhost(client, settings.RABBIT_MQ_PROXY_VHOST)
            assert vhost_client == 201
            print("Init app: Create vhost successfully. ")
        users = await rq.get_users(client)
        if len(users) >= 2:
            print("Init app: Users already exists, will not be created anymore.")
        else:
            client_user = await rq.create_user(
                client, username=settings.RABBIT_MQ_PROXY_USER, password=settings.RABBIT_MQ_PROXY_PASSWORD
            )
            assert client_user == 201
            server_permission = await rq.create_vhost_permission(
                client, vhost=settings.RABBIT_MQ_SERVER_VHOST, username=settings.RABBIT_MQ_SERVER_USER
            )
            client_permission = await rq.create_vhost_permission(
                client, vhost=settings.RABBIT_MQ_PROXY_VHOST, username=settings.RABBIT_MQ_PROXY_USER
            )

            assert server_permission == 204
            assert client_permission == 201
            print("Init app: Create user and set permission successfully. ")
