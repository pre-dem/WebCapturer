package chrome

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestGetScreenShot(t *testing.T) {
	data, err := GetScreenShot("https://www.taobao.com", "")
	screenshotName := "screenshot.png"
	if err = ioutil.WriteFile(screenshotName, data, 0644); err != nil {
		return
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
}
