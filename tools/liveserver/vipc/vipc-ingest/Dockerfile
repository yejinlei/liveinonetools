FROM golang:1.18 as builder
WORKDIR /
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

COPY . .
RUN go mod init yejinlei/ingest && go build -o ingest .

FROM golang:1.18
WORKDIR /
COPY --from=builder /ingest .
COPY --from=builder /h264.mov .
COPY --from=builder /h265.mp4 .
ENTRYPOINT /ingest
