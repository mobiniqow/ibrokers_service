package supplier




func ToSupplierResponse(buyMethod Supplier) SupplierResponse {
    return SupplierResponse {
		Id: &buyMethod.Id,
		Customerid: &buyMethod.Customerid,
    }
}
