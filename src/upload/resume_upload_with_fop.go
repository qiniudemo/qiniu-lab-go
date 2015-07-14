package upload

import (
	"config"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/api.v6/conf"
	"github.com/qiniu/api.v6/resumable/io"
	"github.com/qiniu/api.v6/rs"
)

//该上传演示了带视频流媒体切片功能的
func ResumeUploadWithFop() {
	//设置上传域名
	conf.UP_HOST = config.UP_HOST
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
	//切片结果保存
	m3u8SaveBucket := bucket
	m3u8SaveKey := "2015/demo/qiniu.m3u8"
	//切片队列（私有队列名称，保障速度)
	//https://portal.qiniu.com/mps/pipeline
	persistentPipeline := "test1"
	//处理结果通知地址（可以不填）
	persistentNotifyUrl := "http://demo.qiniu.com/fake/notify"
	persistentOps := "avthumb/m3u8|saveas/" + base64.URLEncoding.EncodeToString([]byte(m3u8SaveBucket+":"+m3u8SaveKey))
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
