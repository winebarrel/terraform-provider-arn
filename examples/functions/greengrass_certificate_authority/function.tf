# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/groups/group-id/certificateauthorities/certificate-authority-id
output "greengrass_certificate_authority" {
  value = provider::arn::greengrass_certificate_authority("group-id", "certificate-authority-id")
}
