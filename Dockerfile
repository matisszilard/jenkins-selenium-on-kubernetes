FROM golang:1.15.3

RUN apt-get update && apt-get install -y  --no-install-recommends \
    build-essential \
    jq
