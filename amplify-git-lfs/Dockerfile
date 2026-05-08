FROM node:22.22.2-slim@sha256:d415caac2f1f77b98caaf9415c5f807e14bc8d7bdea62561ea2fef4fbd08a73c

# https://docs.aws.amazon.com/amplify/latest/userguide/custom-build-image.html
RUN apt-get update && apt-get install -y \
    git \
    git-lfs \
    curl \
    python3 \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT [ "bash", "-c" ]
