# refacgo
A Go-based command-line tool that evaluates the code in a specified Go file and provides refactoring suggestions powered by AI.

## Installation
```shell
$ go install github.com/kakky/refacgo
```

## Usage
```
refacgo <command> [options] <filepath>
```

## Commands
Available commands:
- `eval`  - Evaluates the specified file only.
- `refactor` - Evaluates and refactors the specified file.

### `eval`
```
$ refacgo eval [option] <filepath>
```
This command evaluates the file specified as an argument (provide the relative path from the current directory). The default language is English, but you can use a flag to get a eval in Japanese.

In addition, having a description of the code written in the specified file will allow for a more accurate review. It is recommended that this option be written when business concepts are important.

| Option        | Default | Description |
|-------------|---------|-------------|
| `--japanese`<br>`-j` | English  | Can change language to Japanese |
| `--desc` <br> `-d` |  -  | Write a description of the code in the file as an argument to the `--desc` flag|


### `refactor`
This command evaluates and refactors the file specified as an argument (provide the relative path from the current directory). When the command is executed, the specified file is temporarily modified. You can decide whether to apply the changes based on a `y/n` prompt. If `n` is chosen, the file is reverted to its original state. If `y` is selected, the file will be saved.

The temporarily refactored code will show differences using `+` and `-` symbols, making the changes easy to spot.
Once the changes are confirmed, the `+` and `-` symbols will be removed, and the file will be saved.

| Option        | Default | Description |
|-------------|---------|-------------|
| `--japanese`<br>`-j` | English  | Can change language to Japanese |
| `--desc` <br> `-d` |  -  | Write a description of the code in the file as an argument to the `--desc` flag|

