package apis

//
//type PersonalInformation struct { // types.pb.go文件中有标准的结构体
//	Name   string  `json:"name"`
//	Sex    string  `json:"sex"`
//	Tall   float64 `json:"tall"`
//	Weight float64 `json:"weight"`
//	Age    int     `json:"age"`
//
//
type PersonalInformationFatRate struct {
	Name    string
	FatRate float64
}

//
type PersonalRank struct {
	Name       string
	Sex        string
	RankNumber int
	FatRate    float64
}

func (*PersonalInformation) TableName() string { // 手动维护
	return "personal_information"
}
