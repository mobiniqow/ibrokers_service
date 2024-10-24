package manufacturers




func ToManufacturersResponse(buyMethod Manufacturers) ManufacturersResponse {
    return ManufacturersResponse {
		Id: &buyMethod.Id,
    }
}
