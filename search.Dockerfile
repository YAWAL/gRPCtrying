FROM golang

EXPOSE 50111

ADD ./searchPic /go/src/github.com/YAWAL/gRPCtrying/searchPic
ADD ./api /go/src/github.com/YAWAL/gRPCtrying/api
RUN mkdir -p home/vya/Pictures
COPY ./pics home/vya/Pictures
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
RUN go install github.com/YAWAL/gRPCtrying/searchPic

ENTRYPOINT ["/go/bin/searchPic"]