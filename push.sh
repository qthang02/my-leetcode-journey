#!/bin/bash

message=$1

pushMess=problem" "$message

git add .

git commit -m "$pushMess"

git push

echo "pushed successfully"