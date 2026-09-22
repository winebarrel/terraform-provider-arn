# arn:aws:clouddirectory:ap-northeast-1:111111111111:directory/directory-id
output "clouddirectory_directory" {
  value = provider::arn::clouddirectory_directory("directory-id")
}
