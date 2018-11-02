FROM golang:1.11.0-alpine

EXPOSE 8081

WORKDIR /go/src/github.com/nakabonne/sxds

COPY . /go/src/github.com/nakabonne/sxds

RUN  apk add --no-cache make git \
     && cd server \
     && go build -o sxds \
     && mv ./sxds /usr/local/bin

CMD ["sxds"]
