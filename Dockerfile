FROM yukinying/chrome-headless

MAINTAINER Craig Wang "wangsiyu@qiniu.com"

WORKDIR /root/screenshot

COPY build build

ENTRYPOINT []

CMD "build/start.sh"
