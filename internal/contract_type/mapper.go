package contract_type




func ToContractTypeResponse(buyMethod ContractType) ContractTypeResponse {
    return ContractTypeResponse {
		Id: &buyMethod.Id,
    }
}
