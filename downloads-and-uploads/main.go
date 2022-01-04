package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadItem(sess *session.Session) {

	f, err := os.Open("downloads-and-uploads/my-file.txt")

	if err != nil {
		log.Fatal("Could not open file")
	}
	defer f.Close()

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String("elasticbeanstalk-us-west-1-815606719766"),
		Key:    aws.String("my-file.txt"),
		Body:   f,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Upload result: %+v\n", result)

}

func listItems(sess *session.Session) {

	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("elasticbeanstalk-us-west-1-815606719766"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range resp.Contents {
		log.Printf("Name: %s\n", *item.Key)
		log.Printf("Size: %d\n", *item.Size)
	}
}

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	uploadItem(sess)
	listItems(sess)
}
