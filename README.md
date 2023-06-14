# Feynman

Feynman is a tool to change the proxy URL on a Linux system.

The tool will read the files to be changed from the `config.yml` file. The `config.yml` file is located in the same directory as the `feynman` binary. The `config.yml` file is a YAML file with the following structure:

```yaml
credentials:
  username: "username"
  password: "password"

http_proxy_url: "http://proxy.url:port"
https_proxy_url: "https://proxy.url:port"

files:
  - /etc/environment
    variables:
        - http_proxy
        - https_proxy
        - proxy
        - qualys_http_proxy
  - /etc/wgetrc
    variables:
        - http_proxy
        - https_proxy
  - /home/username/.bashrc
    variables:
        - http_proxy
        - https_proxy
        - proxy
        - qualys_http_proxy
        - all_proxy
  - /home/username/.zshrc
    variables:
        - http_proxy
        - https_proxy
        - proxy
        - qualys_http_proxy
        - all_proxy
```
