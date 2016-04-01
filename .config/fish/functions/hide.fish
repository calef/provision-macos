# Hide hidden files in Finder
function hide
	defaults write com.apple.finder AppleShowAllFiles -bool false
	killall Finder
end
