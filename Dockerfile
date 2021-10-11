FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy
RUN go clean

#Test
RUN CGO_ENABLED=0 GOOS=linux go test -v ./... 

# Go Security
RUN CGO_ENABLED=0 GOOS=linux go get github.com/securego/gosec/v2/cmd/gosec
RUN CGO_ENABLED=0 GOOS=linux gosec ./...

#Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -mod=mod -o main ./cmd/main/
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]