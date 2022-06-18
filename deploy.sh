#!/bin/bash

# dirs for the output
goRemote="/home/remote/Remote"
goServer="/home/remote/Server"


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
server="Server"
remote="Remote"

# gitHome
projectFolder=$(pwd)


# Update repo
printf "fetching from git\n"
git fetch && git pull
git log -1 --pretty=%B

# Update docker-compose
printf "updating docker-compose\n"
docker-compose up

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
