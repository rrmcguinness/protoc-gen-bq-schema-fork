module github.com/GoogleCloudPlatform/protoc-gen-bq-schema

go 1.15

replace github.com/GoogleCloudPlatform/protoc-gen-bq-schema/converter => ./internal/converter

require (
	github.com/golang/glog v1.1.2
	github.com/golang/protobuf v1.5.3
	github.com/googleapis/googleapis v0.0.0-20231019093007-d1609218c37f // indirect
	google.golang.org/protobuf v1.31.0
)
