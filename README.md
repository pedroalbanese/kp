# KP
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/kp/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/kp/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/kp/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/kp?status.png)](http://godoc.org/github.com/pedroalbanese/kp)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/kp)](https://goreportcard.com/report/github.com/pedroalbanese/kp)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/kp)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/kp)](https://github.com/pedroalbanese/kp/releases)

This project is a reimplementation of [kpcli](http://kpcli.sourceforge.net/) with a few additional features thrown in.  It provides a shell-like interface for navigating a KeePass V1 database and manipulating entries. 

## Usage

### Command-line
```
> ./kp -help
Usage of ./kp:
  -db string
        the db to open
  -key string
        a key file to use to unlock the db
  -n string
        execute a given command and exit
  -version
        print version and exit
```

### Program Shell
```
/ > help

Commands:
  attach       attach <get|show|delete> <entry> <filesystem location>
  cd           cd <path>
  clear        clear the screen
  edit         edit <entry>
  exit         exit the program
  help         display help
  ls           ls [path]
  mkdir        mkdir <group name>
  mv           mv <source> <destination>
  new          new <path>
  pwd          pwd
  rm           rm <entry>
  save         save
  saveas       saveas <file.kdb> [file.key]
  search       search <term>
  select       select [-f] <entry>
  show         show [-f] <entry>
  version      version
  xp           xp <entry>
  xu           xu
  xw           xw
  xx           xx
```
Running a command with the argument `help` will display a more detailed usage message
```
/ > attach help

manages the attachment for a given entry

Commands:
  create       attach create <entry> <name> <filesystem location>
  details      attach details <entry>
  get          attach get <entry> <filesystem location>
```

## Overview
There are two main components, the shell and the libraries that interact with the database directly.  The shell interfaces with the database through those abstractions.  The shell works by having individual files for each command which are strung together in `main.go`.

## License

This project is licensed under the ISC License.
#### Copyright (c) 2020-2021 ALBANESE Research Lab.
