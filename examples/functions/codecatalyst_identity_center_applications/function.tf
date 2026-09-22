# arn:aws:codecatalyst:ap-northeast-1:111111111111:/identity-center-applications/identity-center-application-id
output "codecatalyst_identity_center_applications" {
  value = provider::arn::codecatalyst_identity_center_applications("identity-center-application-id")
}
