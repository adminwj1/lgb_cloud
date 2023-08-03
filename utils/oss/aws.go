package oss

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"log"
)

// func NewAws(access_key, secret_key, end_point string) (*session.Session, error) {
//
//		newSession, err := session.NewSession(&aws.Config{
//			Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
//			Endpoint:         aws.String(end_point),
//			Region:           aws.String("us-east-1"),
//			DisableSSL:       aws.Bool(true),
//			S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
//
//		})
//		if err != nil {
//			global.APP.Log.Error(err.Error())
//			return nil, err
//		} else {
//			return newSession, nil
//		}
//	}
//
// /*创建bucket*/
//
//	func CreateBucket(access_key, secret_key, end_point, BucketName string) error {
//		newAws, _ := NewAws(access_key, secret_key, end_point)
//		s := s3.New(newAws)
//		params := &s3.CreateBucketInput{Bucket: aws.String(BucketName)}
//		_, err := s.CreateBucket(params)
//		if err != nil {
//			return err
//		} else {
//			err = s.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: aws.String(BucketName)})
//			if err != nil {
//				return err
//			} else {
//				return nil
//			}
//
//		}
//	}
//
// /*加载当前用户下所有桶数据*/
// func AllBucketLists() {
//
// }
type BucketBasics struct {
	S3Client *s3.Client
}

func (basics BucketBasics) CreateBucket(name string, region string) error {
	_, err := basics.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(name).String(),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		log.Printf("Couldn't create bucket %v in Region %v. Here's why: %v\n",
			name, region, err)
	}
	return err
}
