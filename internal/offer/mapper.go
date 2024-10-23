package offer

import (
    "strconv"
    "time"
)

func ToOfferResponse(buyMethod Offer) Response {
    return Response {
		Buymethodid: strconv.Itoa(buyMethod.Buymethodid),
		Brokerid: strconv.Itoa(buyMethod.Brokerid),
		Commodityid: strconv.Itoa(buyMethod.Commodityid),
		Contracttypeid: strconv.Itoa(buyMethod.Contracttypeid),
		Currencyid: strconv.Itoa(buyMethod.Currencyid),
		Deliveryplaceid: strconv.Itoa(buyMethod.Deliveryplaceid),
		Initprice: strconv.Itoa(buyMethod.Initprice),
		Initvolume: buyMethod.Initvolume,
		Lotsize: strconv.Itoa(buyMethod.Lotsize),
		Manufacturerid: strconv.Itoa(buyMethod.Manufacturerid),
		Maxinitprice: strconv.Itoa(buyMethod.Maxinitprice),
		Maxincoffervol: strconv.Itoa(buyMethod.Maxincoffervol),
		Maxordervol: strconv.Itoa(buyMethod.Maxordervol),
		Maxofferprice: strconv.Itoa(buyMethod.Maxofferprice),
		Measureunitid: strconv.Itoa(buyMethod.Measureunitid),
		Minallocationvol: strconv.Itoa(buyMethod.Minallocationvol),
		Minoffervol: strconv.Itoa(buyMethod.Minoffervol),
		Mininitprice: strconv.Itoa(buyMethod.Mininitprice),
		Minordervol: strconv.Itoa(buyMethod.Minordervol),
		Minofferprice: strconv.Itoa(buyMethod.Minofferprice),
		Offermodeid: strconv.Itoa(buyMethod.Offermodeid),
		Offertypeid: strconv.Itoa(buyMethod.Offertypeid),
		Offervol: strconv.Itoa(buyMethod.Offervol),
		Packagingtypeid: strconv.Itoa(buyMethod.Packagingtypeid),
		Permissibleerror: strconv.Itoa(buyMethod.Permissibleerror),
		Pricediscoveryminordervol: strconv.Itoa(buyMethod.Pricediscoveryminordervol),
		Prepaymentpercent: strconv.Itoa(buyMethod.Prepaymentpercent),
		Securitytypeid: strconv.Itoa(buyMethod.Securitytypeid),
		Settlementtypeid: strconv.Itoa(buyMethod.Settlementtypeid),
		Supplierid: strconv.Itoa(buyMethod.Supplierid),
		Ticksize: strconv.Itoa(buyMethod.Ticksize),
		Tradinghallid: strconv.Itoa(buyMethod.Tradinghallid),
		Weightfactor: strconv.Itoa(buyMethod.Weightfactor),
		Id: strconv.Itoa(buyMethod.Id),
		Description: buyMethod.Description,
		Offerring: buyMethod.Offerring,
		Offersymbol: buyMethod.Offersymbol,
		Securitytypenote: buyMethod.Securitytypenote,
		Tradestatus: buyMethod.Tradestatus,
    }
}
