# arn:aws:odb:ap-northeast-1:111111111111:odb-peering-connection/odb-peering-connection-id
output "odb_odb_peering_connection" {
  value = provider::arn::odb_odb_peering_connection("odb-peering-connection-id")
}
