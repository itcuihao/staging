package screenshot

import (
	"flag"
	"fmt"
	"net/http"
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)


var(
	localFlagp=flag.Int("port",8080,"port")
)

func localhtml() {
	flag.Parse()
	*localFlagp=8080
	go localServer(fmt.Sprintf(":%d",*localFlagp))
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var b1 []byte
	 err := chromedp.Run(ctx,
		getMessage(fmt.Sprintf("http://localhost:%d",*localFlagp),"#chicon",&b1),
	)

	if  err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot3.png", b1, 0644); err != nil {
		log.Fatal(err)
	}
}

func getMessage(host string,sel string,v *[]byte)chromedp.Tasks{
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, v, chromedp.NodeVisible, chromedp.ByID),
	}
}

func localServer(addr string)error{
	mux:=http.NewServeMux()
	mux.HandleFunc("/",func(res http.ResponseWriter,_ *http.Request){
		fmt.Fprint(res,readHtml())
	})	
	return http.ListenAndServe(addr,mux)
}

func readHtml() string {
	path:="test.html"
	data,err:=ioutil.ReadFile(path)
	if err !=nil{
		return err.Error()
	}
	fmt.Println(string(data))
	return string(data)
}
