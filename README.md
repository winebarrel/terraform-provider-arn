# terraform-provider-arn

[![CI](https://github.com/winebarrel/terraform-provider-arn/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/terraform-provider-arn/actions/workflows/ci.yml)
[![terraform docs](https://img.shields.io/badge/terraform-docs-%35835CC?logo=terraform)](https://registry.terraform.io/providers/winebarrel/arn/latest/docs)
[![codecov](https://codecov.io/gh/winebarrel/terraform-provider-arn/graph/badge.svg?token=YsOFKSDseA)](https://codecov.io/gh/winebarrel/terraform-provider-arn)
[![AI Generated](https://img.shields.io/badge/AI%20Generated-Claude-orange?logo=anthropic)](https://claude.com/claude-code)

Terraform functions that build AWS ARNs.

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

becomes

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

`.arn.hcl`, in the directory Terraform runs in. The file is required, though
an empty one is enough for ARNs that need nothing from it.

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

One function per resource type in the
[AWS service reference](https://servicereference.us-east-1.amazonaws.com),
named `<service>_<resource>`. The arguments are the ARN template's
placeholders in order, less the partition, region and account fields, which
come from the configuration file.

```hcl
# arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}
provider::arn::iam_role("my-role")

# arn:${Partition}:s3:::${BucketName}
provider::arn::s3_bucket("my-bucket")

# arn:${Partition}:access-analyzer:${Region}:${Account}:analyzer/${AnalyzerName}/archive-rule/${RuleName}
provider::arn::access_analyzer_archive_rule("my-analyzer", "my-rule")
```

The full list is in
[the documentation](https://registry.terraform.io/providers/winebarrel/arn/latest/docs).

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

An `account` block inherits the top-level values it does not set, so `prod`
above uses `ap-northeast-1`.

## Splitting the configuration

`import` reads one other file first. The path is relative to the file that
names it, or absolute.

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

The imported file is applied first and the importing file second. Above,
`region` is `us-east-1` and `account_id` is `111111111111`.

The position of the `import` line does not matter. HCL decodes the file as a
whole, so one file overrides another, not one line the lines above it.

A value that is not set does not override. An `account` block replaces one of
the same name instead of merging with it, so the replacement starts from the
top-level defaults again.

One file per `import`, no wildcards. An imported file cannot import in turn;
writing one is an error.

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
together. An unrecognized option, or a value set to an empty string, is an
error. Omit an option, or set it to `null`, to use the default.

## Validation

Values are checked for shape, not existence.

An ARN is `arn:partition:service:region:account:resource`. Each field is
checked wherever its value came from:

| ARN field | Set by | Rule |
|---|---|---|
| partition | `partition` | `aws`, `aws-` plus one or more words, or `*` |
| region | `region` | A name like `ap-northeast-1`, including `us-gov-west-1` and `us-iso-east-1`, or `*` |
| account | `account_id`, or the `account_id` of the block named by `account` | Twelve digits, `aws`, or `*` |
| resource | the function's arguments | Not empty |

`account` is not one of these. It names an `account` block, and an undeclared
name reports `unknown account`.

`aws` is the account AWS-managed policies use. `*` is valid in an ARN written
for an IAM policy. IAM Access Analyzer lists the supported partitions as
`*, aws, aws-cn, aws-us-gov`.

```hcl
provider::arn::iam_policy("AdministratorAccess", { account_id = "aws" })
# arn:aws:iam::aws:policy/AdministratorAccess

provider::arn::ec2_vpc("*", { region = "*" })
# arn:aws:ec2:*:111111111111:vpc/*

provider::arn::s3_object("my-bucket", "*", { partition = "*" })
# arn:*:s3:::my-bucket/*
```

A field is identified by its position, not by what AWS named the placeholder.
The account field is `${Account}` in most templates, `${AccountId}` in chime,
datasync and sso, `${ManagementAccountId}` in account and `${VpcOwnerAccount}`
in kafka. All of them come from the configuration:

```hcl
provider::arn::chime_meeting("m1")
# arn:aws:chime:ap-northeast-1:111111111111:meeting/m1
```

The same name in the resource part stays an argument. `organizations_account`
carries both in one template:

```hcl
# arn:${Partition}:organizations::${Account}:account/o-${OrganizationId}/${AccountId}
provider::arn::organizations_account("abc", "222222222222")
# arn:aws:organizations::111111111111:account/o-abc/222222222222
```

An argument in one of the five structural fields cannot contain a `:`. Inside
the resource part a colon is allowed:

```hcl
provider::arn::s3_object("my-bucket", "a:b/c.txt")
# arn:aws:s3:::my-bucket/a:b/c.txt

provider::arn::lambda_function_alias("fn", "PROD")
# arn:aws:lambda:ap-northeast-1:111111111111:function:fn:PROD
```

A slash is always allowed. IAM role paths, S3 object keys and log group names
contain them:

```hcl
provider::arn::iam_role("path/to/my-role")
# arn:aws:iam::111111111111:role/path/to/my-role
```

The configuration file is checked too, not only the call site.

## Development

```console
$ make test          # unit and acceptance tests
$ make tf-console    # terraform console against a locally built provider
$ make gen           # regenerate the function table from the AWS feed
$ make docs          # regenerate docs/
```

`make gen` refetches the AWS service reference and rewrites
`internal/arnspec/spec_<service>_gen.go`, one file per service. The tables are
committed, so the set of functions is fixed at build time. Regenerating is
idempotent, and a service that leaves the feed has its file deleted.
