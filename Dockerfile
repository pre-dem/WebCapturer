FROM node:6.9
MAINTAINER moskize "taozeyu@qiniu.com"

RUN mkdir /root/donwloads
WORKDIR /root/donwloads
RUN wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
RUN tar -xvf go1.7.1.linux-amd64.tar.gz
RUN mv go /usr/local
ENV GOROOT /usr/local/go
ENV GOPATH /root/pili/portal-server
