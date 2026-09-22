# arn:aws:quicksight:ap-northeast-1:111111111111:vpcConnection/resource-id
output "quicksight_vpcconnection" {
  value = provider::arn::quicksight_vpcconnection("resource-id")
}
