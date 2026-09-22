# arn:aws:qldb:ap-northeast-1:111111111111:stream/ledger-name/stream-id
output "qldb_stream" {
  value = provider::arn::qldb_stream("ledger-name", "stream-id")
}
