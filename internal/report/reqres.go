package report



type CreateReportRequest struct {
	Commodityid *int `json:"commodityId"`
	Contracttypeid *int `json:"contractTypeId"`
	Currencyid *int `json:"currencyId"`
	Manufacturerid *int `json:"manufacturerId"`
	Measurementunitid *int `json:"measurementUnitId"`
	Offerid *int `json:"offerId"`
	Sellerbrokerid *int `json:"sellerBrokerId"`
	Supplierid *int `json:"supplierId"`
	Maximumprice *int `json:"maximumPrice"`
	Minimumprice *int `json:"minimumPrice"`
	Offerbaseprice *int `json:"offerBasePrice"`
	Finalweightedaverageprice *int `json:"finalWeightedAveragePrice"`
}

type ReportResponse struct {
    	Commodityid *int `json:"commodityId"`
	Contracttypeid *int `json:"contractTypeId"`
	Currencyid *int `json:"currencyId"`
	Manufacturerid *int `json:"manufacturerId"`
	Measurementunitid *int `json:"measurementUnitId"`
	Offerid *int `json:"offerId"`
	Sellerbrokerid *int `json:"sellerBrokerId"`
	Supplierid *int `json:"supplierId"`
	Maximumprice *int `json:"maximumPrice"`
	Minimumprice *int `json:"minimumPrice"`
	Offerbaseprice *int `json:"offerBasePrice"`
	Finalweightedaverageprice *int `json:"finalWeightedAveragePrice"`
}
