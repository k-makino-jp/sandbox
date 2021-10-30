How to execute main_test.go

```bash
# create busybox container
$ docker run --name busybox -d busybox sleep 3600 

# run main_test.go
$ go test main_test.go
```