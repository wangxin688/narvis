import cryptocode

from src.core.config import settings

def token_decrypt(x_auth_token: str, secret_key: str = settings.SECRET_KEY) -> str:
    result = cryptocode.decrypt(x_auth_token, secret_key)
    if not result:
        return ""
    return result


def token_encrypt(token: str, secret_key: str = settings.SECRET_KEY) -> str:
    return cryptocode.encrypt(token, secret_key)