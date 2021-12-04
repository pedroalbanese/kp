# KP
This project is a reimplementation of [kpcli](http://kpcli.sourceforge.net/) with a few additional features thrown in.  It provides a shell-like interface for navigating a keepass database and manipulating entries.

Currently, a full reimplementation of all of kpcli's features is still under development, but other features, such as `search` have been added.

keepassv2+ has also been implemented, but is still very incomplete - treat it as alpha software.  It is also advisable to not use v2 at all outside of development as it's incomplete and may have edge cases where corruption can occur.

keepassv1, however, is stable and fit for use
# Usage

## Command-line
```
> ./kp -help
Usage of ./kp:
  -db string
        the db to open
  -key string
        a key file to use to unlock the db
  -kpversion int
        which version of keepass to use (1 or 2) (default 1)
  -n string
        execute a given command and exit
  -version
        print version and exit
```

## Program Shell
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
  mv           mv <soruce> <destination>
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
There are two main components, the shell and the libraries that interact with the database directly.  The shell interfaces with the database through those abstractionsso that the actual logic is the same for v1 and v2.  The shell works by having individual files for each command which are strung together in `main.go`.

## License

This project is licensed under the ISC License.
#### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Research Lab.
