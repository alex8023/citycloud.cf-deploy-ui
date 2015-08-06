#/bin/bash
set -x
set -e
git add .
git commit -m "$1"
git push origin dev
