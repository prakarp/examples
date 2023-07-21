# Important points to note

1. import statement, refers to a directory. It does not refer to a package. 
   A directory can have only one package in golang. In the code, always 
   refer to the package name, when trying to use the structs and functions
   in any of the source files of the package.

   The directory name and the package name can be the same and it makes sense
   to keep that convention in all the projects. But golang allows those two
   identifiers to be different, as exemplified here.


2. There is no need to set GOROOT, GOPATH or GO111MODULE variables. In fact,
   do not use them if you want to go with the defaults for your platform.

   On linux, GOROOT is /usr/local/go and GOPATH is ~/go.

   So, the module names (e.g. go.mod) and the import statements refer to the
   directory relative to $GOROOT/src.  


# Directory Structure


   In this example, the directory structure of the medapp program is
   medapp 

```
     README.md
     go.mod
     cmd 
       +- main.go  (imports the directory outpatient)
     pkg
       +- outpatient
            +- labtest.go  (the package name is outpatient_pkg)
   
```

main.go however, refers to a struct and some functions in labtest.go
by using outpatient_pkg package name, and not the directory name.


The directory name and the package name are different in this example
just to illustrate a point, but the convention is that they are the same.

In any case, you can import a package under a different name by giving it
your own name in the import statement. This is sometimes necessary to avoid
namespace collisions.


# For programmers starting with Go

If you are familiar with other programming langauges and are beginning in Go,
the example here should be useful to start with something beyond hello_world.go
(which you need, to ensure your go installation works as expected).

On linux, keep all your source files under ~/go/src, but also note the following.

Typically, you are going to be using third party packages from other code
repositories like github.com or gitlab.com. To work with them seamlessly, 
keep the directory structure to match their code as well as yours.

As an exmaple, you should have
$HOME/go/src/github.com/<github username>/<mmy program directory>


# Modules

Always use modules. A package is a collection of go files and should be used for
encapsulation principles. A module is a collection of related packages with the
version numbers, and also refers to third party packages with their version 
numbers (for linking to your main program, as well as the tools used to build 
your program).  

Your go program source code, does not refer to modules of other programmers. 
It only refers to their packages. 

But your go.mod file can refer to other modules (of other programmers) by
specifying a specific version as a dependency (require statement).

In a short summary, a module is for dependency management of the packages, and 
a package is the source code itself.


