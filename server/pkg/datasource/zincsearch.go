package datasource

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	contractsservice "indexed-mail-search/server/pkg/service/contracts_service"

	"github.com/dghubble/sling"
)

const (
	defaultZincSearchHost = "http://localhost:4080"
	headerContentType     = "Content-Type"
	applicationJSON       = "application/json"
	indexName             = "enron_emails"
)

// ZincsearchClient is the client that will communicate with the zincsearch api
type ZincsearchClient struct {
	client *http.Client
	sling  *sling.Sling
}

// NewZincsearchClient works as the conntrucutor of the ZincsearchClient struc
func NewZincsearchClient(client *http.Client) *ZincsearchClient {
	newSling := sling.New().Client(client).Base(defaultZincSearchHost)
	setBasicAuthentication(newSling)
	return &ZincsearchClient{
		client: client,
		sling:  newSling,
	}
}

// setBasicAuthentication uses two environment variables that must already be set
// to perform an authentication with the zincsearch API
func setBasicAuthentication(newSling *sling.Sling) {
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		panic("ZINC_FIRST_ADMIN_USER and ZINC_FIRST_ADMIN_PASSWORD must be set")
	}
	newSling.SetBasicAuth(username, password)
}

// CreateEmails uses the zincsearch API to create indexed documents
func (zc *ZincsearchClient) CreateEmails(emails interface{}) (*contractsservice.CreateEmailsResponse, error) {
	succesResponse := &contractsservice.CreateEmailsResponse{}
	errorResponse := &contractsservice.ErrorReponse{}
	url := "/api/_bulkv2"
	bodyRequest := contractsservice.CreateEmailsRequest{
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

// IndexedSearch uses the zincsearch API to perform an indexed search for a term within the content of documents
func (zc *ZincsearchClient) IndexedSearch(bodyrequest contractsservice.IndexedSearchRequest) (*contractsservice.IndexedSearchResponse, error) {
	succesResponse := &contractsservice.IndexedSearchResponse{}
	errorResponse := &contractsservice.ErrorReponse{}
	url := "/api/" + indexName + "/_search"

	request, err := makeRequest(zc.sling, http.MethodPost, url, bodyrequest)
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

// makeRequest makes a request to the provided url with the provided POST method and body
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
