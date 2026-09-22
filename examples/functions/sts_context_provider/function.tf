# arn:aws:iam::aws:contextProvider/context-provider-name
output "sts_context_provider" {
  value = provider::arn::sts_context_provider("context-provider-name")
}
