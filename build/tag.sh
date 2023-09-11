#!/usr/bin/env bash

set -e

DIR="$(dirname "$0")"
source "${DIR}"/include.sh

git checkout "${TRAVIS_BRANCH}"

f_info_log "Calculating tag ${TRAVIS_BRANCH} branch..."
if [ "${TRAVIS_BRANCH}" = "main" ]; then
  tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.6.4-debian.9-x64-5.0 /repo -output json -showvariable MajorMinorPatch)
elif [ "${TRAVIS_BRANCH}" = "develop" ]; then
  tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.6.4-debian.9-x64-5.0 /repo -output json -showvariable SemVer)
fi
f_info_log "The tag for current source code is: ${tag}"

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
