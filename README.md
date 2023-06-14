# Feynman

Feynman is a tool to change the proxy URL on a Linux system.

The tool will read the files to be changed from the `config.yml` file.

By default, the `config.yml` file is located in the same directory as the `feynman` binary. This can be changed by passing the `-config` flag to the binary which will take the path to the `config.yml` file as an argument.

The `config.yml` file is a YAML file with the following structure:

```yaml
credentials:
  username: Domain\username
  password: SuperS3cr3t

http_proxy_url: "http://proxy.url:port"
https_proxy_url: "https://proxy.url:port"

files:
  - path: /etc/environment
    variables:
      - http_proxy
      - https_proxy
      - proxy
      - qualys_http_proxy
  - path: /etc/wgetrc
    variables:
      - http_proxy
      - https_proxy
  - path: /home/username/.bashrc
    variables:
      - http_proxy
      - https_proxy
      - proxy
      - qualys_http_proxy
      - all_proxy
  - path: /home/username/.zshrc
    variables:
      - http_proxy
      - https_proxy
      - proxy
      - qualys_http_proxy
      - all_proxy
```
