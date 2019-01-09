package lib

import (
	"fmt"
	"github.com/LindsayBradford/go-dbf/godbf"
	"strconv"
)

const (
	DateKey = "DATE"
	SumKey = "SUM"
	CurrencyKey = "CUR"
)

func Count(source []byte) (float64, error) {
	return count(source)
}

func count(source []byte) (count float64, err error) {
	var dbfTable *godbf.DbfTable
	dbfTable, err = godbf.NewFromByteArray(source, "IBM866")
	if err != nil {
		return
	}
	names := dbfTable.FieldNames()
	for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		var sum string
		sum, err = dbfTable.FieldValueByName(i, SumKey)
		var f float64
		f, err = strconv.ParseFloat(sum, 64)
		if f <= 0 {
			continue
		}
		for _, name := range names {
			var val string
			val, err = dbfTable.FieldValueByName(i, name)
			fmt.Printf("[%s]: '%s'\n", name, val)
		}
		var date string
		date, err = dbfTable.FieldValueByName(i, DateKey)

		var curr string
		curr, err = dbfTable.FieldValueByName(i, CurrencyKey)
		fmt.Printf("[%s] %s %s\n", date, sum, curr)
		fmt.Println("--------------------------------")
	}
	return
}
