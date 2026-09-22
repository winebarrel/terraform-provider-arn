# arn:aws:a4b:ap-northeast-1:111111111111:contact/resource-id
output "a4b_contact" {
  value = provider::arn::a4b_contact("resource-id")
}
