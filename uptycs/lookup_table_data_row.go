package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (T LookupTableDataRow) GetID() string {
	return T.ID
}

func (T LookupTableDataRow) GetName() string {
	return T.Name
}

func (T LookupTableDataRow) KeysToDelete() []string {
	return []string{}
}

func GetAllLookupTableData(c *Client, endpointStr string) ([]LookupTableDataRow, error) {
	urlStr := fmt.Sprintf("%s/%s", c.HostURL, endpointStr)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return []LookupTableDataRow{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return []LookupTableDataRow{}, err
	}

	foundItems := make([]LookupTableDataRow, 1)

	err = json.Unmarshal(body, &foundItems)
	if err != nil {
		return []LookupTableDataRow{}, err
	}

	return foundItems, nil
}

func (c *Client) FindLookupTableDataRow(lookupTable LookupTable, lookupTableDataRow LookupTableDataRow) (LookupTableDataRow, error) {
	tableDataRows, err := GetAllLookupTableData(c, fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
	for _, row := range tableDataRows {
		if row.IDFieldValue == lookupTableDataRow.IDFieldValue {
			return row, err
		}
	}
	return LookupTableDataRow{}, err
}

func (c *Client) CreateLookupTableDataRow(lookupTable LookupTable, lookupTableDataRow LookupTableDataRow) (LookupTableDataRow, error) {
	return lookupTableDataRow, doCreateRaw(c, string(lookupTableDataRow.Data), fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
}

func (c *Client) UpdateLookupTableDataRow(lookupTable LookupTable, lookupTableDataRow LookupTableDataRow) (LookupTableDataRow, error) {
	_lookupTableDataRow, err := c.FindLookupTableDataRow(lookupTable, lookupTableDataRow)
	if err != nil {
		return LookupTableDataRow{}, err
	}
	_, err = doDelete(c, lookupTableDataRow, fmt.Sprintf("lookupTables/%s/data/%s", lookupTable.ID, _lookupTableDataRow.ID))
	if err != nil {
		return LookupTableDataRow{}, err
	}

	_ = doCreateRaw(c, string(lookupTableDataRow.Data), fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
	_newLookupTableDataRow, err := c.FindLookupTableDataRow(lookupTable, lookupTableDataRow)
	return _newLookupTableDataRow, err
}

func (c *Client) GetLookupTableDataRow(lookupTable LookupTable, lookupTableDataRow LookupTableDataRow) (LookupTableDataRow, error) {
	return doGet(c, lookupTableDataRow, fmt.Sprintf("lookupTables/%s/data", lookupTable.ID))
}

func (c *Client) DeleteLookupTableDataRow(lookupTable LookupTable, lookupTableDataRow LookupTableDataRow) (LookupTableDataRow, error) {
	_lookupTableDataRow, err := c.FindLookupTableDataRow(lookupTable, lookupTableDataRow)
	if err != nil {
		return LookupTableDataRow{}, err
	}
	_, err = doDelete(c, lookupTableDataRow, fmt.Sprintf("lookupTables/%s/data/%s", lookupTable.ID, _lookupTableDataRow.ID))
	return LookupTableDataRow{}, err
}
