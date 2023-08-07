package oss

import (
	"clouds.lgb24kcs.cn/global"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewAws(access_key, secret_key, end_point string) *s3.Client {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: "https://oss-cn-chengdu.aliyuncs.com",
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"LTAI4GFMYtZufJVjY4fPpEam", "EOBJsgdHeADPjnp3VbaO0vdcBSa6IM", "",
		)))
	if err != nil {
		global.APP.Log.Error(err.Error())
		panic(err)
	}

	svc := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = false // 这里没有将桶名放在域名中，所以设置为false

	})
	return svc
}

// /*创建bucket*/
func CreateBucket(access_key, secret_key, end_point, BucketName string) error {
	newAws := NewAws(access_key, secret_key, end_point)
	ok := BucketExists(newAws, BucketName)
	if !ok {
		_, err := newAws.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(BucketName),
		})
		if err != nil {
			global.APP.Log.Error(err.Error())
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("该桶已存在")
	}
	//return nil

}

// 检测桶是否存在
func BucketExists(svc *s3.Client, bucketName string) bool {

	//  检测桶是否存在
	_, err := svc.HeadBucket(context.TODO(), &s3.HeadBucketInput{Bucket: aws.String(bucketName)})
	fmt.Println(bucketName)
	if err != nil {
		global.APP.Log.Error(err.Error())
		fmt.Println(err)
		return false

	} else {
		//log.Printf("Bucket %v exists and you already own it.", bucketName)
		return true
	}

}

/*
删除空桶
桶里面没有任何对象存在，否则删除桶汇报错
*/
func DeleteBucket(access_key, secret_key, end_point, BucketName string) error {
	svc := NewAws(access_key, secret_key, end_point)
	_, err := svc.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(BucketName),
	})
	if err != nil {
		return err
	} else {
		return nil
	}

}

/*详细桶信息*/
func GetBucket(access_key, secret_key, end_point, BucketName string) {
	svc := NewAws(access_key, secret_key, end_point)
	acl, err := svc.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{Bucket: aws.String(BucketName)})
	if err != nil {

	} else {
		fmt.Println(acl)
	}
}

/*目录创建*/

func CreateObject(access_key, secret_key, end_point, BucketName string, RootName string) bool {
	svc := NewAws(access_key, secret_key, end_point)
	_, err := svc.PutObject(context.TODO(), &s3.PutObjectInput{Bucket: aws.String(BucketName), Key: aws.String(RootName)})
	if err != nil {
		global.APP.Log.Error(err.Error())
		return false
	} else {
		return true
	}

}
