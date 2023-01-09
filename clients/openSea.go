package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jasper-sw/open-sea-api-tools/dtos"
)

type IOpenSeaAPIClient interface {
	NewOpenSeaAPIClient()
	GetCollections()
	GetCollection()
}

type OpenSeaAPIClient struct {
	ApiBaseUrl string
	IOpenSeaAPIClient
}

func NewOpenSeaAPIClient(ApiBaseUrl string) *OpenSeaAPIClient {
	return &OpenSeaAPIClient{
		ApiBaseUrl: ApiBaseUrl,
	}
}

func (o OpenSeaAPIClient) GetCollections() (error, []dtos.Collection) {
	return nil, []dtos.Collection{}
}

func (o OpenSeaAPIClient) GetCollection(slug string) (error, dtos.Collection) {
	combinedUrl := fmt.Sprintf(`%s/collection/%s`, o.ApiBaseUrl, slug)
	resp, err := http.Get(combinedUrl)
	if err != nil {
		fmt.Printf(`OpenSeaAPIClient hit error when getting collection. Err: %#v\n`, err)
		return err, dtos.Collection{}
	}

	type responseDto struct {
		Collection dtos.Collection `json:"collection"`
	}
	var respDto responseDto
	err = json.NewDecoder(resp.Body).Decode(&respDto)
	if err != nil {
		fmt.Printf(`OpenSeaAPIClient hit error when decoding response to CollectionDto. Err: %#v\n`, err)
		return err, dtos.Collection{}
	}

	return nil, respDto.Collection
}
