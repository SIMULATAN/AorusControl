# AorusControl
A TUI to control your Aorus Motherboard / Laptop Fan settings.

## Features
- Monitor current fan speeds
- See current fan mode
- Set fan speed

## Screenshot
![screenshot of the TUI](screenshot.png)

## Installation
### From source (using go)
```bash
go install github.com/simulatan/aoruscontrol@main
```
If you have go, this is the easiest way to install the tool.

Updating is just as easy:
```bash
GOPROXY=direct go install github.com/simulatan/aoruscontrol@main
```

### From GitHub Releases
Download the latest release from the [releases page](https://github.com/SIMULATAN/AorusControl/releases).

Extract the `tar.gz` file and run the binary inside.

## Usage
```bash
# from source
aoruscontrol # make sure $GOPATH/bin is in your $PATH
# from release
./AorusControl
```
Make sure to run the tool as root, as it needs to access `/sys/kernel/debug/ec/ec0/io`.
