FROM golang:1.10
WORKDIR /go/src/retro-memo/server
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.* /go/src/retro-memo/server/
RUN dep ensure -vendor-only
COPY . /go/src/retro-memo/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' .


FROM scratch
WORKDIR /srv/app
COPY --from=0 /go/src/retro-memo/server/server .
CMD ["./server"]
