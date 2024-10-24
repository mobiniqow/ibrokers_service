package trading_hall




func ToTradingHallResponse(buyMethod TradingHall) TradingHallResponse {
    return TradingHallResponse {
		Id: &buyMethod.Id,
    }
}
