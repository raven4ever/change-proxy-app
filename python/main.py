#!/usr/bin/env python3

import argparse
from pathlib import Path
import urllib.parse
import yaml


def file_path(string) -> str:
    if Path(string).is_file():
        return string
    else:
        raise ValueError(
            f'config file path: {string} does not exist or is not a file.')


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


def edit_file(file: dict, proxy_url: str) -> None:
    file_path = file['path']
    file_vars = file['variables']

    # check if the file exists
    if not Path(file_path).is_file():
        pass

    # using regex, find the lines that start with the variable name and replace the value with the proxy_url value
    # if the variable is not found, add the variable to the end of the file
    with open(file_path, 'r+') as f:
        file_lines = f.readlines()

        for var in file_vars:
            # check if the variable exists
            var_exists = False
            for i, line in enumerate(file_lines):
                if line.startswith(var):
                    var_exists = True
                    file_lines[i] = f'{var}={proxy_url}\n'
                    break

            # if the variable does not exist, add it to the end of the file
            if not var_exists:
                file_lines.append(f'{var}={proxy_url}\n')

        f.seek(0)
        f.writelines(file_lines)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Search & Replace Proxy URL.')

    # add arguments
    parser.add_argument('-config', type=file_path, help='Config file path.')

    # parse the arguments
    args = parser.parse_args()

    CONFIG_FILE_PATH = getattr(args, 'config', './config.yml')

    # load config file from the config.yml path
    with open(CONFIG_FILE_PATH, 'r') as f:
        config = yaml.load(f, Loader=yaml.FullLoader)

    # insert credentials into proxy url
    proxy_url = insert_credentials_into_proxy_url(config)

    print(proxy_url)

    # edit each file in the config file and add the variables defined in the config file
    for file in config['files']:
        edit_file(file, proxy_url)
