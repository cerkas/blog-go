FROM golang:1.8-alpine

RUN apk add git --no-cache && \
    apk add pkgconfig --no-cache && \
    apk add g++ --no-cache && \
    go get -u github.com/kardianos/govendor&& \
    go get -u github.com/go-martini/martini&& \
    go get -u  github.com/lib/pq
ADD . /go/src/blog-go

RUN cd /go/src/blog-go && \
    go build -o /go/bin/blog-go && \
    apk del git

WORKDIR /go/src/blog-go

CMD /go/bin/blog-go


EXPOSE 3000
EXPOSE 9100
EXPOSE 3000
EXPOSE 5432