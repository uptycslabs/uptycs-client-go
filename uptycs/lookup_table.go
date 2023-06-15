package uptycs

import (
	"errors"
	"fmt"
)

func (T LookupTable) GetID() string {
	return T.ID
}

func (T LookupTable) GetName() string {
	return T.Name
}

func (T LookupTable) KeysToDelete() []string {
	return []string{
		"enabled",
		"active",
		"dataLookupTable",
		"fetchRowsquery",
		"forRuleEngine",
		"rowCount",
	}
}

func (c *Client) CreateLookupTable(lookupTable LookupTable) (LookupTable, error) {
	return doCreate(c, lookupTable, "lookupTables", []string{})
}

func (c *Client) UpdateLookupTable(lookupTable LookupTable) (LookupTable, error) {
	_, err := doUpdate(c, lookupTable, "lookupTables", []string{})
	if err != nil {
		return LookupTable{}, err
	}

	_lookupTable, err := doGet(c, lookupTable, "lookupTables")
	_lookupTableDataRows, _ := GetAllLookupTableData(c, fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
	_lookupTable.DataRows = append(_lookupTable.DataRows, _lookupTableDataRows...)
	return _lookupTable, err
}

func (c *Client) GetLookupTables() (LookupTables, error) {
	return doGetMany(c, LookupTables{}, "lookupTables")
}

func (c *Client) GetLookupTable(lookupTable LookupTable) (LookupTable, error) {
	if len(lookupTable.ID) == 0 {
		all, _ := c.GetLookupTables()
		for _, _item := range all.Items {
			if _item.Name == lookupTable.Name {
				_lookupTableDataRows, _ := GetAllLookupTableData(c, fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
				_item.DataRows = append(_item.DataRows, _lookupTableDataRows...)
				return _item, nil
			}
		}
	} else {
		_lookupTable, err := doGet(c, lookupTable, "lookupTables")
		_lookupTableDataRows, _ := GetAllLookupTableData(c, fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
		_lookupTable.DataRows = append(_lookupTable.DataRows, _lookupTableDataRows...)
		return _lookupTable, err
	}
	return lookupTable, errors.New("lookupTable was not found")
}

func (c *Client) DeleteLookupTable(lookupTable LookupTable) (LookupTable, error) {
	return doDelete(c, lookupTable, "lookupTables")
}
