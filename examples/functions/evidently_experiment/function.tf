# arn:aws:evidently:ap-northeast-1:111111111111:project/project-name/experiment/experiment-name
output "evidently_experiment" {
  value = provider::arn::evidently_experiment("project-name", "experiment-name")
}
