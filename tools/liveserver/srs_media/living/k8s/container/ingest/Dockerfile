FROM golang:1.17 as builder
WORKDIR /app
#ENV GOPROXY https://goproxy.cn
ENV GO111MODULE auto
COPY . .
#COPY go.mod go.sum ./
#RUN go mod download
RUN go build -o ingest .

FROM jrottenberg/ffmpeg:4.4-alpine
#设置工作目录
WORKDIR /
#从builder镜像src路径，拷贝到alpine镜像workdir
COPY --from=builder /app/ingest .
COPY --from=builder /app/h264.mov .
COPY --from=builder /app/h265.mp4 .
ENTRYPOINT /ingest
