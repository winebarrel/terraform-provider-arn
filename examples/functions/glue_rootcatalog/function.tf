# arn:aws:glue:ap-northeast-1:111111111111:catalog
output "glue_rootcatalog" {
  value = provider::arn::glue_rootcatalog()
}
