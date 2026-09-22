# arn:aws:fsx:ap-northeast-1:111111111111:volume/file-system-id/volume-id
output "fsx_volume" {
  value = provider::arn::fsx_volume("file-system-id", "volume-id")
}
