// Backup Wallet
package main

// import
import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/kmd"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

// These constants represent the kmd REST endpoint and the corresponding API token. 
// You can retrieve these from the `kmd.net` and `kmd.token` files in the kmd data directory.
const kmdAddress = "http://localhost:7833"
const kmdToken = "42b7482737a77d9e5dffb8493ac8899db5f95cbc744d4fcffc0f1c47a6db0c1e"

// main
func main() {
	// Create a kmd client
	kmdClient, err := kmd.MakeClient(kmdAddress, kmdToken)
	if err != nil {
		fmt.Printf("failed to make kmd client: %s\n", err)
		return
	}
	// print
	fmt.Println("Made a kmd client")

	// Get the list of wallets
	listResponse, err := kmdClient.ListWallets()
	if err != nil {
		fmt.Printf("error listing wallets: %s\n", err)
		return
	}

	// Find our wallet name in the list
	var exampleWalletID string
	fmt.Printf("Got %d wallet(s):\n", len(listResponse.Wallets))
	for _, wallet := range listResponse.Wallets {
		fmt.Printf("ID: %s\tName: %s\n", wallet.ID, wallet.Name)
		if wallet.Name == "testwallet" {
			fmt.Printf("found wallet '%s' with ID: %s\n", wallet.Name, wallet.ID)
			exampleWalletID = wallet.ID
		}
	}

	// Get a wallet handle
	initResponse, err := kmdClient.InitWalletHandle(exampleWalletID, "testpassword")
	if err != nil {
		fmt.Printf("Error initializing wallet handle: %s\n", err)
		return
	}

	// Extract the wallet handle
	exampleWalletHandleToken := initResponse.WalletHandleToken

	// Get the backup phrase
	exportResponse, err := kmdClient.ExportMasterDerivationKey(exampleWalletHandleToken, "testpassword")
	if err != nil {
		fmt.Printf("Error exporting backup phrase: %s\n", err)
		return
	}

	// mdk
	mdk := exportResponse.MasterDerivationKey

	// This string should be kept in a safe place and not shared
	stringToSave, err := mnemonic.FromKey(mdk[:])
	if err != nil {
		fmt.Printf("Error getting backup phrase: %s\n", err)
		return
	}

	// print
	fmt.Printf("Backup Phrase: %s\n", stringToSave)
}