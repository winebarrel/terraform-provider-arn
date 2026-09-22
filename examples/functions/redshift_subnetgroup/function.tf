# arn:aws:redshift:ap-northeast-1:111111111111:subnetgroup:subnet-group-name
output "redshift_subnetgroup" {
  value = provider::arn::redshift_subnetgroup("subnet-group-name")
}
