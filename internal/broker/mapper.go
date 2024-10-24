package broker




func ToBrokerResponse(buyMethod Broker) BrokerResponse {
    return BrokerResponse {
		Id: &buyMethod.Id,
		Description: &buyMethod.Description,
		Persianname: &buyMethod.Persianname,
		Spotid: &buyMethod.Spotid,
		Derivativesid: &buyMethod.Derivativesid,
		Nationalid: &buyMethod.Nationalid,
    }
}
