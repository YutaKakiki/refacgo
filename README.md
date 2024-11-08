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
- `review`  - Evaluates the specified file only.
- `refactor` - Evaluates and refactors the specified file.

### `review`
```
$ refacgo review [option] <filepath>
```
This command reviews the file specified as an argument (provide the relative path from the current directory). The default language is English, but you can use a flag to get a review in Japanese.

| Option        | Default | Description |
|-------------|---------|-------------|
| `-j`, `-en` | `-en`   | Can select language, <br> English or Japanese |

### `refactor`
This command evaluates and refactors the file specified as an argument (provide the relative path from the current directory). When the command is executed, the specified file is temporarily modified. You can decide whether to apply the changes based on a `y/n` prompt. If `n` is chosen, the file is reverted to its original state. If `y` is selected, the file will be saved.

The temporarily refactored code will show differences using `+` and `-` symbols, making the changes easy to spot.
Once the changes are confirmed, the `+` and `-` symbols will be removed, and the file will be saved.

| Option        | Default | Description |
|-------------|---------|-------------|
| `-j`, `-en` | `-en`   | Can select language, <br> English or Japanese |

