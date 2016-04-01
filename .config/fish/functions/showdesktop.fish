# Show all desktop icons (useful when presenting)
function showdesktop
	defaults write com.apple.finder CreateDesktop -bool true
	killall Finder
end
