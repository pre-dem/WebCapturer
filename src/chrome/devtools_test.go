package chrome

import (
	"testing"
	"fmt"
	"io/ioutil"
	"qiniupkg.com/x/log.v7"
	"time"
)

func TestGetScreenShot(t *testing.T) {
	data, err := GetScreenShot("https://www.taobao.com", "", 1500, 1000)
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
