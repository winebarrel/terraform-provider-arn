# arn:aws:a4b:ap-northeast-1:111111111111:address-book/resource-id
output "a4b_addressbook" {
  value = provider::arn::a4b_addressbook("resource-id")
}
