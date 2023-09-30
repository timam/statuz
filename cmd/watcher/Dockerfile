FROM golang:1.21 AS build

WORKDIR /statuz/
COPY ../../go.mod ../../go.sum ./

RUN mkdir /statuz/utils/
COPY ../../utils /statuz/utils/

RUN mkdir /statuz/config/
COPY ../../config /statuz/config/

RUN mkdir -p /statuz/cmd/watcher
COPY ./cmd/watcher /statuz/cmd/watcher

RUN go build -o /statuz/bin/watcher -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v ./cmd/watcher

FROM alpine:edge
COPY --from=build /statuz/bin/watcher /statuz/bin/watcher
ENTRYPOINT ["/statuz/bin/watcher"]
