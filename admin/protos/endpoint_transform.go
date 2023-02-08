package protos

import "github.com/scrapnode/scraphook/entities"

func ConvertEndpointToRecord(endpoint *entities.Endpoint) *EndpointRecord {
	record := &EndpointRecord{
		WorkspaceId: endpoint.WorkspaceId,
		Id:          endpoint.Id,
		Name:        endpoint.Name,
		Uri:         endpoint.Uri,
		CreatedAt:   ConvertMilliToTimestamp(endpoint.CreatedAt),
		UpdatedAt:   ConvertMilliToTimestamp(endpoint.UpdatedAt),
	}
	return record
}
