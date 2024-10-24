package offer_type




func ToOfferTypeResponse(buyMethod OfferType) OfferTypeResponse {
    return OfferTypeResponse {
		Id: &buyMethod.Id,
    }
}
