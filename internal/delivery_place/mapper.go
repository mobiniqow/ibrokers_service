package delivery_place




func ToDeliveryPlaceResponse(buyMethod DeliveryPlace) DeliveryPlaceResponse {
    return DeliveryPlaceResponse {
		Id: &buyMethod.Id,
    }
}
