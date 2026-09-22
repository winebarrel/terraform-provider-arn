# arn:aws:drs:ap-northeast-1:111111111111:source-server/source-server-id
output "drs_source_server_resource" {
  value = provider::arn::drs_source_server_resource("source-server-id")
}
