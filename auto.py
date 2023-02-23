#!env/bin/python

import os
import random
import string
import sys


def generate_key(length: int) -> str:
    return (''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length)))


def put_this_instead_that(this: str, that: str, texts: list, prefix: str) -> str:
    output = ""
    found = False
    for text in texts:
        if that.strip() in text:
            found = True
            text = prefix + this + "\n"
        output += text

    if not found:
        output += "\n" + prefix + this + "\n"

    return output


def read_file(file_address: str) -> list:
    file = open(file_address, 'r')
    content = file.readlines()
    file.close()

    return content


def write_file(file_address: str, content: str) -> None:
    file = open(file_address, 'w')
    file.write(content)
    file.close()


def env_exists() -> bool:
    return os.path.isfile('env.yml') or os.path.isfile('env.yaml')


def env_file_address() -> str:
    if os.path.isfile('env.yml'):
        return "env.yml"
    elif os.path.isfile('env.yaml'):
        return "env.yaml"


def generate_secret_key(do_print: bool) -> None:
    key = generate_key(64)
    if do_print:
        print(key)

    if not env_exists():
        return
    file_address = env_file_address()

    write_file(file_address, put_this_instead_that(
        f"\"{key}\"", "secret_key", read_file(file_address), "secret_key: "))


def install_dependencies() -> None:
    print("===Installing Dependencies===\n")
    if os.system("go mod download -x all") != 0:
        print("installing dependencies failed")
        return

    print("===Done===\n")

    print("===Setup Config Files===\n")
    if os.system("cp dbconfig_example.yml dbconfig.yml && cp env_example.yml env.yml") != 0:
        print("dbconfig_example.yml or env_example.yml does not exists")
        return

    print("===Done===\n")


def create_python_env() -> None:
    print("===Creating env Folder===\n")
    os.system("python3 -m venv env")
    print("===Done===\n")


def activate_githooks() -> None:
    print("===Activating Githooks===\n")
    os.system("python3 .githooks/install.py")
    print("===Done===\n")


def main() -> None:
    if sys.argv[1].lower() == "generate":
        generate_secret_key(True)

    elif sys.argv[1].lower() == "setup":
        install_dependencies()
        generate_secret_key(False)
        create_python_env()
        activate_githooks()


if __name__ == "__main__":
    main()
