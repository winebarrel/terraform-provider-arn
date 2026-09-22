# arn:aws:odb:ap-northeast-1:111111111111:odb-network/odb-network-id
output "odb_odb_network" {
  value = provider::arn::odb_odb_network("odb-network-id")
}
