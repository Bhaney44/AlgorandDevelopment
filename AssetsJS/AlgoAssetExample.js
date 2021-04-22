// Example of asset creation and transfer using Algo
// GitHub https://github.com/algorand/js-algorand-sdk
// Credit @ryanRFox: https://github.com/algorand/docs/blob/master/examples/assets/v2/javascript/AssetExample.js

const algosdk = require('algosdk');

// Retrieve the token, server and port values for your installation in the 
// algod.net and algod.token files within the data directory

// UPDATE THESE VALUES
// const token = "TOKEN";
// const server = "SERVER";
// const port = PORT;

// SANDBOX SERVER
// const token = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1";
// const server = "http://hackathon.algodev.network";
// const port = 9100;

const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;


// Function for a tx confirmation
const waitForConfirmation = async function (algodclient, txId) {
    let response = await algodclient.status().do();
    let lastround = response["last-round"];
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
        if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
            console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
            break;
        }
        lastround++;
        await algodclient.statusAfterBlock(lastround).do();
    }
};

// Function to print created asset
const printCreatedAsset = async function (algodclient, account, assetid) {
    let accountInfo = await algodclient.accountInformation(account).do();
    for (idx = 0; idx < accountInfo['created-assets'].length; idx++) {
        let scrutinizedAsset = accountInfo['created-assets'][idx];
        if (scrutinizedAsset['index'] == assetid) {
            console.log("AssetID = " + scrutinizedAsset['index']);
            let myparms = JSON.stringify(scrutinizedAsset['params'], undefined, 2);
            console.log("parms = " + myparms);
            break;
        }
    }
};

// Function used to print asset holding
const printAssetHolding = async function (algodclient, account, assetid) {
    let accountInfo = await algodclient.accountInformation(account).do();
    for (idx = 0; idx < accountInfo['assets'].length; idx++) {
        let scrutinizedAsset = accountInfo['assets'][idx];
        if (scrutinizedAsset['asset-id'] == assetid) {
            let myassetholding = JSON.stringify(scrutinizedAsset, undefined, 2);
            console.log("assetholdinginfo = " + myassetholding);
            break;
        }
    }
};

// Recover accounts
// Use mnemonic phrases here for each account
// var accountN_mnemonic = "PASTE your phrase for account 1";

var account1_mnemonic = "canal enact luggage language spring similar zoo couple stomach shoe laptop middle wonder eager monitor weather number heavy skirt siren purity spell maze warfare ability eleven";
var account2_mnemonic = "beauty nurse season slip slice cry strategy frozen spy panic hobby strong goose employ review free pride enlist friend enroll clip ability runway";
var account3_mnemonic = "picnic bright know ticket purity pluck prudent stumble destroy quote frame wealth money drift cinnamon resemble shrimp grain dynamic absorb edge";

var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
console.log(recoveredAccount1.addr);
console.log(recoveredAccount2.addr);
console.log(recoveredAccount3.addr);

// Instantiate the algod wrapper
    let algodclient = new algosdk.Algodv2(token, server, port);

// Asset Creation
// The first transaciton is to create a new asset
// Recieve last round and suggested tx fee
// These parameters will be required before every transaction

