    // OPT-IN

    // Account 3 opts in to receive latinum
    // Use previously set transaction parameters and update sending address to account 3
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

    txn, err = transaction.MakeAssetAcceptanceTxn(pks[3], note, txParams, assetID)
    if err != nil {
        fmt.Printf("Failed to send transaction MakeAssetAcceptanceTxn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[3], txn)
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
    waitForConfirmation(txid, algodClient)

    // print created assetholding for this asset and Account 3, showing 0 balance
    fmt.Printf("Asset ID: %d\n", assetID)
    fmt.Printf("Account 3: %s\n", pks[3])
    printAssetHolding(assetID, pks[3], algodClient)