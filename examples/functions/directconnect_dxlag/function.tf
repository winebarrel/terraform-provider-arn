# arn:aws:directconnect:ap-northeast-1:111111111111:dxlag/lag-id
output "directconnect_dxlag" {
  value = provider::arn::directconnect_dxlag("lag-id")
}
