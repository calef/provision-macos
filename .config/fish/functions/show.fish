#Show hidden files in Finder
function show
	defaults write com.apple.finder AppleShowAllFiles -bool true
	killall Finder
end
