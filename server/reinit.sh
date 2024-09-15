#!/bin/zsh

source "$HOME/.zprofile"
cd "$(dirname "$0")" || (echo "Could not cd into the dir"; exit)
rm test.db
sqlite3 test.db < initdb.sql 
go build 
systemctl --user restart split
