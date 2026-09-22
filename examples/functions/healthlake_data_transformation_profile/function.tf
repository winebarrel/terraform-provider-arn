# arn:aws:healthlake:ap-northeast-1:111111111111:dataTransformationProfile/profile-id
output "healthlake_data_transformation_profile" {
  value = provider::arn::healthlake_data_transformation_profile("profile-id")
}
