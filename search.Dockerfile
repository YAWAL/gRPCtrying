FROM golang

EXPOSE 50111

ADD ./search /go/src/github.com/YAWAL/gRPCtrying/search
ADD ./api /go/src/github.com/YAWAL/gRPCtrying/api
ADD ./glide.yaml /go/src/github.com/YAWAL/gRPCtrying

RUN mkdir -p /home/vya/Pictures
COPY ./pics /home/vya/Pictures

RUN go get github.com/Masterminds/glide
WORKDIR /go/src/github.com/YAWAL/gRPCtrying
RUN glide install
RUN go install github.com/YAWAL/gRPCtrying/search

ENTRYPOINT ["/go/bin/search"]