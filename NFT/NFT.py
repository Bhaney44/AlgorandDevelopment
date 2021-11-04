from algosdk import account, encoding, mnemonic,algod
from algosdk.future.transaction import AssetTransferTxn, PaymentTxn, AssetConfigTxn
from algosdk.future.transaction import AssetFreezeTxn
from algosdk.v2client import algod
import json
import urllib3

addresses = []


algod_address = "https://mainnet-algorand.api.purestake.io/ps2"
algod_token = ""
#Initializes Client for node
headers = {"X-API-Key": algod_token }
algod_client = algod.AlgodClient(algod_token,algod_address,headers)
creator_address = " #Put in main creator address here
creator_mnemonic = "" #Put in main creator mnemonic here
creator_key = mnemonic.to_private_key(creator_mnemonic)
#Details of the asset creation transaction.
#Initializes ChoiceCoin, defines manager addresses, and initiaties a main picture/hash
asset_details = {
	"asset_name": "LION,
	"unit_name": "LION",
	"total": 1,
	"decimals": 0,
	"default_frozen": False,
	"manager": creator_address,
	"reserve": creator_address,
	"freeze": '',
	"clawback": '',
	"url": "",
}

def create_buzz():
	params = algod_client.suggested_params()
	transaction = AssetConfigTxn(creator_address, params, **asset_details, strict_empty_address_check = False)
	signature = transaction.sign(creator_key)
    #Signs the transaction with the sender's private key
	algod_client.send_transaction(signature)
	transaction_id = transaction.get_txid()
	transaction_id = str(transaction_id)
	print("Your transaction information is at https://mainnet.algoexplorer.io/tx/" + transaction_id)
