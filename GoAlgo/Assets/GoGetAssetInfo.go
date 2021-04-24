    // note: if you have an indexer instance available it is easier to just search accounts for an asset
    // printAssetHolding utility to print asset holding for account
    func printAssetHolding(assetID uint64, account string, client *algod.Client) {

        act, err := client.AccountInformation(account).Do(context.Background())
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        for _, assetholding := range act.Assets {
            if assetID == assetholding.AssetId {
                prettyPrint(assetholding)
                break
            }
        }
    }

    // printCreatedAsset utility to print created assert for account
    func printCreatedAsset(assetID uint64, account string, client *algod.Client) {

        act, err := client.AccountInformation(account).Do(context.Background())
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        for _, asset := range act.CreatedAssets {
            if assetID == asset.Index {
                prettyPrint(asset)
                break
            }
        }
    }
    ...
    printCreatedAsset(assetID, pks[1], algodClient)
    printAssetHolding(assetID, pks[1], algodClient)   