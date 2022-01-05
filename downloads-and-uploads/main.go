package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const bucketName string = "elasticbeanstalk-us-west-1-815606719766"

func downloadItem(sess *session.Session) {

	file, err := os.Create("downloads-and-uploads/download.txt")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String("my-file.txt"),
		})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully downloaded")
}

func uploadItem(sess *session.Session) {

	f, err := os.Open("downloads-and-uploads/my-file.txt")

	if err != nil {
		log.Fatal("Could not open file")
	}
	defer f.Close()

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(bucketName),
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
		Bucket: aws.String(bucketName),
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
	downloadItem(sess)
}
