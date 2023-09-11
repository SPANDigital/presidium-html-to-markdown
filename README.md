# Presidium HTML to Markdown converter

**html2md** is tool that allow you to convert HTML files into [Presidium](https://presidium.spandigital.net/) markdown articles.

## Getting Started

### Installation

Install homebrew

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Add SPAN's homebrew tap
```bash
brew tap SPANDigital/homebrew-tap https://github.com/SPANDigital/homebrew-tap.git
```

Install converter
```bash
brew install html2md
```

## Usage
```bash
Usage:
  html2md convert [source] [dest] [flags]

Flags:
      --headers strings   the filters for article titles (default [h1,h2])
      --select string     the part of the page to select and convert (default "body")
```

Download and convert the [Presidium](https://presidium.spandigital.net/) website
```bash
html2md convert https://presidium.spandigital.net/ ./presidium --select="#presidium-content"
```

Convert local html files
```bash
html2md convert ./html-files ./presidium --select=.article --headers=h1
```

## Config

Sample config file
```yaml
html:
  remove: ['.article-title .permalink', '#warning'] # CSS selectors for elements that should be removed before conversion.
  replace:
    - match: '.tooltips-term' # CSS selector for the element to be replaced.
      # Below are the arguments to select elements relative to the matched element.
      select: ['href', '.tooltips-text']
      # Replacement pattern with the selected arguments.
      replace: '{{< tooltip "$1" text="$2" >}}'
markdown:
  replace:
    - pattern: '\[([^]]+)\]\(([^\)]+)\)' # Regex pattern used for selecting and capturing specific content.
      # The captured content is then utilized in the replacement pattern below.
      with: "[$1]({{%baseurl%}}/$2)" # This is the replacement pattern for converting Markdown links.

```