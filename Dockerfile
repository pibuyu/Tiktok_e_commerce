FROM golang:1.23

WORKDIR /usr/src/Tiktok_e_commerce

#设置代理
ENV GOPROXY=https://goproxy.io,direct

COPY app/frontend/go.mod   app/frontend/go.sum  ./app/frontend/
#因为frontend中用到replace,将replace依赖的rpc_gen和common文件夹也拷贝进去
COPY rpc_gen rpc_gen
COPY common common

RUN cd app/frontend/ && go mod download && go mod verify

COPY app/frontend    app/frontend

RUN cd app/frontend/ && go build -v -o /opt/Tiktok_e_commerce/frontend/server

#静态文件&配置文件 拷贝
COPY app/frontend/conf /opt/Tiktok_e_commerce/frontend/conf
COPY app/frontend/static /opt/Tiktok_e_commerce/frontend/static
COPY app/frontend/template /opt/Tiktok_e_commerce/frontend/template

#重新定义工作目录，否则找不到conf
WORKDIR /opt/Tiktok_e_commerce/frontend

EXPOSE 8080

CMD ["./server"]