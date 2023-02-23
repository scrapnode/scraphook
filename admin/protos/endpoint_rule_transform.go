package protos

import (
	"github.com/scrapnode/scraphook/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertEndpointRuleToRecord(rule *entities.EndpointRule) *EndpointRuleRecord {
	record := &EndpointRuleRecord{
		EndpointId: rule.EndpointId,
		Id:         rule.Id,
		Rule:       rule.Rule,
		Negative:   rule.Negative,
		Priority:   rule.Priority,
		CreatedAt:  timestamppb.New(time.UnixMilli(rule.CreatedAt)),
		UpdatedAt:  timestamppb.New(time.UnixMilli(rule.UpdatedAt)),
	}
	return record
}
