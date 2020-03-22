package screenshot

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

var (
	localFlagp = flag.Int("port", 8080, "port")
)

func localhtml() {
	flag.Parse()
	*localFlagp = 8080
	go localServer(fmt.Sprintf(":%d", *localFlagp))
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var b1 []byte
	err := chromedp.Run(ctx,
		getMessage(fmt.Sprintf("http://localhost:%d", *localFlagp), "#chicon", &b1),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot3.png", b1, 0644); err != nil {
		log.Fatal(err)
	}
}

func getMessage(host string, sel string, v *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, v, chromedp.NodeVisible, chromedp.ByID),
	}
}

func localServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(res, readHtml())
	})
	return http.ListenAndServe(addr, mux)
}

func readHtml() string {
	path := "test.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	fmt.Println(string(data))
	return string(data)
}

func ScreenShotURL(data string) ([]byte, error) {

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, fullScreenshot(data, 80, &buf)); err != nil {
		// if err := chromedp.Run(ctx, captureScreenshot(data, &buf)); err != nil {
		return nil, err
	}
	return buf, nil
}

func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPad),
		chromedp.EmulateViewport(768, 1024, chromedp.EmulateScale(2)),
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}
			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithFormat(page.CaptureScreenshotFormatJpeg).
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
