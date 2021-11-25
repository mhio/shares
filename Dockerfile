FROM golang:1.17 as build
WORKDIR /go/src/app
COPY go.sum go.mod /go/src/app/
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=build /go/bin/shares /shares
ENTRYPOINT ["/shares"]
