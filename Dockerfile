FROM yukinying/chrome-headless

MAINTAINER WangSiyu "wangsiyu@qiniu.com"

WORKDIR /root/screenshot

COPY build build

ENTRYPOINT []

CMD "build/start.sh"
