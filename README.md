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

`Go` **only** requires semicolons when two or more statements or declarations appear on the same line 

The location of `newlines` in Go **matter**

`gofmt` tool rewrites code into the standard format. `goimports` is another helpful tool

Go does not permit unused local variables which would result in a compiliation error

# Installing packages 
`go install golang.org/x/tools/cmd/goimports@latest`.

`go get` is no longer supported and the version, `@latest` must be explicitly stated

# Best Practices
Describe each package with a comment before the package declartion.
`main` packages have comments describing the entire program as a whole. 

Use either of the two forms to initialize a variable.
> `s: = ""`
> `var s string` 
The first form is when the inital value matters more, the second is when the data type matter more

# Useful Tips
`:=` is a _short variable declaration_ which allows for the declaration of one or more variable and the assignment of the corresponding type based on the initializer value. 

`i++` is a numeric increment statement 

`i--` is a numeric decrement statement

# Looping in Go
The `for` loop os the only loop statement in **Go** with various forms but with a standard structure and the opening brace must be on the same line as the post.
~~~
for initialization; condtion; post {
    //statements here
}
~~~
The `initialization` is optional and is executed before the loop starts. It must be a _simple statement_ being one of the following
>A _short variable declaration_ `i := 1`
>An increment or assignment statemnet 
>A function call 

The `condition` is a boolean expression which is evaluated at the beginning of each loop iteration 

The `post` is executed after the body of the loop then the `condition` is evaluated again. 

The `loop` ends when the condition becomes false. 

`initialization`, `condition`, `post` are all option for a `loop`

A traditional `while` loop can be constructed as follows
~~~
for condition {
    // .. 
}
~~~
A traditional `infinite` loop can be expressed as 
~~~
for {
    // ..
}
~~~
The above can be escaped with a `break` or `return` statement


# Command-Line Arguments
- The `os` package provides functions and other values for platform-independent OS interactions. 
- `Command-line` arguments are available to the program in a variable called `Args` that is part of the `os` package.
>* Use `os.args` to access the variable outside of the `os` package 
>* The variable `os.args` is a _slice_ of strings. **Go** _Slices_ will be discussed later but are functionally like a **Python** _List_ with some differences.
>* The first element of `os.args`, `os.args[0]`, is the name of the command itself. All other elements are the arguments that were present at the start of execution. 
>* The typical desired _slice_ is `os.args[1:]`
