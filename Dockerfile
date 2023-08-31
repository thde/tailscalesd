FROM alpine:3
COPY tailscalesd /usr/bin/tailscalesd
RUN apk --no-cache add ca-certificates
USER guest
ENTRYPOINT ["/usr/local/bin/tailscalesd"]
