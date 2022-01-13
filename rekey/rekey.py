#Imports
import base64
import algosdk
from algosdk.v2client import algod
from algosdk import account, mnemonic
from algosdk.future.transaction import write_to_file
from algosdk.future.transaction import AssetConfigTxn, AssetTransferTxn
from algosdk.future.transaction import PaymentTxn

# Connection
algod_address = "https://mainnet-algorand.api.purestake.io/ps2"
algod_token = ""
headers = {"X-API-Key": algod_token }
client = algod.AlgodClient(algod_token,algod_address,headers)

ADDA = ""
passphrase = ""
ADDB = ""
params = client.suggested_params()
amount = 0
txn = PaymentTxn(sender=ADDA,sp=params,receiver=ADDA,amt=amount,close_remainder_to=None,note=None,lease=None,rekey_to=ADDB)

def wait_for_confirmation(client, transaction_id, timeout):
    start_round = client.status()["last-round"] + 1;
    current_round = start_round
    while current_round < start_round + timeout:
        try:
            pending_txn = client.pending_transaction_info(transaction_id)
        except Exception:
            return 
        if pending_txn.get("confirmed-round", 0) > 0:
            return pending_txn
        elif pending_txn["pool-error"]:  
            raise Exception(
                'pool error: {}'.format(pending_txn["pool-error"]))
        client.status_after_block(current_round)                   
        current_round += 1
    raise Exception(
        'pending tx not found in timeout rounds, timeout value = : {}'.format(timeout))

def sign_and_send(txn, passphrase, client):
	private_key = mnemonic.to_private_key(passphrase)
	stxn = txn.sign(private_key)
	txid = stxn.transaction.get_txid()
	client.send_transaction(stxn)
	wait_for_confirmation(client, txid, 5)
	print('Confirmed TXID: {}'.format(txid))
	txinfo = client.pending_transaction_info(txid)
	return txinfo

if passphrase:
    txinfo = sign_and_send(txn, passphrase, client)
    print("Transaction ID Confirmation: {}".format(txinfo.get("tx")))
else:
    write_to_file([txns], "transfer.txn")
