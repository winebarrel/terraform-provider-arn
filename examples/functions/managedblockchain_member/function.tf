# arn:aws:managedblockchain:ap-northeast-1:111111111111:members/member-id
output "managedblockchain_member" {
  value = provider::arn::managedblockchain_member("member-id")
}
