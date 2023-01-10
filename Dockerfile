############################
# STEP 1 build executable binary
############################
FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o config-shepherd ./cmd/config-shepherd

WORKDIR /app/build

RUN cp -r /app/config-shepherd /app/config.schema.json /app/LICENSE .

############################
# STEP 2 build service image
############################

FROM scratch

ARG COMMIT_SHA=<not-specified>

LABEL maintainer="undefined" \
  name="config-shepherd" \
  description="" \
  eu.mia-platform.url="https://www.mia-platform.eu" \
  vcs.sha="$COMMIT_SHA"

ENV VERSION="1.0.2"

WORKDIR /app

COPY --from=builder /app/build/* /usr/local/bin/
COPY --from=builder /app/build/config.schema.json .
# Use an unprivileged user.
USER 1000

CMD ["config-shepherd"]
