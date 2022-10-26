#!/bin/bash
set -ex

if [ ! -f /.dockerenv ]; then
    docker run -it -v $(pwd):/app -w /app golangci/golangci-lint:latest sh hack/script/build.sh
    docker run --platform linux/amd64 -it -v $(pwd):/app -w /app unsafe/gobuilder sh hack/script/package.sh
    exit 0
fi

ln -sf /sbin/fy _build/sbin/fycli

# embed
cp bin/fy-embed _build/sbin/fy
fpm -f -s dir -t deb -n fy \
    -a amd64 --rpm-os linux \
    --iteration release -v "1.0.0" \
    -C _build -p "bin/fy_embed_1.0.0-release_amd64.deb" \
    --verbose --url 'https://stdout.com.cn' -m 'rwxr@foxmail.com' \
    --description 'find you'

# external
cp bin/fy-external _build/sbin/fy
mkdir -p _build/root
cp -r lib _build/root

fpm -f -s dir -t deb -n fy \
    -a amd64 --rpm-os linux \
    --iteration release -v "1.0.0" \
    -C _build -p "bin/fy_external_1.0.0-release_amd64.deb" \
    --verbose --url 'https://stdout.com.cn' -m 'rwxr@foxmail.com' \
    --description 'find you'

rm -rf _build
