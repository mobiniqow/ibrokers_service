package offer




func ToOfferResponse(buyMethod Offer) OfferResponse {
    return OfferResponse {
		Buymethodid: &buyMethod.Buymethodid,
		Brokerid: &buyMethod.Brokerid,
		Commodityid: &buyMethod.Commodityid,
		Contracttypeid: &buyMethod.Contracttypeid,
		Currencyid: &buyMethod.Currencyid,
		Deliveryplaceid: &buyMethod.Deliveryplaceid,
		Initprice: &buyMethod.Initprice,
		Lotsize: &buyMethod.Lotsize,
		Manufacturerid: &buyMethod.Manufacturerid,
		Maxinitprice: &buyMethod.Maxinitprice,
		Maxincoffervol: &buyMethod.Maxincoffervol,
		Maxordervol: &buyMethod.Maxordervol,
		Maxofferprice: &buyMethod.Maxofferprice,
		Measureunitid: &buyMethod.Measureunitid,
		Minallocationvol: &buyMethod.Minallocationvol,
		Minoffervol: &buyMethod.Minoffervol,
		Mininitprice: &buyMethod.Mininitprice,
		Minordervol: &buyMethod.Minordervol,
		Minofferprice: &buyMethod.Minofferprice,
		Offermodeid: &buyMethod.Offermodeid,
		Offertypeid: &buyMethod.Offertypeid,
		Offervol: &buyMethod.Offervol,
		Packagingtypeid: &buyMethod.Packagingtypeid,
		Permissibleerror: &buyMethod.Permissibleerror,
		Pricediscoveryminordervol: &buyMethod.Pricediscoveryminordervol,
		Prepaymentpercent: &buyMethod.Prepaymentpercent,
		Securitytypeid: &buyMethod.Securitytypeid,
		Settlementtypeid: &buyMethod.Settlementtypeid,
		Supplierid: &buyMethod.Supplierid,
		Ticksize: &buyMethod.Ticksize,
		Tradinghallid: &buyMethod.Tradinghallid,
		Weightfactor: &buyMethod.Weightfactor,
		Id: &buyMethod.Id,
    }
}
