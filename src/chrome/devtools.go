package chrome

import (
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
	"github.com/mafredri/cdp"
	"context"
	"time"
	"github.com/mafredri/cdp/protocol/page"
	"qiniupkg.com/x/log.v7"
	"github.com/mafredri/cdp/protocol/emulation"
	"github.com/mafredri/cdp/protocol/network"
)

func GetScreenShot(url, siteType string, windowWidth, windowHeight int, cookies []network.CookieParam) (data []byte, err error) {
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
	// Initiate a new RPC connection to the Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return
	}
	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)

	// Open a DOMContentEventFired client to buffer this event.
	domContent, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		return
	}
	defer domContent.Close()

	// Enable events on the Page domain, it's often preferrable to create
	// event clients before enabling events so that we don't miss any.
	if err = c.Page.Enable(ctx); err != nil {
		return
	}

	if err = c.Emulation.SetVisibleSize(ctx, emulation.NewSetVisibleSizeArgs(windowWidth, windowHeight)); err != nil {
		return
	}

	err = c.Network.ClearBrowserCookies(ctx)
	if err != nil {
		return
	}


	err = c.Network.SetCookies(ctx, network.NewSetCookiesArgs(cookies))
	if err != nil {
		return
	}

	// Create the Navigate arguments with the optional Referrer field set.
	nav, err := c.Page.Navigate(ctx, page.NewNavigateArgs(url))
	if err != nil {
		return
	}

	// Wait until we have a DOMContentEventFired event.
	if _, err = domContent.Recv(); err != nil {
		return
	}

	log.Infof("Page loaded with frame ID: %s\n", nav.FrameID)

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
