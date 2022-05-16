package psql

import (
	"fmt"

	"github.com/devstackq/bazar/internal/models"
)

func prepareQuery(keys *models.QueryParams) (query string) {
	tempSortKey := ""
	idx := 0

	if len(keys.Filter) > 0 {
		for key, val := range keys.Filter {
			if val != "" {
				if idx >= 1 && idx < len(keys.Filter)-1 {
					query += " AND "
				}
				if key == "yearFrom" {
					query += fmt.Sprintf(" %s", "year >= "+val)
				} else if key == "yearTo" {
					query += fmt.Sprintf(" %s", "year <= "+val)
				} else if key == "priceTo" {
					// filter by price
					query += fmt.Sprintf(" %s", "price <= "+val)
				} else if key == "priceFrom" {
					query += fmt.Sprintf(" %s", "price >= "+val)
				} else {
					if idx < len(keys.Filter)-1 {
						query += fmt.Sprintf(" mch.%s", key+"_id= "+val)
					} else {
						query += fmt.Sprintf(" %s", key+"_id= "+val)
					}
				}
				idx++
			}
		}
	}

	if len(keys.Sort) > 0 {
		for key, val := range keys.Sort {
			if val != "" {
				tempSortKey = fmt.Sprintf(" ORDER BY mch.%s ", key[5:]+" "+val+" ")
			}
		}
	}

	// case if only sort
	if query != "" {
		query = "WHERE " + query
	}
	//default desc
	if tempSortKey == "" {
		query += " ORDER BY mch.created_at DESC "
	}
	// add sort with filter

	query += tempSortKey

	return query
}
