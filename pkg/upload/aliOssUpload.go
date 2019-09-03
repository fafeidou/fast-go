package upload

import (
	"fast-go/conf"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

func Upload(path, filePath string, reader io.Reader) (url string) {
	client, err := oss.New(conf.App.AliOss.AliyunEndPoint, conf.App.AliOss.AliyunAccessKeyId, conf.App.AliOss.AliyunAccessKeySecret)
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket(conf.App.AliOss.AliyunBucketName)
	if err != nil {
		// HandleError(err)
	}

	if filePath != "" {
		err = bucket.PutObjectFromFile(path, filePath)
	} else {
		err = bucket.PutObject(path, reader)
	}
	return conf.App.AliOss.AliyunDomain + "/" + path
}

