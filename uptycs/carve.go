package uptycs

import "errors"

func (T Carve) GetID() string {
	return T.ID
}

func (T Carve) GetName() string {
	return T.Name
}

func (T Carve) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateCarve(carve Carve) (Carve, error) {
	return doCreate(c, carve, "carves")
}

func (c *Client) UpdateCarve(carve Carve) (Carve, error) {
	return doUpdate(c, carve, "carves")
}

func (c *Client) GetCarves() (Carves, error) {
	return doGetMany(c, Carves{}, "carves")
}

func (c *Client) GetCarve(carve Carve) (Carve, error) {
	if len(carve.ID) == 0 {
		all, _ := c.GetCarves()
		for _, _item := range all.Items {
			if _item.Name == carve.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, carve, "carves")
	}
	return carve, errors.New("carve was not found")
}

func (c *Client) DeleteCarve(carve Carve) (Carve, error) {
	return doDelete(c, carve, "carves")
}
