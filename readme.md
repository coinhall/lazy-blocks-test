# Lazy Block Production

## Manual Testing

> Assumes [Ignite CLI](https://docs.ignite.com/welcome/install) is already installed (though it's only used in the first step to help bootstrap all the necessary genesis files and accounts).

1. Initialise chain using Ignite CLI: `git checkout 33b3528 && ignite chain init && git checkout main`
   - For unknown reasons, the `ignite chain init` command does not work while in `main`
   - Note down the address of `alice` and `bob` so that we can use them later
2. In the `~/.test/config.toml` file, set and ensure `create_empty_blocks = false`
3. Build the binary: `go build -o testd cmd/testd/main.go`
4. Start the chain in another terminal: `./testd start`
   - Notice that the block height halts at `2` (use `./testd status | jq` to verify)
5. Send some funds from `alice` to `bob`: `./testd tx bank send ALICE_ADDR BOB_ADDR 1token`
   - Notice that the block height now halts at `4` which means an increment of 2 blocks (this seems like the "correct behaviour" according to this [forum post](https://forum.cosmos.network/t/turning-create-empty-blocks-to-false-has-no-effect/737/7))
   - Use `./testd query bank balances ALICE_ADDR` to confirm that the balance of `alice` is `1token` less

Once done, you may reset all chain state and repeat from step 4 for further tests: `./testd tendermint unsafe-reset-all`

## Code Changes

In chronological order:

1. With reference to commit `33b3528`: this repository was scaffolded using Ignite CLI's `ignite scaffold chain` command (without the `pkg` directory)
2. With reference to commit `8532399`: the Tendermint (`v0.34.24`) and Cosmos SDK (`v0.46.7`) codebases were cloned as is into the`pkg` directory
3. With reference to commit `8ecfc92`: the `replace` directive was added to the `go.mod` and `pkg/cosmos-sdk/go.mod` files to point to the local Tendermint package located at `pkg/tendermint`
   - At this point, there are zero changes to the underlying source code, and the blockchain should run as per normal (ie. producing blocks every second)
4. With reference to commit `1562834`: the local Tendermint and Cosmos SDk pkgs were updated to reflect the following PRs:
   1. Mark "proof blocks" on Tendermint: <https://github.com/tendermint/tendermint/pull/10004>
   2. Ignore "proof blocks" on Cosmos SDK: <https://github.com/cosmos/cosmos-sdk/pull/15148>
