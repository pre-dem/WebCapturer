package chrome

import (
	"fmt"
	"github.com/mafredri/cdp/cdpcmd"
	"io/ioutil"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
	"github.com/mafredri/cdp"
	"context"
	"time"
)

func GetScreenShot(url string) (data []byte, err error) {
	ctx := context.TODO()

	// Use the DevTools json API to get the current page.
	devt := devtool.New("http://127.0.0.1:9222")
	page, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		page, err = devt.Create(ctx)
		if err != nil {
			return
		}
	}

	// Connect to Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, page.WebSocketDebuggerURL)
	if err != nil {
		return
	}
	defer conn.Close() // Must be closed when we are done.

	// Create a new CDP Client that uses conn.
	c := cdp.NewClient(conn)

	err = c.Emulation.SetVisibleSize(ctx, &cdpcmd.EmulationSetVisibleSizeArgs{Width:510, Height:420})

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
	navArgs := cdpcmd.NewPageNavigateArgs(url)
	nav, err := c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return
	}

	// Block until a DOM ContentEventFired event is triggered.
	if _, err = domContentEventFired.Recv(); err != nil {
		return
	}

	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

	for {
		var doc *cdpcmd.DOMGetDocumentReply
		doc, err = c.DOM.GetDocument(ctx, nil)
		if *doc.Root.Children[1].Children[1].ChildNodeCount != 2 {
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	// Capture a screenshot of the current page.
	screenshotName := "screenshot.png"
	screenshot, err := c.Page.CaptureScreenshot(ctx, cdpcmd.NewPageCaptureScreenshotArgs().SetFormat("png"))
	if err != nil {
		return
	}
	if err = ioutil.WriteFile(screenshotName, screenshot.Data, 0644); err != nil {
		return
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
	data = screenshot.Data
	return
}
