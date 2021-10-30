How to execute main_test.go

```bash
# create busybox container
$ docker run --name busybox -d busybox sleep 3600 

# copy shell script
$ docker cp main.sh busybox:/main.sh

# change mode shell script
$ docker exec busybox sh -c 'chmod 755 /main.sh'

# run main_test.go
$ go test main_test.go
```