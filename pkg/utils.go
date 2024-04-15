package pkg

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func SetupDebugger(){
    fmt.Println("SKRT")
}

const contractAddress = "0xYOUR_CONTRACT_ADDRESS"
const variableName = "myUint"
func CheckMonthlyPayment(walletAddress string) {
	// Connect to an Ethereum node
	client, err := ethclient.Dial( /* "http://localhost:8545" */ "YOUR_NODE_URL") // Update URL based on your node provider
	if err != nil {
		log.Fatal(err)
	}

	// Get contract ABI
	contractAbi, err := readContractAbi( /* Path to your contract ABI file */ ) // Update path to your ABI file
	if err != nil {
		log.Fatal(err)
	}

	// Create contract instance
	contract := common.NewContract(common.HexToAddress(contractAddress), contractAbi)

	// Build call data for the getter function
	var getter []byte
	getter, err = contract.Methods.GetMyUint().EncodeABI() // Replace GetMyUint with your actual getter function name
	if err != nil {
		log.Fatal(err)
	}

	// Call the contract
	result, err := client.CallContract(context.Background(), contract, getter, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the result based on the variable type
	var value interface{}
	switch variableName {
	case "myUint":
		value = new(uint64)
	case "myString":
		value = new(string)
	// Add additional cases for other data types
	default:
		log.Fatal("Unsupported variable type:", variableName)
	}
	err = contract.Methods.GetMyUint().Decode(value, result) // Replace GetMyUint with your actual getter function name
	if err != nil {
		log.Fatal(err)
	}

    return true
}

func readContractAbi(path string) ([]byte, error) {
	// Replace with your logic to read the ABI from a file or remote location
	return nil, fmt.Errorf("implement logic to read ABI from %s", path)
}

