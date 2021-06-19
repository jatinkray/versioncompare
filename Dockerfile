############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS build
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/app
ADD . .
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/versioncompare

############################
# STEP 2 binary relase
############################
FROM alpine AS release
WORKDIR /app
ENV PATH=/app:$PATH
COPY --from=build /go/bin/versioncompare /app/versioncompare
USER app
CMD [ "versioncompare" ]