FROM node:22.22.2-slim@sha256:d415caac2f1f77b98caaf9415c5f807e14bc8d7bdea62561ea2fef4fbd08a73c

# https://docs.aws.amazon.com/amplify/latest/userguide/custom-build-image.html
# git is needed for turbo - without it hashes are not updated correctly
# zip is needed for building browser extension bundles
# node-canvas requires additional build dependencies https://github.com/Automattic/node-canvas#compiling
RUN apt-get update && apt-get install -y \
    git \
    git-lfs \
    curl \
    zip \
    python3 \
    build-essential \
    pkg-config \
    libpixman-1-dev \
    libcairo2-dev \
    libpango1.0-dev \
    libjpeg-dev \
    libgif-dev \
    librsvg2-dev \
    && rm -rf /var/lib/apt/lists/*

# corepack provides the pnpm/yarn shims (bundled with node, opt-in)
RUN corepack enable

ENTRYPOINT [ "bash", "-c" ]
