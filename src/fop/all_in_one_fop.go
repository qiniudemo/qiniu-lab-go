package fop

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/rpc"
	"net/http"
)

func SimpleAllFop() {
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	//空间的名称和已有文件的名称
	bucket := "if-pbl"
	key := "qiniu.mp4"

	//m3u8处理指令
	//保存结果
	m3u8SaveBucket := bucket
	m3u8SaveKey := "2015/05/01/qiniu.m3u8"
	m3u8SaveEntry := m3u8SaveBucket + ":" + m3u8SaveKey
	//处理信息
	m3u8Fop := "avthumb/m3u8/noDomain/1|saveas/" +
		base64.URLEncoding.EncodeToString([]byte(m3u8SaveEntry))

	//vsample处理指令
	//vsample保存模版
	pattern := "2015/05/01/qiniu_vsample_$(count).jpg"
	vsampleFop := "vsample/jpg/ss/0/t/180/interval/10/pattern/" +
		base64.URLEncoding.EncodeToString([]byte(pattern))

	//vframe处理指令
	//保存结果
	vframeSaveBucket := bucket
	vframeSaveKey := "2015/05/01/snapshot_1.jpg"
	vframeSaveEntry := vframeSaveBucket + ":" + vframeSaveKey
	//处理信息
	vframeFop := "vframe/jpg/offset/3/w/480/h/480|saveas/" +
		base64.URLEncoding.EncodeToString([]byte(vframeSaveEntry))

	//处理信息
	persistentOps := m3u8Fop + ";" + vsampleFop + ";" + vframeFop
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
