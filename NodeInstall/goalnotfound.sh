# Error
# Command not found: goal
# I am installing a node on a new machine and following the documentation for Install a Node. 
# https://developer.algorand.org/docs/run-a-node/setup/install/
# The installation sequence I followed was as follows.

export ALGORAND_DATA="$HOME/node/data"
export PATH="$HOME/node:$PATH"
mkdir ~/node
cd ~/node
curl https://raw.githubusercontent.com/algorand/go-algorand-doc/master/downloads/installers/update.sh -O
chmod 544 update.sh
./update.sh -i -c stable -p ~/node -d ~/node/data -n
#RETURN: Current Version = 0
#RETURN: This platform arm64 is not supported by updater.
goal node start
#RETURN: zsh: command not found: goal
./goal node start
#RETURN: zsh: no such file or directory: ./goal

#I am not sure where the error is coming from. 
#In my node directory, there are are only two files data and update.sh. 
#I think the problem is the export ALGORAND_DATA function is not exporting the necessary data, but I am not sure why. Any suggestions or advice would be greatly appreciated, thanks!