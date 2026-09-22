# arn:aws:redshift:ap-northeast-1:111111111111:usagelimit:usage-limit-id
output "redshift_usagelimit" {
  value = provider::arn::redshift_usagelimit("usage-limit-id")
}
