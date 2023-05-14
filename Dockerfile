# Simple usage with a mounted data directory:
# > docker build -t fuxchain .
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.fuxchaind:/root/.fuxchaind -v ~/.fuxchaincli:/root/.fuxchaincli fuxchain fuxchaind init mynode
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.fuxchaind:/root/.fuxchaind -v ~/.fuxchaincli:/root/.fuxchaincli fuxchain fuxchaind start
FROM golang:1.20.2-alpine AS build-env

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev

# Set working directory for the build
WORKDIR /go/src/github.com/furyaxyz/fuxchain

# Add source files
COPY . .

ENV GO111MODULE=on \
    GOPROXY=http://goproxy.cn
# Build fuxchain
RUN make install

# Final image
FROM alpine:edge

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/fuxchaind /usr/bin/fuxchaind
COPY --from=build-env /go/bin/fuxchaincli /usr/bin/fuxchaincli

# Run fuxchaind by default, omit entrypoint to ease using container with fuxchaincli
CMD ["fuxchaind"]
