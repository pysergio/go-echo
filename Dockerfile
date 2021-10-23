
FROM golang:1.17.2-alpine 
EXPOSE 8000

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH  

RUN mkdir -p $GOPATH/src/go-echo
ADD . $GOPATH/src/go-echo

WORKDIR $GOPATH/src/go-echo 
RUN go build -o go-echo . 

CMD ["/go/src/go-echo/go-echo"]