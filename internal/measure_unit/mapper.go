package measure_unit




func ToMeasureUnitResponse(buyMethod MeasureUnit) MeasureUnitResponse {
    return MeasureUnitResponse {
		Id: &buyMethod.Id,
    }
}
