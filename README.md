# expert-umbrella
An umbrella project for Golang 

# Driving Force 
Learning Golang with applied applications 

# General Notes 
`main` is special. As a package, it deifnes a standalone executable program 
and not a library. The function `main` is where the execution of the `main` 
program begins. 

`import` declaration is import as one must import the exact needed packages. 
A program will *NOT* compile if there are missing or unnecessary imports. 

`package` declaration *->* `import` declaration *->* then everything else

`func`: Functions require the keyword `func`, the name of the function, a parameter list, 
a body, and a result list 

`Go` *only* requires semicolons when two or more statements or declarations appear on the same line 

The location of `newlines` in Go *matter*

`gofmt` tool rewrites code into the standard format. `goimports` is another helpful tool

# Installing packages 
`go install golang.org/x/tools/cmd/goimports@latest` 
`go get` is no longer supported and the version must be explicitly stated

# Command-Line Arguments
- The `os` package provides functions and other values for platform-independent OS interactions. 
- `Command-line` arguments are available to the program in a variable called `Args` that is part of the `os` package.
>* Use `os.args` to access the variable outside of the `os` package 
>* The variable `os.args` is a _slice_ of strings. **Go** _Slices_ will be discussed later but are functionally like a **Python** _List_ with some differences.
>* The first element of `os.args`, `os.args[0]`, is the name of the command itself. All other elements are the arguments that were present at the start of execution. 
>* The typical desired _slice_ is `os.args[1:]`
