package main

import (
	"bytes"
	"fmt"

	generator "github.com/awslabs/go-config-generator-for-fluentd-and-fluentbit"
)

func sample1() {
	var config generator.FluentConfig
	config = generator.New()
	config.AddInput("tail", "sample1.*", map[string]string{
		"Path":             "/var/log/test.log",
		"Multiline":        "On",
		"Parser_Firstline": "multiline",
	}).AddOutput("cloudwatch_logs", "*", map[string]string{
		"region":            "us-west-2",
		"log_group_name":    "fluent-bit-cloudawatch",
		"log_stream_prefix": "from-fluent-bit-",
		"auto_create_group": "On",
	}).AddIncludeFilter("*failure*", "log", "*").
		AddExcludeFilter("*success*", "log", "*").
		AddExternalConfig("./filters.conf", generator.EndOfFile)

	fluentbitConfig := new(bytes.Buffer)
	config.WriteFluentBitConfig(fluentbitConfig)
	fmt.Println(fluentbitConfig.String())

	fluentdConfig := new(bytes.Buffer)
	config.WriteFluentdConfig(fluentdConfig)
	fmt.Println(fluentdConfig.String())
}

func main() {
	sample1()
}
