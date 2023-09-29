#!/bin/sh

git tag "$1"

git push --tags
git push --all
