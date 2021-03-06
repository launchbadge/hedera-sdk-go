package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type contract struct {
	// ignore the link references since it is empty
	Object    string `json:"object"`
	OpCodes   string `json:"opcodes"`
	SourceMap string `json:"sourceMap"`
}

func main() {
	var client *hedera.Client
	var err error

	if os.Getenv("HEDERA_NETWORK") == "previewnet" {
		client = hedera.ClientForPreviewnet()
	} else {
		client, err = hedera.ClientFromConfigFile(os.Getenv("CONFIG_FILE"))

		if err != nil {
			client = hedera.ClientForTestnet()
		}
	}

	configOperatorID := os.Getenv("OPERATOR_ID")
	configOperatorKey := os.Getenv("OPERATOR_KEY")
	var operatorKey hedera.PrivateKey

	if configOperatorID != "" && configOperatorKey != "" && client.GetOperatorPublicKey().Bytes() == nil {
		operatorAccountID, err := hedera.AccountIDFromString(configOperatorID)
		if err != nil {
			println(err.Error(), ": error converting string to AccountID")
			return
		}

		operatorKey, err = hedera.PrivateKeyFromString(configOperatorKey)
		if err != nil {
			println(err.Error(), ": error converting string to PrivateKey")
			return
		}

		client.SetOperator(operatorAccountID, operatorKey)
	}

	defer func() {
		err = client.Close()
		if err != nil {
			println(err.Error(), ": error closing client")
			return
		}
	}()

	rawContract, err := ioutil.ReadFile("./hello_world.json")
	if err != nil {
		println(err.Error(), ": error reading hello_world.json")
		return
	}

	var contract contract = contract{}

	err = json.Unmarshal([]byte(rawContract), &contract)
	if err != nil {
		println(err.Error(), ": error unmarshaling the json file")
		return
	}

	contractByteCode := []byte(contract.Object)

	fmt.Println("Simple contract example")
	fmt.Printf("Contract bytecode size: %v bytes\n", len(contractByteCode))

	// Upload a file containing the byte code
	byteCodeTransactionID, err := hedera.NewFileCreateTransaction().
		SetMaxTransactionFee(hedera.NewHbar(2)).
		SetKeys(client.GetOperatorPublicKey()).
		SetContents(contractByteCode).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error creating file")
		return
	}

	byteCodeTransactionRecord, err := byteCodeTransactionID.GetRecord(client)
	if err != nil {
		println(err.Error(), ": error getting file creation record")
		return
	}

	fmt.Printf("contract bytecode file upload fee: %v\n", byteCodeTransactionRecord.TransactionFee)

	byteCodeFileID := *byteCodeTransactionRecord.Receipt.FileID

	fmt.Printf("contract bytecode file: %v\n", byteCodeFileID)

	// Instantiate the contract instance
	contractTransactionResponse, err := hedera.NewContractCreateTransaction().
		// Failing to set this to a sufficient amount will result in "INSUFFICIENT_GAS" status
		SetGas(2000).
		SetBytecodeFileID(byteCodeFileID).
		// Setting an admin key allows you to delete the contract in the future
		SetAdminKey(client.GetOperatorPublicKey()).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error creating contract")
		return
	}

	contractRecord, err := contractTransactionResponse.GetRecord(client)
	if err != nil {
		println(err.Error(), ": error retrieving contract creation record")
		return
	}

	contractCreateResult, err := contractRecord.GetContractCreateResult()
	if err != nil {
		println(err.Error(), ": error retrieving contract creation result")
		return
	}

	newContractID := *contractRecord.Receipt.ContractID

	fmt.Printf("Contract create gas used: %v\n", contractCreateResult.GasUsed)
	fmt.Printf("Contract create transaction fee: %v\n", contractRecord.TransactionFee)
	fmt.Printf("Contract: %v\n", newContractID)

	// Call the contract to receive the greeting
	callResult, err := hedera.NewContractCallQuery().
		SetContractID(newContractID).
		SetGas(30000).
		SetFunction("greet", nil).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error executing contract call query")
		return
	}

	fmt.Printf("Call gas used: %v\n", callResult.GasUsed)
	fmt.Printf("Message: %v\n", callResult.GetString(0))

	// delete the transaction
	deleteTransactionResponse, err := hedera.NewContractDeleteTransaction().
		SetContractID(newContractID).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error deleting contract")
		return
	}

	deleteTransactionReceipt, err := deleteTransactionResponse.GetReceipt(client)
	if err != nil {
		println(err.Error(), ": error retrieving contract delete receipt")
		return
	}

	fmt.Printf("Status of transaction deletion: %v\n", deleteTransactionReceipt.Status)
}
