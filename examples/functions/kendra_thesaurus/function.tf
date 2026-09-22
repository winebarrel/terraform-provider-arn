# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/thesaurus/thesaurus-id
output "kendra_thesaurus" {
  value = provider::arn::kendra_thesaurus("index-id", "thesaurus-id")
}
