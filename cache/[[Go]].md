# Go

# 1 Grammer

## 1.1 Program Structure

    package main //current package name
    
    import "fmt" // import
    
    import (
    	"pkg1"
    	"pkg2"
    )            // group import
    
    func main() {
       fmt.Println("Hello, World!")
    }

## 1.2 Basic Syntax

[https://www.tutorialspoint.com/go/go_program_structure.htm](https://www.tutorialspoint.com/go/go_program_structure.htm)

**Data Type**

- Boolean
- Numeric
- String
- Derived

**Variable**

- `var x int`
- `var x = 10`

**Const**

- `const CONS_VAL = 100`

**Operators**

- `&a` provide address of the variable
- `*a` provide pointer to the variable

**Pointer**

- `var ptr *int`
- `vat a int = 10`
`ptr = &a`

**Structure**

    type Books struct {
       title string
       author string
       subject string
       book_id int
    }

## 1.3 Higher level Syntax

[https://www.tutorialspoint.com/go/go_slice.htm](https://www.tutorialspoint.com/go/go_slice.htm)

# 2 CLI

    go run hello.go

# 3 Project Layout

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## 3.1 Go module management

> official guidelines: [https://golang.org/ref/mod](https://golang.org/ref/mod)

`go.mod` file defines all dependent packages in the project. It is placed at the root directory of the go project.

### 3.1.1 `go mod init [package_name]`

- automatically create go.mod file, indicates the directory is a go package/project

        module github.com/chaoszh/blog-crawler
    
    go 1.17

### 3.1.2 `go get [package_name]`

- download package from go repository and install, cache package.
- modify *go.mod* file to add dependency in current project.
- modify *go.sum* file to describe local cached mod status.

### 3.1.3 `go mod tidy`

- automatically detect all dependencies in source code, automatically use `go get` to download all required packages and modify *go.mod* file

### 3.1.4 `go install`

- detect all required dependencies in *go.mod* file, and install which are not exist in local cache.
- run `go build`.

### 3.1.5 `go clean -modcache`

- clear all cached mod stored in *$GOPATH/pkg/mod*.

### 3.1.6 `go build`

- build executable file.