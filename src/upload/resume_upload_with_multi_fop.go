package upload

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api/auth/digest"
	"github.com/qiniu/api/resumable/io"
	"github.com/qiniu/api/rs"
)

//该上传演示了带多个视频流媒体切片功能的代码
func ResumeUploadWithMultiFop() {
	//空间名称
	bucket := "if-pbl"
	//文件在空间中保存名字
	key := "2015/demo/qiniu.mp4"
	//本地路径
	localFile := "/Users/jemy/Documents/qiniu.mp4"
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	//对一个视频做多个处理
	//SD
	sdVideoSaveBucket := bucket
	sdVideoSaveKey := "qiniu_640x320_512k.m3u8"
	sdVideoSaveEntry := sdVideoSaveBucket + ":" + sdVideoSaveKey
	sdVideoFop := "avthumb/m3u8/vb/512k/s/640x320|saveas/" + base64.URLEncoding.EncodeToString([]byte(sdVideoSaveEntry))
	//720P
	highVideoSaveBucket := bucket
	highVideoSaveKey := "qiniu_1280x720_1m.m3u8"
	highVideoSaveEntry := highVideoSaveBucket + ":" + highVideoSaveKey
	highVideoFop := "avthumb/m3u8/vb/1m/s/1280x720|saveas/" + base64.URLEncoding.EncodeToString([]byte(highVideoSaveEntry))
	//1080P
	superVideoSaveBucket := bucket
	superVideoSaveKey := "qiniu_1920x1080_2.5m.m3u8"
	superVideoSaveEntry := superVideoSaveBucket + ":" + superVideoSaveKey
	superVideoFop := "avthumb/m3u8/vb/2.5m/s/1920x1080|saveas/" + base64.URLEncoding.EncodeToString([]byte(superVideoSaveEntry))
	//切片队列（私有队列名称，保障速度)
	//https://portal.qiniu.com/mps/pipeline
	persistentPipeline := "test1"
	//处理结果通知地址（可以不填）
	persistentNotifyUrl := "http://demo.qiniu.com/fake/notify"
	persistentOps := fmt.Sprintf("%s;%s;%s", sdVideoFop, highVideoFop, superVideoFop)
	policy := rs.PutPolicy{
		Scope:               bucket,
		PersistentOps:       persistentOps,
		PersistentPipeline:  persistentPipeline,
		PersistentNotifyUrl: persistentNotifyUrl,
	}
	policy.Expires = 3600 //3600s后过期
	uptoken := policy.Token(&mac)

	//分片上传方式
	putRet := &FopPutRet{}
	err := io.PutFile(nil, putRet, uptoken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(putRet.Hash)
		fmt.Println(putRet.Key)
		//可以使用persistentId去查询处理结果
		//参考http://developer.qiniu.com/docs/v6/api/reference/fop/pfop/prefop.html
		fmt.Println(putRet.PersistentId)
	}
}
