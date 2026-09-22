# arn:aws:redshift:ap-northeast-1:111111111111:redshiftidcapplication:redshift-idc-application-id
output "redshift_redshiftidcapplication" {
  value = provider::arn::redshift_redshiftidcapplication("redshift-idc-application-id")
}
