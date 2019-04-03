# KPM Proxy
KPM Proxy is a dirver for web applications to access small text files on the local hard drive.
It runs as a service, preferable listening on `localhost`.

# How does it work
KPM Proxy offers an endpoint to write files to disk, and has available a websocket to send text files to web browser.

# Environment variables
Example env variables and their default values.
```bash
KPMPROXY_TCP_BIND="127.0.0.1:8080"
KPMPROXY_LONGPOLL="5"
KPMROXY_SHORTPOLL="1"
KPMPROXY_DIRS=""
```
`KPMPROXY_DIRS` variable is a list of directories separated with a colon (`:`) as in:
```bash
KPMPROXY_DIRS="/home/diego/tmp/stuff:/home/diego/tmp/stuff2"
```
This is the only mandatory variable.

# Write
To write files make an HTTP POST to `http://127.0.0.1:8080/kpmproxy/write` 
```json
{
    "body": "Hello, World!",
    "dir": "/home/diego/tmp/stuff",
    "name": "hi.txt"
}
```
An status 200 OK response means writing was a success.

# Read
For an example on how to read files, see `websockets.html` file.

# Build
Clone this repo and run `go build .`

