# shelleditor
Invoke the shell `EDITOR` like `kubectl edit`.

This code is lifted from [kubectl](https://github.com/kubernetes/kubectl/tree/master/pkg/cmd/util/editor).
Thanks lads. 

Why not import the `kubectl` code directly? It has tons of dependencies that are not needed
for this simple task. The codebase has been edited to import fewer packages,
and those that are imported are mostly copied to the `/pkg` dir.


# Example program

There's an example program in `cmd/shelleditor`.

```shell
$ go install github.com/neilotoole/shelleditor/cmd/shelleditor
$ shelleditor hello.txt
```
