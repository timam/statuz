FROM golang:1.21 AS build

WORKDIR /statuz/
COPY ../../go.mod ../../go.sum ./

RUN mkdir /statuz/utils/
COPY ../../utils /statuz/utils/

RUN mkdir -p /statuz/svc/watcher
COPY ./svc/watcher /statuz/svc/watcher

RUN go build -o /statuz/bin/watcher -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v ./svc/watcher

FROM alpine:edge
COPY --from=build /statuz/bin/watcher /statuz/bin/watcher
ENTRYPOINT ["/statuz/bin/watcher"]
