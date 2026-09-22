# arn:aws:drs:ap-northeast-1:111111111111:source-network/source-network-id
output "drs_source_network_resource" {
  value = provider::arn::drs_source_network_resource("source-network-id")
}
