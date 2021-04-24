    // CREATE ASSET

    // Construct the transaction
    // Set parameters for asset creation 
    creator := pks[1]
    assetName := "latinum"
    unitName := "latinum"
    assetURL := "https://path/to/my/asset/details"
    assetMetadataHash := "thisIsSomeLength32HashCommitment"
    defaultFrozen := false
    decimals := uint32(0)
    totalIssuance := uint64(1000)
    manager := pks[2]
    reserve := pks[2]
    freeze := pks[2]
    clawback := pks[2]
    note := []byte(nil)
    txn, err := transaction.MakeAssetCreateTxn(creator,
        note,
        txParams, totalIssuance, decimals,
        defaultFrozen, manager, reserve, freeze, clawback,
        unitName, assetName, assetURL, assetMetadataHash)

    if err != nil {
        fmt.Printf("Failed to make asset: %s\n", err)
        return
    }
    fmt.Printf("Asset created AssetName: %s\n", txn.AssetConfigTxnFields.AssetParams.AssetName)
    // sign the transaction
    txid, stx, err := crypto.SignTransaction(sks[1], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err := algodClient.SendRawTransaction(stx).Do(context.Background())
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Submitted transaction %s\n", sendResponse)
    // Wait for transaction to be confirmed
    waitForConfirmation(txid, algodClient)
    //    response := algodClient.PendingTransactionInformation(txid)
    //    prettyPrint(response)
    // Retrieve asset ID by grabbing the max asset ID
    // from the creator account's holdings.
    act, err := algodClient.AccountInformation(pks[1]).Do(context.Background())
    if err != nil {
        fmt.Printf("failed to get account information: %s\n", err)
        return
    }

    assetID := uint64(0)
    //  find newest (highest) asset for this account
    for _, asset := range act.CreatedAssets {
        if asset.Index > assetID {
            assetID = asset.Index
        }
    }

    // print created asset and asset holding info for this asset
    fmt.Printf("Asset ID: %d\n", assetID)
    printCreatedAsset(assetID, pks[1], algodClient)
    printAssetHolding(assetID, pks[1], algodClient)