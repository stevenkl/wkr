# Wkr

Wkr is simple job-runner, inspired by [Workr](https://github.com/sirikon/workr) but written in Go.
It is primary run on Windows, but running on Linux/MacOS is also planed.

## Build

To build you need [Tsk](https://github.com/stevenkl/tsk) installed:

```shell
go get -u github.com/stevenkl/tsk/cmd/tsk
```

After that call `tsk build`, you can find the executable at `./build/wkr(.exe)`.
