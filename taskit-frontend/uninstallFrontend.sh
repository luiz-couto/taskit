#!/bin/bash
echo "Uninstalling taskit...."

# remove if it was already installed in .profile
sed -i -e '/taskit-frontend/d' ~/.profile

echo "

Taskit uninstalled.
You must restart your PC to apply the changes.

"