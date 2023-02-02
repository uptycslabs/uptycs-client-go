package uptycs

import "errors"

func (T Exception) GetID() string {
	return T.ID
}

func (T Exception) GetName() string {
	return T.Name
}

func (T Exception) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateException(exception Exception) (Exception, error) {
	return doCreate(c, exception, "exceptions", []string{})
}

func (c *Client) UpdateException(exception Exception) (Exception, error) {
	return doUpdate(c, exception, "exceptions", []string{})
}

func (c *Client) GetExceptions() (Exceptions, error) {
	return doGetMany(c, Exceptions{}, "exceptions")
}

func (c *Client) GetException(exception Exception) (Exception, error) {
	if len(exception.ID) == 0 {
		all, _ := c.GetExceptions()
		for _, _item := range all.Items {
			if _item.Name == exception.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, exception, "exceptions")
	}
	return exception, errors.New("exception was not found")
}

func (c *Client) DeleteException(exception Exception) (Exception, error) {
	return doDelete(c, exception, "exceptions")
}
