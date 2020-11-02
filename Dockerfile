FROM golang:1.15-alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o oauth-server .

FROM scratch
COPY --from=builder /build/oauth-server /app/
WORKDIR /app
CMD ["./oauth-server"]