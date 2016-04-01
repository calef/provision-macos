#!/usr/bin/env bash

# Ask for the administrator password upfront.
sudo -v

# Keep-alive: update existing `sudo` time stamp until the script has finished.
while true; do sudo -n true; sleep 60; kill -0 "$$" || exit; done 2>/dev/null &

#install homebrew
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
# Make it so that we can use our Brewfile
brew tap 'homebrew/brewdler'
brew brewdle

#install dotfiles
cd "$(dirname "${BASH_SOURCE}")";
rsync --exclude ".git/" --exclude ".DS_Store" --exclude "Brewfile" \
	--exclude "README.md" --exclude "LICENSE" --exclude "go" \
	--exclude "bin/" --exclude "init/" -avh --no-perms . ~;

#change the shell to fish
chsh -s /usr/local/bin/fish
