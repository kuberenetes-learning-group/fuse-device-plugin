FROM debian:stretch-slim

COPY fuse-device-plugin /usr/bin/fuse-device-plugin

# replace with your desire device count
CMD ["fuse-device-plugin", "--mounts_allowed", "5000"]
