#!/bin/bash
if [[ "$1" == "install" ]]; then
  sudo apt-get install -y python-pip python-dev build-essential git
  git clone https://github.com/eriknelson/govr.git $HOME/govr
  sudo pip install requests
elif [[ "$1" == "test" ]]; then
  export GOVR_PROJECT_ROOT=$GOPATH/src/github.com/eriknelson/govrhub-ex
  python $HOME/govr/main.py
fi
