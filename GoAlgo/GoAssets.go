

// the account issuing the transaction; the asset creator
// change address
addr := "BH55E5RMBD4GYWXGX5W5PJ5JAHPGM5OXKDQH5DC4O2MGI7NW4H6VOE4CP4" 

// the number of microAlgos per byte to pay as a transaction fee
// fees
// assets as transaction fees
// try 100 MicroAlgos
fee := types.MicroAlgos(10) 

// whether user accounts will need to be unfrozen before transacting
defaultFrozen := false 

// hash of the genesis block of the network to be used
// genesis hash
genesisHash, _ := base64.StdEncoding.DecodeString("SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=") 

// total number of this asset in circulation
// total asset in circulation
// try 10
totalIssuance := uint64(100) 

// hint to GUIs for interpreting base unit
// base unit
decimals := uint64(1) 

// specified address is considered the asset reserve 
reserve := addr 

// specified address can freeze or unfreeze user asset holdings
// freeze
freeze := addr 

// specified address can revoke user asset holdings and send them to other addresses
// clawback
// allows manager to reverse transactions
clawback := addr 

// specified address can change reserve, freeze, clawback, and manager
manager := addr 

// used to display asset units to user
unitName := "tst" 

// "friendly name" of asset
assetName := "testcoin" 

// like genesisHash this is used to specify network to be used
// genesis ID 
genesisID := "" 

// first Algorand round on which this transaction is valid
firstRound := types.Round(322575) 

// last Algorand round on which this transaction is valid
lastRound := types.Round(322575) 

// arbitrary data to be stored in the transaction; here, none is stored
note := nil 

 // optional string pointing to a URL relating to the asset 
assetURL := "http://someurl"

// optional hash commitment of some sort relating to the asset. 
// 32 character length.
assetMetadataHash := "thisIsSomeLength32HashCommitment" 

// parameters
params := types.SuggestedParams {
	Fee: fee,
	FirstRoundValid: firstRound,
	LastRoundValid: lastRound,
	GenesisHash: genesisHash, 
	GenesisID: genesisID,
}

// signing and sending "txn" allows "addr" to create an asset
txn, err = MakeAssetCreateTxn(addr, note, params,
	totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
	unitName, assetName, assetURL, assetMetadataHash)