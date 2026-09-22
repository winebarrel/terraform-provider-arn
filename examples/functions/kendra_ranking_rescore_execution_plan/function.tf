# arn:aws:kendra-ranking:ap-northeast-1:111111111111:rescore-execution-plan/rescore-execution-plan-id
output "kendra_ranking_rescore_execution_plan" {
  value = provider::arn::kendra_ranking_rescore_execution_plan("rescore-execution-plan-id")
}
