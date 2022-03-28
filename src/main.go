package main

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	u, _ := url.Parse("https://0xleung-1257283327.cos.ap-singapore.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDnPlqeIfla9RTEjwsXw3tSjkTv61e16vT", // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: "JL6aokPowmLegRv85l9G7h25sOdaGgph",     // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "test/objectPut.go"
	// 1.通过字符串上传对象
	f := strings.NewReader("test")

	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	// // 2.通过本地文件上传对象
	// _, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// // 3.通过文件流上传对象
	// fd, err := os.Open("./test")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fd.Close()
	// _, err = c.Object.Put(context.Background(), name, fd, nil)
	// if err != nil {
	// 	panic(err)
	// }
}
