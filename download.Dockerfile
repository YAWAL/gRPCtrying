FROM golang

EXPOSE 50112

ADD ./download /go/src/github.com/YAWAL/gRPCtrying/download
ADD ./api /go/src/github.com/YAWAL/gRPCtrying/api
RUN mkdir -p /home/vya/Pictures
COPY ./pics /home/vya/Pictures
ADD ./glide.yaml /go/src/github.com/YAWAL/gRPCtrying


RUN go get github.com/Masterminds/glide
WORKDIR /go/src/github.com/YAWAL/gRPCtrying
RUN glide install
RUN go install github.com/YAWAL/gRPCtrying/download

ENTRYPOINT ["/go/bin/download"]