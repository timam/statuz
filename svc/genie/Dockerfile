FROM golang:1.21 AS build

WORKDIR /statuz/
COPY ../../go.mod ../../go.sum ./

RUN mkdir /statuz/utils/
COPY ../../utils /statuz/utils/

RUN mkdir -p /statuz/svc/genie
COPY ./cmd/genie /statuz/cmd/genie

RUN go build -o /statuz/bin/genie -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v ./svc/genie

FROM alpine:edge
COPY --from=build /statuz/bin/genie /statuz/bin/genie
ENTRYPOINT ["/statuz/bin/genie"]
