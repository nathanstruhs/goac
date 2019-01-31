# GOAC

Go AWS Cli

## Getting Started

Set up local environment for Go development: [https://golang.org/doc/install](https://golang.org/doc/install)

Clone repo

Optional: run `go install` from your directory's root to install CLI program on your path. This will enable to you run `goac` from any directory without prepending with `go run`.

## Usage

Get help

`go run goac help`

List Buckets

`go run goac s3 list-buckets`

Upload Song

`go run goac s3 upload-song <bucket> <source-path> <aws-destination>`

Upload Album

`go run goac s3 upload-album <bucket> <source-path> <aws-destination>`

Upload Artist

`go run goac s3 upload-artist <bucket> <source-path> <aws-destination>`
