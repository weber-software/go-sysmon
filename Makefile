NAME=go-sysmon

build-local:
	CGO_ENABLED=0 go build -a -tags netgo -ldflags='-s -w -extldflags "-static"' -i .

build-docker:
	#go build -i .
	docker run --privileged -it --rm -e GOPATH=$(PWD)/deps -e GOCACHE=$(PWD)/cache -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 -v $(PWD):$(PWD) --workdir $(PWD) golang go build -a -tags netgo -ldflags='-s -w -extldflags "-static"' -i .

build-arm:
	docker run --privileged -it --rm -e GOPATH=$(PWD)/deps -e GOCACHE=$(PWD)/cache -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=arm -e GOARM=5 -v $(PWD):$(PWD) --workdir $(PWD) golang go build -a tags netgo -ldflags='-s -w -extldflags "-static"' -i .

run:
	SYSMON_INFLUX_TAG=sys1 SYSMON_INFLUX_HOST=192.168.1.60 ./go-sysmon

run-arm:
	#scp $(NAME) root@192.168.1.141:/tmp/
	#ssh root@192.168.1.141 "/tmp/$(NAME) < /sys/kernel/debug/tracing/trace_pipe"
