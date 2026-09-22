# arn:aws:quicksight:ap-northeast-1:111111111111:extension-access/resource-id
output "quicksight_extensionaccess" {
  value = provider::arn::quicksight_extensionaccess("resource-id")
}
