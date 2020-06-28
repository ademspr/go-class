FROM golang:1.14.3-alpine AS build
RUN apk update \
    && apk add git
WORKDIR $GOPATH/src/github.com/ademspr/go-class
COPY . .
RUN go get ./... \
    && go build -o /out/go-class .

FROM golang:1.14.3-alpine AS bin
COPY --from=build /out/go-class /

CMD ["/go-class"]
