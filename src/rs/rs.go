package rs

import (
	"config"
	"fmt"
	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/api.v6/rs"
	"time"
)

//复制一个文件
//可以是复制到其他空间或者创建本空间中文件的副本
func Copy() {
	srcBucket := "if-pbl"
	dstBucket := "if-pri"
	srcKey := "qiniu.png"
	//如果是不同空间，目标文件名可以和原文件名同名
	//如果是相同空间，目标文件名不可以和原文件名同名
	dstKey := "2015/qiniu.png"

	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	client := rs.New(&mac)
	err := client.Copy(nil, srcBucket, srcKey, dstBucket, dstKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Copy Done!")
	}
}

//移动一个文件
//可以用来给本空间的文件重命名或者移动到其他的空间
func Move() {
	srcBucket := "if-pbl"
	dstBucket := "if-pri"

	srcKey := "qiniu.png"
	//如果是不同空间，目标文件名可以和原文件名相同
	//如果是相同空间，目标文件名不可以和原文件名相同
	//在同一个空间的情况下，该Move操作的实际结果是给文件重命名
	dstKey := "2015/05/15/qiniu.png"
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	client := rs.New(&mac)
	err := client.Move(nil, srcBucket, srcKey, dstBucket, dstKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Move Done!")
	}
}

//获取文件的基本信息
//可以获取文件的hash值，大小，上传时间，类型和customer字段
func Stat() {
	bucket := "if-pbl"
	key := "qiniu.png"
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	client := rs.New(&mac)
	entry, err := client.Stat(nil, bucket, key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Key:", key)
		fmt.Println("Hash:", entry.Hash)
		fmt.Println("Fsize:", entry.Fsize)
		fmt.Println("PutTime:", time.Unix(0, entry.PutTime*100).String())
		fmt.Println("MimeType:", entry.MimeType)
		fmt.Println("Customer:", entry.Customer)
	}
}

//删除空间中的一个文件
func Delete() {
	bucket := "if-pbl"
	key := "qiniu.png"
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	client := rs.New(&mac)
	err := client.Delete(nil, bucket, key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Done!")
	}
}

//修改空间中文件的类型
func ChangeMimeType() {
	bucket := "if-pbl"
	key := "qiniu.png"
	newMimeType := "image/png"
	mac := digest.Mac{
		config.AccessKey,
		[]byte(config.SecretKey),
	}
	client := rs.New(&mac)
	err := client.ChangeMime(nil, bucket, key, newMimeType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Change MimeType Done!")
	}
}
