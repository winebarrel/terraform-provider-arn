# terraform-provider-arn

[![CI](https://github.com/winebarrel/terraform-provider-arn/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/terraform-provider-arn/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/winebarrel/terraform-provider-arn/graph/badge.svg?token=YsOFKSDseA)](https://codecov.io/gh/winebarrel/terraform-provider-arn)

Terraform functions that build AWS ARNs.

Interpolating the account id into every ARN makes a configuration hard to
read:

```hcl
resource "aws_iam_role_policy" "example" {
  # ...
  policy = jsonencode({
    Statement = [{
      Effect   = "Allow"
      Action   = "sts:AssumeRole"
      Resource = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/my-role"
    }]
  })
}
```

With this provider:

```hcl
Resource = provider::arn::iam_role("my-role")
```

## Usage

```hcl
terraform {
  required_providers {
    arn = {
      source = "winebarrel/arn"
    }
  }
}

output "role" {
  value = provider::arn::iam_role("my-role")
}
```

`.arn.hcl`, in the directory Terraform runs in:

```hcl
account_id = "111111111111"
region     = "ap-northeast-1"
```

```console
$ terraform console
> provider::arn::iam_role("my-role")
"arn:aws:iam::111111111111:role/my-role"
```

Set `ARN_CONFIG` to read the configuration from somewhere else.

## Functions

There is one function per resource type in the
[AWS service reference](https://servicereference.us-east-1.amazonaws.com),
named `<service>_<resource>`. The arguments are the ARN template's
placeholders, in order, with `${Partition}`, `${Region}` and `${Account}`
taken from the configuration file.

```hcl
# arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}
provider::arn::iam_role("my-role")

# arn:${Partition}:s3:::${BucketName}
provider::arn::s3_bucket("my-bucket")

# arn:${Partition}:access-analyzer:${Region}:${Account}:analyzer/${AnalyzerName}/archive-rule/${RuleName}
provider::arn::access_analyzer_archive_rule("my-analyzer", "my-rule")
```

See [the documentation](https://registry.terraform.io/providers/winebarrel/arn/latest/docs)
for the full list.

## Multiple accounts

Name the other accounts in the configuration file:

```hcl
account_id = "111111111111"
region     = "ap-northeast-1"

account "prod" {
  account_id = "222222222222"
}

account "us" {
  account_id = "333333333333"
  region     = "us-east-1"
}
```

```hcl
provider::arn::iam_role("my-role", { account = "prod" })
# arn:aws:iam::222222222222:role/my-role

provider::arn::sqs_queue("my-queue", { account = "us" })
# arn:aws:sqs:us-east-1:333333333333:my-queue
```

An `account` block inherits every field it does not set, so `prod` above keeps
the top-level `ap-northeast-1`.

## Splitting the configuration

`import` reads one other file first. The path is relative to the file that
names it, or absolute:

```hcl
# .arn.hcl
import = "common.hcl"

region = "us-east-1"
```

```hcl
# common.hcl
account_id = "111111111111"
region     = "ap-northeast-1"

account "prod" {
  account_id = "222222222222"
}
```

The imported file is applied first and the importing file second, so a value
set closer to where you are reading wins. Above, `region` is `us-east-1`, and
`account_id` is the imported `111111111111` because nothing overrode it.

Where the `import` line sits in the file makes no difference. HCL decodes a
file as a whole, so this is one file overriding another, not one line
overriding the lines above it.

Only values that are actually set take part, so a file naming an account id
and nothing else leaves the region it inherited alone. An `account` block of
the same name replaces the imported one rather than merging with it, and so
starts again from the top-level defaults.

The path is one file. There is no wildcard, and an imported file cannot
itself import; saying so is an error rather than being ignored, since
otherwise its contents would be quietly unused with nothing to say why.

## Options

Every function takes an optional trailing map.

| Option | Meaning |
|---|---|
| `account` | Name of an `account` block in the configuration file |
| `account_id` | A literal account id, bypassing the `account` blocks |
| `region` | Region to use instead of the resolved one |
| `partition` | Partition to use instead of the resolved one |

```hcl
provider::arn::sqs_queue("my-queue", { region = "us-east-1" })
provider::arn::s3_bucket("my-bucket", { partition = "aws-cn" })
provider::arn::iam_role("my-role", { account_id = "999999999999" })
```

`partition` defaults to `aws`. `account` and `account_id` cannot be given
together. An unrecognized option is an error rather than a silent fallback to
the default account, and so is an explicitly empty value: write nothing, or
`null`, to use the default.

## Validation

The values are checked for shape, not existence. An account that does not
exist, or a region AWS has not built yet, is not something a string can be
asked about. What the checks catch is the value that could never be right,
which would otherwise be interpolated into a syntactically valid ARN and fail
much later, at apply time, somewhere unhelpful.

An ARN is `arn:partition:service:region:account:resource`. Each field is
checked wherever its value came from:

| ARN field | Set by | Rule |
|---|---|---|
| partition | `partition` | `aws`, `aws-` plus one or more words, or `*` |
| region | `region` | A name like `ap-northeast-1`, including `us-gov-west-1` and `us-iso-east-1`, or `*` |
| account | `account_id`, or the `account_id` of the block named by `account` | Twelve digits, `aws`, or `*` |
| resource | the function's arguments | Not empty |

`account` itself is not one of these: it names an `account` block, and a name
that is not declared is its own error rather than a malformed value.

`aws` is the account AWS-managed policies carry, and `*` is how an ARN written
for an IAM policy wildcards a field. IAM Access Analyzer reports the supported
partitions as `*, aws, aws-cn, aws-us-gov`.

```hcl
provider::arn::iam_policy("AdministratorAccess", { account_id = "aws" })
# arn:aws:iam::aws:policy/AdministratorAccess

provider::arn::ec2_vpc("*", { region = "*" })
# arn:aws:ec2:*:111111111111:vpc/*

provider::arn::s3_object("my-bucket", "*", { partition = "*" })
# arn:*:s3:::my-bucket/*
```

A rule applies wherever its field comes from. Most templates take the account
from the configuration, but a few spell it as a placeholder, so it arrives as
an argument instead: chime and datasync write `${AccountId}` rather than
`${Account}`. Those arguments are checked like any other account id, and which
ones they are follows from the field's position in the template rather than
from what AWS named the placeholder.

```hcl
provider::arn::chime_meeting("abc", "m1")
# invalid account id "abc"
```

An argument in one of the five structural fields cannot contain a `:` either,
because that would shift every field after it and name something else
entirely.

Inside the resource part a colon is just a character, and what it separates is
up to the service, so it is left alone:

```hcl
provider::arn::s3_object("my-bucket", "a:b/c.txt")
# arn:aws:s3:::my-bucket/a:b/c.txt

provider::arn::lambda_function_alias("fn", "PROD")
# arn:aws:lambda:ap-northeast-1:111111111111:function:fn:PROD
```

A slash is never a field separator, and IAM role paths, S3 object keys and log
group names all contain them:

```hcl
provider::arn::iam_role("path/to/my-role")
# arn:aws:iam::111111111111:role/path/to/my-role
```

The configuration file goes through the same check, so a typo there is
reported even when the call site overrides nothing.

## Why a file instead of a provider block

Provider-defined functions cannot read provider configuration, so values in a
`provider "arn"` block would never reach the functions. A file keeps them
visible and diffable, which an environment variable or an STS call made behind
your back would not.

The file is required. A few ARN shapes need nothing from it, an S3 bucket ARN
carrying neither account nor region, but letting those work without it would
mean the provider behaves differently depending on which function you happen
to call first:

```console
> provider::arn::s3_bucket("my-bucket")
Call to function "provider::arn::s3_bucket" failed: .arn.hcl not found:
create it, or point ARN_CONFIG at another path.
```

The contents are another matter. An empty file is enough for an ARN that needs
nothing, and a value only has to be set once something asks for it:

```console
> provider::arn::iam_role("my-role")
Call to function "provider::arn::iam_role" failed: iam_role needs an account
id: set account_id in the configuration file, or pass { account = ... }.
```

## Development

```console
$ make test          # unit and acceptance tests
$ make tf-console    # terraform console against a locally built provider
$ make gen           # regenerate the function table from the AWS feed
$ make docs          # regenerate docs/
```

`make gen` refetches the AWS service reference and rewrites
`internal/arnspec/spec_<service>_gen.go`, one file per service. The tables are
committed so that the set of functions is fixed at build time: the provider
declares its functions at startup, before any network call could be allowed to
fail, and a plan should not change because AWS published a new resource type
this morning. Regenerating is idempotent, and a service that leaves the feed
has its file deleted.
