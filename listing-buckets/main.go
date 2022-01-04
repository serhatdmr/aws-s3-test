package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	fmt.Println("Listing Buckets")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, bucket := range result.Buckets {

		fmt.Printf("Bucket %s\n", aws.StringValue(bucket.Name))
	}

}
