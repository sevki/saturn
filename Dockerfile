FROM sevki/saturn:base

ADD . /go/src/sevki.org/saturn

RUN go get -v sevki.org/saturn/cmd/...
RUN mkdir /x # pan/pan.go://x//

ENTRYPOINT saturn

EXPOSE	8080
