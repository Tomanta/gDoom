package wad

type Thing struct {
	XPos   int16
	YPOS   int16
	Angle  int16
	TypeID int16
	Flags  int16
}

const (
	TH_SKILL_1_2 int16 = 1 << iota
	TH_SKILL_3
	TH_SKILL_4_5
	TH_AMBUSH
	TH_MULTI_ONLY
)
