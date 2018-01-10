FROM yukinying/chrome-headless

MAINTAINER WangSiyu "wangsiyu@qiniu.com"

WORKDIR /root/WebCapturer

COPY build build

RUN apt-get update \
	&& apt-get install fonts-arphic-ukai fonts-arphic-uming

ENTRYPOINT []

CMD "build/start.sh"
