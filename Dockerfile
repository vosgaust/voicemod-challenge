FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/vosgaust/voicemod
COPY . .
RUN cd cmd && CGO_ENABLED=0 go build -o /go/bin/voicemod

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/voicemod /go/bin/voicemod
ENTRYPOINT ["/go/bin/voicemod"]