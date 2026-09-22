# arn:aws:fis:ap-northeast-1:111111111111:safety-lever/id
output "fis_safety_lever" {
  value = provider::arn::fis_safety_lever("id")
}
