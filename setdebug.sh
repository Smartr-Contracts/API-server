##!/usr/bin/env bash

if command -v python &>/dev/null; then
    echo "Python is installed."
else
   sudo apt-get install python3.6 
fi

git clone https://github.com/smartbugs/smartbugs
cd smartbugs
install/setup-venv.sh
ln -s "`pwd`/smartbugs" "$HOME/bin/smartbugs"
ln -s "`pwd`/reparse" "$HOME/bin/reparse"
ln -s "`pwd`/results2csv" "$HOME/bin/results2csv"
cd ..
