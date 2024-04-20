package pkg

type DebugContractBody struct {
	// json tag to de-serialize json body
	Contract string  `json:"contract"`
	WalletAddress string  `json:"walletaddress"`
	Network string  `json:"network"`
}

type GenerateContractBody struct {
	// json tag to de-serialize json body
	Prompt string  `json:"prompt"`
	WalletAddress string  `json:"walletaddress"`
	Network string  `json:"network"`
}

type DebugContractRes struct {
	// json tag to de-serialize json body
	Bugs []string  `json:"bugs"`
}

type GenerateContractRes struct {
	// json tag to de-serialize json body
	Contract string  `json:"contract"`
}
