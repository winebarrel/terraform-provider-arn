# AGENTS.md

A Terraform provider whose only contents are functions that build AWS ARNs.
There is one function per resource type in the AWS service reference, 2332 of
them, named `<service>_<resource>`. The account, region and partition come
from `.arn.hcl` in the directory Terraform runs in.

`README.md` is the user-facing description. `performance.md` has measurements
of what 2332 functions cost.

## Generated directories

**`docs/` and `examples/` are generated. Do not edit them by hand.**

| Directory | Written by | Source |
|---|---|---|
| `docs/` | `make docs` | `templates/`, plus the schema the built provider reports |
| `examples/functions/` | `make gen` | the AWS service reference |
| `internal/arnspec/spec_*_gen.go` | `make gen` | the AWS service reference |

`docs/index.md` comes from `templates/index.md.tmpl`; edit the template. The
2332 pages under `docs/functions/` come from each function's `Definition` in
`internal/provider/arn_function.go` and the matching example under
`examples/functions/<name>/function.tf`; edit those.

`make gen` refetches the feed, so it can pick up AWS changes unrelated to what
you are working on. Check the diff to `internal/arnspec/` before committing,
and keep a feed update as its own commit.

## Layout

```
main.go                     provider server
cmd/gen/                    reads the AWS service reference, writes the tables and examples
internal/arnconf/           .arn.hcl: parsing, import, and account resolution
internal/arnspec/           ARN templates: parsing and building
internal/arnspec/spec_*_gen.go   generated, one file per AWS service
internal/arnvalue/          shape checks for the partition, region and account fields
internal/provider/          the provider and the one function type every function is
templates/                  source for docs/index.md
tools/                      go:generate entry point for tfplugindocs
```

Every function is the same `ARNFunction` type with a different spec, so there
is one implementation and 2332 instances.

## Where a value comes from

An ARN is `arn:partition:service:region:account:resource`. Fields are
identified by counting separators, not by the placeholder name AWS used: the
account field is `${Account}` in most templates but `${AccountId}` in chime,
datasync and sso, `${ManagementAccountId}` in account and `${VpcOwnerAccount}`
in kafka. The partition, region and account fields come from the
configuration; everything else is a function argument.

`TestGeneratedPlaceholderPositions` pins that correspondence across every
template. It fails when the feed adds one, which is intended: the addition
should be looked at rather than absorbed.

## Commands

```console
$ make test          # unit and acceptance tests
$ make lint          # golangci-lint
$ make tf-console    # terraform console against a locally built provider
$ make gen           # regenerate the tables and examples from the AWS feed
$ make docs          # regenerate docs/
```

The acceptance tests run real Terraform through the plugin test harness, so
the Terraform CLI has to be on PATH. They write a temporary `.arn.hcl` and
point `ARN_CONFIG` at it.

`make tf-console` reads `arn.tf` and `.arn.hcl` in the repository root. Both
are gitignored; `arn.tf.sample` and `.arn.hcl.sample` are there to copy.

## Conventions

Comments and documentation are in English. Say what the code does and why a
non-obvious choice was made, not what a reader can see from the line itself.

Verify behaviour rather than describing what it should be. The provider is
easy to run: `make tf-console` gives a prompt in a second.
