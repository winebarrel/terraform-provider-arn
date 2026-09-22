# arn:aws:memorydb:ap-northeast-1:111111111111:snapshot/snapshot-name
output "memorydb_snapshot" {
  value = provider::arn::memorydb_snapshot("snapshot-name")
}
