function ejectall
	osascript -e 'tell application "Finder" to eject (every disk whose ejectable is true)'
end
