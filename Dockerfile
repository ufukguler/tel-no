#https://klotzandrew.com/blog/smallest-golang-docker-image
FROM golang:1.19-bullseye AS base

RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid 65532 app-user

WORKDIR $GOPATH/app/
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM scratch

WORKDIR /

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .
COPY .env .env
COPY images/ images/

USER app-user:app-user

CMD ["./main"]