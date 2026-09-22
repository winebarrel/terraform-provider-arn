# arn:aws:supportauthz:ap-northeast-1:111111111111:supportpermit/resource-id
output "supportauthz_supportpermit" {
  value = provider::arn::supportauthz_supportpermit("resource-id")
}
