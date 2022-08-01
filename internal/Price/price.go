package Price

import (
	"MailSenderG/data/StructData"
	"strconv"	
)

func ReturnPrice(price string)int {
	reTPrice := 0
	if price != "-" {
		reTPrice , _ = strconv.Atoi(price)
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

func CheckDiffPrice(totalMiPrice, totalTaPrice int) int{
	diffPrice := totalMiPrice - totalTaPrice
	if totalMiPrice < totalTaPrice {
		diffPrice = totalTaPrice - totalMiPrice
	}
	return diffPrice
}

func CheckCost(costTotalPrice string) string{
	if costTotalPrice == "-" {		
		return "光熱費を入力してください"
	}
	return costTotalPrice
}