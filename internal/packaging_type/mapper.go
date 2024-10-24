package packaging_type




func ToPackagingTypeResponse(buyMethod PackagingType) PackagingTypeResponse {
    return PackagingTypeResponse {
		Id: &buyMethod.Id,
    }
}
