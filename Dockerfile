FROM golang:1.9-alpine
MAINTAINER Danniel Magno

EXPOSE 8080

WORKDIR /go/src/github.com/DennyLoko/go-thumbor

ENTRYPOINT [ "go-thumbor" ]
CMD [ "--help" ]

COPY . /go/src/github.com/DennyLoko/go-thumbor

RUN apk add --update \
        ca-certificates \
        gcc \
        git \
        make \
        musl-dev \
        pcre \
        tzdata \
    && go get -u -v github.com/kardianos/govendor \
    && govendor sync \
    && go install

VOLUME [ "/var/www/img" ]
