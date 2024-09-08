import tomllib
from enum import StrEnum
from pathlib import Path
from typing import Literal
from uuid import UUID

from pydantic import Field
from pydantic_settings import BaseSettings, SettingsConfigDict

PROJECT_DIR = Path(__file__).parent.parent.parent
with Path.open(Path(f"{PROJECT_DIR}/pyproject.toml"), "rb") as f:
    PYPROJECT_CONTENT = tomllib.load(f)["project"]


class Settings(BaseSettings):
    SECRET_KEY: str = Field(default="26c78673faa8bbad298ad025e09a4e61")
    PUBLIC_AUTH_KEY: str = Field(default="9fdee1b8a3f18b7f86673938beec96e6")

    CURRENT_ENV: Literal["PROD", "DEV"] = Field(default="DEV")
    BASE_URL: str = Field(default="http://127.0.0.1:8000")

    ZBX_URL: str = Field(default="http://127.0.0.1:8088")
    ZBX_USERNAME: str = Field(default="narvis")
    ZBX_PASSWORD: str = Field(default="50a8c8858b1ddca756db990053830303")
    ZBX_TOKEN: str = Field(...)


    RABBIT_MQ_SERVER_USER: str = Field(default="narvis-server")
    RABBIT_MQ_SERVER_PASSWORD: str = Field(default="26cc7abbea97a17b9f7860ee0dabb051")
    RABBIT_MQ_CLIENT_USER: str = Field(default="narvis-client")
    RABBIT_MQ_CLIENT_PASSWORD: str = Field(default="851b090b967a89f802e72a0baf1d230e")
    RABBIT_MQ_URL: str = Field(default="http://127.0.0.1:15672")
    RABBIT_MQ_SERVER_VHOST: str = Field(default="narvis-server")
    RABBIT_MQ_CLIENT_VHOST: str = Field(default="narvis-client")

    model_config = SettingsConfigDict(
        env_file=f"{PROJECT_DIR}/.env",
        case_sensitive=True,
        extra="ignore"
    )

settings = Settings()