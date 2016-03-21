#!/usr/bin/env bash

# Ask for the administrator password upfront.
sudo -v

# Keep-alive: update existing `sudo` time stamp until the script has finished.
while true; do sudo -n true; sleep 60; kill -0 "$$" || exit; done 2>/dev/null &

#bootstrap a system
function doIt() {
	#install homebrew
	ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	# Make it so that we can use our Brewfile
	brew tap 'homebrew/brewdler'
	brew brewdle

	#install dotfiles
	cd "$(dirname "${BASH_SOURCE}")";

	git pull origin master;

	rsync --exclude ".git/" --exclude ".DS_Store" --exclude "Brewfile" \
		--exclude "README.md" --exclude "LICENSE" --exclude "go" \
		--exclude "bin/" --exclude "init/" -avh --no-perms . ~;
	source ~/.bash_profile;

	#install perlbrew
	curl -L http://install.perlbrew.pl | bash

	#install perl
	perlbrew install perl-5.22.1
	perlbrew switch perl-5.22.1

	#install projects
	mkdir -p ~/Projects
	cd ~/Projects
	git clone https://github.com/MotoFish/pretty-dam-quick.git

	#install pretty-dam-quick
	cd pretty-dam-quick
	git pull origin master
	./go

	cd "$(dirname "${BASH_SOURCE}")";
}

if [ "$1" == "--force" -o "$1" == "-f" ]; then
	doIt;
else
	read -p "This may overwrite existing files in your home directory. Are you sure? (y/n) " -n 1;
	echo "";
	if [[ $REPLY =~ ^[Yy]$ ]]; then
		doIt;
	fi;
fi;
unset doIt;
