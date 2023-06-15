package uptycs

import (
	"errors"
	"fmt"
)

func (T QueryJobResult) GetID() string {
	return ""
}

func (T QueryJobResult) GetName() string {
	return T.Name
}

func (T QueryJobResult) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateQueryJobResults(queryJobResult QueryJobResult) (QueryJobResult, error) {
	return QueryJobResult{}, errors.New("CREATE is not supported for query job results")
}

func (c *Client) UpdateQueryJobResults(queryJobResult QueryJobResult) (QueryJobResult, error) {
	return QueryJobResult{}, errors.New("UPDATE is not supported for query job results")
}

func (c *Client) GetQueryJobResults(queryJobResult QueryJobResult) (QueryJobResult, error) {
	return doGetMany(c, queryJobResult, fmt.Sprintf("queryJobs/%s/results", queryJobResult.ID))
}

func (c *Client) DeleteQueryJobResult(queryJobResult QueryJobResult) (QueryJobResult, error) {
	return QueryJobResult{}, errors.New("DELETE is not supported for query job results")
}
