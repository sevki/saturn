FROM sevki/saturn:base as builder

ADD . /sevki.org/saturn

WORKDIR /sevki.org/saturn

RUN go install -mod vendor -v ./cmd/...

FROM sevki/saturn:base

ADD templates /sevki.org/saturn/templates

COPY --from=builder /go/bin/saturn /go/bin/saturn
COPY --from=builder /go/bin/atlas /go/bin/atlas


RUN mkdir /x
RUN mkdir /pub

ENTRYPOINT saturn

EXPOSE	8080