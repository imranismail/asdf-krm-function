<div align="center">

# asdf-krm-function [![Test](https://github.com/imranismail/asdf-krm-function/actions/workflows/test.yml/badge.svg)](https://github.com/imranismail/asdf-krm-function/actions/workflows/test.yml) [![Lint](https://github.com/imranismail/asdf-krm-function/actions/workflows/lint.yml/badge.svg)](https://github.com/imranismail/asdf-krm-function/actions/workflows/lint.yml)

[krm-function-registry](https://github.com/imranismail/krm-function-registry) plugin for the [asdf version manager](https://asdf-vm.com).

</div>

# Contents

- [Dependencies](#dependencies)
- [Install](#install)
- [Contributing](#contributing)
- [License](#license)

# Dependencies

- `bash`, `curl`, `tar`: generic POSIX utilities.

# Install

Plugin:

```shell
asdf plugin add krm-function
# or
asdf plugin add krm-function https://github.com/imranismail/asdf-krm-function.git
```

krm-function:

```shell
# Show all installable versions
asdf list-all krm-function

# Install specific version
asdf install krm-function latest

# Set a version globally (on your ~/.tool-versions file)
asdf global krm-function latest

# Now krm-function commands are available
krm-function --help
```

Check [asdf](https://github.com/asdf-vm/asdf) readme for more instructions on how to
install & manage versions.

# Contributing

Contributions of any kind welcome! See the [contributing guide](contributing.md).

[Thanks goes to these contributors](https://github.com/imranismail/asdf-krm-function/graphs/contributors)!

# License

See [LICENSE](LICENSE) © [Imran Ismail](https://github.com/imranismail/)
