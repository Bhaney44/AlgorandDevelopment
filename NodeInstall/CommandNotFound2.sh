# I am still getting an error here. For Homebrew, running:
# Link: https://medium.com/swlh/run-x86-terminal-apps-like-homebrew-on-your-new-m1-mac-73bdc9b0f343

Command Line Input: arch -x86_64 /bin/bash -c "$(curl -fsSL [https://raw.githubusercontent.com/Homebrew/install/master/in...](https://raw.githubusercontent.com/Homebrew/install/master/install.sh))"

Returns: curl: (22) The requested URL returned error: 404

# For the Node installation, running:

Command Line Input: arch -x86_64 /bin/bash -c "curl https://raw.githubusercontent.com/algorand/go-algorand-doc/master/downloads/installers/update.sh -O"

Returns: curl: (23) Failed writing body (0 != 853)

#I got an error installing for both Homebrew and the Algorand Node. 
# Running ls in the node directory yields two files: data and update.sh.

# However, the second option, installing the Rosetta-Terminal worked for Homebrew installation. 
# But I am still getting the same error for the Algorand Node. Running:

Command Line Input: curl https://raw.githubusercontent.com/algorand/go-algorand-doc/master/downloads/installers/update.sh -O

Return: curl: (23) Failure writing output to destination

# I have been iterating between installing in a virtual environment using Anaconda and running directly, however neither has returned a successful installation. 
# I also was sure to follow the Install Instructions. I am running on macOS Big Sur 11.3.1.

# Any suggestions or advice would be appreciated. Thank you!
