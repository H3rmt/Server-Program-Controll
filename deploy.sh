#!/bin/sh

# githome
#home="/home/enrico/Documents/Server-Program-Controll"
home=$(pwd)

# dirs for the output
html="/home/enrico/Programs/Website"
goserver="/home/enrico/Programs/Server"
goremote="/home/enrico/Programs/Remote"


# Update repo
pwd=$(pwd)
printf "WebsiteHome:%s\n" "$pwd"

pull=$(git pull)
printf "pulling from git:%s\n\n" "$pull"


Copy () {
  cd $home

  if [ ! -d "$1/" ]; then
    echo "source $home/$1 not existing"
    return
  fi

  if [ ! -d "$2" ]; then
    mkdir "$2"
    echo "created $2 dir"
  else
    if [ "$(ls -A $2)" ]; then
      cd $2
      sudo rm -r *
      echo "removed files from $2"
    else
      echo "$2 empty"
    fi
  fi

  cd "$home/$1"

  printf "coping files:  "
  printf "%s; " $(ls)

  pwd=$(pwd)
  printf "\nfrom:%s\n" "$pwd"

  sudo cp -r * "$2"

  echo ""
}

# dirs inside the $home dir
website="Website"
server="Server"
remote="Remote"

Copy $website $html
Copy $server $goserver
Copy $remote $goremote