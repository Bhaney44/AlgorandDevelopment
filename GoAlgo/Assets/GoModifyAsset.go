    // CHANGE MANAGER
    // Change Asset Manager from Account 2 to Account 1
    // assetID := uint64(332920)
    // Get network-related transaction parameters and assign
    txParams, err = algodClient.SuggestedParams().Do(context.Background())
    if err != nil {
        fmt.Printf("Error getting suggested tx params: %s\n", err)
        return
    }
    // comment out the next two (2) lines to use suggested fees
    txParams.FlatFee = true
    txParams.Fee = 1000

    manager = pks[1]
    oldmanager := pks[2]
    strictEmptyAddressChecking := true
    txn, err = transaction.MakeAssetConfigTxn(oldmanager, note, txParams, assetID, manager, reserve, freeze, clawback, strictEmptyAddressChecking)
    if err != nil {
        fmt.Printf("Failed to send txn: %s\n", err)
        return
    }

    txid, stx, err = crypto.SignTransaction(sks[2], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID raw: %s\n", txid)

    // Wait for transaction to be confirmed
    waitForConfirmation(txid,algodClient )
    // print created assetinfo for this asset
    fmt.Printf("Asset ID: %d\n", assetID)
    printCreatedAsset(assetID, pks[1], algodClient)