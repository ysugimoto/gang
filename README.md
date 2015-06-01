# gang

Gang is a command snippet tool on any architecture.

## Installtion

Get the command:

```
$ go get github.com/ysugimoto/gang/cmd/gang
```

## Usage

Show help:

```
$ gang -h
====================================================
 Gang - A console command snippet management tool
====================================================
Usage:
    gang [subcommand|snippet-name] [paramters...] [option]

Options:
    -h, --help : Show this help

Subcommands:
    mode [ls|peco]                  : Change list mode
    ammo                            : Show list sorted by call times
    kill [snippet-name]             : Remove snippet
    bullet [snippet-name] [command] : Register the snippet
    [snippet-name] [bind,...]       : Run the snippet command
```

- `gang mode [ls|peco]`: change list mode.
  - `ls`: show command list and execute selected number's command.
  - `peco`: select command from `peco`, and execute.
- `ammo`: sort by times whitch executed command, and show list.
- `kill`: purge command snippet
- `bullet`: regist new command sunippet

## Feature

### Support placeholder

Gang supports placeholder snippet like `{:name}`, apply paramter dynamically on execute.

```
# Regist command with placeholder
$ gang bullet sample "echo \"{:word}\" | grep world"
> [gang] Command "sample" Bulleted.

# execute
$ gang sample
> [gang] Execute command needs parameter: echo "{:word}" | grep World
> [gang] Bind Parameter "word" is: World # input "World"
World
```

## License

MIT

## Author

Yoshiaki Sugimoto

## Thanks

This tool was Insprired from:

- <a href="https://github.com/jimmycuadra/bang">jimmycuadra/bang</a>
- <a href="https://github.com/holman/boom">holman/boom</a>
