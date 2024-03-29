# I/O over networking

Low level networking remains a lot reading and writing a file. We can even use
readers and writes, just as seen in [0005-todo-list-file][1] and
[0004-check-triangles][2].

For this one i decided to use a full feature go module and test cases because
I/O has interesting caveats that crosses languages and hit the modern OS'es
boundaries.

## The 'sleep' 

Network connections are limited OS resources. They take time to get provisioned
and therefore test it must take the OS par into account when testing it.

This is why we have this line on the testcase:

```go
time.Sleep(time.Duration(1) * time.Second)
```

A similar line would be needed in other languages as well, like Java, C/C++ and
so on.

[1]: ../0005-todo-list-file/todo-list-v1.go
[2]: ../0004-check-triangles/check-triangles.go

## How to run this example

```bash
go test -v -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```
