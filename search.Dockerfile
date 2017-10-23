FROM golang

EXPOSE 50111

ADD ./search /go/src/github.com/YAWAL/gRPCtrying/search
ADD ./api /go/src/github.com/YAWAL/gRPCtrying/api
RUN mkdir -p home/vya/Pictures
COPY ./pics home/vya/Pictures
RUN go install github.com/YAWAL/gRPCtrying/search

ENTRYPOINT ["/go/bin/search"]