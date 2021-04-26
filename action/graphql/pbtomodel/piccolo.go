package pbtomodel

import (
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/pb"
)

// PbGroupToModelGroup : Change group of proto type to model
func PbGroupToModelGroup(group *pb.Group) *model.Group {
	modelGroup := &model.Group{
		ID:   group.Id,
		Name: group.Name,
	}

	return modelGroup
}
