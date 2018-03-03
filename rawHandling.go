package protocol

import (
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
)

//RawMessage handles the conversion of the unexpected json input to expected format
type RawMessage interface {
	proto.Message
	Parse() proto.Message
}

//Parse RawUser to User
func (u *RawUser) Parse() proto.Message {
	return &User{
		Id:                        toInt32(u.Id),
		Jobroles:                  toInt32Slice(u.Jobroles),
		CareerLevel:               toInt32(u.CareerLevel),
		DisciplineId:              toInt32(u.DisciplineId),
		IndustryId:                toInt32(u.IndustryId),
		Country:                   u.Country,
		Region:                    toInt32(u.Region),
		ExperienceNEntriesClass:   toInt32(u.ExperienceNEntriesClass),
		ExperienceYearsExperience: toInt32(u.ExperienceYearsExperience),
		ExperienceYearsInCurrent:  toInt32(u.ExperienceYearsInCurrent),
		EduDegree:                 toInt32(u.EduDegree),
		EduFieldofstudies:         toInt32Slice(u.EduFieldofstudies),
		Wtcj:                      toInt32(u.Wtcj),
		Premium:                   toInt32(u.Premium),
	}
}

//Parse RawItem to Item
func (u *RawItem) Parse() proto.Message {
	return &Item{
		Id:           toInt32(u.Id),
		Title:        toInt32(u.Title),
		CareerLevel:  toInt32(u.CareerLevel),
		DisciplineId: toInt32(u.DisciplineId),
		IndustryId:   toInt32(u.IndustryId),
		Country:      u.Country,
		IsPayed:      toInt32(u.IsPayed),
		Region:       toInt32(u.Region),
		Latitude:     toFloat32(u.Latitude),
		Longitude:    toFloat32(u.Longitude),
		Employment:   toInt32(u.Employment),
		Tags:         toInt32Slice(u.Tags),
		CreatedAt:    toInt64(u.CreatedAt),
	}
}

//Parse RawInteraction to Interaction
func (u *RawInteraction) Parse() proto.Message {
	return &Interaction{
		UserId:          toInt32(u.UserId),
		ItemId:          toInt32(u.ItemId),
		InteractionType: toInt32(u.InteractionType),
		CreatedAt:       toInt64(u.CreatedAt),
	}
}

//Parse RawTargetUser to TargetUser
func (u *RawTargetUser) Parse() proto.Message {
	return &TargetUser{
		UserId: toInt32(u.UserId),
	}
}

//Parse RawTargetItem to TargetItem
func (u *RawTargetItem) Parse() proto.Message {
	return &TargetItem{
		ItemId: toInt32(u.ItemId),
	}
}

func toInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

func toInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func toInt32Slice(s string) []int32 {
	if s == "" {
		return []int32{}
	}
	parts := strings.Split(s, ",")
	result := []int32{}
	for _, p := range parts {
		result = append(result, toInt32(p))
	}
	return result
}

func toFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0
	}
	return float32(f)
}
