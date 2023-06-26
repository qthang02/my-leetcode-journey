#!/bin/bash

message = $1

git add .

git commit -m "problem $message"

git push

echo "pushed successfully"