# WebCapturer

## 简介
WebCapturer 是基于 `headless chrome` 开发的 web 截图工具，对外提供 `Restful API`

## 接口

```$xslt
GET /v1/get_screenshot
Content-Type: application/json
```
url 参数：

| 参数 | 类型 | 必填 | 说明 |
| - | - | - | - |
| url | string | 是 | 需要截图的 url |
| site_type | string | 否 | 站点类型，当前只支持默认类型或 grafana |
| window_width | int | 是 | 截图窗口宽度 |
| window_height | int | 是 | 截图窗口高度 |
| cookies | int | 是 | 截图之前需要设置的 cookies，样例参考 `src/chrome/cookies.json` |

## 快速开始

### 本地调试
首先参考 [headless chrome 配置指南](https://developers.google.com/web/updates/2017/04/headless-chrome) 在本地运行 headless chrome

然后在本地运行如下命令服务便启动完毕

```$xslt
go run src/app/main.go screenshot_local.conf
```
- 您可以在 `screenshot_local.conf` 中修改监听的端口以及 `chrome devtools` 对应的 url

### 线上部署
推荐使用 [docker](https://www.docker.com/) 进行部署，在根目录运行 `make docker` 等待片刻，即可得到 build 完成的镜像，之后您便可以轻松地将改镜像部署于任意支持 docker 的环境
```$xslt
docker run -d -p 8080:80 screenshotd:v0.0.1
```
