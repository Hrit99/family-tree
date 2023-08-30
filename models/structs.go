package models

type GenderEnum int
type WedlockPositionEnum int

const (
	Male    GenderEnum = iota + 1 // EnumIndex = 1
	Female                        // EnumIndex = 2
	Unknown                       // EnumIndex = 3
)

// const (
// 	Father  WedlockPositionEnum = iota + 1 // EnumIndex = 1
// 	Mother                                 // EnumIndex = 2
// 	Child                                  // EnumIndex = 3
// )

func (g GenderEnum) String() string {
	return [...]string{"Male", "Female", "Unknown"}[g-1]
}

// func (g WedlockPositionEnum) String() string {
// 	return [...]string{"Father", "Mother", "Child"}[g-1]
// }

func (g GenderEnum) EnumIndex() int {
	return int(g)
}

// func (g WedlockPositionEnum) EnumIndex() int {
// 	return int(g)
// }

type Person struct {
	Id                 int                       `json:"id"`
	Name               string                    `json:"name"`
	Gender             string                    `json:"gender,omitempty"`
	RelationMapForward map[string]map[string]int `json:"relationMapForward,omitempty"`
}

// type Wedlock struct {
// 	Id        int   `json:"id"`
// 	Male      int   `json:"male,omitempty"`
// 	Female    int   `json:"female,omitempty"`
// 	Sons      []int `json:"sons,omitempty"`
// 	Daughters []int `json:"daughters,omitempty"`
// }

type Relation struct {
	Name           string `json:"name"`
	Opposite       string `json:"Opposite"`
	MaleOpposite   string `json:"maleOpposite"`
	FemaleOpposite string `json:"femaleOpposite"`
	Gender         string `json:"gender"`
}
