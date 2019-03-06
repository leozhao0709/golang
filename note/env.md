# Golang Env

## 1. install and setup GOPATH

-   install: `brew install golang`
-   upgrade: `brew upgrade golang`
-   Setup GOPATH: add the following script to your bashfile:

    ```shell
    export GOROOT=/usr/local/Cellar/go/1.11.2/libexec
    export PATH=$PATH:$GOROOT/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```

-   You may need to change `GOROOT` if you update your golang

## 2. [go modules(dependecy management) ](https://www.jianshu.com/p/c5733da150c6)

With go modules, you **don't need** to initial project in `$GOPATH`. Create a project folder in any folder, then:

-   Use `go mod init (github.com/userName/project)` to initial go module
-   Use `go mod download <path@version>` to download modules to local cache, this won't update mod file.
-   Use `go get <path@version>` will also get the depency but will also update the mod file.
-   Use `go list -m` to list current version and dependecy
-   Use just `go mod` to see the help list
-   Use `go tidy` to add missing and remove unused modules

