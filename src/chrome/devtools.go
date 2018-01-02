package chrome

import (
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
	"github.com/mafredri/cdp"
	"context"
	"time"
	"github.com/mafredri/cdp/protocol/page"
)

func GetScreenShot(url, siteType string) (data []byte, err error) {
	ctx := context.TODO()

	// Use the DevTools json API to get the current page.
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return
		}
	}

	// Connect to Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return
	}
	defer conn.Close() // Must be closed when we are done.

	// Create a new CDP Client that uses conn.
	c := cdp.NewClient(conn)

	// Enable events on the Page domain.
	if err = c.Page.Enable(ctx); err != nil {
		return
	}
	// New DOMContentEventFired client will receive and buffer
	// ContentEventFired events from now on.
	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		return
	}
	defer domContentEventFired.Close()

	// Create the Navigate arguments with the optional Referrer field set.
	navArgs := page.NewNavigateArgs(url)
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return
	}

	// Block until a DOM ContentEventFired event is triggered.
	if _, err = domContentEventFired.Recv(); err != nil {
		return
	}

	// wait 2 second until the full images been rendered
	time.Sleep(2 * time.Second)

	// Capture a screenshot of the current page.
	screenshot, err := c.Page.CaptureScreenshot(ctx, page.NewCaptureScreenshotArgs().SetFormat("png"))
	if err != nil {
		return
	}
	data = screenshot.Data
	return
}
