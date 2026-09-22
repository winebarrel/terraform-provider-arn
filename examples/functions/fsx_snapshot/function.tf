# arn:aws:fsx:ap-northeast-1:111111111111:snapshot/volume-id/snapshot-id
output "fsx_snapshot" {
  value = provider::arn::fsx_snapshot("volume-id", "snapshot-id")
}
