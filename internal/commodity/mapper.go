package commodity




func ToCommodityResponse(buyMethod Commodity) CommodityResponse {
    return CommodityResponse {
		Id: &buyMethod.Id,
		Parentid: &buyMethod.Parentid,
    }
}
