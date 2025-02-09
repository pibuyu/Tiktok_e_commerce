FROM golang:1.23

WORKDIR /usr/src/Tiktok_e_commerce

#设置代理
ENV GOPROXY=https://goproxy.io.direct

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD["app"]