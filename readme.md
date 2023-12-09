# Iggy
A non-standard way of combining `.ignore` files.

## Usage
### Create `.iggy` File
`iggy init`

### Combine `.ignore` Files
`iggy combine`

### Build `.ignore` Files
`iggy build`

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
The above will generate two `.ignore` files; `.gitignore` and `.dockerignore`. Technically, anything between the curly braces will be considered part of the `.ignore` file. So that includes the comments and any other plaintext.

