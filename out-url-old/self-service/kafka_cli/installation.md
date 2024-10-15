---
title: "Installation¶"
weight: 2
---

Note

Make sure to update `pie-kafka-cli` whenever it tells you that a new version is available. These updates provide bug fixes and new functionality.

### Using `pip` [¶](#using-pip "Link to this heading")

Please follow the [guidelines on setting up a virtual environment](https://pages.github.pie.apple.com/python-frameworks/python-dot-apple/venv/), to set up one and isolate this CLI from your system packages.

```
$ pip install --upgrade --index-url https://pypi.apple.com/simple pie-kafka-cli

```

### Using `pipx` [¶](#using-pipx "Link to this heading")

Installation with [`pipx`](https://github.com/pypa/pipx/) is possible. It allows you to install the CLI without requiring the manual setup of a virtual environment:

```
$ pipx install --index-url https://pypi.apple.com/simple --preinstall setuptools pie-kafka-cli

```

