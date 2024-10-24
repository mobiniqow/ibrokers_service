package currency_unit




func ToCurrencyUnitResponse(buyMethod CurrencyUnit) CurrencyUnitResponse {
    return CurrencyUnitResponse {
		Id: &buyMethod.Id,
    }
}
