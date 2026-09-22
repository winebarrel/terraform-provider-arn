# arn:aws:mobiletargeting:ap-northeast-1:111111111111:recommenders/*
output "mobiletargeting_recommenders" {
  value = provider::arn::mobiletargeting_recommenders()
}
