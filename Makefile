NAME=go-sysmon

build-local:
	go build -gcflags=-trimpath=$(PWD) -ldflags "-s -w" .

build-docker:
	#go build -i .
	docker run --privileged -it --rm -e GOPATH=$(PWD)/deps -e GOCACHE=$(PWD)/cache -e GOOS=linux -v $(PWD):$(PWD) --workdir $(PWD) golang go build -ldflags="-s -w" -i .

build-arm:
	docker run --privileged -it --rm -e GOPATH=$(PWD)/deps -e GOCACHE=$(PWD)/cache -e GOOS=linux -e GOARCH=arm -e GOARM=5 -v $(PWD):$(PWD) --workdir $(PWD) golang go build -ldflags="-s -w" -i .

run:
	SYSMON_TAG=sys1 SYSMON_INFLUX_HOST=192.168.1.60 ./go-sysmon

run-arm:
	#scp $(NAME) root@192.168.1.141:/tmp/
	#ssh root@192.168.1.141 "/tmp/$(NAME) < /sys/kernel/debug/tracing/trace_pipe"
