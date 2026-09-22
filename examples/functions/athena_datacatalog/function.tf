# arn:aws:athena:ap-northeast-1:111111111111:datacatalog/data-catalog-name
output "athena_datacatalog" {
  value = provider::arn::athena_datacatalog("data-catalog-name")
}
