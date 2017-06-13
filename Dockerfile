FROM yukinying/chrome-headless

MAINTAINER Craig Wang "wangsiyu@qiniu.com"

COPY release_screenshot /root/screenshot
WORKDIR /root/screenshot
CMD ["/root/screenshot/start.sh"]
