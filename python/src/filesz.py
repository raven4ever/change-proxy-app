from pathlib import Path


def file_path(string) -> str:
    if Path(string).is_file():
        return string
    else:
        raise ValueError(
            f'config file path: {string} does not exist or is not a file.')


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
                # check if line is empty
                if not line.strip():
                    continue

                if line.startswith(var):
                    var_exists = True
                    file_lines[i] = f'{var}={proxy_url}\n'
                    break

            # if the variable does not exist, add it to the end of the file
            if not var_exists:
                file_lines.append(f'{var}={proxy_url}\n')

        f.seek(0)
        f.writelines(file_lines)
