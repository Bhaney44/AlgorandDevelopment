# Modify this file as you run through the tutorial.
# Initial variables, {address:passphrase}, {creator:reciever}.
# The address is the place to and from which the funds are transferred.
# The passphrase is the secret key.
creator_address = "TZVTFROZLE6IFGRETS2ULK5JOWOEUMA2LFPYOW2BAFNF27WNMDGQIOJKZM"
creator_passphrase = "ask lucky match spoon drastic long future muscle farm female soon hill balcony bunker man prefer novel unaware that tomato brief window next absorb license"
receiver_address = "66PWZ4YF7MT6SRAFESWZ44KA262BLNSH6PTE7CUULPIHMNWYX4ZGPMZJMQ"
receiver_passphrase = "excite profit arch either banner movie open tool model must step zone wisdom deal innocent truck test rocket angle glare between dismiss spell able local"

# Credentials to connect through an algod client
# To find address and token run the following:
# 	cat $ALGORAND_DATA/algod.net
# 	cat $ALGORAND_DATA/algod.token
# algod_address = "http://localhost:4001"

algod_address = "https://testnet.algoexplorerapi.io"
algod_token = "a115bfe5a9c583b04d63f675ca8e397c0777f334b022d66c86c8702cd7c34534"

# Details of the asset creation transaction.
asset_details = {
	"asset_name": "DogCoin",
	"unit_name": "Woof",
	"total": 888888888,
	"decimals": 2,
	"default_frozen": False,
	"manager": creator_address,
	"reserve": creator_address,
	"freeze": creator_address,
	"clawback": creator_address,
	"url": "LaylaGyoza.jpg",
	"metadata_hash": b'O\x88\xfd\xf2\xd1\xfe\xee\x96+\xf9\xf0\xb6\xb2\x8d\r\xb5\xced)#\x9bV\xce\xa4\x81\xa6\xb9\xbd\x0e\xf7al'
}

metadata_file = "LaylaGyoza.jpg"
metadatahash_b64 = "T4j98tH+7pYr+fC2so0Ntc5kKSObVs6kgaa5vQ73YWw="

# The asset ID is available after the asset is created.
asset_id = 0 # change this

