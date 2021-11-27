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