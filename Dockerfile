FROM golang:1.21 AS build

WORKDIR /src/
COPY . /src/

RUN go build -o /bin/statuz -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v ./cmd/api

FROM alpine:edge
COPY --from=build /bin/statuz /bin/statuz
ENTRYPOINT ["/bin/statuz"]
