# arn:aws:tnb:ap-northeast-1:111111111111:network-package/network-package-id
output "tnb_network_package" {
  value = provider::arn::tnb_network_package("network-package-id")
}
