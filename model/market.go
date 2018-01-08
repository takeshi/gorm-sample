package model

import "github.com/jinzhu/gorm"

type Market struct {
	gorm.Model

	Name string
}

func (this *Market) addAsset(name string) *FinancialAsset {
	asset := FinancialAsset{
		Name: name,
	}
	asset.Market = this
	return &asset
}

type FinancialAsset struct {
	gorm.Model

	Name string

	Market   *Market
	MarketId int
}
