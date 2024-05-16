#!/bin/bash

cd ..
pwd
echo

# Loop through each project folder starting with 'c...' within the root folder
for dir in "c"*/
do
  pwd
  # go into the project folder
  cd "$dir"

  # Check if the current directory contains a git repository
  if [[ -d ".git" ]]; then
    # Push changes to the remote repository (try 10 times)
    n=0
    until [ $n -ge 10 ]
    do
      git push -v && break
      n=$[$n+1]
      sleep 1
    done
    sleep 1
  fi
  
  echo

  # go back into the workspace folder
  cd ..
done

# go back into cbasepkgs
cd cbasepkgs
