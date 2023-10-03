FROM golang:1.20 AS build

WORKDIR /statuz/
COPY ../../go.mod ../../go.sum ./

RUN mkdir /statuz/utils/
COPY ../../utils /statuz/utils/

RUN mkdir -p /statuz/svc/genie
COPY ./svc/genie /statuz/svc/genie

RUN go build -o /statuz/bin/genie -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v ./svc/genie

ENTRYPOINT ["/statuz/bin/genie"]
