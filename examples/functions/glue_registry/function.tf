# arn:aws:glue:ap-northeast-1:111111111111:registry/registry-name
output "glue_registry" {
  value = provider::arn::glue_registry("registry-name")
}
