package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"io/ioutil"

	"net/http"

	cos "github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func main() {
	u, _ := url.Parse(fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com", os.Getenv("COS_BUCKET"), os.Getenv("COS_APPID"), os.Getenv("COS_REGION")))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test.txt"
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))

	// range
	opt := &cos.ObjectGetOptions{
		ResponseContentType: "text/html",
		Range:               "bytes=0-3",
	}
	resp, err = c.Object.Get(context.Background(), name, opt)
	if err != nil {
		panic(err)
	}
	bs, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
}
