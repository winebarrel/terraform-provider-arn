// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kendra
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kendra/kendra.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kendra_access_control_configuration", Service: "kendra", Resource: "access-control-configuration", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/access-control-configuration/${AccessControlConfigurationId}"},
		{Name: "kendra_data_source", Service: "kendra", Resource: "data-source", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/data-source/${DataSourceId}"},
		{Name: "kendra_experience", Service: "kendra", Resource: "experience", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/experience/${ExperienceId}"},
		{Name: "kendra_faq", Service: "kendra", Resource: "faq", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/faq/${FaqId}"},
		{Name: "kendra_featured_results_set", Service: "kendra", Resource: "featured-results-set", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/featured-results-set/${FeaturedResultsSetId}"},
		{Name: "kendra_index", Service: "kendra", Resource: "index", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}"},
		{Name: "kendra_query_suggestions_block_list", Service: "kendra", Resource: "query-suggestions-block-list", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/query-suggestions-block-list/${QuerySuggestionsBlockListId}"},
		{Name: "kendra_thesaurus", Service: "kendra", Resource: "thesaurus", Template: "arn:${Partition}:kendra:${Region}:${Account}:index/${IndexId}/thesaurus/${ThesaurusId}"},
	})
}
