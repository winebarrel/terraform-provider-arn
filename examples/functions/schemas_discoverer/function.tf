# arn:aws:schemas:ap-northeast-1:111111111111:discoverer/discoverer-id
output "schemas_discoverer" {
  value = provider::arn::schemas_discoverer("discoverer-id")
}
