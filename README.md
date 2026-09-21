# terraform-provider-arn

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
together, and an unrecognized option is an error rather than a silent
fallback to the default account.

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

## License

[MIT](LICENSE)
