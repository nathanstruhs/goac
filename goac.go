package main

import (
	"goac/commands"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		if c.NArg() < 2 {
			log.Println("Must include service and command")
			return nil
		}

		service := c.Args().Get(0)
		command := c.Args().Get(1)
		bucket := c.Args().Get(2)
		source := c.Args().Get(3)
		destination := c.Args().Get(4)
		sess := session.Must(session.NewSession())

		switch service {
		case "s3":
			executeS3(command, bucket, source, destination, sess)
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

func executeS3(command string, bucket string, source string, destination string, sess *session.Session) {
	svc := s3.New(sess)

	switch command {
	case "list-buckets":
		commands.ListBuckets(svc)
	case "upload-song":
		commands.UploadSong(svc, bucket, source, destination)
	case "upload-album":
		commands.UploadAlbum(svc, bucket, source, destination)
	case "upload-artist":
		commands.UploadAlbum(svc, bucket, source, destination)
	default:
		log.Println("Invalid command", command)
	}
}
