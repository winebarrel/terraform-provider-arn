# arn:aws:fsx:ap-northeast-1:111111111111:association/file-system-id-or-file-cache-id/data-repository-association-id
output "fsx_association" {
  value = provider::arn::fsx_association("file-system-id-or-file-cache-id", "data-repository-association-id")
}
