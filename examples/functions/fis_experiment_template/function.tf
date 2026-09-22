# arn:aws:fis:ap-northeast-1:111111111111:experiment-template/id
output "fis_experiment_template" {
  value = provider::arn::fis_experiment_template("id")
}
