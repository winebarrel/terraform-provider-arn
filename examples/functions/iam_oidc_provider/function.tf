# arn:aws:iam::111111111111:oidc-provider/oidc-provider-name
output "iam_oidc_provider" {
  value = provider::arn::iam_oidc_provider("oidc-provider-name")
}
