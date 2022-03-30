package psql

import (
	"fmt"
)

func prepareQuery(keys map[string]string) (query string) {
	idx := 0
	tempSortKey := ""

	for key, val := range keys {
		if val != "" {
			idx++
			if key != "sort_created_at" && key != "sort_price" && key != "sort_year" && key != "sort_odometer" {
				if idx > 1 {
					query += " AND "
				}
				// filter by price
				if key == "priceTo" {
					query += fmt.Sprintf(" %s", "price <="+val)
				} else if key == "priceFrom" {
					query += fmt.Sprintf(" %s", "price >= "+val)
				} else {
					query += fmt.Sprintf(" %s", key+"_id="+val)
				}
			} else {
				// add last key, sort
				tempSortKey = fmt.Sprintf("ORDER BY %s ", key[5:]+" "+val)
			}
		}
	}
	//case if only sort
	if query != "" {
		query = "WHERE " + query
	}
	// add sort with filter
	query += tempSortKey

	return query
}
