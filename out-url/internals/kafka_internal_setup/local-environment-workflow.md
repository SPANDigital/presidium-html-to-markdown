---
title: "Local environment and workflow¶"
weight: 5
---

### Setup[Internal Link: ¶](\#setup)

Before building, you must configure the Apple artifactory server(s). See
the [External Link: rio
docs](https://docs.aci.apple.com/artifactory/gradle.html)
for information on how to set this up.

### Remotes[Internal Link: ¶](\#remotes)

First, create your own fork of our [External Link: ACI Kafka
fork](https://github.pie.apple.com/pie/apache-kafka). Then: :

```
git clone git@github.pie.apple.com:your-github-username/apache-kafka.git
cd apache-kafka
git remote add upstream git@github.pie.apple.com:pie/apache-kafka.git
git remote add apache-kafka https://github.com/apache/kafka

```

### Build[Internal Link: ¶](\#build)

You can build the binaries with the following command: :

```
./gradlew clean package

```

### Official releases[Internal Link: ¶](\#official-releases)

For purposes of illustration, let’s assume we have already created an
internal release of Kafka version “x” with our internal patches
applied, and we now want to create an internal release for Kafka version
“y”.

- Fetch tags from apache :


  ```
  git fetch --tags apache-kafka

  ```

- Determine the commit for the corresponding release tag :


  ```
  git rev-list -n 1 <y>

  ```

- Create a new “apache-” tag for the release and push it upstream :


  ```
  git tag apache-<y> <commit>
  git push upstream apache-<y>

  ```

- Create a new “develop-” branch from the tag and push it upstream :


  ```
  git checkout -b develop-<y> apache-<y>
  git push upstream develop-<y>

  ```

- Create a new “release-” branch from the “develop-” branch and
  push it upstream :


  ```
  git checkout -b release-<y> develop-<y>
  git push upstream release-<y>

  ```

- Create apache\_y\_with\_our\_patches from develop-y :


  ```
  git checkout -b apache_<y>_with_our_patches develop-<y>

  ```

- Check the patch list we’ve maintained for release “x” so that you
  can apply it to release “y” :


  ```
  git log develop-<x> ^apache-<x> --no-merges --oneline | tail -r

  ```

- Rebase apache\_y\_with\_our\_patches and insert the commits from our
  patch list :


  ```
  git rebase -i apache-<y>
  # Insert each commit from the "log" command and prepend "p" for pick
  # If a commit is directly related to the version, prepend "e" for edit rather than "p"

  ```

- There will inevitably be conflicts with the rebase, so those will
  need to be resolved. The initial version should be y.0
- Run tests :


  ```
  ./gradlew cleanTest test

  ```

- Assuming tests pass, push apache\_y\_with\_our\_patches to origin and
  create a PR against develop-y
- Merge this PR to create the release

### Making a change[Internal Link: ¶](\#making-a-change)

Steps to make a change to Kafka version 1.0.1.

- Checkout develop-1.0.1 :


  ```
  git checkout --track origin/develop-1.0.1

  ```

- Checkout a local work branch `myStuff` on top of develop-1.0.1 :


  ```
  git checkout -b myStuff develop-1.0.1

  ```

- Make changes, including incrementing the version from 1.0.1.x. Push
  to origin. Create pull request (PR) against upstream/develop-1.0.1.
- If [External Link: prb
  job](https://rio.apple.com/projects/pie-apache-kafka/pipeline-specs/pie-apache-kafka-develop-1.0.1-prb/pipelines/)
  passes, you can merge your PR.
- Next few steps are automatic if everything goes well.
  - [External Link: merge
    job](https://rio.apple.com/projects/pie-apache-kafka/pipeline-specs/pie-apache-kafka-develop-1.0.1-merge/pipelines/)
    gets triggered.
  - If merge job passes, changes gets merged to _release-1.0.1_
    branch.
  - [External Link: release
    job](https://rio.apple.com/projects/pie-apache-kafka/pipeline-specs/pie-apache-kafka-develop-1.0.1-release/pipelines)
    triggers and push artifacts to artifactory.

### Back ports of upstream fixes[Internal Link: ¶](\#back-ports-of-upstream-fixes)

As we run a version that is a bit behind the latest upstream release it
sometimes makes sense to back port relevant changes to our version. This
is typically done by cherry-picking commits to either the `trunk` branch
or some release branch such as `2.1`. When doing this there are a few
things to think about:

- We would like to keep amount of change that we put on top of an
  upstream release as small as possible. If the changes you want to
  back port are significant, discuss this with the rest of the team to
  make sure that we make the right trade-off.
- To keep track of which upstream commit was cherry picked onto our
  branch, please add a header such as this one to the cherry pick
  commit message:
  `X-Apple-Cherry-Picked: 70828cea49ab8a3ceb54a9017618b84e0d9c1420`.
- Be aware of the fact that new commits that gets merged onto the
  `release-` branch will trigger the publish pipeline, so please be
  sure to bump the version number with a separate commit in the same
  PR as the cherry-picked commit(s).

