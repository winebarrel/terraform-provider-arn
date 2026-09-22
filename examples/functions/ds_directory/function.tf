# arn:aws:ds:ap-northeast-1:111111111111:directory/directory-id
output "ds_directory" {
  value = provider::arn::ds_directory("directory-id")
}
