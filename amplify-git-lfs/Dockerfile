FROM node:18.17.0-slim


# https://docs.aws.amazon.com/amplify/latest/userguide/custom-build-image.html
RUN apt-get update && apt-get install -y \
    git \
    git-lfs \
    curl \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT [ "bash", "-c" ]