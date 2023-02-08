package protos

import (
	"github.com/scrapnode/scraphook/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertEndpointToRecord(endpoint *entities.Endpoint) *EndpointRecord {
	record := &EndpointRecord{
		WorkspaceId: endpoint.WorkspaceId,
		WebhookId:   endpoint.WebhookId,
		Id:          endpoint.Id,
		Name:        endpoint.Name,
		Uri:         endpoint.Uri,
		CreatedAt:   timestamppb.New(time.UnixMilli(endpoint.CreatedAt)),
		UpdatedAt:   timestamppb.New(time.UnixMilli(endpoint.UpdatedAt)),
	}
	return record
}
