# arn:aws:logs:ap-northeast-1:111111111111:lookup-table:lookup-table-name
output "logs_lookup_table" {
  value = provider::arn::logs_lookup_table("lookup-table-name")
}
