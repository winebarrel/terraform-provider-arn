# arn:aws:directconnect:ap-northeast-1:111111111111:dxvif/virtual-interface-id
output "directconnect_dxvif" {
  value = provider::arn::directconnect_dxvif("virtual-interface-id")
}
