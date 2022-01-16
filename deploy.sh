#!/bin/bash

# dirs for the output
html="/var/www/html"
goRemote="/home/enrico/Remote"
goServer="/home/enrico/Server"


# copy/override config files
read -p "initial deploy? [y/N] " r
initial=False
if [ "$r" = "y" ] || [ "$r" = "Y" ]; then
	echo "selected initial deploy"
	initial=True
else
	echo "selected redeploy"
fi
echo

# dirs inside the $projectFolder dir
website="Website"
server="Server"
remote="Remote"

# gitHome
projectFolder=$(pwd)

# Update repo
fetch=$(git fetch)
printf "fetching from git:%s\n" "$fetch"
reset=$(git reset --hard)
printf "reset git:%s\n\n" "$reset"

CopyFolder() {
	if [ ! -d "$1" ]; then
		echo "source $projectFolder/$1 not existing"
		return
	fi

	if [ ! -d "$2" ]; then
		mkdir -p "$2"
		echo "created $2 dir"
	else
		if [ "$(ls -A "$2")" ]; then
			rm -r "$2/"*
			echo "removed files from $2"
		else
			echo "$2 empty"
		fi
	fi

	echo "coping files from $projectFolder/$1 to $2"
	cp -r "$projectFolder/$1/"* "$2"

	echo ""
}

BuildGo() {
	if [ ! -d "$1/" ]; then
		echo "source $projectFolder/$1 not existing"
		return
	fi
	if [ ! -d "$2/" ]; then
		mkdir -p "$2"
		echo "created $2 dir"
	fi

	cd "$projectFolder/$1" || return

	echo "building go $1"
	go build -o "$1"

	cd "$projectFolder" || return

	echo "moving file $1 from $1 to $2"
	mv "$projectFolder/$1/$1" "$2/$1"

	echo ""
}

CopyFiles() {
	arr=("$@")

	# start from index 2 to exclude source and output
	for ((i = 2; i < ${#arr[@]}; i++)); do
		a="${arr[i]}"
		echo "coping file $a from $projectFolder/$1 to $2"
		cp "$projectFolder/$1/$a" "$2/$a"
	done
	echo ""
}

CopyFolder $website $html

BuildGo $server $goServer
serverFiles=("config.json")
if [ $initial = True ]; then
	CopyFiles $server $goServer "${serverFiles[@]}"
fi

BuildGo $remote $goRemote
remoteFiles=("programs.json" "config.json")
if [ $initial = True ]; then
	CopyFiles $remote $goRemote "${remoteFiles[@]}"
fi
