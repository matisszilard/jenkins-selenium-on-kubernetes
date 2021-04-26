FROM golang:1.16.3

RUN apt-get update && apt-get install -y  --no-install-recommends \
    build-essential \
    jq
