package chrome

import (
	"testing"
	"fmt"
	"io/ioutil"
	"qiniupkg.com/x/log.v7"
	"time"
	"github.com/mafredri/cdp/protocol/network"
)

func TestGetScreenShot(t *testing.T) {
	url := "http://zw0de1gx.nq.cloudappl.com"
	cookies := []network.CookieParam{
		{
			Name:"grafana_remember",
			Value:"e0f8a0242cda66ef54b89d5fb745a0b55f753cb8101ca6fea2365323ad4a933899",
			URL:&url,
		},
		{
			Name:"grafana_sess",
			Value:"26b2cf2831cd416f",
			URL:&url,
		},
		{
			Name:"grafana_user",
			Value:"admin",
			URL:&url,
		},
	}
	data, err := GetScreenShot("http://zw0de1gx.nq.cloudappl.com/dashboard/db/predem-nginx-logs?refresh=30s&orgId=1", "grafana", 1500, 1000, cookies)
	if err != nil {
		log.Error(err)
		return
	}
	screenshotName := fmt.Sprintf("/Users/wangsiyu/Desktop/screenshot/screen_shot_%s.png", time.Now().Format("2006-01-02 15:04:05"))
	if err = ioutil.WriteFile(screenshotName, data, 0644); err != nil {
		log.Error(err)
		return
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
}
