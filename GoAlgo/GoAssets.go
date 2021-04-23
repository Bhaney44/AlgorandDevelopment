addr := "BH55E5RMBD4GYWXGX5W5PJ5JAHPGM5OXKDQH5DC4O2MGI7NW4H6VOE4CP4" // the account issuing the transaction; the asset creator
fee := types.MicroAlgos(10) // the number of microAlgos per byte to pay as a transaction fee
defaultFrozen := false // whether user accounts will need to be unfrozen before transacting
genesisHash, _ := base64.StdEncoding.DecodeString("SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=") // hash of the genesis block of the network to be used
totalIssuance := uint64(100) // total number of this asset in circulation
decimals := uint64(1) // hint to GUIs for interpreting base unit
reserve := addr // specified address is considered the asset reserve (it has no special privileges, this is only informational)
freeze := addr // specified address can freeze or unfreeze user asset holdings
clawback := addr // specified address can revoke user asset holdings and send them to other addresses
manager := addr // specified address can change reserve, freeze, clawback, and manager
unitName := "tst" // used to display asset units to user
assetName := "testcoin" // "friendly name" of asset
genesisID := "" // like genesisHash this is used to specify network to be used
firstRound := types.Round(322575) // first Algorand round on which this transaction is valid
lastRound := types.Round(322575) // last Algorand round on which this transaction is valid
note := nil // arbitrary data to be stored in the transaction; here, none is stored
assetURL := "http://someurl" // optional string pointing to a URL relating to the asset 
assetMetadataHash := "thisIsSomeLength32HashCommitment" // optional hash commitment of some sort relating to the asset. 32 character length.

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