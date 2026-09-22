# arn:aws:managedblockchain:ap-northeast-1:111111111111:invitations/invitation-id
output "managedblockchain_invitation" {
  value = provider::arn::managedblockchain_invitation("invitation-id")
}
