# arn:aws:ds:ap-northeast-1:111111111111:directory/directory-id
output "ds_data_directory" {
  value = provider::arn::ds_data_directory("directory-id")
}
