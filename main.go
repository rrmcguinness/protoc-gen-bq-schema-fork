/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
protoc plugin which converts .proto to schema for BigQuery.
It is spawned by protoc and generates schema for BigQuery, encoded in JSON.

usage:
$ bin/protoc --bq-schema_out=path/to/outdir foo.proto
*/

// Protobuf code for extensions are generated --
//go:generate protoc --go_out=. --go_opt=module=github.com/GoogleCloudPlatform/protoc-gen-bq-schema bq_table.proto bq_field.proto

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/protoc-gen-bq-schema/converter"
	"github.com/golang/glog"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
)

// a command line parameter to chop the package name if desired, since most tables are bound for a single dataset.
var ignorePrefix = flag.Bool("ignorePrefix", false, "Tells the writer to ignore the package prefix for the output file.")

func main() {
	flag.Parse()
	ok := true
	glog.Infof("Processing code generator request, ignoring prefix: %s", ignorePrefix)
	res, err := converter.ConvertFrom(os.Stdin, *ignorePrefix)
	if err != nil {
		ok = false
		if res == nil {
			message := fmt.Sprintf("Failed to read input: %v", err)
			res = &plugin.CodeGeneratorResponse{
				Error: &message,
			}
		}
	}

	glog.Info("Serializing code generator response")
	data, err := proto.Marshal(res)
	if err != nil {
		glog.Fatal("Cannot marshal response", err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		glog.Fatal("Failed to write response", err)
	}

	if ok {
		glog.Info("Succeeded to process code generator request")
	} else {
		glog.Info("Failed to process code generator but successfully sent the error to protoc")
	}

	// ensure the log is written before the program exits
	glog.Flush()
}
