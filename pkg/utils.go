package pkg

import (
    "fmt"
    "log"
    "io/ioutil"
    "bytes"
    "os"
    "time"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

)


const GnosisContractAddress = "0xAddress"
const ArbitrumContractAddress = "0xAddress"
const variableName = "payementMap"
const abiPath = "paymentContractABI.json"
const nodeUrl = ""
func CheckMonthlyPayment(walletAddress string, network string) bool {
	// Connect to an Ethereum node
	client, err := ethclient.Dial(nodeUrl) // Update URL based on your node provider
	if err != nil {
		log.Fatal(err)
	}

	// Get contract ABI
	//contractAbi, err := readContractAbi() // Update path to your ABI file
	if err != nil {
		log.Fatal(err)
	}

	switch contractAddress := ""; network {
	case "arbitrum":
        contractAddress = ArbitrumContractAddress
	case "gnosis":
        contractAddress = GnosisContractAddress
	default:
        contractAddress = ArbitrumContractAddress
        fmt.Println(contractAddress)
	}
	// Create contract instance
    contractInstance, err := NewYourContract(common.HexToAddress(contractAddress), client)
    
    if err != nil {
        log.Fatal(err)
    }    
    value, err := GetVal(walletAddress)
    if err != nil {
        log.Fatal(err)
    }
    
    return isMoreThanMonthOld(int64(value))
}

func readContractAbi() ([]byte, error) {
	// Replace with your logic to read the ABI from a file or remote location
	return nil, fmt.Errorf("implement logic to read ABI from %s", path)
}

func isMoreThanMonthOld(timestamp int64) bool {
    // Convert the timestamp to a time.Time object
    t := time.Unix(timestamp, 0)

    // Get the current time
    currentTime := time.Now()

    // Subtract a month from the current time
    oneMonthAgo := currentTime.AddDate(0, -1, 0)

    // Check if the timestamp is older than one month
    return t.Before(oneMonthAgo)
}

func NewMyContract(address common.Address, client *ethclient.Client) (MyContract, error) {
    // Instantiate the Go binding with the address and the Ethereum client
    // Example:
     contract, err := NewMyContract(contractAddress, client)
     return contract, err
}

func GetVal(walletAddress string) (int64, error) {
    value, err := contractInstance.Get(&bind.CallOpts{}, walletAddress)
    return value, err
}

func ExecuteCommand(command string) (string, error) {
	// Create a new command with the specified command and arguments
	cmd := exec.Command(command, args...)

	// Create a buffer to store the command output
	var outBuffer bytes.Buffer

	// Set the buffer to store the command output
	cmd.Stdout = &outBuffer

	// Run the command
	err := cmd.Run()
	if err != nil {
		// If an error occurred, return the error
		return "", err
	}

	// If the command was successful, return the output as a string
	return outBuffer.String(), nil
}

func WriteStringToFile(filename, data string) error {
	// Write the string data to the file with the specified filename
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		// If an error occurred while writing to the file, return the error
		return err
	}
	// Return nil if the write operation was successful
	return nil
}
