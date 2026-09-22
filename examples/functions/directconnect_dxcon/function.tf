# arn:aws:directconnect:ap-northeast-1:111111111111:dxcon/connection-id
output "directconnect_dxcon" {
  value = provider::arn::directconnect_dxcon("connection-id")
}
