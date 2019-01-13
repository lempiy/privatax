package lib

import (
	"encoding/json"
	"github.com/LindsayBradford/go-dbf/godbf"
	"strconv"
)

const (
	DateKey = "DATE"
	SumKey = "SUM"
	CurrencyKey = "CUR"
)

func Parse(source []byte) (string, error) {
	return parse(source)
}

func parse(source []byte) (parsed string, err error) {
	var dbfTable *godbf.DbfTable
	dbfTable, err = godbf.NewFromByteArray(source, "IBM866")
	if err != nil {
		return
	}
	var arr []map[string]interface{}
	for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		data := make(map[string]interface{})
		var sum string
		sum, err = dbfTable.FieldValueByName(i, SumKey)
		var f float64
		f, err = strconv.ParseFloat(sum, 64)
		if f <= 0 {
			continue
		}
		names := dbfTable.FieldNames()
		for _, n := range names {
			data[n], err = dbfTable.FieldValueByName(i, n)
			if err != nil {
				return
			}
		}
		arr = append(arr, data)
	}
	bts, err := json.Marshal(arr)
	if err != nil {
		return
	}
	parsed = string(bts)
	return
}
