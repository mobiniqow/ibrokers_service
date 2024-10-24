package offer



type CreateOfferRequest struct {
	Buymethodid *int `json:"buyMethodId"`
	Brokerid *int `json:"brokerId"`
	Commodityid *int `json:"commodityId"`
	Contracttypeid *int `json:"contractTypeId"`
	Currencyid *int `json:"currencyId"`
	Deliveryplaceid *int `json:"deliveryPlaceId"`
	Initprice *int `json:"initPrice"`
	Lotsize *int `json:"lotSize"`
	Manufacturerid *int `json:"manufacturerId"`
	Maxinitprice *int `json:"maxInitPrice"`
	Maxincoffervol *int `json:"maxIncOfferVol"`
	Maxordervol *int `json:"maxOrderVol"`
	Maxofferprice *int `json:"maxOfferPrice"`
	Measureunitid *int `json:"measureUnitId"`
	Minallocationvol *int `json:"minAllocationVol"`
	Minoffervol *int `json:"minOfferVol"`
	Mininitprice *int `json:"minInitPrice"`
	Minordervol *int `json:"minOrderVol"`
	Minofferprice *int `json:"minOfferPrice"`
	Offermodeid *int `json:"offerModeId"`
	Offertypeid *int `json:"offerTypeId"`
	Offervol *int `json:"offerVol"`
	Packagingtypeid *int `json:"packagingTypeId"`
	Permissibleerror *int `json:"permissibleError"`
	Pricediscoveryminordervol *int `json:"priceDiscoveryMinOrderVol"`
	Prepaymentpercent *int `json:"prepaymentPercent"`
	Securitytypeid *int `json:"securityTypeId"`
	Settlementtypeid *int `json:"settlementTypeId"`
	Supplierid *int `json:"supplierId"`
	Ticksize *int `json:"tickSize"`
	Tradinghallid *int `json:"tradingHallId"`
	Weightfactor *int `json:"weightFactor"`
	Id *int `json:"id"`
}

type OfferResponse struct {
    	Buymethodid *int `json:"buyMethodId"`
	Brokerid *int `json:"brokerId"`
	Commodityid *int `json:"commodityId"`
	Contracttypeid *int `json:"contractTypeId"`
	Currencyid *int `json:"currencyId"`
	Deliveryplaceid *int `json:"deliveryPlaceId"`
	Initprice *int `json:"initPrice"`
	Lotsize *int `json:"lotSize"`
	Manufacturerid *int `json:"manufacturerId"`
	Maxinitprice *int `json:"maxInitPrice"`
	Maxincoffervol *int `json:"maxIncOfferVol"`
	Maxordervol *int `json:"maxOrderVol"`
	Maxofferprice *int `json:"maxOfferPrice"`
	Measureunitid *int `json:"measureUnitId"`
	Minallocationvol *int `json:"minAllocationVol"`
	Minoffervol *int `json:"minOfferVol"`
	Mininitprice *int `json:"minInitPrice"`
	Minordervol *int `json:"minOrderVol"`
	Minofferprice *int `json:"minOfferPrice"`
	Offermodeid *int `json:"offerModeId"`
	Offertypeid *int `json:"offerTypeId"`
	Offervol *int `json:"offerVol"`
	Packagingtypeid *int `json:"packagingTypeId"`
	Permissibleerror *int `json:"permissibleError"`
	Pricediscoveryminordervol *int `json:"priceDiscoveryMinOrderVol"`
	Prepaymentpercent *int `json:"prepaymentPercent"`
	Securitytypeid *int `json:"securityTypeId"`
	Settlementtypeid *int `json:"settlementTypeId"`
	Supplierid *int `json:"supplierId"`
	Ticksize *int `json:"tickSize"`
	Tradinghallid *int `json:"tradingHallId"`
	Weightfactor *int `json:"weightFactor"`
	Id *int `json:"id"`
}
