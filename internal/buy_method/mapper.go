package buy_method




func ToBuyMethodResponse(buyMethod BuyMethod) BuyMethodResponse {
    return BuyMethodResponse {
		Id: &buyMethod.Id,
		Description: &buyMethod.Description,
		Persianname: &buyMethod.Persianname,
    }
}
