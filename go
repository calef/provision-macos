#!/usr/bin/env bash

# Ask for the administrator password upfront.
sudo -v

# Keep-alive: update existing `sudo` time stamp until the script has finished.
while true; do sudo -n true; sleep 60; kill -0 "$$" || exit; done 2>/dev/null &

#install homebrew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew bundle

#install dotfiles
cd "$(dirname "${BASH_SOURCE}")";
rsync --exclude ".git/" --exclude ".DS_Store" --exclude "Brewfile" \
	--exclude "README.md" --exclude "LICENSE" --exclude "go" \
	--exclude "osx.bash" --exclude "init/" -avh --no-perms . ~;

#install GPG keys used to verify RVM installation package
gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB

#install RVM
curl -sSL https://get.rvm.io | bash -s stable
