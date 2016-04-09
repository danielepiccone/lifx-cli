#  Lifx cli

This is a cli client for lifx wifi light bulbs using the [remote control API](https://api.developer.lifx.com/)

## Installation and usage

Get Golang [here](https://golang.org/doc/install) or ```apt-get install golang-1.6```

### Installation

```
$ go get ./...
$ make
$ ./bin/lifx
...
# to install in /usr/local/bin
$ make install
...
$ make uninstall

```
### Usage

```
NAME:
   Lifx - lifx [command] [...params]

USAGE:
   lifx [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
    config      Gets the configuration
    select      Select a light
    list        List all bulbs
    toggle      Toggle all bulbs
    on          Turn light(s) on
    off         Turn light(s) off
    brightness  Change the brightness of a light
    color       Change the color of a light

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version

```

## Disclaimer

This is a project I did for personal use, not all features from the original API are implemented.

## License (MIT)

Copyright (c) 2016 Daniele Piccone

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.




