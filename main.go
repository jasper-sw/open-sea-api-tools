package main

import (
	"fmt"

	"github.com/jasper-sw/open-sea-api-tools/clients"
)

func main() {
	const baseUrl = "https://api.opensea.io/api/v1"
	openSeaClient := clients.NewOpenSeaAPIClient(baseUrl)
	err, collection := openSeaClient.GetCollection("doodles-official")

	if err != nil {
		return
	}
	fmt.Printf("%#v", collection)
}
