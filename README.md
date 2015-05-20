# Example app vendored using `godep save`

* The use of `godep save` (without the `-r` option) does not rewrite import paths.
* Application can only be built using the `godep build` and not `go build`

# To build this app:

```
$ mkdir godepsave
$ cd godepsave/
$ export GOPATH=`pwd`
$ export GOBIN=$GOPATH/bin
$ go get github.com/kr/godep
$ go install github.com/kr/godep
$ go get github.com/jbuberel/example_godep_save
$ cd src/github.com/jbuberel/example_godep_save/godepserver/
$ godep go build 
```
