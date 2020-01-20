FROM golang as build

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags='-s -w -extldflags "-static"' -i .

FROM scratch

COPY --from=build /build/go-sysmon /go-sysmon
ENTRYPOINT [ "/go-sysmon" ]
