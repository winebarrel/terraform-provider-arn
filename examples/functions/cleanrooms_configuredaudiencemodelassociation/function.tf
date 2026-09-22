# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/configuredaudiencemodelassociation/configured-audience-model-association-id
output "cleanrooms_configuredaudiencemodelassociation" {
  value = provider::arn::cleanrooms_configuredaudiencemodelassociation("membership-id", "configured-audience-model-association-id")
}
