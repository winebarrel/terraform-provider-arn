# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/subscriptions/subscription-definition-id/versions/version-id
output "greengrass_subscription_definition_version" {
  value = provider::arn::greengrass_subscription_definition_version("subscription-definition-id", "version-id")
}
