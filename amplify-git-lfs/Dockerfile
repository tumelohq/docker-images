FROM node:18.17.0-alpine


# https://docs.aws.amazon.com/amplify/latest/userguide/custom-build-image.html
RUN apk update && apk add \
    bash \
    git \
    git-lfs \
    curl \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT bin/bash
