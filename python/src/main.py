#!/usr/bin/env python3

import argparse
import logging

import yaml

from filesz import edit_file, file_path
from utils import insert_credentials_into_proxy_url

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)

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

    logging.info(f"Using as proxy URL: {proxy_url}")

    # edit each file in the config file and add the variables defined in the config file
    for file in config['files']:
        edit_file(file, proxy_url)
