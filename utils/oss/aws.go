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
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func NewAws(access_key, secret_key, end_point string) *s3.Client {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: end_point,
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			access_key, secret_key, "",
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
func GetBucket(access_key, secret_key, end_point, BucketName string) bool {
	svc := NewAws(access_key, secret_key, end_point)
	//acl, err := svc.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{Bucket: aws.String(BucketName)}) // 获取桶的详细数据
	if exists := BucketExists(svc, BucketName); exists {
		return true
	} else {
		return false
	}

}

/*目录创建*/

func CreateObject(access_key, secret_key, end_point, BucketName string, RootName string) bool {
	svc := NewAws(access_key, secret_key, end_point)
	if ok := ObjectExists(svc, BucketName, RootName); ok {
		return false
	} else {
		// 创建目录，如果要创建嵌套目录，请在目录路径中包含完整的路径，例如 directoryPath := "folder/subfolder/"
		directoryPath := BucketName + "/" + RootName + "/" // 一定要加最后一根斜杠否则创建的是以字节数为0的文件
		input := &s3.PutObjectInput{
			Bucket: &BucketName,
			Key:    &directoryPath,
		}

		_, err := svc.PutObject(context.TODO(), input)
		if err != nil {
			global.APP.Log.Error(err.Error())
			return false
		} else {
			return true
		}

	}

}

/* 删除多个目录*/
func DelObjects(access_key, secret_key, end_point, BucketName string, objectKeys []string) bool {
	svc := NewAws(access_key, secret_key, end_point)
	var objectIds []types.ObjectIdentifier
	for _, key := range objectKeys {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}
	_, err := svc.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{Bucket: aws.String(BucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		global.APP.Log.Error(err.Error())
		return false
	} else {
		return true
	}
}

/*检查目录对象是否存在*/
func CatalogueObjectExists(svc *s3.Client, BucketName, CatalogueName string) bool {
	// s3中没有目录的概念，这里实现目录功能是通过创建没有数据的空对象实现的
	_, err := svc.HeadObject(context.TODO(), &s3.HeadObjectInput{Bucket: aws.String(BucketName), Key: aws.String(BucketName + "/" + CatalogueName + "/")})
	if err != nil {
		global.APP.Log.Error(err.Error())
		return false
	} else {
		return true
	}

}

/*目录对象详情信息，这个方法存在问题*/
func GetObject(access_key, secret_key, end_point, BucketName string, objectKeys string) error {
	svc := NewAws(access_key, secret_key, end_point)

	//	检查对象是否存在
	if exists := CatalogueObjectExists(svc, BucketName, objectKeys); !exists {
		return errors.New("对象不存在")

	} else {
		return nil
	}
}

/*删除单个目录对象*/
func DelCatalogue(access_key, secret_key, end_point, BucketName string, objectKeys string) bool {
	svc := NewAws(access_key, secret_key, end_point)
	//检查目录对象是否存在

	exists := CatalogueObjectExists(svc, BucketName, objectKeys)
	if !exists {
		return false
	} else {
		_, err := svc.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: aws.String(BucketName), Key: aws.String(BucketName + "/" + objectKeys + "/")})
		if err != nil {
			global.APP.Log.Error(err.Error())
			return false
		} else {
			return true
		}
	}

}

/*检查对象是否存在*/
func ObjectExists(svc *s3.Client, BucketName, CatalogueName string) bool {
	_, err := svc.HeadObject(context.TODO(), &s3.HeadObjectInput{Bucket: aws.String(BucketName), Key: aws.String(CatalogueName)})
	if err != nil {
		global.APP.Log.Error(err.Error())
		return false
	} else {
		return true
	}

}

/*删除单个文件对象*/
func DelObject(access_key, secret_key, end_point, BucketName string, objectKeys string) bool {
	svc := NewAws(access_key, secret_key, end_point)
	//检查目录对象是否存在

	exists := CatalogueObjectExists(svc, BucketName, objectKeys)
	if !exists {
		return false
	} else {
		_, err := svc.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: aws.String(BucketName), Key: aws.String(objectKeys)})
		if err != nil {
			global.APP.Log.Error(err.Error())
			return false
		} else {
			return true
		}
	}

}
