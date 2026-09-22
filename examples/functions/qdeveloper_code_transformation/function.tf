# arn:aws:qdeveloper:ap-northeast-1:111111111111:codeTransformation/identifier
output "qdeveloper_code_transformation" {
  value = provider::arn::qdeveloper_code_transformation("identifier")
}
