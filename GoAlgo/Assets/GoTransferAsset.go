    // TRANSFER ASSET

    // Send  10 latinum from Account 1 to Account 3
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

    sender := pks[1]
    recipient := pks[3]
    amount := uint64(10)
    closeRemainderTo := ""
    txn, err = transaction.MakeAssetTransferTxn(sender, recipient, amount, note, txParams, closeRemainderTo, 
        assetID)
    if err != nil {
        fmt.Printf("Failed to send transaction MakeAssetTransfer Txn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[1], txn)
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
    waitForConfirmation(txid,algodClient)

    // print created assetholding for this asset and Account 3 and Account 1
    // You should see amount of 10 in Account 3, and 990 in Account 1
    fmt.Printf("Asset ID: %d\n", assetID)
    fmt.Printf("Account 3: %s\n", pks[3])
    printAssetHolding(assetID, pks[3], algodClient)
    fmt.Printf("Account 1: %s\n", pks[1])
    printAssetHolding(assetID, pks[1], algodClient)