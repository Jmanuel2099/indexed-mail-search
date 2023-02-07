package datasource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"indexed-mail-search/server/pkg/handlers/contracts"
	"io"
	"net/http"
	"os"

	"github.com/dghubble/sling"
)

const (
	defaultZincSearchHost = "http://localhost:4080"
	headerContentType     = "Content-Type"
	applicationJSON       = "application/json"
	indexName             = "enron_emails"
)

type ZincsearchClient struct {
	client *http.Client
	sling  *sling.Sling
}

func NewZincsearchClient(client *http.Client) *ZincsearchClient {
	newSling := sling.New().Client(client).Base(defaultZincSearchHost)
	setBasicAuthentication(newSling)
	return &ZincsearchClient{
		client: client,
		sling:  newSling,
	}
}

func setBasicAuthentication(newSling *sling.Sling) {
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		panic("ZINC_FIRST_ADMIN_USER and ZINC_FIRST_ADMIN_PASSWORD must be set")
	}
	newSling.SetBasicAuth(username, password)
}

func (zc *ZincsearchClient) CreateEmails(emails interface{}) (*contracts.CreateEmailsResponse, error) {
	succesResponse := &contracts.CreateEmailsResponse{}
	errorResponse := &contracts.ErrorReponse{}
	url := "/api/_bulkv2"
	bodyRequest := contracts.CreateEmailsRequest{
		Index:   indexName,
		Records: emails,
	}

	request, err := makeRequest(zc.sling, http.MethodPost, url, bodyRequest)
	if err != nil {
		return nil, err
	}

	response, err := zc.sling.Do(request, succesResponse, errorResponse)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	return succesResponse, nil
}

func (zc *ZincsearchClient) IndexedSearch(bodyrequest contracts.IndexedSearchRequest) (*contracts.IndexedSearchResponse, error) {
	succesResponse := &contracts.IndexedSearchResponse{}
	errorResponse := &contracts.ErrorReponse{}
	url := "/api/" + indexName + "/_search"

	request, err := makeRequest(zc.sling, http.MethodPost, url, bodyrequest)
	if err != nil {
		// fmt.Println("Make request error: " + err.Error())
		return nil, err
	}
	response, err := zc.sling.Do(request, succesResponse, errorResponse)
	if err != nil {
		fmt.Println("Do error: " + err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, err
	}
	fmt.Println(succesResponse.Hits.Total.Value)
	return succesResponse, nil
}

func makeRequest(sling *sling.Sling, method string, path string, body interface{}) (*http.Request, error) {
	bodyRequest := makeBodyRequest(body)

	requestSilng, err := sling.New().Post(path).
		Set(headerContentType, applicationJSON).
		Body(bodyRequest).
		Request()

	if err != nil {
		return nil, err
	}

	return requestSilng, nil
}

func makeBodyRequest(bodyRequest interface{}) io.Reader {
	if bodyRequest == nil {
		return nil
	}
	body, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(body)
}
