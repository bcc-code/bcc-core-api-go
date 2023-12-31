#!/bin/bash

# set -x

# Usage:
#   
#   release X.X.X
# 
# Creates a new version of the package. This script creates a new release branch
# where it updates the version variable. It then creates a matching tag and 
# pushes both to GitHub.

# The file holding a version variable.
VERSION_FILE=bcccoreapi/version.go

function create_release_branch {
  local version="$1"
  git checkout -b "releases/${version}"
}

function update_version {
  local version="$1"
  sed -i.backup "s/var Version =.*/var Version = \"${version}\"/" "${VERSION_FILE}"
  rm -rf "${VERSION_FILE}.backup"
}

function commit_changes {
  local version="$1"
  git add "${VERSION_FILE}"
  git commit -m "v${version}"
  git tag -a "v${version}" -m "v${version}"
}

function push_changes {
  local version="$1"
  git push origin "releases/${version}"
  git push origin "v${version}"
}

function main {
  local version="$1"
  if [[ $version = v* ]]; then
    echo "version should be in the format X.Y.Z"
    exit 1
  fi
  if git rev-parse "v$version" > /dev/null 2>&1; then
    echo "version $version already exists"
    exit 1
  fi
  
  local current_branch=$(git rev-parse --symbolic-full-name --abbrev-ref HEAD)
  
  create_release_branch $version
  update_version $version
  commit_changes $version
  push_changes $version
  
  git checkout $current_branch
}

main "$@"