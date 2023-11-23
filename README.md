# cli-plugin-network

`cli-plugin-network` is an app developed for [Ignite CLI](https://github.com/ignite/cli).

The Ignite App adds `ignite network` commands that allow launching new Cosmos blockchains by interacting with the Ignite Chain to coordinate with validators.

The app is installed into Ignite CLI by default.

[**Check out our documentation for launching chains with the commands**](https://docs.ignite.com/nightly/network/introduction)

## Developer instruction

- Clone this repo locally
- Run `ignite app install -g /absolute/path/to/cli-plugin-network` to add the app to global config
- The `ignite network` command is now available with the local version of the app

Then repeat the following loop:

- Hack on the plugin code
- Rerun `ignite network` to automatically recompile the app and test
