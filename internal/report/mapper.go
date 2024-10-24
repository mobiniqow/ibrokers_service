package report




func ToReportResponse(buyMethod Report) ReportResponse {
    return ReportResponse {
		Commodityid: &buyMethod.Commodityid,
		Contracttypeid: &buyMethod.Contracttypeid,
		Currencyid: &buyMethod.Currencyid,
		Manufacturerid: &buyMethod.Manufacturerid,
		Measurementunitid: &buyMethod.Measurementunitid,
		Offerid: &buyMethod.Offerid,
		Sellerbrokerid: &buyMethod.Sellerbrokerid,
		Supplierid: &buyMethod.Supplierid,
		Maximumprice: &buyMethod.Maximumprice,
		Minimumprice: &buyMethod.Minimumprice,
		Offerbaseprice: &buyMethod.Offerbaseprice,
		Finalweightedaverageprice: &buyMethod.Finalweightedaverageprice,
    }
}
