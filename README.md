# :sunglasses: Feynman :nerd_face:

[![Publish Binaries](https://github.com/raven4ever/change-proxy-app/actions/workflows/release.yml/badge.svg)](https://github.com/raven4ever/change-proxy-app/actions/workflows/release.yml)

Feynman is a tool to find and replace the proxy URL in a given list of files on both Windows and Linux systems.

The tool will read the files to be changed from the `config.yml` file.

By default, the `config.yml` file is located in the same directory as the `feynman` binary. This can be changed by passing the `-config` flag to the binary which will take the path to the `config.yml` file as an argument.

The `config.yml` file is a YAML file with the following structure:

```yaml
username: Domain\username
password: Su#per@S3cr3t!&

proxy_url: https://proxy.url:8080

files:
  - path: C:\Users\Adrian\Documents\Projects\change-proxy-app\testfile.txt
    variables:
      - http_proxy
      - https_proxy
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

All the files listed in the `config.yml` file will be searched for the variables listed under the `variables` key. If a variable is found, it will be replaced with the value of the `proxy_url` key. If the variable is not found, it will be added to end of the file with the value of the `proxy_url` key.

The `proxy_url` value that will be used to replace the variables will contain the username and password from the `username` and `password` keys in a URL encoded format (`https://username:password@url.com:8080`). The special characters in the `username` and `password` keys will be translated to their [URL encoded format](https://www.w3schools.com/tags/ref_urlencode.ASP).

## Installation

To install and use Feynman across your system you need to:

- from the [releases](https://github.com/raven4ever/change-proxy-app/releases) page, download the binary for your operating system and architecture.
- rename the binary to `feynman` and make it executable.
- add the binary to your `PATH` environment variable.

## Usage

**_!!!Before running the binary, make sure you have sufficient permissions to change the files listed in the `config.yml` file!!!_**

- If your `config.yml` file is located in the same location as the binary, just run the binary without any arguments:

  ```bash
  feynman
  ```

- If your `config.yml` file is located in a different directory, you can pass the `-config` flag to the binary:

  ```bash
  feynman -config /path/to/config.yml
  ```

## Note

The tool is written in both Go and Python.

Both versions are functionally identical and they will be released together.

On the [releases](https://github.com/raven4ever/change-proxy-app/releases) page you will find the binaries for both versions under the form `feynman-go-vX.Y.Z-os-arch` and `feynman-py-vX.Y.Z-os-arch`.
