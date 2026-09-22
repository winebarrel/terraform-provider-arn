# arn:aws:aws-external-anthropic:ap-northeast-1:111111111111:workspace/resource-id
output "aws_external_anthropic_workspace" {
  value = provider::arn::aws_external_anthropic_workspace("resource-id")
}
