# HDD Proxy
HDD Proxy is a dirver for web applications to access small text files on the local hard drive.
It runs as a service, preferable listening on `localhost`.

# How does it work
HDD Proxy offers an endpoint to write files to disk, and has available a websocket to send text files to web browser.

# Environment variables
Example env variables and their default values.
```bash
HDDPROXY_TCP_BIND="127.0.0.1:8080"
HDDPROXY_LONGPOLL="5"
HDDROXY_SHORTPOLL="1"
HDDPROXY_DIRS_SEP=":"
HDDPROXY_DIRS=""
```
`HDDPROXY_DIRS` variable is a list of directories separated (by default) with a colon (`:`) as in:
```bash
HDDPROXY_DIRS="/home/diego/tmp/stuff:/home/diego/tmp/stuff2"
```
This is the only mandatory variable.
Directories separator (`:`) can be changed setting `HDDPROXY_DIRS_SEP` variable.

# Write
To write files make an HTTP POST to `http://127.0.0.1:8080/hddproxy/write` 
```json
{
    "body": "Hello, World!",
    "dir": "/home/diego/tmp/stuff",
    "name": "hi.txt"
}
```
A response with status 200 OK means writing was successful.

# Read
For an example on how to read files, see `websockets.html` file.

# Build
Clone this repo and run `go build .`

