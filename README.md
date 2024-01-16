# Azareal

Azareal is a command-line tool for managing schedules for teachers. It allows you to register teachers for classes and subjects, and map subjects to classes.

## Installation

To install Azareal, clone the repository and build the project using the Makefile or Go compiler:

```bash
git clone https://github.com/LitFill/azareal.git
cd azareal
```

### Makefile (recomended)

#### Linux

```bash
make build
```

#### Windows

```bash
make win
```

### Go compiler

```bash
go build -o azareal

go build -o azareal.exe
```

## Usage

After installation, you can run Azareal from the command line:

```bash
./azareal
```

## Some Commands

### guru

Prints the teacher's data based on the provided code.

```bash
azareal guru [teacher code]
```

### version

Prints the version number of Azareal.

```bash
azareal version
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the terms of the MIT license. See the LICENSE file for details.
