# arn:aws:mgn:ap-northeast-1:111111111111:vcenter-client/vcenter-client-id
output "mgn_vcenter_client_resource" {
  value = provider::arn::mgn_vcenter_client_resource("vcenter-client-id")
}
