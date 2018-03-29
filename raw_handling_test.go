package protocol

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	fuzz "github.com/google/gofuzz"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserHandling(t *testing.T) {
	Convey("Given a user", t, func() {
		input := `{"id":"320","jobroles":"579315,1745826,3637189","career_level":"4","discipline_id":"0","industry_id":"0","country":"de","region":"0","experience_n_entries_class":"3","experience_years_experience":"6","experience_years_in_current":"4","edu_degree":"0","edu_fieldofstudies":"3","wtcj":"1","premium":"0"}`
		r := strings.NewReader(input)
		ru := &RawUser{}

		jsonpb.Unmarshal(r, ru)
		u := ru.Parse()
		user := *(u.(*User))
		expected := User{
			Id:                        320,
			Jobroles:                  []int32{579315, 1745826, 3637189},
			CareerLevel:               4,
			DisciplineId:              0,
			IndustryId:                0,
			Country:                   "de",
			Region:                    0,
			ExperienceNEntriesClass:   3,
			ExperienceYearsExperience: 6,
			ExperienceYearsInCurrent:  4,
			EduDegree:                 0,
			EduFieldofstudies:         []int32{3},
			Wtcj:                      1,
			Premium:                   0,
		}
		So(user, ShouldResemble, expected)
		So(user.Parse().(*RawUser), ShouldResemble, ru)
	})
}
func TestConversionFuncs(t *testing.T) {
	Convey("Conversion Funcs", t, func() {
		Convey("Given an empty input", func() {
			So(toInt32(""), ShouldBeZeroValue)
			So(toInt64(""), ShouldBeZeroValue)
			So(toFloat32(""), ShouldBeZeroValue)
			So(toInt32Slice(""), ShouldBeEmpty)
			So(formatInt32(0), ShouldEqual, "0")
			So(formatInt64(0), ShouldEqual, "0")
			So(formatFloat32(0.0), ShouldStartWith, "0.0")
			So(formatInt32Slice([]int32{}), ShouldEqual, "")
		})

		Convey("Given an input", func() {
			So(toInt32("11"), ShouldEqual, 11)
			So(toInt64("22"), ShouldEqual, 22)
			So(toFloat32("33.89"), ShouldEqual, 33.89)
			So(toInt32Slice("1,55,88"), ShouldResemble, []int32{1, 55, 88})
			So(formatInt32(11), ShouldEqual, "11")
			So(formatInt64(22), ShouldEqual, "22")
			So(formatFloat32(33.3), ShouldStartWith, "33.")
			So(formatInt32Slice([]int32{22, 33, 64}), ShouldEqual, "22,33,64")
		})
	})
}

var rU proto.Message

func BenchmarkUserConversion(b *testing.B) {
	var user proto.Message
	random := fuzz.New().NilChance(0)
	rawUser := &RawUser{}
	random.Fuzz(rawUser)
	for n := 0; n < b.N; n++ {
		user = rawUser.Parse()
	}
	rU = user
}

var rS string

func BenchmarkFormatInt32Slice(b *testing.B) {
	var result string
	input := &[]int32{}
	random := fuzz.New().NumElements(100, 100)
	random.Fuzz(input)
	for n := 0; n < b.N; n++ {
		result = formatInt32Slice(*input)
	}
	rS = result
}
