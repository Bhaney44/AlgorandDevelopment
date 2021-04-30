import util
import main
import config

#util.generate_new_account()

#from util import hash_file_data
#hash_file_data('LaylaGyoza.jpg')
#hash_file_data('LaylaGyoza.jpg', 'base64')

from main import create
from config import creator_passphrase

create(creator_passphrase)
