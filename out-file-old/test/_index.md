---
title: "Blah Tests"
weight: 1
---

### external links test case

[adir://groups/1227975](adir://groups/1227975)
 `adir://groups/1227975`
[Send email]({{%baseurl%}}/files/mailto:someone@example.com)

### 1\. Extra space between open paren

A period ( `.`) means one of any character.


### 2\. Broken table - it's just missing the second row of a markdown table definition

| Pattern | Meaning |
| `libgit2/pom\.xml$` | A file called `pom.xml` inside a `libgit2` directory |
| `.+\.md$` | Any file with an `md` file extension |
| `dir/foo/.*` | Any file under a `dir/foo` directory structure |

### 3\. Malformed table - an empty header row

|     |     |     |     |
| --- | --- | --- | --- |
| PR | File Path | Owners | Watchers |
| 1 | `root/file.txt<br>root/app1/install.sh<br>root/recipes.txt<br>` |  | `amp-delivery-engineering` |
| 2 | `root/app1/pom.xml` | `ssmall` | `amp-delivery-engineering` |
| 3 | `root/app2/README.md<br>                  <br>` | `dre` | `amp-delivery-engineering, Yuanling` |
| 4 | `root/app2/recipes.txt<br>` | `ssmall, dre` | `amp-delivery-engineering` |
| 5 | `root/app2/docs/history.md<br>                  ` | `dre` | `amp-delivery-engineering, Yuanling` |

### 4\. Some images didn't make it

#### An image that was correctly converted:

![image](https://media.github.pie.apple.com/user/3143/files/d33e516f-f7ba-483f-b626-d74dba3a317f)

#### An image that was not converted:

![]({{%baseurl%}}/download/attachments/4439017778/rerun.png?version=3&modificationDate=1672780660942&api=v2)
5\. Inconsistent spacing of `<li>` items within an `<p>`
1. A regex pattern that tries to match the changed file paths in the PR (string)
2. the policy being configured for the above pattern. either owners or watchers (map) target of the policy in the form name: user | team (map)
    1. target of the policy in the form name: user | team (map)

6\. Example: 2
1. Mix flour, baking powder, sugar, and salt.
2. In another bowl, mix eggs, milk, and oil.
3. Stir both mixtures together.
4. Fill muffin tray 3/4 full.
5. Bake for 20 minutes.

6\. Example: 3
1. first item
2. second item second item first subitem second item second subitem second item third subitem
    1. second item first subitem
    2. second item second subitem
    3. second item third subitem
3. third item

6\. Example: 4
- first item
- second item second item first subitem second item second subitem second item second subitem first sub-subitem second item second subitem second sub-subitem second item second subitem third sub-subitem second item third subitem
    1. second item first subitem
    2. second item second subitem second item second subitem first sub-subitem second item second subitem second sub-subitem second item second subitem third sub-subitem
        - second item second subitem first sub-subitem
        - second item second subitem second sub-subitem
        - second item second subitem third sub-subitem
    3. second item third subitem
    - second item second subitem first sub-subitem
    - second item second subitem second sub-subitem
    - second item second subitem third sub-subitem
- third item

7\. Don't format links in code blocks

`curl -X POST "` [https://ase-stewards-api.itunes.apple.com/rest/v3/owners/search](https://ase-stewards-api.itunes.apple.com/rest/v3/owners/search) `" -H "accept: application/json" -H "Content-Type: application/json" -d '{"host":"github.pie.apple.com","repo_owner":"its","repo_name":"Jingle","paths":["ragglefraggle","owners.yaml"]}'`
