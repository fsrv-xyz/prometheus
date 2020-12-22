package prometheus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// QueryResponse - The return data from a simple prometheus time series query
type QueryResponse struct {
	Endpoint string
	Status   string `json:"status"`
	Data     struct {
		ResultType string `json:"resultType"`
		Result []QueryResult `json:"result"`
	} `json:"data"`
}

// QueryResult - Key-Value formated output of the query
type QueryResult struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
}

// Results - Only returns the Results of the Query
func (obj *QueryResponse) Results() []QueryResult {
  return obj.Data.Results
}

// Query - fetch the data from prometheus API
func (c Client) Query(metric string) QueryResponse {
	if c.Endpoint == "" {
		c.Endpoint = "http://127.0.0.1:9090"
	}
	res, _ := http.Get(fmt.Sprintf("%v/api/v1/query?query=%v", c.Endpoint, metric))
	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var obj QueryResponse
	json.Unmarshal(data, &obj)
	return obj
}

