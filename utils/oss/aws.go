package oss

import (
	"clouds.lgb24kcs.cn/global"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewAws(access_key, secret_key, end_point string) (*session.Session, error) {

	newSession, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:         aws.String(end_point),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改

	})
	if err != nil {
		global.APP.Log.Error(err.Error())
		return nil, err
	} else {
		return newSession, nil
	}
}

/*创建bucket*/
func CreateBucket(access_key, secret_key, end_point, BucketName string) error {
	newAws, _ := NewAws(access_key, secret_key, end_point)
	s := s3.New(newAws)
	params := &s3.CreateBucketInput{Bucket: aws.String(BucketName)}
	_, err := s.CreateBucket(params)
	if err != nil {
		return err
	} else {
		err = s.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: aws.String(BucketName)})
		if err != nil {
			return err
		} else {
			return nil
		}

	}
}

/*加载当前用户下所有桶数据*/
func AllBucketLists(access_key, secret_key, end_point string, bucket string) ([]interface{}, error) {
	newAws, _ := NewAws(access_key, secret_key, end_point)
	s := s3.New(newAws)
	buckets, err := s.ListBuckets(nil)
	if err != nil {
		return nil, err
	} else {
		tmp := []interface{}{}
		for _, b := range buckets.Buckets {
			tmp = append(tmp, aws.StringValue(b.Name))
		}
		return tmp, nil
	}
}
