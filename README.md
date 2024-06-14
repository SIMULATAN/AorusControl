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

### On Arch Linux
If you are on Arch Linux, you can use the AUR package.
It is called [aoruscontrol-bin](https://aur.archlinux.org/packages/aoruscontrol-bin/).

Use your favorite AUR helper to install it, e.g.:
```bash
# using yay
yay -S aoruscontrol-bin
# using paru
paru -S aoruscontrol-bin
```

## Usage
```bash
# if you installed from source using `go install`
aoruscontrol # make sure $GOPATH/bin is in your $PATH
# otherwise
./AorusControl
```
Make sure to run the tool as root, as it needs to access `/sys/kernel/debug/ec/ec0/io`.
