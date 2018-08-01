package hbase2go

func GenTGet(rowKey, TableCF, Qual string) (tget *TGet) {
	tget = &TGet{
		Row:     []byte(rowKey),
		Columns: []*TColumn{},
	}
	if TableCF != "" {
		Tcol := &TColumn{
			Family: []byte(TableCF),
		}
		if Qual != "" {
			Tcol.Qualifier = []byte(Qual)
		}
		tget.Columns = append(tget.Columns, Tcol)
	}
	return
}

func GenTPut(rowKey, tableCF, qualifier string, rowValue []byte) (tput *TPut) {
	tput = &TPut{
		Row: []byte(rowKey),
		ColumnValues: []*TColumnValue{
			&TColumnValue{
				Family:    []byte(tableCF),
				Qualifier: []byte(qualifier),
				Value:     rowValue,
			},
		},
	}
	return
}
