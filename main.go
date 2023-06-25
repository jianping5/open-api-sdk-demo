package main

import (
	"demo/config"
	"fmt"

	openapisdkgo "github.com/jianping5/open-api-sdk-go"
)

func main() {
	// get key pairs from config
	accessKey := config.Config.OpenApi.AccessKey
	secretKey := config.Config.OpenApi.SecretKey

	// New openApiClient
	openApiClient := openapisdkgo.NewClient(accessKey, secretKey)
	
	// demo1：invoke "GET" api
	res := openApiClient.SayHelloUsingGet()
	fmt.Println(res)

	// demo2: invoke "POST" api
	res = openApiClient.GetNameByIdUsingPost("{\"id\": 1}")
	fmt.Println(res)
}