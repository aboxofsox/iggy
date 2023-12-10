# Iggy
A non-standard way of combining `.ignore` files.

## Install
### With Go
`go install -u github.com/aboxofsox/iggy@latest`

### Without Go
Download the binary for your system, or run the `install` script relevant for your system.

## Usage
```
Combine or create .ignore files from a single .iggy file.

Usage:
  iggy [command]

Available Commands:
  build       Builds ignore files from .iggy file.
  combine     Combine multiple .gitignore files into one
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize a new .iggy file
  update      Check for updates

Flags:
  -h, --help     help for iggy
  -t, --toggle   Help message for toggle

Use "iggy [command] --help" for more information about a command.
```

## The `.iggy` File
The `.iggy` file is a type of `.ignore` file with a defined structure. `.ignore` files are defined by the `@` symbol and the contents of each file is contained between curly braces.

### Example
```
@gitignore{
    # Python
    *.pyc
    *.pyo
    *.pyd
    __pycache__/
    *.so
    *.egg
    *.egg-info/
    dist/
    build/
    *.tar.gz
    
    # IDEs
    .vscode/
    .idea/
    *.swp
    *.swo
    *.swn
    
    # OS
    .DS_Store
    Thumbs.db
}

@dockerignore{
    # Ignore everything
    **

    # Except these files/directories
    !Dockerfile
    !src/	
}
```

The above will generate two `.ignore` files; `.gitignore` and `.dockerignore`. Technically, anything between the curly braces will be considered part of the `.ignore` file. So that includes the comments and any other plaintext. The indention on each line is optional. It's included here for readability. `iggy` will add the indention when you run `iggy combine` or `iggy init`.

