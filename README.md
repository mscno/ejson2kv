# ejson2kv

Ejson2kv is a simple tool to convert ejson files to key-value pairs.

## Installation

```bash
go install github.com/mscno/ejson2kv@latest
```

## Usage

```bash
ejson decrypt secrets.ejson | ejson2kv
```

will tranform the ejson file to key-value pairs.

## Example

```bash
$ cat secrets.ejson
{
    "_public_key": "kP1YIjY...",
    "password": "password",
    "username": "username"
    
}

$ ejson decrypt secrets.ejson | ejson2kv
password=password
username=username
```