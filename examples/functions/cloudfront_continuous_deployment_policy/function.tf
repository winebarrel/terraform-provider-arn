# arn:aws:cloudfront::111111111111:continuous-deployment-policy/id
output "cloudfront_continuous_deployment_policy" {
  value = provider::arn::cloudfront_continuous_deployment_policy("id")
}
