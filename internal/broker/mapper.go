package broker




func ToBrokerResponse(buyMethod Broker) BrokerResponse {
    return BrokerResponse {
		Id: &buyMethod.Id,
		Spotid: &buyMethod.Spotid,
		Derivativesid: &buyMethod.Derivativesid,
    }
}
