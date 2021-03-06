package main

import (
	"goac/commands"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/urfave/cli"
)

func main() {
	os.Setenv("AWS_REGION", "us-east-1")

	var service, command, bucket, source, destination, genre string

	app := cli.NewApp()
	app.Name = "GOAC - Go AWS Cli"
	app.Usage = "goac [aws-service] [command] (command options...)"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "service, svc",
			Value:       "s3",
			Usage:       "AWS service",
			Destination: &service,
		},
		cli.StringFlag{
			Name:        "command, c",
			Usage:       "Options: [list-buckets, upload-song, upload-album, upload-artist]",
			Destination: &command,
		},
		cli.StringFlag{
			Name:        "bucket, b",
			Value:       "struhs-spotify-clone",
			Usage:       "AWS S3 bucket",
			Destination: &bucket,
		},
		cli.StringFlag{
			Name:        "source, s",
			Usage:       "Source of local file, requires absolute path",
			Destination: &source,
		},
		cli.StringFlag{
			Name:        "destination, d",
			Usage:       "S3 or Dynamo storage destination of uploaded file or directory",
			Destination: &destination,
		},
		cli.StringFlag{
			Name:        "genre, g",
			Usage:       "Dynamo hash key",
			Destination: &genre,
		},
	}

	app.Action = func(c *cli.Context) error {
		if command == "" {
			log.Println("Must include a command")
			return nil
		}

		sess := session.Must(session.NewSession())

		switch service {
		case "s3":
			executeS3(command, bucket, source, destination, sess)
		case "dynamo":
			executeDynamo(command, genre, destination, sess)
		default:
			log.Println("Invalid service")
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func executeDynamo(command string, genre string, destination string, sess *session.Session) {
	svc := dynamodb.New(sess)

	switch command {
	case "put-item":
		commands.PutItem(svc, genre, destination)
	default:
		log.Println("Invalid Command", command)
	}
}

func executeS3(command string, bucket string, source string, destination string, sess *session.Session) {
	svc := s3.New(sess)

	switch command {
	case "list-buckets":
		commands.ListBuckets(svc)
	case "upload-song":
		commands.UploadSong(svc, bucket, source, destination)
	case "upload-album":
		commands.UploadDirectory(svc, bucket, source, destination)
	case "upload-artist":
		commands.UploadDirectory(svc, bucket, source, destination)
	default:
		log.Println("Invalid command", command)
	}
}
