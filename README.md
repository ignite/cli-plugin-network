# cli-plugin-network

`cli-plugin-network` is an Ignite CLI plugin that allows to launch new Cosmos blockchains by interacting with the Ignite chain to coordinate with validators.

The plugin is integrated into Ignite CLI by default.

[**Check out our documentation for launching chains with the commands**](https://docs.ignite.com/nightly/network/introduction)

## Developer instruction

- clone this repo locally
- Run `ignite plugin add -g /absolute/path/to/cli-plugin-network` to add the plugin to global config
- `ignite network` command is now available with the local version of the plugin.

Then repeat that loop :

- Hack plugin code
- Rerun `ignite network` to recompile the plugin and test
