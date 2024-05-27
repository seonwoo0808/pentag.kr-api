FROM golang:1.21-alpine AS backend-builder

WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ .
RUN go build -o main .

FROM alpine:3.18

COPY --from=backend-builder /build/main .

# Install Doppler CLI
RUN wget -q -t3 'https://packages.doppler.com/public/cli/rsa.8004D9FF50437357.key' -O /etc/apk/keys/cli@doppler-8004D9FF50437357.rsa.pub && \
    echo 'https://packages.doppler.com/public/cli/alpine/any-version/main' | tee -a /etc/apk/repositories && \
    apk add doppler

CMD ["doppler", "run", "--", "/main"]