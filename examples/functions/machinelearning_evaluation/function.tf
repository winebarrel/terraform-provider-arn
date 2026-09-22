# arn:aws:machinelearning:ap-northeast-1:111111111111:evaluation/evaluation-id
output "machinelearning_evaluation" {
  value = provider::arn::machinelearning_evaluation("evaluation-id")
}
