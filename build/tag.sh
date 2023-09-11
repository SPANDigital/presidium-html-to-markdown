#!/usr/bin/env bash

set -e

DIR="$(dirname "$0")"
source "${DIR}"/include.sh

# Determine the current branch in GitHub Actions
current_branch="${GITHUB_REF#refs/heads/}"

f_info_log "Calculating tag ${current_branch} branch..."
# Calculate the tag based on the branch
if [ "$current_branch" = "main" ]; then
  tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.6.4-debian.9-x64-5.0 /repo -output json -showvariable MajorMinorPatch)
elif [ "$current_branch" = "develop" ]; then
  tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.6.4-debian.9-x64-5.0 /repo -output json -showvariable SemVer)
else
  echo "Branch '$current_branch' is not configured for tagging."
  exit 1
fi

echo "The tag for the current source code is: $tag"


# Create tag in github
f_info_log "Pushing tag to github..."
curl -s -X POST https://api.github.com/repos/presidium-html-to-markdown/git/refs -H "Authorization: token $GITHUB_TOKEN" \
  -d @- <<EOF
{
  "ref": "refs/tags/$tag",
  "sha": "$TRAVIS_COMMIT"
}
EOF
f_info_log "Source code tagged successfully as ${tag}"