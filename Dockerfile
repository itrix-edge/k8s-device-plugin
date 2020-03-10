FROM nvcr.io/nvidia/l4t-base:r32.3.1

RUN apt-get update && apt-get install -y --no-install-recommends \
        g++ \
        ca-certificates \
        wget && \
    rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.10.3
RUN wget -nv -O - https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-arm64.tar.gz \
    | tar -C /usr/local -xz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV LD_LIBRARY_PATH=/usr/local/cuda/lib:$LD_LIBRARY_PATH

WORKDIR /go/src/semaphore-device-plugin
COPY . .

RUN go build

//FROM debian:stretch-slim

ENV NVIDIA_VISIBLE_DEVICES=all
ENV NVIDIA_DRIVER_CAPABILITIES=utility
ENV LD_LIBRARY_PATH=/usr/local/cuda/lib:$LD_LIBRARY_PATH

//COPY --from=0 /go/src/semaphore-device-plugin/semaphore-device-plugin /usr/bin/semaphore-device-plugin

//CMD ["semaphore-device-plugin"]
CMD ["/go/src/semaphore-device-plugin/semaphore-device-plugin"]