package fop

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api/auth/digest"
	"github.com/qiniu/rpc"
	"net/http"
)

func SimpleAvthumbForDJ() {
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	//空间的名称和已有文件的名称
	bucket := "skypixeltest"
	key := "4838_848ca76a004611e584f8d73f7f367d12.f0.mp4"

	//点播转码
	//SD
	sdVideoSaveBucket := bucket
	sdVideoSaveKey := "dj_640x360_512k.mp4"
	sdVideoSaveEntry := sdVideoSaveBucket + ":" + sdVideoSaveKey
	sdVideoFop := "avthumb/mp4/vb/512k/s/640x360|saveas/" + base64.URLEncoding.EncodeToString([]byte(sdVideoSaveEntry))
	//720P
	highVideoSaveBucket := bucket
	highVideoSaveKey := "dj_1280x720_1m.mp4"
	highVideoSaveEntry := highVideoSaveBucket + ":" + highVideoSaveKey
	highVideoFop := "avthumb/mp4/vb/1m/s/1280x720|saveas/" + base64.URLEncoding.EncodeToString([]byte(highVideoSaveEntry))
	//1080P
	superVideoSaveBucket := bucket
	superVideoSaveKey := "dj_1920x1080_2.5m.mp4"
	superVideoSaveEntry := superVideoSaveBucket + ":" + superVideoSaveKey
	superVideoFop := "avthumb/mp4/vb/2.5m/s/1920x1080|saveas/" + base64.URLEncoding.EncodeToString([]byte(superVideoSaveEntry))
	//流媒体切片
	//SD
	sdM3u8SaveBucket := bucket
	sdM3u8SaveKey := "dj_640x360_512k.m3u8"
	sdM3u8SaveEntry := sdM3u8SaveBucket + ":" + sdM3u8SaveKey
	sdM3u8Fop := "avthumb/m3u8/vb/512k/s/640x360|saveas/" + base64.URLEncoding.EncodeToString([]byte(sdM3u8SaveEntry))
	//720P
	highM3u8SaveBucket := bucket
	highM3u8SaveKey := "dj_1280x720_1m.m3u8"
	highM3u8SaveEntry := highM3u8SaveBucket + ":" + highM3u8SaveKey
	highM3u8Fop := "avthumb/m3u8/vb/1m/s/1280x720|saveas/" + base64.URLEncoding.EncodeToString([]byte(highM3u8SaveEntry))
	//1080P
	superM3u8SaveBucket := bucket
	superM3u8SaveKey := "dj_1920x1080_2.5m.m3u8"
	superM3u8SaveEntry := superM3u8SaveBucket + ":" + superM3u8SaveKey
	superM3u8Fop := "avthumb/m3u8/vb/2.5m/s/1920x1080|saveas/" + base64.URLEncoding.EncodeToString([]byte(superM3u8SaveEntry))
	//处理信息
	persistentOps := fmt.Sprintf("%s;%s;%s;%s;%s;%s", sdVideoFop, highVideoFop, superVideoFop, sdM3u8Fop, highM3u8Fop, superM3u8Fop)
	persistentPipeline := "p1"
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
