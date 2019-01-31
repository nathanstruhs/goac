# GOAC

Go AWS Cli

## Getting Started

Set up local environment for Go development: [https://golang.org/doc/install](https://golang.org/doc/install)

Clone repo

Optional: run `go install` from your directory's root to install CLI program on your path. This will enable to you run `goac` from any directory without prepending with `go run`.

## Usage

### Get help

`go run goac help`

### Flags

```
--service value, --svc value   AWS service (default: "s3")
--command value, -c value      Options: [list-buckets, upload-song, upload-album, upload-artist]
--bucket value, -b value       AWS S3 bucket (default: "my-very-first-bucket-yeyeyaya")
--source value, -s value       Source of local file, requires absolute path
--destination value, -d value  S3 storage destination of uploaded file or directory
--help, -h                     show help
--version, -v                  print the version
```

### Examples

List Buckets

`go run goac --service s3 --command list-buckets`

Upload Song

`go run goac --svc s3 --c upload-song -b <bucket> -s <source-path> -d <aws-destination>`

Upload Album

`go run goac --service s3 --command upload-album --bucket <bucket> --source <source-path> --destination <aws-destination>`

Upload Artist

`go run goac --svc s3 -c upload-artist -b <bucket> -s <source-path> -d <aws-destination>`
