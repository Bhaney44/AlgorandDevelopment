# Imports
import hashlib
import base64
import algosdk
from algosdk import account, mnemonic

#from algosdk import account, mnemonic

#Change
## import algosdk

#Generate a new Algorand account and print the public address and private key mnemonic.
def generate_new_account():
	private_key, public_address = account.generate_account()
	passphrase = mnemonic.from_private_key(private_key)
	print("Address: {}\nPassphrase: \"{}\"".format(public_address, passphrase))

# Utility to function to wait until the transaction is confirmed before proceeding.
# @txid: C{string} - The transaction id of the asset
def wait_for_confirmation(client, txid):
	last_round = client.status().get('lastRound')
	while True:
		txinfo = client.pending_transaction_info(txid)
		if txinfo.get('round') and txinfo.get('round') > 0:
			print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
			return txinfo
		else:
			print("Waiting for confirmation...")
			last_round += 1
			client.status_after_block(last_round)
			
#Takes any byte data and returns the SHA512/256 hash in base64.	
def hash_file_data(filename, return_type="bytes"):
	filebytes = open(filename, 'rb').read()
	h = hashlib.sha256()
	h.update(filebytes)
	if return_type == "bytes":
		return h.digest()
	elif return_type == "base64":
		return base64.b64encode(h.digest())

# Adds network-related parameters to transaction data.
def add_network_params(tx_data, client):
	params = client.suggested_params()
	tx_data["fee"] = params.get("fee")
	tx_data["first"] = params.get("lastRound")
	tx_data["last"] = params.get("lastRound") + 1000
	tx_data["gh"] = params.get("genesishashb64")
	tx_data["gen"] = params.get("genesisID")
	return tx_data

# Signs and sends the transaction to the network.
# Returns transaction info.
def sign_and_send(txn, passphrase, client):
	private_key = mnemonic.to_private_key(passphrase)
	stxn = txn.sign(private_key)
	txid = stxn.transaction.get_txid()
	client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
	txinfo = wait_for_confirmation(client, txid)
	return txinfo

# Returns the formatted units for a given asset and amount. 
def balance_formatter(amount, asset_id, client):
	asset_info = client.asset_info(asset_id)
	decimals = asset_info.get("decimals")
	unit = asset_info.get("unitname")
	formatted_amount = amount/10**decimals
	return "{} {}".format(formatted_amount, unit)
