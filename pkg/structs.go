package pkg

type DebugContractBody struct {
	// json tag to de-serialize json body
	Contract string  `json:"contract"`
}

type GenerateContractBody struct {
	// json tag to de-serialize json body
	Prompt string  `json:"prompt"`
}

type LoginBody struct {
	// json tag to de-serialize json body
	WalletAddress string  `json:"walletaddress"`
}
