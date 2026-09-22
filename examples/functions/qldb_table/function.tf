# arn:aws:qldb:ap-northeast-1:111111111111:ledger/ledger-name/table/table-id
output "qldb_table" {
  value = provider::arn::qldb_table("ledger-name", "table-id")
}
