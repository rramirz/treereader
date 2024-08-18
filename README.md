
# TreeReader

`TreeReader` is a Go-based command-line tool that recursively opens and reads files in a specified directory while preserving the directory hierarchy. It allows you to selectively ignore specific files or directories and focus on particular files of interest.

## Features

- Recursively reads and prints the contents of files in a directory.
- Allows ignoring specific files or directories (e.g., `.git`).
- Optionally selects specific files to include in the output.
- Skips execution for files with executable permissions.

## Usage

### Basic Usage

To run `TreeReader` with the default configuration (starting from the current directory):

```sh
./treereader
```

### Ignoring Files or Directories

You can ignore specific files or directories by using the `-ignore` flag:

```sh
./treereader -ignore=go.mod,config.yaml
```

In this example, `TreeReader` will skip the `.git` directory and the `config.yaml` file.

### Selecting Specific Files

If you only want to process specific files, use the `-select` flag:

```sh
./treereader -select=README.md,main.go
```

In this example, `TreeReader` will only read the `README.md` and `main.go` files.

### Combining Options

You can combine the `-ignore` and `-select` options to customize the file processing:

```sh
./treereader -ignore=go.mod -select=README.md
```

### Specifying a Root Directory

To start reading from a specific directory, provide the directory path as an argument:

```sh
./treereader /path/to/directory
```

## Installation

1. Clone the repository:
    ```sh
    git clone git@github.com:rramirz/treereader.git
    ```
2. Navigate to the directory:
    ```sh
    cd treereader
    ```
3. Build the application:
    ```sh
    go build -o treereader
    ```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
