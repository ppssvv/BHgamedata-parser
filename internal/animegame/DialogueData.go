package animegame

import (
	"encoding/binary"
	"encoding/json"

	bin "github.com/streamingfast/binary"
)

type DialogueData struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []DialogueDataEntry
}

func (s *DialogueData) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type DialogueDataEntry struct {
	ID               int32
	Addr1            uint32 `json:"-"`
	JumpID           int32
	Addr2            uint32 `json:"-"`
	DialogueType     int8
	InputID          int32
	CgRawPos         int32
	Addr3            uint32 `json:"-"`
	AvatarID         int32
	DressID          int32
	AvatarViceKey    int32
	ScreenSide       int8
	FaceID           int8
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	AnimationID      int8
	Distortion       int8
	Transparency     float32
	Addr7            uint32 `json:"-"`
	Addr8            uint32 `json:"-"`
	Addr9            uint32 `json:"-"`
	ImageID          int32
	Addr10           uint32 `json:"-"`
	Addr11           uint32 `json:"-"`
	Addr12           uint32 `json:"-"`
	BackgroundEffect int8
	EnterEffect      int8
	Addr13           uint32 `json:"-"`
	AudioDelay       float32
	Addr14           uint32 `json:"-"`
	FaceVersion      int8
	Addr15           uint32 `json:"-"`
	Addr16           uint32 `json:"-"`
	Addr17           uint32 `json:"-"`
	Addr18           uint32 `json:"-"`
	Unknown4         int32
	Unknown5         int32
	Addr19           uint32 `json:"-"`
	Addr20           uint32 `json:"-"`
	Addr21           uint32 `json:"-"`

	PreDialogueIDList    ArrUint32
	LeafIDList           ArrUint32
	AvatarName           StrWithPrefix16
	Face                 StrWithPrefix16
	FaceType             StrWithPrefix16
	FaceEffect           StrWithPrefix16
	BGMCover             StrWithPrefix16
	BGMVolumeControlList ArrStrWithPrefix16
	CgID                 StrWithPrefix16
	Content              DialogueContent
	Background           StrWithPrefix16
	BackgroundCG         StrWithPrefix16
	AudioID              StrWithPrefix16
	LipMotion            StrWithPrefix16
	ScreenEffect         StrWithPrefix16
	Unknown1             StrWithPrefix16
	Unknown2             ArrStrWithPrefix16
	Unknown3             ArrStrWithPrefix16
	PostDialogueIDList   ArrUint32
	TalkerName           StrWithPrefix16
	QuestionContent      StrWithPrefix16
}

type DialogueContent []struct {
	Text     string
	Duration float32
}

func (c *DialogueContent) UnmarshalBinary(decoder *bin.Decoder) error {
	count, err := decoder.ReadInt32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*c = make([]struct {
		Text     string
		Duration float32
	}, count)

	for i := int32(0); i < count; i++ {
		var tmpStr StrWithPrefix16
		err := tmpStr.UnmarshalBinary(decoder)
		if err != nil {
			return err
		}
		b, err := decoder.ReadFloat32(binary.LittleEndian)
		if err != nil {
			return err
		}

		(*c)[i] = struct {
			Text     string
			Duration float32
		}{string(tmpStr), b}
	}

	return nil
}
