# arn:aws:iam::111111111111:instance-profile/instance-profile-name-with-path
output "iam_instance_profile" {
  value = provider::arn::iam_instance_profile("instance-profile-name-with-path")
}
