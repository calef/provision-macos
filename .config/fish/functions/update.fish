# Get OS X Software Updates, and update installed Ruby gems, Homebrew, npm, and their installed packages
function update
	sudo softwareupdate -i -a
	brew update
	brew upgrade
	brew cleanup
end
