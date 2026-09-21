# Performance

A provider that declares 2321 functions is unusual, so this is what that
costs. Short version: about 66 ms and 23 MB per Terraform command, paid once
at startup, plus roughly half a millisecond per ARN built.

Measured on an Apple M4 Pro, macOS 26.7, Go 1.27.0, Terraform v1.14.8, against
2321 functions. Every figure is a median of seven runs. Treat them as orders
of magnitude rather than exact values; the shape is what matters.

## Startup

Terraform asks the plugin for its schema when it starts, which means
constructing all 2321 function definitions and sending them over gRPC.

| | `terraform plan`, one ARN | Plugin startup |
|---|---|---|
| No provider (plain strings) | 23 ms | |
| A two-function provider | 87 ms | 63 ms |
| This provider | 153 ms | 129 ms |

Most of the plugin startup cost is not the function count: launching any Go
plugin and completing the handshake is 63 ms of it. The 2319 extra function
definitions add **66 ms**.

The schema itself is **2.3 MB** of JSON, against 3.8 KB for the two-function
provider. That is the payload crossing the gRPC connection at startup.

In-process, building the definitions is far cheaper than shipping them:

```
BenchmarkProviderSchema    1.48 ms    3.8 MB    46039 allocs   // all 2321, name + definition
BenchmarkFunctionsOnly     22.5 us     75 KB     2322 allocs   // the closures alone
BenchmarkParseAll          0.93 ms    1.1 MB    28891 allocs   // parsing every template, once
```

So of the 66 ms, only about 1.5 ms is the provider doing work. The rest is
serialisation and transport.

## Per call

| ARNs built | `terraform plan` | Same outputs as plain strings | Difference |
|---|---|---|---|
| 1 | 154 ms | 23 ms | 131 ms |
| 100 | 215 ms | 31 ms | 184 ms |
| 1000 | 748 ms | 107 ms | 641 ms |
| 5000 | 3099 ms | 677 ms | 2422 ms |

The marginal cost is flat at about **0.6 ms per call**, of which roughly
0.13 ms is what Terraform spends on any output at all. So each function call
costs around **0.46 ms** over writing the string by hand.

Almost none of that is this provider. Building an ARN takes 40 nanoseconds:

```
BenchmarkBuild                        40.1 ns    96 B    1 alloc
BenchmarkBuildWithRegionAndAccount    42.4 ns    96 B    1 alloc
BenchmarkLookup                        5.3 ns     0 B    0 allocs
BenchmarkAll                           1.9 ns     0 B    0 allocs   // after the sync.Once
```

The gap between 40 ns and 0.46 ms is the gRPC round trip Terraform makes for
every function call. Nothing in this provider can shrink it.

The configuration file is read once per provider process, behind a
`sync.Once`, so it does not appear in the per-call cost.

## Memory

| | Peak RSS |
|---|---|
| No provider, one output | 54 MB |
| A two-function provider | 59 MB |
| This provider, one ARN | 82 MB |
| Plain strings, 5000 outputs | 301 MB |
| This provider, 5000 ARNs | 420 MB |

**23 MB** over the two-function provider, for the generated tables and the
2321 definitions. At 5000 calls the extra is 119 MB, which is the 2.3 MB
schema and the plan values Terraform holds either way, not the provider
growing with use.

The binary is 27.7 MB, against 26.2 MB for the two-function provider: the
generated tables add about 1.5 MB.

## What this means

For a typical configuration, a few dozen ARNs, the whole thing is 66 ms of
startup and 15 ms of calls. It is not worth thinking about.

The number to watch is the fixed 66 ms, since it is paid on every plan and
every apply whether or not the configuration calls a single function. It is
the price of one function per AWS resource type. Emitting only the common
services would cut it roughly in proportion, at the cost of having to decide
which services are common.

Building ARNs is not the slow part and never will be. A `data
"aws_caller_identity"` block, which this provider exists to avoid, is one call
to STS: a bare HTTPS round trip to the regional endpoint measures 27 ms from
here, before signing, and it is a dependency the graph has to wait on. This
provider makes no network calls at all.

## Reproducing

The Go benchmarks are in the tree:

```console
$ go test -run XXX -bench . -benchmem ./internal/arnspec/ ./internal/provider/
```

The Terraform figures come from timing `terraform plan -refresh=false` over
generated configurations with N outputs, each calling
`provider::arn::iam_role`, against configurations with the same outputs as
literal strings. Peak RSS is `/usr/bin/time -l`, which on macOS includes the
plugin process. The schema payload is:

```console
$ terraform providers schema -json | wc -c
```
