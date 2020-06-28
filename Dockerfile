FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/build .
FROM scratch AS bin
COPY --from=build /out/build /
