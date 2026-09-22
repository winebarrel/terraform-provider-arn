# arn:aws:entityresolution:ap-northeast-1:111111111111:schemamapping/schema-name
output "entityresolution_schema_mapping" {
  value = provider::arn::entityresolution_schema_mapping("schema-name")
}
