FROM debian:jessie
MAINTAINER jeanepaul@gmail.com

# USAGE
# $ docker build -t weather101 .
# $ docker run -d --name weather101 -v "$(pwd)":/go/src/weather101 -p 49167:8000 weather101

# SCMs for "go get", gcc for cgo
RUN apt-get update && apt-get install -y \
		ca-certificates curl gcc libc6-dev make \
		bzr git mercurial \
		openssh-client \
		--no-install-recommends \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.4.2

RUN curl -sSL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz \
		| tar -v -C /usr/src -xz

RUN cd /usr/src/go/src && ./make.bash --no-clean 2>&1

ENV PATH /usr/src/go/bin:$PATH

RUN mkdir -p /go/src/weather101
ENV GOPATH /go
ENV PATH /go/bin:$PATH

WORKDIR /go/src/weather101

EXPOSE 8000
ENTRYPOINT ls -lh && go get ./... && go build -v && ls -lh && ./weather101
