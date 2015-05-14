package fop

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api/auth/digest"
	"github.com/qiniu/rpc"
	"net/http"
)

func SimpleVframeFop() {
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	//空间的名称和已有文件的名称
	bucket := "if-pbl"
	key := "qiniu.mp4"

	//保存结果
	saveBucket := bucket
	saveKey := "2015/s1/snapshot_1.jpg"
	saveEntry := saveBucket + ":" + saveKey
	//处理信息
	persistentOps := "vframe/jpg/offset/3/w/480/h/480|saveas/" +
		base64.URLEncoding.EncodeToString([]byte(saveEntry))
	persistentPipeline := "test1"
	persistentNotifyUrl := "http://demo.qiniu.com/fake/notifyURL"
	//组织接口的参数
	pfopParams := map[string][]string{
		"bucket":    []string{bucket},
		"key":       []string{key},
		"fops":      []string{persistentOps},
		"notifyURL": []string{persistentNotifyUrl},
		"pipeline":  []string{persistentPipeline},
	}
	//创建client
	t := digest.NewTransport(&mac, nil)
	client := &http.Client{Transport: t}
	rpcClient := rpc.Client{client}

	pfopResult := PfopResult{}
	err := rpcClient.CallWithForm(nil, &pfopResult, PFOP_URL, pfopParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("PersistentId:", pfopResult.PersistentId)
	}
}
