#!/bin/bash
echo "Installing taskit...."

firstPart='export PATH="$PATH:'
adressToThisFolder=`pwd`
lastPart='/bin"'

# remove if it was already installed in .profile
sed -i -e '/taskit-frontend/d' ~/.profile

# add PATH for bin in .profile
echo ${firstPart}${adressToThisFolder}${lastPart} >> ~/.profile

echo "Taskit installed!"
echo "

You must restart your PC to apply the changes. But, if you like, 
you can also run the command

'source ~/.profile'

to apply the changes to this terminal. Test it by just running

'taskit'

"