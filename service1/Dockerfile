FROM golang:1.14 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o upstream-app main.go

# runtime image
FROM alpine

RUN mkdir -p /data/pdf/handout
COPY --from=buildImage /usr/src/app/upstream-app /
ENTRYPOINT ["/upstream-app"]