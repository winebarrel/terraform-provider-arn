# arn:aws:qldb:ap-northeast-1:111111111111:ledger/ledger-name/information_schema/user_tables
output "qldb_catalog" {
  value = provider::arn::qldb_catalog("ledger-name")
}
