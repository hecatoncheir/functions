package function

import (
	"encoding/json"
	"fmt"
	"github.com/hecatoncheir/Storage"
)

type Request struct{ DatabaseGateway, ProductID, CompanyID string }
type Response struct{ Message, Data, Error string }

// Handle a serverless request
func Handle(req []byte) string {
	request := Request{}

	err := json.Unmarshal(req, &request)
	if err != nil {
		warning := fmt.Sprintf(
			"Unmarshal request error: %v. Error: %v", request, err)

		fmt.Println(warning)

		errorResponse := Response{Message: warning, Data: string(req), Error: err.Error()}

		response, err := json.Marshal(errorResponse)
		if err != nil {
			fmt.Println(err)
		}

		return string(response)
	}

	executor := Executor{Store: &storage.Store{DatabaseGateway: request.DatabaseGateway}}

	err = executor.AddCompanyToProduct(request.ProductID, request.CompanyID)
	if err != nil {
		warning := fmt.Sprintf(
			"Add company to product error: %v", err)

		fmt.Println(warning)

		errorResponse := Response{Message: warning, Data: string(req), Error: err.Error()}

		response, err := json.Marshal(errorResponse)
		if err != nil {
			fmt.Println(err)
		}

		return string(response)
	}

	response := Response{}

	encodedResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	return string(encodedResponse)
}
