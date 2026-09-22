# arn:aws:supportauthz:ap-northeast-1:111111111111:supportpermitrequest/resource-id
output "supportauthz_supportpermitrequest" {
  value = provider::arn::supportauthz_supportpermitrequest("resource-id")
}
