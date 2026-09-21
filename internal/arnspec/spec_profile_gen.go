// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: profile
// Source: https://servicereference.us-east-1.amazonaws.com/v1/profile/profile.json
// Functions: 12
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "profile_calculated_attributes", Service: "profile", Resource: "calculated-attributes", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/calculated-attributes/${CalculatedAttributeName}"},
		{Name: "profile_domain_object_types", Service: "profile", Resource: "domain-object-types", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/domain-object-types/${ObjectTypeName}"},
		{Name: "profile_domains", Service: "profile", Resource: "domains", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}"},
		{Name: "profile_event_streams", Service: "profile", Resource: "event-streams", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/event-streams/${EventStreamName}"},
		{Name: "profile_event_triggers", Service: "profile", Resource: "event-triggers", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/event-triggers/${EventTriggerName}"},
		{Name: "profile_integrations", Service: "profile", Resource: "integrations", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/integrations/${Uri}"},
		{Name: "profile_layouts", Service: "profile", Resource: "layouts", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/layouts/${LayoutDefinitionName}"},
		{Name: "profile_object_types", Service: "profile", Resource: "object-types", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/object-types/${ObjectTypeName}"},
		{Name: "profile_recommender_filters", Service: "profile", Resource: "recommender-filters", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/recommender-filters/${RecommenderFilterName}"},
		{Name: "profile_recommender_schemas", Service: "profile", Resource: "recommender-schemas", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/recommender-schemas/${RecommenderSchemaName}"},
		{Name: "profile_recommenders", Service: "profile", Resource: "recommenders", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/recommenders/${RecommenderTypeName}"},
		{Name: "profile_segment_definitions", Service: "profile", Resource: "segment-definitions", Template: "arn:${Partition}:profile:${Region}:${Account}:domains/${DomainName}/segment-definitions/${SegmentDefinitionName}"},
	})
}
