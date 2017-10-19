FROM golang

EXPOSE 50112

ADD ./downloadPic /go/src/github.com/YAWAL/downloadPic
ADD ./api /go/src/github.com/YAWAL/gRPCtrying/api
RUN mkdir -p home/vya/Pictures
COPY ./pics home/vya/Pictures
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
RUN go install github.com/YAWAL/gRPCtrying/downloadPic

ENTRYPOINT ["/go/bin/downloadPic"]