# arn:aws:ssm:ap-northeast-1:111111111111:patchbaseline/patch-baseline-id-resource-id
output "ssm_patchbaseline" {
  value = provider::arn::ssm_patchbaseline("patch-baseline-id-resource-id")
}
