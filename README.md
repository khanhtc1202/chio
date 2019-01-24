# Chio

[![][goreportcard-svg]][goreportcard] 
[![][CodeFactor]](https://www.codefactor.io/repository/github/khanhtc1202/chio)
[![][Build Status]](https://travis-ci.org/khanhtc1202/chio)

[Build Status]: https://travis-ci.org/khanhtc1202/chio.svg?branch=master
[CodeFactor]: https://www.codefactor.io/repository/github/khanhtc1202/chio/badge
[goreportcard]: https://goreportcard.com/report/github.com/khanhtc1202/chio
[goreportcard-svg]: https://goreportcard.com/badge/github.com/khanhtc1202/chio

Chio (chan) exports metrics of a specific module for checking it's stability and flexibility.  

## What can it does

Main idea from [Clean Architecture book](https://books.google.co.jp/books/about/Clean_Architecture.html?id=uGE1DwAAQBAJ&source=kp_cover&redir_esc=y) Chapter 14. 

Summarization go [here](https://github.com/khanhtc1202/til/issues/7).

Three metrics we care about:

1. Abstractness (A):
This metric has the range [0, 1]. A value of 0 implies that the module has no abstract classes at all. A value of 1 implies that the module contains nothing but abstract classes.
2. Instability (I):
This metric has the range [0, 1]. I = 0 indicates a maximally stable module. I = 1 indicates a maximally unstable module.
3. Distance (D):
There is a place, where A and I have a balance, it's called `The Main Sequence`.
This metric has the range [0, 1]. D = 0 implies that the module lies in that place, and on the other side (D = 1) means far from balance line, locates in `Zone of Pain ` or `Zone of Uselessness`.  

![](https://user-images.githubusercontent.com/32532742/43695838-e64283f0-9975-11e8-8a9d-8d6d64f87437.png)

Example output metrics extracted from [boogeyman repo](https://github.com/khanhtc1202/boogeyman)

```bash
+---------------------------------+-------+----------+----------+-------+--------+--------------+-------------+----------+
|           MODULE PATH           | FILES | CONCRETE | ABSTRACT | FANIN | FANOUT | ABSTRACTNESS | INSTABILITY | DISTANCE |
+---------------------------------+-------+----------+----------+-------+--------+--------------+-------------+----------+
| /infrastructure/cmd/data/       |     1 |        1 |        0 |     1 |      0 |        0.000 |       0.000 |    1.000 |
| /infrastructure/meta_info/      |     1 |        1 |        0 |     0 |      1 |        0.000 |       1.000 |    0.000 |
| /usecase/presenter/             |     1 |        0 |        1 |     2 |      1 |        1.000 |       0.333 |    0.333 |
| /adapter/persistent/repository/ |     2 |        1 |        0 |     0 |      3 |        0.000 |       1.000 |    0.000 |
| /domain/                        |    12 |        4 |        0 |    11 |     12 |        0.000 |       0.522 |    0.478 |
| /infrastructure/io/             |     1 |        1 |        0 |     4 |      3 |        0.000 |       0.429 |    0.571 |
| /usecase/interactor/            |     2 |        1 |        0 |     1 |      5 |        0.000 |       0.833 |    0.167 |
| /usecase/repository/            |     1 |        0 |        1 |     4 |      1 |        1.000 |       0.200 |    0.200 |
| /adapter/presenter/console/     |     1 |        1 |        0 |     0 |      4 |        0.000 |       1.000 |    0.000 |
| /adapter/persistent/service/    |     1 |        0 |        1 |     1 |      1 |        1.000 |       0.500 |    0.500 |
| /config/                        |     5 |        5 |        0 |     1 |     13 |        0.000 |       0.929 |    0.071 |
| /cross_cutting/io/              |     2 |        0 |        1 |     4 |      2 |        1.000 |       0.333 |    0.333 |
| /infrastructure/service/        |     8 |        4 |        0 |     0 |     26 |        0.000 |       1.000 |    0.000 |
| /adapter/controller/            |     1 |        1 |        0 |     0 |      4 |        0.000 |       1.000 |    0.000 |
| /infrastructure/cmd/            |     1 |        1 |        0 |     0 |      2 |        0.000 |       1.000 |    0.000 |
+---------------------------------+-------+----------+----------+-------+--------+--------------+-------------+----------+
```

## Usage

In case exec file you downloaded's name is `chio`.

Sample full params command

```bash
$ ./chio -l go -p ./
```

Type `-h` to get help. Return value be like

```$xslt
Usage of ./chio:
  -l string
    language(s): go (default "go")
  -p string
    path to module (default ".")
```


## TODO list

Support language(s)
- [x] Golang
- [ ] Java
- [ ] NodeJs

Load module strategies
- [x] n-depth Directory level (group files in same dir as a module)
- [ ] 1-depth Directory level (file as module)
- [ ] File level (file as module)

## Run on local

Chio development environment requires: 

1. Golang (1.9.2 or higher). Install go [here](https://golang.org/doc/install).
2. dep (Go dependency management tool). Install go [here](https://github.com/golang/dep).

Run by `go`

```bash
$ go run main.go
```

or check [Makefile](https://github.com/khanhtc1202/chio/blob/master/Makefile) for building bin on your local.

## Contribution

All contributions will be welcome in this project.

## Download

For linux x64 : [download](https://github.com/khanhtc1202/chio/releases/download/0.0.1/chio-linux-64)

For MacOS x64 : [download](https://github.com/khanhtc1202/chio/releases/download/0.0.1/chio-darwin-64)

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
