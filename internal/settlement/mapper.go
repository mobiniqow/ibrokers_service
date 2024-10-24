package settlement




func ToSettlementResponse(buyMethod Settlement) SettlementResponse {
    return SettlementResponse {
		Id: &buyMethod.Id,
    }
}
