<div align="center">

# asdf-krm-functions [![Build](https://github.com/imranismail/asdf-krm-functions/actions/workflows/build.yml/badge.svg)](https://github.com/imranismail/asdf-krm-functions/actions/workflows/build.yml) [![Lint](https://github.com/imranismail/asdf-krm-functions/actions/workflows/lint.yml/badge.svg)](https://github.com/imranismail/asdf-krm-functions/actions/workflows/lint.yml)

[krm-functions](https://github.com/imranismail/asdf-krm-functions) plugin for the [asdf version manager](https://asdf-vm.com).

</div>

# Contents

- [Dependencies](#dependencies)
- [Install](#install)
- [Contributing](#contributing)
- [License](#license)

# Dependencies

**TODO: adapt this section**

- `bash`, `curl`, `tar`: generic POSIX utilities.
- `SOME_ENV_VAR`: set this environment variable in your shell config to load the correct version of tool x.

# Install

Plugin:

```shell
asdf plugin add krm-functions
# or
asdf plugin add krm-functions https://github.com/imranismail/asdf-krm-functions.git
```

krm-functions:

```shell
# Show all installable versions
asdf list-all krm-functions

# Install specific version
asdf install krm-functions latest

# Set a version globally (on your ~/.tool-versions file)
asdf global krm-functions latest

# Now krm-functions commands are available
krm-functions --help
```

Check [asdf](https://github.com/asdf-vm/asdf) readme for more instructions on how to
install & manage versions.

# Contributing

Contributions of any kind welcome! See the [contributing guide](contributing.md).

[Thanks goes to these contributors](https://github.com/imranismail/asdf-krm-functions/graphs/contributors)!

# License

See [LICENSE](LICENSE) © [Imran Ismail](https://github.com/imranismail/)
