FROM public.ecr.aws/docker/library/golang:1.19.6 as builder

COPY go.mod go.sum /workspace/
WORKDIR /workspace
ENV GOPROXY="https://goproxy.io"
RUN go mod download

COPY main.go main.go
COPY service service

RUN GO111MODULE=on CGO_ENABLED=0 go build -a -o echo-server

FROM public.ecr.aws/amazonlinux/amazonlinux:2023
WORKDIR /root/
COPY --from=builder /workspace/echo-server /root/
COPY config.yaml /root/

RUN chmod +x /root/echo-server
CMD /root/echo-server -f config.yaml