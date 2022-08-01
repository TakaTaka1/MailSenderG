package Usecase

import (
	_"fmt"
	"MailSenderG/data/StructData"
	"MailSenderG/infrastructure"	
)

type sheetService struct {
	SheetRepo *infrastructure.SheetRepository
}

func NewSheetService(repos *infrastructure.SheetRepository) *sheetService {
	return &sheetService{
		SheetRepo: repos,
	}	
}

func (s *sheetService) Read(sheetNameRange string) map[int]StructData.SheetData{
	return s.SheetRepo.RetSheetData(sheetNameRange)
}

func (s *sheetService) SetCost(sheetName string)string{
	costRange := sheetName + "!A2:E13"
	return costRange
}