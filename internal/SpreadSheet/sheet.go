package SpreadSheet

import (
	_"fmt"	
)

type SheetCostName struct {
	FoodCost string
	HouseHoldCost string
	GeneralCost string
	WaterCost string
	GasCost string
	RentCost string
}

func RetSheetNameStruct () *SheetCostName{
	return &SheetCostName {
		FoodCost : "食費",
		HouseHoldCost : "日用品",
		GeneralCost : "雑費",
		WaterCost : "水道費",
		GasCost : "光熱費",
		RentCost : "家賃",
	}
}