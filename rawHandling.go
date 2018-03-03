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

//Parse User to RawUser
func (u User) Parse() proto.Message {
	return &RawUser{
		Id:                        formatInt32(u.Id),
		Jobroles:                  formatInt32Slice(u.Jobroles),
		CareerLevel:               formatInt32(u.CareerLevel),
		DisciplineId:              formatInt32(u.DisciplineId),
		IndustryId:                formatInt32(u.IndustryId),
		Country:                   u.Country,
		Region:                    formatInt32(u.Region),
		ExperienceNEntriesClass:   formatInt32(u.ExperienceNEntriesClass),
		ExperienceYearsExperience: formatInt32(u.ExperienceYearsExperience),
		ExperienceYearsInCurrent:  formatInt32(u.ExperienceYearsInCurrent),
		EduDegree:                 formatInt32(u.EduDegree),
		EduFieldofstudies:         formatInt32Slice(u.EduFieldofstudies),
		Wtcj:                      formatInt32(u.Wtcj),
		Premium:                   formatInt32(u.Premium),
	}
}

//Parse Item to RawItem
func (u *Item) Parse() proto.Message {
	return &RawItem{
		Id:           formatInt32(u.Id),
		Title:        formatInt32(u.Title),
		CareerLevel:  formatInt32(u.CareerLevel),
		DisciplineId: formatInt32(u.DisciplineId),
		IndustryId:   formatInt32(u.IndustryId),
		Country:      u.Country,
		IsPayed:      formatInt32(u.IsPayed),
		Region:       formatInt32(u.Region),
		Latitude:     formatFloat32(u.Latitude),
		Longitude:    formatFloat32(u.Longitude),
		Employment:   formatInt32(u.Employment),
		Tags:         formatInt32Slice(u.Tags),
		CreatedAt:    formatInt64(u.CreatedAt),
	}
}

//Parse Interaction to RawInteraction
func (u *Interaction) Parse() proto.Message {
	return &RawInteraction{
		UserId:          formatInt32(u.UserId),
		ItemId:          formatInt32(u.ItemId),
		InteractionType: formatInt32(u.InteractionType),
		CreatedAt:       formatInt64(u.CreatedAt),
	}
}

//Parse TargetUser to RawTargetUser
func (u *TargetUser) Parse() proto.Message {
	return &RawTargetUser{
		UserId: formatInt32(u.UserId),
	}
}

//Parse TargetItem to RawTargetItem
func (u *TargetItem) Parse() proto.Message {
	return &RawTargetItem{
		ItemId: formatInt32(u.ItemId),
	}
}

func formatInt32(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func formatInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func formatInt32Slice(i []int32) string {
	s := []string{}
	for _, element := range i {
		s = append(s, formatInt32(element))
	}
	return strings.Join(s, ",")
}

func formatFloat32(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', 10, 64)
}
