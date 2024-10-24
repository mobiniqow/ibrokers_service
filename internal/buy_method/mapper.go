package buy_method




func ToBuyMethodResponse(buyMethod BuyMethod) BuyMethodResponse {
    return BuyMethodResponse {
		Id: &buyMethod.Id,
    }
}
