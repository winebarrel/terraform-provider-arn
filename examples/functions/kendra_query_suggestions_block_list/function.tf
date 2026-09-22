# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/query-suggestions-block-list/query-suggestions-block-list-id
output "kendra_query_suggestions_block_list" {
  value = provider::arn::kendra_query_suggestions_block_list("index-id", "query-suggestions-block-list-id")
}
