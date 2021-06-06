FROM golang:1.16.5

RUN apt-get update && apt-get install -y  --no-install-recommends \
    build-essential \
    jq
