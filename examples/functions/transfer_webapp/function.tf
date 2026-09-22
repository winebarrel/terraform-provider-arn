# arn:aws:transfer:ap-northeast-1:111111111111:webapp/web-app-id
output "transfer_webapp" {
  value = provider::arn::transfer_webapp("web-app-id")
}
