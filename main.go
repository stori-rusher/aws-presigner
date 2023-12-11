package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:embed VERSION
var version string

func main() {
	ctx := context.Background()

	var (
		b, k, m string
		d       int64
		v       bool
	)
	flag.StringVar(&b, "b", "", "bucket")
	flag.Int64Var(&d, "d", 0, "duration in seconds (max 604,800)")
	flag.StringVar(&k, "k", "", "key")
	flag.StringVar(&m, "m", "", "method (get|put)")
	flag.BoolVar(&v, "v", false, "version")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Printf("\nVersion:\n  %s\n", version)
	}

	flag.Parse()

	if v {
		fmt.Printf("v%s\n", version)
		os.Exit(0)
	}

	if d < 1 || d > 604_800 {
		log.Fatalln("duration must be between 1 and 604,800 seconds (7 days)")
	}

	awsCFG, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("error loading S3 config: %v\n", err)
	}

	s3Client := s3.NewFromConfig(awsCFG)

	loc, err := s3Client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{Bucket: aws.String(b)})
	if err != nil {
		log.Fatalf("error getting bucket location: %v\n", err)
	}

	p := NewS3Presigner(s3.NewPresignClient(s3Client), string(loc.LocationConstraint))

	var obj *v4.PresignedHTTPRequest
	switch strings.ToLower(m) {
	case "get":
		obj, err = p.GetObject(ctx, b, k, time.Duration(d)*time.Second)
	case "put":
		obj, err = p.PutObject(ctx, b, k, time.Duration(d)*time.Second)
	default:
		log.Fatalf("invalid method: %v\n", m)
	}

	if err != nil {
		log.Fatalf("error presigning URL: %v\n", err)
	}

	fmt.Println(obj.URL)
}
