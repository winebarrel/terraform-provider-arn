# arn:aws:managedblockchain:ap-northeast-1:111111111111:nodes/node-id
output "managedblockchain_node" {
  value = provider::arn::managedblockchain_node("node-id")
}
