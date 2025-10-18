#!/usr/bin/env bash

# Loop through each project folder starting with 'c...' within the root folder
for dir in */
do
  if [[ -d "$dir" && ! -L "$dir" ]]; then

    # if the folder name doesn't start with 'c', skip it
    if [[ ${dir:0:1} != "c" ]]; then
     continue
    fi

    echo "-------------------"
    dir=${dir%*/}
    echo $dir

    # go into the project folder
    cd "$dir"

    # show the project folder
    pwd

    # remove origin
    git remote remove origin
    # add SSH origin 
    git remote add origin git@github.com:jurgen-kluft/$dir.git
    # fetch origin
    git fetch origin
    # get the name of the branch
    currentbranch=$(git rev-parse --abbrev-ref HEAD)
    echo "On branch: $currentbranch"
    # set upstream to origin/$branch
    echo git branch --set-upstream-to=origin/$currentbranch $currentbranch
    git branch --set-upstream-to=origin/$currentbranch $currentbranch

    # Check if the current directory contains a git repository
    if [[ -d ".git" ]]; then
      # Push changes to the remote repository (try 10 times)
      n=0
      until [ $n -ge 10 ]
      do
        # pull changes
        echo "Attempting to pull... (try #$n)"
        git pull -v && break
        n=$[$n+1]
        sleep 1
      done
      sleep 1
    fi

    echo "Done with $dir"
    echo "-------------------"
    echo ""

    # go back into the workspace folder
    cd ..
  fi
done

