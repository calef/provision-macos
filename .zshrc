export CLICOLOR=1
export EDITOR="/usr/local/bin/mate -w"
export HISTCONTROL=ignoreboth
export HISTFILESIZE=32768 
export HISTSIZE=$HISTFILESIZE
export LANG=en_US.UTF-8
export LC_ALL=$LANG
export LS_COLORS='no=00:fi=00:di=01;34:ln=01;36:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.gz=01;31:*.bz2=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.avi=01;35:*.fli=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.ogg=01;35:*.mp3=01;35:*.wav=01;35:'
export MANPAGER="less -X"
export NODE_REPL_HISTORY=~/.node_history
export NODE_REPL_HISTORY_SIZE=$HISTFILESIZE
export NODE_REPL_MODE=sloppy
export NVM_DIR="$HOME/.nvm"
export PATH=./node_modules/.bin/:~/perl5/perlbrew/bin/perlbrew:~/.rvm/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin

alias afk='/System/Library/CoreServices/Menu\ Extras/User.menu/Contents/Resources/CGSession -suspend'
alias p='cd ~/projects'
alias update='sudo softwareupdate -i -a && brew update && brew upgrade && brew cleanup'
alias xyzzy='echo "Nothing happens"'

[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
[ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && . "/usr/local/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion

source ~/perl5/perlbrew/etc/bashrc
