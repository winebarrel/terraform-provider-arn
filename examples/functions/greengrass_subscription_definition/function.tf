# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/subscriptions/subscription-definition-id
output "greengrass_subscription_definition" {
  value = provider::arn::greengrass_subscription_definition("subscription-definition-id")
}
