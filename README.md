# GOAC

Go AWS Cli

## Getting Started

Set up local environment for Go development: [https://golang.org/doc/install](https://golang.org/doc/install)

Clone repo

Optional: run `go install` from this directory's root to install cli program on your path. This will enable to you run `goac` from any directory with prepending with `go run`.

## Usage

List Buckets
`go run goac s3 list-buckets`

Upload Song
`go run goac s3 upload-song <source-path> <aws-destination>`

Upload Album
`go run goac s3 upload-album <source-path> <aws-destination>`

Upload Artist
`go run goac s3 upload-artist <source-path> <aws-destination>`
