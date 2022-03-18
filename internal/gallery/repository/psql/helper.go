package psql

import "fmt"

func prepareQuery(keys map[string]string) (query string){
	idx := 0
	for key, val := range keys {
		if val != "" {
			idx++
			if idx >  1 {
			query +=	fmt.Sprintf(" %s", " AND " + key + "_id="+val)
			}else {
			query +=	fmt.Sprintf(" %s",  key + "_id="+val)
			}
		}
	}
	return query
}