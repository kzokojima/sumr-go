# sumr-go

calculate a message-digest fingerprint (checksum) for a file on recursive

## Requirement

* Go 1.11

## Installation

```
$ go get github.com/kzokojima/sumr-go/sumr
```

## Usage

```
$ sumr [-a algo] [dir]
```

### Options

* -a algo<br>
Name of selected hashing algorithm (i.e. "md5" (default), "sha256")

### Output example

```
PATH TAB SUM (ALGO)
./path/to/file1 TAB XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
./path/to/file2 TAB XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

* sumr ignore special files ('desktop.ini', 'Thumbs.db', '.DS_Store').

## Test

```
$ make test
```

## Changes

### v0.1.1 - 2019-03-06

* fix

### v0.1.0 - 2019-03-05

* Initial release

## Author

[Kazuo Kojima](https://github.com/kzokojima)

## License

[MIT](LICENSE)