(async () => {

    let params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    params.fee = 1000;
    // flat fee parameter
    params.flatFee = true;
    // log parameter
    console.log(params);
    // Arbitrary data to be stored in the transaction
    let note = undefined; 
    // Asset creation specific parameters
    let addr = recoveredAccount1.addr;
    // Whether user accounts will need to be unfrozen before transacting    
    let defaultFrozen = false;
    // Integer number of decimals for asset unit calculation
    let decimals = 0;
    // Total number of this asset available for circulation   
    let totalIssuance = 1000;
    // Used to display asset units to user    
    let unitName = "Lions";
    // Ticker name of the asset    
    let assetName = "LEO";
    // Optional string pointing to a URL relating to the asset
    let assetURL = "https://github.com/LionsRoarRealLoud";

    // Optional hash commitment relating to the asset with 32 character length.
    let assetMetadataHash = "2010wpialchamps44and63ontheend38";

    // The following parameters must be changed
    // Specified address can change reserve, freeze, clawback, and manager
    let manager = recoveredAccount2.addr;
    // Specified address is the asset reserve
    let reserve = recoveredAccount2.addr;
    // Specified address can freeze or unfreeze user asset holdings 
    let freeze = recoveredAccount2.addr;
    // Specified address can revoke or send user asset holdings    
    let clawback = recoveredAccount2.addr;

    // Signing and sending "txn" allows "addr" to create an asset
    let txn = algosdk.makeAssetCreateTxnWithSuggestedParams(addr, note,
         totalIssuance, decimals, defaultFrozen, manager, reserve, freeze,
        clawback, unitName, assetName, assetURL, assetMetadataHash, params);
    let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + tx.txId);
    let assetID = null;

    // Wait for transaction to be confirmed
    await waitForConfirmation(algodclient, tx.txId);
    // Get the new asset's information from the creator account
    let ptx = await algodclient.pendingTransactionInformation(tx.txId).do();
    assetID = ptx["asset-index"];
   // console.log("AssetID = " + assetID);
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);

    
    // Change Asset Configuration
    // Change the manager using an asset configuration transaction
    // First update changing transaction parameters
    params = await algodclient.getTransactionParams().do();
    //Comment out the next two lines to use suggested fee
    //Suggested fee 1000
    params.fee = 1000;
    //Flate rate param
    params.flatFee = true;

    // Asset configuration specific parameters
    manager = recoveredAccount1.addr;
    let ctxn = algosdk.makeAssetConfigTxnWithSuggestedParams(recoveredAccount2.addr, note, 
        assetID, manager, reserve, freeze, clawback, params);

    // This transaction must be signed by the current manager
    rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
    let ctx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + ctx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ctx.txId);

    // Get the asset information using indexer
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
  
    // Opting in to an Asset

    // First update changing transaction parameters
    params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    params.fee = 1000;
    params.flatFee = true;

    let sender = recoveredAccount3.addr;
    let recipient = sender;
    // We set revocationTarget to undefined as 
    // This is not a clawback operation
    let revocationTarget = undefined;
    // CloseReaminerTo is set to undefined as
    // we are not closing out an asset
    let closeRemainderTo = undefined;
    // We are sending 0 assets
    amount = 0;

    // signing and sending txn allows sender to begin accepting asset specified by creator and index
    let opttxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
         amount, note, assetID, params);
    // Must be signed by the account wishing to opt in to the asset    
    rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
    let opttx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + opttx.txId);

    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, opttx.txId);
    // New asset listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    // Transfer New Asset 
    // Tranfer tokens in from the creator to other accounts
    // First update changing transaction parameters
    params = await algodclient.getTransactionParams().do();
    // fee
    params.fee = 1000;
    // flat fee param
    params.flatFee = true;
    sender = recoveredAccount1.addr;
    recipient = recoveredAccount3.addr;
    revocationTarget = undefined;
    closeRemainderTo = undefined;
    //Amount of the asset to transfer
    amount = 10;

    // signing and sending
    let xtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
         amount,  note, assetID, params);
    // Must be signed by the account sending the asset  
    rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
    let xtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + xtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, xtx.txId);
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    // Freeze asset
    // First update changing transaction parameters
    params = await algodclient.getTransactionParams().do();
    //fee
    params.fee = 1000;
    //flat
    params.flatFee = true;
    from = recoveredAccount2.addr;
    freezeTarget = recoveredAccount3.addr;
    freezeState = true;

    // The freeze transaction needs to be signed by the freeze account
    let ftxn = algosdk.makeAssetFreezeTxnWithSuggestedParams(from, note,
        assetID, freezeTarget, freezeState, params)
    rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
    let ftx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + ftx.txId);
    await waitForConfirmation(algodclient, ftx.txId);
    // You should now see the asset is frozen listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    // Revoke an Asset
    // The asset was also created with the ability for revocation 
    // First update changing transaction parameters
    params = await algodclient.getTransactionParams().do();
    // fee
    params.fee = 1000;
    // flat
    params.flatFee = true;   
    sender = recoveredAccount2.addr;
    recipient = recoveredAccount1.addr;
    revocationTarget = recoveredAccount3.addr;
    closeRemainderTo = undefined;
    amount = 10;

    // signing and sending 
    let rtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
       amount, note, assetID, params);
    // must be signed by the account that is the clawback address    
    rawSignedTxn = rtxn.signTxn(recoveredAccount2.sk)
    let rtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + rtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, rtx.txId);
    // You should now see 0 assets listed in the account
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    // Destroy and Asset
    // First update changing transaction parameters
    params = await algodclient.getTransactionParams().do();
    // fee
    params.fee = 1000;
    // flat
    params.flatFee = true;
    addr = recoveredAccount1.addr;
    note = undefined;
    let dtxn = algosdk.makeAssetDestroyTxnWithSuggestedParams(addr, note, assetID, params);
    // The transaction must be signed
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + dtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, dtx.txId);
    console.log("Asset ID: " + assetID);
    console.log("Account 1 = " + recoveredAccount1.addr);
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);  
    // metadata retained

})().catch(e => {
    console.log(e);
    console.trace();
});