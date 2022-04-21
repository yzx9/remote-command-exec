# remote-command-exec

A toy for kidding, DO NOT FOR PRODUCTION.

## Usage

1. Run `go build`
2. Run `hiddenrun.vbs`
3. POST it:

```http
POST 

{
  "command": "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe",
  "args": ["https://google.com/"]
}
```
