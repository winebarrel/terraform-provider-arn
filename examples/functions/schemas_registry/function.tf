# arn:aws:schemas:ap-northeast-1:111111111111:registry/registry-name
output "schemas_registry" {
  value = provider::arn::schemas_registry("registry-name")
}
