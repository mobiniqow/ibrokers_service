package offer

import "time";type Offer struct {
    Buymethodid int `form:"buyMethodId"`
    Brokerid int `form:"brokerId"`
    Commodityid int `form:"commodityId"`
    Contracttypeid int `form:"contractTypeId"`
    Currencyid int `form:"currencyId"`
    Deliveryplaceid int `form:"deliveryPlaceId"`
    Initprice int `form:"initPrice"`
    Initvolume string `form:"initVolume"`
    Lotsize int `form:"lotSize"`
    Manufacturerid int `form:"manufacturerId"`
    Maxinitprice int `form:"maxInitPrice"`
    Maxincoffervol int `form:"maxIncOfferVol"`
    Maxordervol int `form:"maxOrderVol"`
    Maxofferprice int `form:"maxOfferPrice"`
    Measureunitid int `form:"measureUnitId"`
    Minallocationvol int `form:"minAllocationVol"`
    Minoffervol int `form:"minOfferVol"`
    Mininitprice int `form:"minInitPrice"`
    Minordervol int `form:"minOrderVol"`
    Minofferprice int `form:"minOfferPrice"`
    Offermodeid int `form:"offerModeId"`
    Offertypeid int `form:"offerTypeId"`
    Offervol int `form:"offerVol"`
    Packagingtypeid int `form:"packagingTypeId"`
    Permissibleerror int `form:"permissibleError"`
    Pricediscoveryminordervol int `form:"priceDiscoveryMinOrderVol"`
    Prepaymentpercent int `form:"prepaymentPercent"`
    Securitytypeid int `form:"securityTypeId"`
    Settlementtypeid int `form:"settlementTypeId"`
    Supplierid int `form:"supplierId"`
    Ticksize int `form:"tickSize"`
    Tradinghallid int `form:"tradingHallId"`
    Weightfactor int `form:"weightFactor"`
    Id            int    `form:"id" gorm:"primary_key"`
    Deliverydate time.Time `form:"deliveryDate"`
    Description string `form:"description"`
    Offerdate time.Time `form:"offerDate"`
    Offerring string `form:"offerRing"`
    Offersymbol string `form:"offerSymbol"`
    Securitytypenote string `form:"securityTypeNote"`
    Tradestatus string `form:"tradeStatus"`
}
