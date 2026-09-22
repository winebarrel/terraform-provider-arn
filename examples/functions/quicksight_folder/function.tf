# arn:aws:quicksight:ap-northeast-1:111111111111:folder/resource-id
output "quicksight_folder" {
  value = provider::arn::quicksight_folder("resource-id")
}
