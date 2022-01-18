ARG GO_VERSION=1.17.6
FROM golang:${GO_VERSION}-alpine as builder-go
WORKDIR /goBuild
COPY ./src .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app -ldflags="-s -w" ./main.go


FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=builder-go /goBuild/app .
ENTRYPOINT ["./app"]
