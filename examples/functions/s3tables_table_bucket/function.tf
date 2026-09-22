# arn:aws:s3tables:ap-northeast-1:111111111111:bucket/table-bucket-name
output "s3tables_table_bucket" {
  value = provider::arn::s3tables_table_bucket("table-bucket-name")
}
