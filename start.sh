#!/bin/sh
/chrome/headless_shell --no-sandbox --remote-debugging-address=0.0.0.0 --remote-debugging-port=9222 & /root/screenshot/screenshotd /root/screenshot/screenshot.conf