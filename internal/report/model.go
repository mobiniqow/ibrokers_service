package report

import "time"

type Report struct {
	Id                        int       `gorm:"primary_key"`
	Commodityid               int       `json:"commodityId"`
	Contracttypeid            int       `json:"contractTypeId"`
	Currencyid                int       `json:"currencyId"`
	Manufacturerid            int       `json:"manufacturerId"`
	Measurementunitid         int       `json:"measurementUnitId"`
	Offerid                   int       `json:"offerId"`
	Sellerbrokerid            int       `json:"sellerBrokerId"`
	Supplierid                int       `json:"supplierId"`
	Demandvolume              string    `json:"demandVolume"`
	Maximumprice              int       `json:"maximumPrice"`
	Minimumprice              int       `json:"minimumPrice"`
	Offerbaseprice            int       `json:"offerBasePrice"`
	Offervolume               string    `json:"offerVolume"`
	Finalweightedaverageprice int       `json:"finalWeightedAveragePrice"`
	Buymethod                 string    `json:"buyMethod"`
	Duedate                   time.Time `json:"dueDate"`
	Offermode                 string    `json:"offerMode"`
	Offersymbol               string    `json:"offerSymbol"`
	Offertype                 string    `json:"offerType"`
	Tradedate                 time.Time `json:"tradeDate"`
	Tradevalue                string    `json:"tradeValue"`
	Tradevolume               string    `json:"tradeVolume"`
}
