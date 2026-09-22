# arn:aws:redshift:ap-northeast-1:111111111111:parametergroup:parameter-group-name
output "redshift_parametergroup" {
  value = provider::arn::redshift_parametergroup("parameter-group-name")
}
