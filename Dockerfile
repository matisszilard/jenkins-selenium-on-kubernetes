FROM golang:1.15.2

RUN apt-get update && apt-get install -y  --no-install-recommends \
    build-essential \
    jq
