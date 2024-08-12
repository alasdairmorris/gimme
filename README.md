# gimme

A super-simple file downloader.

## Installation

`gimme` will run on most Linux, MacOS and Windows systems.

To install it, just `cd` into the directory in which you wish to install it and then copy-paste the appropriate one-liner from below.

### Linux (32-bit)

```
curl -s -L -o gimme https://github.com/alasdairmorris/gimme/releases/latest/download/gimme-linux-386 && chmod +x gimme
```

### Linux (64-bit)

```
curl -s -L -o gimme https://github.com/alasdairmorris/gimme/releases/latest/download/gimme-linux-amd64 && chmod +x gimme
```

### Mac OS X (Intel)

```
curl -s -L -o gimme https://github.com/alasdairmorris/gimme/releases/latest/download/gimme-darwin-amd64 && chmod +x gimme
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o gimme https://github.com/alasdairmorris/gimme/releases/latest/download/gimme-darwin-arm64 && chmod +x gimme
```

### Windows (32-bit)

```
curl -s -L -o gimme.exe https://github.com/alasdairmorris/gimme/releases/latest/download/gimme-windows-386.exe
```

### Windows (64-bit)

```
curl -s -L -o gimme.exe https://github.com/alasdairmorris/gimme/releases/latest/download/gimme-windows-amd64.exe
```


### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/gimme@latest
```


## Usage

```
A super-simple file downloader.

Usage:
  gimme -o FILE URL
  gimme -h | --help
  gimme --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
  -o, --outfile FILE     File to download to.

Homepage: https://github.com/alasdairmorris/gimme

```

## Examples

```
$ gimme -o unixtime https://github.com/alasdairmorris/unixtime/releases/latest/download/unixtime-darwin-amd64
```

## License

[MIT](LICENSE)
