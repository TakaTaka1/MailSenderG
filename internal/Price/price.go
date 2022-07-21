package Price

import (
	"MailSenderG/internal/StructData"
	"strconv"
)

func PrintPrice (price int) int {
	return price
}

func ReturnPrice(price string)int {
	var reTPrice int
	if price != "-" {
		reTPrice , _ = strconv.Atoi(price)
	} else {
		reTPrice = 0
	}
	return reTPrice
}

func GetLastMonthPrice(costMap map[int]StructData.SheetData, lastMonth string, costName string) map[string]StructData.SheetData {		
	var row = make(map[string]StructData.SheetData)
	for _,v := range costMap {		
		if lastMonth == v.Date {
			row[costName] = StructData.SheetData{
				v.Date,
				v.TotalPrice,
				v.TPrice,
				v.MPrice,
				v.Status}
		}
	}
	return row
}