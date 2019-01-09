package lib

import (
	"fmt"
	"github.com/LindsayBradford/go-dbf/godbf"
)

func Count(sources []string) (float64, error) {
	return count(sources)
}

func count(sources []string) (count float64, err error) {
	for _, path := range sources {
		var dbfTable *godbf.DbfTable
		dbfTable, err = godbf.NewFromFile(path, "UTF8")
		if err != nil {
			return
		}
		names := dbfTable.FieldNames()
		for i := 0; i < dbfTable.NumberOfRecords(); i++ {
			for _, name := range names {
				var val string
				val, err = dbfTable.FieldValueByName(i, name)
				fmt.Printf("[%s]: '%s'\n", name, val)
			}
			fmt.Println("--------------------------------")
		}
	}
	return
}
