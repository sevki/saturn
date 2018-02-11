FROM golang:1.10

RUN apt-get update -y

RUN apt-get install git texlive -y
RUN apt-get install xz-utils -y
RUN apt-get install graphviz -y

RUN go get \
 cloud.google.com/go/compute/metadata \
 golang.org/x/net/html \
 gopkg.in/russross/blackfriday.v2 \
 github.com/pkg/errors \
 upspin.io/...

ADD https://github.com/jgm/pandoc/releases/download/2.1.2/pandoc-2.1.2-1-amd64.deb /pandoc-2.1.2-1-amd64.deb
RUN dpkg -i /pandoc-2.1.2-1-amd64.deb

ADD http://cdsoft.fr/pp/pp-linux-x86_64.txz /pp-2.3.3.txz
RUN tar -xf /pp-2.3.3.txz -C /
RUN mv /pp /usr/bin/pp
ADD . /go/src/sevki.org/saturn

RUN which pp

RUN go get -v sevki.org/saturn/cmd/saturn

ENTRYPOINT saturn

EXPOSE 8080
