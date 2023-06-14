import urllib.parse

def insert_credentials_into_proxy_url(config: dict) -> str:
    proxy_url = config['proxy_url']

    # URL encode username and password
    username = urllib.parse.quote(config['username'])
    password = urllib.parse.quote(config['password'])

    auth_string = f'{username}:{password}@'

    # insert credentials into proxy url
    if proxy_url.startswith('http') or proxy_url.startswith('https'):
        proxy_url = proxy_url.replace('://', f'://{auth_string}')
    else:
        proxy_url = 'https://' + auth_string + proxy_url

    return proxy_url