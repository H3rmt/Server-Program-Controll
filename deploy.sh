#!/bin/bash

# Update repo
printf "fetching from git\n"
git fetch && git pull
git log -1

# Update repo
printf "starting deploy\n"
./.deploy.sh