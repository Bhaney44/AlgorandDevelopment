# Creator and reciever
creator_address = "6U4X6KSDTVNVLNKTQAQUWSVSI4LQGPOMSSO34XUJCDW5CXCUSYUBJLTVCI"
creator_passphrase = "one online sponsor radio kangaroo horn license else spot pistol isolate kite arch elite author genre seminar scrap sponsor another original stand radio abandon guilt"
receiver_address = "DILZEWFPINKNAIXL4T4Z63SDVJ6DWK3KMK43S2XXJGVVC52M54MP4RMV5A"
receiver_passphrase = "session comfort ice manual lazy thrive crumble oak tissue quiz region broken giant obtain metal expose burden north core fiber giraffe bird multiply absent quarter"

# Credentials to connect through an algod client

#Try one from developer docs
#algod_address = "https://api.host.com"
#Try two from reciever address in article
algod_address = "http://localhost:8080"

algod_token = "a115bfe5a9c583b04d63f675ca8e397c0777f334b022d66c86c8702cd7c34534"

# Details of the asset creation transaction
asset_details = {
	"sender": creator_address,
	"asset_name": "LaylaCoin",
	"unit_name": "Gyoza",
	"total": 1000000,
	"decimals": 6,
	"default_frozen": False,
	"manager": creator_address,
	"reserve": creator_address,
	"freeze": creator_address,
	"clawback": creator_address,
	"url": "LaylaGyoza.jpg",
	"metadata_hash": b'O\x88\xfd\xf2\xd1\xfe\xee\x96+\xf9\xf0\xb6\xb2\x8d\r\xb5\xced)#\x9bV\xce\xa4\x81\xa6\xb9\xbd\x0e\xf7al'
}

#Meta data
metadata_file = "LaylaGyoza.jpg"
metadatahash_b64 = "T4j98tH+7pYr+fC2so0Ntc5kKSObVs6kgaa5vQ73YWw="

# The asset ID is available after the asset is created.
asset_id = 1
