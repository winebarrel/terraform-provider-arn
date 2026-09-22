# arn:aws:qldb:ap-northeast-1:111111111111:ledger/ledger-name
output "qldb_ledger" {
  value = provider::arn::qldb_ledger("ledger-name")
}
