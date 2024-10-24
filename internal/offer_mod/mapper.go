package offer_mod




func ToOfferModResponse(buyMethod OfferMod) OfferModResponse {
    return OfferModResponse {
		Id: &buyMethod.Id,
    }
}
