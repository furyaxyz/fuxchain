FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

WORKDIR /fuxchain
# Add source files
COPY . .

# build fuxchain
RUN make build-linux

# Final image
FROM golang:1.20 as final

WORKDIR /fuxchaind
# Copy over binaries from the build-env
COPY --from=build-env /fuxchain/build/fuxchaind /usr/bin/fuxchaind
COPY --from=build-env /fuxchain/build/fuxchaincli /usr/bin/fuxchaincli
COPY --from=build-env /fuxchain/networks/local/node/wrapper.sh /usr/bin/wrapper.sh

EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

