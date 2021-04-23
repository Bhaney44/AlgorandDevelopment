// main package
package main

// import
import (
	"fmt"
	"io/ioutil"

	"github.com/algorand/go-algorand-sdk/client/algod"
)

// CHANGE ME
const algodAddress = "http://localhost:8080"
const algodToken = "f1dee49e36a82face92fdb21cd3d340a1b369925cd12f3ee7371378f1665b9b1"

// main
func main() {

	rawTx, err := ioutil.ReadFile("./signed.tx")
	if err != nil {
		fmt.Printf("failed to open signed transaction: %s\n", err)
		return
	}

	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}

	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(rawTx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
}