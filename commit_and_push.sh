#!/bin/bash

# Define the list of subdirectories
# IMPORTANT: Replace these with the actual paths to your subdirectories
subdirectories=(
    "c3dff"
    "cactor"
    "callocator"
    "catomic"
    "cbase"
    "cbasepkgs"
    "cbenchmark"
    "ccmdline"
    "ccode"
    "ccompress"
    "ccore"
    "ccrypto"
    "cd3d12"
    "cdag"
    "cdeptrackr"
    "cdocs"
    "cds"
    "cecs"
    "centry"
    "cfibers"
    "cfile"
    "cfilesystem"
    "cfort"
    "cframegraph"
    "cgamelogic"
    "cgenerics"
    "cgfx"
    "cglfw"
    "charon"
    "chash"
    "chistogram"
    "chlslpp"
    "chshg"
    "cimgui"
    "cjobs"
    "cjson"
    "clang"
    "cmacos"
    "cmath"
    "cmimalloc"
    "cmsg"
    "cp2p"
    "cpair"
    "cpath"
    "crandom"
    "cred"
    "crendergraph"
    "cscode"
    "csocket"
    "csort"
    "cstring"
    "csuperalloc"
    "csuperalloc_analysis"
    "csystem"
    "ctext"
    "ctff"
    "cthread"
    "ctime"
    "cunittest"
    "cuuid"
    "cvkmem"
    "cvmem"
    "cvolk"
    "cvulkan"
    "cwindow"
    "rdno_core"
    "rdno_wifi"
    "rdno_network"
    "rdno_blinky"
    "rdno_sensors"
    "rdno_bedpresence"
)

# Commit message
commit_message="* adjust to changes in ccode"

# Store the current directory to return to it later
original_dir=$(pwd)

# Iterate over each subdirectory
for dir in "${subdirectories[@]}"; do
    echo "-----------------------------------------------------"
    echo "Processing directory: $dir"
    echo "-----------------------------------------------------"

    # Check if the directory exists
    if [ ! -d "$dir" ]; then
        echo "Error: Directory $dir does not exist. Skipping."
        continue
    fi

    # Change to the subdirectory
    cd "$dir" || { echo "Error: Could not change to directory $dir. Skipping."; continue; }

    pwd # Print the current working directory

    # Check if there are any changes to commit
    if [ -z "$(git status --porcelain)" ]; then
        echo "No changes to commit in $dir."
    else

        # Remove origin and add the SSH remote URL
        git remote remove origin
        git remote add origin git@github.com:jurgen-kluft/$dir.git

        # Figure out if we are on main or master branch and set upstream accordingly
        current_branch=$(git rev-parse --abbrev-ref HEAD)
        if [ "$current_branch" != "main" ] && [ "$current_branch" != "master" ]; then
            echo "Warning: You are not on the main or master branch in $dir. Skipping push."
            cd "$original_dir"
            continue
        fi

        git branch --set-upstream-to=origin/$current_branch $current_branch

        # Add all changes to the staging area
        echo "Adding changes to staging area in $dir..."
        git add .
        if [ $? -ne 0 ]; then
            echo "Error: Failed to add changes in $dir. Skipping."
            cd "$original_dir"
            continue
        fi

        echo "Attempting to commit changes..."
        # Perform git commit
        git commit -m "$commit_message"
        if [ $? -ne 0 ]; then
            echo "Git commit failed or nothing to commit in $dir."
            # We can choose to continue to push or skip. For now, let's try to push.
        else
            echo "Git commit successful in $dir."
        fi
    fi

    # Check if there are any changes to push
    # We can be either on main or master branch
    current_branch=$(git rev-parse --abbrev-ref HEAD)
    if [ "$current_branch" != "main" ] && [ "$current_branch" != "master" ]; then
        echo "Warning: You are not on the main or master branch in $dir. Skipping push."
        cd "$original_dir"
        continue
    fi
    # Check if there are any changes to push
    if [ -z "$(git log origin/$current_branch..$current_branch)" ]; then
        echo "No changes to push in $dir."
        cd "$original_dir"
        continue
    fi

    while true; do
        # Attempt to push changes
        echo "Attempting to push changes to remote repository..."
        git push -v
        if [ $? -eq 0 ]; then
            echo "Git push successful in $dir."
            break
        else
            echo "Git push failed. Retrying..."
            sleep 5
        fi
    done

    # Change back to the original directory
    cd "$original_dir" || { echo "Error: Could not return to original directory $original_dir. Exiting."; exit 1; }
    echo ""
done

echo "-----------------------------------------------------"
echo "All specified repositories processed."
echo "-----------------------------------------------------"
