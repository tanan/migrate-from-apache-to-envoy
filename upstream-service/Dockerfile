FROM golang:1.14 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o upstream-service main.go

FROM alpine
COPY --from=buildImage /usr/src/app/upstream-service /
ENTRYPOINT ["/upstream-service"]