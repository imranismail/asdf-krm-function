# Contributing

Testing Locally:

```shell
asdf plugin test <plugin-name> <plugin-url> [--asdf-tool-version <version>] [--asdf-plugin-gitref <git-ref>] [test-command*]

# TODO: adapt this
asdf plugin test krm-functions https://github.com/imranismail/asdf-krm-functions.git "krm-functions --help"
```

Tests are automatically run in GitHub Actions on push and PR.
