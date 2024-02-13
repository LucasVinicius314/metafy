#
# Builder
#
FROM golang:1.21 AS builder

WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./

RUN go mod download && go mod verify

COPY ./server .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o metafy

#
# Client Builder
#
# FROM ubuntu:20.04 AS client-builder

# ARG DEBIAN_FRONTEND=noninteractive

# ARG API_AUTHORITY

# RUN apt-get update && apt-get install -y curl git wget unzip fonts-droid-fallback python3
# RUN apt-get clean

# RUN apt-get update && apt-get install -y build-essential
# RUN apt-get clean

# RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter

# ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"

# WORKDIR /usr/local/flutter

# RUN flutter channel master

# RUN flutter upgrade

# RUN flutter config --enable-web

# WORKDIR /app

# COPY ./client .

# RUN flutter pub get

# RUN flutter build web --release --dart-define API_AUTHORITY=${API_AUTHORITY}

#
# Runner
#
FROM alpine

COPY --from=builder /app/cmd/metafy /app/bin/metafy
COPY --from=builder /app/public /app/public
# COPY --from=client-builder /app/build/web /app/public

WORKDIR /app/bin

CMD [ "./metafy" ]
