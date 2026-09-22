# arn:aws:panorama:ap-northeast-1:111111111111:package/package-id
output "panorama_package" {
  value = provider::arn::panorama_package("package-id")
}
