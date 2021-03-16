function hittheroad
	osascript -e 'tell application "Finder" to eject (every disk whose ejectable is true)'
	sudo pmset -a hibernatemode 25
	sudo pmset sleepnow
end
