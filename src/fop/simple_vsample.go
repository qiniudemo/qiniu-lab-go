package fop

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api/auth/digest"
	"github.com/qiniu/rpc"
	"net/http"
)

func SimpleVsampleFop() {
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	//空间的名称和已有文件的名称
	bucket := "if-pbl"
	key := "qiniu.mp4"

	//保存模版
	pattern := "2015/vsample/demo/qiniu_vsample_$(count).jpg"
	//处理信息
	persistentOps := "vsample/jpg/ss/0/t/180/interval/10/pattern/" +
		base64.URLEncoding.EncodeToString([]byte(pattern))
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
