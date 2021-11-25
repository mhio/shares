FROM golang:1.17 as build
WORKDIR /app
COPY go.sum go.mod /app/
RUN go mod download
COPY . /app/
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=build /app/shares /shares
ENTRYPOINT ["/shares"]
