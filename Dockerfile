# Build PKI Server
FROM golang:1
ENV version 1.0.0
RUN mkdir -p $GOPATH/src/github.com/srleyva
WORKDIR /go/src/github.com/srleyva
ADD . ./CertAPI
RUN  cd CertAPI && CGO_ENABLED=0 GOOS=linux go build main.go && cp main /main

# Inject PKI Server into container
FROM scratch
COPY config.yaml /etc/.pkiconf/config.yaml
COPY --from=0 /main /bin/main
ENTRYPOINT ["main"]