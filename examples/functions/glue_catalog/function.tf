# arn:aws:glue:ap-northeast-1:111111111111:catalog/catalog-name
output "glue_catalog" {
  value = provider::arn::glue_catalog("catalog-name")
}
