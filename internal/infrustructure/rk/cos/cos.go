package cos

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
)

const domain = "https://lukeqqin-1302523197.cos.ap-chongqing.myqcloud.com"

var client *cos.Client

type Cos struct {
	SecretId  string `yaml:"secretId" json:"secretId"`
	SecretKey string `yaml:"secretKey" json:"secretKey"`
}

func New(cs Cos) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(domain)
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cs.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: cs.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 3.通过文件流上传对象
	fd, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	rsp, err := client.Object.Put(context.Background(), "test", fd, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
	opt := &cos.BucketGetOptions{
		Prefix:  "test",
		MaxKeys: 3,
	}
	v, _, err := client.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}
}
