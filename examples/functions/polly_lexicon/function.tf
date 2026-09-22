# arn:aws:polly:ap-northeast-1:111111111111:lexicon/lexicon-name
output "polly_lexicon" {
  value = provider::arn::polly_lexicon("lexicon-name")
}
