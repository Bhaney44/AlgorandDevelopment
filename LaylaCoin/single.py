# Imports
from config import *
from util import sign_and_send, balance_formatter

from algosdk.v2client import algod
from algosdk import account, mnemonic
from algosdk.future.transaction import write_to_file
from algosdk.future.transaction import AssetConfigTxn, AssetTransferTxn
from algosdk.future.transaction import AssetTransferTxn
import hashlib
import base64
from algosdk import account, mnemonic

# Variables
creator_address = "..."
creator_passphrase = "..."
receiver_address = "..."
receiver_passphrase = "..."
algod_address = "https://testnet.algoexplorerapi.io"
algod_token = "a115bfe5a9c583b04d63f675ca8e397c0777f334b022d66c86c8702cd7c34534"
asset_id = 17103284

## 1. Generate the LaylaCoin Creator and Receiver Accounts
## Create New Accounts
def generate_new_account():
	private_key, public_address = account.generate_account()
	passphrase = mnemonic.from_private_key(private_key)
	print("Address: {}\nPassphrase: \"{}\"".format(public_address, passphrase))
generate_new_account()

## 2. Define the LaylaCoin Parameters
## Hash image
# Takes any byte data and returns the SHA512/256 hash in base64.
def hash_file_data(filename, return_type="bytes"):
	filebytes = open(filename, 'rb').read()
	h = hashlib.sha256()
	h.update(filebytes)
	if return_type == "bytes":
		return h.digest()
	elif return_type == "base64":
		return base64.b64encode(h.digest())
hash_file_data('LaylaGyoza.jpg')
hash_file_data('LaylaGyoza.jpg', 'base64')

## 3. Create LaylaCoin
## Create Asset
creator_passphrase = "ask lucky match spoon drastic long future muscle farm female soon hill balcony bunker man prefer novel unaware that tomato brief window next absorb license"
def create(passphrase=None):
	params = client.suggested_params()
	txn = AssetConfigTxn(creator_address, params, **asset_details)
	if passphrase:
		txinfo = sign_and_send(txn, passphrase, client)
		asset_id = txinfo.get('asset-index')
		print("Asset ID: {}".format(asset_id))
	else:
		write_to_file([txn], "create_coin.txn")
create(creator_passphrase)

## 4. Opt-In to Receive LaylaCoin
## Optin
# from main import optin
# from config import receiver_passphrase
receiver_passphrase = "..."
def optin(passphrase=None):
	params = client.suggested_params()
	txn = AssetTransferTxn(sender=receiver_address, sp=params, receiver=receiver_address, amt=0, index=asset_id)
	if passphrase:
		txinfo = sign_and_send(txn, passphrase, client)
		print("Opted in to asset ID: {}".format(asset_id))
	else:
		write_to_file([txns], "optin.txn")

optin(receiver_passphrase)

## 5. Read the LaylaCoin Balance
## Balance
receiver_address = "..."
asset_id = 17103284
def check_holdings(asset_id, address):
	account_info = client.account_info(address)
	assets = account_info.get("assets")
	for asset in assets:
		if asset['asset-id'] == asset_id:
			amount = asset.get("amount")
			print("Account {} has {}.".format(address, balance_formatter(amount, asset_id, client)))
			return
	print("Account {} must opt-in to Asset ID {}.".format(address, asset_id))
check_holdings(asset_id, receiver_address)

## 6. Transfer LaylaCoin to the Receiver Account
## Transfer
creator_passphrase = "..."
def transfer(passphrase=None):
    amount = 6000
    params = client.suggested_params()
    txn = AssetTransferTxn(sender=creator_address, sp=params, receiver=receiver_address, amt=amount, index=asset_id)
    if passphrase:
        txinfo = sign_and_send(txn, passphrase, client)
        formatted_amount = balance_formatter(amount, asset_id, client)
        print("Transferred {} from {} to {}".format(formatted_amount, 
            creator_address, receiver_address))
        print("Transaction ID Confirmation: {}".format(txinfo.get("tx")))
    else:
        write_to_file([txns], "transfer.txn")
transfer(creator_passphrase)

## Confirm
from config import creator_address, receiver_address
creator_address = "..."
receiver_address = "..."
# Checks the asset balance for the specific address and asset id.
def check_holdings(asset_id, address):
    account_info = client.account_info(address)
    assets = account_info.get("assets")
    for asset in assets:
        if asset['asset-id'] == asset_id:
            amount = asset.get("amount")
            print("Account {} has {}.".format(address, balance_formatter(amount, asset_id, client)))
            return
    print("Account {} must opt-in to Asset ID {}.".format(address, asset_id))

check_holdings(asset_id, creator_address)
check_holdings(asset_id, receiver_address)
