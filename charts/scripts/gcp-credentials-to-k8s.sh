#!/bin/sh
DIR="$HOME/.gcp/credentials/*"

echo "Directory: $DIR"
echo "Using files:"
for d in $DIR ; do
  echo "$d"
  FROM_FILES="--from-file=$(basename "$d")=${d} $FROM_FILES"
done
echo "Executing: kubectl create secret generic gcp-credentials $FROM_FILES"
kubectl create secret generic gcp-credentials $FROM_FILES
