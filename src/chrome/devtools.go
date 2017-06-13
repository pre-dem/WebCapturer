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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

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

	// Fetch the document root node. We can pass nil here
	// since this method only takes optional arguments.
	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		return
	}

	// Get the outer HTML for the page.
	result, err := c.DOM.GetOuterHTML(ctx, cdpcmd.NewDOMGetOuterHTMLArgs(doc.Root.NodeID))
	if err != nil {
		return
	}

	fmt.Printf("HTML: %s\n", result.OuterHTML)

	time.Sleep(8 * time.Second)

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
