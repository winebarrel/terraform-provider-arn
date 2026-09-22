# arn:aws:managedblockchain:ap-northeast-1:111111111111:accessors/accessor-id
output "managedblockchain_accessor" {
  value = provider::arn::managedblockchain_accessor("accessor-id")
}
