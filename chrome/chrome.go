package chrome

import (
	"context"
	"fmt"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/emulation"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
	"os"
	"time"
	"trackingtracker/internal/config"
)

type Request struct {
	URL      string `json:"url"`
	LoadWait int    `json:"load_wait"`
	Domain   string `json:"domain"`
}

type Response struct {
	Text string `json:"text"`
}

func Process(task Request) (Response, error) {
	var response Response
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Second)
	defer cancel()

	devt := devtool.New(config.Settings.CDPURL)
	pt, err := devt.Get(ctx, devtool.BackgroundPage)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return response, err
		}
	}

	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return response, err
	}
	defer conn.Close()
	c := cdp.NewClient(conn)

	viewportWidth := 1920
	viewportHeight := 1080

	err = c.Emulation.SetDeviceMetricsOverride(ctx, emulation.NewSetDeviceMetricsOverrideArgs(viewportWidth, viewportHeight, 1, false))
	if err != nil {
		return response, err
	}

	userAgent := "Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36"

	err = c.Emulation.SetUserAgentOverride(ctx, &emulation.SetUserAgentOverrideArgs{
		UserAgent:      userAgent,
		AcceptLanguage: nil,
		Platform:       nil,
	})
	if err != nil {
		return response, err
	}

	domContent, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		return response, err
	}
	defer domContent.Close()

	if err = c.Page.Enable(ctx); err != nil {
		return response, err
	}

	navArgs := page.NewNavigateArgs(task.URL)
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return response, err
	}

	if _, err = domContent.Recv(); err != nil {
		return response, err
	}
	time.Sleep(time.Duration(task.LoadWait) * time.Second)

	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		return response, err
	}

	result, err := c.DOM.GetOuterHTML(ctx, &dom.GetOuterHTMLArgs{
		NodeID: &doc.Root.NodeID,
	})
	if err != nil {
		return response, err
	}

	response.Text = result.OuterHTML

	if err != nil {
		return response, err
	}

	defer c.Page.Close(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return response, nil
}
