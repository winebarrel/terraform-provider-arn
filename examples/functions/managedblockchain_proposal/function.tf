# arn:aws:managedblockchain:ap-northeast-1::proposals/proposal-id
output "managedblockchain_proposal" {
  value = provider::arn::managedblockchain_proposal("proposal-id")
}
