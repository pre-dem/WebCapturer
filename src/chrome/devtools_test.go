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
	url := "YOUR_URL"
	cookies := []network.CookieParam{
		{
			Name:"grafana_remember",
			Value:"YOUR_grafana_remember",
			URL:&url,
		},
		{
			Name:"grafana_sess",
			Value:"YOUR_grafana_sess",
			URL:&url,
		},
		{
			Name:"grafana_user",
			Value:"YOUR_grafana_user",
			URL:&url,
		},
	}
	data, err := GetScreenShot("YOUR_URL", "grafana", 1500, 1000, cookies)
	if err != nil {
		log.Error(err)
		return
	}
	screenshotName := fmt.Sprintf("screen_shot_%s.png", time.Now().Format("2006-01-02 15:04:05"))
	if err = ioutil.WriteFile(screenshotName, data, 0644); err != nil {
		log.Error(err)
		return
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
}
