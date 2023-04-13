package uptycs

import "errors"

func (T QueryJob) GetID() string {
	return T.ID
}

func (T QueryJob) GetName() string {
	return T.Name
}

func (T QueryJob) KeysToDelete() []string {
	return []string{
		"enabled",
	}
}

func (c *Client) CreateQueryJob(queryJob QueryJob) (QueryJob, error) {
	return doCreate(c, queryJob, "queryJobs", []string{})
}

func (c *Client) UpdateQueryJob(queryJob QueryJob) (QueryJob, error) {
	return QueryJob{}, errors.New("UPDATE is not supported for query jobs")
}

func (c *Client) GetQueryJobs() (QueryJobs, error) {
	return doGetMany(c, QueryJobs{}, "queryTables")
}

func (c *Client) GetQueryJob(queryJob QueryJob) (QueryJob, error) {
	if len(queryJob.ID) == 0 {
		all, _ := c.GetQueryJobs()
		for _, _item := range all.Items {
			if _item.Name == queryJob.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, queryJob, "queryJobs")
	}
	return queryJob, errors.New("queryJob was not found")
}

func (c *Client) DeleteQueryJob(queryJob QueryJob) (QueryJob, error) {
	return doDelete(c, queryJob, "queryJobs")
}
