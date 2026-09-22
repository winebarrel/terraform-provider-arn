# arn:aws:s3tables:ap-northeast-1:111111111111:bucket/table-bucket-name/table/table-id
output "s3tables_table" {
  value = provider::arn::s3tables_table("table-bucket-name", "table-id")
}
