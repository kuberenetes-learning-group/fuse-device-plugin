FROM image-registry:5000/debian:stretch-slim

COPY fuse-device-plugin /usr/bin/fuse-device-plugin

CMD ["fuse-device-plugin"]
