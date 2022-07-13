package uptycs

import "errors"

func (T Destination) GetID() string {
	return T.ID
}

func (T Destination) GetName() string {
	return T.Name
}

func (T Destination) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateDestination(destination Destination) (Destination, error) {
	return doCreate(c, destination, "destinations")
}

func (c *Client) UpdateDestination(destination Destination) (Destination, error) {
	return doUpdate(c, destination, "destinations")
}

func (c *Client) GetDestinations() (Destinations, error) {
	return doGetMany(c, Destinations{}, "destinations")
}

func (c *Client) GetDestination(destination Destination) (Destination, error) {
	if len(destination.ID) == 0 {
		all, _ := c.GetDestinations()
		for _, _item := range all.Items {
			if _item.Name == destination.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, destination, "destinations")
	}
	return destination, errors.New("Destination was not found")
}

func (c *Client) DeleteDestination(destination Destination) (Destination, error) {
	return doDelete(c, destination, "destinations")
}
