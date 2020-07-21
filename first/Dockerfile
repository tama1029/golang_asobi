FROM golang:1.14 AS build-env
ADD . /go/src/github.com/tama1029/golang_asobi/first
WORKDIR /go/src/github.com/tama1029/golang_asobi/first
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' main.go

FROM alpine
RUN apk update
COPY --from=build-env /go/src/github.com/tama1029/golang_asobi/first/main /usr/local/bin/main
EXPOSE 1323
CMD /usr/local/bin/main