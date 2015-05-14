package download

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api/auth/digest"
	"net/url"
	"time"
)

func b64encode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

//这里的域名可以替换为自己的自定义域名
func PublicDownload() {
	domain := "http://if-pbl.qiniudn.com"
	key := "qiniu.png"
	dnlink := fmt.Sprintf("%s/%s", domain, url.QueryEscape(key))
	fmt.Println(dnlink)
}

func PublicDownloadWithFops() {
	domain := "http://if-pbl.qiniudn.com"
	key := "qiniu.png"
	fops := "watermark/2/fontsize/1200/text/" + b64encode("七牛云存储") + "|imageView2/0/w/100"
	dnlink := fmt.Sprintf("%s/%s?%s", domain, url.QueryEscape(key), fops)
	fmt.Println(dnlink)
}

func PrivateDownload() {
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	domain := "http://if-pri.qiniudn.com"
	key := "qiniu.png"

	expires := 3600
	deadline := time.Now().Add(time.Duration(expires) * time.Second).Unix()
	urlToSign := fmt.Sprintf("%s/%s?e=%d", domain, url.QueryEscape(key), deadline)
	token := digest.Sign(&mac, []byte(urlToSign))

	dnlink := fmt.Sprintf("%s&token=%s", urlToSign, token)
	fmt.Println(dnlink)
}

func PrivateDownloadWithFops() {
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	domain := "http://if-pri.qiniudn.com"
	key := "qiniu.png"
	fops := "watermark/2/fontsize/1200/text/" + b64encode("七牛云存储") + "|imageView2/0/w/100"

	expires := 3600
	deadline := time.Now().Add(time.Duration(expires) * time.Second).Unix()
	urlToSign := fmt.Sprintf("%s/%s?%s&e=%d", domain, url.QueryEscape(key), fops, deadline)
	token := digest.Sign(&mac, []byte(urlToSign))

	dnlink := fmt.Sprintf("%s&token=%s", urlToSign, token)
	fmt.Println(dnlink)
}
