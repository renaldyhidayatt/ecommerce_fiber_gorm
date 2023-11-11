FROM golang:1.21.4-alpine AS builder

RUN apk update && \
    apk add git && \
    apk add build-base upx

WORKDIR /src/app
COPY . .
# Fetch dependencies.
# Using go get.
# Build the binary.
RUN go build  -o /go/bin/app cmd/app/main.go
RUN upx /go/bin/app

############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
RUN apk update && apk add --no-cache  vips-dev
COPY --from=builder /go/bin/app /go/bin/app


# Run the hello binary.
ENTRYPOINT ["/go/bin/app"]
