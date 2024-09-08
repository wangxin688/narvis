import asyncio

from src.bootstrap.init_zbx import init_monitor_system
from src.bootstrap.init_rmq import init_rabbit_mq


async def init_app()->None:
    await init_rabbit_mq()
    await init_monitor_system()


if __name__ == "__main__":
    asyncio.run(init_app())