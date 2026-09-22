# arn:aws:medialive:ap-northeast-1:111111111111:offering:offering-id
output "medialive_offering" {
  value = provider::arn::medialive_offering("offering-id")
}
