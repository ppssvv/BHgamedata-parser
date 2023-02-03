package dump

type EditorUniqueMonsterMetaData struct {
	// Fields
	UniqueID uint32

	// Properties
	Addr1          uint32 `json:"-"`
	Group          uint32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	HPRatio        float32
	AttackRatio    float32
	DefenseRatio   float32
	MoveSpeedRatio float32
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	Addr10         uint32 `json:"-"`
	Addr11         uint32 `json:"-"`
	HPSegmentNum   int32
	Addr12         uint32 `json:"-"`
	BossRank       int32
	HandBookId     int32

	// Objects
	Info          StrWithPrefix16
	Category      StrWithPrefix16
	Name          Hash
	MonsterName   StrWithPrefix16
	TypeName      StrWithPrefix16
	ATKRatios     []float32
	ConfigType    StrWithPrefix16
	AIName        StrWithPrefix16
	AttackCDNames []StrWithPrefix16
	AttackCDs     []float32
	Abilities     StrWithPrefix16
	Scale         []float32
}

type MonsterActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ActivityName StrWithPrefix16
}

type MonsterProtoTypeMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	HPRatio        float32
	AttackRatio    float32
	DefenseRatio   float32
	MoveSpeedRatio float32
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	Addr10         uint32 `json:"-"`
	Addr11         uint32 `json:"-"`
	Addr12         uint32 `json:"-"`
	HPSegmentNum   int32
	Addr13         uint32 `json:"-"`
	BossRank       int32
	HandBookId     int32

	// Objects
	Info          StrWithPrefix16
	Details       StrWithPrefix16
	Category      StrWithPrefix16
	Name          Hash
	MonsterName   StrWithPrefix16
	TypeName      StrWithPrefix16
	ATKRatios     []float32
	ConfigType    StrWithPrefix16
	AIName        StrWithPrefix16
	AttackCDNames []StrWithPrefix16
	AttackCDs     []float32
	Abilities     StrWithPrefix16
	Scale         []float32
}

type PropListMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	PropName StrWithPrefix16
	Info     StrWithPrefix16
	Details  StrWithPrefix16
	Category StrWithPrefix16
}

type EditorUniqueMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []EditorUniqueMonsterMetaData
}

type MonsterActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonsterActivityMetaData
}

type MonsterProtoTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonsterProtoTypeMetaData
}

type PropListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []PropListMetaData
}

type ArtifactSkillNameOverrideMetaData struct {
	// Fields
	SubSkillID uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	OverrideName Hash
}

type AvatarArtifactLevelMetaData struct {
	// Fields
	ID int32

	// Properties
	Level             int32
	UnlockAvatarStar  int32
	UnlockAvatarLevel int32
	Addr1             uint32 `json:"-"`

	// Objects
	UnlockMaterial []AvatarArtifactLevelMetaData_UpgradeMaterialItem
}

type AvatarArtifactLevelMetaData_UpgradeMaterialItem struct {
	// Fields
	Amount int32
	ItemID int32
}

type AvatarArtifactMetaData struct {
	// Fields
	ArtifactID int32

	// Properties
	ShowPlayLevel     int32
	UnlockAvatarLevel int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	Addr12            uint32 `json:"-"`

	// Objects
	ArtifactName        Hash
	ArtifactDes         Hash
	ArtifactPrefab      StrWithPrefix16
	LevelList           []int32
	ReplaceSkillList    []AvatarArtifactMetaData_ReplaceSkillPair
	DressBanner         StrWithPrefix16
	DressBannerDes1     Hash
	DressBannerDes2     Hash
	BannerAvatarImgPath StrWithPrefix16
	BannerTitleImgPath  StrWithPrefix16
	ShowSkillList       []AvatarArtifactMetaData_LockedDisplaySkillInfo
	TagUnlockList       []int32
}

type AvatarArtifactMetaData_ReplaceSkillPair struct {
	// Fields
	NewSkillID int32
	OldSkillID int32
}

type AvatarArtifactMetaData_LockedDisplaySkillInfo struct {
	// Fields
	SkillID int32
	Addr1   uint32 `json:"-"`

	// Objects
	SkillDescriptionTextID Hash
}

type AvatarCardMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity              int32
	MaxRarity           int32
	Cost                int32
	MaxLv               int32
	SellPriceBase       float32
	SellPriceAdd        float32
	GearExpProvideBase  float32
	GearExpPorvideAdd   float32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	CharacterExpProvide float32
	GachaDisplayRarity  int32
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`

	// Objects
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
}

type AvatarFragmentMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity              int32
	MaxRarity           int32
	Cost                int32
	MaxLv               int32
	SellPriceBase       float32
	SellPriceAdd        float32
	GearExpProvideBase  float32
	GearExpPorvideAdd   float32
	BaseType            int32
	SortID              uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	CharacterExpProvide float32
	QuantityLimit       int32
	Addr5               uint32 `json:"-"`
	ReturnMaterialID    int32
	ReturnNum           int32
	TagType             uint8
	GachaDisplayRarity  int32
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`

	// Objects
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	LinkIDList                 []uint32
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
}

type AvatarLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp              int32
	Cost             int32
	AvatarAssistConf float32
	SubSkillScoin    int32
}

type AvatarMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	ClassID             int32
	RoleID              int32
	AvatarType          uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	UnlockStar          int32
	Addr7               uint32 `json:"-"`
	Attribute           int32
	InitialWeapon       int32
	AvatarCardID        int32
	AvatarFragmentID    int32
	ArtifactFragmentID  int32
	UltraSkillID        int32
	CaptainSkillID      int32
	SKL01SP             float32
	SKL01SPNeed         float32
	SKL01Charges        int32
	SKL01CD             float32
	SKL02SP             float32
	SKL02SPNeed         float32
	SKL02Charges        int32
	SKL02CD             float32
	SKL03SP             float32
	SKL03SPNeed         float32
	SKL03Charges        int32
	SKL03CD             float32
	SKL02ArtifactCD     float32
	SKL02ArtifactSP     float32
	SKL02ArtifactSPNeed float32
	BaseAvatarID        int32
	Addr8               uint32 `json:"-"`
	Addr9               uint32 `json:"-"`
	Addr10              uint32 `json:"-"`
	Addr11              uint32 `json:"-"`
	Addr12              uint32 `json:"-"`
	Addr13              uint32 `json:"-"`
	Addr14              uint32 `json:"-"`
	Addr15              uint32 `json:"-"`
	ArtifactID          int32
	IsEasterner         bool
	Addr16              uint32 `json:"-"`
	Addr17              uint32 `json:"-"`
	Addr18              uint32 `json:"-"`
	DefaultDressId      int32
	AvatarStarUpType    uint8
	Addr19              uint32 `json:"-"`
	IsCollaboration     bool
	Addr20              uint32 `json:"-"`

	// Objects
	FullName               Hash
	ShortName              Hash
	RomaName               Hash
	Desc                   Hash
	AvatarRegistryKey      StrWithPrefix16
	WeaponBaseTypeList     []int32
	SkillList              []int32
	FirstName              Hash
	LastName               Hash
	EnFirstName            Hash
	EnLastName             Hash
	UISelectVoice          StrWithPrefix16
	UILevelUpVoice         StrWithPrefix16
	DA_Name                StrWithPrefix16
	DA_Type                StrWithPrefix16
	FaceAnimationGroupName StrWithPrefix16
	AvatarEffects          []int32
	TagUnlockList          []int32
	AvatarStarSourceID     []int32
	StarUpBG               StrWithPrefix16
}

type AvatarModelMetaData struct {
	// Fields
	ID int32

	// Properties
	ModelID          int32
	AvatarID         uint16
	AvatarLV         uint8
	AvatarStar       uint8
	Addr1            uint32 `json:"-"`
	AvatarDress      int32
	WeaponID         int32
	WeaponLV         uint8
	StigmataID1      int32
	StigmataLV1      uint8
	StigmataID2      int32
	StigmataLV2      uint8
	StigmataID3      int32
	StigmataLV3      uint8
	IsAvatarArtifact bool
	ArtifactStar     int32
	UseGrandKey      bool
	UseMedal         bool
	UseIsland        bool
	UseNewIslandBuff bool

	// Objects
	AvatarSubSkills []AvatarModelMetaData_SkillInfoItem
}

type AvatarModelMetaData_SkillInfoItem struct {
	// Fields
	SkillID int32
	SkillLV int32
}

type AvatarSampleMetaData struct {
	// Fields
	AvatarSampleID int32

	// Properties
	ModelID         int32
	AvatarID        int32
	MaterialID      int32
	Type            int32
	ForceLevel      int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	StageTypeEffect int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`

	// Objects
	StageTypeList  []uint8
	StageTypeList2 []uint8
	StageIDList    []int32
	SampleDesc     StrWithPrefix16
}

type AvatarSkillMetaData struct {
	// Fields
	SkillId int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	ShowOrder       int32
	UnlockLv        int32
	UnlockStar      int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	ParamBase_1     float32
	ParamLogic_1    string
	ParamSubID_1    int32
	ParamSubIndex_1 int32
	ParamBase_2     float32
	ParamLogic_2    string
	ParamSubID_2    int32
	ParamSubIndex_2 int32
	ParamBase_3     float32
	ParamLogic_3    string
	ParamSubID_3    int32
	ParamSubIndex_3 int32
	CanTry          int32
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`

	// Objects
	Name            Hash
	Info            Hash
	SkillStep       Hash
	IconPath        StrWithPrefix16
	IconPathInLevel StrWithPrefix16
	ButtonName      StrWithPrefix16
	TagList         []AvatarSkillMetaData_Tag
	UnlockItemList  []AvatarSkillMetaData_UnlockItem
}

type AvatarSkillMetaData_Tag struct {
	// Fields
	Strength uint8
	TagID    int32
	Addr1    uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type AvatarSkillMetaData_UnlockItem struct {
	// Fields
	Amount int32
	ItemID int32
}

type AvatarStarMetaData struct {
	// Fields
	AvatarID int32
	Star     int32
	SubStar  int32

	// Properties
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`
	HpBase  float32
	HpAdd   float32
	SpBase  float32
	SpAdd   float32
	AtkBase float32
	AtkAdd  float32
	DfsBase float32
	DfsAdd  float32
	CrtBase float32
	CrtAdd  float32
	Addr4   uint32 `json:"-"`
	Addr5   uint32 `json:"-"`
	Addr6   uint32 `json:"-"`
	Addr7   uint32 `json:"-"`
	Addr8   uint32 `json:"-"`
	Addr9   uint32 `json:"-"`
	Addr10  uint32 `json:"-"`

	// Objects
	IconPath              StrWithPrefix16
	IconPathInLevel       StrWithPrefix16
	FigurePath            StrWithPrefix16
	ChibiIconPath         StrWithPrefix16
	SkillID               []int32
	SkillUnlockType       []int32
	SkillNodeName         Hash
	SkillNodeDesc         Hash
	SkillUnlockSketchList []Hash
	SkillUnlockDescList   []Hash
}

type AvatarStarTypeMetaData struct {
	// Fields
	Star             uint8
	SubStar          uint8
	AvatarType       uint8
	AvatarStarUpType uint8

	// Properties
	Upgrade uint32
	Addr1   uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type AvatarStarUpSubSkillData struct {
	// Fields
	SkillID int32
	Star    uint8
	SubStar uint8

	// Properties
	ParamBase_1 float32
	ParamBase_2 float32
	ParamBase_3 float32
	ParamAdd_1  float32
	ParamAdd_2  float32
	ParamAdd_3  float32
}

type AvatarSubSkillMetaData struct {
	// Fields
	AvatarSubSkillId uint16

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	ShowOrder       uint8
	SkillId         uint16
	IgnoreLeader    bool
	Addr4           uint32 `json:"-"`
	UnlockStar      uint8
	UnlockSubStar   uint8
	UnlockLv        uint8
	UnlockLvAdd     uint8
	MaxLv           uint8
	Addr5           uint32 `json:"-"`
	ScoinCalc       bool
	UnlockScoin     int32
	ScoinLvAdd      uint8
	ItemType        uint8
	SkillToggle     bool
	ParamBase_1     float32
	ParamAdd_1      float32
	ParamBase_2     float32
	ParamAdd_2      float32
	ParamBase_3     float32
	ParamAdd_3      float32
	CanTry          bool
	ArtifactSkillID uint16
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`

	// Objects
	Name                    Hash
	Info                    Hash
	Brief                   Hash
	IconPath                StrWithPrefix16
	UpLevelSubStarNeedList  []AvatarSubSkillMetaData_UpLevelStarNeed
	UpLevelArtifactNeedList []AvatarSubSkillMetaData_ArtifactSubSkillLevelUpCondition
	TagList                 []AvatarSubSkillMetaData_Tag
}

type AvatarSubSkillMetaData_UpLevelStarNeed struct {
	// Fields
	Level       uint8
	StarNeed    uint8
	SubStarNeed uint8
}

type AvatarSubSkillMetaData_ArtifactSubSkillLevelUpCondition struct {
	// Fields
	ArtifactLevel uint8
	SubSkillLevel uint8
}

type AvatarSubSkillMetaData_Tag struct {
	// Fields
	Strength uint8
	TagID    int32
	Addr1    uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type AvatarTagUnLockMetaData struct {
	// Fields
	AvatarTagID int32

	// Properties
	TagID int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	TagSecondDesc        Hash
	UnlockSkillIDList    []int32
	UnlockSubSkillIDList []int32
}

type CollaborationStigmataMetaData struct {
	// Fields
	SetID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	OptionalStigmataList      []int32
	CollaborationStigmataList []int32
	StigmataSetName           Hash
	StigmataSetIcon           StrWithPrefix16
	OptionalStigmataTextmap   Hash
	OptionalStigmataShowTime  RemoteTime
}

type CustomHeadMetaData struct {
	// Fields
	HeadID int32

	// Properties
	IndexID     int32
	Page        int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Show        bool
	HeadType    int32
	HeadParaInt int32
	TimeType    int32
	During      int32
	Addr5       uint32 `json:"-"`
	Addr6       uint32 `json:"-"`
	Rarity      int32

	// Objects
	HeadTitle       StrWithPrefix16
	HeadDescription StrWithPrefix16
	IconPath        StrWithPrefix16
	ImagePath       StrWithPrefix16
	EndTime         StrWithPrefix16
	BgColorPath     StrWithPrefix16
}

type DLCAvatarCardMetaData struct {
	// Fields
	DLCCardID int32

	// Properties
	Rarity int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`

	// Objects
	DisplayTitle       Hash
	DisplayDescription Hash
	IconPath           StrWithPrefix16
	ImagePath          StrWithPrefix16
}

type DLCAvatarMetaData struct {
	// Fields
	DLCAvatarID int32

	// Properties
	Index             int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	UnlockHintStoryID int32
	RotationFront     float32
	RotationSide      float32
	Addr3             uint32 `json:"-"`

	// Objects
	StoryBGImg      StrWithPrefix16
	StoryDesc       Hash
	TalentTattooImg StrWithPrefix16
}

type DormitoryFurnitureMetaData struct {
	// Fields
	FurnitureID int32

	// Properties
	Rarity          uint8
	Addr1           uint32 `json:"-"`
	Type            uint8
	EditType        uint8
	InteractiveType uint8
	InitialUnLock   uint8
	SuitId          uint8
	Addr2           uint32 `json:"-"`
	ManualScale     float32
	ComfortType     uint8
	ComfortValue    uint8
	HcoinCost       uint16
	HcoinUnlock     uint16
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	UnlockAdd       bool
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`
	KeepLimit       uint8
	Addr10          uint32 `json:"-"`
	Addr11          uint32 `json:"-"`
	CollectionType  uint8

	// Objects
	FurnitureTag            []uint8
	Size                    []uint8
	CostItemList            []DormitoryFurnitureMetaData_Material
	UnlockItemList          []DormitoryFurnitureMetaData_Material
	FurnitureModPath        StrWithPrefix16
	FurnitureNameText       Hash
	FurnitureIconPath       StrWithPrefix16
	FurnitureDescText       Hash
	FurnitureSourceDescText Hash
	DeleteMaterialList      []DormitoryFurnitureMetaData_Material
	VerandaPath             StrWithPrefix16
}

type DormitoryFurnitureMetaData_Material struct {
	// Fields
	MaterialID     int32
	MaterialNumber int32
}

type DressMetaData struct {
	// Fields
	DressID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Rarity          int32
	Addr2           uint32 `json:"-"`
	RoleID          int32
	Addr3           uint32 `json:"-"`
	DressType       int32
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Coin            int32
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`
	Addr10          uint32 `json:"-"`
	Addr11          uint32 `json:"-"`
	Addr12          uint32 `json:"-"`
	Addr13          uint32 `json:"-"`
	Show            int32
	PreviewStageID  int32
	Addr14          uint32 `json:"-"`
	Addr15          uint32 `json:"-"`
	Addr16          uint32 `json:"-"`
	Addr17          uint32 `json:"-"`
	ArtifactID      int32
	Addr18          uint32 `json:"-"`
	Addr19          uint32 `json:"-"`
	Addr20          uint32 `json:"-"`
	Addr21          uint32 `json:"-"`
	Addr22          uint32 `json:"-"`
	IsCollaboration bool

	// Objects
	Name                       Hash
	AvatarIDList               []int32
	IconPath                   StrWithPrefix16
	GetWay                     Hash
	Desc                       Hash
	DressResource              StrWithPrefix16
	CardPath                   StrWithPrefix16
	TachiePath                 StrWithPrefix16
	AvatarIconPath             StrWithPrefix16
	AvatarIconSidePath         StrWithPrefix16
	ImagePath                  StrWithPrefix16
	LinkIDList                 []uint32
	AvatarGachaFigure          StrWithPrefix16
	DressTagList               []int32
	RechargeShowName           Hash
	RechargeShowAvatarName     Hash
	RechargeShowPicPath        StrWithPrefix16
	MVPVoicePattern            StrWithPrefix16
	UISelectVoice              StrWithPrefix16
	UILevelupVoice             StrWithPrefix16
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
}

type ElfBreakMetaData struct {
	// Fields
	ElfID     int32
	BreakStep int32

	// Properties
	MaxLevel        int32
	Addr1           uint32 `json:"-"`
	NeedPlayerLevel int32

	// Objects
	BreakCostMaterialList []ElfBreakMetaData_ElfBreakMaterial
}

type ElfBreakMetaData_ElfBreakMaterial struct {
	// Fields
	MaterialID int32
	Number     int32
}

type ElfCaptainSkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	Addr1          uint32 `json:"-"`
	RestrictExtend int32

	// Objects
	Restricts []int32
}

type ElfCardMetaData struct {
	// Fields
	CardID int32

	// Properties
	ElfID              int32
	Rarity             int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	GachaDisplayRarity int32
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`

	// Objects
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
}

type ElfFragmentMetaData struct {
	// Fields
	FragmentID int32

	// Properties
	ElfID              int32
	Rarity             int32
	BaseType           int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	QuantityLimit      int32
	Addr5              uint32 `json:"-"`
	ReturnMaterialID   int32
	ReturnNum          int32
	ReturnMinStar      int32
	TagType            uint8
	GachaDisplayRarity int32
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`

	// Objects
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	LinkIDList                 []uint32
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
}

type ElfGalEventMetaData struct {
	// Fields
	GalEventID int32

	// Properties
	ElfID                  int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	AudioDelayTime         int32
	SelfCD                 int32
	Priority               int32
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`
	FaceAnimationDelayTime int32

	// Objects
	TriggerName       StrWithPrefix16
	Audio             StrWithPrefix16
	TriggerEffectList []ElfGalEventMetaData_ElfUIEffectPattern
	FaceAnimationKey  StrWithPrefix16
}

type ElfGalEventMetaData_ElfUIEffectPattern struct {
	// Fields
	BeginTime float32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	AttachPoint StrWithPrefix16
	EffectPath  StrWithPrefix16
}

type ElfMetaData struct {
	// Fields
	ElfID int32

	// Properties
	Index             int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	IsFlight          int32
	UnlockStar        int32
	Rarity            int32
	ElfCardID         int32
	ElfFragmentID     int32
	CarryPlayerLevel  int32
	Addr12            uint32 `json:"-"`
	Addr13            uint32 `json:"-"`
	RGBCGID           int32
	AlphaCGID         int32
	UltraSkillCD      int32
	UltraSkillSPCost  int32
	AttackSkillCD     float32
	Addr14            uint32 `json:"-"`
	Addr15            uint32 `json:"-"`
	Addr16            uint32 `json:"-"`
	Addr17            uint32 `json:"-"`
	CombatPointRarity float32
	Addr18            uint32 `json:"-"`
	Addr19            uint32 `json:"-"`
	Addr20            uint32 `json:"-"`
	Addr21            uint32 `json:"-"`

	// Objects
	FullName               Hash
	ElfRegistryKey         StrWithPrefix16
	ElfUIModelPath         StrWithPrefix16
	ElfUIModelPosition     ElfMetaData_VectorFloat3
	ElfUIModelRotation     ElfMetaData_VectorFloat3
	IconPath               StrWithPrefix16
	ElfChibiIcon           StrWithPrefix16
	PortraitCommonIcon     StrWithPrefix16
	StoryBGImg             StrWithPrefix16
	PresentBGImg           StrWithPrefix16
	AIName                 StrWithPrefix16
	UISoundbankName        StrWithPrefix16
	StoryDesc              Hash
	SkillTabBGImg          StrWithPrefix16
	TalentTabBGImg         StrWithPrefix16
	SelectAudio            StrWithPrefix16
	LevelUpAudio           StrWithPrefix16
	TagList                []ElfMetaData_Tag
	FeatureBrief           Hash
	TrialStageActivityList []int32
	CaptainSkillIDs        []int32
}

type ElfMetaData_MatCost struct {
	// Fields
	ID     int32
	Number int32
}

type ElfMetaData_VectorFloat3 struct {
	// Fields
	X float32
	Y float32
	Z float32
}

type ElfMetaData_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type ElfRarityMetaData struct {
	// Fields
	RarityID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	RarityImg StrWithPrefix16
}

type ElfSkillMetaData struct {
	// Fields
	ElfSkillID int32

	// Properties
	ElfID                   int32
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	Addr3                   uint32 `json:"-"`
	MaxLv                   int32
	SkillType               int32
	Addr4                   uint32 `json:"-"`
	IconType                int32
	UnlockStar              int32
	UIPointRow              int32
	UIPointColumn           int32
	Addr5                   uint32 `json:"-"`
	IconSpecial             int32
	AbilityParamBase_1      float32
	AbilityParamAdd_1       float32
	AbilityParamBase_2      float32
	AbilityParamAdd_2       float32
	AbilityParamBase_3      float32
	AbilityParamAdd_3       float32
	HasNoRestrictionAbility int32

	// Objects
	Name         Hash
	Info         Hash
	SkillTypeTag Hash
	IconPath     StrWithPrefix16
	TagList      []ElfSkillMetaData_Tag
}

type ElfSkillMetaData_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type ElfSkillTreeMetaData struct {
	// Fields
	ElfSkillID int32
	ElfSkillLv int32

	// Properties
	Addr1        uint32 `json:"-"`
	LevelUpStar  int32
	NeedElfLevel int32
	Addr2        uint32 `json:"-"`

	// Objects
	LevelUpPreSkill     []ElfSkillTreeMetaData_ElfPreSkill
	LevelUpMaterialList []ElfSkillTreeMetaData_ElfSkillLvUpMaterial
}

type ElfSkillTreeMetaData_ElfSkillLvUpMaterial struct {
	// Fields
	MaterialID int32
	Number     int32
}

type ElfSkillTreeMetaData_ElfPreSkill struct {
	// Fields
	Lv      int32
	SkillID int32
}

type ElfStarMetaData struct {
	// Fields
	ElfID int32
	Star  int32

	// Properties
	StarDisplay       int32
	StarDisplayMax    int32
	SubStarDisplay    int32
	SubStarDisplayMax int32
	UpgradeFragment   int32
	UpgradeElfLevel   int32
	SPBase            float32
	SPAdd             float32
	ATKBase           float32
	ATKAdd            float32
	CRTBase           float32
	CRTAdd            float32
	ElementType       int32
}

type ElfStoryMetaData struct {
	// Fields
	StoryId int32

	// Properties
	ElfId  int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	SelfCD float32

	// Objects
	StoryTitle  Hash
	StoryDetail Hash
	Audio       StrWithPrefix16
}

type EntryThemeItemDataMetaData struct {
	// Fields
	ThemeItemID int32

	// Properties
	UnlockType   uint8
	UnlockPartID int32
	Rarity       int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`

	// Objects
	ItemName  StrWithPrefix16
	ItemDesc  StrWithPrefix16
	IconPath  StrWithPrefix16
	ImagePath StrWithPrefix16
}

type EntryThemeSampleDataMetaData struct {
	// Fields
	SampleID int32

	// Properties
	SpaceShipConfigID int32
	Rarity            uint8
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`

	// Objects
	IconPath   StrWithPrefix16
	ImagePath  StrWithPrefix16
	SampleName Hash
	SampleDesc Hash
}

type EquipFragMetaData struct {
	// Fields
	FragID int32

	// Properties
	Cost      int32
	Equipment int32
}

type EquipmentSetMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Prop1ID         int32
	SpellEffectNum1 int32
	Prop1Param1     float32
	Prop1Param2     float32
	Prop1Param3     float32
	Prop1Param1Add  float32
	Prop1Param2Add  float32
	Prop1Param3Add  float32
	Prop2ID         int32
	SpellEffectNum2 int32
	Prop2Param1     float32
	Prop2Param2     float32
	Prop2Param3     float32
	Prop2Param1Add  float32
	Prop2Param2Add  float32
	Prop2Param3Add  float32
	Prop3ID         int32
	SpellEffectNum3 int32
	Prop3Param1     float32
	Prop3Param2     float32
	Prop3Param3     float32
	Prop3Param1Add  float32
	Prop3Param2Add  float32
	Prop3Param3Add  float32

	// Objects
	SetName Hash
	SetDesc Hash
}

type EquipSkillMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	SkillCD         int32
	SPCost          float32
	SPNeed          float32
	MaxChargesCount int32
	Addr4           uint32 `json:"-"`

	// Objects
	SkillName     Hash
	SkillDisplay  Hash
	SkillIconPath StrWithPrefix16
	TagList       []EquipSkillMetaData_Tag
}

type EquipSkillMetaData_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type FrameDataMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1       uint32 `json:"-"`
	Rarity      int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	IsShow      bool
	UIShowOrder int32
	Type        uint8

	// Objects
	Name      Hash
	Desc      Hash
	IconPath  StrWithPrefix16
	ImagePath StrWithPrefix16
}

type GrandKeyBuffMetaData struct {
	// Fields
	GrandKeySkill int32

	// Properties
	GrandKeyID          int32
	UnlockgrandKeyLevel int32
	BreachLevel         int32
	HP                  int32
	SP                  int32
	Attack              int32
	Defence             int32
	Critical            int32
	Type                int32
	IsSynchron          int32
	Param1              float32
	Param2              float32
	Param3              float32
	Side                int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`

	// Objects
	BuffNameText  Hash
	BuffDescStrId Hash
	BuffIcon      StrWithPrefix16
	BuffTypeIcon  StrWithPrefix16
}

type HalfBalanceModeAttrMetaData struct {
	// Fields
	AttrPoolId int32

	// Properties
	HP              int32
	SP              int32
	Atk             int32
	Dfs             int32
	CritChance      int32
	CritDamage      int32
	AtkSpeed        int32
	PhyDam_add      int32
	FireDam_add     int32
	IceDam_add      int32
	ThunderDam_add  int32
	PhyDam_add2     int32
	FireDam_add2    int32
	IceDam_add2     int32
	ThunderDam_add2 int32
	AllDam_add      int32
	AllDam_add2     int32
}

type HalfBalanceModeDataMetaData struct {
	// Fields
	StageId int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ExceptFunctionId []int32
	AttrPoolId       []HalfBalanceModeDataMetaData_AvatarPoolKV
	DescInfo         Hash
}

type HalfBalanceModeDataMetaData_AvatarPoolKV struct {
	// Fields
	AvatarId int32
	PoolId   int32
}

type ItemMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity                  uint8
	MaxRarity               uint8
	Cost                    uint8
	MaxLv                   uint8
	SellPriceBase_1         int32
	SellPriceAdd            uint16
	ServantExpProvide       uint16
	GearExpProvideBase      uint16
	GearExpPorvideAdd       uint16
	Addr1                   uint32 `json:"-"`
	UseID                   int32
	BaseType                uint8
	Addr2                   uint32 `json:"-"`
	Addr3                   uint32 `json:"-"`
	Addr4                   uint32 `json:"-"`
	Addr5                   uint32 `json:"-"`
	CharacterExpProvide     uint16
	Addr6                   uint32 `json:"-"`
	Addr7                   uint32 `json:"-"`
	Addr8                   uint32 `json:"-"`
	QuantityLimit           int32
	SortID                  uint8
	AffixTrainType          uint8
	AffixRandomValueIncress uint8
	AffixTitleExp           uint16
	QuickBuyType            uint8
	ShopType                uint8
	IdShopGoods             int32
	QuickBuyConfirm         bool
	HideInInventory         bool
	HideNumInTips           bool
	SellPriceID_1           uint8
	CostVitality            uint16
	EnableQuickSell         bool
	AdvSellBonusNum         uint16
	TagType                 uint8
	Addr9                   uint32 `json:"-"`
	Addr10                  uint32 `json:"-"`
	AlwaysShowPopUp         bool

	// Objects
	ItemType                   StrWithPrefix16
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	LinkIDList                 []uint32
	ShopUseList                []StrWithPrefix16
	DisplayBGDescription       Hash
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
}

type MaterialDeleteData struct {
	// Fields
	Id int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	ReturnMaterialID int32
	ReturnNum        float32
	RoundType        int32

	// Objects
	CallbackTime RemoteTime
	ReturnList   StrWithPrefix16
}

type MechMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Type              int32
	SubType           int32
	Addr4             uint32 `json:"-"`
	CanRide           int32
	Rarity            int32
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	HP                int32
	Attack            int32
	Defence           int32
	Critical          int32
	SP                int32
	HP_Tower          int32
	Attack_Tower      int32
	Defence_Tower     int32
	Critical_Tower    int32
	EffectParam       int32
	DisjoinScoinCost  int32
	Addr12            uint32 `json:"-"`
	Addr13            uint32 `json:"-"`
	Addr14            uint32 `json:"-"`
	ModeRenderYOffset float32
	ModeRenderScale   float32

	// Objects
	MonsterName           StrWithPrefix16
	MonsterTypeName       StrWithPrefix16
	Name_TextMap          Hash
	RandomType_TextMap    Hash
	IconPath              StrWithPrefix16
	TowerStandby_IconPath StrWithPrefix16
	HPGrade               StrWithPrefix16
	AttackGrade           StrWithPrefix16
	DefenceGrade          StrWithPrefix16
	CriticalGrade         StrWithPrefix16
	SPGrade               StrWithPrefix16
	DisjoinAddMaterial    []MechMetaData_DisjoinOutputItem
	Story                 Hash
	BestWeather           []int32
}

type MechMetaData_DisjoinOutputItem struct {
	// Fields
	ID  int32
	Num int32
}

type MedalMetaData struct {
	// Fields
	MedalID int32

	// Properties
	Rarity          int32
	TimeType        int32
	During          int32
	IndexID         int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	HP              int32
	SP              int32
	Attack          int32
	Defence         int32
	Critical        int32
	Prop1ID         int32
	Prop1param1     float32
	Prop1param2     float32
	Prop1param3     float32
	Prop2ID         int32
	Prop2param1     float32
	Prop2param2     float32
	Prop2param3     float32
	Prop3ID         int32
	Prop3param1     float32
	Prop3param2     float32
	Prop3param3     float32
	Show            int32
	Addr7           uint32 `json:"-"`
	ShowMission     int32
	Addr8           uint32 `json:"-"`
	IsCollaboration bool

	// Objects
	MedalTitle       Hash
	MedalDescription Hash
	MedalGet         Hash
	IconPath         StrWithPrefix16
	ImagePath        StrWithPrefix16
	SmallPath        StrWithPrefix16
	EndTime          StrWithPrefix16
	ShowTime         RemoteTime
}

type PhonePendantMetaData struct {
	// Fields
	PendantId int32

	// Properties
	Rarity     int32
	TimeType   int32
	TimeDuring int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`
	Addr7      uint32 `json:"-"`

	// Objects
	PendantModPath        StrWithPrefix16
	PendantIconPath       StrWithPrefix16
	PendantImagePath      StrWithPrefix16
	PendantNameText       Hash
	PendantDescText       Hash
	PendantSourceDescText Hash
	PendantUnlockDescText Hash
}

type ScDLCAvatarMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	IsFake      bool
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	HpBase      int32
	SpBase      int32
	AtkBase     int32
	DfsBase     int32
	TalentPoint int32
	NPCID       int32
	Addr6       uint32 `json:"-"`
	Addr7       uint32 `json:"-"`

	// Objects
	AvatarRegistryKey StrWithPrefix16
	FullName          Hash
	Desc              Hash
	IconPath          StrWithPrefix16
	ImagePath         StrWithPrefix16
	AvatarHead        StrWithPrefix16
	UIOffset          []float32
}

type ScDLCCombineTalentMetaData struct {
	// Fields
	ID int32

	// Properties
	CombineTalentID int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`

	// Objects
	CombineTalentSource []int32
	SubIcon             StrWithPrefix16
	SubTalentTitle      Hash
	SubTalentDesc       Hash
}

type ScDLCTalentDataMetaData struct {
	// Fields
	TalentID int32

	// Properties
	AvatarID     int32
	TalentType   int32
	ParentTalent int32
	UniqueTag    int32
	ForceLock    bool
	OWStoryID    int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	IsSupport    bool
	Addr6        uint32 `json:"-"`

	// Objects
	UIPointName    StrWithPrefix16
	UILineName     []StrWithPrefix16
	TalentName     Hash
	IconPath       StrWithPrefix16
	CGID           []int32
	SupportDerived []int32
}

type ScDLCTalentLevelMetaData struct {
	// Fields
	TalentID    int32
	TalentLevel int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	FeverLevelLimit    int32
	Addr5              uint32 `json:"-"`
	AbilityParamBase_1 float32
	AbilityParamBase_2 float32
	AbilityParamBase_3 float32
	IsCoopTalent       bool
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`

	// Objects
	TalentDesc         Hash
	NextLevelDesc      Hash
	SubTalentLevelDesc []StrWithPrefix16
	PreTalent          []ScDLCTalentLevelMetaData_LevelUpTalentRequire
	LevelUpCost        []ScDLCTalentLevelMetaData_LevelUpMatRequire
	CoopAbilityName    StrWithPrefix16
	CoopAbilityParam   []StrWithPrefix16
}

type ScDLCTalentLevelMetaData_LevelUpMatRequire struct {
	// Fields
	MatID int32
	Num   int32
}

type ScDLCTalentLevelMetaData_LevelUpTalentRequire struct {
	// Fields
	Level    int32
	TalentID int32
}

type SpaceShipConfigMetaData struct {
	// Fields
	SpaceShipID int32

	// Properties
	Addr1        uint32 `json:"-"`
	IsSkyBoxShow bool

	// Objects
	ResPath StrWithPrefix16
}

type SpecialAvatarDataMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	SpecialAvatarPrefab StrWithPrefix16
}

type StageRestrictExtendMetaData struct {
	// Fields
	RestrictID uint32

	// Properties
	AvatartNumMax    uint8
	RequestAvatarNum uint8
	ForceSelect      bool
	RequestType1     uint8
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	RequestType2     uint8
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	Addr7            uint32 `json:"-"`
	Addr8            uint32 `json:"-"`
	RequestType3     uint8
	Addr9            uint32 `json:"-"`
	Addr10           uint32 `json:"-"`
	Addr11           uint32 `json:"-"`
	Addr12           uint32 `json:"-"`

	// Objects
	RequestParam1   []int32
	RequsetTitle1   Hash
	RequestDesc1    Hash
	RequestPosName1 Hash
	RequestParam2   []int32
	RequestTitle_2  Hash
	RequestDesc2    Hash
	RequestPosName2 Hash
	RequestParam3   []int32
	RequestTitle3   Hash
	RequestDesc3    Hash
	RequestPosName3 Hash
}

type StageRestrictMetaData struct {
	// Fields
	RestrictID int32

	// Properties
	Addr1          uint32 `json:"-"`
	RequestType1   int32
	Argument1Type1 int32
	Argument2Type1 int32
	Addr2          uint32 `json:"-"`
	RequestType2   int32
	Argument1Type2 int32
	Argument2Type2 int32
	Addr3          uint32 `json:"-"`
	RequestType3   int32
	Argument1Type3 int32
	Argument2Type3 int32
	Addr4          uint32 `json:"-"`

	// Objects
	DisplayNameID     Hash
	ArgumentListType1 []int32
	ArgumentListType2 []int32
	ArgumentListType3 []int32
}

type StigmataMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity             uint8
	MaxRarity          uint8
	SubRarity          uint8
	SubMaxRarity       uint8
	Cost               uint8
	PowerType          uint8
	MaxLv              uint8
	ExpType            uint8
	SellPriceBase      float32
	SellPriceAdd       float32
	GearExpProvideBase float32
	GearExpPorvideAdd  float32
	BaseType           uint8
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	DisplayNumber      int32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	HPBase             float32
	HPAdd              float32
	SPBase             float32
	SPAdd              float32
	AttackBase         float32
	AttackAdd          float32
	DefenceBase        float32
	DefenceAdd         float32
	CriticalBase       float32
	CriticalAdd        float32
	DurabilityMax      float32
	Addr6              uint32 `json:"-"`
	EvoID              int32
	Prop1ID            int32
	Prop1Param1        float32
	Prop1Param2        float32
	Prop1Param3        float32
	Prop1Param1Add     float32
	Prop1Param2Add     float32
	Prop1Param3Add     float32
	Prop2ID            int32
	Prop2Param1        float32
	Prop2Param2        float32
	Prop2Param3        float32
	Prop2Param1Add     float32
	Prop2Param2Add     float32
	Prop2Param3Add     float32
	Prop3ID            int32
	Prop3Param1        float32
	Prop3Param2        float32
	Prop3Param3        float32
	Prop3Param1Add     float32
	Prop3Param2Add     float32
	Prop3Param3Add     float32
	Protect            bool
	SetID              int32
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	OffsetX            float32
	OffsetY            float32
	Scale              float32
	AffixTreeId        uint8
	CanRefine          bool
	RecycleID          uint8
	DisjoinScoinCost   int32
	Addr9              uint32 `json:"-"`
	Addr10             uint32 `json:"-"`
	Quality            uint8
	StigmataMainID     int32
	Addr11             uint32 `json:"-"`
	SellPriceID        uint8
	Transcendent       bool
	Target             int32
	IsSecurityProtect  bool
	Addr12             uint32 `json:"-"`
	Addr13             uint32 `json:"-"`
	Addr14             uint32 `json:"-"`
	CollaborationSetID int32

	// Objects
	LabelPath                  StrWithPrefix16
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	EvoMaterial                []StigmataMetaData_Material
	SmallIcon                  StrWithPrefix16
	TattooPath                 StrWithPrefix16
	DisjoinAddMaterial         []StigmataMetaData_Material
	LinkIDList                 []uint32
	ShortName                  Hash
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
	StigmataFilterList         []int32
}

type StigmataMetaData_Material struct {
	// Fields
	ID  int32
	Num int32
}

type TagListDataMetaData struct {
	// Fields
	TagID uint16

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Type         uint8
	ShowInScreen bool
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`

	// Objects
	TagName             Hash
	TagDec              Hash
	Icon                StrWithPrefix16
	DisplayWeaponList   []int32
	DisplayStigmataList []int32
}

type ThemeBGMConfigMetaData struct {
	// Fields
	ThemeBGMID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	StateName       StrWithPrefix16
	SwitchName      StrWithPrefix16
	SwitchGroupName StrWithPrefix16
	ShowName        Hash
}

type VirtualResourcesMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1         uint32 `json:"-"`
	DisplayRarity int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	ProtoID       int32
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`

	// Objects
	DisplayTitle         Hash
	DisplayDescription   Hash
	IconPath             StrWithPrefix16
	ImagePath            StrWithPrefix16
	LinkIDList           []uint32
	DisplayBGDescription Hash
}

type WeaponMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity                uint8
	MaxRarity             uint8
	SubRarity             uint8
	SubMaxRarity          uint8
	Cost                  uint8
	PowerType             uint8
	MaxLv                 uint8
	ExpType               uint8
	SellPriceBase         float32
	SellPriceAdd          float32
	GearExpProvideBase    float32
	GearExpPorvideAdd     float32
	BaseType              uint8
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	HPBase                float32
	HPAdd                 float32
	SPBase                float32
	SPAdd                 float32
	AttackBase            float32
	AttackAdd             float32
	DefenceBase           float32
	DefenceAdd            float32
	CriticalBase          float32
	CriticalAdd           float32
	ResistanceBase        float32
	ResistanceAdd         float32
	Addr6                 uint32 `json:"-"`
	EvoPlayerLevel        int32
	EvoID                 int32
	Prop1ID               int32
	Prop1Param1           float32
	Prop1Param2           float32
	Prop1Param3           float32
	Prop1Param1Add        float32
	Prop1Param2Add        float32
	Prop1Param3Add        float32
	Prop2ID               int32
	Prop2Param1           float32
	Prop2Param2           float32
	Prop2Param3           float32
	Prop2Param1Add        float32
	Prop2Param2Add        float32
	Prop2Param3Add        float32
	Prop3ID               int32
	Prop3Param1           float32
	Prop3Param2           float32
	Prop3Param3           float32
	Prop3Param1Add        float32
	Prop3Param2Add        float32
	Prop3Param3Add        float32
	Protect               bool
	ExDisjoinCurrencyCost int32
	Addr7                 uint32 `json:"-"`
	DisjoinScoinCost      int32
	Addr8                 uint32 `json:"-"`
	WeaponMainID          int32
	Addr9                 uint32 `json:"-"`
	WeaponQuality         uint8
	SellPriceID           uint8
	Transcendent          bool
	Target                int32
	Addr10                uint32 `json:"-"`
	Addr11                uint32 `json:"-"`
	Addr12                uint32 `json:"-"`
	Addr13                uint32 `json:"-"`
	CollaborationWeaponID int32
	AvatarCustomDisplayID int32

	// Objects
	BodyMod                    StrWithPrefix16
	DisplayTitle               Hash
	DisplayDescription         Hash
	IconPath                   StrWithPrefix16
	ImagePath                  StrWithPrefix16
	EvoMaterial                []WeaponMetaData_Material
	ExDisjoinAddMaterial       []WeaponMetaData_Material
	DisjoinAddMaterial         []WeaponMetaData_Material
	LinkIDList                 []uint32
	GachaMainDropDisplayConfig []float32
	GachaGiftDropDisplayConfig []float32
	PreloadEffectFolderPath    StrWithPrefix16
	WeaponFilterList           []int32
}

type WeaponMetaData_Material struct {
	// Fields
	ID  int32
	Num int32
}

type AbilitySpecialMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	AbilityName StrWithPrefix16
	ParamList   []AbilitySpecialMetaData_ParamEntry
}

type AbilitySpecialMetaData_ParamEntry struct {
	// Fields
	ParamValue float32
	Addr1      uint32 `json:"-"`

	// Objects
	Paramkey StrWithPrefix16
}

type AccountBuffEffect struct {
	// Fields
	EffectID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	LimitType int32
	LimitNum  int32

	// Objects
	EffectDisplayIcon StrWithPrefix16
	EffectNameText    Hash
	EffectDisplayText Hash
}

type AccumulatorUIConfigMetaData struct {
	// Fields
	ID uint8

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`
	Addr11 uint32 `json:"-"`
	Addr12 uint32 `json:"-"`
	Addr13 uint32 `json:"-"`
	Addr14 uint32 `json:"-"`
	Addr15 uint32 `json:"-"`
	Addr16 uint32 `json:"-"`
	Addr17 uint32 `json:"-"`

	// Objects
	AccDebuff           StrWithPrefix16
	IconPath            StrWithPrefix16
	IconColor           StrWithPrefix16
	CycleColor          StrWithPrefix16
	RootColor           StrWithPrefix16
	ResWaveColor        StrWithPrefix16
	StarColor           StrWithPrefix16
	GlowColor           StrWithPrefix16
	MaskSliderFillColor StrWithPrefix16
	LayerTextColor      StrWithPrefix16
	MaxLayerTextColor   StrWithPrefix16
	MaskColor           StrWithPrefix16
	BaseColor           StrWithPrefix16
	FillColor           StrWithPrefix16
	FlashColor          StrWithPrefix16
	TipsName            Hash
	TipsDesc            Hash
}

type AchievementTagMetaData struct {
	// Fields
	Tag int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Type  int32
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	TagName    Hash
	TagColor   StrWithPrefix16
	LineColor  StrWithPrefix16
	TypeIcon   StrWithPrefix16
	TypeName   StrWithPrefix16
	TypeNameEn StrWithPrefix16
}

type ActChallengeMetaData struct {
	// Fields
	ActId      int32
	Difficulty int32

	// Properties
	ChallengeNum1 int32
	RewardID1     int32
	ChallengeNum2 int32
	RewardID2     int32
	ChallengeNum3 int32
	RewardID3     int32
	IsShow        int32
	FirstStageID  int32
	LastStageID   int32
}

type ActiveCollectionItemMetaData struct {
	// Fields
	ID int32

	// Properties
	ButtonType uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`
	MinLevel   int32
	MaxLevel   int32

	// Objects
	ButtonTextMap   Hash
	ButtonImagePath StrWithPrefix16
	Link            StrWithPrefix16
	RewardIDList    []int32
	StartTime       RemoteTime
	EndTime         RemoteTime
}

type ActiveCollectionScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	Addr7    uint32 `json:"-"`
	MinLevel int32
	MaxLevel int32

	// Objects
	BGPath             StrWithPrefix16
	TitleTextMap       StrWithPrefix16
	TitleIconImagePath StrWithPrefix16
	CurrencyIDList     []int32
	ActiveItemList     []int32
	StartTime          RemoteTime
	EndTime            RemoteTime
}

type ActivityBbqMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	CostItemID           int32
	CostItemNum          int32
	GratuityTriggerCount int32
}

type ActivityBingoMetaData struct {
	// Fields
	BingoID int32

	// Properties
	PreBingo        int32
	BingoSize       int32
	BingoNum        int32
	ResetTimes      int32
	ResetItemID     int32
	ResetItemNum    int32
	CostItemID      int32
	CostItemNum     int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	BingoColorTheme int32
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`
	Addr10          uint32 `json:"-"`
	Addr11          uint32 `json:"-"`
	Addr12          uint32 `json:"-"`
	Addr13          uint32 `json:"-"`
	Addr14          uint32 `json:"-"`
	Addr15          uint32 `json:"-"`
	BingoLinkType1  int32
	Addr16          uint32 `json:"-"`
	Addr17          uint32 `json:"-"`
	Addr18          uint32 `json:"-"`
	Addr19          uint32 `json:"-"`
	BingoLinkType2  int32
	Addr20          uint32 `json:"-"`
	Addr21          uint32 `json:"-"`
	Addr22          uint32 `json:"-"`
	Addr23          uint32 `json:"-"`
	Addr24          uint32 `json:"-"`

	// Objects
	UnlockTime           StrWithPrefix16
	HighlightTimeUp      StrWithPrefix16
	HighlightTimeDown    StrWithPrefix16
	TagText              Hash
	TagImage             StrWithPrefix16
	BingoCardFrontCircle StrWithPrefix16
	BingoCardFrontCross  StrWithPrefix16
	BingoCardBack        StrWithPrefix16
	TitleText            Hash
	TipText              Hash
	BingoPromotText      StrWithPrefix16
	BingoPromotImage     StrWithPrefix16
	BingoPromotAvatar    StrWithPrefix16
	BingoLinkText1       Hash
	BingoLinkImage1      StrWithPrefix16
	BingoLinkParams1     []int32
	BingoLinkParamStr1   StrWithPrefix16
	BingoLinkText2       Hash
	BingoLinkImage2      StrWithPrefix16
	BingoLinkParams2     []int32
	BingoLinkParamStr2   StrWithPrefix16
	BingoRulesImage      StrWithPrefix16
	BingoWebImage        StrWithPrefix16
	BingoWebLink         StrWithPrefix16
}

type ActivityBossRushMetaData struct {
	// Fields
	GeneralActivityID int32

	// Properties
	Addr1            uint32 `json:"-"`
	AvatarInitial    int32
	AvatarUnlockTime int32
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`

	// Objects
	AvatarTrial          []int32
	RushDailyMission     []int32
	RushChallengeMission []int32
	BossRushImg          []StrWithPrefix16
}

type ActivityCardDataMetaData struct {
	// Fields
	ActiveID int32

	// Properties
	Addr1          uint32 `json:"-"`
	TriggerMission int32
	DisplayMission int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Xcoordinate    int32
	Ycoordinate    int32

	// Objects
	MissionIDList      []int32
	MissionListName    Hash
	MissionListDisplay Hash
	MissionListReward  []int32
	ActivePic          StrWithPrefix16
	ActivePic1         StrWithPrefix16
	TextmapRuleTitle   Hash
	TextmapRule        Hash
}

type ActivityChallengeMetaData struct {
	// Fields
	ActivityId int32

	// Properties
	ChallengeNum1 int32
	RewardID1     int32
	ChallengeNum2 int32
	ChallengeNum3 int32
}

type ActivityChargeLevelMetaData struct {
	// Fields
	AcitivityID int32
	ChargeID    uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ChargeAttrList []int32
}

type ActivityDropLimitData struct {
	// Fields
	Id uint16

	// Properties
	MaterialId int32
	Amount     uint16
	Type       uint8
	Addr1      uint32 `json:"-"`

	// Objects
	TypeParam []int32
}

type ActivityDropLimitScheduleMetaData struct {
	// Fields
	Id uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	DataIdList []int32
}

type ActivityLoginBonusMetaData struct {
	// Fields
	ID   int32
	Days int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	ImagePath StrWithPrefix16
	BonusDesc Hash
}

type ActivityLoginMetaData struct {
	// Fields
	ID int32

	// Properties
	Type         int32
	Addr1        uint32 `json:"-"`
	ActivityType int32
	Addr2        uint32 `json:"-"`
	PlayerLevel  int32
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	Addr7        uint32 `json:"-"`
	Addr8        uint32 `json:"-"`
	Addr9        uint32 `json:"-"`
	Addr10       uint32 `json:"-"`
	Addr11       uint32 `json:"-"`
	AvatarPosX   int32
	AvatarPosY   int32
	Addr12       uint32 `json:"-"`
	CanSharePost bool

	// Objects
	TabTextMapID    Hash
	EndTime         StrWithPrefix16
	RewardList      []int32
	KeyRewardList   []int32
	LoginTitleText  Hash
	LoginRewardText Hash
	ImgAvatar       StrWithPrefix16
	ImgTheme        StrWithPrefix16
	ImgTitle        StrWithPrefix16
	ImgBG           StrWithPrefix16
	ImgBGLight      StrWithPrefix16
	WebLink         StrWithPrefix16
}

type ActivityMosaicMetaData struct {
	// Fields
	MosaicID int32

	// Properties
	Addr1        uint32 `json:"-"`
	MosaicNum    int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	MosaicReward int32
	Xcoordinate  int32
	Ycoordinate  int32

	// Objects
	MissionIDList []int32
	MosaicPic     StrWithPrefix16
	TextmapRule   StrWithPrefix16
}

type ActivityPanelMetaData struct {
	// Fields
	Id uint16

	// Properties
	GroupID        int32
	Type           uint8
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	CompleteDelete bool
	CurrencyHide   bool

	// Objects
	TypeParam      []int32
	Description    Hash
	ImagePath      StrWithPrefix16
	TextImagePath  StrWithPrefix16
	SuperTitleName StrWithPrefix16
	TimeTitleName  Hash
	BGPath         StrWithPrefix16
	PanelColor     StrWithPrefix16
	TabLabelIcon   StrWithPrefix16
}

type ActivityPanelScoreData struct {
	// Fields
	PanelID   int32
	MissionID int32

	// Properties
	Score int32
}

type ActivityPanelScoreRewardData struct {
	// Fields
	PanelID  int32
	Progross int32

	// Properties
	TotalScore int32
	RewardID   int32
	StageID    int32
	LinkType   int32
	Addr1      uint32 `json:"-"`

	// Objects
	LinkParams []int32
}

type ActivityPictureMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	MaxStep      int32
	MaxScore     int32
	Addr1        uint32 `json:"-"`
	JumpActivity int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`

	// Objects
	Explain_Textmap Hash
	BackgroundPic   StrWithPrefix16
	ContentPrefab   StrWithPrefix16
}

type ActivityQuestionnaireMetaData struct {
	// Fields
	ID int32

	// Properties
	LinkType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	Website    StrWithPrefix16
	ImgTheme   StrWithPrefix16
	ImgBG      StrWithPrefix16
	ImgBGLight StrWithPrefix16
}

type ActivityRandomEffectChargeMetaData struct {
	// Fields
	ChargeID uint8

	// Properties
	ChargeValue uint8
}

type ActivityRandomEffectMetaData struct {
	// Fields
	ActivityBuffID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	BuffChargeList  []int32
	DisplayBuffList []int32
}

type ActivityRewardShowMetaData struct {
	// Fields
	RewardShowID uint8

	// Properties
	ShopType uint32
	ShopID   uint32
	Addr1    uint32 `json:"-"`

	// Objects
	ShowGoodsIDList []int32
}

type ActivityScheduleDisplayData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	ItemDisplay []ActivityScheduleDisplayData_ItemDisplay
	BeginTime   StrWithPrefix16
	EndTime     StrWithPrefix16
	ImagePath   StrWithPrefix16
	Text        Hash
	Link        StrWithPrefix16
}

type ActivityScheduleDisplayData_ItemDisplay struct {
	// Fields
	ItemID  int32
	ItemNum int32
}

type ActivitySchedulePageMetaData struct {
	// Fields
	ID int16

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	ShowReward int32
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`

	// Objects
	ActivityTitle Hash
	TitleImgPath  StrWithPrefix16
	ActivityDesc  Hash
	BgImgPath     StrWithPrefix16
	EndTime       StrWithPrefix16
	DescBgImgPath StrWithPrefix16
}

type ActivityStageRankMetaData struct {
	// Fields
	StageID int32

	// Properties
	IsDisplay   bool
	RankGroupID uint32
}

type ActivityTowerMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	ActivityType int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	UnlockStage  int32
	MaterialID   int32
	LinkChapter  int32
	Addr3        uint32 `json:"-"`
	DailyCount   int32
	ClearReward  int32
	Addr4        uint32 `json:"-"`

	// Objects
	EnterImg            StrWithPrefix16
	EnterImgColor       StrWithPrefix16
	StageConfigList     []int32
	SpecialMissionPanel StrWithPrefix16
}

type ActivityTypeManagementMetaData struct {
	// Fields
	ActivityID uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ManagementTypeList []int32
}

type ActMetaData struct {
	// Fields
	ActId int32

	// Properties
	ChapterId    int32
	NumInChapter int32
	Addr1        uint32 `json:"-"`
	ActType      int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	Addr7        uint32 `json:"-"`
	Addr8        uint32 `json:"-"`
	Addr9        uint32 `json:"-"`

	// Objects
	ActName             Hash
	LevelPannelPath     StrWithPrefix16
	HardLevelPannelPath StrWithPrefix16
	HellLevelPannelPath StrWithPrefix16
	SmallImgPath        StrWithPrefix16
	BgImgPath           StrWithPrefix16
	BgPrefabPath        StrWithPrefix16
	OpenTime            StrWithPrefix16
	EventGroup          []ActMetaData_ActEventData
}

type ActMetaData_ActEventData struct {
	// Fields
	EventID    int32
	PreLevelID int32
}

type ActRitaRankingMetaData struct {
	// Fields
	LetterRankID int32
	RankLevel    int32

	// Properties
	Rankscore int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	ImagePath StrWithPrefix16
	RankName  Hash
}

type AddParamGroupMetaData struct {
	// Fields
	SkillID int32
	GroupID uint8

	// Properties
	LevelMin uint8
	LevelMax uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	BasicList  []float32
	AddList    []float32
	Desc       Hash
	DetailDesc Hash
}

type AdventureAvatarMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	GrowupTestID int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`

	// Objects
	TechnicTestID     []int32
	Mastery           []int32
	AvatarQuestSucess []int32
}

type AdventureAvatarMissionListMetaData struct {
	// Fields
	AdventureRecordID int32

	// Properties
	MissionId   int32
	AchieveType int32
	AvatarID    int32
	BadgeReward int32
}

type AdventureAvatarSkillLevelMetaData struct {
	// Fields
	ID int32

	// Properties
	AvatarSkillID int32
	Level         int32
	Par           float32
}

type AdventureAvatarSkillMetaData struct {
	// Fields
	AvatarSkillID int32

	// Properties
	Type  int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	SkillName             Hash
	SkillDesc             Hash
	DisplayAdventureQuest Hash
	Icon                  StrWithPrefix16
}

type AdventureAvatarUnlockMetaData struct {
	// Fields
	ID int32

	// Properties
	AvatarID int32
	Unlock   int32
	Skill    int32
}

type AdventureDecorateMetaData struct {
	// Fields
	ID uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	CollectionTypeList []uint8
}

type AdventureElfMetaData struct {
	// Fields
	ElfID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ElfQuestSucess []int32
}

type AdventureElfUnlockMetaData struct {
	// Fields
	ID int32

	// Properties
	ElfID  int32
	Unlock int32
	Skill  int32
}

type AdventureLevelDataMetaData struct {
	// Fields
	AdventureLevel int32

	// Properties
	BadgeNeed     int32
	ExaminationID int32
	Reward        int32
	HP            int32
	SP            int32
	Attack        int32
	Defence       int32
	Critical      int32
}

type AdventureQuestConfigMetaData struct {
	// Fields
	QuestConfigId int32

	// Properties
	DailyTimes            int32
	AdventureMaxGrain     int32
	AdventureGrainRecover int32
	QuestRefreshTimes     int32
}

type AdventureQuestMetaData struct {
	// Fields
	QuestId int32

	// Properties
	Rarity           int32
	QuestRare        uint8
	GrainCost        int32
	TimeCost         int32
	Addr1            uint32 `json:"-"`
	BonusDrop        int32
	QuestDropShow    int32
	BonusDropShow    int32
	FailReward       int32
	MinMember        int32
	MaxMember        int32
	Addr2            uint32 `json:"-"`
	TraitSuccessMax  int32
	AvatarSuccessMax int32

	// Objects
	NeedTraitList []AdventureQuestMetaData_ADQuestSkill
	QuestNameText Hash
}

type AdventureQuestMetaData_ADQuestSkill struct {
	// Fields
	SkillID       int32
	SkillStrength int32
}

type AdventureQuestPoolMetaData struct {
	// Fields
	QuestPoolId int32

	// Properties
	MapPosition int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	QuestPoolBgPath   StrWithPrefix16
	QuestPoolIconPath StrWithPrefix16
	QuestPoolDescText Hash
}

type AdventureWelfareMetaData struct {
	// Fields
	WelfareID int32

	// Properties
	AdventureLevel int32
	Type           int32
	Par1           int32
	Par2           int32
	Par3           int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`

	// Objects
	Title Hash
	Dec   Hash
	Icon  StrWithPrefix16
}

type AffixEffectMetaData struct {
	// Fields
	AffixEffectID int32

	// Properties
	Attributetype int32
	Addr1         uint32 `json:"-"`

	// Objects
	RelatedAffixTitleEffectTemplateList []int32
}

type AffixRecycleMetaData struct {
	// Fields
	RecycleId int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MaterialList []AffixRecycleMetaData_IntIntPair
}

type AffixRecycleMetaData_IntIntPair struct {
	// Fields
	Item1 int32
	Item2 int32
}

type AffixTitleEffectMetaData struct {
	// Fields
	AffixTitleEffectID int32

	// Properties
	AffixTitleEffectTemplateID int32
	EffectInitValue            float32
	EffectValuePerSubLevel     float32
}

type AffixTitleEffectTemplateMetaData struct {
	// Fields
	EffectName int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	EffectContentTextMap Hash
}

type AffixTitleLevelMetaData struct {
	// Fields
	QualityId int32

	// Properties
	BreakExpNeed    int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	InheritNeedCoin int32
	Addr3           uint32 `json:"-"`
	ResetNeedCoin   int32
	Addr4           uint32 `json:"-"`

	// Objects
	BreakMaterialNeed    []AffixTitleLevelMetaData_IntIntPair
	DisjointGetMaterials []AffixTitleLevelMetaData_IntIntPair
	InheritItemNeed      []AffixTitleLevelMetaData_IntIntPair
	ResetItemNeed        []AffixTitleLevelMetaData_IntIntPair
}

type AffixTitleLevelMetaData_IntIntPair struct {
	// Fields
	Item1 int32
	Item2 int32
}

type AffixTreeNodeMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	RefineCost_0to1  []AffixTreeNodeMetaData_MaterialNeed
	RefineCost_1     []AffixTreeNodeMetaData_MaterialNeed
	RefineCost_2     []AffixTreeNodeMetaData_MaterialNeed
	RefineCost_1to2  []AffixTreeNodeMetaData_MaterialNeed
	RefineCost_core  []AffixTreeNodeMetaData_MaterialNeed
	RefineCost_lock1 []AffixTreeNodeMetaData_MaterialNeed
}

type AffixTreeNodeMetaData_MaterialNeed struct {
	// Fields
	MaterialID  int32
	MaterialNum int32
}

type AffixTypeConfigMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TypeName Hash
	Icon     StrWithPrefix16
}

type AiCyberActivityScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	MapID            int32
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	TicketID         int32
	CurrencyID       int32
	MaxTicketNum     int32
	RepairMaterialID int32
	RewardScheduleID int32
	PerformID        int32

	// Objects
	BeginTime              RemoteTime
	EndTime                RemoteTime
	AreaIDList             []int32
	MainMissionIDList      []int32
	DailyMissionIDList     []int32
	ChallengeMissionIDList []int32
}

type AiCyberAreaMetaData struct {
	// Fields
	AreaID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	MainStageGroupID      int32
	DailyStageGroupID     int32
	ChallengeStageGroupID int32
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	PuzzleReward          int32

	// Objects
	AreaName            Hash
	UnlockConditionList StrWithPrefix16
	QAvatarID           []int32
	QAvatarSkillID      []int32
	PuzzleEventIDList   []int32
}

type AiCyberDailyStageDropLimitMetaData struct {
	// Fields
	ActivityScheduleID uint32
	Priority           uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	UnlockConditionList   StrWithPrefix16
	MaterialDropLimitList []AiCyberDailyStageDropLimitMetaData_MaterialDropLimit
}

type AiCyberDailyStageDropLimitMetaData_MaterialDropLimit struct {
	// Fields
	MaterialID int32
	Number     int32
}

type AiCyberDailyStageScoreDropMetaData struct {
	// Fields
	StageGroupID uint32
	MinScore     int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	DropMaterialList []AiCyberDailyStageScoreDropMetaData_DropMaterial
}

type AiCyberDailyStageScoreDropMetaData_DropMaterial struct {
	// Fields
	MaterialID int32
	Number     int32
}

type AiCyberGalTouchMetaData struct {
	// Fields
	GalEventID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Dialogue    Hash
	StartMotion StrWithPrefix16
}

type AiCyberHyperionMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	SpaceshipPath       StrWithPrefix16
	AvatarModelPath     StrWithPrefix16
	UnlockConditionList StrWithPrefix16
	GalEvents           []int32
}

type AiCyberMainUIRepairMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	ID            int32
	ProgressLimit int32

	// Objects
	UIPrefabPath int32
}

type AiCyberProgressMetaData struct {
	// Fields
	Progress int32

	// Properties
	SlotType              uint8
	SlotTypeParam         int32
	CostRepairMaterialNum int32
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	IsCritical            bool
	IsShow                bool

	// Objects
	SlotName            Hash
	SlotDec             Hash
	SlotIcon            StrWithPrefix16
	UnlockConditionList StrWithPrefix16
}

type AiCyberPuzzleEventMetaData struct {
	// Fields
	PuzzleEventID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	PuzzleSpawnPoint int32

	// Objects
	PuzzleDes                 Hash
	RewardUnlockConditionList StrWithPrefix16
}

type AiCyberQavatarMetaData struct {
	// Fields
	QavatarID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`
	Addr8 uint32 `json:"-"`

	// Objects
	QAvatarName           Hash
	QAvatarIcon           StrWithPrefix16
	QAvatarModel          StrWithPrefix16
	QAvatarExploreSkillID []int32
	QAvatarAttackSkillID  []int32
	QAvatarReaction       []AiCyberQavatarMetaData_Reaction
	QAvatarColor          StrWithPrefix16
	UnlockConditionList   StrWithPrefix16
}

type AiCyberQavatarMetaData_Reaction struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Bubble  Hash
	AnimKey StrWithPrefix16
}

type AiCyberQavatarSkillMetaData struct {
	// Fields
	QAvatarSkillID int32

	// Properties
	SkillType uint8
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	SkillName    Hash
	SkillDec     Hash
	SkillIcon    StrWithPrefix16
	SkillAbility StrWithPrefix16
}

type AiCyberStageGroupMetaData struct {
	// Fields
	StageGroupID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	StageIDList []int32
}

type AiCyberStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	QAvatarID           []int32
	QAvatarSkillID      []int32
	AvatarTrialID       []int32
}

type AiCyberStoryDataMetaData struct {
	// Fields
	StageGroupID int32
	Index        uint8

	// Properties
	StageID int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	QAvatarID           []int32
	QAvatarSkillID      []int32
}

type AnniversaryFifthBoxDataMetaData struct {
	// Fields
	BoxID uint32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	MaterialID uint32
	BoxType    uint32
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`

	// Objects
	InProgressImagePath StrWithPrefix16
	FinishedImagePath   StrWithPrefix16
	BoxStartTime        RemoteTime
	BoxEndTime          RemoteTime
	IntParams           []uint32
}

type AnniversaryFifthDetailItemsMetaData struct {
	// Fields
	DetailItemID uint32

	// Properties
	Addr1    uint32 `json:"-"`
	MinLevel uint32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`

	// Objects
	RewardImagePath StrWithPrefix16
	StartTime       RemoteTime
	EndTime         RemoteTime
	IntParam        []uint32
	ProgressDesc    Hash
}

type AnniversaryFifthShowItemsMetaData struct {
	// Fields
	ShowItemID uint32

	// Properties
	Addr1    uint32 `json:"-"`
	MinLevel uint32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	RewardImagePath StrWithPrefix16
	StartTime       RemoteTime
	EndTime         RemoteTime
	ProgressDesc    Hash
}

type AnniversaryIconintegrateMetaData struct {
	// Fields
	IconID int16

	// Properties
	IconX          int16
	IconY          int16
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	IconScheduleID int32
	Addr6          uint32 `json:"-"`
	MinPlayerLevel uint8
	MaxPlayerLevel uint8

	// Objects
	IconTitle Hash
	StartTime RemoteTime
	EndTime   RemoteTime
	IconGo    StrWithPrefix16
	IconImage StrWithPrefix16
	IconRule  Hash
}

type AppointmentDownloadScheduleMetaData struct {
	// Fields
	ID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	DownloadLimitTime StrWithPrefix16
	NoticeTextMapId   Hash
	ErrorTextMapId    Hash
}

type ArmadaBossActivityDataMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1  uint32 `json:"-"`
	ShopID int32
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`

	// Objects
	EndTime              RemoteTime
	LevelPanelPrefabPath StrWithPrefix16
	StageIDList          StrWithPrefix16
}

type ArmadaBossActivityScoreMetaData struct {
	// Fields
	ActivityID    int32
	ScoreType     int32
	ActivityScore int32

	// Properties
	Reward int32
}

type ArmadaBossData struct {
	// Fields
	Level int32

	// Properties
	ArmadaLevel int32
	Addr1       uint32 `json:"-"`
	CostFund    int32
	CostTime    int32
	ApCap       int32

	// Objects
	Prefab StrWithPrefix16
}

type ArmadaContactData struct {
	// Fields
	Level int32

	// Properties
	ArmadaLevel int32
	Addr1       uint32 `json:"-"`
	CostFund    int32
	CostTime    int32
	Capacity    int32
	PoplCap     int32

	// Objects
	Prefab StrWithPrefix16
}

type ArmadaExchequerData struct {
	// Fields
	Level int32

	// Properties
	ArmadaLevel int32
	Addr1       uint32 `json:"-"`
	CostFund    int32
	CostTime    int32
	FundCap     int32

	// Objects
	Prefab StrWithPrefix16
}

type ArmadaGridData struct {
	// Fields
	GridID int32

	// Properties
	X int32
	Y int32
}

type ArmadaHangarData struct {
	// Fields
	Level int32

	// Properties
	ArmadaLevel int32
	Addr1       uint32 `json:"-"`
	CostFund    int32
	CostTime    int32
	ApCap       int32

	// Objects
	Prefab StrWithPrefix16
}

type ArmadaLinerRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	NeedScore int32
	RewardID  int32
}

type ArmadaMainData struct {
	// Fields
	Level int32

	// Properties
	Grids        int32
	OfficerNum   int32
	ContactNum   int32
	HangarNum    int32
	BossNum      int32
	WorkshopNum  int32
	ExchequerNum int32
	CostFund     int32
	RequirePopl  int32
	CostTime     int32
}

type ArmadaPersonalizedLabelMetaData struct {
	// Fields
	PersonalizedLabelID int32

	// Properties
	Type  int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TextMap   StrWithPrefix16
	MutexList []int32
}

type ArmadaRecruitTipsMetaData struct {
	// Fields
	ArmadaLevel int32
	TYPE        int32

	// Properties
	Value1       int32
	Value2       int32
	VisibleRange int32
	Priority     int32
	Addr1        uint32 `json:"-"`

	// Objects
	TextMap_Tip StrWithPrefix16
}

type ArmadaReunionLevelMetaData struct {
	// Fields
	ScheduleID int32
	Level      int32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Displayreward1 int32
	Displayreward2 int32
	Displayreward3 int32
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`

	// Objects
	Demand               []ArmadaReunionLevelMetaData_LevelMaterialDemand
	RewardList           []int32
	PicPath              StrWithPrefix16
	TextMap_Level        StrWithPrefix16
	OriginalPaintingPath StrWithPrefix16
	ActivePaintingPath   StrWithPrefix16
}

type ArmadaReunionLevelMetaData_LevelMaterialDemand struct {
	// Fields
	MaterialID int32
	Number     int32
}

type ArmadaReunionMissionMetaData struct {
	// Fields
	ScheduleID             int32
	ArmadaReunionMissionID int32

	// Properties
	Addr1    uint32 `json:"-"`
	CountMax int32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	FinishTime        StrWithPrefix16
	MissionIntroImage StrWithPrefix16
	RewardDisplay     []ArmadaReunionMissionMetaData_MissionRewardDisplay
}

type ArmadaReunionMissionMetaData_MissionRewardDisplay struct {
	// Fields
	MaterialID int32
	Num        int32
}

type ArmadaReunionPointMetaData struct {
	// Fields
	Material int32

	// Properties
	Point int32
}

type ArmadaReunionScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	CurrencyID     int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	CoolingTime    int32
	MaxAwardsTimes int32
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`

	// Objects
	OpenTime        StrWithPrefix16
	RewardTime      StrWithPrefix16
	RewardEndTime   StrWithPrefix16
	DailyOpenTime   RemoteTimeSpan
	CloseTime       StrWithPrefix16
	PlayInstruction StrWithPrefix16
}

type ArmadaStageBossMetaData struct {
	// Fields
	BossID int32

	// Properties
	Addr1    uint32 `json:"-"`
	UniqueID int32
	Addr2    uint32 `json:"-"`

	// Objects
	Name      Hash
	ImagePath StrWithPrefix16
}

type ArmadaStageHardMetaData struct {
	// Fields
	HardLevel int32

	// Properties
	Addr1          uint32 `json:"-"`
	BossLevel      int32
	SummonPop      int32
	Addr2          uint32 `json:"-"`
	Stage1RewardID int32
	Stage2RewardID int32
	Stage3RewardID int32

	// Objects
	HardText   Hash
	KillReward []int32
}

type ArmadaWelcomeMetaData struct {
	// Fields
	WelcomeID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TextMap StrWithPrefix16
}

type ArmadaWorkshopData struct {
	// Fields
	Level int32

	// Properties
	ArmadaLevel int32
	Addr1       uint32 `json:"-"`
	CostFund    int32
	CostTime    int32
	TechLevel   int32

	// Objects
	Prefab StrWithPrefix16
}

type ArtifactTryEnterMetaData struct {
	// Fields
	ArtifactID int32

	// Properties
	Level_id          int32
	UnlockLevel       int32
	IsGeneralActivity bool
	GeneralActivityID int32
}

type AudioPackageData struct {
	// Fields
	ID int32

	// Properties
	Addr1                         uint32 `json:"-"`
	Addr2                         uint32 `json:"-"`
	PckType                       uint8
	AppointmentDownloadScheduleID uint32

	// Objects
	PackageName StrWithPrefix16
	Version     StrWithPrefix16
}

type AvatarAttackPunishMetaData struct {
	// Fields
	LevelDifference int32

	// Properties
	DamageReduceRate float32
}

type AvatarAttackSpeedParamMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	IsFixAttackSpeed int32
	K1               float32
	K2               float32
	K3               float32
	K4               float32
	K5               float32
	K6               float32
	K7               float32
	K8               float32
	K9               float32
	K10              float32
	K11              float32
	K12              float32
}

type AvatarBriefMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`
	Addr11 uint32 `json:"-"`
	Addr12 uint32 `json:"-"`

	// Objects
	Birthday     Hash
	Sex          Hash
	Organization Hash
	Height       Hash
	Weight       Hash
	HomePlace    Hash
	Story01      Hash
	Story02      Hash
	Story03      Hash
	Dialogue01   Hash
	Dialogue02   Hash
	Dialogue03   Hash
}

type AvatarDefencePunishMetaData struct {
	// Fields
	LevelDifference int32

	// Properties
	DamageIncreaseRate float32
}

type AvatarEffectInfoData struct {
	// Fields
	EffectID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	EffectIcon StrWithPrefix16
}

type AvatarExMetaData struct {
	// Fields
	SetID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	AvatarIDList       []int32
	CommonSubSkillList []AvatarExMetaData_CommonSkill
}

type AvatarExMetaData_CommonSkill struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Skill1 StrWithPrefix16
	Skill2 StrWithPrefix16
}

type AvatarFilterMetaData struct {
	// Fields
	Order int32

	// Properties
	Type            int32
	Value           int32
	Addr1           uint32 `json:"-"`
	IsCollaboration bool

	// Objects
	OrderDec Hash
}

type AvatarForgeRecommendMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	DefaultType uint8
	Addr1       uint32 `json:"-"`

	// Objects
	HeadPicPath StrWithPrefix16
}

type AvatarForgeWhiteListMetaData struct {
	// Fields
	AvatarID     int32
	LevelgroupID int32

	// Properties
	Addr1    uint32 `json:"-"`
	IsEffect bool

	// Objects
	ForgeList []int32
}

type AvatarGachaDisplayMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	Addr7            uint32 `json:"-"`
	Addr8            uint32 `json:"-"`
	WeaponID         int32
	Addr9            uint32 `json:"-"`
	Addr10           uint32 `json:"-"`
	TextPositionType int32
	TextShowDelay    float32
	Addr11           uint32 `json:"-"`

	// Objects
	AvatarUIPath                 StrWithPrefix16
	AvatarWinAnim                StrWithPrefix16
	AvatarIdleAnim               StrWithPrefix16
	AvatarAddWinAnim             StrWithPrefix16
	AvatarAddIdleAnim            StrWithPrefix16
	CameraAnimPath               StrWithPrefix16
	StagePath                    StrWithPrefix16
	AvatarAnimatorControllerPath StrWithPrefix16
	AudioPattern                 []AvatarGachaDisplayMetaData_AudioPatternEvent
	FaceAnimation                []AvatarGachaDisplayMetaData_FaceAnimationEvent
	StageOffset                  []float32
}

type AvatarGachaDisplayMetaData_AudioPatternEvent struct {
	// Fields
	Time  float32
	Addr1 uint32 `json:"-"`

	// Objects
	AudioPattern StrWithPrefix16
}

type AvatarGachaDisplayMetaData_FaceAnimationEvent struct {
	// Fields
	Time  float32
	Addr1 uint32 `json:"-"`

	// Objects
	FaceAnimation StrWithPrefix16
}

type AvatarGoodFeelMetaData struct {
	// Fields
	AvatarId      int32
	GoodFeelLevel int32
}

type AvatarLevelGroupMetaData struct {
	// Fields
	LevelGroup int32

	// Properties
	LevelMin int32
	LevelMax int32
}

type AvatarMemoirMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Unlock int32
}

type AvatarMissionActivityMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	Addr3                    uint32 `json:"-"`
	DailyContractPoint       int32
	DailyContractPointReward int32
	Addr4                    uint32 `json:"-"`

	// Objects
	BeginTime          StrWithPrefix16
	EndTime            StrWithPrefix16
	AvatarSampleIDList []int32
	MissionList        []int32
}

type AvatarMissionActivityRewardMetaData struct {
	// Fields
	ID              int32
	AccumulatedDays int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	RewardIDList []AvatarMissionActivityRewardMetaData_SampleReward
}

type AvatarMissionActivityRewardMetaData_SampleReward struct {
	// Fields
	RewardID int32
	SampleID int32
}

type AvatarNewbieMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TextMapID Hash
}

type AvatarPracticeMainMetaData struct {
	// Fields
	MainID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	StepIDs []int32
}

type AvatarPracticeStepMetaData struct {
	// Fields
	StepID int32

	// Properties
	Type         int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	UseDuration  bool
	Duration     float32
	HoldDuration float32
	Addr6        uint32 `json:"-"`
	TimeOffset   float32
	UseSyncBling bool
	Addr7        uint32 `json:"-"`

	// Objects
	Text              Hash
	Param             []StrWithPrefix16
	PrefabType        StrWithPrefix16
	SkillPrefabType   StrWithPrefix16
	PrefabImagePath   StrWithPrefix16
	AnimatorStateName StrWithPrefix16
	SKLButton         StrWithPrefix16
}

type AvatarRelayMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	RelayPhaseList []uint16
	StageImgPath   StrWithPrefix16
	StageDesc      Hash
}

type AvatarRewardGroupMetaData struct {
	// Fields
	GroupID int32

	// Properties
	NeedProcess int32
	RewardType  int32
	RewardID    int32
}

type AvatarRoleMetaData struct {
	// Fields
	RoleID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	AvatarRoleIcon   StrWithPrefix16
	AvatarRoleName   Hash
	AvatarRoleNameEn Hash
}

type AvatarSubSkillLevelMetaData struct {
	// Fields
	ItemType   int32
	SkillLevel int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ItemList_1 []AvatarSubSkillLevelMetaData_SkillUpLevelNeedItem
}

type AvatarSubSkillLevelMetaData_SkillUpLevelNeedItem struct {
	// Fields
	ItemMetaID int32
	ItemNum    int32
}

type AvatarTrialMetaData struct {
	// Fields
	ID uint16

	// Properties
	AvatarID         uint16
	AvatarLV         uint8
	AvatarStar       uint8
	UseGrandKey      bool
	UseMedal         bool
	UseIsland        bool
	UseNewIslandBuff bool
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	WeaponID         int32
	WeaponLV         uint8
	StigmataID1      int32
	StigmataLV1      uint8
	StigmataID2      int32
	StigmataLV2      uint8
	StigmataID3      int32
	StigmataLV3      uint8
	IsAvatarArtifact bool
	ArtifactStar     int32
	AvatarDress      int32

	// Objects
	AvatarSubSkills  []AvatarTrialMetaData_SkillInfoItem
	DLCAvatarTalents []AvatarTrialMetaData_TalentInfoItem
}

type AvatarTrialMetaData_SkillInfoItem struct {
	// Fields
	SkillID int32
	SkillLV int32
}

type AvatarTrialMetaData_TalentInfoItem struct {
	// Fields
	Equiped  int32
	TalentID uint32
	TalentLV uint32
}

type AvatarTryEnterMetaData struct {
	// Fields
	Avatar_id uint16

	// Properties
	Level_id          int32
	Addr1             uint32 `json:"-"`
	UnlockLevel       int32
	IsGeneralActivity bool
	GeneralActivityID int32

	// Objects
	AvatarTry_textmap Hash
}

type AvatarTutorialMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	AvatarID        int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Width           int32
	Addr6           uint32 `json:"-"`
	EntryWeight     uint32
	IsActivityEntry bool
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`

	// Objects
	ArtifactPath    StrWithPrefix16
	TitleName       Hash
	SiteList        []uint32
	AvatarImgOffset []float32
	MissionList     []int32
	GachatypeList   []int32
	TagBeginTime    StrWithPrefix16
	TagEndTime      StrWithPrefix16
	InfoTextMap     Hash
}

type AvatarTutorialSiteMetaData struct {
	// Fields
	SiteID uint32

	// Properties
	Site          uint8
	Addr1         uint32 `json:"-"`
	SiteStage     int32
	ShowSiteLevel int32
	Addr2         uint32 `json:"-"`

	// Objects
	SiteName  Hash
	VideoLink StrWithPrefix16
}

type AvatarTwinsMetaData struct {
	// Fields
	MainAvatarID int32

	// Properties
	Addr1          uint32 `json:"-"`
	SecondAvatarID int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`

	// Objects
	StandByMotion1     StrWithPrefix16
	StandByMotion2     StrWithPrefix16
	ExcludeDressIDList []int32
}

type AvatarVideoPopupMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Video1LinkOpenType uint8
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Video2LinkOpenType uint8
	Addr7              uint32 `json:"-"`
	MaxTimes           int32

	// Objects
	MiddleImageSource StrWithPrefix16
	Video1Icon        StrWithPrefix16
	Video1Link        StrWithPrefix16
	Video1Position    []int32
	Video2Icon        StrWithPrefix16
	Video2Link        StrWithPrefix16
	Video2Position    []int32
}

type BattleDisplayData struct {
	// Fields
	BattleDisplayID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	BuffText    Hash
	LoadingTips Hash
}

type BattlePassBannerMetaData struct {
	// Fields
	ID uint32

	// Properties
	ScheduleID uint16
	MinLevel   uint16
	MaxLevel   uint16
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	LinkDataID uint32
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`

	// Objects
	BPIdentification []uint8
	BannerImgPath    StrWithPrefix16
	Description      StrWithPrefix16
	Description2     StrWithPrefix16
	EndTime          RemoteTime
}

type BattlePassChallengeDataMetaData struct {
	// Fields
	MissionID int32

	// Properties
	ThemeExpReward  uint16
	SeasonExpReward uint16
	CountMax        uint8
	SortID          uint8
	Addr1           uint32 `json:"-"`
	BattlePassType  uint8

	// Objects
	MissionIntroImage StrWithPrefix16
}

type BattlePassScheduleMetaData struct {
	// Fields
	SeasonScheduleID uint8

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	SeasonLevelConfigID   uint8
	BasicTypeID           int32
	AdvancedTypeID        int32
	LuxuryTypeID          int32
	ExtraRewardID         int32
	ShowRewardID          int32
	Addr3                 uint32 `json:"-"`
	ExpHcoinPrice         uint16
	SeasonExpMaterialID   int32
	Addr4                 uint32 `json:"-"`
	IntroImageAnimTime    uint16
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	Addr7                 uint32 `json:"-"`
	Addr8                 uint32 `json:"-"`
	Addr9                 uint32 `json:"-"`
	BpCoinID              int32
	CrystalID             int32
	Addr10                uint32 `json:"-"`
	Addr11                uint32 `json:"-"`
	Addr12                uint32 `json:"-"`
	MaxSeasonLevel        uint16
	MaxDisplaySeasonLevel uint16
	Addr13                uint32 `json:"-"`

	// Objects
	BeginTime            RemoteTime
	EndTime              RemoteTime
	ShopEnterIcon        StrWithPrefix16
	IntroImgPathList     []StrWithPrefix16
	PurchaseBeginTime    StrWithPrefix16
	PurchaseEndTime      StrWithPrefix16
	DisplayBeginTime     RemoteTime
	AdvancedRewardFormat []BattlePassScheduleMetaData_Format
	LuxuryRewardFormat   []BattlePassScheduleMetaData_Format
	AvatarSprite         StrWithPrefix16
	LastingDesc          Hash
	RewardnoticeFreePic  StrWithPrefix16
	DefaultPic           StrWithPrefix16
}

type BattlePassScheduleMetaData_Format struct {
	// Fields
	Type  uint8
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Text  Hash
	Param StrWithPrefix16
}

type BattlePassSeasonLevelConfigMetaData struct {
	// Fields
	SeasonLevelComfigID uint8
	SeasonLevel         uint8

	// Properties
	NextLevelExp         uint16
	IsNotice             bool
	Addr1                uint32 `json:"-"`
	BasicRewardID        int32
	AdvancedRewardID     int32
	EliteRewardID        int32
	IsConvenientPurchase bool

	// Objects
	RewardNoticeImgPath StrWithPrefix16
}

type BattlePassTypeMetaData struct {
	// Fields
	TypeID int32

	// Properties
	OriginalPrice int32
	DiscountPrice int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`

	// Objects
	DisplayIcon          StrWithPrefix16
	DisplayImage         StrWithPrefix16
	DisplayExpBoxIcon    StrWithPrefix16
	DisplayExpBoxTextMap StrWithPrefix16
}

type BattleTypeMetaData struct {
	// Fields
	BattleTypeId int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type BbqBonusSetMetaData struct {
	// Fields
	Type int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	MovingSpd int32

	// Objects
	Bonus_1 []float32
	Bonus_2 []float32
}

type BbqLevelMetaData struct {
	// Fields
	BbqLevel int32

	// Properties
	LevelUpExp      int32
	LevelUpRewardID int32
}

type BeastStageDisplayMetaData struct {
	// Fields
	BeastStageID int32
	Level        int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	LevelDesc            StrWithPrefix16
	LevelDropDisplayList []BeastStageDisplayMetaData_LevelDropDisplay
}

type BeastStageDisplayMetaData_LevelDropDisplay struct {
	// Fields
	Count      int32
	MaterialID int32
}

type BeastStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	StageHpTotal int32
}

type BeastTreasureMetaData struct {
	// Fields
	TreasureRankID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TreasurePath StrWithPrefix16
}

type BossChallengeActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	StageIDList []int32
}

type BossChallengeRewardMetaData struct {
	// Fields
	RewardConfigID int32
	Index          uint16

	// Properties
	Time     int32
	RewardID int32
}

type BossChallengeStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Index               int32
	Addr1               uint32 `json:"-"`
	CollectionMonsterID int32
	RewardConfigID      int32
	MaxFloor            uint16
	IsShowEffect        bool
	Addr2               uint32 `json:"-"`

	// Objects
	UnlockTime RemoteTime
	BossName   Hash
}

type BossPointScoreRewardMetaData struct {
	// Fields
	ID uint32

	// Properties
	BossPointScheduleID uint32
	Score               uint64
	RewardID            uint32
}

type BossRushBuffConfigMetaData struct {
	// Fields
	BuffConfigID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	BuffGridUnlockTimeOffsets []uint32
	BuffPoolIDs               []uint32
}

type BossRushBuffMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	BuffType int32

	// Objects
	BossRushBuffName    Hash
	BossRushBuffDes     Hash
	BossRushAbilityName StrWithPrefix16
	IconPath            StrWithPrefix16
}

type BossRushBuffPoolMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	BossRushBuffID []int32
}

type BossRushStageScheduleMetaData struct {
	// Fields
	ID uint32

	// Properties
	StartTimeOffset uint32
	EndTimeOffset   uint32
	Addr1           uint32 `json:"-"`
	BuffConfigID    uint32
	CoinMaterialID  int32

	// Objects
	DisplayedStageGroups []int32
}

type BoxGachaMetaData struct {
	// Fields
	GachaID int32

	// Properties
	IsHide bool
}

type BridgeLinkMetaData struct {
	// Fields
	LinkID uint8

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	LinkMinLevel uint8
	LinkMaxLevel uint8
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`

	// Objects
	BeginTime     StrWithPrefix16
	EndTime       StrWithPrefix16
	LinkParam     StrWithPrefix16
	BridgeLinkPic StrWithPrefix16
}

type BubbleBridgeMetaData struct {
	// Fields
	Id int32

	// Properties
	TriggerMissionID int32
	AvatarID         int32
	Type             int32
	GalEventID       int32
}

type BuffAssistanceActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	Addr3                   uint32 `json:"-"`
	Addr4                   uint32 `json:"-"`
	BuffPoolMaxRefreshTimes uint8
	Addr5                   uint32 `json:"-"`
	Addr6                   uint32 `json:"-"`
	Addr7                   uint32 `json:"-"`

	// Objects
	ActivityName        Hash
	ActivityIconPath    StrWithPrefix16
	UnlockConditionList StrWithPrefix16
	LockTips            Hash
	MissionIDList       []int32
	BeginTime           RemoteTime
	EndTime             RemoteTime
}

type BuffAssistanceBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	BuffType uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`

	// Objects
	BuffName              Hash
	BuffDescription       Hash
	BuffSimpleDescription Hash
	BuffIconPath          StrWithPrefix16
	BuffStringName        StrWithPrefix16
	BuffParamList         []StrWithPrefix16
}

type BuffAssistanceLevelMetaData struct {
	// Fields
	Level uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Title Hash
}

type BuffAssistanceNPCRobotMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	NPCIconPath    StrWithPrefix16
	NPCPicturePath StrWithPrefix16
	BuffList       []int32
}

type BuffAssistanceScheduleMetaData struct {
	// Fields
	BuffScheduleID int32

	// Properties
	Addr1     uint32 `json:"-"`
	DeBuffNum uint8
	Addr2     uint32 `json:"-"`

	// Objects
	ScheduleName Hash
	EndTime      RemoteTime
}

type BuffAssistanceStageGroupMetaData struct {
	// Fields
	StageGroupID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	NormalStageIDList       []int32
	ChallengeStageIDList    []int32
	BuffSlotPreStageList    []BuffAssistanceStageGroupMetaData_BuffAssistanceStageBuffNum
	DebuffSlotPreStageList  []BuffAssistanceStageGroupMetaData_BuffAssistanceStageBuffNum
	StageRewardNumRangeList []int32
}

type BuffAssistanceStageGroupMetaData_BuffAssistanceStageBuffNum struct {
	// Fields
	BuffNum uint8
	StageID int32
}

type BuffAssistanceWordMetaData struct {
	// Fields
	BuffAssistWordID int32

	// Properties
	Addr1      uint32 `json:"-"`
	CanBeFirst uint8
	Addr2      uint32 `json:"-"`
	CanBeLast  uint8

	// Objects
	AssistWordTextmapID Hash
	ContinueWordID      []int32
}

type BulletinEffectMetaData struct {
	// Fields
	ConfigID int32

	// Properties
	LoadAnimeLastTime uint16
}

type BurdenAlleviationMetaData struct {
	// Fields
	ID int32

	// Properties
	Type          int32
	TypeParam     int32
	Addr1         uint32 `json:"-"`
	MaxNum        int32
	Priority      int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	RewardDisplay int32
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	Addr8         uint32 `json:"-"`
	Addr9         uint32 `json:"-"`
	ShowPopup     int32

	// Objects
	TypeParamStrClient []int32
	EndTime            RemoteTime
	BeginTime          RemoteTime
	BGPath             StrWithPrefix16
	DisplayInfo        StrWithPrefix16
	NotOpenInfoText    StrWithPrefix16
	CloseInfoText      StrWithPrefix16
	BurdenDialogTitle  StrWithPrefix16
	BurdenDialogDesc   StrWithPrefix16
}

type BurialPointMetaData struct {
	// Fields
	ID int32

	// Properties
	TutoriaStepId int32
}

type CgCategoryMetaData struct {
	// Fields
	CgCategory int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	CgCategory_Name   Hash
	CgCategory_NameEn Hash
	CgCategory_Icon   StrWithPrefix16
}

type CgGroupDataMetaData struct {
	// Fields
	CgGroupID int32

	// Properties
	PackageID      int32
	CgDownLoadMode int16
}

type CgMetaData struct {
	// Fields
	CgID int32

	// Properties
	UnlockType                    uint8
	UnlockCondition               uint32
	LevelIDBegin                  int32
	LevelIDEnd                    int32
	CgCategory                    int32
	CgSubCategory                 int32
	Addr1                         uint32 `json:"-"`
	WikiCgScore                   int32
	InitialUnlock                 bool
	Addr2                         uint32 `json:"-"`
	Addr3                         uint32 `json:"-"`
	Addr4                         uint32 `json:"-"`
	InStreamingAssets             int32
	CgPlayMode                    int32
	Addr5                         uint32 `json:"-"`
	FileSize                      int32
	PckType                       uint8
	Addr6                         uint32 `json:"-"`
	AppointmentDownloadScheduleID uint32

	// Objects
	CgGroupID         []int32
	CgPath            StrWithPrefix16
	CgIconSpritePath  StrWithPrefix16
	CgLockHint        Hash
	CgExtraKey        StrWithPrefix16
	DownloadLimitTime StrWithPrefix16
}

type CGPackageDataMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Title Hash
	Desc  Hash
}

type ChallengeMissionLink struct {
	// Fields
	MissionID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	DisplayIconPath StrWithPrefix16
	ComeFromList    []StrWithPrefix16
}

type ChallengeMissionPanel struct {
	// Fields
	TypeID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TypePanelTexture StrWithPrefix16
}

type ChallengeMissionStepData struct {
	// Fields
	StepID int32

	// Properties
	TypeID                     int32
	Addr1                      uint32 `json:"-"`
	ChallengeMissionPreStep    int32
	ChallengeMissionStepReward int32
	StepUnlockLevel            int32
	ChallengeMissionStepBonus  int32

	// Objects
	ChallengeMissionList []int32
}

type ChallengeWarAchieveMissionMetaData struct {
	// Fields
	MissionID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	DisplaySiteName Hash
}

type ChallengeWarAreaInfoMetaData struct {
	// Fields
	AreaID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	AreaDisplayBG StrWithPrefix16
}

type ChallengeWarBuffDataMetaData struct {
	// Fields
	BuffID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	BuffParam1    float32
	BuffParam1Add float32
	BuffParam2    float32
	BuffParam2Add float32
	BuffParam3    float32
	BuffParam3Add float32

	// Objects
	BuffName       Hash
	BuffDesc       Hash
	BuffAddDesc    Hash
	BuffIcon       StrWithPrefix16
	BuffDetailDesc StrWithPrefix16
}

type ChallengeWarSeasonMetaData struct {
	// Fields
	SeasonID int32

	// Properties
	SeasonMissionGroupID int32
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`

	// Objects
	AchieveTaleIDList    []int32
	RPGScheduleIDList    []int32
	AreaAchieveGroupList []int32
	SeasonDisplayTitle   Hash
	EndTime              RemoteTime
}

type ChallengeWarStageDataMetaData struct {
	// Fields
	StageID int32

	// Properties
	GroupID          int32
	Addr1            uint32 `json:"-"`
	SupportAvatarNum uint8
	UseSubStar       bool
	Addr2            uint32 `json:"-"`
	SubStarDisplay   bool

	// Objects
	SeriesIDList     []uint16
	StageCoverBGPath StrWithPrefix16
}

type ChallengeWarStageFloorMetaData struct {
	// Fields
	StageID int32
	Floor   int16
}

type ChallengeWarTagEffectPoolMetaData struct {
	// Fields
	SeriesID uint16

	// Properties
	PoolType uint16
	TagID    uint16
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	MainTag  bool
	Addr4    uint32 `json:"-"`

	// Objects
	RequireParam []int32
	RequireList  []uint8
	BuffList     []int32
	DisplayText  StrWithPrefix16
}

type ChanllengeWarAvatarSubStarMetaData struct {
	// Fields
	AvatarUnlockStar uint8
	Star             uint8
	SubStar          uint8

	// Properties
	Value uint8
}

type ChapterActivityRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	ChapterID    int32
	EventLevel   int32
	NeedEventExp int32
	RewardID     int32
	Addr1        uint32 `json:"-"`

	// Objects
	UnlockTime StrWithPrefix16
}

type ChapterActivityRpgAvatarMetaData struct {
	// Fields
	RpgAreaID uint16

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	AvatarType          uint8
	Addr3               uint32 `json:"-"`
	InitiativeTalentID  int32
	ChallengeSiteTyle   uint8
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	ThemeAvatarID       int32
	Addr9               uint32 `json:"-"`
	AvatarTutorialID    int32
	Addr10              uint32 `json:"-"`
	SweepLevelDropLimit int32

	// Objects
	NormalStageRowPrefabPath    StrWithPrefix16
	ChallengeStageRowPrefabPath StrWithPrefix16
	PassiveTalentIDList         []int32
	AvatarImgPath               StrWithPrefix16
	AvatarThumbnail             StrWithPrefix16
	EntryAvatarBGPath           StrWithPrefix16
	AvatarName                  Hash
	AvatarRestrictDisplayList   []int32
	CurrencyDisplay             []int32
	WebLink                     StrWithPrefix16
}

type ChapterActivitySectionMetaData struct {
	// Fields
	ID int32

	// Properties
	ChapterID      int32
	IconPos        int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	SPActivityType int32
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	UnlockLevel    int32
	ShowLevel      int32
	Addr8          uint32 `json:"-"`
	PreLevelID     int32

	// Objects
	SPActName            Hash
	SPActDesc            Hash
	UnlockDesc           Hash
	EntryBG              []StrWithPrefix16
	ActList              []int32
	ActivityScheduleList []int32
	UnlockLink           StrWithPrefix16
	EndTime              StrWithPrefix16
}

type ChapterActivityStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	IsAllowNormalAvatar bool
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`

	// Objects
	StageEventList             []StrWithPrefix16
	StageEventHideTextList     []StrWithPrefix16
	StageEventCompleteTextList []StrWithPrefix16
	StageEventDescTextList     []StrWithPrefix16
	StageEventIconList         []StrWithPrefix16
	StageMark                  StrWithPrefix16
}

type ChapterActivityZoneBuffDataMetaData struct {
	// Fields
	BuffID uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	BuffDesc StrWithPrefix16
}

type ChapterActivityZoneDataMetaData struct {
	// Fields
	ZoneID int32

	// Properties
	ChapterID    int32
	NumInChapter int32
	PreZone      int32
	Addr1        uint32 `json:"-"`
	DropLimit    int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	Addr7        uint32 `json:"-"`
	Addr8        uint32 `json:"-"`
	Addr9        uint32 `json:"-"`
	Addr10       uint32 `json:"-"`
	Addr11       uint32 `json:"-"`

	// Objects
	OpenTime        RemoteTime
	RepeatStages    []int32
	ChallengeStages []int32
	BuffRequireList []int32
	BuffList        []int32
	BuffTitle       StrWithPrefix16
	BuffDesc        StrWithPrefix16
	BuffDescShort   StrWithPrefix16
	BuffIcon        StrWithPrefix16
	ZoneOpenHint    StrWithPrefix16
	ZoneRefuseHint  StrWithPrefix16
}

type ChapterBuffDisplayMetaData struct {
	// Fields
	ID uint16

	// Properties
	EnterTime uint16
	Buff1     uint16
	Buff2     uint16
	Buff3     uint16
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	Desc        Hash
	UpgradeDesc Hash
}

type ChapterChallengeMetaData struct {
	// Fields
	ChapterID    int32
	ChallengeNum int32

	// Properties
	RewardID int32
}

type ChapterCultivateAttrMetaData struct {
	// Fields
	CultivateID    uint32
	CultivateLevel uint32

	// Properties
	HP             uint32
	ATK            uint32
	DEF            uint32
	AllDamageRatio uint32
}

type ChapterCultivateBuffMetaData struct {
	// Fields
	ID    uint32
	Level uint32

	// Properties
	ActID                 uint32
	BuffGoodsID           uint32
	RequireBuffMaterialID uint32
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	RowGroupID            uint8
	Phase                 uint8

	// Objects
	IconPath StrWithPrefix16
	Name     Hash
	Desc     Hash
}

type ChapterCultivateBuffShopMetaData struct {
	// Fields
	ChapterID uint32

	// Properties
	ShopID         uint32
	ShopCurrencyID uint32
	Addr1          uint32 `json:"-"`

	// Objects
	Titles []Hash
}

type ChapterCultivateConfigMetaData struct {
	// Fields
	CultivateID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Title    StrWithPrefix16
	SubTitle StrWithPrefix16
	Desc     StrWithPrefix16
}

type ChapterCultivateJumpMetaData struct {
	// Fields
	CultivateLevel uint32
	ChapterID      int32

	// Properties
	CultivateSiteID uint32
}

type ChapterCultivateLevelDataMetaData struct {
	// Fields
	CultivateUseID uint32
	Level          uint32

	// Properties
	CostMaterialID uint32
	CostNum        uint32
}

type ChapterCultivateMonsterMetaData struct {
	// Fields
	SkillID    uint32
	SkillLevel uint32

	// Properties
	MonsterID uint32
}

type ChapterCultivateSiteEffectMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1                uint32 `json:"-"`
	MaxFloor             int32
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	Addr6                uint32 `json:"-"`
	UpAvatarConditionNum int32
	Addr7                uint32 `json:"-"`
	Addr8                uint32 `json:"-"`

	// Objects
	LevelJumpDisplay      []ChapterCultivateSiteEffectMetaData_LevelJump
	LevelEffectList       []ChapterCultivateSiteEffectMetaData_LevelEffect
	EnterTimeEffectList   StrWithPrefix16
	RandomEffectList      []int32
	BaseEffectList        []int32
	UpAvatarList          []int32
	UpAvatarConditionIcon StrWithPrefix16
	UpAvatarConditionText Hash
}

type ChapterCultivateSiteEffectMetaData_LevelEffect struct {
	// Fields
	BuffID int32
	Level  int32
}

type ChapterCultivateSiteEffectMetaData_LevelJump struct {
	// Fields
	FloorBlock int32
	NumHint    int32
}

type ChapterCultivateSkillMetaData struct {
	// Fields
	CultivateID    uint32
	CultivateLevel uint32

	// Properties
	SkillType  uint8
	SkillID    uint32
	SkillLevel uint32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
	Name     StrWithPrefix16
	Desc     StrWithPrefix16
}

type ChapterCultivateUseMetaData struct {
	// Fields
	ID uint32

	// Properties
	ChapterID uint32
	UseType   uint8
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	CultivateUseIDList []uint32
	CultivateIDList    []uint32
}

type ChapterGroupConfigMetaData struct {
	// Fields
	ID uint8

	// Properties
	GroupType          uint8
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	BeginShowLevel     uint8
	Addr3              uint32 `json:"-"`
	UnlockLevel        uint8
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	LeftFuncBtn        uint16
	Addr7              uint32 `json:"-"`
	AutoShowPreviously bool
	Addr8              uint32 `json:"-"`
	Addr9              uint32 `json:"-"`

	// Objects
	BeginShowTime     RemoteTime
	BeginTime         RemoteTime
	SiteList          []uint8
	Title             Hash
	SubTitle          Hash
	BGPath            StrWithPrefix16
	FuncBtnList       []uint16
	PreviouslyImgPath StrWithPrefix16
	PreviouslyDesc    Hash
}

type ChapterGroupFuncButtonMetaData struct {
	// Fields
	ID uint16

	// Properties
	Type  uint8
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ImagePath StrWithPrefix16
	Name      Hash
	FuncParam StrWithPrefix16
}

type ChapterGroupSiteMetaData struct {
	// Fields
	SiteID uint8

	// Properties
	SiteType    uint8
	Addr1       uint32 `json:"-"`
	SiteChapter int32
	Addr2       uint32 `json:"-"`
	PopUpType   uint8

	// Objects
	SiteName Hash
	UIConfig StrWithPrefix16
}

type ChapterMetaData struct {
	// Fields
	ChapterId int32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	ChapterType       int32
	VirtualGroupID    int32
	ChapterMaxLevelID int32
	DefaultSection    int32
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	Addr12            uint32 `json:"-"`
	Addr13            uint32 `json:"-"`
	LoginDayLimit     uint8

	// Objects
	Title                  Hash
	ChapterDetail          Hash
	MonsterList            []int32
	CoverPicture           StrWithPrefix16
	UnlockDisplay          Hash
	ActiviyDesc            Hash
	EventBGPrefabPath      StrWithPrefix16
	EventSectionPrefabPath StrWithPrefix16
	BgImgPath              StrWithPrefix16
	CoverChapter           StrWithPrefix16
	CoverChapterColor      StrWithPrefix16
	VersionTagTitle        Hash
	MissionList            []int32
}

type ChapterOW31DialogueMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Duration float32

	// Objects
	TipsTextMap StrWithPrefix16
}

type ChapterOWAchievementMetaData struct {
	// Fields
	MissionID uint32

	// Properties
	MapID          uint32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	IsHide         bool
	IsShowProgress bool
	SequenceID     int32

	// Objects
	Name       Hash
	LockedName Hash
	Desc       Hash
	LockedDesc Hash
	ImgPath    StrWithPrefix16
}

type ChapterOWActivityPanelMetaData struct {
	// Fields
	ItemID uint16

	// Properties
	MapID           int32
	ItemType        uint8
	MaterialType    uint8
	MonitorMaterial int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`

	// Objects
	NumberList   []int32
	ItemTilte    Hash
	ItemIconPath StrWithPrefix16
	ScoreDisplay Hash
	LinkParam    StrWithPrefix16
}

type ChapterOWAntiGravityGroupMetaData struct {
	// Fields
	ID int32

	// Properties
	MapID            int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	LocationID       int32
	DisplayCondition int32
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`

	// Objects
	LevelIDList         []int32
	TalentIDList        []int32
	MaterialID          []int32
	UnlockConditionTips Hash
	Title               Hash
}

type ChapterOWAntiGravityLevelMetaData struct {
	// Fields
	ID int32

	// Properties
	MapID     int32
	GroupID   uint8
	StageID   int32
	LevelType uint8
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`
	Addr7     uint32 `json:"-"`
	Addr8     uint32 `json:"-"`
	Addr9     uint32 `json:"-"`
	Addr10    uint32 `json:"-"`

	// Objects
	ChallengeShowList      []ChapterOWAntiGravityLevelMetaData_ChallengeShowCondition
	SpecialChallengeIDList []int32
	SeniorCoinIndexList    []int32
	LowerCoinIndexList     []int32
	LowerCoinMaterial      ChapterOWAntiGravityLevelMetaData_CoinMaterialData
	SeniorCoinMaterial     ChapterOWAntiGravityLevelMetaData_CoinMaterialData
	UnlockConditionList    StrWithPrefix16
	UnlockConditionTips    []StrWithPrefix16
	ImagePath              StrWithPrefix16
	Title                  Hash
}

type ChapterOWAntiGravityLevelMetaData_CoinMaterialData struct {
	// Fields
	MaterialID int32
	RatioNum   int32
}

type ChapterOWAntiGravityLevelMetaData_ChallengeShowCondition struct {
	// Fields
	ChallengeID int32
	ConditionID int32
}

type ChapterOWBagMetaData struct {
	// Fields
	BagID uint32

	// Properties
	Type  uint8
	Param uint32
}

type ChapterOWBuildingLevelMetaData struct {
	// Fields
	MainID uint32
	Level  uint32

	// Properties
	BuildingType        uint8
	Addr1               uint32 `json:"-"`
	CostTime            uint32
	NeedTerminalLevel   uint32
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	Addr9               uint32 `json:"-"`
	Addr10              uint32 `json:"-"`
	ProgressPanelRadius float32
	ProgressPanelHeight float32

	// Objects
	Tips                 Hash
	CostMaterialList     []ChapterOWBuildingLevelMetaData_CostMaterialData
	PrefabPath           StrWithPrefix16
	UpgradeAddPrefabList []StrWithPrefix16
	UpgradeAddPathList   []StrWithPrefix16
	UpgradeOpenPathList  []StrWithPrefix16
	UpgradeClosePathList []StrWithPrefix16
	AnimPathList         []StrWithPrefix16
	IconPath             StrWithPrefix16
	RemindTipsList       []ChapterOWBuildingLevelMetaData_RemindTipData
}

type ChapterOWBuildingLevelMetaData_CostMaterialData struct {
	// Fields
	CostNum    int32
	MaterialID int32
}

type ChapterOWBuildingLevelMetaData_RemindTipData struct {
	// Fields
	ConditionID int32
	Addr1       uint32 `json:"-"`

	// Objects
	ConditionTextID Hash
}

type ChapterOWBuildingMetaData struct {
	// Fields
	MainID uint32

	// Properties
	MapID           int32
	BuildingType    uint8
	IsLobbyBuilding uint8
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	InitLevel       uint32
	MaxLevel        uint32
	Addr6           uint32 `json:"-"`
	IsShow          bool
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`
	Addr10          uint32 `json:"-"`
	Addr11          uint32 `json:"-"`

	// Objects
	BuildingName              Hash
	RepairDesc                Hash
	RepairSuccDesc            Hash
	StateTypeList             []int32
	ConditionList             []ChapterOWBuildingMetaData_ConditionStatePair
	MaterialList              []int32
	PrefabPath                StrWithPrefix16
	IconPath                  StrWithPrefix16
	MainStoryIDList           []int32
	BranchStoryIDList         []int32
	NPCHeadMarkOffsetPosition []float32
}

type ChapterOWBuildingMetaData_ConditionStatePair struct {
	// Fields
	ConditionID int32
	StateID     int32
}

type ChapterOWBuildingStateMetaData struct {
	// Fields
	MainID    int32
	StateType uint8

	// Properties
	Addr1    uint32 `json:"-"`
	PopType  uint8
	Addr2    uint32 `json:"-"`
	ShowType uint8

	// Objects
	Prefabpath      StrWithPrefix16
	TerminalTextmap Hash
}

type ChapterOWChallengeBuffMetaData struct {
	// Fields
	BuffID uint32

	// Properties
	Type  uint8
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Icon StrWithPrefix16
	Desc Hash
}

type ChapterOWChallengeDataMetaData struct {
	// Fields
	ConfigID       uint32
	ChallengeIndex uint32

	// Properties
	ShowRewardID uint32
	RewardID     uint32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	RecommendBP  int32
	Addr4        uint32 `json:"-"`

	// Objects
	Title         Hash
	MonsterIDList []int32
	BuffList      []int32
	FailedNote    []Hash
}

type ChapterOWChallengeGroupMetaData struct {
	// Fields
	GroupID uint32

	// Properties
	LocationID        uint32
	ChallengeConfigID uint32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	RecommendHero     uint32
	Addr4             uint32 `json:"-"`

	// Objects
	UnlockConditionList   StrWithPrefix16
	OpenTimeScheduleText  Hash
	Title                 Hash
	HighlightMaterialList []uint32
}

type ChapterOWChallengeSiteMetaData struct {
	// Fields
	ChallengeSiteID int32

	// Properties
	MapID              int32
	EventID            int32
	ChallengeHardLevel int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	ChallengeRewardID  int32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	MonsterImgPath      StrWithPrefix16
	RecommendBuffList   []int32
	Title               Hash
	Icon                StrWithPrefix16
	Desc                Hash
	MapDisplayTitle     Hash
}

type ChapterOWCollectionTabMetaData struct {
	// Fields
	ID uint32

	// Properties
	MapID          uint32
	CollectionType uint8
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`

	// Objects
	TabName StrWithPrefix16
	TabDesc StrWithPrefix16
}

type ChapterOWConfigMetaData struct {
	// Fields
	MapID int32

	// Properties
	ChapterId                int32
	Type                     int32
	PerformID                int32
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	Addr3                    uint32 `json:"-"`
	MainstageEventId         int32
	Addr4                    uint32 `json:"-"`
	RewardCountMaterial      int32
	Addr5                    uint32 `json:"-"`
	Addr6                    uint32 `json:"-"`
	Addr7                    uint32 `json:"-"`
	Addr8                    uint32 `json:"-"`
	GeniusTowerConutMaterial int32
	Addr9                    uint32 `json:"-"`
	CostMaterial             int32
	FameMaterial             int32
	FameShopMaterial         int32
	Addr10                   uint32 `json:"-"`
	DailyChallengeNumber     int32
	Addr11                   uint32 `json:"-"`
	Addr12                   uint32 `json:"-"`
	Addr13                   uint32 `json:"-"`
	DailyChallengeTowerTimes int32

	// Objects
	GroupTypeList          []int32
	PhaseIDList            []int32
	FortIDList             []int32
	EquipmentPartIDList    []int32
	RewardLimitEndTime     RemoteTime
	AchievementMissionList []int32
	UIManager              StrWithPrefix16
	GeniusTowerList        []uint8
	JsonPath               StrWithPrefix16
	PhotoIdList            []int32
	PlotGroupList          []int32
	MemoryIDList           []int32
	BagDataIDList          []int32
}

type ChapterOWDailyStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Multiple   int32
	Group      int32
	LocationID int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`

	// Objects
	UnlockConditionList      StrWithPrefix16
	HighLightDropDisplayList []int32
	MonsterList              []int32
	StageInfoText            StrWithPrefix16
}

type ChapterOWDigsiteMetaData struct {
	// Fields
	MainID uint32

	// Properties
	Addr1   uint32 `json:"-"`
	EventID int32
	Addr2   uint32 `json:"-"`

	// Objects
	DisplayRewardList []int32
	ProgramList       []uint16
}

type ChapterOWDigsiteProgramMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1           uint32 `json:"-"`
	ProduceMaterial uint32
	ProduceNumber   uint16
	CostTime        uint32
	LimitTimes      uint16

	// Objects
	ProgramName Hash
}

type ChapterOWEndlessChallengeMetaData struct {
	// Fields
	ChallengeID int32

	// Properties
	MapID                  int32
	Group                  int32
	PreChallengeID         int32
	FirstReward            int32
	Reward                 int32
	DisplayFirstReward     int32
	DisplayReward          int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	MaxScore               int32
	Difficult              int32
	Addr3                  uint32 `json:"-"`
	ChallengeHeroRecommend int32
	Addr4                  uint32 `json:"-"`
	RecommendBP            int32
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`

	// Objects
	MonsterIDList []int32
	CostMaterial  StrWithPrefix16
	ChallengeDesc StrWithPrefix16
	LockTips      StrWithPrefix16
	GroupName     StrWithPrefix16
	FailedNote    []Hash
}

type ChapterOWEntryPlotLineMetaData struct {
	// Fields
	ID int32

	// Properties
	ConditionID int32
	Addr1       uint32 `json:"-"`

	// Objects
	PlotLinePath StrWithPrefix16
}

type ChapterOWEquipmentBuffMetaData struct {
	// Fields
	ID    int32
	Level int32

	// Properties
	MapID      int32
	Type       uint16
	Cost       int32
	BuffWeight int32
	Addr1      uint32 `json:"-"`
	SuitID     int32
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`
	MaterialId int32
	Addr7      uint32 `json:"-"`

	// Objects
	IconPath    StrWithPrefix16
	Name        Hash
	Desc        Hash
	AbilityName StrWithPrefix16
	ConfigType  StrWithPrefix16
	ParamList   []StrWithPrefix16
	SourceDesc  Hash
}

type ChapterOWEquipmentPartMetaData struct {
	// Fields
	PartID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	InitialCost int32
	Addr3       uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	SlotIDList          []int32
	UINodePath          StrWithPrefix16
}

type ChapterOWEquipmentSlotMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Type  uint16
	Addr2 uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	UINodePath          StrWithPrefix16
}

type ChapterOWEventTabConfigMetaData struct {
	// Fields
	TabID uint32

	// Properties
	MapID uint32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	EventIDs      []uint32
	TabIconPath   StrWithPrefix16
	TabDescTextID Hash
}

type ChapterOWFameMetaData struct {
	// Fields
	FameLevel uint8
	MapID     int32

	// Properties
	FameNumber int32
}

type ChapterOWFortDataMetaData struct {
	// Fields
	FortID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	QuestIDList  []int32
	FortBuffList []int32
}

type ChapterOWFurnaceFormulaMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1           uint32 `json:"-"`
	ProduceMaterial int32
	ProduceNumber   int32
	FurnaceLevel    uint32
	NeedMaterial_1  int32
	NeedNumber_1    int32
	NeedMaterial_2  int32
	NeedNumber_2    int32
	NeedMaterial_3  int32
	NeedNumber_3    int32
	CostTime        int32
	MaxUseTimes     int32

	// Objects
	FormulaName Hash
}

type ChapterOWFurnitureMetaData struct {
	// Fields
	ID uint8

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	NPCHeadMarkScale float32
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`

	// Objects
	PrefabPath                StrWithPrefix16
	ConditionList             []ChapterOWFurnitureMetaData_ConditionStatePair
	NPCHeadMark               []ChapterOWFurnitureMetaData_ConditionStatePair
	NPCHeadMarkOffsetPosition []float32
	MainStoryIDList           []int32
	BranchStoryIDList         []int32
}

type ChapterOWFurnitureMetaData_ConditionStatePair struct {
	// Fields
	ConditionID int32
	StateID     int32
}

type ChapterOWHeroCardLevelMetaData struct {
	// Fields
	CardID int32
	Level  uint8

	// Properties
	NeedExp     int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	BattlePoint int32

	// Objects
	Desc        StrWithPrefix16
	AbilityName StrWithPrefix16
	ParamList   []float32
}

type ChapterOWHeroCardMetaData struct {
	// Fields
	CardID int32

	// Properties
	MapID    int32
	HeroID   int32
	MaxLevel uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Rarity   int32
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	Addr7    uint32 `json:"-"`

	// Objects
	Name                        StrWithPrefix16
	CardContentTextID           Hash
	ShowConditionList           StrWithPrefix16
	OverflowConvertMaterialList StrWithPrefix16
	SymbiosisCardsList          []int32
	CounterCardsList            []int32
	ActivationDesc              StrWithPrefix16
}

type ChapterOWHeroCardRarityMetaData struct {
	// Fields
	RarityID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	RarityBasePath   StrWithPrefix16
	IconOutLineColor StrWithPrefix16
	ProgressBarColor StrWithPrefix16
}

type ChapterOWHeroDisplayMetaData struct {
	// Fields
	ID         uint32
	HeroStatus uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	UnlockPhaseRange    ChapterOWHeroDisplayMetaData_PhaseRangeData
	HeroName            StrWithPrefix16
	HeroDesc            StrWithPrefix16
	HeroOrderDesc       StrWithPrefix16
	HeroImagePath       StrWithPrefix16
	HeroImageForDisplay StrWithPrefix16
	HeroIconPath        StrWithPrefix16
}

type ChapterOWHeroDisplayMetaData_PhaseRangeData struct {
	// Fields
	BeginPhaseID int32
	EndPhaseID   int32
}

type ChapterOWHeroLevelMetaData struct {
	// Fields
	HeroID int32
	Level  uint8

	// Properties
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	SpecialCardSlotUpgrade uint8
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`
	Addr7                  uint32 `json:"-"`
	Addr8                  uint32 `json:"-"`
	Addr9                  uint32 `json:"-"`
	Addr10                 uint32 `json:"-"`
	Addr11                 uint32 `json:"-"`
	ActiveSkillCD          float32
	Addr12                 uint32 `json:"-"`
	Addr13                 uint32 `json:"-"`
	BattlePoint            int32

	// Objects
	UnlockSlotIndex            []int32
	SpecialSlotLimitHeroList   []int32
	LevelUpCostMaterialList    StrWithPrefix16
	ActiveSkillName            StrWithPrefix16
	ActiveSkillNextLevelDesc   StrWithPrefix16
	ActiveSkillDesc            StrWithPrefix16
	AttributeDisplayDesc       StrWithPrefix16
	BuildingSkillDesc          StrWithPrefix16
	BuildingSkillNextLevelDesc StrWithPrefix16
	ActiveSkill                StrWithPrefix16
	ActiveSkillParam           []float32
	AttributeSkill             StrWithPrefix16
	AttributeSkillParam        []int32
}

type ChapterOWHeroMetaData struct {
	// Fields
	HeroID int32

	// Properties
	Type         uint8
	MapID        int32
	MaterialID   int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	Addr7        uint32 `json:"-"`
	Addr8        uint32 `json:"-"`
	Addr9        uint32 `json:"-"`
	HeroMaxLevel uint8
	Addr10       uint32 `json:"-"`
	Addr11       uint32 `json:"-"`
	Addr12       uint32 `json:"-"`
	Addr13       uint32 `json:"-"`

	// Objects
	Name                 StrWithPrefix16
	BuffName             StrWithPrefix16
	HeroImagePath        StrWithPrefix16
	HeroHeadImagePath    StrWithPrefix16
	ShadowImagePath      StrWithPrefix16
	HeroIconPath         StrWithPrefix16
	HeroCardIconPath     StrWithPrefix16
	IconColor            StrWithPrefix16
	IconLightColor       StrWithPrefix16
	SpecialCardSlotIndex []int32
	LockHeroName         StrWithPrefix16
	LockHeroDesc         StrWithPrefix16
	LockHeroComfrom      StrWithPrefix16
}

type ChapterOWHeroSpMetaData struct {
	// Fields
	HeroID int32

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Reward int32
	Addr3  uint32 `json:"-"`

	// Objects
	ShowConditionList   StrWithPrefix16
	UnlockConditionList StrWithPrefix16
	SubLineStegeList    []int32
}

type ChapterOWInteractActionMetaData struct {
	// Fields
	ID int32

	// Properties
	Type  uint8
	Addr1 uint32 `json:"-"`

	// Objects
	Param []StrWithPrefix16
}

type ChapterOWInteractStateMetaData struct {
	// Fields
	ID int32

	// Properties
	NPCID               int32
	IsShow              bool
	Interactive         bool
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	InteractRadius      float32
	InteractBelowHeight float32
	InteractAboveHeight float32
	InteractAngle       float32
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`

	// Objects
	InteractName            Hash
	InteractIcon            StrWithPrefix16
	InteractEnterActionList []int32
	InteractExitActionList  []int32
}

type ChapterOWLandMarkMetaData struct {
	// Fields
	EventID int32

	// Properties
	MapID           int32
	DisplayType     uint16
	AreaID          int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	UnlockCondition int32
	LocationID      int32

	// Objects
	Icon             StrWithPrefix16
	Title            Hash
	Desc             Hash
	DisplayCondition StrWithPrefix16
}

type ChapterOWLevelSiteMetaData struct {
	// Fields
	SiteID int32

	// Properties
	MapID   int32
	Addr1   uint32 `json:"-"`
	LevelId int32
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`
	Addr4   uint32 `json:"-"`

	// Objects
	UINodePath      StrWithPrefix16
	ShowCondition   StrWithPrefix16
	UnlockCondition StrWithPrefix16
	CloseCondition  StrWithPrefix16
}

type ChapterOWMainLineMetaData struct {
	// Fields
	ID uint32

	// Properties
	PhaseID uint32
	Type    uint8
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`
	Addr4   uint32 `json:"-"`
	Addr5   uint32 `json:"-"`
	Addr6   uint32 `json:"-"`

	// Objects
	Param               []uint32
	UnlockConditionList StrWithPrefix16
	EntryImagePath      StrWithPrefix16
	BackgroundIconPath  StrWithPrefix16
	ChapterDisplay      StrWithPrefix16
	OrderDisplay        StrWithPrefix16
}

type ChapterOWMemoryContentMetaData struct {
	// Fields
	ContentID int32

	// Properties
	ContentType int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`

	// Objects
	ContentParams []int32
	Title         Hash
	Desc          Hash
	ImagePath     StrWithPrefix16
}

type ChapterOWMemoryMetaData struct {
	// Fields
	MemoryID int32

	// Properties
	MemoryGroup int32
	MemoryType  int32
	Addr1       uint32 `json:"-"`

	// Objects
	UIPanelPrefabPath StrWithPrefix16
}

type ChapterOWMemorySiteMetaData struct {
	// Fields
	SiteID int32

	// Properties
	MemoryID  int32
	SiteType  int32
	Addr1     uint32 `json:"-"`
	ContentID int32
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	PreSiteList     []int32
	UnlockCondition StrWithPrefix16
	Title           Hash
	ImagePath       StrWithPrefix16
}

type ChapterOWMissionMetaData struct {
	// Fields
	MissionID uint32

	// Properties
	Tab      uint32
	Priority int32
}

type ChapterOWMissionTabMetaData struct {
	// Fields
	TabID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TabName    Hash
	UpdateDesc Hash
}

type ChapterOWMoonConditionMetaData struct {
	// Fields
	ID int32

	// Properties
	ConditionType uint8
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	Parameter       []int32
	StringParameter []StrWithPrefix16
}

type ChapterOWMoonTowerFloorMetaData struct {
	// Fields
	FloorConfigID uint16
	FloorOrder    uint16

	// Properties
	FirstDropMaterialID  int32
	FirstDropMaterialNum uint16
	FirstDropTerminalExp uint32
	IsSaveFloor          bool
	MonsterID            int32
	Addr1                uint32 `json:"-"`
	RecommendTalentLevel uint16
	MaxScore             uint32
	IsDifficult          bool
	Addr2                uint32 `json:"-"`

	// Objects
	Title             Hash
	MonsterNatureList []int32
}

type ChapterOWMoonTowerMetaData struct {
	// Fields
	TowerID uint16

	// Properties
	FloorConfigID     uint16
	MapID             int32
	LocationID        uint32
	DisplayCondition  int32
	Addr1             uint32 `json:"-"`
	ScoreDropConfigID uint16
	DropLimitConfigID uint16
	Addr2             uint32 `json:"-"`
	SweepUnlockFloor  uint16
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`

	// Objects
	OpenTime            RemoteTime
	RecommendTalentList []ChapterOWMoonTowerMetaData_RecommendTalent
	DescList            []Hash
	Title               Hash
	DisplayMaterialList []int32
}

type ChapterOWMoonTowerMetaData_RecommendTalent struct {
	// Fields
	TalentTag uint8
	TalentID  int32
}

type ChapterOWMoonTowerScoreDropMetaData struct {
	// Fields
	ScoreDropConfigID uint16
	Score             uint32

	// Properties
	DropMaterialID  int32
	DropMaterialNum uint32
}

type ChapterOWNPCMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1                uint32 `json:"-"`
	PrefabScale          float32
	IsNeedLookAtPosition uint8
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	Addr6                uint32 `json:"-"`

	// Objects
	PrefabPath                StrWithPrefix16
	ConditionList             []ChapterOWNPCMetaData_ConditionStatePair
	NPCHeadMark               []ChapterOWNPCMetaData_ConditionStatePair
	NPCHeadMarkOffsetPosition []float32
	MainStoryIDList           []int32
	BranchStoryIDList         []int32
}

type ChapterOWNPCMetaData_ConditionStatePair struct {
	// Fields
	ConditionID int32
	StateID     int32
}

type ChapterOWNPCPositionGroupMetaData struct {
	// Fields
	Priority int32

	// Properties
	ConditionID int32
	Addr1       uint32 `json:"-"`

	// Objects
	PositionIDList []int32
}

type ChapterOWNPCPositionMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	ConditionID int32

	// Objects
	NPCPosList []ChapterOWNPCPositionMetaData_Positions
}

type ChapterOWNPCPositionMetaData_Positions struct {
	// Fields
	AnimID int32
	NpcID  int32
	Addr1  uint32 `json:"-"`

	// Objects
	PostionName StrWithPrefix16
}

type ChapterOWPhaseMetaData struct {
	// Fields
	PhaseID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	Addr6       uint32 `json:"-"`
	PhaseNumber int32

	// Objects
	UnlockConditionList     StrWithPrefix16
	ProgressText            Hash
	CoverImagePath          StrWithPrefix16
	MainStoryEntryImagePath StrWithPrefix16
	ActivityEntryImagePath  StrWithPrefix16
	BGMStateName            StrWithPrefix16
}

type ChapterOWPlotSiteMetaData struct {
	// Fields
	SiteID int32

	// Properties
	MapID  int32
	Addr1  uint32 `json:"-"`
	PlotID int32
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`

	// Objects
	UINodePath      StrWithPrefix16
	SiteTitle       Hash
	SiteSubTitle    Hash
	WindowImgPath   StrWithPrefix16
	WindowDesc      Hash
	WindowBtnText   Hash
	ShowCondition   StrWithPrefix16
	UnlockCondition StrWithPrefix16
	PassedCondition StrWithPrefix16
	CloseCondition  StrWithPrefix16
}

type ChapterOWQTECircleMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1      uint32 `json:"-"`
	StartAngle uint16
	Addr2      uint32 `json:"-"`

	// Objects
	CircleRangeIDList []uint16
	CircleModeList    []ChapterOWQTECircleMetaData_CircleMode
}

type ChapterOWQTECircleMetaData_CircleMode struct {
	// Fields
	Angle int16
	Speed float32
}

type ChapterOWQTECircleRangeMetaData struct {
	// Fields
	ID uint16

	// Properties
	MinAngle         float32
	MaxAngle         float32
	MinPerfectAngle  float32
	MaxPerfectAngle  float32
	BasicRewardNum   int32
	PerfectRewardNum int32
}

type ChapterOWQTELastRewardMetaData struct {
	// Fields
	ID uint16

	// Properties
	PullTime      float32
	QTEReduceTime float32
	EnergyCost    float32
	MinInterval   float32
}

type ChapterOWQTEMapGroupMetaData struct {
	// Fields
	ID uint16

	// Properties
	GroupType uint8
	Addr1     uint32 `json:"-"`

	// Objects
	MapList []ChapterOWQTEMapGroupMetaData_Map
}

type ChapterOWQTEMapGroupMetaData_Map struct {
	// Fields
	ID     uint16
	Weight uint16
}

type ChapterOWQTEMapMetaData struct {
	// Fields
	ID uint16

	// Properties
	MapType              uint8
	MapDepth             uint16
	NormalDropMaterialID int32
	MaxNormalDropNum     int32
	EnergyCost           uint16
	Addr1                uint32 `json:"-"`
	LastRewardID         uint16
	LastRewardMaterialID int32
	LastRewardNum        int32
	Addr2                uint32 `json:"-"`
	RecommendEnergy      uint16

	// Objects
	CircleIDList     []ChapterOWQTEMapMetaData_Circle
	ChallengeTextmap Hash
}

type ChapterOWQTEMapMetaData_Circle struct {
	// Fields
	Depth uint16
	ID    uint16
}

type ChapterOWQuestDataMetaData struct {
	// Fields
	ID uint32

	// Properties
	IsSpecialQuest  bool
	QuestLevel      uint32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	QuestLocationID int32
	Addr3           uint32 `json:"-"`

	// Objects
	QuestConsigner StrWithPrefix16
	QuestInfo      StrWithPrefix16
	RarityTips     StrWithPrefix16
}

type ChapterOWQuestLevelMetaData struct {
	// Fields
	Level uint32

	// Properties
	MapID                int32
	SpecialQuestNum      uint32
	MaxDailyAcceptTimes  uint32
	MaxSlotNum           uint32
	MaxDailyRefreshTimes uint32
}

type ChapterOWRandomStageMetaData struct {
	// Fields
	LevelID int32

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	MainDropCardDisplay int32
	ShowWeight          float32
	Addr3               uint32 `json:"-"`

	// Objects
	HardLevelPointRange []int32
	DisplayIcon         StrWithPrefix16
	MonsterIDList       []int32
}

type ChapterOWRangePlotGroupMetaData struct {
	// Fields
	PlotID uint32

	// Properties
	PlotGroupID uint32
	RangeType   uint32
	Addr1       uint32 `json:"-"`

	// Objects
	Range ChapterOWRangePlotGroupMetaData_PlotGroupRangeData
}

type ChapterOWRangePlotGroupMetaData_PlotGroupRangeData struct {
	// Fields
	Max uint32
	Min uint32
}

type ChapterOWRefiningLevelUPMetaData struct {
	// Fields
	Level uint32

	// Properties
	Addr1         uint32 `json:"-"`
	FieldNumber   int32
	ReducePercent int32

	// Objects
	FormulaList []int32
}

type ChapterOWRelicsEventMetaData struct {
	// Fields
	ID uint32

	// Properties
	Reward int32
}

type ChapterOWRelicsMetaData struct {
	// Fields
	ID uint16

	// Properties
	MapID                    int32
	RelicsMaterialID         int32
	RelicsFragmentMaterialID int32
	RelicsChangeNumber       uint16
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	ShowCondition            int32
	RelicsReward             int32
	Addr3                    uint32 `json:"-"`

	// Objects
	SiteIDList []uint16
	UnlockTime RemoteTime
	Title      Hash
}

type ChapterOWRelicsSiteMetaData struct {
	// Fields
	ID uint16

	// Properties
	IsSeniorSite bool
	SectionID    int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	LocationID   int32
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`

	// Objects
	UINodeName             StrWithPrefix16
	EventIDList            []uint32
	AntiGravityLevelIDList []int32
	Title                  Hash
	SiteDesc               Hash
}

type ChapterOWRewardMetaData struct {
	// Fields
	MapID       int32
	CountNumber int32

	// Properties
	LimitRewardID    int32
	ResidentRewardID int32
	Addr1            uint32 `json:"-"`
	IsPhasesReward   bool
	IsKeyReward      bool
	Addr2            uint32 `json:"-"`
	IsAvatarReward   bool
	Addr3            uint32 `json:"-"`

	// Objects
	RewardBeginTime RemoteTime
	KeyRewardName   StrWithPrefix16
	KeyRewardImage  StrWithPrefix16
}

type ChapterOWShopGoodsMetaData struct {
	// Fields
	GoodsID int32

	// Properties
	RewardID   int32
	ShopID     int32
	CostNumber int32
	NeedFame   int32
}

type ChapterOWShopLinkMetaData struct {
	// Fields
	ShopID int32

	// Properties
	Type          int32
	MapID         int32
	Addr1         uint32 `json:"-"`
	ChangePhase   int32
	ChangePhaseV2 int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`

	// Objects
	SellerImagePath  StrWithPrefix16
	TabDesc          StrWithPrefix16
	CurrencyShowList []int32
}

type ChapterOWShopMetaData struct {
	// Fields
	ShopID int32

	// Properties
	MapID   int32
	EventID int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	Title               Hash
}

type ChapterOWSpecialStoryMetaData struct {
	// Fields
	ID int32

	// Properties
	MapID      int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	CanExplore bool

	// Objects
	SubTitle     StrWithPrefix16
	Introduction StrWithPrefix16
}

type ChapterOWTalentBuffLevelMetaData struct {
	// Fields
	ID uint8

	// Properties
	LevelRange     uint8
	MapID          int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	RobotPropState uint8
	Addr4          uint32 `json:"-"`
	RobotScale     float32
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`

	// Objects
	DisplayDesc         StrWithPrefix16
	AbilityName         StrWithPrefix16
	AbilityParamList    []float32
	RobotPrefabPath     StrWithPrefix16
	RobotPosition       []float32
	RobotRotation       []float32
	RobotImgPath        StrWithPrefix16
	RobotUpgradeImgPath StrWithPrefix16
	RobotOutlineBGPath  StrWithPrefix16
}

type ChapterOWTalentDataMetaData struct {
	// Fields
	TalentID int32

	// Properties
	MapID             int32
	TalentType        uint8
	IsDefaultActivate uint8
	AutoActivate      int32
	UniqueTag         uint8
	MaxLevel          uint8
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`

	// Objects
	RelateTalentIDList  []int32
	UIChangePath        StrWithPrefix16
	UIPointPath         StrWithPrefix16
	UILinePath          []StrWithPrefix16
	TalentName          Hash
	TalentDesc          Hash
	IconPath            StrWithPrefix16
	CGID                []int32
	UnlockConditionList StrWithPrefix16
}

type ChapterOWTalentLevelMetaData struct {
	// Fields
	TalentID    int32
	TalentLevel uint8

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	NeedWorkShopLevel uint8
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`

	// Objects
	TalentDesc              Hash
	NextLevelDesc           Hash
	PreTalent               []ChapterOWTalentLevelMetaData_PreTalentRequirement
	CostMaterialList        []ChapterOWTalentLevelMetaData_CostMaterial
	AbilityName             StrWithPrefix16
	AbilityParamList        []float32
	AbilityParamDisplayList []float32
}

type ChapterOWTalentLevelMetaData_CostMaterial struct {
	// Fields
	MaterialID int32
	Num        int32
}

type ChapterOWTalentLevelMetaData_PreTalentRequirement struct {
	// Fields
	TalentLevel uint8
	TalentID    int32
}

type ChapterOWTerminalLevelMetaData struct {
	// Fields
	MapID int32
	Level uint8

	// Properties
	LevelExp               int32
	IsBreak                uint8
	BattleTalentTotalLevel uint8
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	HardLevel              uint8
	Addr4                  uint32 `json:"-"`
	Addr5                  uint32 `json:"-"`

	// Objects
	ConsumeMaterialList []ChapterOWTerminalLevelMetaData_ConsumeMaterial
	PreStoryList        []int32
	PlotlinePath        StrWithPrefix16
	PrefabPath          StrWithPrefix16
	IconPath            StrWithPrefix16
}

type ChapterOWTerminalLevelMetaData_ConsumeMaterial struct {
	// Fields
	MatrialID int32
	Num       int32
}

type ChapterOWTerminalTipsMetaData struct {
	// Fields
	ID int32

	// Properties
	MapID       int32
	Type        uint8
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	ConditionID int32
	Weight      int32

	// Objects
	Title Hash
	Desc  Hash
}

type ChapterOWTipsMetaData struct {
	// Fields
	ID int32

	// Properties
	MapID      int32
	Type       int16
	PopTiming  int16
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	TutorialID int32
	MaterialID int32
	Addr5      uint32 `json:"-"`

	// Objects
	UnlockConditionList StrWithPrefix16
	Icon                StrWithPrefix16
	Title               Hash
	Desc                Hash
	Params              []int32
}

type ChapterRpgAvatarRewardLineMetaData struct {
	// Fields
	ID int32

	// Properties
	StageID       int32
	RequiredScore int32
	RewardID      int32
}

type ChapterRpgBuffLevelMetaData struct {
	// Fields
	ID int32

	// Properties
	LevelRange int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	AbilityName      StrWithPrefix16
	AbilityParamList []StrWithPrefix16
}

type ChapterStageCompensationMetaData struct {
	// Fields
	StageID int32

	// Properties
	FinishReward   int32
	UnfinishReward int32
}

type ChapterSwitchAnimConfigMetaData struct {
	// Fields
	ConfigID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	ChapterImage         StrWithPrefix16
	ChapterBgImage       StrWithPrefix16
	ChapterTitleTextID   Hash
	ChapterContentTextID Hash
}

type ChatGroupIconMetaData struct {
	// Fields
	IconID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type ChatLobbyActionMetaData struct {
	// Fields
	ID int16

	// Properties
	Addr1       uint32 `json:"-"`
	Type        uint8
	ActionIndex int16
	Addr2       uint32 `json:"-"`
	Loop        bool

	// Objects
	Name Hash
	Icon StrWithPrefix16
}

type ChatLobbyActivityMetaData struct {
	// Fields
	ActivityId int32

	// Properties
	Type int32
	Para int32
}

type ChatLobbyActivityNoticeMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`
	Addr8 uint32 `json:"-"`
	Addr9 uint32 `json:"-"`

	// Objects
	OpeningShowText  Hash
	ActivityIcon     StrWithPrefix16
	ActivityNameText Hash
	ActivityTypeText Hash
	LevelText        Hash
	TimeText         Hash
	MainRewardList   []int32
	ActivityDescText Hash
	RewardDetailList []int32
}

type ChatLobbyActivityScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	ActivityId int32
}

type ChatLobbyAnnouncementMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	AnnouncementType int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	ShowTimes        int16
	ShowSeconds      int16
	Addr4            uint32 `json:"-"`

	// Objects
	BeginTime StrWithPrefix16
	Title     Hash
	Desc      Hash
	EndTime   StrWithPrefix16
}

type ChatLobbyBeastMetaData struct {
	// Fields
	BeastID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BeastModelPath    StrWithPrefix16
	BeastMaterialPath StrWithPrefix16
	Slogan            StrWithPrefix16
	HeadPicPath       StrWithPrefix16
}

type ChatLobbyBoxActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1                uint32 `json:"-"`
	PlayerDailyOpenLimit int32

	// Objects
	BoxSeriesID []int32
}

type ChatLobbyBoxSeriesMetaData struct {
	// Fields
	SeriesID int32

	// Properties
	BoxObject      int32
	BoxTriggerType int32
	Addr1          uint32 `json:"-"`

	// Objects
	Para02 Hash
}

type ChatLobbyDishLimitedRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	BeginTime              StrWithPrefix16
	EndTime                StrWithPrefix16
	DishDailyRewardDisplay []int32
}

type ChatLobbyDishMetaData struct {
	// Fields
	ActivityID uint16

	// Properties
	Addr1             uint32 `json:"-"`
	RewardGetLimit    uint8
	RewardGetColdDown uint16

	// Objects
	RewardIDList []ChatLobbyDishMetaData_RewardTurn
}

type ChatLobbyDishMetaData_RewardTurn struct {
	// Fields
	RewardID1st int32
	RewardID2nd int32
}

type ChatLobbyDishRateMetaData struct {
	// Fields
	ID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Rate  uint8

	// Objects
	TotalStartTime StrWithPrefix16
	TotalEndTime   StrWithPrefix16
}

type ChatLobbyFightExpressionMetaData struct {
	// Fields
	EmojiFaceID int32
}

type ChatLobbyFishActivityMetaData struct {
	// Fields
	ActivityId int32

	// Properties
	FishCurrencyId int32
	Addr1          uint32 `json:"-"`

	// Objects
	FishSpotList []int32
}

type ChatLobbyFishMetaData struct {
	// Fields
	FishID int32

	// Properties
	FishGroup      uint8
	Weight         uint8
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	IsRare         uint8
	FishType       uint8
	AddCurrencyID  int32
	AddCurrencyNum int32

	// Objects
	FishPicID    StrWithPrefix16
	Textmap_Name Hash
	Textmap_Desc Hash
}

type ChatLobbyFishShowMetaData struct {
	// Fields
	FishId int32

	// Properties
	FishType uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	FishPic  StrWithPrefix16
	FishName Hash
	FishDesc Hash
}

type ChatLobbyItemMetaData struct {
	// Fields
	ID int32

	// Properties
	Type                        uint8
	IsBroadcastUsingInformation bool
	Addr1                       uint32 `json:"-"`
	IsUsed                      uint8
	IsShowUseTips               bool
	TargetType                  uint8
	Addr2                       uint32 `json:"-"`

	// Objects
	Effect  StrWithPrefix16
	UseInfo Hash
}

type ChatLobbyKillingStreakMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Name         Hash
	Icon         StrWithPrefix16
	AudioPattern StrWithPrefix16
}

type ChatLobbyLanternDifficultyMetaData struct {
	// Fields
	LanternDifficulty int32

	// Properties
	LanternSignSpeed    float32
	LanternFailedScale  float32
	LanternGoodScale    float32
	LanternPerfectScale float32
}

type ChatLobbyLayoutMetaData struct {
	// Fields
	ID int32

	// Properties
	RoomID     int32
	ObjectType uint8
	ObjectID   int32
	X          float32
	Y          float32
	Z          float32
	Angle      float32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	StartTime StrWithPrefix16
	EndTime   StrWithPrefix16
}

type ChatLobbyMainPageNoticeMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	NotificationIcon StrWithPrefix16
	NotificationText Hash
}

type ChatLobbyMissionMetaData struct {
	// Fields
	MissionID int32

	// Properties
	NavigateX      float32
	NavigateY      float32
	NavigateZ      float32
	NavigateRadius float32
}

type ChatLobbyNPCMetaData struct {
	// Fields
	ID int32

	// Properties
	QAvatarID               int32
	Addr1                   uint32 `json:"-"`
	NameOffset              float32
	IsNameShown             bool
	Addr2                   uint32 `json:"-"`
	Interactable            bool
	Addr3                   uint32 `json:"-"`
	IdleIntervalMin         uint8
	IdleIntervalMax         uint8
	DialogOffset            float32
	Addr4                   uint32 `json:"-"`
	Addr5                   uint32 `json:"-"`
	BubbleDialogIntervalMin uint8
	BubbleDialogIntervalMax uint8

	// Objects
	Name             Hash
	HeadIcon         StrWithPrefix16
	IdleList         []ChatLobbyNPCMetaData_Idle
	NormalDialogList []ChatLobbyNPCMetaData_Dialog
	BubbleDialogList []ChatLobbyNPCMetaData_Dialog
}

type ChatLobbyNPCMetaData_Idle struct {
	// Fields
	Weight float32
	Addr1  uint32 `json:"-"`

	// Objects
	Path StrWithPrefix16
}

type ChatLobbyNPCMetaData_Dialog struct {
	// Fields
	Duration float32
	Weight   float32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	Text  Hash
	Voice StrWithPrefix16
}

type ChatLobbyNPCSystemFunctionMetaData struct {
	// Fields
	NPCID       int32
	UIPostionID int32
	Type        int32

	// Properties
	LinkType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	LinkParams   []int32
	LinkParamStr StrWithPrefix16
}

type ChatLobbyObjectMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1                  uint32 `json:"-"`
	NameHight              float32
	IsNameShown            bool
	Addr2                  uint32 `json:"-"`
	InteractiveRange       float32
	InteractiveButtonRange float32
	Addr3                  uint32 `json:"-"`
	NoticeEffect           int32
	NoticeEffectRange      float32
	Addr4                  uint32 `json:"-"`

	// Objects
	Name                  Hash
	AssetPath             StrWithPrefix16
	InteractiveButtonIcon StrWithPrefix16
	NoticeEffectPath      StrWithPrefix16
}

type ChatLobbyPrayMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	RewardID int32

	// Objects
	PrayMainTitle Hash
	PrayText1     Hash
	PrayText2     Hash
}

type ChatLobbyQuestionActivityMetaData struct {
	// Fields
	ActivityId uint16

	// Properties
	RightBuff            uint16
	WrongBuff            uint16
	QuestionActivityType uint16
}

type ChatLobbyQuestionMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Question_Text StrWithPrefix16
	AnswerList    []ChatLobbyQuestionMetaData_AnswerItem
}

type ChatLobbyQuestionMetaData_AnswerItem struct {
	// Fields
	ID    int32
	Addr1 uint32 `json:"-"`

	// Objects
	Text StrWithPrefix16
}

type ChatLobbyRoomMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1                uint32 `json:"-"`
	Type                 uint8
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	PlayerRotation       float32
	CameraRotation       float32
	IsCameraLocked       bool
	CameraRotationReset  bool
	CameraRotationMin    float32
	CameraRotationMax    float32
	Addr6                uint32 `json:"-"`
	Addr7                uint32 `json:"-"`
	ChatLobbySupportJump int32
	Addr8                uint32 `json:"-"`

	// Objects
	Name                Hash
	LeftBoardPageList   []int32
	ShowMissionList     []int32
	QAvatarList         []int32
	PlayerSpawnPoint    []float32
	StagePath           StrWithPrefix16
	Weather             StrWithPrefix16
	ChatLobbyMainUIPath StrWithPrefix16
}

type ChatLobbySceneItemMetaData struct {
	// Fields
	SceneItemID int32

	// Properties
	Type          int32
	Addr1         uint32 `json:"-"`
	InteractRange float32
	Addr2         uint32 `json:"-"`
	IsNameShow    bool
	Addr3         uint32 `json:"-"`

	// Objects
	DestroyEffectPattern StrWithPrefix16
	SceneItemNameText    Hash
	SceneItemModel       StrWithPrefix16
}

type ChatLobbySkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	CDTime        float32
	CostBulletNum int32
}

type ChatLobbyStanceMetaData struct {
	// Fields
	StanceID int32

	// Properties
	Addr1                      uint32 `json:"-"`
	Addr2                      uint32 `json:"-"`
	StanceOccupiedAnnouncement int32
	ProcessBarOffset           float32
	StanceRange                float32

	// Objects
	StanceName  Hash
	StanceModel StrWithPrefix16
}

type ChatLobbyTreasureMetaData struct {
	// Fields
	TreasureID int32

	// Properties
	Addr1         uint32 `json:"-"`
	ProcessOffset float32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	OpenNum       int32
	Addr4         uint32 `json:"-"`

	// Objects
	TreasureNameText    Hash
	TreasureModel       StrWithPrefix16
	CostMaterialList    []ChatLobbyTreasureMetaData_Cost
	DropDisplayItemList []ChatLobbyTreasureMetaData_DisplayItem
}

type ChatLobbyTreasureMetaData_Cost struct {
	// Fields
	Count      int32
	MaterialID int32
}

type ChatLobbyTreasureMetaData_DisplayItem struct {
	// Fields
	Count  int32
	ItemID int32
}

type ChatLobbyUsableItemMetaData struct {
	// Fields
	UsableItemID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	UsableItemIcon StrWithPrefix16
}

type ChatLobbyVoiceMetaData struct {
	// Fields
	ID int16

	// Properties
	Addr1    uint32 `json:"-"`
	RoleID   int32
	FromHour uint8
	ToHour   uint8
	Addr2    uint32 `json:"-"`
	Weight   float32

	// Objects
	AudioID      StrWithPrefix16
	AudioPattern StrWithPrefix16
}

type ChatLobbyWayPointMetaData struct {
	// Fields
	Id int32

	// Properties
	Width float32
	Depth float32
	Addr1 uint32 `json:"-"`

	// Objects
	Poslist []ChatLobbyWayPointMetaData_FishWayPoint
}

type ChatLobbyWayPointMetaData_FishWayPoint struct {
	// Fields
	X float32
	Z float32
}

type ChatMessageContentMetaData struct {
	// Fields
	ChatMessageID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ChatContent StrWithPrefix16
}

type ChatMessageDataMetaData struct {
	// Fields
	ChatID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ChatMessageList []int32
}

type ChatReportMetaData struct {
	// Fields
	ReportType int32

	// Properties
	SortWeight int32
	Addr1      uint32 `json:"-"`

	// Objects
	ReportTypeName Hash
}

type CityEventPhotoMetaData struct {
	// Fields
	PhotoID int32

	// Properties
	PlotID     int32
	IsDefault  bool
	PhotoType  int32
	ActivityID int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Type       int32
	Role       int32
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`

	// Objects
	PlotName   Hash
	UnlockTips Hash
	RolePic    StrWithPrefix16
	Face       StrWithPrefix16
	BackGround StrWithPrefix16
}

type ClickDialogBGCloseBlackListMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	ContextName StrWithPrefix16
	Hierachies  []StrWithPrefix16
}

type ClientLogDataMetaData struct {
	// Fields
	ContextID uint32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	NeedEntrance bool

	// Objects
	ContextName StrWithPrefix16
	FeatureKey  StrWithPrefix16
}

type CollaborationScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	CollaborationStigmataIDList []int32
	CollaborationWeaponIDList   []int32
	AvatarIDList                []int32
	RoleIDList                  []int32
	DressIDList                 []int32
	MedalIDList                 []int32
}

type CollaborationWeaponMetaData struct {
	// Fields
	CollaborationWeaponID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Priority int32

	// Objects
	ImagePath StrWithPrefix16
}

type CollectionDialogueMetaData struct {
	// Fields
	CollectionDialogueID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	Name      Hash
	NameBg    StrWithPrefix16
	BankName  StrWithPrefix16
	AudioName StrWithPrefix16
}

type CollectionEventMetaData struct {
	// Fields
	CollectionEventID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	PlotID                int32
	HideContentWhenLocked bool
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`

	// Objects
	Name       Hash
	NameBg     StrWithPrefix16
	LockedText Hash
	Desc       Hash
}

type CollectionFileMetaData struct {
	// Fields
	CollectionFileID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	HideContentWhenLocked bool
	Addr4                 uint32 `json:"-"`

	// Objects
	Name       Hash
	NameBg     StrWithPrefix16
	Desc       Hash
	LockedText Hash
}

type CollectionItemDataMetaData struct {
	// Fields
	CollectionItemID int32

	// Properties
	MaterialID int32
}

type CollectionMonsterDetailInfoMetaData struct {
	// Fields
	InfoID int32

	// Properties
	Addr1   uint32 `json:"-"`
	Mission int32
	Addr2   uint32 `json:"-"`

	// Objects
	Desc       Hash
	LockedDesc Hash
}

type CollectionMonsterMetaData struct {
	// Fields
	CollectionMonsterID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	AllowNormal           bool
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	Scale                 float32
	Addr7                 uint32 `json:"-"`
	Addr8                 uint32 `json:"-"`
	Health                int32
	Attack                int32
	Addr9                 uint32 `json:"-"`
	Addr10                uint32 `json:"-"`
	HideContentWhenLocked bool
	Addr11                uint32 `json:"-"`

	// Objects
	UniqueMonsterID []int32
	Name            Hash
	TypeName        Hash
	Desc            Hash
	Icon            StrWithPrefix16
	PrefabPath      StrWithPrefix16
	Position        []float32
	Rotation        []float32
	AbilitiesList   []int32
	DetailInfoIDs   []int32
	LockedText      Hash
}

type CollectionPictureMetaData struct {
	// Fields
	CollectionPictureID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	HideContentWhenLocked bool
	Addr4                 uint32 `json:"-"`

	// Objects
	Name       Hash
	Icon       StrWithPrefix16
	ShowImg    StrWithPrefix16
	LockedText Hash
}

type CollectionTypeMetaData struct {
	// Fields
	SysType        uint8
	CollectionType uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Title Hash
	Desc  Hash
}

type CollectionVisualNovelDataMetaData struct {
	// Fields
	VisualNovelID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	LuaPath StrWithPrefix16
}

type CompensationOfDormitoryMetaData struct {
	// Fields
	FacilityType int32
	Level        int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	FacilityName Hash
	ItemList     []CompensationOfDormitoryMetaData_CompensationOfDormitoryRewardMetaData
}

type CompensationOfDormitoryMetaData_CompensationOfDormitoryRewardMetaData struct {
	// Fields
	Num      int32
	RewardID int32
}

type CompensationOfElfBreakMetaData struct {
	// Fields
	ElfID     int32
	BreakStep int32

	// Properties
	MaxLevel int32
	Addr1    uint32 `json:"-"`

	// Objects
	ItemList []CompensationOfElfBreakMetaData_Item
}

type CompensationOfElfBreakMetaData_Item struct {
	// Fields
	ID     int32
	Number int32
}

type CompensationOfElfSlotUnlockMetaData struct {
	// Fields
	ElfID   int32
	SlotNum int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ItemList []CompensationOfElfSlotUnlockMetaData_Item
}

type CompensationOfElfSlotUnlockMetaData_Item struct {
	// Fields
	ID     int32
	Number int32
}

type CompensationOfElfTalentMetaData struct {
	// Fields
	TalentID    int32
	TalentLevel int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ItemList []CompensationOfElfTalentMetaData_Item
}

type CompensationOfElfTalentMetaData_Item struct {
	// Fields
	ID     int32
	Number int32
}

type CompensationOfExtendGradeMetaData struct {
	// Fields
	CabinType   int32
	ExtendGrade int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ItemList []CompensationOfExtendGradeMetaData_CompensationOfExtendGradeRewardMetaData
}

type CompensationOfExtendGradeMetaData_CompensationOfExtendGradeRewardMetaData struct {
	// Fields
	Num      int32
	RewardID int32
}

type CompensationOfIslandMetaData struct {
	// Fields
	CabinType int32
	Level     int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	CabinName Hash
	ItemList  []CompensationOfIslandMetaData_CompensationOfIslandRewardMetaData
}

type CompensationOfIslandMetaData_CompensationOfIslandRewardMetaData struct {
	// Fields
	Num      int32
	RewardID int32
}

type ConstValueMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Name  Hash
	Type  StrWithPrefix16
	Value StrWithPrefix16
}

type CoupleTowerRewardMetaData struct {
	// Fields
	ActivityID int32
	MinFloor   int32

	// Properties
	MaxFloor int32
	RewardID int32
}

type CoupleTowerScoreMetaData struct {
	// Fields
	StageID int32
	Floor   int32

	// Properties
	ScoreType int32
	Score     int32
}

type CoupleTowerStageMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	StageIDList []int32
}

type CouponMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Addr1           uint32 `json:"-"`
	CouponNeedPrice int32
	CouponValue     int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	SortId          int32
	IsShow          int32

	// Objects
	CouponTypeList []int32
	BeginTime      StrWithPrefix16
	EndTime        StrWithPrefix16
}

type CreditRankMetaData struct {
	// Fields
	RankID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	ShowWarningHint int32
	Addr3           uint32 `json:"-"`
	MinScore        int32
	MaxScore        int32

	// Objects
	RankName    StrWithPrefix16
	RankColor   StrWithPrefix16
	WarningHint StrWithPrefix16
}

type CreditRegeditMetaData struct {
	// Fields
	AccountID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	TypeList  []int32
	Title     StrWithPrefix16
	RuleTitle StrWithPrefix16
	RuleInfo  StrWithPrefix16
}

type CrisisModeActivityConfigMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	PreLevelID int32
	Addr1      uint32 `json:"-"`

	// Objects
	AvatarPool []int32
}

type CrisisModeStageConfigMetaData struct {
	// Fields
	StageID int32

	// Properties
	Reward int32
}

type CrisisModeStageGroupConfigMetaData struct {
	// Fields
	StageGroupID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TagList []int32
}

type CurrencyIconMetaData struct {
	// Fields
	CurrencyID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type CurrencyQuickExchangeMetaData struct {
	// Fields
	TargetID int32

	// Properties
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	Addr3                    uint32 `json:"-"`
	Addr4                    uint32 `json:"-"`
	Addr5                    uint32 `json:"-"`
	Addr6                    uint32 `json:"-"`
	Addr7                    uint32 `json:"-"`
	Addr8                    uint32 `json:"-"`
	EnableDiamondExchange    int32
	EnablePurpleJadeExchange int32
	EnableMCoinExchange      int32
	MaxBuyNum                int32
	UnitCost                 int32

	// Objects
	DialogTitle        Hash
	ExchangeMethodDesc Hash
	CurrentNumberDesc  Hash
	ExchangeNumberDesc Hash
	ObtainNumberDesc   Hash
	CancelDesc         Hash
	OkDesc             Hash
	TipsDesc           Hash
}

type CustomGachaExFragmentMetaData struct {
	// Fields
	EventID uint16

	// Properties
	MaterialAmountRequest int32
	FragmentAmountByOnce  int32
	BonusGachaTimes       int32
	BonusRate             int32
}

type CustomHeadPageMetaData struct {
	// Fields
	Page int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	PageDec StrWithPrefix16
}

type CycleMissionManagementMetaData struct {
	// Fields
	MissionID int32

	// Properties
	HaveLink bool
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	PreDisplayCycleIDList   []int32
	RewardUnlockCycleIDList []int32
	RewardUnlockTime        StrWithPrefix16
}

type DailyMissionMaterialIconMetaData struct {
	// Fields
	MaterialID   int32
	MaterialType int16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type DailyMPStageDataMetaData struct {
	// Fields
	Level int32

	// Properties
	Addr1         uint32 `json:"-"`
	RewardDisplay int32

	// Objects
	CoverBGPath StrWithPrefix16
}

type DailyRechargeRewardGroupMetaData struct {
	// Fields
	GroupID  int32
	Progress int32

	// Properties
	RewardID int32
	Addr1    uint32 `json:"-"`

	// Objects
	GroupDetail Hash
}

type DailyRecommendMeta struct {
	// Fields
	PanelID uint16

	// Properties
	Addr1                 uint32 `json:"-"`
	SortID                uint8
	Addr2                 uint32 `json:"-"`
	DefaultMissionDisplay int32
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	LinkType              uint16
	Addr7                 uint32 `json:"-"`
	Addr8                 uint32 `json:"-"`
	Addr9                 uint32 `json:"-"`
	Addr10                uint32 `json:"-"`
	Addr11                uint32 `json:"-"`
	Addr12                uint32 `json:"-"`
	Addr13                uint32 `json:"-"`
	MentoringRelationship uint8
	IsRetainGoBtn         bool
	ShowEndCountdown      bool

	// Objects
	TextMapTitle          StrWithPrefix16
	HangMissionList       []int32
	InfoButtonIcon        StrWithPrefix16
	InfoButtonFigure      StrWithPrefix16
	TextMapTag            StrWithPrefix16
	TagColor              StrWithPrefix16
	LinkParams            []int32
	LinkParamStr          StrWithPrefix16
	TextMapActivityTime   StrWithPrefix16
	TextMapActivityType   StrWithPrefix16
	TextMapLevelLimit     StrWithPrefix16
	DisplayRewardItemList []int32
	DisplayMaterialList   []StrWithPrefix16
}

type DanmakuMetaData struct {
	// Fields
	DanmakuID uint16

	// Properties
	MinLevel uint16
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	BeginTime  StrWithPrefix16
	EndTime    StrWithPrefix16
	SlotIDList []uint16
}

type DanmakuSlotMetaData struct {
	// Fields
	SlotID uint16

	// Properties
	ChannelType uint8
	ChannelID   uint16
}

type DeviceFPSLimitMetaData struct {
	// Fields
	TargetFPS int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	DisplayTextID           Hash
	DeviceWhiteList         []StrWithPrefix16
	GraphicsDeviceWhiteList []StrWithPrefix16
	DeviceBlackList         []StrWithPrefix16
}

type DialogImageDataMetaData struct {
	// Fields
	ImageId int32

	// Properties
	Addr1     uint32 `json:"-"`
	X         float32
	Y         float32
	Width     float32
	Height    float32
	Rotationy float32
	Addr2     uint32 `json:"-"`

	// Objects
	Path  StrWithPrefix16
	Title Hash
}

type DialogMetaData struct {
	// Fields
	DialogID int32
	Addr1    uint32 `json:"-"`

	// Properties
	Addr2      uint32 `json:"-"`
	DialogType uint8
	AvatarID   uint16
	ScreenSide uint8
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`
	Addr7      uint32 `json:"-"`

	// Objects
	PostDialogIdList []int32
	PreDialogIDList  []int32
	Source           StrWithPrefix16
	Content          []DialogMetaData_PlotChatNode
	AudioID          StrWithPrefix16
	Emotion          StrWithPrefix16
	LipMotion        StrWithPrefix16
}

type DialogMetaData_PlotChatNode struct {
	// Fields
	ChatDuration float32
	Addr1        uint32 `json:"-"`

	// Objects
	ChatContent Hash
}

type DiceyDungeonActivityMetaData struct {
	// Fields
	DiceyDungeonActivityID int32

	// Properties
	Addr1                  uint32 `json:"-"`
	DiceyDungeonMedalID    int32
	DailyDungeonUnlockSite int32

	// Objects
	WeaponGachaMaterial DiceyDungeonActivityMetaData_MaterialInfo
}

type DiceyDungeonActivityMetaData_MaterialInfo struct {
	// Fields
	ID  int32
	Num int32
}

type DiceyDungeonBubbleMetaData struct {
	// Fields
	RoleID   int32
	BubbleID uint8

	// Properties
	TriggerCondition uint8
	Addr1            uint32 `json:"-"`
	TriggerTimes     uint8
	TriggerPR        uint8
	Weight           uint16
	Addr2            uint32 `json:"-"`

	// Objects
	ConditionParameter []int32
	DialogTextList     []Hash
}

type DiceyDungeonBuffMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	IsHide bool
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`

	// Objects
	BuffKey  StrWithPrefix16
	BuffIcon StrWithPrefix16
	BuffName Hash
	BuffDesc Hash
}

type DiceyDungeonDailyScheduleMetaData struct {
	// Fields
	DungeonDailyScheduleID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	DungeonID             int32
	MaxOrnamentNum        uint8
	MaxPassiveOrnamentNum uint8
	Addr3                 uint32 `json:"-"`

	// Objects
	StartTime                    RemoteTime
	EndTime                      RemoteTime
	DailyDungeonMaterialNumLimit []DiceyDungeonDailyScheduleMetaData_MatLimitInfo
}

type DiceyDungeonDailyScheduleMetaData_MatLimitInfo struct {
	// Fields
	ID       int32
	LimitNum int32
}

type DiceyDungeonDailyScoreMetaData struct {
	// Fields
	Index int32

	// Properties
	DungeonID            int32
	DailyDungeonMinScore int32
	Addr1                uint32 `json:"-"`

	// Objects
	DailyDungeonMaterial []DiceyDungeonDailyScoreMetaData_RewardInfo
}

type DiceyDungeonDailyScoreMetaData_RewardInfo struct {
	// Fields
	ID  int32
	Num int32
}

type DiceyDungeonDataMetaData struct {
	// Fields
	DungeonID int32

	// Properties
	Addr1                      uint32 `json:"-"`
	DungeonType                uint8
	BeginDungeonRoomID         int32
	Addr2                      uint32 `json:"-"`
	Addr3                      uint32 `json:"-"`
	Addr4                      uint32 `json:"-"`
	Addr5                      uint32 `json:"-"`
	Addr6                      uint32 `json:"-"`
	OrnamentPoolID             int32
	FristDropRewardDisplayList int32
	RewardDisplayList          int32
	Addr7                      uint32 `json:"-"`
	Addr8                      uint32 `json:"-"`
	RecoveryRoomHP             float32
	CheckEventHP               float32
	OrnamentChoiceHP           float32
	Addr9                      uint32 `json:"-"`
	OrnamentRefreshNum         uint8

	// Objects
	DungeonName         Hash
	RoleRestrictList    []int32
	RoleRecommendList   []int32
	EnemyDisplayList    []int32
	OrnamentDisplayList []int32
	TipDisplayList      []Hash
	UnlockRoleList      []DiceyDungeonDataMetaData_RoleInfo
	UnlockWeaponList    []DiceyDungeonDataMetaData_WeaponInfo
	BgPath              StrWithPrefix16
}

type DiceyDungeonDataMetaData_RoleInfo struct {
	// Fields
	ID    int32
	Level int32
}

type DiceyDungeonDataMetaData_WeaponInfo struct {
	// Fields
	ID    int32
	Level int32
}

type DiceyDungeonGachaPoolMetaData struct {
	// Fields
	GachaPoolID int32

	// Properties
	GachaPoolType uint8
	Addr1         uint32 `json:"-"`
	UnlockSite    int32

	// Objects
	GachaPoolItemList []DiceyDungeonGachaPoolMetaData_WeaponInfo
}

type DiceyDungeonGachaPoolMetaData_WeaponInfo struct {
	// Fields
	ID       int32
	Priority int32
}

type DiceyDungeonMonsterBehaviorMetaData struct {
	// Fields
	BehaviorID int32

	// Properties
	SkillID      int32
	BeginTurn    uint8
	CastTurn     uint8
	CastPriority uint8
	Addr1        uint32 `json:"-"`

	// Objects
	CastLimit StrWithPrefix16
}

type DiceyDungeonMonsterMetaData struct {
	// Fields
	MonsterID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	IsBoss    bool
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`
	MonsterHP int32
	Addr6     uint32 `json:"-"`
	Addr7     uint32 `json:"-"`
	Addr8     uint32 `json:"-"`
	Addr9     uint32 `json:"-"`
	Addr10    uint32 `json:"-"`

	// Objects
	MonsterName             Hash
	MonsterIcon             StrWithPrefix16
	MonsterDetailText       Hash
	MonsterFigurePath       StrWithPrefix16
	MonsterPrefabPath       StrWithPrefix16
	BehaviorIDList          []int32
	MonsterPassiveSKillList []int32
	SkillDisplayType1       StrWithPrefix16
	SkillDisplayType2       StrWithPrefix16
	SkillDisplayType3       StrWithPrefix16
}

type DiceyDungeonOrnamentMetaData struct {
	// Fields
	OrnamentID int32
	Level      int32

	// Properties
	MaxLevel   int32
	Addr1      uint32 `json:"-"`
	UnlockSite int32
	Addr2      uint32 `json:"-"`

	// Objects
	DungeonSkillID []int32
	BuffBriefDes   []Hash
}

type DiceyDungeonPassiveSkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	PassiveskillType uint8
	PassiveskillTag  uint8
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	ShowDetailDes    bool

	// Objects
	PassiveskillName Hash
	PassiveskillDes  Hash
	TriggerCondition StrWithPrefix16
	TriggerParam     StrWithPrefix16
	ResultOnExecute  []StrWithPrefix16
	IconPath         StrWithPrefix16
}

type DiceyDungeonRoleMetaData struct {
	// Fields
	RoleID int32
	Level  int32

	// Properties
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	Addr3                    uint32 `json:"-"`
	Addr4                    uint32 `json:"-"`
	MaxLevel                 int32
	Addr5                    uint32 `json:"-"`
	Health                   int32
	Strength                 int32
	Agility                  int32
	Intelligence             int32
	Addr6                    uint32 `json:"-"`
	Addr7                    uint32 `json:"-"`
	Addr8                    uint32 `json:"-"`
	Addr9                    uint32 `json:"-"`
	Addr10                   uint32 `json:"-"`
	Addr11                   uint32 `json:"-"`
	Addr12                   uint32 `json:"-"`
	Addr13                   uint32 `json:"-"`
	DamageIncrease           float32
	Addr14                   uint32 `json:"-"`
	OrnamentInitiativePoolID int32
	OrnamentPassivePoolID    int32
	UnlockSiteID             int32
	Addr15                   uint32 `json:"-"`

	// Objects
	RoleName              Hash
	RoleChibiIcon         StrWithPrefix16
	RoleFigurePath        StrWithPrefix16
	RolePrefabPath        StrWithPrefix16
	LevelUpCostMaterial   DiceyDungeonRoleMetaData_CostInfo
	DungeonSkillID        []int32
	SupportAbilityName    StrWithPrefix16
	SupportParamList      []float32
	SupportSkillDesc      Hash
	RecommendWeaponIDList []int32
	SkillDisplayType1     StrWithPrefix16
	SkillDisplayType2     StrWithPrefix16
	SkillDisplayType3     StrWithPrefix16
	BattleDicePool        []DiceyDungeonRoleMetaData_CostInfo
	DailyRecommendList    []int32
}

type DiceyDungeonRoleMetaData_CostInfo struct {
	// Fields
	ID  int32
	Num int32
}

type DiceyDungeonRoomMetaData struct {
	// Fields
	DungeonRoomID int32

	// Properties
	Floor                    uint8
	Position                 uint8
	Addr1                    uint32 `json:"-"`
	DungeonRoomType          uint8
	DungeonRoomTypeParameter int32
	CheckEventType           uint8
	CheckEventParameter      int32
	EnterPlotID              int32
	PassPlotID               int32
	PassReward               int32
	Addr2                    uint32 `json:"-"`
	Addr3                    uint32 `json:"-"`
	FristDropReward          int32
	Addr4                    uint32 `json:"-"`
	Addr5                    uint32 `json:"-"`

	// Objects
	NextRoomList         []int32
	DisplayDialogTitle   Hash
	DisplayDialogContent Hash
	CheckEventDesc       Hash
	CheckEventOption     Hash
}

type DiceyDungeonSiteDisplayMetaData struct {
	// Fields
	SiteID int32

	// Properties
	RoleID   int32
	RewardID int32
}

type DiceyDungeonSkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	SkillType         uint8
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	SlotType          uint8
	SlotTypeParameter uint8
	UseTimeType       uint8
	UseTimeParameter  uint8
	DisplayType       uint8
	ShowDetailDes     bool

	// Objects
	SkillName       Hash
	SkillDes        Hash
	ResultOnExecute []StrWithPrefix16
}

type DiceyDungeonTowerMetaData struct {
	// Fields
	DungeonRoomID int32

	// Properties
	FloorID         int16
	Medal           int16
	ShowTowerReward bool
}

type DiceyDungeonTutorialDataMetaData struct {
	// Fields
	DungeonRoomID int32
	TriggerType   uint8

	// Properties
	TutorialDataID int32
}

type DiceyDungeonWeaponMetaData struct {
	// Fields
	WeaponID int32
	Level    int32

	// Properties
	Addr1    uint32 `json:"-"`
	MaxLevel int32
	Addr2    uint32 `json:"-"`

	// Objects
	WeaponIcon     StrWithPrefix16
	DungeonSkillID []int32
}

type DLC2ConditionMetaData struct {
	// Fields
	ID int32

	// Properties
	ConditionType uint8
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	Parameter       []int32
	StringParameter []StrWithPrefix16
}

type DLC2DailyQuestInfoMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	SpecialRewardID  int32
	SpecialQuestTime int32

	// Objects
	QuestConsigner Hash
	QuestInfo      Hash
}

type DLC2EntryPlotLineMetaData struct {
	// Fields
	ID int32

	// Properties
	ConditionID int32
	Addr1       uint32 `json:"-"`

	// Objects
	PlotLinePath StrWithPrefix16
}

type DLC2PlotUIConfigMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	ShowHistoryBtn bool
	ShowHideBtn    bool
	ShowAutoBtn    bool
	ShowSkipBtn    bool

	// Objects
	PlotName int32
}

type DLC2SupportGiftMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	SupportExp []DLC2SupportGiftMetaData_SupportExpPair
	UpNPC      []uint8
}

type DLC2SupportGiftMetaData_SupportExpPair struct {
	// Fields
	NpcID uint8
	Exp   int32
}

type DLCGachaTabMetaData struct {
	// Fields
	ID int32

	// Properties
	ActivityID int32
}

type DLCLevelDiffMetaData struct {
	// Fields
	LevelDiff uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	AbilityName StrWithPrefix16
	IconBg      StrWithPrefix16
}

type DLCLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp         int32
	Reward      int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	IsHighlight bool
	Addr3       uint32 `json:"-"`

	// Objects
	UnlockSysTips              []Hash
	UnlockTalentTips           []DLCLevelMetaData_TalentTip
	MaterialDropLimitAddedList []DLCLevelMetaData_ItemLimit
}

type DLCLevelMetaData_TalentTip struct {
	// Fields
	AvatarID int32
	TalentID int32
}

type DLCLevelMetaData_ItemLimit struct {
	// Fields
	ItemID int32
	Limit  int32
}

type DLCNPCMetaData struct {
	// Fields
	NPCID int32

	// Properties
	IsHidden       bool
	UnlockStoryID  int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	AudioDelayTime int32
	Addr9          uint32 `json:"-"`
	Addr10         uint32 `json:"-"`
	Addr11         uint32 `json:"-"`
	Addr12         uint32 `json:"-"`

	// Objects
	HiddenList      []int32
	Name            Hash
	Portrait        StrWithPrefix16
	AvatarIcon      StrWithPrefix16
	ParentTransName StrWithPrefix16
	NPCPerfabPath   StrWithPrefix16
	ActionName      StrWithPrefix16
	Audio           StrWithPrefix16
	Lines           Hash
	DialogStoryList []int32
	NPCSupportDesc  Hash
	NPCSupportInfo  Hash
}

type DLCRecommendLevelMetaData struct {
	// Fields
	StageID int32

	// Properties
	RecommendLevel int16
	WarningLevel   int16
}

type DLCReviveCostMetaData struct {
	// Fields
	ReviveTimes int32
	ReviveType  int32

	// Properties
	CostCoin int32
}

type DLCStoryMetaData struct {
	// Fields
	DLCStoryID int32

	// Properties
	NPCID            int32
	PlotID           int32
	MissionTag       uint8
	UnlockDLCStoryID int32
	BlockDLCStoryID  int32
}

type DLCStoryPreviewMetaData struct {
	// Fields
	DLCStoryID int32

	// Properties
	PreviewStoryID int32
}

type DLCSupportDataMetaData struct {
	// Fields
	SupportType  uint8
	SupportParam int32

	// Properties
	NPCID        uint8
	SupportPoint uint16
	Addr1        uint32 `json:"-"`

	// Objects
	SupportTypeEnum int32
}

type DLCSupportLevelDataMetaData struct {
	// Fields
	SupportLevel uint8

	// Properties
	UpgradeSupportPoint int16
	Addr1               uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type DLCSupportRewardMetaData struct {
	// Fields
	NPCID        uint8
	SupportLevel uint8

	// Properties
	PlotID int32
}

type DLCTalentAffixMetaData struct {
	// Fields
	AffixID uint16

	// Properties
	PropID    uint16
	ValueMin  float32
	ValueMax  float32
	ValueStep float32
}

type DLCTalentAffixSetMetaData struct {
	// Fields
	AffixSetID uint8

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Prop1Param1 float32
	Addr5       uint32 `json:"-"`
	Prop1Param2 float32
	Addr6       uint32 `json:"-"`
	Prop1Param3 float32
	Addr7       uint32 `json:"-"`
	Addr8       uint32 `json:"-"`
	Addr9       uint32 `json:"-"`
	Prop2Param1 float32
	Addr10      uint32 `json:"-"`
	Prop2Param2 float32
	Addr11      uint32 `json:"-"`
	Prop2Param3 float32

	// Objects
	Name            Hash
	Ability1Name    StrWithPrefix16
	Prop1Desc       Hash
	Prop1ParamName1 StrWithPrefix16
	Prop1ParamName2 StrWithPrefix16
	Prop1ParamName3 StrWithPrefix16
	Ability2Name    StrWithPrefix16
	Prop2Desc       Hash
	Prop2ParamName1 StrWithPrefix16
	Prop2ParamName2 StrWithPrefix16
	Prop2ParamName3 StrWithPrefix16
}

type DLCTalentLevelMetaData struct {
	// Fields
	DLCTalentID int32
	TalentLevel int32

	// Properties
	Addr1              uint32 `json:"-"`
	LevelUpDLCLv       int32
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	HPAdd              float32
	SPAdd              float32
	ATKAdd             float32
	AbilityParamBase_1 float32
	AbilityParamBase_2 float32
	AbilityParamBase_3 float32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`

	// Objects
	LevelUpPreTalent    []DLCTalentLevelMetaData_PreTalentCondition
	LevelUpMaterialList []DLCTalentLevelMetaData_MatCost
	RefreshMaterialList []DLCTalentLevelMetaData_MatCost
	RelateAvatarList    []DLCTalentLevelMetaData_RelateAvatar
	TalentInfo          Hash
	TalentInfo2         Hash
}

type DLCTalentLevelMetaData_MatCost struct {
	// Fields
	ID     int32
	Number int32
}

type DLCTalentLevelMetaData_PreTalentCondition struct {
	// Fields
	ID    int32
	Level int32
}

type DLCTalentLevelMetaData_RelateAvatar struct {
	// Fields
	AvatarID int32
	Star     int32
}

type DLCTalentMetaData struct {
	// Fields
	DLCTalentID int32

	// Properties
	AvatarID           int32
	TalentType         int32
	UniqueTag          int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	CGID               int32
	IsNeedSwitch       bool
	UpgradeDisplayType uint8
	Addr6              uint32 `json:"-"`
	IsMultiDisplay     bool
	DisplayPriority    uint8

	// Objects
	TalentName               Hash
	TalentShortName          Hash
	IconPath                 StrWithPrefix16
	UIPointName              StrWithPrefix16
	UILineName               []StrWithPrefix16
	UpgradeAdditionalDisplay Hash
}

type DLCTalentMultiDisplayMetaData struct {
	// Fields
	ID uint8

	// Properties
	DLCTalentID     int32
	TalentLevel     uint8
	CGID            int32
	DisplayPriority uint8
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`

	// Objects
	TalentInfo               Hash
	TalentInfo2              Hash
	UpgradeAdditionalDisplay Hash
}

type DLCTalentTagMetaData struct {
	// Fields
	UniqueTag int32

	// Properties
	Weight int32
}

type DLCTowerBonusMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BuffList []int32
	Name     StrWithPrefix16
	Desc     StrWithPrefix16
	IconPath StrWithPrefix16
}

type DLCTowerBuffMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	BuffType int32

	// Objects
	BuffName    StrWithPrefix16
	OverrideMap []StrWithPrefix16
}

type DLCTowerFloorMetaData struct {
	// Fields
	Floor      int32
	ConfigType int32

	// Properties
	RecommendLevel int32
	WarningLevel   int32
	HardLevel      int32
	SpawnType      int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	TimeBonus      int32
	IsCheckPoint   int32
	Reward         int32
	Addr4          uint32 `json:"-"`
	HardLevelGroup int32
	MaxCoin        int32

	// Objects
	SpawnTypeDesc StrWithPrefix16
	Parameter     []int32
	WaveListPool  []StrWithPrefix16
	StageSection  StrWithPrefix16
}

type DLCTowerMonsterMetaData struct {
	// Fields
	UniqueID int32

	// Properties
	AddTime int32
	AddCoin int32
}

type DLCTowerRankMetaData struct {
	// Fields
	Rank int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ScheduleReward []DLCTowerRankMetaData_RewardData
}

type DLCTowerRankMetaData_RewardData struct {
	// Fields
	ConfigType int32
	RewaredID  int32
}

type DLCTowerScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	NextID     int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	ConfigType int32
	LevelID    int32
	MaxFloor   int32
	Weather    int32
	Addr4      uint32 `json:"-"`

	// Objects
	StartTime  StrWithPrefix16
	EndTime    StrWithPrefix16
	SettleTime StrWithPrefix16
	Bonus      []int32
}

type DLCTowerStyleBonusMetaData struct {
	// Fields
	Rank int32

	// Properties
	Rate float32
}

type DLCTowerWaveMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MonsterConfig []int32
}

type DLCTowerWeatherMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BuffList []int32
	Name     StrWithPrefix16
	Desc     StrWithPrefix16
	IconPath StrWithPrefix16
}

type DormitoryAvatarMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	RoleID               int32
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	FaceAnimType         int32
	MoveType             float32
	KneelToAngry         float32
	PairAvatarID         int32
	JumpPar              float32
	SitPar               float32
	SitTime              float32
	SitToSubAction       float32
	SitToSleepPar        float32
	SitToDrinkPar        float32
	SitToEatPar          float32
	SleepPar             float32
	SleepTime            float32
	GroundSitPar         float32
	GroundSitTime        float32
	GroundKneelPar       float32
	GroundKneelTime      float32
	LieDownPar           float32
	LieDownTime          float32
	LieDownToSubAction   float32
	LieDownToSleepPar    float32
	UseDeskPar           float32
	UseDeskTime          float32
	StandGamePar         float32
	StandGameTime        float32
	OpenFridgePar        float32
	OpenFridgeTime       float32
	OpenFridgeToDrinkPar float32
	OpenFridgeToEatPar   float32
	WatchTVPar           float32
	WatchTVTime          float32
	StandBookPar         float32
	StandBookTime        float32
	StandArcadePar       float32
	StandArcadeTime      float32
	HandWritePar         float32
	HandWriteTime        float32
	PlayZhengPar         float32
	PlayZhengTime        float32
	PlayGoPar            float32
	PlayGoTime           float32
	StandEatPar          float32
	StandEatIcePar       float32
	BaseType             float32
	SleepType            float32
	DrinkType            float32
	EatType              float32
	TVShowType           float32
	ComputerType         float32
	ArcadeType           float32
	LieDownType          float32
	BookType             float32
	HandWriteType        float32
	UiOrder              int32

	// Objects
	Name        StrWithPrefix16
	QAvatarIcon StrWithPrefix16
	NameText    Hash
}

type DormitoryComfortMetaData struct {
	// Fields
	ComfortLevel int32

	// Properties
	NeedComfortValues          int32
	PowerCostBonus             int32
	Addr1                      uint32 `json:"-"`
	Addr2                      uint32 `json:"-"`
	Addr3                      uint32 `json:"-"`
	EnergyCost                 int32
	AdventureGrainRecoverBonus int32
	StaminaReward              int32

	// Objects
	ComfortTitleText Hash
	ComfortIcon      StrWithPrefix16
	ComfortIconColor StrWithPrefix16
}

type DormitoryDecorationEffectMetaData struct {
	// Fields
	EffectID int32

	// Properties
	EffectType    int32
	EffectPara    int32
	EffectSubPara int32
}

type DormitoryDecorationMetaData struct {
	// Fields
	DecorationLevel int32
	TargetHouse     int32

	// Properties
	TargetRoom         int32
	DecorationPhase    int32
	DecorationStepShow int32
	DecorationMaxStep  int32
	NeedComfort        int32
	Addr1              uint32 `json:"-"`
	NeedTime           int32
	Addr2              uint32 `json:"-"`

	// Objects
	DecorationEffectList []int32
	MaterialList         []DormitoryDecorationMetaData_DecorationMaterialNode
}

type DormitoryDecorationMetaData_DecorationMaterialNode struct {
	// Fields
	MaterialID int32
	Number     int32
}

type DormitoryDialogMetaData struct {
	// Fields
	DialogID int32

	// Properties
	RoleID        int32
	AvatarID      int32
	FromHour      int32
	ToHour        int32
	NeedFurniture int32
	NoFurniture   int32
	NeedRole      int32
	NoRole        int32
	Weight        int32
	Addr1         uint32 `json:"-"`

	// Objects
	DialogText Hash
}

type DormitoryEventDialogMetaData struct {
	// Fields
	ID int32

	// Properties
	EventKey int32
	RoleID   int32
	AvatarID int32
	Addr1    uint32 `json:"-"`
	Weight   float32

	// Objects
	TextMap Hash
}

type DormitoryEventSequenceMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Avatar        int8
	Wait          float32
	WaitRandomAdd float32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	TalkToAvatar  int8
	FaceAnimStop  bool
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`

	// Objects
	Type             StrWithPrefix16
	TalkPop          StrWithPrefix16
	TalkTxt          StrWithPrefix16
	FaceAnimType     StrWithPrefix16
	TriggerAction    StrWithPrefix16
	TriggerSubAction StrWithPrefix16
}

type DormitoryEventWeightMetaData struct {
	// Fields
	ID int32

	// Properties
	RoleID   int32
	AvatarID int32
	EventID  int32
	Weight   float32
}

type DormitoryFacilityBonusDropMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	FacilityType  int32
	FacilityLevel int32

	// Objects
	BeginTime   StrWithPrefix16
	EndTime     StrWithPrefix16
	DisplayList []int32
}

type DormitoryFacilityMetaData struct {
	// Fields
	FacilityId int32

	// Properties
	FacilityType int32
	HouseID      int32
	Addr1        uint32 `json:"-"`
	SwitchOrder  int32
	Addr2        uint32 `json:"-"`

	// Objects
	FacilityModPath  StrWithPrefix16
	FacilityNameText Hash
}

type DormitoryFurnitureCollectRewardMetaData struct {
	// Fields
	MissionID int32
}

type DormitoryFurnitureSuitMetaData struct {
	// Fields
	SuitId int32

	// Properties
	Rarity   int32
	SuitType int32
	UIOrder  int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`

	// Objects
	SuitNameText      Hash
	SuitNameEn        Hash
	SuitVersionText   Hash
	SuitIcon          StrWithPrefix16
	SmallSuitIcon     StrWithPrefix16
	RewardMissionList []int32
}

type DormitoryHouseMetaData struct {
	// Fields
	HouseID int32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	UnlockLv       int32
	UnlockMaterial int32
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	Addr10         uint32 `json:"-"`

	// Objects
	TotalRoomList          []int32
	HouseAssetPath         StrWithPrefix16
	HouseFacilityList      []DormitoryHouseMetaData_Facility
	HouseNameText          Hash
	HouseTypeText          Hash
	HouseEnterImgPath      StrWithPrefix16
	HouseEnterImgLineColor StrWithPrefix16
	HouseIconPath          StrWithPrefix16
	HouseDescText          Hash
	HouseIntro             StrWithPrefix16
}

type DormitoryHouseMetaData_Facility struct {
	// Fields
	Col int32
	Dir int32
	Id  int32
	Row int32
}

type DormitoryInteractEventMetaData struct {
	// Fields
	EventID int32

	// Properties
	InteractEvent int32
	TouchExp      int32
}

type DormitoryRoomMetaData struct {
	// Fields
	RoomID int32

	// Properties
	RoomType      int32
	SizeType      int32
	BornPositionX float32
	BornPositionY float32
	BornPositionZ float32
	RoomSquare    int32
	AvatarLimit   int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`

	// Objects
	RoomNameText     Hash
	RoomUnlockText   Hash
	InitialLimitList []int32
}

type DormitoryTalkEventMetaData struct {
	// Fields
	PlotId int32

	// Properties
	StartDialogId int32
	Addr1         uint32 `json:"-"`
	Reward        int32
	TouchExp      int32

	// Objects
	FinishDialogIdList []int32
}

type DormitoryVoiceMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1    uint32 `json:"-"`
	RoleID   int32
	FromHour uint8
	ToHour   uint8
	Addr2    uint32 `json:"-"`
	Weight   uint8

	// Objects
	AudioID      StrWithPrefix16
	AudioPattern StrWithPrefix16
}

type DormitoryVoteActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	ActivityName         Hash
	ActivityDes          Hash
	ProductionList       []int32
	ActivityEnterBtnIcon StrWithPrefix16
	ActivityShareBtnIcon StrWithPrefix16
}

type DormitoryVotePrizeMetaData struct {
	// Fields
	PrizeType int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	PrizeName Hash
	PrizeIcon StrWithPrefix16
}

type DormitoryVoteProductionMetaData struct {
	// Fields
	ID int32

	// Properties
	PrizeType  int32
	Rank       int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	VotingNum  int32
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`
	RoomConfig int32

	// Objects
	Author         StrWithPrefix16
	ProductionName StrWithPrefix16
	Story          Hash
	Thumbnail01    StrWithPrefix16
	Thumbnail02    StrWithPrefix16
	Thumbnail03    StrWithPrefix16
}

type DownLoadConfigMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Icontype uint8
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Isshow   bool
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	HaveRect bool

	// Objects
	AccountType   []uint8
	Icon          StrWithPrefix16
	Icontext      Hash
	LinkUrl       StrWithPrefix16
	SupplyImgPath StrWithPrefix16
	Rect          DownLoadConfigMetaData_RectTransformData
}

type DownLoadConfigMetaData_RectTransformData struct {
	// Fields
	AnchoredPositionX float32
	AnchoredPositionY float32
	AnchorMaxX        float32
	AnchorMaxY        float32
	AnchorMinX        float32
	AnchorMinY        float32
	PivotX            float32
	PivotY            float32
	SizeDeltaX        float32
	SizeDeltaY        float32
}

type DownloadInterfaceMetaData struct {
	// Fields
	ID int32

	// Properties
	BackGroundType uint8
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`

	// Objects
	BackGroundInfo   StrWithPrefix16
	ProgressBarColor StrWithPrefix16
	ActionsRect      DownloadInterfaceMetaData_RectTransformData
	ProgressBarRect  DownloadInterfaceMetaData_RectTransformData
}

type DownloadInterfaceMetaData_RectTransformData struct {
	// Fields
	AnchoredPositionX float32
	AnchoredPositionY float32
	AnchorMaxX        float32
	AnchorMaxY        float32
	AnchorMinX        float32
	AnchorMinY        float32
	PivotX            float32
	PivotY            float32
	SizeDeltaX        float32
	SizeDeltaY        float32
}

type DreamMetaData struct {
	// Fields
	DreamID int32

	// Properties
	Addr1         uint32 `json:"-"`
	ScoreMaterial int32

	// Objects
	DreamRewardList []int32
}

type DreamMissionMetaData struct {
	// Fields
	MissionID int32

	// Properties
	Score             int32
	MinLevel          int32
	MaxLevel          int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	DisplayPreMission int32

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
}

type DreamRewardMetaData struct {
	// Fields
	DreamRewardID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	RewardID      int32
	Addr5         uint32 `json:"-"`
	Score         int32
	ExchangeOpen  bool
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	ExchangeScore int32
	Addr8         uint32 `json:"-"`
	RewardShowNum uint8

	// Objects
	TitleName        Hash
	Description      Hash
	ImagePath        StrWithPrefix16
	BgPath           StrWithPrefix16
	MissionList      []int32
	ExchangeOpentime StrWithPrefix16
	ExchangeEndtime  StrWithPrefix16
	ExchangeCost     []int32
}

type DreamUnlockMetaData struct {
	// Fields
	UnlockID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	UnlockHint Hash
}

type DressTagDataMetaData struct {
	// Fields
	TagID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	TagName Hash
	TagDesc Hash
	TagIcon StrWithPrefix16
}

type DungeonsMirrorRecoveryMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity    int32
	TicketID  int32
	TicketNum float32
}

type DutyDailyMetaData struct {
	// Fields
	DutyId int32

	// Properties
	NeedDuty int32
	UnlockLv int32
	RewardId int32
}

type DutyWeeklyMetaData struct {
	// Fields
	DutyId int32

	// Properties
	NeedDuty int32
	UnlockLv int32
	RewardId int32
}

type DynamicHardLvMetaData struct {
	// Fields
	TeamLv int32

	// Properties
	HardLv int32
}

type ElevatorConfigMetaData struct {
	// Fields
	ElevatorID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ResPath StrWithPrefix16
}

type ElfLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp int32
}

type ElfStoryActMetaData struct {
	// Fields
	StoryId     int32
	TriggerTime int32

	// Properties
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	FaceAnimationDelayTime int32

	// Objects
	TriggerName       StrWithPrefix16
	TriggerEffectList []ElfStoryActMetaData_ElfUIEffectPattern
	FaceAnimationKey  StrWithPrefix16
}

type ElfStoryActMetaData_ElfUIEffectPattern struct {
	// Fields
	BeginTime float32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	AttachPoint StrWithPrefix16
	EffectPath  StrWithPrefix16
}

type ElfTalentLevelMetaData struct {
	// Fields
	TalentID    int32
	TalentLevel int32

	// Properties
	AttackValue        int32
	CritValue          int32
	NeedExp            int32
	BonusScore         float32
	LevelUpGroupLv     int32
	LevelUpElfID       int32
	LevelUpElfLv       int32
	LevelUpElfStar     int32
	AbilityParamBase_1 float32
	AbilityParamBase_2 float32
	AbilityParamBase_3 float32
}

type ElfTalentMetaData struct {
	// Fields
	ElfTalentID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	TalentGroup int32
	TalentCost  int32
	MaxLv       int32
	Addr3       uint32 `json:"-"`
	SortWeight  int32

	// Objects
	Name     Hash
	Info     Hash
	IconPath StrWithPrefix16
}

type ElfTrialMetaData struct {
	// Fields
	ID uint8

	// Properties
	ElfID   uint8
	ElfLV   uint8
	ElfStar uint8
	Addr1   uint32 `json:"-"`

	// Objects
	EquipSkillList []ElfTrialMetaData_SkillData
}

type ElfTrialMetaData_SkillData struct {
	// Fields
	ID int32
	LV int32
}

type EliteAbilityMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	AbilityName      StrWithPrefix16
	EliteAbilityDesc Hash
	EliteAbilityIcon StrWithPrefix16
	EliteAbilityText Hash
	EliteAbilityTag  StrWithPrefix16
}

type EliteChapterMetaData struct {
	// Fields
	EliteChapterID int32

	// Properties
	SortID       int32
	ChapterType  int32
	ContentID    int32
	UnlockLevel  int32
	PreviewLevel int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	RewardShow   int32
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`

	// Objects
	ChaperImgPath       StrWithPrefix16
	StageImgPath        StrWithPrefix16
	EliteChapterName    StrWithPrefix16
	KeyItemIDList       []int32
	KeyResourceTypeList []int32
}

type EliteStageCompensationMetaData struct {
	// Fields
	StageID int32

	// Properties
	FirstDropRewardID int32
}

type EliteStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1             uint32 `json:"-"`
	SortID            int32
	EliteChapterID    int32
	PreStageID        int32
	UnlockLevel       int32
	UnlockMainStoryID int32

	// Objects
	StageName StrWithPrefix16
}

type EmojiFaceDataMetaData struct {
	// Fields
	FaceID int32

	// Properties
	Addr1      uint32 `json:"-"`
	Page       int32
	Addr2      uint32 `json:"-"`
	UnlockType int32
	Addr3      uint32 `json:"-"`
	Show       int32

	// Objects
	FaceDec   Hash
	Pic       StrWithPrefix16
	UnlockDec Hash
}

type EmojiFacePageMetaData struct {
	// Fields
	Page int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	PageDec    Hash
	PageNameEn Hash
	Pic        StrWithPrefix16
}

type EntryPerformMetaData struct {
	// Fields
	PerformID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	EventList       []uint16
	EventListFollow []uint16
}

type EntryThemeDataMetaData struct {
	// Fields
	SpaceShipConfigID int32

	// Properties
	Display            bool
	ThemeUIConfig      int32
	Addr1              uint32 `json:"-"`
	RandomThemeBGM     int32
	Addr2              uint32 `json:"-"`
	WeatherGroupConfig uint8
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	Addr9              uint32 `json:"-"`
	Addr10             uint32 `json:"-"`

	// Objects
	ThemeBGMConfigList []int32
	SpecialConfigList  []int32
	ThemeName          Hash
	ThemeDesc          Hash
	ThemeGetDesc       Hash
	ThemeBGSmall       StrWithPrefix16
	ThemeBGLarge       StrWithPrefix16
	ThemeTagList       []uint16
	GiftDesc           Hash
	GiftBGPath         StrWithPrefix16
}

type EntryThemeMetaData struct {
	// Fields
	ThemeID uint8

	// Properties
	ElevatorConfig  int32
	SpaceShipConfig int32
	ThemeUIConfig   int32
	ThemeBGMConfig  int32
}

type EntryThemeScheduleMetaData struct {
	// Fields
	ThemeScheduleID uint8

	// Properties
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	ThemeID uint8

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
}

type EntryThemeTagDataMetaData struct {
	// Fields
	ThemeTagID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ThemeTagName Hash
	ThemeTagDesc Hash
	ThemeTagIcon StrWithPrefix16
}

type EquipFilterTagMetaData struct {
	// Fields
	TagID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TagName Hash
	Icon    StrWithPrefix16
}

type EquipForgeExchangeMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	GoodsID int32
}

type EquipForgeGenerationMetaData struct {
	// Fields
	GenerationID int32

	// Properties
	Priority int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	IconPath        StrWithPrefix16
	NameTextMap     Hash
	BackGroundColor StrWithPrefix16
	TagDisplay      StrWithPrefix16
}

type EquipForgeKeepAffixMetaData struct {
	// Fields
	AffixScore uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ConsumeNumList []EquipForgeKeepAffixMetaData_EFInheritCost
}

type EquipForgeKeepAffixMetaData_EFInheritCost struct {
	// Fields
	ItemMetaID int32
	ItemNum    int32
}

type EquipForgeMeta struct {
	// Fields
	ForgeId int32

	// Properties
	TargetId        int32
	TargetInitLevel int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	MaxForgeType    int32
	MaxBuyTimes     int32
	EquipmentType   int32
	SortType        int32
	Evolution       bool
	SortPriority    int32
	RefreshTimeType int32
	DisplayMinLevel int32
	MinLevel        int32
	MaxLevel        int32
	Addr5           uint32 `json:"-"`
	Generation      int32
	InferiorEquip   int32
	Addr6           uint32 `json:"-"`
	IsProgressShow  bool

	// Objects
	CurrencyShow        []int32
	CurrencyList        []EquipForgeMeta_ItemMaterial
	ForgeParaStr        []EquipForgeMeta_EFCostMatItem
	PreEquipmentList    []EquipForgeMeta_WeaponMaterial
	UnlockMainIDList    []int32
	ForgeParaStrDisplay []int32
}

type EquipForgeMeta_EFCostMatItem struct {
	// Fields
	ItemLevel  int32
	ItemMetaID int32
	ItemNum    int32
}

type EquipForgeMeta_WeaponMaterial struct {
	// Fields
	Level    int32
	WeaponID int32
}

type EquipForgeMeta_ItemMaterial struct {
	// Fields
	ItemID int32
	Number int32
}

type EquipForgeRemindMeta struct {
	// Fields
	RemindID int32

	// Properties
	MinLevel int32
	MaxLevel int32
	Addr1    uint32 `json:"-"`

	// Objects
	ForgeList []int32
}

type EquipForgeShadowMetaData struct {
	// Fields
	ShadowID int32

	// Properties
	TargetMainID               int32
	TypeIconStigmataIDRelation int32
}

type EquipmentActivitySetMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1      uint32 `json:"-"`
	SetNum     int32
	Addr2      uint32 `json:"-"`
	IsShow     bool
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	RewardType int32
	Reward     int32

	// Objects
	SetName          Hash
	DisPlayList      []int32
	BeginShowingTime StrWithPrefix16
	SetIcon          StrWithPrefix16
	SourceTextMap    Hash
}

type EquipmentLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Addr1               uint32 `json:"-"`
	WeaponUpgradeCost   int32
	WeaponEvoCost       int32
	StigmataUpgradeCost int32
	StigmataEvoCost     int32

	// Objects
	Type1 []int32
}

type EquipTypeDetailMetaData struct {
	// Fields
	EquipmentType int32
	SortType      int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	PicPath StrWithPrefix16
}

type EvaluateDialogMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	DialogTitle   Hash
	DialogContent Hash
	DialogAgree   Hash
	DialogRefuse  Hash
	DialogClose   Hash
	DialogPic     StrWithPrefix16
}

type EvaluateIntroMetaData struct {
	// Fields
	ID int32

	// Properties
	IsEnable         int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	EvaluateDialogID int32

	// Objects
	ChannelList []StrWithPrefix16
	Param1      []int32
	Param2      []int32
	Param3      []int32
}

type EventCollectionMetaData struct {
	// Fields
	CollectionID int32

	// Properties
	SysType         uint8
	CollectionType  uint8
	CollectionSubID int32
	UnlockType      uint8
	UnlockParam     int32
}

type EventDialogDataMetaData struct {
	// Fields
	DialogId int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Properties
	Addr4              uint32 `json:"-"`
	JumpID             int32
	Addr5              uint32 `json:"-"`
	DialogType         uint8
	InputID            int32
	CGRawPos           int32
	Addr6              uint32 `json:"-"`
	AvatarId           int32
	DressId            int32
	AvatarViceKey      int32
	ScreenSide         uint8
	Face               uint8
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	Addr9              uint32 `json:"-"`
	AnimID             uint8
	Distortion         uint8
	Transparency       float32
	Addr10             uint32 `json:"-"`
	Addr11             uint32 `json:"-"`
	Addr12             uint32 `json:"-"`
	ImageId            int32
	Addr13             uint32 `json:"-"`
	Addr14             uint32 `json:"-"`
	Addr15             uint32 `json:"-"`
	Backgroundeffect   uint8
	EnterEffect        uint8
	Addr16             uint32 `json:"-"`
	LaterForAudio      float32
	Addr17             uint32 `json:"-"`
	FaceVersion        uint8
	Addr18             uint32 `json:"-"`
	Addr19             uint32 `json:"-"`
	Addr20             uint32 `json:"-"`
	Addr21             uint32 `json:"-"`
	DialogueDesignType int32
	ImageChar2DID      uint32

	// Objects
	PostDialogIdList     []int32
	TalkerName           StrWithPrefix16
	QuestionContent      StrWithPrefix16
	PreDialogIdList      []int32
	LeafIDList           []int32
	AvatarName           StrWithPrefix16
	FaceAnimation        StrWithPrefix16
	FaceExpression       StrWithPrefix16
	FaceEffect           StrWithPrefix16
	BgmCover             StrWithPrefix16
	BGMVolumeControlList []StrWithPrefix16
	CgId                 StrWithPrefix16
	Content              []EventDialogDataMetaData_PlotChatNode
	Background           StrWithPrefix16
	BackgroundCG         StrWithPrefix16
	AudioId              StrWithPrefix16
	LipMotion            StrWithPrefix16
	ScreenEffect         StrWithPrefix16
	DialogueSubTitle     StrWithPrefix16
	ImagePathList        []StrWithPrefix16
	ImageDescList        []StrWithPrefix16
}

type EventDialogDataMetaData_PlotChatNode struct {
	// Fields
	ChatDuration float32
	Addr1        uint32 `json:"-"`

	// Objects
	ChatContent StrWithPrefix16
}

type ExaminationAnswerMetaData struct {
	// Fields
	AnswerID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TextMap Hash
}

type ExaminationQuestionMetaData struct {
	// Fields
	ExaminationQuestionID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	CorrectAnswerID int32
	Addr3           uint32 `json:"-"`
	LinkType        int32
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`

	// Objects
	QuestionTextMap Hash
	WrongTipTextMap Hash
	HeadImgPath     StrWithPrefix16
	LinkPram        []int32
	LinkParamStr    StrWithPrefix16
}

type ExaminationRewardMetaData struct {
	// Fields
	ConfigID                 int32
	ExaminationQuestionIndex int32

	// Properties
	RewardID int32
}

type ExaminationScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	ExaminationRewardConfigID int32
}

type ExBossMonsterLevelConfigMetaData struct {
	// Fields
	ID           int32
	MonsterLevel uint8

	// Properties
	Addr1             uint32 `json:"-"`
	TimeScoreK        int32
	BaseScoreK        int32
	Addr2             uint32 `json:"-"`
	ContinueChallenge bool
	UnlockWindow      bool

	// Objects
	TurboModeBuffName StrWithPrefix16
	TurboModeDesc     StrWithPrefix16
}

type ExBossMonsterScoreRewardMetaData struct {
	// Fields
	RewardConfigID int32
	ScoreLineNode  int32

	// Properties
	Score    int32
	RewardID int32
	Addr1    uint32 `json:"-"`

	// Objects
	RewardDesc Hash
}

type ExBossMonsterScoreRewardMetaData_RewardShowList struct {
	// Fields
	MaterialID int32
	Num        int32
}

type ExBossMonsterWeatherMetaData struct {
	// Fields
	WeatherID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	IconPath     StrWithPrefix16
	WeatherTitle Hash
	WeatherDesc  Hash
}

type ExBossTransferInfoMetaData struct {
	// Fields
	UpBossID int32

	// Properties
	BossID int32
}

type ExhibitionMetaData struct {
	// Fields
	ExhibitionID uint16

	// Properties
	MuseumID uint8
	ItemType uint8
	ItemID   int32
	LinkType uint16
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	LinkParams     []int32
	LinkParamStr   StrWithPrefix16
	LinkColor      StrWithPrefix16
	SpecialTextmap StrWithPrefix16
}

type ExposureRateDataMetaData struct {
	// Fields
	ID int32

	// Properties
	ExposureAreaRate int32
	Addr1            uint32 `json:"-"`

	// Objects
	ShopGoodsList []int32
}

type ExScheduleTextmapListMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	Id          StrWithPrefix16
	TextmapList []int32
}

type ExScheduleTextmapMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Begintime  RemoteTime
	Endtime    RemoteTime
	Newtextmap StrWithPrefix16
}

type ExtraShortStoryMetaData struct {
	// Fields
	ShortStoryID int32

	// Properties
	StoryModeChapter     int32
	ChallangeModeChapter int32
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`

	// Objects
	ShortStoryDescTitle   Hash
	ShortStoryDescContent Hash
}

type ExtraStoryAchieveDisplayMetaData struct {
	// Fields
	DisplayId int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	AreaID    int32
	Type      int32
	ChapterID int32
	Addr4     uint32 `json:"-"`
	PreAreaID int32

	// Objects
	DisplayName    Hash
	DisplayNameEng Hash
	DisplayPicPath StrWithPrefix16
	GroupIDList    []int32
}

type ExtraStoryAchieveGroupyMetaData struct {
	// Fields
	GroupID int32

	// Properties
	RewardID   int32
	GroupType  int32
	AreaID     int32
	Difficulty int32
	ChapterID  int32
}

type ExtraStoryAchieveMetaData struct {
	// Fields
	MissionId int32

	// Properties
	GroupId         int32
	AchieveLocation int32
}

type ExtraStoryActMetaData struct {
	// Fields
	ActId int32

	// Properties
	ActType       int32
	AreaId        int32
	Addr1         uint32 `json:"-"`
	Location      int32
	Addr2         uint32 `json:"-"`
	CgID          int32
	EnhanceOn     int32
	EnhanceReward int32
	Addr3         uint32 `json:"-"`

	// Objects
	StageIdList         []int32
	NameTextId          Hash
	EnhanceRestrictList []int32
}

type ExtraStoryAreaMetaData struct {
	// Fields
	AreaId int32

	// Properties
	ChapterId int32
	WithMap   int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	NameTextId        Hash
	RogueEntryPicPath StrWithPrefix16
}

type ExtraStoryChallengeModeMetaData struct {
	// Fields
	ChapterId  int32
	Difficulty int32

	// Properties
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	RewardPreview          int32
	ActRewardShow          int32
	UnlockTeamLv           int32
	FirstAreaID            int32
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`
	Addr5                  uint32 `json:"-"`
	MaxAvatarNumberPerArea int32

	// Objects
	DifficultyText       Hash
	DifficultyUnlockText Hash
	AreaList             []int32
	RecTeamLv            []int32
	RecGroupText         Hash
}

type ExtraStoryChapterMetaData struct {
	// Fields
	ChapterId int32

	// Properties
	MinPlayerLevel     int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	ChapterType        int32
	ShopType           int32
	IsFeatured         int32
	AvatarSampleSwitch bool

	// Objects
	CurrenceIDList []int32
	NameTextId     Hash
	EntryPicPath   StrWithPrefix16
	InfoTextID     Hash
}

type ExtraStoryLineMetaData struct {
	// Fields
	StoryLineId int32

	// Properties
	StoryChapterId      int32
	Weight              int32
	ActivityWeight      int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	ChallengeChapterId  int32
	UnlockStageId       int32
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	StoryFinshReward    int32
	RewardShow          int32
	ChallengeRewardShow int32
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	CgStoryLinkType     int32
	Addr9               uint32 `json:"-"`
	EnhanceUnlockLevel  int32

	// Objects
	ActivityBeginTime StrWithPrefix16
	ActivityEndTime   StrWithPrefix16
	ActivityText      Hash
	CGGroup           []int32
	StoryLineName     Hash
	UnlockText        Hash
	StoryLineDesc     Hash
	CgStoryBg         StrWithPrefix16
	CgStoryLink       StrWithPrefix16
}

type ExtraStoryThemeData struct {
	// Fields
	ThemeID int32

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	StigmataGroupBuff1  int32
	StigmataGroupBuff2  int32
	StigmataGroupBuff3  int32
	StigmataGroup2Buff1 int32
	StigmataGroup2Buff2 int32
	StigmataGroup2Buff3 int32
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	ThemeAvatarBuff     int32
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`

	// Objects
	ThemeName         Hash
	ThemeBuffList     []int32
	StigmataGroupList []int32
	ThemeAvatarList   []int32
	ThemeRoleList     []int32
	ThemeAvatarTitle  Hash
	ThemeImg          StrWithPrefix16
	ThemeDisplay      Hash
}

type ExtraStoryThemeSchedule struct {
	// Fields
	ID int32

	// Properties
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`
	ThemeID int32

	// Objects
	BeginTime   StrWithPrefix16
	EndTime     StrWithPrefix16
	ChapterList []int32
}

type FakeAvatarMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	FakeAvatarName StrWithPrefix16
	IconPath       StrWithPrefix16
	DesTextMap     Hash
}

type FarmBuffScheduleMetaData struct {
	// Fields
	BuffID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	BuffSpeedUpTime int32

	// Objects
	BuffIcon  StrWithPrefix16
	BeginTime StrWithPrefix16
}

type FarmEventMetaData struct {
	// Fields
	PhotoID int32

	// Properties
	UnlockEventExp int32
}

type FarmLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	RewardId      int32
	UnlockFarmExp int32
	MaxSlotNum    int32
}

type FarmMaterialMetaData struct {
	// Fields
	MaterialId int32

	// Properties
	EveryNum        int32
	LimitTimes      int32
	CostTime        int32
	UnlockFarmLevel int32
	RewardFarmExp   int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`

	// Objects
	FoodIcon   StrWithPrefix16
	FoodDesc   Hash
	FoodNameEn Hash
}

type FarmScheduleMetaData struct {
	// Fields
	ScheduleId int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	Addr7            uint32 `json:"-"`
	MinLevel         int32
	MaxLevel         int32
	Addr8            uint32 `json:"-"`
	Addr9            uint32 `json:"-"`
	Addr10           uint32 `json:"-"`
	PointRewardID    int32
	PointRewardLevel int32
	Addr11           uint32 `json:"-"`
	Addr12           uint32 `json:"-"`
	Addr13           uint32 `json:"-"`
	Addr14           uint32 `json:"-"`
	Addr15           uint32 `json:"-"`

	// Objects
	BeginTime        StrWithPrefix16
	ProduceEndTime   StrWithPrefix16
	EndTime          StrWithPrefix16
	DisplayEndTime   StrWithPrefix16
	BeginDayTime     StrWithPrefix16
	EndDayTime       StrWithPrefix16
	MissionList      []int32
	StagePath        StrWithPrefix16
	CGList           []StrWithPrefix16
	SpeedUpMaterLink StrWithPrefix16
	SfxActionStart   StrWithPrefix16
	SfxActionLoop    StrWithPrefix16
	UISkinList       []int32
	CameraPosition   []int32
	CameraRotation   []int32
}

type FarmSlotMetaData struct {
	// Fields
	SlotId int32

	// Properties
	Addr1           uint32 `json:"-"`
	UnlockFarmLevel int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Scale           float32

	// Objects
	CostMaterialStr StrWithPrefix16
	Position        []float32
	Rotation        []float32
}

type FarmSpeedUpMetaData struct {
	// Fields
	MaterialId int32

	// Properties
	SpeedUpTime int32
}

type FarmUISkinMetaData struct {
	// Fields
	ID int32

	// Properties
	UIType  int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	ResType int32
	Addr3   uint32 `json:"-"`

	// Objects
	UIPath  StrWithPrefix16
	ResPath StrWithPrefix16
	Color   StrWithPrefix16
}

type FastPassLevelMetaData struct {
	// Fields
	LevelId int32

	// Properties
	ID int32
}

type FastPassMetaData struct {
	// Fields
	ID uint8

	// Properties
	MaterialCost uint8
	MaterialID   uint16
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	LimitType    uint8
	LimitNum     uint8
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`

	// Objects
	BeginTime    StrWithPrefix16
	EndTime      StrWithPrefix16
	TextID       Hash
	MaterialLink StrWithPrefix16
}

type FeaturedContentMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	PreviewType           uint8
	Addr3                 uint32 `json:"-"`
	Type                  uint8
	Addr4                 uint32 `json:"-"`
	SpareType             uint8
	Addr5                 uint32 `json:"-"`
	ShareChannelContentID int32
	Addr6                 uint32 `json:"-"`
	Addr7                 uint32 `json:"-"`
	Period                uint8
	MinLevel              uint8
	MaxLevel              uint8
	Addr8                 uint32 `json:"-"`

	// Objects
	Title             Hash
	Picture           StrWithPrefix16
	PreviewContent    StrWithPrefix16
	ShareContent      StrWithPrefix16
	SpareShareContent StrWithPrefix16
	BeginTime         RemoteTime
	EndTime           RemoteTime
	DisabledChannel   []int32
}

type FeatureSubPakConfigMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Properties
	Addr3 uint32 `json:"-"`

	// Objects
	ContextName StrWithPrefix16
	Tag         StrWithPrefix16
	SubPakList  []int32
}

type FixedPlotUIConfigMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	ShowHistoryBtn bool
	ShowHideBtn    bool
	ShowAutoBtn    bool
	ShowSkipBtn    bool

	// Objects
	PlotName StrWithPrefix16
}

type FlopActivityPanel struct {
	// Fields
	ShowID int32

	// Properties
	CostItemID          int32
	NoCostTimes         int32
	CostItemNum         int32
	DailyMaxFlopTimes   int32
	AssitentActiveTimes int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`

	// Objects
	RewardList          []int32
	ImagePath           []StrWithPrefix16
	ImagePathBackground StrWithPrefix16
}

type FoundationRewardMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Level int32

	// Properties
	Reward int32
	Addr2  uint32 `json:"-"`

	// Objects
	Product_name StrWithPrefix16
	Textmap      Hash
}

type FoundationTypeMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	Minlevel int32
	Maxlevel int32

	// Objects
	Product_name StrWithPrefix16
	BGpic        StrWithPrefix16
	BGTitle      Hash
	Smallpic     StrWithPrefix16
	Smalltext    Hash
	Reference    Hash
}

type FOWEffectMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	HeightRange       float32
	XSize             int32
	ZSize             int32
	TexWidth          int32
	TexHeigth         int32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	CutOff            float32
	Luminance         float32
	StepScale         float32
	StepMinValue      float32
	Addr4             uint32 `json:"-"`
	PlayerSectorAngle float32
	PlayerRadius      float32
	FogValue          float32

	// Objects
	StageName  StrWithPrefix16
	PositionWS []float32
	RotationWS []float32
	ImagePath  StrWithPrefix16
}

type FrontEndlessBattleConfigMetaData struct {
	// Fields
	FrontEndlessFloor int32

	// Properties
	WeatherID     int32
	MechanismID   int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	TimeLimit     int32
	ChallengeTime int32
	FloorType     int32

	// Objects
	MonsterTable        []int32
	StageMonsterDisplay []int32
}

type FrontEndlessFloorConfigMetaData struct {
	// Fields
	FrontEndlessFloor int32

	// Properties
	BaseScore               int32
	FloorDisplayType        uint8
	FrontEndlessFloorReward int32
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	ChallengeRewardDisplay  int32
	Addr3                   uint32 `json:"-"`
	Addr4                   uint32 `json:"-"`
	Addr5                   uint32 `json:"-"`
	Addr6                   uint32 `json:"-"`
	IsAutoContinue          uint8

	// Objects
	FloorName          Hash
	FloorEntrancePath  StrWithPrefix16
	ChallengeFloorName Hash
	RankIconPath       StrWithPrefix16
	RankTitle          Hash
	RankDesc           Hash
}

type FrontEndlessStageMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	StageID int32
}

type FurnitureTypeMetaData struct {
	// Fields
	EditType int32

	// Properties
	ListType int32
	UIOrder  int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	TypeNameText Hash
	TypeIcon     StrWithPrefix16
}

type GachaAdventureDisplayMetaData struct {
	// Fields
	AvatarItemID int32

	// Properties
	Addr1        uint32 `json:"-"`
	AvatarWeight int32

	// Objects
	AvatarImg StrWithPrefix16
}

type GachaBoxConfigMetaData struct {
	// Fields
	BoxTypeID int32

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`
	Addr11 uint32 `json:"-"`

	// Objects
	BoxModel                StrWithPrefix16
	BoxModelMaterialList    []StrWithPrefix16
	StageModel              StrWithPrefix16
	StageOffset             []float32
	StageModelMaterialList  []StrWithPrefix16
	CameraAnim              StrWithPrefix16
	BoxModelPathInfoOneDraw StrWithPrefix16
	BoxModelPathInfoTenDraw StrWithPrefix16
	BoxPicPathInfo          StrWithPrefix16
	BoxAudioPattern         StrWithPrefix16
	BoxEffectPath           StrWithPrefix16
}

type GachaEntranceManageMetaData struct {
	// Fields
	GachaTypeID int32

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	Addr9               uint32 `json:"-"`
	UnlockLevel         int32
	Addr10              uint32 `json:"-"`
	Addr11              uint32 `json:"-"`
	Addr12              uint32 `json:"-"`
	Addr13              uint32 `json:"-"`
	Addr14              uint32 `json:"-"`
	GachaMainPageIsShow bool
	PriorityValue       int32
	SortID              int32
	GachaDefaultBoxType int32

	// Objects
	Gacha1Title                Hash
	Gacha1Desc                 Hash
	Gacha1ButtonText           Hash
	Gacha10Title               Hash
	Gacha10Desc                Hash
	Gacha10ButtonText          Hash
	GachaNameText              StrWithPrefix16
	GachaTenDrawTips           StrWithPrefix16
	Gacha10SpecificPic         StrWithPrefix16
	TitleGachaContent_Avatar   Hash
	TitleGachaContent_Weapon   Hash
	TitleGachaContent_Stigmata Hash
	TitleGachaContent_Material Hash
	TitleGachaContent_Elf      Hash
}

type GachaFirstDiscountData struct {
	// Fields
	ID int32

	// Properties
	IsEnableRedHint bool
	IsHideTimer     bool
}

type GachaGroupManageMetaData struct {
	// Fields
	GachaGroupID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	GachaTypeList      []int32
	GachaGroupName     StrWithPrefix16
	GachaGroupIconPath StrWithPrefix16
}

type GachaPanelMetaData struct {
	// Fields
	ID int32

	// Properties
	GachaType int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Postion   int32
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	BeginTime StrWithPrefix16
	EndTime   StrWithPrefix16
	ImagePath StrWithPrefix16
	Text      Hash
}

type GalAvatarStandByMetaData struct {
	// Fields
	AvatarID int32
	DressID  int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	StandByMotion StrWithPrefix16
}

type GalEventEffectMetaData struct {
	// Fields
	EffectID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Duration float32

	// Objects
	TriggerPath   StrWithPrefix16
	TriggerEffect StrWithPrefix16
}

type GalEventMetaData struct {
	// Fields
	GalEventID int32

	// Properties
	GalEventType      uint8
	Addr1             uint32 `json:"-"`
	RoleID            uint16
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	StartMotionSpeed  float32
	Addr5             uint32 `json:"-"`
	LoopMotionSpeed   float32
	LoopTimes         uint8
	Addr6             uint32 `json:"-"`
	EndMotionSpeed    float32
	AudioType         uint8
	Addr7             uint32 `json:"-"`
	AudioDelayTime    float32
	Addr8             uint32 `json:"-"`
	EffectDelayTime   float32
	Addr9             uint32 `json:"-"`
	FaceDelayTime     float32
	Priority          uint16
	BubbleSwitch      uint8
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	TouchExp          uint8
	Condition1        uint8
	Addr12            uint32 `json:"-"`
	Condition2        uint8
	Addr13            uint32 `json:"-"`
	Condition3        uint8
	Addr14            uint32 `json:"-"`
	SubsidiaryEventID int32
	PortraitActive    bool

	// Objects
	Dialogue         Hash
	AvatarID         []int32
	DressID          []int32
	StartMotion      StrWithPrefix16
	LoopMotion       StrWithPrefix16
	EndMotion        StrWithPrefix16
	Audio            StrWithPrefix16
	Effect           StrWithPrefix16
	Face             StrWithPrefix16
	Bubble           StrWithPrefix16
	BubbleCoordinate StrWithPrefix16
	Parameter1       StrWithPrefix16
	Parameter2       StrWithPrefix16
	Parameter3       StrWithPrefix16
}

type GalEventMetaData_RawData struct {
	// Fields
	GalEventID        int32
	GalEventType      uint8
	Addr1             uint32 `json:"-"`
	RoleID            uint16
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	StartMotionSpeed  float32
	Addr5             uint32 `json:"-"`
	LoopMotionSpeed   float32
	LoopTimes         uint8
	Addr6             uint32 `json:"-"`
	EndMotionSpeed    float32
	AudioType         uint8
	Addr7             uint32 `json:"-"`
	AudioDelayTime    float32
	Addr8             uint32 `json:"-"`
	EffectDelayTime   float32
	Addr9             uint32 `json:"-"`
	FaceDelayTime     float32
	Priority          uint16
	BubbleSwitch      uint8
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	TouchExp          uint8
	Condition1        uint8
	Addr12            uint32 `json:"-"`
	Condition2        uint8
	Addr13            uint32 `json:"-"`
	Condition3        uint8
	Addr14            uint32 `json:"-"`
	SubsidiaryEventID int32
	PortraitActive    bool

	// Objects
	Dialogue         Hash
	AvatarID         []int32
	DressID          []int32
	StartMotion      StrWithPrefix16
	LoopMotion       StrWithPrefix16
	EndMotion        StrWithPrefix16
	Audio            StrWithPrefix16
	Effect           StrWithPrefix16
	Face             StrWithPrefix16
	Bubble           StrWithPrefix16
	BubbleCoordinate StrWithPrefix16
	Parameter1       StrWithPrefix16
	Parameter2       StrWithPrefix16
	Parameter3       StrWithPrefix16
}

type GalEventTwinsMetaData struct {
	// Fields
	TwinsEventID int32

	// Properties
	GalEventID1 int32
	GalEventID2 int32
}

type GardenCropDataMetaData struct {
	// Fields
	CropID   int32
	GardenID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Delay    int32
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	PlayPlot int32

	// Objects
	WeatherList       StrWithPrefix16
	UnlockCropNumList StrWithPrefix16
	DelayIcon         StrWithPrefix16
	Icon              StrWithPrefix16
	Name              Hash
	Desc              Hash
}

type GardenScheduleMetaData struct {
	// Fields
	GardenID int32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	MinLevel       int32
	MaxLevel       int32
	DefaultWeather int32
	Addr3          uint32 `json:"-"`
	ProductCD      int32
	ProductMax     int32
	DailyMax       int32
	Addr4          uint32 `json:"-"`

	// Objects
	BeginTime        StrWithPrefix16
	EndTime          StrWithPrefix16
	MissionList      StrWithPrefix16
	GrowUpMaterialID StrWithPrefix16
}

type GardenWeatherDataMetaData struct {
	// Fields
	WeatherID int32
	GardenID  int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	NeedCropNum int32
	PlayPlot    int32

	// Objects
	IconPath      StrWithPrefix16
	SmallIconPath StrWithPrefix16
	WeatherName   StrWithPrefix16
	ShadowColor   StrWithPrefix16
}

type GenActivityRewardScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ShowItemList []GenActivityRewardScheduleMetaData_Item
}

type GenActivityRewardScheduleMetaData_Item struct {
	// Fields
	ID     int32
	Number int32
}

type GenActivityRewardShowItemsMetaData struct {
	// Fields
	ShowID int32

	// Properties
	ItemID     int32
	ShowRarity int32
	Addr1      uint32 `json:"-"`

	// Objects
	ComeFromList []StrWithPrefix16
}

type GeneralActivityActUIMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1        uint32 `json:"-"`
	LetterRankID int32

	// Objects
	ConfigPath StrWithPrefix16
}

type GeneralActivityDisplayDataMetaData struct {
	// Fields
	DisplayData int32

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	MainReward          int32
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	IsRankGroupHidden   bool
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	Addr9               uint32 `json:"-"`
	Addr10              uint32 `json:"-"`
	IsMP                bool
	DisplayLeftTime     bool
	DisplayScoreOnEntry bool
	Addr11              uint32 `json:"-"`
	Addr12              uint32 `json:"-"`
	IsInfoDialogHidden  bool

	// Objects
	ActivityName       Hash
	AcitivityDes       Hash
	AcitivityDetailDes Hash
	AcitivityBG        StrWithPrefix16
	RankTitle          Hash
	RankGroup          Hash
	RankIcon           StrWithPrefix16
	RankDes            Hash
	RankDetailDes      Hash
	EnterImg           StrWithPrefix16
	SpecialImgBtnPath  StrWithPrefix16
	SpecialImgPath     StrWithPrefix16
}

type GeneralActivityLinkDataMetaData struct {
	// Fields
	LinkData int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	ActivityPanelID int32
	PlotID          int32
	Addr4           uint32 `json:"-"`
	BulletId        int32
	Addr5           uint32 `json:"-"`

	// Objects
	MissionList []int32
	GachaList   []int32
	ShopList    []int32
	WebLink     StrWithPrefix16
	CgExtraKey  StrWithPrefix16
}

type GeneralActivityMetaData struct {
	// Fields
	AcitivityID int32

	// Properties
	Series         int32
	AcitivityType  int32
	MinLevel       int32
	MaxLevel       int32
	LinkData       int32
	DisplayData    int32
	ScoreData      int32
	RankData       int32
	ShowRank       int32
	TeamListID     int32
	PreCondType    int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	ActivityBuffID int32
	Addr3          uint32 `json:"-"`

	// Objects
	PreCondParaStr StrWithPrefix16
	PreUnlockHint  Hash
	TicketIDList   []uint32
}

type GeneralActivityOptionalBuffMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	StageID      int32
	Addr1        uint32 `json:"-"`
	OriginScore  int32
	WarningScore int32
	Addr2        uint32 `json:"-"`

	// Objects
	BuffPool []int32
	BGPath   StrWithPrefix16
}

type GeneralActivityRankDataMetaData struct {
	// Fields
	RankData int32
	Rank     int32

	// Properties
	Reward int32
}

type GeneralActivityScoreDataMetaData struct {
	// Fields
	ID int32

	// Properties
	RewardType      int32
	ScoreData       int32
	Score           int32
	Reward          int32
	IsSpecialReward int32
	RewardShow      int32
}

type GeneralActivityStageGroupMetaData struct {
	// Fields
	StageGroupID int32

	// Properties
	AcitivityID     int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	ChallengeNum    int32
	ChallengeReward int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`

	// Objects
	StageIDList          []int32
	StepStageIDList      []int32
	LevelPanelPrefabPath StrWithPrefix16
	GroupImagePath       StrWithPrefix16
	GroupTitle           Hash
	GroupDescTitle       Hash
	GroupDesc            Hash
	GroupBackImagePath   StrWithPrefix16
}

type GeneralActivityStageHomuMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	HomuPanel    StrWithPrefix16
	BriefPicPath StrWithPrefix16
}

type GeneralActivityStageMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	ActivityPanelType int32
	PanelSizeRatio    float32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	ShowAllGroup      bool
	LevelPanelDragDir int32
	DailyCount        int32
	Token             int32
	TokenType         int32
	TokenCount        int32
	ShowReward        int32
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	LinkType          int32
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`

	// Objects
	ActivityDesBlockPic     StrWithPrefix16
	ActivityTitlePic        StrWithPrefix16
	ActivityLevelBlockPanel StrWithPrefix16
	ActivityLevelBlockPic   StrWithPrefix16
	ShowRewardTips          Hash
	LinkIcon                StrWithPrefix16
	LinkName                StrWithPrefix16
	LinkParams              []int32
	LinkParamStr            StrWithPrefix16
}

type GeneralActivityStageRewardMetaData struct {
	// Fields
	StageID int32
}

type GeneralActivityTicketMetaData struct {
	// Fields
	TicketID uint32

	// Properties
	MaterialID int32
	MaxNum     int32
}

type GeneralRarityDisplayConfigMetaData struct {
	// Fields
	RarityID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	GachaBoxPlateEffectPath StrWithPrefix16
	GachaUIColorType        StrWithPrefix16
	GachaDropNameTextColor  StrWithPrefix16
	GachaDropIconBgPath     StrWithPrefix16
}

type GeneralScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	VersionTagTitle Hash
}

type GeniusCommonTowerMetaData struct {
	// Fields
	SiteID int16

	// Properties
	TowerID       uint8
	Type          uint8
	SiteContent   int32
	PreSite       int16
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	RecommendCost int32
	Addr3         uint32 `json:"-"`
	HardTag       bool
	IsBossLevel   bool

	// Objects
	SiteTitle         StrWithPrefix16
	MonsterIDList     []int32
	RecommendBuffList []int32
}

type GeniusEndlessTowerMetaData struct {
	// Fields
	SiteID int16

	// Properties
	TowerID uint8
	StageID int32
	Addr1   uint32 `json:"-"`
	HardTag bool

	// Objects
	RecommendAvatarIDList []int32
}

type GeniusEndlessTowerRewardMetaData struct {
	// Fields
	ItemID int16

	// Properties
	SiteID     int16
	TotalScore int32
	RewardID   int32
}

type GeniusTowerMetaData struct {
	// Fields
	TowerID uint8

	// Properties
	MapID uint16
	Type  uint8
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	TowerTitle           StrWithPrefix16
	TowerSubTitle        StrWithPrefix16
	TowerDifficultyTitle StrWithPrefix16
	ImagePath            StrWithPrefix16
	BGImagePath          StrWithPrefix16
	UnlockConditionList  StrWithPrefix16
	UnlockConditionTips  []StrWithPrefix16
}

type GerenalRulesEffectMetaData struct {
	// Fields
	RuleID int32
	Value  int32
}

type GiftPackAdditionalDisplayMetaData struct {
	// Fields
	GiftPackID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	HorizontalBg StrWithPrefix16
}

type GiftPackMetaData struct {
	// Fields
	GiftPackID int32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Rarity                int32
	DisplaySpecialEffect  int32
	ImmediateReward       int32
	ExtraRewardDays       int32
	ExtraRewardID         int32
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	SelectableDisplayType uint8

	// Objects
	GiftPackName         Hash
	GiftPackTips         Hash
	GiftPackPicPath      StrWithPrefix16
	MainRewarMatrixList  []int32
	SelectableRewardList []int32
}

type GlobalArmadaReunionLevelMetaData struct {
	// Fields
	ScheduleID int32
	Level      int32

	// Properties
	Addr1          uint32 `json:"-"`
	Percent        float32
	Displayreward1 int32
	Displayreward2 int32
	Displayreward3 int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`

	// Objects
	TextMap_Level        StrWithPrefix16
	OriginalPaintingPath StrWithPrefix16
	ActivePaintingPath   StrWithPrefix16
}

type GlobalArmadaReunionScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	OpenTime        StrWithPrefix16
	WaitingTime     StrWithPrefix16
	RewardTime      StrWithPrefix16
	RewardEndTime   StrWithPrefix16
	DailyOpenTime   RemoteTimeSpan
	CloseTime       StrWithPrefix16
	PlayInstruction StrWithPrefix16
}

type GlobalExploreActionDisplayMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Name          Hash
	PicPath       StrWithPrefix16
	DicePointList []int32
}

type GlobalExploreActionMetaData struct {
	// Fields
	ActionID int32

	// Properties
	ActionType uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`

	// Objects
	ActionName        Hash
	FlagIDList        []int32
	EventIDList       []int32
	ActionDisplayList []int32
}

type GlobalExploreAreaMetaData struct {
	// Fields
	AreaID int32

	// Properties
	AreaBuffID         int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	ThumbnailScale     float32
	Addr5              uint32 `json:"-"`
	AreaTpFrontMission int32
	DropLimitID        int32
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`

	// Objects
	AreaName           Hash
	ShowQuestList      []int32
	PreAreaList        []int32
	ThumbnailDeviation []float32
	ThumbnailPath      StrWithPrefix16
	AreaIconPath       StrWithPrefix16
	AudioSwitchEvent   StrWithPrefix16
}

type GlobalExploreBuffDataMetaData struct {
	// Fields
	AreaBuffID int32
	Level      int32

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	FlagID int32
	Addr4  uint32 `json:"-"`

	// Objects
	AbilityNameID StrWithPrefix16
	AreaBuffName  Hash
	AreaBuffDesc  Hash
	Params        []float32
}

type GlobalExploreDataMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	RpgScheduleID          int32
	ExploreStaminaCost     int32
	DonateMaterialID       int32
	DonateExchangeRatio    int32
	WinRewardList          int32
	LoseRewardList         int32
	DanmakuID              int32
	StaminaRecoverNum      int32
	StaminaRecoverInterval int32
	DropLimitScheduleID    int32
}

type GlobalExploreEntityMetaData struct {
	// Fields
	EntityID uint32

	// Properties
	Addr1                        uint32 `json:"-"`
	Addr2                        uint32 `json:"-"`
	EntityType                   int32
	Addr3                        uint32 `json:"-"`
	Addr4                        uint32 `json:"-"`
	ActiveProgress               int32
	IsGlobalProgress             bool
	IsBuffGrow                   bool
	Addr5                        uint32 `json:"-"`
	Addr6                        uint32 `json:"-"`
	GlobalExploreBloodID         uint8
	QuestID                      int32
	EntityTriggerType            uint8
	Addr7                        uint32 `json:"-"`
	Addr8                        uint32 `json:"-"`
	Addr9                        uint32 `json:"-"`
	IsCanOverlap                 bool
	ActiveExtraStaminaLimit      int32
	ActiveExtraStaminaRecoverNum int32
	ActiveStaminaItemNum         int32
	RecommendedLevel             int32
	ViewRadius                   uint8
	Addr10                       uint32 `json:"-"`
	DanmakuID                    int32

	// Objects
	EntityName           Hash
	EntityDesc           Hash
	ActionIDList         []uint32
	InteractActionIDList []uint32
	ActiveRewardTypeList []uint8
	ActiveEventIDList    []uint32
	EntityFigure         StrWithPrefix16
	EntityStrokeFigure   StrWithPrefix16
	EntityPrefab         StrWithPrefix16
	TipsIcon             StrWithPrefix16
}

type GlobalExploreEventMetaData struct {
	// Fields
	EventID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	HintTitle Hash
}

type GlobalExploreFlagMetaData struct {
	// Fields
	FlagID int32

	// Properties
	FlagType uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	ValidNum int32

	// Objects
	FlagDesc  Hash
	ParamList []int32
}

type GlobalExploreGridMapMetaData struct {
	// Fields
	GridID uint16

	// Properties
	PosX     int32
	PosY     int32
	EntityID uint32
	AreaID   uint8
	Addr1    uint32 `json:"-"`

	// Objects
	TilePic StrWithPrefix16
}

type GlobalExploreKingdomMetaData struct {
	// Fields
	KingdomID int32

	// Properties
	StartPosX       int32
	StartPosY       int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	DanmakuID       int32
	EntranceEventID int32

	// Objects
	AreaIDList         []int32
	KingdomName        Hash
	KingdomIcon        StrWithPrefix16
	KingdomFigure      StrWithPrefix16
	MissionDisplayList []int32
}

type GlobalExploreLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp   int32
	Addr1 uint32 `json:"-"`

	// Objects
	Attributes []float32
}

type GlobalExploreMapPathMetaData struct {
	// Fields
	PosX int32
	PosY int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MapPath StrWithPrefix16
}

type GlobalExploreMessageBoardMetaData struct {
	// Fields
	ID int32

	// Properties
	UnlockArea        int32
	AreaSpotActiveNum int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	MissionCondition  int32

	// Objects
	HeadPath            StrWithPrefix16
	LayerMaster         StrWithPrefix16
	MessageBoardTitle   Hash
	MessageBoardContent Hash
	Conciliations       StrWithPrefix16
	CommentList         StrWithPrefix16
	BGPath              StrWithPrefix16
	PublishTime         Hash
}

type GlobalExplorePlotMetaData struct {
	// Fields
	PlotID int32

	// Properties
	Addr1         uint32 `json:"-"`
	CorrectAction int32
	ErrorAction   int32

	// Objects
	CorrectAnswer []int32
}

type GlobalExploreQuestMetaData struct {
	// Fields
	QuestID int32

	// Properties
	QuestType           uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	QuestRecommendLevel int32
	Progress            int32
	RewardDiasplay      int32
	FinishTimesLimit    int32
	Addr3               uint32 `json:"-"`
	Order               int32

	// Objects
	QuestName        Hash
	QuestDesc        Hash
	AssociatedEntity []int32
}

type GlobalExploreStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	StageType    uint8
	EntityID     int32
	AreaBuffID   int32
	StageGroupID int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	TimeLimit    int32
	Addr5        uint32 `json:"-"`

	// Objects
	MonstersUIDList    []int32
	MonsterSpawnList   []StrWithPrefix16
	StageName          StrWithPrefix16
	BgmName            StrWithPrefix16
	TimeLimitIncrement []int32
}

type GlobalExploreStageScoreMetaData struct {
	// Fields
	TierID uint8

	// Properties
	MaxScore              int32
	ConvertContributeRate float32
}

type GlobalSupportRewardConfigMetaData struct {
	// Fields
	GlobalSupportRewardConfigID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	BeginTime              RemoteTime
	EndTime                RemoteTime
	GiftImgPath            StrWithPrefix16
	ShareImgPath           StrWithPrefix16
	ShareBackGroundImgPath StrWithPrefix16
	ShareButtonImgPath     StrWithPrefix16
}

type GlobalWarAreaMetaData struct {
	// Fields
	AreaID uint8

	// Properties
	ClearPointNum uint8
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	AreaName      Hash
	AreaUnlockTip Hash
}

type GlobalWarAvatarBuffMetaData struct {
	// Fields
	BuffLevel uint16

	// Properties
	Hp             int32
	Atk            int32
	Def            int32
	AllDamageRatio int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`

	// Objects
	BuffIcon        StrWithPrefix16
	BuffName        Hash
	BuffDesc        Hash
	BuffEffectDesc  StrWithPrefix16
	SpecialBuffList []GlobalWarAvatarBuffMetaData_BuffTypeLevel
}

type GlobalWarAvatarBuffMetaData_BuffTypeLevel struct {
	// Fields
	BuffLevel uint16
	BuffType  uint16
}

type GlobalWarAvatarCollectionMetaData struct {
	// Fields
	ID int32

	// Properties
	ScheduleId int32
	MissionId  int32
	Addr1      uint32 `json:"-"`

	// Objects
	PointList []int32
}

type GlobalWarBuffMetaData struct {
	// Fields
	BuffID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	BuffDesc Hash
	BuffIcon StrWithPrefix16
}

type GlobalWarDamageTextMetaData struct {
	// Fields
	DamageMin int32

	// Properties
	DamageMax int32
	Addr1     uint32 `json:"-"`

	// Objects
	TextID Hash
}

type GlobalWarExchangeMetaData struct {
	// Fields
	ExchangeID uint16

	// Properties
	Addr1          uint32 `json:"-"`
	GetCurrencyNum int32
	CostStamina    int32

	// Objects
	CostMaterialList []GlobalWarExchangeMetaData_CostMaterial
}

type GlobalWarExchangeMetaData_CostMaterial struct {
	// Fields
	ID  int32
	Num int32
}

type GlobalWarIsolatePointMetaData struct {
	// Fields
	IsolatePointID uint16

	// Properties
	UIIndex         int32
	Addr1           uint32 `json:"-"`
	ClientPrePlotID int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`

	// Objects
	StageList    []int32
	Lockedtext   StrWithPrefix16
	Finishedtext StrWithPrefix16
}

type GlobalWarPhotoMetaData struct {
	// Fields
	ID uint16

	// Properties
	PhotoID             int32
	PhotoTag            uint8
	StartPointID        int32
	EndPointID          int32
	StartIsolatePointID int32
	EndIsolatePointID   int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	IsShow              bool

	// Objects
	PointTypelist StrWithPrefix16
	BeginTime     StrWithPrefix16
	EndTime       StrWithPrefix16
}

type GlobalWarPointMetaData struct {
	// Fields
	PointID uint16

	// Properties
	Addr1                 uint32 `json:"-"`
	UIIndex               uint8
	PointType             uint8
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	ExchangeID            uint16
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	RewardID              int32
	Addr7                 uint32 `json:"-"`
	Addr8                 uint32 `json:"-"`
	LockAvatarId          int32
	PointHP               int32
	AreaID                int32
	Addr9                 uint32 `json:"-"`
	PreEventID            int32
	FinishEventID         int32
	DanmakuID             int32
	PointClientType       uint8
	SweepCurrencyCD       int32
	SweepCurrencyNumPerCD int32
	SweepCurrencyMaxNum   int32

	// Objects
	UnlockTime     StrWithPrefix16
	StageList      []int32
	StageBuffIcon  StrWithPrefix16
	PrePointList   []int32
	PointTitle     Hash
	PointDesc      Hash
	BuffList       []int32
	AvatarList     []int32
	PointImagePath StrWithPrefix16
}

type GlobalWarScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	GlobalRewardMinCurrency int32
	TicketID                int32
	TicketDailyMax          int32
	CurrencyID              int32
	CurrencyMax             int32
	ClientExchangeCoinID    int32
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	Addr3                   uint32 `json:"-"`
	ShopType                int32
	RewardShowID            int32
	Addr4                   uint32 `json:"-"`
	DanmakuID               int32
	PerformID               int32
	BackgroundCGID          int32
	CGID                    int32
	PlotID                  int32
	Addr5                   uint32 `json:"-"`
	Addr6                   uint32 `json:"-"`
	Addr7                   uint32 `json:"-"`

	// Objects
	AreaList         []uint8
	IsolatePointList []uint16
	PhotoOpenTime    StrWithPrefix16
	MissionList      []int32
	WebLink          StrWithPrefix16
	GachaLink        StrWithPrefix16
	Tutorial         []GlobalWarScheduleMetaData_ConfigTutorial
}

type GlobalWarScheduleMetaData_ConfigTutorial struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Content StrWithPrefix16
	Sprite  StrWithPrefix16
	Title   StrWithPrefix16
}

type GlobalWarScoreMetaData struct {
	// Fields
	Score int32

	// Properties
	Currency int32
}

type GlobalWarSpecialBuffMetaData struct {
	// Fields
	SpecialBuffType  uint16
	SpecialBuffLevel uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	StageDetailEffectList []int32
}

type GlobalWarStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	MinLevel    uint8
	MaxLevel    uint8
	MinCurrency int32
	MaxCurrency int32
	Addr1       uint32 `json:"-"`

	// Objects
	StageBuffTypeList []uint16
}

type GoBackFreeSelectMetaData struct {
	// Fields
	FreeSelectID int32

	// Properties
	Addr1             uint32 `json:"-"`
	PreSelectedRarity int32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`

	// Objects
	PreSelectedIcon StrWithPrefix16
	TitleText       Hash
	DescText        Hash
	RewardIDList    []int32
}

type GoBackFundMetaData struct {
	// Fields
	FundID int32

	// Properties
	FundPrice       int32
	FundShowMission int32
	Addr1           uint32 `json:"-"`
	ShowRewardID    int32
	Addr2           uint32 `json:"-"`

	// Objects
	DiscountText      Hash
	FundBannerLeftPic StrWithPrefix16
}

type GoBackFundRewardMetaData struct {
	// Fields
	FundID int32
	Sort   int32

	// Properties
	MissionID int32
	RewardID  int32
}

type GoBackGachaMetaData struct {
	// Fields
	GachaConfigID uint32

	// Properties
	GachaTypeID  uint32
	ShowMaterial uint32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	DescriptionText  Hash
	CurrentGachaLink StrWithPrefix16
}

type GoBackGamePlayConfigMetaData struct {
	// Fields
	ConfigID uint32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	ProgressType  uint16
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	IsEveryDayNew bool

	// Objects
	PicPath           StrWithPrefix16
	Title             Hash
	Name              Hash
	TagDesc           Hash
	Goto              StrWithPrefix16
	ProgressParam     []int32
	RewardMissionList []int32
}

type GoBackGrowUpActivityMetaData struct {
	// Fields
	GrowUpConfigID uint8
	Rank           uint8

	// Properties
	Addr1              uint32 `json:"-"`
	RankMissionGroupID int32
	EndLevel           uint8

	// Objects
	MissionChainDisplay StrWithPrefix16
}

type GoBackLoginImgMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	AvatarPosX int32
	AvatarPosY int32
	Addr6      uint32 `json:"-"`
	Addr7      uint32 `json:"-"`

	// Objects
	ImgAvatar     StrWithPrefix16
	ImgTheme      StrWithPrefix16
	ImgTitle      StrWithPrefix16
	ImgBG         StrWithPrefix16
	ImgBGLight    StrWithPrefix16
	NewLoginTitle Hash
	NewLoginInfo  Hash
}

type GoBackMissionDayMetaData struct {
	// Fields
	ScheduleID int32
	TabId      int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	MissionList []int32
	PageText    Hash
}

type GoBackMissionScoreMetaData struct {
	// Fields
	MissionID int32

	// Properties
	DisplayPreMissionID int32
	Score               int32
}

type GoBackScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	LoginCGID           int32
	ShopType            int32
	MallType            int32
	LoginRewardUIType   int32
	GoBackMissionUIType int32
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	GrowUpConfigID      uint8
	Addr9               uint32 `json:"-"`

	// Objects
	MissionList             []int32
	MissionTitleName        Hash
	MissionBanner           StrWithPrefix16
	ScoreRewardList         []int32
	LoginRewardID           []int32
	GachaConfigIDList       []int32
	MissionShopShowCurrency []int32
	FundMallShowCurrency    []int32
	GoBackGamePlayConfigID  []uint32
}

type GoBackScoreRewardMetaData struct {
	// Fields
	ScoreRewardID int32

	// Properties
	TotalScore int32
	RewardID   int32
}

type GoBackTabMetaData struct {
	// Fields
	TabID int32

	// Properties
	SortId int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`

	// Objects
	TabName          StrWithPrefix16
	TabIcon          StrWithPrefix16
	Currency         []int32
	Params           []int32
	OpenScheduleList []int32
}

type GoBackWebMetaData struct {
	// Fields
	ScheduleID int32
	TabID      int32

	// Properties
	Addr1        uint32 `json:"-"`
	MinGobackDay int32
	MaxGobackDay int32

	// Objects
	WebLink StrWithPrefix16
}

type GodWarAreaMetaData struct {
	// Fields
	AreaID int32

	// Properties
	Addr1             uint32 `json:"-"`
	PreAreaID         int32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	AreaUnlockMission int32
	Addr6             uint32 `json:"-"`
	AreaSelectEvent   int32
	TaleID            int32
	Index             int32

	// Objects
	SiteIDList    []int32
	AreaName      Hash
	AreaImage     StrWithPrefix16
	AreaLevelDesc StrWithPrefix16
	AreaBossList  []GodWarAreaMetaData_BossDataItem
	LockHintText  Hash
}

type GodWarAreaMetaData_BossDataItem struct {
	// Fields
	BossID      int32
	ShowMission int32
}

type GodWarAvatarLevelDataMetaData struct {
	// Fields
	AvatarLevelID uint32

	// Properties
	ATK uint16
	HP  uint16
}

type GodWarAvatarTaleScheduleMetaData struct {
	// Fields
	AvatarScheduleID uint16

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	VirtualGroup uint32
	Addr4        uint32 `json:"-"`
	PunishStepID int32

	// Objects
	BeginTime         RemoteTime
	EndTime           RemoteTime
	MainAvatarList    []uint32
	SupportAvatarList []uint32
}

type GodWarAvatarUpScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	DisplayUpMaterialList []int32
}

type GodWarBuffMetaData struct {
	// Fields
	BuffID    uint32
	BuffLevel uint32

	// Properties
	BuffSuit           uint32
	BuffQuality        uint32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	ShowBuffIconEffect bool

	// Objects
	AbilityName    StrWithPrefix16
	ParamList      []StrWithPrefix16
	BuffIcon       StrWithPrefix16
	BuffName       Hash
	BuffDesc       Hash
	BuffUpDesc     Hash
	SimpleBuffDesc Hash
	BuffTagList    []uint16
}

type GodWarBuffPoolMetaData struct {
	// Fields
	PoolID uint32

	// Properties
	BuffSuit uint32
}

type GodWarBuffSuitMetaData struct {
	// Fields
	SuitID int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	ShowSuitIconEffect bool
	ShowFilterOption   bool

	// Objects
	SuitIcon  StrWithPrefix16
	SuitName  Hash
	SuitDesc  Hash
	SuitImage StrWithPrefix16
}

type GodWarChallengeBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	BuffCost int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	BuffDesc      Hash
	BuffAbility   StrWithPrefix16
	BuffParamList []StrWithPrefix16
}

type GodWarChallengeLevelMetaData struct {
	// Fields
	ID int32

	// Properties
	Floor       int32
	Addr1       uint32 `json:"-"`
	ChallengeID int32
	MissionFlag int32
	ReplaceID   int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`

	// Objects
	Stage StrWithPrefix16
	Wave1 []int32
	Wave2 []int32
	Wave3 []int32
	Wave4 []int32
}

type GodWarChallengeMetaData struct {
	// Fields
	ChallengeID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	BuffCostLimit int32

	// Objects
	BossPreviewMap []GodWarChallengeMetaData_FloorBossDataItem
	EffectIDList   []int32
	BuffList       []int32
}

type GodWarChallengeMetaData_FloorBossDataItem struct {
	// Fields
	BossID         int32
	DifficultyType int32
	FloorNum       int32
}

type GodWarChallengeRewardMetaData struct {
	// Fields
	TaleID int32
	Step   int32

	// Properties
	Mission int32
	Score   int32
	Addr1   uint32 `json:"-"`

	// Objects
	Icon StrWithPrefix16
}

type GodWarChapterMetaData struct {
	// Fields
	ChapterID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	LobbyID  int32
	GodWarID int32
	Addr7    uint32 `json:"-"`
	Addr8    uint32 `json:"-"`

	// Objects
	ChapterName        Hash
	ChapterSubName     Hash
	ChapterEnterImage  StrWithPrefix16
	LockHintText       Hash
	BeginTime          RemoteTime
	EndTime            RemoteTime
	TaleIDList         []int32
	KeyNodeMissionList []int32
}

type GodWarChar2DMetaData struct {
	// Fields
	ImageChar2DID uint32

	// Properties
	CharID   uint32
	FaceType uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	ImageChar2DPath   StrWithPrefix16
	SpImageChar2DPath StrWithPrefix16
}

type GodWarClientDataMetaData struct {
	// Fields
	GodWarID uint16

	// Properties
	Addr1                  uint32 `json:"-"`
	ShopLink               uint32
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	EntryPerformID         int32
	Addr4                  uint32 `json:"-"`
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`
	Addr7                  uint32 `json:"-"`
	Addr8                  uint32 `json:"-"`
	SkipMissionMaterialID  uint32
	WeeklyRewardPreview    int32
	SpecialStageOverallKey int32

	// Objects
	PhotoIDMap              GodWarClientDataMetaData_CollectionData
	EntryIconPath           StrWithPrefix16
	WebLink                 StrWithPrefix16
	GodWarShowPopDropIDList []int32
	BlockRoleIDList         []uint32
	TopRoleIDList           []uint32
	ProgressMissionList     []int32
	MaterialList            []int32
}

type GodWarClientDataMetaData_CollectionData struct {
	// Fields
	Num          int32
	BeginPhotoID uint32
}

type GodWarCollectionDataMetaData struct {
	// Fields
	CollectionID uint32

	// Properties
	GodWarCollectionType  uint8
	GodWarSubType         uint8
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	MissionID             int32
	GodWarID              uint32
	CollectionStarNum     uint32
	Addr7                 uint32 `json:"-"`
	GodWarChapterID       int32
	CollectionMissionType int32
	CollectionOrder       int32

	// Objects
	LockName           Hash
	CollectionSubTitle Hash
	LockSubTitle       Hash
	DetailDesc         Hash
	BriefDesc          Hash
	LockDesc           Hash
	ShowImgOffset      []float32
}

type GodWarCollectionSuitDataMetaData struct {
	// Fields
	ID uint32

	// Properties
	GodWarID  uint32
	Addr1     uint32 `json:"-"`
	SuitType  int32
	Addr2     uint32 `json:"-"`
	MissionID int32

	// Objects
	CollectionIDList []int32
	SuitTextID       Hash
}

type GodWarEventFlagMetaData struct {
	// Fields
	FlagID uint32

	// Properties
	FlagType uint16
	Operator uint16
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	ParamsVar       []uint32
	ConditionTextID Hash
}

type GodWarEventMetaData struct {
	// Fields
	EventID uint32

	// Properties
	EventType uint8
	Addr1     uint32 `json:"-"`

	// Objects
	ParamsVar []uint32
}

type GodWarEventPanelMetaData struct {
	// Fields
	PanelID uint32

	// Properties
	Addr1    uint32 `json:"-"`
	IconType uint32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Star     uint8

	// Objects
	Icon StrWithPrefix16
	Name Hash
	Desc Hash
}

type GodWarEventPlotMetaData struct {
	// Fields
	PlotID uint32

	// Properties
	Addr1                uint32 `json:"-"`
	DefaultFollowEventID uint32

	// Objects
	EventParams []GodWarEventPlotMetaData_GodWarDialogOverData
}

type GodWarEventPlotMetaData_GodWarDialogOverData struct {
	// Fields
	EventID        int32
	FinishDialogID int32
}

type GodWarEventPropMetaData struct {
	// Fields
	PropID uint16

	// Properties
	Addr1     uint32 `json:"-"`
	PropEvent uint32
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`

	// Objects
	PropName     StrWithPrefix16
	InteractIcon StrWithPrefix16
	InteractDesc Hash
}

type GodWarEventScoreMetaData struct {
	// Fields
	EventID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ScoreName Hash
}

type GodWarEventStageMetaData struct {
	// Fields
	StageID uint32

	// Properties
	StartCoinEventID uint32
	FollowEventID    uint32
	FailEventID      uint32
	Score            int32
}

type GodWarExtraItemMetaData struct {
	// Fields
	ExtraItemID uint32

	// Properties
	GodWarID              uint16
	ExtraItemSkillType    uint16
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	ExtraItemStar         uint16
	Addr3                 uint32 `json:"-"`
	ExtraItemSkillOverall uint32
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	ExtraItemType         int32
	ExtraItemSuitID       int32

	// Objects
	AbilityName         StrWithPrefix16
	ParamList           []StrWithPrefix16
	ExtraItemName       Hash
	ExtraItemIcon       StrWithPrefix16
	UnlockExtraItemHint Hash
	ExtraItemSkill      Hash
}

type GodWarExtraItemTextMetaData struct {
	// Fields
	ExtraItemSkillOverallID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TextMap        []GodWarExtraItemTextMetaData_OverallTextMap
	DefaultTextMap Hash
}

type GodWarExtraItemTextMetaData_OverallTextMap struct {
	// Fields
	OverallValue uint32
	Addr1        uint32 `json:"-"`

	// Objects
	Desc Hash
}

type GodWarHardLevelMetaData struct {
	// Fields
	HardLevelID int32

	// Properties
	DamageRaito float32
}

type GodWarHardLevelScheduleMetaData struct {
	// Fields
	ID uint32

	// Properties
	AvatarScheduleID     uint16
	BeginLevel           uint32
	EndLevel             uint32
	HPRatio              float32
	ATKRatio             float32
	DEFRatio             float32
	ElementalResistRatio float32
}

type GodWarLobbyAreaTriggerDataMetaData struct {
	// Fields
	ID int32

	// Properties
	LobbyID        int32
	Addr1          uint32 `json:"-"`
	OnEnterEventID uint32
	OnExitEventID  uint32

	// Objects
	PrefabPath StrWithPrefix16
}

type GodWarLobbyAvatarDataMetaData struct {
	// Fields
	LobbyCharID uint32

	// Properties
	Addr1           uint32 `json:"-"`
	PositionID      uint32
	SubPositionID   uint32
	SceneConfigID   uint32
	AvatarAnimation int32
	Addr2           uint32 `json:"-"`
	RoleID          int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`

	// Objects
	AvatarPrefabPath StrWithPrefix16
	AvatarEventGroup []uint32
	NpcName          StrWithPrefix16
	IconPath         StrWithPrefix16
}

type GodWarLobbyButtonMetaData struct {
	// Fields
	ButtonConfigID uint32

	// Properties
	ButtonUnlockMissionID int32
	Addr1                 uint32 `json:"-"`
	IsButtonActive        bool
	SpeculativeEventID    uint32
	ButtonType            int32

	// Objects
	TargetUIPath StrWithPrefix16
}

type GodWarLobbyDataMetaData struct {
	// Fields
	LobbyID uint32

	// Properties
	Addr1  uint32 `json:"-"`
	TaleID int32

	// Objects
	SceneConfigIDList []uint32
}

type GodWarLobbyInteractActionMetaData struct {
	// Fields
	ActionID int32

	// Properties
	ActionType     int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	DisplayMission int32
	Addr4          uint32 `json:"-"`
	DeleteMission  int32

	// Objects
	ActionParam      []int32
	ImgPath          StrWithPrefix16
	InteractPropDesc Hash
	UnlockMission    []int32
}

type GodWarLobbyInteractPropMetaData struct {
	// Fields
	PropID int32

	// Properties
	LobbyID int32
	Type    int32
	Addr1   uint32 `json:"-"`
	Radius  float32
	Angle   float32
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`

	// Objects
	PrefabPath   StrWithPrefix16
	ActionIDList []int32
	HintOffset   []float32
}

type GodWarMainAvatarMetaData struct {
	// Fields
	GodWarID     uint16
	MainAvatarID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`
	Addr8 uint32 `json:"-"`

	// Objects
	AbilityName          StrWithPrefix16
	ParamList            []StrWithPrefix16
	StepUnlockMissionMap []GodWarMainAvatarMetaData_StepUnLockMissionDataItem
	SkillName            Hash
	SkillDesc            Hash
	SkillIcon            StrWithPrefix16
	AvatarMissionList    []int32
	RecommendLink        StrWithPrefix16
}

type GodWarMainAvatarMetaData_StepUnLockMissionDataItem struct {
	// Fields
	MissionID int32
	StepID    int32
}

type GodWarMainMissionNodeMetaData struct {
	// Fields
	NodeID int32

	// Properties
	ChapterID           int32
	NodeType            int32
	PreNodeID           int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	SkipCostMaterialID  int32
	SkipCostMaterialNum uint32
	IsShowTime          bool
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	RewardPreview       int32

	// Objects
	NodeName              Hash
	AssociatedMissionList []int32
	UnlockTime            RemoteTime
	Icon                  StrWithPrefix16
}

type GodWarMapMetaData struct {
	// Fields
	MapID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	AreaList     []int32
	PrefabPath   StrWithPrefix16
	LastAreaList []int32
}

type GodWarMaterialMetaData struct {
	// Fields
	GodWarID   int32
	MaterialID int32

	// Properties
	MaterialType  uint8
	QuantityLimit int32
}

type GodWarMissionDataMetaData struct {
	// Fields
	MissionID int32

	// Properties
	GodWarID           uint32
	GodWarChapterID    int32
	ParentTabID        uint32
	SubTabID           uint32
	ParentNode         uint32
	Addr1              uint32 `json:"-"`
	ShowTimeRemain     bool
	LobbyCharID        int32
	DisplayMissionType int32

	// Objects
	MissionIconPath StrWithPrefix16
}

type GodWarNodeMetaData struct {
	// Fields
	NodeID uint32

	// Properties
	PreNodeID   uint32
	Addr1       uint32 `json:"-"`
	ParentTabID uint32
	SubTabID    uint32

	// Objects
	NodeIcon StrWithPrefix16
}

type GodWarParentTabDataMetaData struct {
	// Fields
	TabID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TabIconPath   StrWithPrefix16
	TabNameTextID Hash
}

type GodWarPlotStageMetaData struct {
	// Fields
	PlotID int32

	// Properties
	StageID   int32
	MissionID int32
}

type GodWarPunishBuffMetaData struct {
	// Fields
	PunishBuffID uint32

	// Properties
	PunishLevelCost   int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	PunishBuffType    uint8
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	UnlockPunishLevel uint8
	Addr5             uint32 `json:"-"`
	UnlockMissionID   uint32
	Addr6             uint32 `json:"-"`

	// Objects
	PunishBuffName      Hash
	PunishBuffDesc      Hash
	PunishBuffAbility   StrWithPrefix16
	PunishBuffParamList []StrWithPrefix16
	UnlockHint          Hash
	PunishBuffIcon      StrWithPrefix16
}

type GodWarPunishRewardMetaData struct {
	// Fields
	PunishLevel uint16
	PunishScore uint32

	// Properties
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	AvatarExpMaterialAddNum int32
	ShowRewardMission       int32

	// Objects
	PunishGradeDesc Hash
	PunishMaterial  []GodWarPunishRewardMetaData_MaterialNumPair
}

type GodWarPunishRewardMetaData_MaterialNumPair struct {
	// Fields
	Material int32
	Num      int32
}

type GodWarPunishScoreMetaData struct {
	// Fields
	PunishLevel uint32

	// Properties
	PunishExtraScore  float32
	Addr1             uint32 `json:"-"`
	MaxPunishMaterial uint32

	// Objects
	PunishRecommendLevel []GodWarPunishScoreMetaData_AvatarLevelRatio
}

type GodWarPunishScoreMetaData_AvatarLevelRatio struct {
	// Fields
	AvatarLevel uint32
	Ratio       uint32
}

type GodWarPunishStepMetaData struct {
	// Fields
	PunishStepConfigID uint32
	PunishStep         uint8

	// Properties
	Addr1                  uint32 `json:"-"`
	PunishStepFirstReward  uint32
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	PunishStepType         uint8
	Addr4                  uint32 `json:"-"`
	PunishStepMaxLevel     uint32
	PunishStepMaxHardLevel uint32
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`
	ChallengeID            int32
	MainMissionNode        int32
	MapID                  int32
	Addr7                  uint32 `json:"-"`

	// Objects
	PunishStepTitle           Hash
	PunishStepInitialBuffList []uint32
	PunishStepUnlockHint      Hash
	PunishParamMissionLink    []uint32
	PunishStepPrefabPath      StrWithPrefix16
	StoryHint                 Hash
	SupportAvatarLimit        []int32
}

type GodWarRelationDataMetaData struct {
	// Fields
	GodWarID uint16
	RoleID   uint32
	Level    uint16

	// Properties
	AvatarID           uint32
	AvatarLevelID      uint32
	MaxLevel           uint16
	DisplayLevelCap    uint16
	Exp                uint32
	ExpMaterialID      uint32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	RewardID           uint32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	Addr9              uint32 `json:"-"`
	Addr10             uint32 `json:"-"`
	Addr11             uint32 `json:"-"`
	AvatarArtifactType uint8
	RoleType           uint8

	// Objects
	LvUpMaterial         GodWarRelationDataMetaData_UpgradeMaterial
	LvUpMissionLinkParam []uint32
	LockHint             Hash
	IconPath             StrWithPrefix16
	IconName             Hash
	IconDesc             Hash
	FigurePath           StrWithPrefix16
	FigureName           Hash
	FigureSubName        Hash
	ArchiveLinkParam     []uint32
	IconOffset           []float32
}

type GodWarRelationDataMetaData_UpgradeMaterial struct {
	// Fields
	MaterialID uint32
	Num        uint32
}

type GodWarRoleSkillDataMetaData struct {
	// Fields
	SkillID uint16

	// Properties
	RoleID              uint16
	NeedRelationLevel   uint16
	UnlockAvatarStar    uint8
	UnlockAvatarSubStar uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	ShowSkillDesc       bool

	// Objects
	AbilityName StrWithPrefix16
	ParamList   []StrWithPrefix16
	SkillName   Hash
	SkillDesc   Hash
}

type GodWarRoleStoryDataMetaData struct {
	// Fields
	StoryID uint32

	// Properties
	RoleID            uint32
	NeedRelationLevel uint32
	Addr1             uint32 `json:"-"`
	StoryType         uint8
	TypeParam         uint32
	AutoPlay          bool
	RewardID          uint32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`

	// Objects
	UnlockMissionList []uint32
	StoryTitle        Hash
	StoryDetail       Hash
	Picture           StrWithPrefix16
}

type GodWarSceneConfigMetaData struct {
	// Fields
	SceneConfigID uint32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	StartEventGroup uint32
	Addr3           uint32 `json:"-"`

	// Objects
	OpenSound          StrWithPrefix16
	CloseSound         StrWithPrefix16
	ButtonConfigIDList []uint32
}

type GodWarScheduleMetaData struct {
	// Fields
	GodWarID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	SupportAvatarSlotMaterialList []uint32
	BeginTime                     RemoteTime
	EndTime                       RemoteTime
	TicketList                    []uint32
}

type GodWarSiteMetaData struct {
	// Fields
	SiteID uint32

	// Properties
	Addr1        uint32 `json:"-"`
	SiteIndex    uint32
	SiteMaxIndex uint32
	SiteEvent    uint32
	BuffEvent    uint32
	AreaID       int32

	// Objects
	SiteName Hash
}

type GodWarStageHintMetaData struct {
	// Fields
	HintID int32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	RewardShow int32

	// Objects
	HintName           Hash
	HintImage          StrWithPrefix16
	LevelHintText      Hash
	LevelChallengeText Hash
}

type GodWarStageSkillMetaData struct {
	// Fields
	StageSkillID uint32

	// Properties
	StageSkillType uint8
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	OffsetX        float32
	OffsetY        float32
	Scale          float32

	// Objects
	StageSkillRandomAudio []StrWithPrefix16
	StageSkillImage       StrWithPrefix16
	StageSkillIcon        StrWithPrefix16
	StageSkillName        Hash
	StageSkillDesc        Hash
}

type GodWarSubTabDataMetaData struct {
	// Fields
	TabID uint32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	ImageType uint32

	// Objects
	TabIconPath   StrWithPrefix16
	TabNameTextID Hash
}

type GodWarSupportAvatarMetaData struct {
	// Fields
	GodWarID        uint16
	SupportAvatarID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	CD    uint32
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	SkillFunctionName       StrWithPrefix16
	AbilityName             StrWithPrefix16
	ParamList               []StrWithPrefix16
	SkillName               Hash
	SkillDesc               Hash
	SkillIcon               StrWithPrefix16
	UnlockSupportAvatarHint Hash
}

type GodWarSupportLevelMetaData struct {
	// Fields
	GodWarID uint16
	Level    uint32

	// Properties
	TalentPointExp uint32
	SupportType    uint16
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
	Name     Hash
	Desc     Hash
}

type GodWarTaleClientDataMetaData struct {
	// Fields
	TaleID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	RewardPreview int32
	Addr3         uint32 `json:"-"`

	// Objects
	TaleName       Hash
	TaleEnterImage StrWithPrefix16
	LockHintText   StrWithPrefix16
}

type GodWarTaleMetaData struct {
	// Fields
	TaleID uint16

	// Properties
	RogueCoin                 uint32
	StartEvent                uint32
	TeleportResetNum          uint32
	CloseVirtualAvatarMission uint32
	Addr1                     uint32 `json:"-"`
	TaleType                  int32

	// Objects
	StageList []int32
}

type GodWarTalentClientDataMetaData struct {
	// Fields
	ID uint32

	// Properties
	GodWarID   uint16
	TalentArea uint8
	TalentCol  uint16
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	TalentLineList  []uint16
	TalentIndexList []GodWarTalentClientDataMetaData_TalentIndex
}

type GodWarTalentClientDataMetaData_TalentIndex struct {
	// Fields
	Index    uint16
	TalentID uint32
}

type GodWarTalentDataMetaData struct {
	// Fields
	TalentID uint32
	GodWarID uint16

	// Properties
	Addr1            uint32 `json:"-"`
	TalentEffectType uint16
	MaxLevel         uint16
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`

	// Objects
	TaleIDList []uint16
	TalentName Hash
	IconPath   StrWithPrefix16
}

type GodWarTalentLevelDataMetaData struct {
	// Fields
	TalentID    uint32
	TalentLevel uint16

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	TalentPoint uint16

	// Objects
	TalentDesc          Hash
	LvUpFlagList        []uint32
	UpGradeMaterialList []GodWarTalentLevelDataMetaData_UpGradeMaterial
	TalentEffect        StrWithPrefix16
	EffectParamList     []StrWithPrefix16
}

type GodWarTalentLevelDataMetaData_UpGradeMaterial struct {
	// Fields
	Cost       uint32
	MaterialID uint32
}

type GodWarTaleScheduleMetaData struct {
	// Fields
	TaleScheduleID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	TaleIDList []uint16
	BeginTime  RemoteTime
	EndTime    RemoteTime
}

type GodWarTeleportMetaData struct {
	// Fields
	TeleportID uint16

	// Properties
	TeleportType  uint8
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	TeleportEvent uint32
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	Addr8         uint32 `json:"-"`

	// Objects
	TeleportProp    StrWithPrefix16
	TeleportIcon    StrWithPrefix16
	TeleportTitle   Hash
	TeleportContent Hash
	TeleportDesc    Hash
	TeleportAbility StrWithPrefix16
	InteractIcon    StrWithPrefix16
	InteractDesc    Hash
}

type GodWarTicketMetaData struct {
	// Fields
	TicketID uint32

	// Properties
	MaterialID int32
	MaxNum     int32
}

type GodWarUseAvatarMetaData struct {
	// Fields
	ChapterID int32
	Step      int32

	// Properties
	StepUnlockMission int32
	AvatarID          int32
	Addr1             uint32 `json:"-"`

	// Objects
	LobbyAvatarPrefabPath StrWithPrefix16
}

type GrandKeyBuffActiveInfoMetaData struct {
	// Fields
	UnlockGrandKeyLevel int32
	Order               int32

	// Properties
	Duration int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	Cost      []StrWithPrefix16
	CostLabel Hash
}

type GrandKeyLevel struct {
	// Fields
	GrandKeyID int32
	Level      int32

	// Properties
	UnlockWeaponLevel int32
	PlayerLevel       int32
	Addr1             uint32 `json:"-"`

	// Objects
	UpgradeMaterial []GrandKeyLevel_Item
}

type GrandKeyLevel_Item struct {
	// Fields
	ID  int32
	Num int32
}

type GrandKeyMainWeapon struct {
	// Fields
	MainWeaponID int32

	// Properties
	WeaponType       int32
	Addr1            uint32 `json:"-"`
	Open             int32
	PositionCorrect1 float32
	PositionCorrect2 float32
	PositionCorrect3 float32
	Addr2            uint32 `json:"-"`
	StrokeThickness  float32
	Addr3            uint32 `json:"-"`

	// Objects
	EffectPath  StrWithPrefix16
	StrokeColor StrWithPrefix16
	Material    StrWithPrefix16
}

type GrandKeyMetaData struct {
	// Fields
	GrandKeyID int32

	// Properties
	MaxGrandKeyLv int32
	PositionNum   int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`

	// Objects
	GrandkeyName        Hash
	MainWeaponIdList    []int32
	SubWeaponmainIDlist []int32
	BreachIconList      []StrWithPrefix16
	WebLinks            StrWithPrefix16
}

type GrandKeySkillLimitMetaData struct {
	// Fields
	LimitID int32

	// Properties
	Minlevel         int32
	GrandKeySkillNum int32
	SynchronLimit    int32
}

type GrandKeyWeaponContrast struct {
	// Fields
	ID int32

	// Properties
	WeaponMainIDBefore int32
	WeaponIDBefore     int32
	WeaponMainIDAfter  int32
	WeaponIDAfter      int32
}

type GrandKeyWeaponRaidSkillMetaData struct {
	// Fields
	WeaponID int32

	// Properties
	PropRaidID     int32
	PropRaidParam1 float32
	PropRaidParam2 float32
	PropRaidParam3 float32
}

type GratuityMessageMetaData struct {
	// Fields
	MonsterID int32

	// Properties
	Initial uint32
}

type GratuityScheduleMetaData struct {
	// Fields
	ID uint32

	// Properties
	CurrencyID     int32
	TeamListID     int32
	RewardLinkType int32
	Addr1          uint32 `json:"-"`

	// Objects
	RewardLinkParams []int32
}

type GratuityStageClassMetaData struct {
	// Fields
	StageClassID       int32
	GratuityScheduleID int32

	// Properties
	Addr1               uint32 `json:"-"`
	ReciveStageMinLevel int32
	Addr2               uint32 `json:"-"`
	TeamSubList         int32
	Addr3               uint32 `json:"-"`
	DefeatRewardNum     int32
	StaminaCoef         int32
	DamageCoef          float32
	BonusStaminaRatio   float32
	BonusDamageRatio    float32

	// Objects
	StageClassDec    Hash
	Colors           []StrWithPrefix16
	TeamSubListTitle Hash
}

type GratuityStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	EnemyNumType uint8

	// Objects
	MonsterWaveList []GratuityStageMetaData_MessageIDWaveData
	SpecialBuffText Hash
}

type GratuityStageMetaData_MessageIDWaveData struct {
	// Fields
	Wave            uint16
	MessageID       int32
	UniqueMonsterID int32
}

type GreedyEndlessBattleConfigMetaData struct {
	// Fields
	BattleConfig int32
	GroupLevel   int32
	Floor        int32

	// Properties
	StageID       int32
	WeatherID     int32
	MechanismID   int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	TimeLimit     int32
	ChallengeTime int32
	FloorType     int32

	// Objects
	MonsterTable        []int32
	StageMonsterDisplay []int32
}

type GreedyEndlessBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	Addr1    uint32 `json:"-"`
	BuffType int32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	BuffName        StrWithPrefix16
	InLevelIconPath StrWithPrefix16
	Effect          StrWithPrefix16
}

type GreedyEndlessFloorConfigMetaData struct {
	// Fields
	FloorConfigID int32
	GroupLevel    int32
	Floor         int32

	// Properties
	FloorDisplayType       int16
	FloorReward            int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`
	ChallengeRewardDisplay int32
	Addr5                  uint32 `json:"-"`
	Challenge              int32
	BaseScore              int32
	ExtraScore             int32
	Addr6                  uint32 `json:"-"`
	ChangeType             uint8
	IsAutoContinue         uint8

	// Objects
	FloorName          Hash
	RankIconPath       StrWithPrefix16
	RankTitle          Hash
	RankDesc           Hash
	ChallengeFloorName Hash
	FloorEntrancePath  StrWithPrefix16
}

type GreedyEndlessGroupMetaData struct {
	// Fields
	GroupLevel int32

	// Properties
	PlayerGroup     int32
	IsInactiveGroup bool
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`

	// Objects
	GroupName        Hash
	GroupDesc        Hash
	GroupIcon        StrWithPrefix16
	BgColor          StrWithPrefix16
	SelectPrefabPath StrWithPrefix16
}

type GreedyEndlessMechanismMetaData struct {
	// Fields
	MechanismID int32

	// Properties
	BuffID int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`

	// Objects
	MechanismNameText Hash
	MechanismDescText Hash
	MechanismIconPath StrWithPrefix16
}

type GreedyEndlessPlayerGroupMetaData struct {
	// Fields
	PlayerGroup int32

	// Properties
	MinPlayerLevel int32
	MaxPlayerLevel int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`

	// Objects
	Icon StrWithPrefix16
	Name Hash
}

type GreedyEndlessRankRewardMetaData struct {
	// Fields
	GroupLevel  int32
	RankPercent int32

	// Properties
	RankReward int32
}

type GreedyEndlessScheduleMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	SettleRewardConfig int32

	// Objects
	StartTime RemoteTime
}

type GreedyEndlessSettleConfigMetaData struct {
	// Fields
	SettleRewardConfigID int32
	GroupLevel           int32

	// Properties
	PrototeRewardID    int32
	NormalRewardID     int32
	DemoteRewardID     int32
	GroupRewardDispaly int32
}

type GreedyEndlessWeatherMetaData struct {
	// Fields
	WeatherID int32

	// Properties
	BuffID int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`

	// Objects
	WeatherNameText Hash
	WeatherDescText Hash
	WeatherIconPath StrWithPrefix16
}

type HoukaiTownBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	BuffType      uint8
	Addr1         uint32 `json:"-"`
	DurationType  uint8
	DurationNum   uint8
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	IsShowBuffDes bool

	// Objects
	ParamList     []int32
	BuffName      Hash
	BuffDes       Hash
	BuffImagePath StrWithPrefix16
}

type HoukaiTownBuildingMetaData struct {
	// Fields
	BuildingID int32

	// Properties
	ModelPath       int32
	ModelType       uint8
	Addr1           uint32 `json:"-"`
	UItype          uint8
	IsCanSell       bool
	Price           int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	EffectType      uint8
	Addr6           uint32 `json:"-"`
	Hp              int32
	BuildingLimit   int32
	BackgroundColor uint8

	// Objects
	Blueprint    StrWithPrefix16
	IconPath     StrWithPrefix16
	ImagePath    StrWithPrefix16
	BuildingName Hash
	BuildingDes  Hash
	TriggerHint  []int32
}

type HoukaiTownChallengeMetaData struct {
	// Fields
	ChallengeID int32

	// Properties
	ChallengeFinishWay uint8
	CompareType        uint8
	ChallengeParam     int32
	TargetValue        int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`

	// Objects
	Textmap       Hash
	ChallengeHint []int32
	HintTextmap   Hash
}

type HoukaiTownEventMetaData struct {
	// Fields
	EventID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	EventName   Hash
	EventDes    Hash
	ImagePath   StrWithPrefix16
	QavatarFace StrWithPrefix16
}

type HoukaiTownItemMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Type                uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	IsForbidOnFullBrick bool
	AutoUse             bool

	// Objects
	UseHintText       Hash
	PurchaseLimitText Hash
	Param             []int32
}

type HoukaiTownLuckLevelMetaData struct {
	// Fields
	Luck int32

	// Properties
	GetCoinRatio int32
	CriticalRate int32
}

type HoukaiTownMapMetaData struct {
	// Fields
	TowerID int32

	// Properties
	HonkaiTownsID      int32
	FloorID            int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	BossBornLocation   int32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	BrickRefreshTime   int32
	BrickLimitNum      int32
	CoinMaterial       int32
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	Addr9              uint32 `json:"-"`
	Addr10             uint32 `json:"-"`
	Addr11             uint32 `json:"-"`
	TutorialStepID     int32
	FirstTimePlotID    int32
	WinPlotID          int32
	RpgSiteID          int32
	RecommendStrength  int32
	RecommendAttribute int32
	Addr12             uint32 `json:"-"`
	Addr13             uint32 `json:"-"`
	StartScaleStrength int32
	MaxScaleStrength   int32
	TalentMaterial     int32
	Addr14             uint32 `json:"-"`
	IsCanRefreshBrick  bool

	// Objects
	QavatarList           []int32
	BuildingList          StrWithPrefix16
	BossList              []int32
	InitialBuildingList   StrWithPrefix16
	ChallengeList         []HoukaiTownMapMetaData_Int2
	EnterTitle            Hash
	EnterDialog           Hash
	StageName             Hash
	StageDes              Hash
	StageImagePath        StrWithPrefix16
	TalentPointRewardShow []int32
	RecommendTarget       []int32
	WebLink               StrWithPrefix16
	RefreshBrickCostList  []int32
}

type HoukaiTownMapMetaData_Int2 struct {
	// Fields
	Item1 int32
	Item2 int32
}

type HoukaiTownPathMetaData struct {
	// Fields
	TowerID     int32
	PrePosition int32

	// Properties
	DefaultNextPosition int32
	AlterNextPosition   int32
	AlterCondition      int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`

	// Objects
	AlterConditionParamList []int32
	ConditionText           Hash
}

type HoukaiTownQAvatarMetaData struct {
	// Fields
	QavatarID int32

	// Properties
	DormitoryAvatarID  int32
	Addr1              uint32 `json:"-"`
	Strength           int32
	Speed              int32
	Luck               int32
	Attribute          uint8
	HealTime           int32
	HealPricePerSecond int32
	Addr2              uint32 `json:"-"`
	LevelType          uint8
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`

	// Objects
	Skill                []int32
	HealPriceUpgradeList []int32
	QavatarName          Hash
	QavatarDesc          Hash
	ImagePath            StrWithPrefix16
	CardPicPath          StrWithPrefix16
	ChibiIcons           StrWithPrefix16
	L2DPath              StrWithPrefix16
}

type HoukaiTownQBattleSkillMetaData struct {
	// Fields
	QavatarSkillID int32

	// Properties
	Addr1          uint32 `json:"-"`
	PredicatesPara int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	ActionPara     int32
	IsOccupyARound bool
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	Addr10         uint32 `json:"-"`
	IsShowBuffDes  bool

	// Objects
	SkillPredicates      StrWithPrefix16
	PredicatesTarget     StrWithPrefix16
	PredicatesTargetPara []int32
	ActionTarget         StrWithPrefix16
	ActionTargetPara     []int32
	SkillAction          StrWithPrefix16
	StateName            StrWithPrefix16
	SkillName            Hash
	SkillDes             Hash
	SkillImagePath       StrWithPrefix16
}

type HoukaiTownQbossMetaData struct {
	// Fields
	QmonsterID int32

	// Properties
	AttackBuildingDamage int32
	AttackInterval       int32
	BornTime             int32
	Addr1                uint32 `json:"-"`

	// Objects
	StateName StrWithPrefix16
}

type HoukaiTownQMonsterMetaData struct {
	// Fields
	QmonsterID int32

	// Properties
	ModelPath int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Strength  int32
	Attribute uint8
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	LevelType uint8
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`

	// Objects
	ImagePath       StrWithPrefix16
	Skill           []int32
	DisplayDropList []int32
	QmonsterName    Hash
	L2DPath         StrWithPrefix16
	ImagePath2      StrWithPrefix16
}

type HoukaiTownRefreshMetaData struct {
	// Fields
	RefreshTagID int32

	// Properties
	Type            uint8
	RefreshInterval int32
	TowerID         int32
}

type HoukaiTownSpeedLevelMetaData struct {
	// Fields
	Speed int32

	// Properties
	MoveSpeed      int32
	CriticalDamage int32
}

type HoukaiTownStrengthLevelMetaData struct {
	// Fields
	LevelType uint8
	Strength  int32

	// Properties
	HP  int32
	ATK int32
}

type HoukaiTownTerrainSkillMetaData struct {
	// Fields
	TerrainType int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	TerrainName    Hash
	TerrainDes     Hash
	ImagePath      StrWithPrefix16
	QavatarSkillID []int32
}

type HoukaiTownTutorialStepMetaData struct {
	// Fields
	StepID int32

	// Properties
	TriggerRound    int32
	TutorialType    uint8
	TutorialID      int32
	FinishWay       uint8
	Addr1           uint32 `json:"-"`
	NotStopTutorial bool
	NextStepID      int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`

	// Objects
	FinishParam     []int32
	TutorialTextmap Hash
	FailedTextmap   Hash
	StartLocation   []HoukaiTownTutorialStepMetaData_Pointer
	TargetLocation  []HoukaiTownTutorialStepMetaData_Pointer
	TargetPrefab    []HoukaiTownTutorialStepMetaData_Pointer
}

type HoukaiTownTutorialStepMetaData_Pointer struct {
	// Fields
	ID    int32
	Addr1 uint32 `json:"-"`

	// Objects
	Path StrWithPrefix16
}

type HybridRelayPhaseMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	UpAvatarTagList []uint16
	MainEnemyList   []uint8
	Name            Hash
	Desc            Hash
}

type HybridSiteBossDataMetaData struct {
	// Fields
	BossID uint32

	// Properties
	StageID         uint32
	MinLevel        int32
	MaxLevel        int32
	StageHpAll      int32
	ChallengeOnceHp int32
}

type HybridSiteCameraMetaData struct {
	// Fields
	HybridSiteID uint32
	ChapterID    uint32

	// Properties
	CameraSite uint32
	OffsetX    float32
}

type HybridSiteContentMetaData struct {
	// Fields
	SiteContentID uint32

	// Properties
	Priority            uint32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	RelationContentID   uint32
	ExclusiveGroupID    uint32
	Difficulty          int32
	Addr4               uint32 `json:"-"`
	LevelRecommendation int32
	DropLimitID         uint16

	// Objects
	OpenTime         StrWithPrefix16
	UnlockTime       StrWithPrefix16
	CloseTime        StrWithPrefix16
	HardLevelPicPath StrWithPrefix16
}

type HybridSiteDataMetaData struct {
	// Fields
	SiteID uint32

	// Properties
	ChapterId          uint32
	ActID              uint32
	SiteType           uint32
	PreSite            uint32
	Addr1              uint32 `json:"-"`
	SiteFirstContentID uint32
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	RewardDisplayType  uint8
	RewardDisplay      uint32
	Addr6              uint32 `json:"-"`

	// Objects
	PreSiteTower        []HybridSiteDataMetaData_PreTower
	SiteContentID       []uint32
	SiteNameTextmapID   Hash
	SiteBackgroundPath  StrWithPrefix16
	SiteBackgroundPath2 StrWithPrefix16
	DisplayTimeLimit    StrWithPrefix16
}

type HybridSiteDataMetaData_PreTower struct {
	// Fields
	PreSiteID    uint32
	PreSiteScore uint32
}

type HybridSiteRemindMetaData struct {
	// Fields
	SiteID uint32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	LevelRequire   uint32
	RemindPreSite  uint32
	StopRemindSite uint32

	// Objects
	RemindStart StrWithPrefix16
	RemindClose StrWithPrefix16
}

type ImgPanelMetaData struct {
	// Fields
	PanelID_Img uint32

	// Properties
	SubType uint8
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`

	// Objects
	ImgIDList_Position1 []int32
	ImgIDList_Position2 []int32
	ImgIDList_Position3 []int32
}

type InControlActionMapMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	ActionGroup         int32
	Addr2               uint32 `json:"-"`
	BindingSettingsType int32
	HideInMobile        bool
	ActionType          int32
	ActionSubType       int32
	IsSettingComboKey   bool
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`

	// Objects
	ActionName            StrWithPrefix16
	ActionGroupString     StrWithPrefix16
	DefaultKeys           StrWithPrefix16
	DefaultInControlTypes []StrWithPrefix16
}

type InControlActionSubTypeInfoMetaData struct {
	// Fields
	ActionSubType int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	SubTypeString StrWithPrefix16
}

type InControlBattleOverrideMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	ActionName         StrWithPrefix16
	OverrideActionName StrWithPrefix16
}

type InControlControlTypeInfoMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	IsSettingControlType bool
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`

	// Objects
	ControlType       StrWithPrefix16
	PrefixControlList []InControlControlTypeInfoMetaData_PrefixControl
	IconForSony       StrWithPrefix16
	IconForXBox       StrWithPrefix16
	IconForSwitch     StrWithPrefix16
}

type InControlControlTypeInfoMetaData_PrefixControl struct {
	// Fields
	IsSettingKey bool
	Addr1        uint32 `json:"-"`

	// Objects
	PrefixControlString StrWithPrefix16
}

type InControlKeyInfoMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	IsSettingKey bool
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`

	// Objects
	Key           StrWithPrefix16
	PrefixKeyList []InControlKeyInfoMetaData_PrefixKey
	KeyString     StrWithPrefix16
}

type InControlKeyInfoMetaData_PrefixKey struct {
	// Fields
	IsSettingKey bool
	Addr1        uint32 `json:"-"`

	// Objects
	PrefixKeyString StrWithPrefix16
}

type InControlUIButtonConfigMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Properties
	Addr3               uint32 `json:"-"`
	KeyTipsType         int32
	IsDynamicallyLoaded bool

	// Objects
	ContextName      StrWithPrefix16
	PlayerActionName StrWithPrefix16
	ButtonPath       StrWithPrefix16
}

type InLevelBuffUIMetaData struct {
	// Fields
	StateID int32

	// Properties
	Addr1          uint32 `json:"-"`
	BGColor        int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	DisplayType    int32
	DisplayTypeTwo int32
	IsActive       int32
	ShowStacking   bool

	// Objects
	IconPath     StrWithPrefix16
	Description  Hash
	BuffNameDesc Hash
}

type InLevelDialogMetaData struct {
	// Fields
	DialogID int32

	// Properties
	PostDialogID int32
	Duration     float32
	TxtDuration  float32
	Addr1        uint32 `json:"-"`
	AvatarID     int32
	DressID      int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	AnimID       int32
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	ScreenEffect int32
	Addr7        uint32 `json:"-"`

	// Objects
	AvatarName    Hash
	FaceAnimation StrWithPrefix16
	FaceEffect    StrWithPrefix16
	LipMotion     StrWithPrefix16
	Content       Hash
	AudioID       StrWithPrefix16
	ScreenPic     StrWithPrefix16
}

type InviteeActivityGobackConfigMetaData struct {
	// Fields
	GobackInviteeConfigID uint32

	// Properties
	InviteeReward       int32
	InviteeMissionGroup int32
}

type InviteeActivityNewbieConfigMetaData struct {
	// Fields
	NewbieInviteeConfigID uint32

	// Properties
	InviteeReward       int32
	InviteeMissionGroup int32
}

type InviteeActivityScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	NewbieInviteeConfigID uint32
	GobackInviteeConfigID uint32
	Addr1                 uint32 `json:"-"`

	// Objects
	EndTime RemoteTime
}

type InviterActivityRewardConfigMetaData struct {
	// Fields
	ConfigID   int32
	InviteeNum int32

	// Properties
	RewardID int32
}

type InviterActivityScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	InviterRewardConfigID int32
	InviteType            uint8
	Addr1                 uint32 `json:"-"`

	// Objects
	EndTime RemoteTime
}

type JigsawGroupMetaData struct {
	// Fields
	JigsawID int32
	GroupID  int32

	// Properties
	Addr1         uint32 `json:"-"`
	GroupRewardID int32
	Addr2         uint32 `json:"-"`
	ShowReward    bool

	// Objects
	PiecesID       []int32
	GroupRewardPic StrWithPrefix16
}

type JigsawMetaData struct {
	// Fields
	JigsawID int32

	// Properties
	Width          int32
	Height         int32
	PreJigsaw      int32
	Addr1          uint32 `json:"-"`
	CostItemID     int32
	CostItemNum    int32
	FinishedReward int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	TagMaterialID  int32
	Addr7          uint32 `json:"-"`

	// Objects
	UnlockTime    StrWithPrefix16
	JigsawBgPic   StrWithPrefix16
	BackGroundPic StrWithPrefix16
	PuzzlePic     StrWithPrefix16
	TagTextID     StrWithPrefix16
	TagPic        StrWithPrefix16
	RulesTextID   StrWithPrefix16
}

type KingdomsAreaMetaData struct {
	// Fields
	ID uint16

	// Properties
	PhaseID     int16
	RewardID    uint32
	BuffID      uint16
	Addr1       uint32 `json:"-"`
	BonusRenown uint16
	Addr2       uint32 `json:"-"`

	// Objects
	PointList []uint32
	UIIndex   StrWithPrefix16
}

type KingdomsBossPointScheduleMetaData struct {
	// Fields
	BossPointScheduleID uint32

	// Properties
	PointID     uint32
	AdvanceTime uint32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`

	// Objects
	BeginTime     StrWithPrefix16
	EndTime       StrWithPrefix16
	Belief        []uint16
	RewardDisplay []uint32
	StageList     []uint32
}

type KingdomsMetaData struct {
	// Fields
	KingdomsID int16

	// Properties
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`
	Addr4   uint32 `json:"-"`
	Addr5   uint32 `json:"-"`
	Addr6   uint32 `json:"-"`
	Qavatar int32
	Addr7   uint32 `json:"-"`
	Addr8   uint32 `json:"-"`
	Addr9   uint32 `json:"-"`

	// Objects
	IconPath        StrWithPrefix16
	FlagIconPath    StrWithPrefix16
	QavatarIconPath StrWithPrefix16
	Color1          StrWithPrefix16
	Color2          StrWithPrefix16
	ColorName       StrWithPrefix16
	KingdomsName    Hash
	PhotoTitle      Hash
	PhotoDesc       StrWithPrefix16
}

type KingdomsPhaseMetaData struct {
	// Fields
	PhaseID int16

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	TagChangeNum     uint16
	PhaseType        int16
	ServerRefreshCD  uint16
	StepNum          int32
	MinScore         int32
	Addr3            uint32 `json:"-"`
	VoteGroup        int16
	VoteAddScore     int32
	RedPiontRenown   int16
	ShowPlayerNumMin uint32
	EnterTims        int32
	Addr4            uint32 `json:"-"`

	// Objects
	PhaseEndTime StrWithPrefix16
	ChangeTime   StrWithPrefix16
	RewardList   StrWithPrefix16
	AddBelief    []KingdomsPhaseMetaData_RankAddBelief
}

type KingdomsPhaseMetaData_RankAddBelief struct {
	// Fields
	BeliefAddon uint16
	Rank        uint16
}

type KingdomsPhaseTypeMetaData struct {
	// Fields
	PhaseTypeID uint8

	// Properties
	MapType          uint8
	IsMove           bool
	IsQavatarShow    bool
	RealTimeReportCD uint32
}

type KingdomsPointMetaData struct {
	// Fields
	PointID uint32

	// Properties
	Addr1                        uint32 `json:"-"`
	PointType                    uint8
	Addr2                        uint32 `json:"-"`
	Addr3                        uint32 `json:"-"`
	Renown                       uint16
	KingdomsID                   int16
	BeVoted                      bool
	Addr4                        uint32 `json:"-"`
	Addr5                        uint32 `json:"-"`
	Addr6                        uint32 `json:"-"`
	ExchangeStaminaNum           uint16
	ExchangeKindomsWarStaminaNum uint16
	ExchangeScore                uint32
	OccupyMinScore               uint32
	DisplayDifficult             uint8
	DanmakuID                    uint16

	// Objects
	UIIndex                  StrWithPrefix16
	ParaList                 []int32
	NearPointList            []int32
	PointTitle               Hash
	PointDesc                Hash
	ExchangeCostMaterialList []KingdomsPointMetaData_Material
}

type KingdomsPointMetaData_Material struct {
	// Fields
	ID  int32
	Num int32
}

type KingdomsStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	MinLevel      int16
	MaxLevel      int16
	MinScore      int32
	MaxScore      int32
	ResetCoinID   int32
	ResetCoinCost int32
	Difficult     uint8
	WaveGroupId   int32
}

type KingdomsWarRankRewardMetaData struct {
	// Fields
	ID int16

	// Properties
	GroupID  int16
	Type     int32
	Rank     int32
	RewardID int32
}

type KingdomsWarScheduleMetaData struct {
	// Fields
	ID int16

	// Properties
	Addr1                     uint32 `json:"-"`
	MinLevel                  uint8
	MaxLevel                  uint8
	Addr2                     uint32 `json:"-"`
	ScoreRewardGroup          int16
	RankRewardGroup           int16
	ScoreRewardMark           int32
	OpenEventID               uint32
	ShopType                  int16
	KingdomStaminaRecoverTime uint16
	KingdomStaminaLimit       uint8
	StaminaMaterialID         uint32
	DailyRewardMinScore       uint32
	Addr3                     uint32 `json:"-"`
	RefreshCD                 uint16
	Addr4                     uint32 `json:"-"`
	IsPictureTutorialShow     bool
	MinReportCurrencyNum      uint32
	RewardShowID              int32

	// Objects
	PhotoOpenTime StrWithPrefix16
	KingdomsList  []int32
	MissionList   []uint32
	WebLink       StrWithPrefix16
}

type KingdomsWarScoreRewardMetaData struct {
	// Fields
	ID int16

	// Properties
	GroupID  int16
	Score    int32
	RewardID int32
}

type KingdomsWarStageCurrencyMetaData struct {
	// Fields
	ID uint32

	// Properties
	PhaseID     uint32
	MinScore    uint32
	CurrencyID  uint32
	CurrencyNum uint32
}

type KingdomsWarStoryMetaData struct {
	// Fields
	ID int16

	// Properties
	PhotoType       int16
	PhotoID         int32
	PrePhotoID      int32
	PointID         uint32
	KingdomsID      int16
	NeedScore       int32
	NeedCurrencyID  int32
	NeedCurrencyNum int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	IsShow          int32

	// Objects
	BeginTime StrWithPrefix16
	EndTime   StrWithPrefix16
	LockText  Hash
}

type KingdomsWarVoteMetaData struct {
	// Fields
	ID int16

	// Properties
	Score    uint32
	Addition int16
	GroupID  int16
}

type KingdomsWarWaveMetaData struct {
	// Fields
	GroupID int32
	Wave    int32

	// Properties
	CostItem    int32
	Num         int32
	StaminaCost int32
	AddScore    int32
	Addr1       uint32 `json:"-"`

	// Objects
	DisplayID StrWithPrefix16
}

type LevelChallengeMetaData struct {
	// Fields
	ChallengeId int32

	// Properties
	ConditionId int32
	Addr1       uint32 `json:"-"`
	HintPeriod  int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Difficulty  int32

	// Objects
	ParamList     []int32
	DiaplayTarget Hash
	Explanation   Hash
}

type LevelChallengeMetaData_RawData struct {
	// Fields
	ChallengeId int32
	ConditionId int32
	Addr1       uint32 `json:"-"`
	HintPeriod  int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Difficulty  int32

	// Objects
	ParamList     []int32
	DiaplayTarget Hash
	Explanation   Hash
}

type LevelMatrixArmadaMetaData struct {
	// Fields
	ArmadaLevel int32

	// Properties
	ExplorPointLimit int32
}

type LevelMatrixBlockMetaData struct {
	// Fields
	BlockID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	BlockNameText Hash
}

type LevelMatrixBuffMetaData struct {
	// Fields
	BuffId int32

	// Properties
	Addr1       uint32 `json:"-"`
	Rarity      int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	MaxLv       int32
	ParamBase_1 float32
	ParamAdd_1  float32
	ParamBase_2 float32
	ParamAdd_2  float32
	ParamBase_3 float32
	ParamAdd_3  float32
	Addr5       uint32 `json:"-"`

	// Objects
	Name        Hash
	Desc        Hash
	IconPath    StrWithPrefix16
	AbilityName StrWithPrefix16
	BuffLink    []int32
}

type LevelMatrixEventDialogMetaData struct {
	// Fields
	Id int32

	// Properties
	PlotId int32
	Addr1  uint32 `json:"-"`

	// Objects
	EventParams []LevelMatrixEventDialogMetaData_IntIntPair
}

type LevelMatrixEventDialogMetaData_IntIntPair struct {
	// Fields
	DialogId int32
	EventId  int32
}

type LevelMatrixFloorMetaData struct {
	// Fields
	FloorID int32

	// Properties
	GroupID    int32
	BlockID    int32
	FloorDepth int32
	Addr1      uint32 `json:"-"`
	Fatigue    int32
	RecTeamLv  int32
	ThemeID    int32
	Addr2      uint32 `json:"-"`

	// Objects
	FloorNameText   Hash
	FloorPlayerShow []int32
}

type LevelMatrixFloorRequirementMetaData struct {
	// Fields
	FloorRequirementID int32

	// Properties
	RequirementProgress int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`

	// Objects
	RequirementDisplay1 Hash
	RequirementDisplay2 Hash
	RequirementHint     Hash
}

type LevelMatrixGoodsMetaData struct {
	// Fields
	GoodID int32

	// Properties
	Type        int32
	SubID       int32
	HonkaiPrice int32
	Addr1       uint32 `json:"-"`

	// Objects
	Description Hash
}

type LevelMatrixGridIndexMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	HaveFloor   bool
	ClientParam int32
	Addr2       uint32 `json:"-"`
	PlaceType   int32
	PlaceParam  float32

	// Objects
	ModPath  StrWithPrefix16
	ModScale Vector3
}

type LevelMatrixLevelLogicMetaData struct {
	// Fields
	HardLevel int32

	// Properties
	HPRatio              float32
	ATKRatio             float32
	DEFRatio             float32
	ElementalResistRatio float32
}

type LevelMatrixMapInfoMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	InfoText     Hash
	InfoIconPath StrWithPrefix16
}

type LevelMatrixMessageMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Textmap Hash
}

type LevelMatrixMonsterMetaData struct {
	// Fields
	MonsterID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	HandBookID   int32
	MonsterLevel int32

	// Objects
	MonsterName StrWithPrefix16
	ConfigType  StrWithPrefix16
}

type LevelMatrixPresetMetaData struct {
	// Fields
	PresetID int32

	// Properties
	IsBoss     bool
	Addr1      uint32 `json:"-"`
	IsBuffItem bool
	PresetType int32
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	Addr6      uint32 `json:"-"`

	// Objects
	Name                  Hash
	PresetEffectList      []int32
	LevelMatrixMonsterMod StrWithPrefix16
	PresetLua             StrWithPrefix16
	PresetLuaPara         StrWithPrefix16
	StageChallengeDataid  []int32
}

type LevelMatrixScheduleMetaData struct {
	// Fields
	MatrixScheduleID int32

	// Properties
	DisplayReward int32
	MainReward    int32
}

type LevelMatrixScoreBonusMetaData struct {
	// Fields
	ID int32

	// Properties
	ScheduleID int32
	Type       uint8
	RewardID   int32
	NeedScore  int32
}

type LevelMatrixSuddenEventMetaData struct {
	// Fields
	ID int32

	// Properties
	EventType int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	DetailDesc     Hash
	EventImagePath StrWithPrefix16
	EventName      Hash
	EventDesc      Hash
}

type LevelMatrixThemeMetaData struct {
	// Fields
	ThemeID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	FloorMod      StrWithPrefix16
	BgMod         StrWithPrefix16
	WallTallModel StrWithPrefix16
	WallLowModel  StrWithPrefix16
}

type LevelMatrixTreasureBoxMetaData struct {
	// Fields
	ID int32

	// Properties
	DropType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	EventImagePath StrWithPrefix16
	EventName      Hash
	EventDesc      Hash
}

type LevelMatrixUseItemMetaData struct {
	// Fields
	ID int32

	// Properties
	ItemType      int32
	EffectType    int32
	ParamsVar     int32
	SubParams     int32
	ItemMaxNum    int32
	Rarity        int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	IsShowOutside bool

	// Objects
	IconPath StrWithPrefix16
	ItemName Hash
	ItemDesc Hash
}

type LevelMetaData struct {
	// Fields
	LevelId int32

	// Properties
	Addr1                uint32 `json:"-"`
	ChapterId            uint8
	ActId                uint16
	SectionId            uint8
	Difficulty           uint8
	Type                 uint8
	Addr2                uint32 `json:"-"`
	BattleType           uint8
	EnterTimes           uint16
	ResetType            uint8
	ResetCoinID          int32
	ResetCostType        uint8
	ResetTimes           uint8
	StaminaCost          uint8
	AvatarExpReward      uint16
	AvatarExpInside      float32
	ScoinReward          int32
	ScoinInside          float32
	MaxScoinReward       int32
	MaxProgress          uint8
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	RecommendPlayerLevel uint8
	UnlockPlayerLevel    uint8
	UnlockStarNum        uint8
	Addr6                uint32 `json:"-"`
	Addr7                uint32 `json:"-"`
	Addr8                uint32 `json:"-"`
	Addr9                uint32 `json:"-"`
	Addr10               uint32 `json:"-"`
	Addr11               uint32 `json:"-"`
	Addr12               uint32 `json:"-"`
	IsActChallenge       bool
	FastBonusTime        uint16
	SonicBonusTime       uint16
	HardLevel            uint16
	HardLevelGroup       uint16
	ReviveTimes          uint16
	ReviveCostType       uint8
	Addr13               uint32 `json:"-"`
	TeamNum              uint8
	Addr14               uint32 `json:"-"`
	Addr15               uint32 `json:"-"`
	Addr16               uint32 `json:"-"`
	RecordLevelType      uint8
	UseDynamicHardLv     uint8
	IsTrunk              bool
	Addr17               uint32 `json:"-"`
	PlayerGetAllDrops    bool
	HardCoeff            uint16
	EnterTimesType       uint8
	IsEnterWithElf       uint8
	Addr18               uint32 `json:"-"`
	Addr19               uint32 `json:"-"`
	PreMissionLink       uint16
	Addr20               uint32 `json:"-"`
	Addr21               uint32 `json:"-"`
	Addr22               uint32 `json:"-"`
	UnlockedLink         uint16
	Addr23               uint32 `json:"-"`
	Addr24               uint32 `json:"-"`
	CostMaterialId       int32
	CostMaterialNum      int32
	FirstCostMaterialNum int32
	BalanceModeType      int32
	Addr25               uint32 `json:"-"`
	Addr26               uint32 `json:"-"`
	IsBattleYLevel       bool

	// Objects
	Name                         Hash
	Tag                          []int32
	HighlightDisplayDropIdList_1 []int32
	HighlightDisplayDropIdList   []int32
	DropList                     []int32
	PreLevelID                   []int32
	DisplayTitle                 Hash
	DisplayDetail                Hash
	BriefPicPath                 StrWithPrefix16
	DetailPicPath                StrWithPrefix16
	LuaFile                      StrWithPrefix16
	ChallengeList                []LevelMetaData_LevelChallengeMetaNode
	ReviveUseTypeList            []uint16
	MaxNumList                   []uint8
	RestrictList                 []uint16
	LoseDescList                 []Hash
	MonsterAttrShow              []uint8
	PreMissionList               []int32
	LockedText                   Hash
	PreMissionLinkParams         []int32
	PreMissionLinkParamStr       StrWithPrefix16
	UnlockedText                 Hash
	UnlockedLinkParams           []int32
	UnlockedLinkParamStr         StrWithPrefix16
	StageEntryNameList           []StrWithPrefix16
	FloatDrop                    []LevelMetaData_FloatDropParam
}

type LevelMetaData_LevelChallengeMetaNode struct {
	// Fields
	ChallengeId int32
	RewardId    int32
}

type LevelMetaData_FloatDropParam struct {
	// Fields
	MaterialId int32
	MaxNum     int32
	MinNum     int32
}

type LevelResetCostMetaData struct {
	// Fields
	Times int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Type1 []int32
}

type LevelSaveMetaData struct {
	// Fields
	LevelID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Icon        StrWithPrefix16
	Description Hash
}

type LevelStatisticsMetaData struct {
	// Fields
	StatisticsID int32

	// Properties
	ModuleID   int32
	TypeInt    int32
	HintPeriod int32
	HintRank   int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	DisplayTarget Hash
	OverrideList  []LevelStatisticsMetaData_IntIntPair
}

type LevelStatisticsMetaData_IntIntPair struct {
	// Fields
	Item2 float32
	Addr1 uint32 `json:"-"`

	// Objects
	Item1 StrWithPrefix16
}

type LevelTrialMetaData struct {
	// Fields
	ID uint16

	// Properties
	StageID             int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	AvatarTrialType     int32
	ElfTrialID          uint8
	FakeAvatarID        int32
	AllowedPlayerAvatar uint8
	Addr3               uint32 `json:"-"`
	IsAutoSelect        bool
	IsSaveTeam          bool

	// Objects
	AvatarList       []int32
	AvatarTrialGroup []LevelTrialMetaData_AvatarTrialLimit
	AvatarDress      []LevelTrialMetaData_DressInfo
}

type LevelTrialMetaData_DressInfo struct {
	// Fields
	AvatarID int32
	DressID  int32
}

type LevelTrialMetaData_AvatarTrialLimit struct {
	// Fields
	AvatarTrialID int32
	MaxLevel      int32
	MinLevel      int32
}

type LevelTutorialMetaData struct {
	// Fields
	TutorialId int32

	// Properties
	LevelId int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`

	// Objects
	ParamList     []float32
	DiaplayTarget []Hash
}

type LinearMissionData struct {
	// Fields
	MissionID int32

	// Properties
	PreMissionId         int32
	AchievementTag       uint8
	AchievementType      uint8
	WikiAchievementScore uint8
}

type LinkDataMetaData struct {
	// Fields
	LinkID uint32

	// Properties
	IsShow   bool
	LinkType uint32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	LinkParams  StrWithPrefix16
	URL         StrWithPrefix16
	LinkTextMap StrWithPrefix16
}

type LoadingTipMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	MinLevel        int32
	MaxLevel        int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Weight          int32
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	IsIgnoreZeroTag int32
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`

	// Objects
	ImgPath                StrWithPrefix16
	ImgPath_2              StrWithPrefix16
	LoadingTitle           Hash
	LoadingTagList         []int32
	LoadingTipsList        []int32
	BeginTime              StrWithPrefix16
	EndTime                StrWithPrefix16
	ExtraTriggerTypeList   []int32
	ExtraTriggerParamsList []int32
}

type LoadingTipStringMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TipTextID Hash
}

type LoginWishActivityLinkMetaData struct {
	// Fields
	LoginLinkID uint8

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	LinkMinLevel uint8
	LinkMaxLevel uint8
	Addr3        uint32 `json:"-"`

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
	LinkPic   StrWithPrefix16
}

type LoginWishActivityMetaData struct {
	// Fields
	ActivityId uint16

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Logindayisopen bool

	// Objects
	UnselectPicPath        StrWithPrefix16
	SpecialUnselectPicPath StrWithPrefix16
	TextmapRule            StrWithPrefix16
	PanelBGColor           StrWithPrefix16
	PaneldecoratingColor   StrWithPrefix16
}

type LoginWishActivityRewardMetaData struct {
	// Fields
	ActivityId uint16
	LoginDay   uint8

	// Properties
	LoginReward   int32
	SpecialReward int32
	WishId        uint8
	Addr1         uint32 `json:"-"`

	// Objects
	LoginPic StrWithPrefix16
}

type LoginWishActivityWishMetaData struct {
	// Fields
	WishId uint8

	// Properties
	SpecialWish     bool
	Addr1           uint32 `json:"-"`
	RewardDayOffset uint8

	// Objects
	CandidateRewardList []int32
}

type MailDataMetaData struct {
	// Fields
	MailID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	MailStyle int32
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`

	// Objects
	MailSender    StrWithPrefix16
	MailTitle     StrWithPrefix16
	MailContent   StrWithPrefix16
	ImgPath       StrWithPrefix16
	MailLabelPic  StrWithPrefix16
	MailLabelText Hash
}

type MailStyleMetaData struct {
	// Fields
	MailStyle int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	QuickDelete     bool
	DeleteConfirm   int32
	BanRewardGetAll bool

	// Objects
	BaseColor   StrWithPrefix16
	SelectColor StrWithPrefix16
}

type MainEventPanelMetaData struct {
	// Fields
	PanelID_MainEvent uint32

	// Properties
	Addr1                       uint32 `json:"-"`
	Addr2                       uint32 `json:"-"`
	Addr3                       uint32 `json:"-"`
	GenActivityRewardScheduleID uint16

	// Objects
	ImgIDList       []int32
	GoToLink_Raffle StrWithPrefix16
	ShopLink        []int32
}

type MainlineStepMissionDataMetaData struct {
	// Fields
	StepID int32

	// Properties
	Addr1            uint32 `json:"-"`
	ForceClosedLevel int32

	// Objects
	MissionIDList []int32
}

type MainlineStepMissionDisplayDataMetaData struct {
	// Fields
	SectionID int32

	// Properties
	SectionType          int32
	UnlockStepID         int32
	SectionDisplayReward int32
	SectionDisplayLevel  int32
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`

	// Objects
	SectionDisplayStr      Hash
	SectionDisplayHonor    Hash
	SectionDisplayIconPath StrWithPrefix16
	SectionDescription     Hash
}

type MainPageActivityEntryMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	MinLevel uint8
	MaxLevel uint8
	Priority uint8
	Type     uint8
	Addr4    uint32 `json:"-"`

	// Objects
	EntryPrefab StrWithPrefix16
	BeginTime   RemoteTime
	EndTime     RemoteTime
	Param1      []int32
}

type MainPageSequenceDialogInfoMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Priority    int32
	Addr2       uint32 `json:"-"`
	ShowType    int32
	Review      bool
	NewbieGuide bool

	// Objects
	Name      StrWithPrefix16
	CheckType []int32
}

type MainstageEventMetaData struct {
	// Fields
	Id int32

	// Properties
	ChapterID       int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	UnlockLevel     int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	ShowPointReward int32
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`

	// Objects
	BeginShowTime       RemoteTime
	EndTime             RemoteTime
	EnterIconPrefabPath StrWithPrefix16
	RewardBGPath        StrWithPrefix16
	CurrencyList        []int32
	CycleMissionIDList  []int32
	Ticket              []int32
	TicketDailyNumLimit []int32
	ShopLink            []int32
}

type MallEntranceDataMetaData struct {
	// Fields
	ID int32

	// Properties
	MallType      int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	SortID        int32
	Addr3         uint32 `json:"-"`
	MallTypeGroup uint8
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`

	// Objects
	MallTabBGPath   StrWithPrefix16
	Name            Hash
	CurrencyBarList []int32
	BeginTime       StrWithPrefix16
	EndTime         StrWithPrefix16
}

type MapGrowthMetaData struct {
	// Fields
	MapID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	MapGrowthName                     Hash
	MapGrowthInfo                     Hash
	PrefabPath                        StrWithPrefix16
	ServantChapterDescription         Hash
	UpgradeResultDialogTitleIcon      StrWithPrefix16
	UpgradeResultDialogTitleIconColor StrWithPrefix16
}

type MapLevelMetaData struct {
	// Fields
	Level int32
	MapId int32

	// Properties
	Exp        int32
	Atk        int32
	SkillPoint int32
	KeyQuestID int32
	Addr1      uint32 `json:"-"`

	// Objects
	KeyQuestDesc Hash
}

type MapSiteQualityConfigMetaData struct {
	// Fields
	QualityID uint8

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	MaterialID      int32
	MaterialDrop    int32
	ChargeIMaterial int32
	WarningLevel    int32

	// Objects
	QualityIcon       StrWithPrefix16
	QualityFrameColor StrWithPrefix16
}

type MapSiteScheduleDataMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	IsDailyRaid   bool
	ScoreType     int32
	ScoreTypePara int32
	ActivityType  int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	DailyTimesList []int32
	StageGroupList []int32
}

type MapSiteStageDataMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	PreQuality         bool
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	RandomEffectCharge int32
	MapSiteType        int32
	Addr6              uint32 `json:"-"`

	// Objects
	PreStageShowList []int32
	PreStageList     []int32
	UnlockDec        StrWithPrefix16
	MapSitePosition  []int32
	SiteQuality      []int32
	SiteBg           StrWithPrefix16
}

type MapSkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	MapID         int32
	Addr1         uint32 `json:"-"`
	PreSkillID    int32
	UnlockMapLv   int32
	UnlockStoryId int32
	UnlockType    int32
	UnlockCost    int32
	SkillType     int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`

	// Objects
	MapXY         []int32
	SkillPar      StrWithPrefix16
	SkillIconPath StrWithPrefix16
	SkillName     Hash
	SkillInfo     Hash
}

type MapSkillTypeMetaData struct {
	// Fields
	SkillType int32

	// Properties
	SkillTypeSeries int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`

	// Objects
	SkillTypeName  Hash
	SkillTypeImg   StrWithPrefix16
	SkillTypeIcon  StrWithPrefix16
	SkillTypeColor StrWithPrefix16
}

type MassiveWarDamageNeedMetaData struct {
	// Fields
	PlayerLevel uint8

	// Properties
	DamageBasic int32
	Addr1       uint32 `json:"-"`

	// Objects
	DamageNeedList []MassiveWarDamageNeedMetaData_DamageNeed
}

type MassiveWarDamageNeedMetaData_DamageNeed struct {
	// Fields
	Rank   uint8
	Damage int32
}

type MassiveWarDamageRewardMetaData struct {
	// Fields
	ScheduleID uint8
	DamageRank uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	DamageRankIconPath StrWithPrefix16
	DamageRankText     Hash
	ScoreDamageList    []MassiveWarDamageRewardMetaData_RewardScore
}

type MassiveWarDamageRewardMetaData_RewardScore struct {
	// Fields
	PlayerGroupID uint8
	Score         int32
}

type MassiveWarMonsterDataMetaData struct {
	// Fields
	BossId uint8

	// Properties
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	MessageId            uint8
	StageDetailMonsterId int32
	Addr3                uint32 `json:"-"`

	// Objects
	BossName       StrWithPrefix16
	BossPrefabPath StrWithPrefix16
	BossDetail     StrWithPrefix16
}

type MassiveWarPlayerGroupMetaData struct {
	// Fields
	PlayerGroupID uint8

	// Properties
	MinPlayerLevel uint8
	MaxPlayerLevel uint8
}

type MassiveWarRankRewardMetaData struct {
	// Fields
	ScheduleID uint8
	RankId     uint8

	// Properties
	MaxRankRatio uint16
	Addr1        uint32 `json:"-"`

	// Objects
	ScoreRankList []MassiveWarRankRewardMetaData_RewardScore
}

type MassiveWarRankRewardMetaData_RewardScore struct {
	// Fields
	PlayerGroupID uint8
	Score         int32
}

type MassiveWarRewardMetaData struct {
	// Fields
	PlayerGroupID uint16
	ScoreLevel    uint16

	// Properties
	NeedScore int32
	RewardID  int32
}

type MassiveWarScheduleMetaData struct {
	// Fields
	ScheduleID uint8

	// Properties
	MaxEnterTimes    uint8
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	RewardLinkType   int32
	RewardLinkParams int32
	Addr4            uint32 `json:"-"`

	// Objects
	ScoreBasicShowList  []MassiveWarScheduleMetaData_ScoreRange
	ScoreRankShowList   []MassiveWarScheduleMetaData_ScoreRange
	ScoreDamageShowList []MassiveWarScheduleMetaData_ScoreRange
	LinkBtnImgPath      StrWithPrefix16
}

type MassiveWarScheduleMetaData_ScoreRange struct {
	// Fields
	PlayerGroupID uint8
	MaxScore      int32
	MinScore      int32
}

type MassiveWarStageInfoMetaData struct {
	// Fields
	StageInfoId int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	IconPath        StrWithPrefix16
	DescriptionText StrWithPrefix16
}

type MassiveWarStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	DamageNeedRatio float32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`

	// Objects
	MonsterWaveList           []MassiveWarStageMetaData_MessageIDWaveData
	BossIdList                []uint8
	StageImagePath            StrWithPrefix16
	StageImagePathRankingPage StrWithPrefix16
	StageInfoIdList           []int32
}

type MassiveWarStageMetaData_MessageIDWaveData struct {
	// Fields
	Wave            uint16
	MessageID       int32
	UniqueMonsterID int32
}

type MaterialProtectMetaData struct {
	// Fields
	ID int32
}

type MaterialProvideElfExpMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	ProvideExp int32
	ScoinCost  int32
}

type MaterialUseConversionMetaData struct {
	// Fields
	ConvertIndex int32

	// Properties
	ItemID   int32
	CostNum  int32
	RewardID int32
}

type MaterialUseMetaData struct {
	// Fields
	UseID int32

	// Properties
	UseType                int32
	MultiUse               int32
	MaterialUseConfirmType uint8
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	UseMin                 int32
	UseMax                 int32
	Addr3                  uint32 `json:"-"`
	EventID                int32

	// Objects
	ParaStr      []StrWithPrefix16
	UseTip       Hash
	AdditionDesc StrWithPrefix16
}

type MatrixFloorMetaData struct {
	// Fields
	FloorID int32

	// Properties
	ThemeID   int32
	MapDataID int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	FloorNameText  Hash
	CfgName        StrWithPrefix16
	AudioNameEnter StrWithPrefix16
	AudioNameExit  StrWithPrefix16
}

type MatrixGridIndexMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	HaveFloor   bool
	ClientParam int32
	Addr3       uint32 `json:"-"`
	PlaceType   int32
	PlaceParam  float32
	CanRotate   int32

	// Objects
	ModPath           StrWithPrefix16
	EffectPatternName StrWithPrefix16
	ModScale          Vector3
}

type MatrixMapMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ForkList StrWithPrefix16
}

type MatrixPanelLinkMetaData struct {
	// Fields
	ID int32

	// Properties
	LinkType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	LinkParams   []int32
	LinkParamStr []StrWithPrefix16
	TextMapList  []Hash
}

type MatrixSuddenEventMetaData struct {
	// Fields
	ID int32

	// Properties
	EventType uint8
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	ParamsVar      []int32
	EventName      Hash
	EventDesc      Hash
	EventImagePath StrWithPrefix16
}

type MatrixThemeMetaData struct {
	// Fields
	ThemeID int32

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`

	// Objects
	FloorMod    []StrWithPrefix16
	BgMod       StrWithPrefix16
	WallEdge1   []StrWithPrefix16
	WallEdge2   []StrWithPrefix16
	WallEdge3   []StrWithPrefix16
	WallEdge4   []StrWithPrefix16
	WallCorner1 StrWithPrefix16
	WallCorner2 StrWithPrefix16
	WallCorner3 []StrWithPrefix16
	WallCorner4 StrWithPrefix16
}

type MechPaperMetaData struct {
	// Fields
	PaperID int32

	// Properties
	Addr1       uint32 `json:"-"`
	UnlockLevel int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	InfoTextMap  Hash
	PaperIcon    StrWithPrefix16
	NeedMaterial []MechPaperMetaData_BuildNeedMaterial
}

type MechPaperMetaData_BuildNeedMaterial struct {
	// Fields
	ID  int32
	Num int32
}

type MechTDDiffcultyMetaData struct {
	// Fields
	OWTechLevel int32

	// Properties
	Difficulty int32
	Weather    int32
}

type MechTDRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	Stage      int32
	Difficulty int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Reward     int32

	// Objects
	Icon          StrWithPrefix16
	Desc          Hash
	TimeDisplayed Hash
}

type MechTDWeatherMetaData struct {
	// Fields
	WeatherID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	WeatherName Hash
	WeatherIcon StrWithPrefix16
	WeatherDesc Hash
}

type MentorCompanyMissionRewardMetaData struct {
	// Fields
	MissionID int32

	// Properties
	StudentRewardID int32
	MentorRewardID  int32
}

type MentorConst struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Name  StrWithPrefix16
	Type  StrWithPrefix16
	Value StrWithPrefix16
}

type MentorDailyMissionReward struct {
	// Fields
	MissionID int32

	// Properties
	IdRewardStudent int32
	IdRewardMentor  int32
}

type MentorEvaluate struct {
	// Fields
	EvaluateID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	EvaluateTextmap Hash
}

type MentorExamData struct {
	// Fields
	ExamID int32

	// Properties
	NextExamID      int32
	Addr1           uint32 `json:"-"`
	IdRewardStudent int32
	IdRewardMentor  int32
	RefLevel        int32
	Addr2           uint32 `json:"-"`

	// Objects
	NameExam          Hash
	IdExamMissionList []int32
}

type MentorFameLevelData struct {
	// Fields
	Level int32

	// Properties
	FameForNextLevel int32
	RewardObtain     int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`

	// Objects
	LevelTitle Hash
	LevelDesc  Hash
}

type MentorLobbyTagDisplay struct {
	// Fields
	PlanID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	IsDisplay int32

	// Objects
	PathIconStudent StrWithPrefix16
	PathIconMentor  StrWithPrefix16
}

type MentorMainSceneNotice struct {
	// Fields
	NoticeEventID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	NoticeTextMap Hash
}

type MentorRewardAttenuation struct {
	// Fields
	LevelDiff int32

	// Properties
	RatioExamStudent  float32
	RatioExamMentor   float32
	RatioDailyStudent float32
	RatioDailyMentor  float32
	AttenuationType   int32
	Addr1             uint32 `json:"-"`

	// Objects
	AttenuationTipsTextmap Hash
}

type MiniBossExamExplainInfoMetaData struct {
	// Fields
	GroupID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	OpenDate StrWithPrefix16
	Desc     StrWithPrefix16
}

type MiniBossExamHardStageDataMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TagList       []MiniBossExamHardStageDataMetaData_Tag
	MissionIDList []int32
}

type MiniBossExamHardStageDataMetaData_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	TagDesc Hash
	TagName Hash
	TagIcon StrWithPrefix16
}

type MiniMonopolyActivityMetaData struct {
	// Fields
	ActivityId int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	StartMapId       uint16
	CreditMaterialID int32
	RewardLineID     int32
	ThrowDiceScore   uint16
	DiceMaterialID   int32
	DiceSpeed        uint16
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	Addr7            uint32 `json:"-"`

	// Objects
	ActivityTitle       Hash
	ActivityDescription Hash
	MapPool             []uint16
	SectorIdList        []uint16
	ItemIDList          []uint16
	MissionList         []int32
	AvatarImagePath     StrWithPrefix16
}

type MiniMonopolyGridMetaData struct {
	// Fields
	MapId  uint16
	GridId uint8

	// Properties
	GridType       uint8
	RewardId       int32
	GridMaterialID int32
	IsKeyReward    bool
	AvatarOrient   uint8
}

type MiniMonopolyItemMetaData struct {
	// Fields
	ItemId uint16

	// Properties
	ItemType   uint8
	MaterialID int32
}

type MiniMonopolyMapMetaData struct {
	// Fields
	MapId uint16

	// Properties
	NextMapId uint16
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	MapPrefabPath StrWithPrefix16
	GridList      []uint8
}

type MiniMonopolySectorMetaData struct {
	// Fields
	SectorId uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	NumberList []uint8
	SectorText Hash
}

type MiniSkyFireInfoDesMetaData struct {
	// Fields
	InfoDesID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Date    StrWithPrefix16
	TextMap Hash
}

type MiniSkyFireInfoMetaData struct {
	// Fields
	AcitivityID int32

	// Properties
	InfoIconShow        bool
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	TicketIconShow      int32
	RewardShowTpye      int32
	GenActivityRewardID int32

	// Objects
	ShowShopIconTime  RemoteTime
	RewardShowTextMap Hash
	InfoDesList       []int32
}

type MiniSkyFireStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Index      int32
	HardType   int32
	EnterTimes int32
	MinScore   int32
	MaxScore   int32
	Addr1      uint32 `json:"-"`

	// Objects
	CoverTextmap Hash
}

type MinuteChallengeMetaData struct {
	// Fields
	ID int32

	// Properties
	StageID   int32
	MissionID int32
}

type MissionCategoryMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	CategoryTitle StrWithPrefix16
	CategoryDesc  StrWithPrefix16
	MissionList   []int32
}

type MissionData struct {
	// Fields
	Id int32

	// Properties
	Type          uint8
	SubType       uint16
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	FinishWay     uint16
	FinishParaInt int32
	Addr4         uint32 `json:"-"`
	TotalProgress int32
	RewardId      int32
	LinkType      uint16
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	PreviewTime   int32
	NoDisplay     bool
	IsNeedDisplay bool
	Priority      int32

	// Objects
	Title         Hash
	Description   Hash
	Thumb         StrWithPrefix16
	FinishParaStr StrWithPrefix16
	LinkParams    []int32
	TextmapTag    Hash
	LinkParamStr  StrWithPrefix16
}

type MissionExtraRewardMetaData struct {
	// Fields
	MissionID int32

	// Properties
	ExtraRewardID int32
}

type MissionGroupMetaData struct {
	// Fields
	GroupID int32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Reward         int32
	Addr3          uint32 `json:"-"`
	NeedMissionNum int32

	// Objects
	MissionIDList   []int32
	RPGAreaList     []int32
	AreaDisplayText Hash
}

type MissionPanelConversationMetaData struct {
	// Fields
	ActivityID     int32
	ConversationID uint16

	// Properties
	Popup_x          int16
	Popup_y          int16
	ConversationType uint16
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`

	// Objects
	ConversationTypeParaInt []int32
	ConversationTextMapList []Hash
	CharacterImagePathList  []StrWithPrefix16
}

type MissionThemeMetaData struct {
	// Fields
	ThemeID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	AvatarID         int32
	DressID          int32
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	Addr6            uint32 `json:"-"`
	LinkType         int32
	Addr7            uint32 `json:"-"`
	Addr8            uint32 `json:"-"`
	Addr9            uint32 `json:"-"`
	UpgradeCostHCoin int32
	UpgradeCostMCoin int32
	Addr10           uint32 `json:"-"`
	Addr11           uint32 `json:"-"`

	// Objects
	EndTime                 RemoteTime
	ThemeTitle              Hash
	BGPath                  StrWithPrefix16
	CameraPos               []float32
	ImagePath               StrWithPrefix16
	IconPath                StrWithPrefix16
	LinkParams              []int32
	LinkParamStr            StrWithPrefix16
	MissionIDList           []int32
	UpgradeCostMaterialList []MissionThemeMetaData_CostMaterial
	UpgradeCostProductName  StrWithPrefix16
}

type MissionThemeMetaData_CostMaterial struct {
	// Fields
	Count      int32
	MaterialID int32
}

type MonopolyBuffInfoMetaData struct {
	// Fields
	BuffTextID uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	SpecialBuffText StrWithPrefix16
}

type MonopolyBuffMetaData struct {
	// Fields
	BuffID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	BuffName     StrWithPrefix16
	BuffDesc     StrWithPrefix16
	BuffIcon     StrWithPrefix16
	Para         StrWithPrefix16
	PetJoinText  StrWithPrefix16
	PetLeaveText StrWithPrefix16
}

type MonopolyDiceMetaData struct {
	// Fields
	DiceID uint16

	// Properties
	DiceItemID   int32
	DiceMaxValue int32
}

type MonopolyGuessMetaData struct {
	// Fields
	GuessID int32

	// Properties
	RewardID     int32
	DrawRewardID int32
	LoseRewardID int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	ChatTextLeft  []Hash
	ChatTextRight []Hash
}

type MonopolyLevelInfoMetaData struct {
	// Fields
	BuffLevelID uint16
	MonopolyID  uint16

	// Properties
	SpecialText1 uint16
	HPValue      int32
	SpecialText2 uint16
	ATKValue     int32
}

type MonopolyLotteryMetaData struct {
	// Fields
	LotteryID uint16

	// Properties
	Addr1        uint32 `json:"-"`
	LotteryTime  int32
	RewardID     int32
	LoseRewardID int32

	// Objects
	LotteryIcon StrWithPrefix16
}

type MonsterCardActivityDataMetaData struct {
	// Fields
	ActivityID uint32

	// Properties
	SweepCostMaterialID   int32
	DailyRewardLimitTimes int32
	Addr1                 uint32 `json:"-"`
	InitCardCostLimit     int32
	MaxCardCostLimit      int32
	ExtraCardCostMaterial int32
	CardNumLimit          uint8
	SecretAreaID          uint32
	TowerAreaID           uint32
	ActivityMedalID       int32
	Addr2                 uint32 `json:"-"`

	// Objects
	CardCollectionMission []int32
	PlayerMonsterCardList []uint32
}

type MonsterCardBuffDisplayMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	Buffname     StrWithPrefix16
	BuffIconPath StrWithPrefix16
}

type MonsterCardLevelLimitMetaData struct {
	// Fields
	ActivityID uint32
	Level      uint8

	// Properties
	LimitType uint8
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	Param         StrWithPrefix16
	LimitTextHint Hash
}

type MonsterCardLevelMetaData struct {
	// Fields
	UniqueID uint32
	Level    uint8

	// Properties
	Addr1               uint32 `json:"-"`
	Atk                 float32
	Def                 float32
	HP                  float32
	ResistElementAttack float32
	Addr2               uint32 `json:"-"`

	// Objects
	LevelUpCost MonsterCardLevelMetaData_Cost
	SkillList   []uint32
}

type MonsterCardLevelMetaData_Cost struct {
	// Fields
	Amount     int32
	MaterialID int32
}

type MonsterCardMetaData struct {
	// Fields
	UniqueID uint32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Scale       float32
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	CardCost    int32
	MonsterSize uint8
	Addr6       uint32 `json:"-"`
	InitThreat  float32
	AttackRange uint8
	Profession  uint8
	Addr7       uint32 `json:"-"`
	Addr8       uint32 `json:"-"`
	Addr9       uint32 `json:"-"`
	Index       int32

	// Objects
	CardPicPath     StrWithPrefix16
	CardLongPicPath StrWithPrefix16
	PrefabPath      StrWithPrefix16
	Position        []float32
	Rotation        []float32
	CardSkill       []uint32
	ProfessionName  Hash
	ProfessionIcon  StrWithPrefix16
	CardInformation Hash
}

type MonsterCardSkillMetaData struct {
	// Fields
	SkillID uint32

	// Properties
	SkillType    uint8
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	IsAddAbility bool
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`

	// Objects
	AbilityName     StrWithPrefix16
	AbilityParaList []MonsterCardSkillMetaData_SkillParam
	Icon            StrWithPrefix16
	Name            Hash
	Desc            Hash
	UnlockDesc      Hash
}

type MonsterCardSkillMetaData_SkillParam struct {
	// Fields
	Param float32
	Addr1 uint32 `json:"-"`

	// Objects
	Name StrWithPrefix16
}

type MonsterCardStageMetaData struct {
	// Fields
	StageID uint32

	// Properties
	StageType             uint8
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	StageTime             float32
	Addr5                 uint32 `json:"-"`
	IsSweep               bool
	SweepNeedChallengeNum uint8
	SweepRewardID         int32
	IsNextStage           bool

	// Objects
	StageScene  StrWithPrefix16
	Wave1       []MonsterCardStageMetaData_MonsterSlot
	Wave2       []MonsterCardStageMetaData_MonsterSlot
	Wave3       []MonsterCardStageMetaData_MonsterSlot
	StageEffect []MonsterCardStageMetaData_Effect
}

type MonsterCardStageMetaData_MonsterSlot struct {
	// Fields
	Level    uint8
	Slot     uint8
	Star     uint8
	UniqueID uint32
}

type MonsterCardStageMetaData_Effect struct {
	// Fields
	EffectSide uint8
	EffectID   uint32
}

type MonsterCardStarMetaData struct {
	// Fields
	UniqueID uint32
	Star     uint8

	// Properties
	LevelCanStarUp           uint8
	LevelLimit               uint8
	StarUpCost               uint32
	AtkRatio                 float32
	DefRatio                 float32
	HPRatio                  float32
	ResistElementAttackRatio float32
	Addr1                    uint32 `json:"-"`

	// Objects
	SkillList []uint32
}

type MonsterCardTalentBookMetaData struct {
	// Fields
	MaterialID int32
}

type MonsterCardTowerMetaData struct {
	// Fields
	StageID int32

	// Properties
	FloorID uint8
	IsHard  bool
	Medal   uint8
}

type MonsterConfigMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Properties
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	HP        float32
	Attack    float32
	Defense   float32
	Nature    int32
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`
	Addr7     uint32 `json:"-"`
	EliteType int32
	Addr8     uint32 `json:"-"`

	// Objects
	MonsterName  StrWithPrefix16
	TypeName     StrWithPrefix16
	CategoryName StrWithPrefix16
	SubTypeName  StrWithPrefix16
	ConfigFile   StrWithPrefix16
	ConfigType   StrWithPrefix16
	AIName       StrWithPrefix16
	DisplayTitle StrWithPrefix16
}

type MonsterResistanceMetaData struct {
	// Fields
	UniqueMonsterID uint32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`

	// Properties
	DamageResist            float32
	FireDamageResist        float32
	IceDamageResist         float32
	ThunderDamageResist     float32
	ResistAllDamageRatio    float32
	AllDamageTakeRatio      float32
	AllDamageTakeRatio2     float32
	DamageTakeRatio         float32
	IceAttackTakeRatio      float32
	FireAttackTakeRatio     float32
	ThunderAttackTakeRatio  float32
	Weak                    float32
	Fragile                 float32
	Bleed                   float32
	Burn                    float32
	Poisoned                float32
	Stun                    float32
	Paralyze                float32
	Frozen                  float32
	MoveSpeedDown           float32
	AttackSpeedDown         float32
	WitchTimeSlowed         float32
	WeakDuration            float32
	FragileDuration         float32
	BleedDuration           float32
	BurnDuration            float32
	PoisonedDuration        float32
	StunDuration            float32
	ParalyzeDuration        float32
	FrozenDuration          float32
	MoveSpeedDownDuration   float32
	AttackSpeedDownDuration float32
	WitchTimeSlowedDuration float32

	// Objects
	MonsterName StrWithPrefix16
	TypeName    StrWithPrefix16
}

type MonsterWikiDataMetaData struct {
	// Fields
	MonsterWikiID int32

	// Properties
	Page        uint16
	MonsterID   int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	UnlockLevel uint16

	// Objects
	Name            Hash
	IconPath        StrWithPrefix16
	UnlockStageList []int32
	UnlockDec       Hash
}

type MonsterWikiDescConfigMetaData struct {
	// Fields
	DisplayDescID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	DisplayTitle StrWithPrefix16
	BGPath       StrWithPrefix16
	IconPath     StrWithPrefix16
	DisplayDesc  StrWithPrefix16
	MessageList  []StrWithPrefix16
	AuthorList   []StrWithPrefix16
}

type MonsterWikiPageMetaData struct {
	// Fields
	Page uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Show  bool

	// Objects
	PageName Hash
	Pic      StrWithPrefix16
}

type MPBuffMetaData struct {
	// Fields
	MPBuffID uint32

	// Properties
	BuffType uint8
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	Icon         StrWithPrefix16
	Name         StrWithPrefix16
	Desc         StrWithPrefix16
	DescInDetail StrWithPrefix16
}

type MPLevelData struct {
	// Fields
	MpLevel int32

	// Properties
	Addr1 uint32 `json:"-"`
	MpExp int32

	// Objects
	Name Hash
}

type MPListMetaData struct {
	// Fields
	ListMainID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Title Hash
}

type MPPveActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	RewardShowID       uint8
	IsRankShow         bool
	IsAllowSelectSkill bool
}

type MPPveActivityStageGroupMetaData struct {
	// Fields
	StageGroupID int32

	// Properties
	StageType uint8
}

type MPRaidActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	SpecialReward     int32
	RankID            int32
	ResetSwitch       bool
	ResetTimes        int32
	Addr1             uint32 `json:"-"`
	RefreshCostType   int32
	RefreshCostBase   int32
	RefreshCostOffset int32
	Difficulty        int32
	UnlockActivityID  int32
	UnlockTeamLv      int32
	UnlockMpLv        int32
	ActRewardShow     int32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	BalanceModeType   int32

	// Objects
	WeeklyRefreshDay []int32
	RecTeamLv        []int32
	MapBG            StrWithPrefix16
}

type MPRaidActMetaData struct {
	// Fields
	ActID int32

	// Properties
	ActivityID int32
	Addr1      uint32 `json:"-"`
	ActMission int32
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`

	// Objects
	LevelList            []int32
	LevelPanelPrefabPath StrWithPrefix16
	ActNameText          Hash
	ActTitleText         Hash
}

type MPRaidRankMetaData struct {
	// Fields
	RankID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	RankBG       StrWithPrefix16
	RankTab      StrWithPrefix16
	RankTabTitle Hash
}

type MPRaidRankRewardMetaData struct {
	// Fields
	Id int32

	// Properties
	RaidRank int32
	RewardID int32
}

type MPRaidScoreRewardMetaData struct {
	// Fields
	Id int32

	// Properties
	RaidScore int32
	RewardID  int32
}

type MPRaidSeriesMetaData struct {
	// Fields
	SeriesID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	UiAssetPath StrWithPrefix16
	CgExtraKey  StrWithPrefix16
}

type MPSkillsData struct {
	// Fields
	MpSkillsId int32

	// Properties
	TypeOfSkill   int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	UnlockMPLevel int32
	Proportion1   float32
	Proportion2   float32
	Proportion3   float32
	Proportion4   float32
	MaxSkillLevel int32

	// Objects
	MpSkillName Hash
	MpSkillDesc Hash
}

type MPStageBuffMetaData struct {
	// Fields
	LevelID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MPBuffIDs []uint32
}

type MPStageMetaData struct {
	// Fields
	LevelID int32

	// Properties
	HostExtraStamina    uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	OnlyCostHostTimes   bool
	UnlockMPLevel       uint8
	WildPlayerExitType  uint8
	ExitButtonType      uint8
	Addr3               uint32 `json:"-"`
	MatchType           uint8
	MaxAssistTimes      uint8
	EnterTimeType       uint8
	EnableGrandKeyBuff  bool
	EnableElf           bool
	ExitSettleType      uint8
	MPTeamSkillType     uint8
	RevivalDuration     int32
	MPReviveTimes       int32
	MemberIdentifyLimit int32
	MPMode              uint8
	MinPlayer           uint8
	MaxPlayer           uint8

	// Objects
	HostExtraDisplayDropList   []int32
	MemberExtraDisplayDropList []int32
	StageMissionList           []int32
}

type MPStagePlayerLevelDropMetaData struct {
	// Fields
	StageID        int32
	MinPlayerLevel int32

	// Properties
	MaxPlayerLevel int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`

	// Objects
	HostExtraDisplayDropList   []int32
	MemberExtraDisplayDropList []int32
}

type MPTagDetectMetaData struct {
	// Fields
	LevelId int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TagList    []int32
	TagNumList []int32
}

type MPTeamSkillMetaData struct {
	// Fields
	AvatarId        int32
	MPTeamSkillType int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	MPTeamSkillName Hash
	MPTeamSkillDesc Hash
	MPTeamSkillIcon StrWithPrefix16
}

type MPTrophyMetaData struct {
	// Fields
	LevelGroupID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TrophyScore []int32
}

type MultiPlatforms_UserInfoMetaData struct {
	// Fields
	UserType    uint8
	AccountType uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	UserLabel Hash
}

type MuseumMetaData struct {
	// Fields
	MuseumID uint8

	// Properties
	ActivityID      uint8
	ActivityPanelID uint16
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	ShopId          int32

	// Objects
	Image           StrWithPrefix16
	ShopGoodsIdList []int32
}

type NatureCounterMetaData struct {
	// Fields
	PlayerNatureType int32
	MonsterNature    int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	NatureCounterTitle Hash
	NatureCounterDesc  Hash
}

type NetworkErrCodeMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Properties
	Addr3 uint32 `json:"-"`

	// Objects
	ErrType   StrWithPrefix16
	RetCode   StrWithPrefix16
	TextMapID Hash
}

type NewbieActivityPanelMetaData struct {
	// Fields
	ScheduleID int32
	ActivityID int32

	// Properties
	ActivityType    int32
	ActivityPanelID uint16
	GroupID         int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	CurrencyHide    bool
	ShowLevel       int32
	OpenLevel       int32
	SortID          int32

	// Objects
	TabTitle     Hash
	SubTabTitle  Hash
	TitleText    Hash
	DescText     Hash
	InfoText     Hash
	StringParam  []StrWithPrefix16
	CurrencyList []uint32
}

type NewbieActivityScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	StageMissionID              int32
	LoginRewardSchedule         int32
	NewbiePanelSchedule         int32
	GrowUpActivityID            int32
	WorldMapNewbieChallengeHide bool
	CgGroupID                   int32
	Addr1                       uint32 `json:"-"`
	Addr2                       uint32 `json:"-"`
	DevelopRewardLineID         uint16
	NewbieLevelRush             int32
	Addr3                       uint32 `json:"-"`
	Addr4                       uint32 `json:"-"`
	Addr5                       uint32 `json:"-"`

	// Objects
	NewbieShopTypeList        []int32
	NewbieSubMallID           []int32
	HighlightMaterialList     []int32
	HighlightResourceTypeList []uint8
	NewbieLoginPostIDList     []uint16
}

type NewbieActivityStageMissionMetaData struct {
	// Fields
	ID         int32
	StageLevel int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	MissionList   []int32
	BannerImgPath StrWithPrefix16
}

type NewbieAvatarGuideData struct {
	// Fields
	AvatarID         int32
	IsAvatarArtifact bool

	// Properties
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	DefaultRecommandType int32
	Addr6                uint32 `json:"-"`

	// Objects
	AvatarBriefText  Hash
	AvatarTraitsText Hash
	RadarParams      []int32
	GroupAdvsID      []int32
	NewbieGuideList  []int32
	GachaVideoLink   StrWithPrefix16
}

type NewbieBattleBuff struct {
	// Fields
	MinLevel int32
	MaxLevel int32

	// Properties
	EvadeEffectCDReduce float32
	Addr1               uint32 `json:"-"`
	ExtraEvadePropID    int32
	Addr2               uint32 `json:"-"`
	EvadePropParam1     float32
	EvadePropParam2     float32
	EvadePropParam3     float32
	SpBuffPropID        int32
	Addr3               uint32 `json:"-"`
	SpBuffPropParam1    float32
	SpBuffPropParam2    float32

	// Objects
	EvadeEffectCDReduceText Hash
	ExtraEvadePropText      Hash
	SpBuffPropText          Hash
}

type NewbieCumulativeLotteryMetaData struct {
	// Fields
	ID uint8

	// Properties
	KeyRewardMission int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`

	// Objects
	KeyRewardImage      StrWithPrefix16
	NormalMissionIDList []StrWithPrefix16
}

type NewbieDevelopRewardPanelMetaData struct {
	// Fields
	ID uint16

	// Properties
	OpenLevel uint8
	Addr1     uint32 `json:"-"`

	// Objects
	SpecialRewardList []NewbieDevelopRewardPanelMetaData_SpecialReward
}

type NewbieDevelopRewardPanelMetaData_SpecialReward struct {
	// Fields
	ShowType uint8
	RankID   uint16
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	Tag       Hash
	ImagePath StrWithPrefix16
	ParamStr  StrWithPrefix16
}

type NewbieDormQAvatarMetaData struct {
	// Fields
	ID int32

	// Properties
	IsOpen int32
	Type   int32
	Addr1  uint32 `json:"-"`

	// Objects
	BigImgPath StrWithPrefix16
}

type NewbieEquipAdvData struct {
	// Fields
	EquipAdvID int32

	// Properties
	WeaponAdvID    int32
	Type           int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	StigmataAdv1ID int32
	StigmataAdv2ID int32
	StigmataAdv3ID int32
	Addr4          uint32 `json:"-"`

	// Objects
	Title           Hash
	CopyName        Hash
	WeaponAdvText   Hash
	StigmataAdvText Hash
}

type NewbieEquipTypeMetaData struct {
	// Fields
	Type int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Desc Hash
}

type NewbieFirstPaymentMetaData struct {
	// Fields
	ID uint16

	// Properties
	OpenLevel  uint8
	MissionID  int32
	MaterialID int32
	Addr1      uint32 `json:"-"`

	// Objects
	RewardIDList []int32
}

type NewbieFirstPaymentRewardMetaData struct {
	// Fields
	RewardID int32

	// Properties
	AvatarCardID int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	Name Hash
	Tag  Hash
}

type NewbieGlobalBuff struct {
	// Fields
	MaxLevel int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	BuffEffectList []int32
}

type NewbieGroupAdvData struct {
	// Fields
	GroupAdvID int32

	// Properties
	GroupAvatarID int32
	RecommendStar int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	RecommendText      Hash
	RecommendSkillIcon StrWithPrefix16
}

type NewbieGrowUpMetaData struct {
	// Fields
	ScheduleID uint32
	Rank       uint8

	// Properties
	UseOpenLevel       bool
	OpenLevel          uint8
	DesiredLevel       uint8
	IsOld              bool
	Addr1              uint32 `json:"-"`
	RankMissionGroupID uint32
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`

	// Objects
	RankName            Hash
	MissionChainDisplay StrWithPrefix16
	PreDisplayDesc      Hash
}

type NewbieGuideDialogueData struct {
	// Fields
	DialogueID int32

	// Properties
	Tag        int32
	SubTag     int32
	RemainTime int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	DialogueText Hash
	AiExpression StrWithPrefix16
}

type NewbieGuideFaqData struct {
	// Fields
	FaqID int32

	// Properties
	LevelPhase  int32
	UnlockLevel int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Recommend   int32
	LinkType    int32
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`

	// Objects
	TitleText       Hash
	DescText        Hash
	LinkParamList   []int32
	LinkParamString StrWithPrefix16
	FaqBGIcon       StrWithPrefix16
}

type NewbieGuideFaqListData struct {
	// Fields
	FaqListID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	TitleText    Hash
	DescText     Hash
	RelatedFaqID []int32
	FaqBGIcon    StrWithPrefix16
	Icon         StrWithPrefix16
}

type NewbieGuideGrowthQuestData struct {
	// Fields
	QuestID int32

	// Properties
	LevelPhase       int32
	RelatedMissionID int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`

	// Objects
	TitleText Hash
	DescText  Hash
}

type NewbieGuideTutorialLevelData struct {
	// Fields
	TutorialID int32

	// Properties
	StageLevelID        int32
	IsNewMechanismEntry bool
}

type NewbieLevelRushMetaData struct {
	// Fields
	ID int32

	// Properties
	OpenLevel         uint8
	PurchaseCostMCoin uint16
	RewardConfigID    int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`

	// Objects
	SpecialReward NewbieLevelRushMetaData_SpecialRewardData
	Tag           Hash
}

type NewbieLevelRushMetaData_SpecialRewardData struct {
	// Fields
	IsPurchaseReward bool
	UnlockLevel      uint8
}

type NewbieLevelRushRewardMetaData struct {
	// Fields
	RewardConfigID int32
	UnlockLevel    uint8

	// Properties
	FreeRewardID     int32
	PurchaseRewardID int32
}

type NewbieLoginPostMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Priority uint8

	// Objects
	ImgPath             StrWithPrefix16
	LinkData            StrWithPrefix16
	BegineConditionList []StrWithPrefix16
	EndConditionList    []StrWithPrefix16
}

type NewbieRecurrenceLoginRewardData struct {
	// Fields
	LoginRewardID int32

	// Properties
	LoginDay     int32
	RewardID     int32
	FreeSelectID int32
	IsKeyReward  bool
	Type         int32
}

type NewbieWebsitesMetaData struct {
	// Fields
	ID int32

	// Properties
	IsOpen      int32
	JumpingType int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`

	// Objects
	BigImgPath StrWithPrefix16
	WebsiteURL StrWithPrefix16
}

type NewFeatureGuideData struct {
	// Fields
	FeatureID int32

	// Properties
	Addr1           uint32 `json:"-"`
	IsActive        bool
	Type            int32
	MinLv           int32
	MaxLv           int32
	Weight          int32
	Addr2           uint32 `json:"-"`
	FeatureLinkType int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`

	// Objects
	Desc             Hash
	UrlPath          StrWithPrefix16
	LinkParam        []int32
	FeatureBeginTime StrWithPrefix16
	FeatureEndTime   StrWithPrefix16
}

type NinjaAreaMetaData struct {
	// Fields
	AreaID uint8

	// Properties
	PreArea         uint8
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	SlotID          uint8
	Addr8           uint32 `json:"-"`
	RecommendSlotID uint8
	RecommendLv     uint8

	// Objects
	UIIndex     StrWithPrefix16
	TextUIIndex StrWithPrefix16
	UnlockTime  StrWithPrefix16
	AreaName    Hash
	AreaDesc    Hash
	BossIcon    StrWithPrefix16
	BossImg     StrWithPrefix16
	UICameraID  []float32
}

type NinjaAttrMetaData struct {
	// Fields
	AttrID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	AttrName Hash
	AttrDesc Hash
	AttrIcon StrWithPrefix16
}

type NinjaMissionCGMetaData struct {
	// Fields
	MissionID int32

	// Properties
	Addr1   uint32 `json:"-"`
	EventID uint16
	SiteID  uint16

	// Objects
	UnlockCGTime StrWithPrefix16
}

type NinjaScheduleMetaData struct {
	// Fields
	NinjaID uint8

	// Properties
	MinLevel         uint8
	MaxLevel         uint8
	CurrencyID       int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	RewardShowID     uint8
	EntryPerformID   uint16
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	TalkRoomIconShow bool
	DailyLimitShow   bool

	// Objects
	AreaList     []int32
	MissionList  []int32
	CGID         []int32
	EntryEventID []int32
	WebLink      StrWithPrefix16
}

type NinjaSiteMetaData struct {
	// Fields
	SiteID uint16

	// Properties
	AreaID          uint8
	SiteType        uint8
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	StageID         int32
	NextStageID     int32
	Addr3           uint32 `json:"-"`
	UnlockSlotLevel uint8
	PreSiteID       uint16
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`

	// Objects
	UIIndex     StrWithPrefix16
	LineUIIndex StrWithPrefix16
	ChangeText  Hash
	LockText    Hash
	PointTitle  Hash
	PointDesc   Hash
}

type NinjaSlotLevelMetaData struct {
	// Fields
	SlotID uint8
	Level  int32

	// Properties
	StrengthenExp int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	SpEffectID    uint8
	SpValue       float32

	// Objects
	AttrList  []int32
	AttrValue []int32
}

type NinjaSlotMetaData struct {
	// Fields
	SlotID uint8

	// Properties
	CostMaterialID int32
	StrengthenExp  uint8
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`

	// Objects
	OptionalEffect []int32
	OptionValue    []float32
	SpUnlockLevel  []int32
	SlotName       Hash
	SlotDesc       Hash
	SlotIcon       StrWithPrefix16
}

type NinjaSpEffectMetaData struct {
	// Fields
	SpEffectID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	SpEffectName Hash
	SpEffectDesc Hash
	SpEffectIcon StrWithPrefix16
}

type NPCLevelMetaData struct {
	// Fields
	HardLevelGroup int32
	HardLevel      int32

	// Properties
	HPRatio              float32
	ATKRatio             float32
	DEFRatio             float32
	ElementalResistRatio float32
}

type OpenWorldArea struct {
	// Fields
	AreaId int32

	// Properties
	MapId         int32
	IsStoryOnly   int32
	IsMarkOnly    bool
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	UnlockStoryId int32
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	LevelID       int32
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`

	// Objects
	AreaList       []StrWithPrefix16
	AreaRGB        StrWithPrefix16
	NameTextId     Hash
	AreaIcon       StrWithPrefix16
	SwitchPoint    StrWithPrefix16
	UnlockAreaTips Hash
	UnlockAreaDesc Hash
}

type OpenWorldBoss030MetaData struct {
	// Fields
	Phase int32

	// Properties
	Addr1         uint32 `json:"-"`
	UnlockStoryId int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`

	// Objects
	ProgressText Hash
	UnlockParts  StrWithPrefix16
	UnlockTips   Hash
}

type OpenWorldContentMetaData struct {
	// Fields
	ID    int32
	MapID int32

	// Properties
	JudgeType          int32
	ContentJudgeWeight int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	ContentRewardShow  int32
	Addr5              uint32 `json:"-"`

	// Objects
	BeginTime       StrWithPrefix16
	EndTime         StrWithPrefix16
	ContentTitle    Hash
	ContentIconPath StrWithPrefix16
	ContentDesc     Hash
}

type OpenWorldCookData struct {
	// Fields
	CookID int32

	// Properties
	UnlockQuestLv int32
	Addr1         uint32 `json:"-"`
	EffectType    int32
	Addr2         uint32 `json:"-"`
	DisplayStar   int32
	UseType       int32
	CookTab       int32
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	DailyTimes    int32
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	Addr8         uint32 `json:"-"`

	// Objects
	DisplayCookName   Hash
	EffectParam       []int32
	DisplayCookDesc   Hash
	DisplayEffectDesc Hash
	CostMaterialsList []OpenWorldCookData_IntIntPair
	IconPath          StrWithPrefix16
	CookBuffIcon      StrWithPrefix16
	CookBuffEfx       StrWithPrefix16
}

type OpenWorldCookData_IntIntPair struct {
	// Fields
	Item1 int32
	Item2 int32
}

type OpenWorldCycleActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	MapID             int32
	MinLevel          int32
	MaxLevel          int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	ContentRewardShow int32

	// Objects
	BeginTime StrWithPrefix16
	EndTime   StrWithPrefix16
}

type OpenworldCycleData struct {
	// Fields
	Cycle int32

	// Properties
	Level         int32
	Needstory     int32
	Finishreward  int32
	CycleMap      int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	EntranceScene int32
	Addr3         uint32 `json:"-"`

	// Objects
	HardLvKey    StrWithPrefix16
	CycleName    Hash
	EntranceImg1 StrWithPrefix16
}

type OpenWorldDestinyLevelDiffMetaData struct {
	// Fields
	LevelDiff int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	AbilityName StrWithPrefix16
	IconBg      StrWithPrefix16
}

type OpenWorldDestinyLineLinkMetaData struct {
	// Fields
	LineLinkID int32

	// Properties
	Type           int32
	EventID        int32
	LineLinkTypeID int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	HintDuration   float32

	// Objects
	OverrideMap []StrWithPrefix16
	HintTitle   Hash
	HintContent Hash
	HintDesc    Hash
}

type OpenWorldDynamicHardLvMetaData struct {
	// Fields
	TeamLv int32

	// Properties
	Cycle1HardLv int32
	Cycle2HardLv int32
	Cycle3HardLv int32
	Cycle4HardLv int32
	Map2HardLv   int32
}

type OpenWorldEventActivityMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BeginTime  StrWithPrefix16
	EndTime    StrWithPrefix16
	EventList  []int32
	TipsBanner StrWithPrefix16
}

type OpenWorldEventData struct {
	// Fields
	EventID int32

	// Properties
	Type          uint8
	Addr1         uint32 `json:"-"`
	JudgeType     uint8
	EventMap      uint16
	EventArea     uint16
	UnlockEvent   int32
	UnlockCycle   uint8
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	EventInitType uint16
	RefreshType   uint8
	Reward        int32

	// Objects
	Subtype          []int32
	EventNameTextMap Hash
	Area             StrWithPrefix16
	SpawnPointName   StrWithPrefix16
	PropName         StrWithPrefix16
	AreaName         StrWithPrefix16
}

type OpenWorldEventExtraDropMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	BeginTime  StrWithPrefix16
	EndTime    StrWithPrefix16
	TipsBanner StrWithPrefix16
}

type OpenWorldExplorePointMetaData struct {
	// Fields
	ExploreID int32

	// Properties
	MapID         int32
	AreaID        int32
	Type          int32
	LocationID    int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	DefaultStatus int32
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	UnlockLevel   uint8
	Rotation      float32

	// Objects
	BornPointName   StrWithPrefix16
	PointNameTextID Hash
	PointDescTextID Hash
	UIPath          []StrWithPrefix16
	PrefabPath      []StrWithPrefix16
}

type OpenWorldFunctionMetaData struct {
	// Fields
	Function int32
	MapId    int32

	// Properties
	UnlockType int32
	UnlockPara int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`

	// Objects
	UnlockTextmap Hash
	UnlockTips    Hash
	FunctionIcon  StrWithPrefix16
	FunctionName  Hash
	FunctionDes   Hash
}

type OpenWorldLocation struct {
	// Fields
	LocationID int32

	// Properties
	AreaId         int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	LocationBuffId int32
	Addr4          uint32 `json:"-"`
	LocationType   uint8

	// Objects
	LocationNameTextId Hash
	LocationDescTextId Hash
	SpawnPointName     StrWithPrefix16
	LocationIcon       StrWithPrefix16
}

type OpenWorldMap struct {
	// Fields
	MapId int32

	// Properties
	MapType           int32
	UnlockLv          int32
	UnlockStoryId     int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	QuestSlotNum      int32
	QuestSettleType   int32
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	HpRecoverInterval int32
	HpRecoverPercent  int32
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	Addr12            uint32 `json:"-"`

	// Objects
	ShowTime             StrWithPrefix16
	UnlockTime           StrWithPrefix16
	MapNameText          Hash
	MapContentText       Hash
	QuestMapUIManager    StrWithPrefix16
	SelectDailyQuestPage StrWithPrefix16
	SettlementPage       StrWithPrefix16
	ShopOpenWorldPage    StrWithPrefix16
	ShopTypeList         []int32
	MapInfoText          Hash
	QuestInfoText        Hash
	MapSelectPath        StrWithPrefix16
}

type OpenWorldOverViewData struct {
	// Fields
	MissionID int32

	// Properties
	Addr1      uint32 `json:"-"`
	ScoreRatio int32

	// Objects
	NameTextID Hash
}

type OpenWorldQuestActivityMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	MapId              int32
	Addr3              uint32 `json:"-"`
	MaxQuestInOneRound int32
	ShowRewardId       int32
	LinkType           int32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`

	// Objects
	BeginTime            StrWithPrefix16
	EndTime              StrWithPrefix16
	QuestList            []int32
	LinkParams           []int32
	ActivityTagName      StrWithPrefix16
	ActivityQuestTagName StrWithPrefix16
	TipsBanner           StrWithPrefix16
}

type OpenworldQuestBuffData struct {
	// Fields
	BuffId int32

	// Properties
	BuffType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	BuffNameStrId StrWithPrefix16
	BuffDescStrId Hash
	BuffIcon      StrWithPrefix16
}

type OpenworldQuestChallengeData struct {
	// Fields
	ChallengeId int32

	// Properties
	ProgressType int32
	Score        int32
	Addr1        uint32 `json:"-"`

	// Objects
	ParamList []int32
}

type OpenworldQuestDataCfg struct {
	// Fields
	QuestId int32

	// Properties
	QuestType       uint8
	QuestSubType    uint16
	QuestMap        int32
	LevelLockNeed   int32
	QuestArea       int32
	Addr1           uint32 `json:"-"`
	ShowDesc        bool
	Addr2           uint32 `json:"-"`
	Rarity          uint8
	Difficulty      int32
	RecTeamLv       uint8
	StaminaCost     uint8
	Addr3           uint32 `json:"-"`
	RewardId        int32
	RewardStar      int32
	SlotCost        uint8
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	LocationID      int32
	Progress        int32
	FinishWay       int32
	Addr8           uint32 `json:"-"`
	ActivityScore   int32
	IsActivityQuest bool
	Addr9           uint32 `json:"-"`

	// Objects
	QuestNameId   Hash
	QuestDescText Hash
	ChallengeList []int32
	IconPath      StrWithPrefix16
	LuaPara       []int32
	LuaFileName   StrWithPrefix16
	QuestTipsText Hash
	FinishParaStr StrWithPrefix16
	EnemyList     []int32
}

type OpenWorldQuestDeleteRule struct {
	// Fields
	FinishWay  int32
	IsActivate int32

	// Properties
	CanDelete int32
}

type OpenWorldQuestJudgeMetaData struct {
	// Fields
	JudgeLv int32
	QuestLv int32
	MapId   int32

	// Properties
	NeedScore int32
	Addr1     uint32 `json:"-"`

	// Objects
	JudgeIcon StrWithPrefix16
}

type OpenworldQuestLevelDataMetaData struct {
	// Fields
	QuestLevel int32

	// Properties
	LevelupNeedStar        int32
	Addr1                  uint32 `json:"-"`
	OWQuestLevelEffectMask int32

	// Objects
	Title Hash
}

type OpenWorldQuestMapLevelMetaData struct {
	// Fields
	QuestLevel int32
	MapId      int32

	// Properties
	CompleteReward int32
	WeeklyReward   int32
	KeyQuestID     int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`

	// Objects
	Title    Hash
	RecLevel []int32
}

type OpenWorldQuestMonsterPowerData struct {
	// Fields
	MonsterID int32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	LimitCount int32

	// Objects
	MonsterName StrWithPrefix16
	ConfigType  []StrWithPrefix16
	Power       []float32
}

type OpenWorldQuestRarityMetaData struct {
	// Fields
	QuestLevel int32
	MapId      int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	WeightGroup []OpenWorldQuestRarityMetaData_IntIntPair
}

type OpenWorldQuestRarityMetaData_IntIntPair struct {
	// Fields
	Item1 int32
	Item2 int32
}

type OpenWorldQuestRarityRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	Rarity      int32
	RewardID    int32
	StaminaCost int32
}

type OpenWorldQuestSlotData struct {
	// Fields
	SlotPosition int32

	// Properties
	StarReward    int32
	UnlockStoryID int32
}

type OpenWorldQuestTheme struct {
	// Fields
	ThemeID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	QuestThemeBuffList []int32
	ThemeNameText      Hash
	ThemeDescText      Hash
	ThemeIconPath      StrWithPrefix16
}

type OpenWorldQuestThemeSchedule struct {
	// Fields
	Step   int32
	AreaID int32

	// Properties
	ThemeID int32
	MapID   int32
}

type OpenWorldRewardUpListMetaData struct {
	// Fields
	ActivityID int32
	QuestLevel int32

	// Properties
	UpReward int32
}

type OpenWorldSection struct {
	// Fields
	SectionID int32

	// Properties
	AreaId            int32
	IsOutSide         int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	SitePanelRotation float32
	RotationOffset    float32
	RotationSign      float32
	Addr7             uint32 `json:"-"`

	// Objects
	SpawnPointAssetPath  StrWithPrefix16
	MapPrefabPath        StrWithPrefix16
	UiAssetPath          StrWithPrefix16
	MapSize              []float32
	CoordinateOffsetList []float32
	CoordinateList       []float32
	AxisMapping          []uint8
}

type OpenWorldStoryData struct {
	// Fields
	StoryID int32

	// Properties
	StorySeriesID         int32
	StorySeriesStep       int32
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	StoryType             int32
	GroupType             int32
	Cycle                 int32
	Addr3                 uint32 `json:"-"`
	UnlockMapLv           int32
	UnlockQuestLv         int32
	Addr4                 uint32 `json:"-"`
	IsUseNewCondition     bool
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`
	Addr7                 uint32 `json:"-"`
	StoryMap              int32
	StoryArea             int32
	Addr8                 uint32 `json:"-"`
	HuntRewardItem        int32
	HuntRewardItemDisplay int32
	Addr9                 uint32 `json:"-"`
	Addr10                uint32 `json:"-"`
	MaxCount              int32
	LocationID            int32
	DLCChallengeMode      bool
	IsTaskAnimation       int32
	IsHideUI              bool
	IsTutorial            bool
	Addr11                uint32 `json:"-"`

	// Objects
	StorySeriesTitle    Hash
	PreStory            []int32
	StoryStartDate      StrWithPrefix16
	ShowConditionList   StrWithPrefix16
	UnlockConditionList StrWithPrefix16
	PreviewRelateSeries []uint32
	UnlockConditionTips Hash
	Name                Hash
	Description         Hash
	Target              Hash
	PreStage            []int32
}

type OpenWorldStoryDetailMetaData struct {
	// Fields
	GroupType uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	GroupIcon        StrWithPrefix16
	InlevelGroupIcon StrWithPrefix16
	GroupTitle       Hash
}

type OpenWorldTaskConfigMetaData struct {
	// Fields
	MapID int32

	// Properties
	AutoChaseType int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	BgPath        StrWithPrefix16
	GroupTypeList []int32
}

type OpenWorldTimerMetaData struct {
	// Fields
	ID int32

	// Properties
	IntervalSecond int32
}

type OptionalBuffListMetaData struct {
	// Fields
	BuffID int32

	// Properties
	BuffType   int32
	RestrictID int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Score      int32
	GroupType  int32
	Level      int32

	// Objects
	BuffName      Hash
	BuffDesc      Hash
	BuffDescBrief Hash
	BuffIcon      StrWithPrefix16
}

type OptionalBuffPoolMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	ForceSelect bool

	// Objects
	PoolName        Hash
	BuffList        []int32
	UnlockScore     []int32
	UnlockCountDown []int32
}

type OptionalGachaDetailDataMetaData struct {
	// Fields
	GachaID int32

	// Properties
	BonusFragmentEventID int32
}

type OrderClientMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Name Hash
}

type OrderItemMetaData struct {
	// Fields
	OrderId int32

	// Properties
	ItemId       int32
	ItemNum      int32
	OrderCredits int32
}

type OverlappingCatMetaData struct {
	// Fields
	CatID int32

	// Properties
	Group int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	ImgPath     StrWithPrefix16
	TextmapName StrWithPrefix16
}

type OverlappingMetaData struct {
	// Fields
	OverlappingID int32

	// Properties
	SummonItemID  int32
	SummonItemNum int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	CatIDList    []int32
	TextMapRules StrWithPrefix16
}

type OverlappingRewardMetaData struct {
	// Fields
	Order int32

	// Properties
	CatID         int32
	OverlappingID int32
	RewardID      int32
}

type OWActivityBossMetaData struct {
	// Fields
	BossID int32

	// Properties
	Addr1          uint32 `json:"-"`
	Hardlevel      int32
	MonsterHp      int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	StaminaCost    int32
	ClueProgress   int32
	ClueTimeLimit  int32
	ClueBuffID     int32
	MutationBuffID int32
	ActivityExp    int32
	BossReward     int32
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	StageID        int32
	MPDamageRatio  float32
	StageMonsterID int32

	// Objects
	MonsterIdList     []int32
	ClueEventIdList   []int32
	ClueLocationList  []int32
	ClueTimeLimitTips Hash
	BossName          Hash
	BossDesc          Hash
	BossLocationList  []int32
	ImagePath         StrWithPrefix16
}

type OWActivityBossRatingMetaData struct {
	// Fields
	RatingID int32

	// Properties
	TimeLimit    int32
	RatingReward int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`

	// Objects
	RatingName         Hash
	RewardBoxPath      StrWithPrefix16
	RewardBoxAnimation StrWithPrefix16
}

type OWActivityClueMetaData struct {
	// Fields
	EventID int32

	// Properties
	ClueProgressReward int32
	ActivityExpReward  int32
}

type OWActivityLevelMetaData struct {
	// Fields
	Level        int32
	OWActivityID int32

	// Properties
	NeedExp     int32
	LevelReward int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	LevelNameText Hash
	LevelDescText Hash
	LevelIconPath StrWithPrefix16
}

type OWActivityPhaseMetaData struct {
	// Fields
	PhaseID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	BeginTime  StrWithPrefix16
	EndTime    StrWithPrefix16
	ParaString []int32
}

type OWActivityScheduleMetaData struct {
	// Fields
	OWActivityID int32
	MapID        int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	ThemeBuffPara int32
	Addr7         uint32 `json:"-"`
	Addr8         uint32 `json:"-"`

	// Objects
	BeginTime       StrWithPrefix16
	EndTime         StrWithPrefix16
	PhaseIDList     []int32
	BigImgPath      StrWithPrefix16
	RatingIDList    []int32
	ThemeAvatarList []int32
	GroupName       Hash
	GroupIcon       StrWithPrefix16
}

type OWAvatarActivityDataMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	MapID                   uint16
	VAGroup                 uint16
	RestrictAvatarID        uint16
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	TutorialActivityID      int32
	ChallengeActivityID     int32
	QuestDailyLimit         uint8
	QuestRefreshLimit       uint8
	QuestCostStamina        uint8
	QuestTotalLimit         uint8
	LevelMaterialID         int32
	Addr3                   uint32 `json:"-"`
	GenActivityRewardShowID uint8
	Addr4                   uint32 `json:"-"`
	Addr5                   uint32 `json:"-"`
	Addr6                   uint32 `json:"-"`
	Addr7                   uint32 `json:"-"`
	Addr8                   uint32 `json:"-"`
	GachaID                 int32
	ExploreAreaID           uint16
	Addr9                   uint32 `json:"-"`
	Addr10                  uint32 `json:"-"`
	PerformID               int16
	Addr11                  uint32 `json:"-"`

	// Objects
	BanSelfAvatarIDList []int32
	UIManager           StrWithPrefix16
	CurrencyLimit       []OWAvatarActivityDataMetaData_CurrencyLimitItem
	CultivateImage      StrWithPrefix16
	ShowMaterialList    []int32
	ShopLink            []int32
	WebLink             StrWithPrefix16
	GachaLink           StrWithPrefix16
	ImageInfo           []OWAvatarActivityDataMetaData_Tutorial
	JsonCfg             StrWithPrefix16
	MissionList         []int32
}

type OWAvatarActivityDataMetaData_CurrencyLimitItem struct {
	// Fields
	CurrencyID int32
	Max        int32
}

type OWAvatarActivityDataMetaData_Tutorial struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Content StrWithPrefix16
	Sprite  StrWithPrefix16
	Title   StrWithPrefix16
}

type OWAvatarActivityFilesData struct {
	// Fields
	FileID uint32

	// Properties
	ActivityID uint32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`

	// Objects
	FileTitle    Hash
	FileText     Hash
	FileLockText Hash
	FileLockHint Hash
}

type OWAvatarActivityLevelMetaData struct {
	// Fields
	ActivityLevel uint8
	ActivityID    uint16

	// Properties
	NeedNum    int32
	ClientType uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	UnlockAvatarShowList []uint16
	UnlockTalentShowList []OWAvatarActivityLevelMetaData_TalentKey
}

type OWAvatarActivityLevelMetaData_TalentKey struct {
	// Fields
	TalentID    int32
	TalentLevel int32
}

type OWAvatarActivityQuestMetaData struct {
	// Fields
	QuestId int32

	// Properties
	ActivityID uint16
	Addr1      uint32 `json:"-"`

	// Objects
	UnlockHint Hash
}

type OWAvatarActivityScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	ActivityID int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
}

type OWAvatarCultivateDataMetaData struct {
	// Fields
	CultivateID int32

	// Properties
	Type       uint8
	MaxLevel   uint8
	ActivityID uint16
}

type OWAvatarCultivateLevelMetaData struct {
	// Fields
	CultivateID int32
	Level       uint8

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	VirtualAvatarID int32
	Addr8           uint32 `json:"-"`
	VAWeaponID      int32

	// Objects
	ConditionList     StrWithPrefix16
	CostContent       []OWAvatarCultivateLevelMetaData_CostContentItem
	AvatarImage       StrWithPrefix16
	StigmataImage     StrWithPrefix16
	WeaponImage       StrWithPrefix16
	AvatarDetailImage StrWithPrefix16
	WeaponDetailImage StrWithPrefix16
	VAStigmataIDList  []int32
}

type OWAvatarCultivateLevelMetaData_CostContentItem struct {
	// Fields
	MaterialID int32
	Num        int32
}

type OWAvatarUnlockMetaData struct {
	// Fields
	AvatarID   uint16
	ActivityID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	AvatarIconPath StrWithPrefix16
	Desc           StrWithPrefix16
	ConditionList  StrWithPrefix16
}

type OWEndlessAreaDetailMetaData struct {
	// Fields
	AreaId int32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	AreaBuffId int32

	// Objects
	AreaNameTextId StrWithPrefix16
	AreaDescTextId Hash
	AreaIcon       StrWithPrefix16
}

type OWEndlessAreaScoreConfig struct {
	// Fields
	AreaScoreConfigID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Type  uint8
	Addr3 uint32 `json:"-"`

	// Objects
	MonsterScoreConfig []StrWithPrefix16
	BossScoreConfig    []int32
	SplitBattles       OWEndlessAreaScoreConfig_MonsterGroup
}

type OWEndlessAreaScoreConfig_MonsterGroup struct {
	// Fields
	GroupScoreId int32
	IsBoss       bool
	LayerId      int32
}

type OWEndlessBattleAreaMetaData struct {
	// Fields
	BattleID int32

	// Properties
	Addr1      uint32 `json:"-"`
	AreaId     int32
	TrapMaxNum int32
	Addr2      uint32 `json:"-"`

	// Objects
	BattleArea     StrWithPrefix16
	BattleAreaName Hash
}

type OWEndlessBattleConfig struct {
	// Fields
	BattleConfigID    int32
	EndlessConfigType uint8

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	MapID             int32
	AreaScoreConfigID int32

	// Objects
	EnvironmentIDList []int32
	AreaIDList        []int32
	BossAreaIDList    []int32
}

type OWEndlessBossData struct {
	// Fields
	BossConfigID int32

	// Properties
	MonsterTypeID int32
	Addr1         uint32 `json:"-"`
	BossLevel     int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`

	// Objects
	GroupLevelRange   []int32
	UniqueIDList      []int32
	LuaFilePath       StrWithPrefix16
	BOSSDisplayIDList []int32
}

type OWEndlessBuff struct {
	// Fields
	BuffID int32

	// Properties
	Addr1    uint32 `json:"-"`
	BuffType int32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	BuffName        StrWithPrefix16
	Effect          StrWithPrefix16
	InLevelIconPath StrWithPrefix16
}

type OWEndlessEnvironmentMeta struct {
	// Fields
	EnvironmentID int32

	// Properties
	MonsterTypeID int32
	BuffID        int32
	ExtraBuffID   int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`

	// Objects
	EnvironmentNameText Hash
	EnvironmentDescText Hash
	EnvironmentIconPath StrWithPrefix16
	UpTagList           []OWEndlessEnvironmentMeta_Tag
	DownTagList         []OWEndlessEnvironmentMeta_Tag
}

type OWEndlessEnvironmentMeta_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type OWEndlessGroupMeta struct {
	// Fields
	GroupLevel int32
	Type       uint8

	// Properties
	PlayerGroup    int32
	Addr1          uint32 `json:"-"`
	GroupHardLevel int32
	EliteRatio     float32
	ExtraScore     int32
	ExtraScoreTime int32
	MaxCapacity    int32
	PromoteRank    int32
	DemoteRank     int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	MinSkipLevel   uint8
	MaxSkipLevel   uint8
	SkipLevelTime  uint16

	// Objects
	GroupName          Hash
	PromoteDesc        Hash
	NormalDesc         Hash
	DemoteDesc         Hash
	UnSelectColor      StrWithPrefix16
	BgColor            StrWithPrefix16
	UnselectPrefabPath StrWithPrefix16
	SelectPrefabPath   StrWithPrefix16
	UnopenPrefabPath   StrWithPrefix16
}

type OWEndlessInvadeBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	BuffCondition   int32
	ConditionVarInt int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`

	// Objects
	ConditionVarList []int32
	BuffName         StrWithPrefix16
	BuffDesc         Hash
	BuffIconPath     StrWithPrefix16
	ConditionDesc    Hash
}

type OWEndlessInvadeMetaData struct {
	// Fields
	InvadeID int32

	// Properties
	Addr1       uint32 `json:"-"`
	InvadeScore int32
	Addr2       uint32 `json:"-"`

	// Objects
	BossConfigIDList []OWEndlessInvadeMetaData_IDPointData
	BuffList         []OWEndlessInvadeMetaData_IDPointData
}

type OWEndlessInvadeMetaData_IDPointData struct {
	// Fields
	ID    int32
	Point int32
}

type OWEndlessItemEffectMetaData struct {
	// Fields
	TypeID int32

	// Properties
	LimitType    int32
	BagAmount    int32
	EffectAmount int32
}

type OWEndlessItemMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	IndicatorOffset float32
	Rarity          int32
	Type            int32
	MaxUseRange     int32
	MinUseRange     int32
	ParamScore      int32
	RelateBuff      int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`
	Addr10          uint32 `json:"-"`

	// Objects
	Name             Hash
	PropName         StrWithPrefix16
	ParamStr         StrWithPrefix16
	IconPath         StrWithPrefix16
	SmallIconPath    StrWithPrefix16
	DescriptionShort Hash
	Description      Hash
	BuffDesc         Hash
	DebuffDesc       Hash
	ShowHint         StrWithPrefix16
}

type OWEndlessMonsterData struct {
	// Fields
	MonsterID int32

	// Properties
	MonsterTypeID uint8
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	UniqueID      uint16
	Nature        uint8
	Score         uint16
	Weight        uint8
	BossRank      uint8

	// Objects
	InitTypeList    []uint16
	GroupLevelRange []uint8
	Name            StrWithPrefix16
	ConfigType      StrWithPrefix16
}

type OWEndlessMonsterGroupScore struct {
	// Fields
	MonsterGroupID int32

	// Properties
	Score           uint16
	MonsterInitType int32
	Addr1           uint32 `json:"-"`
	RandomType      uint8
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`

	// Objects
	ExtraInitTypes          []StrWithPrefix16
	DifficultyDescription   Hash
	ExtraInitTypesProcessed []OWEndlessMonsterGroupScore_InitEntry
}

type OWEndlessMonsterGroupScore_InitEntry struct {
	// Fields
	InitType  int32
	Weight    int32
	WeightSum int32
}

type OWEndlessMonsterInitTypeMetaData struct {
	// Fields
	MonsterInitTypeID int32

	// Properties
	MonsterScoreRatio float32
	Addr1             uint32 `json:"-"`
	Type              int32
	Addr2             uint32 `json:"-"`
	IsNegate          bool
	Addr3             uint32 `json:"-"`

	// Objects
	InitTypeDesc      Hash
	WaveList          []int32
	EnvironmentIDList []int32
}

type OWEndlessPlayerBaseRewardMetaData struct {
	// Fields
	GroupLevel int32
	Type       uint8
	MinScore   int32

	// Properties
	RewardID int32
}

type OWEndlessPlayerGroupMeta struct {
	// Fields
	PlayerGroup int32
	Type        uint8

	// Properties
	MinPlayerLevel int32
	MaxPlayerLevel int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`

	// Objects
	Icon StrWithPrefix16
	Name Hash
}

type OWEndlessRewardConfigMetaData struct {
	// Fields
	ConfigID   uint32
	GroupLevel uint32

	// Properties
	PrototeRewardID uint32
	NormalRewardID  uint32
	DemoteRewardID  uint32
}

type OWEndlessScheduleMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	RewardConfig uint32

	// Objects
	StartTime RemoteTime
}

type OWEndlessSingleFloorDataMetaData struct {
	// Fields
	ActivityID int32
	StageID    int32
	Floor      uint8

	// Properties
	Score            int32
	ShowReset        bool
	ResetFloor       uint8
	MonsterGroupID   int32
	HardLevel        uint16
	ResetCostStamina uint8
	Addr1            uint32 `json:"-"`

	// Objects
	ResetCostMaterialList []OWEndlessSingleFloorDataMetaData_IntIntPair
}

type OWEndlessSingleFloorDataMetaData_IntIntPair struct {
	// Fields
	Item1 int32
	Item2 int32
}

type OWEndlessSingleModeStageMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	StageID  int32
	Type     int32
	Addr1    uint32 `json:"-"`
	MapID    int32
	Addr2    uint32 `json:"-"`
	MaxScore uint16
	Addr3    uint32 `json:"-"`

	// Objects
	EnvironmentID               []StrWithPrefix16
	AreaIDList                  []int32
	EnvironmentIDArrayProcessed []OWEndlessSingleModeStageMetaData_EnvironmentEntry
}

type OWEndlessSingleModeStageMetaData_EnvironmentEntry struct {
	// Fields
	EnvironmentID int32
	Weight        int32
	WeightSum     int32
}

type OWEndlessSingleMonsterGroupMetaData struct {
	// Fields
	ActivityID int32
	BattleID   int32

	// Properties
	Addr1          uint32 `json:"-"`
	AreaID         int32
	MonsterGroupID int32
	IsBoss         bool
	MaxScore       int32

	// Objects
	BattleArea StrWithPrefix16
}

type OWEventDisplayMetaData struct {
	// Fields
	EventID int32

	// Properties
	Type  uint8
	Level int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	EventName     StrWithPrefix16
	EventLockHint Hash
	EventImage    StrWithPrefix16
}

type OWHuntActivityDataMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	RewardScheduleID int32

	// Objects
	MapIDList            []int32
	MissionIdList        []int32
	BGPicPath            StrWithPrefix16
	TalentBGPicPath      StrWithPrefix16
	CollectionsBGPicPath StrWithPrefix16
}

type OWHuntActivityFloorMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	MapID     int32
	Floor     uint8
	CurFloor  uint8
	NextFloor uint8
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`

	// Objects
	SpawnPointName StrWithPrefix16
	CurFloorArea   StrWithPrefix16
	NextFloorArea  StrWithPrefix16
}

type OWHuntActivityHunterMetaData struct {
	// Fields
	HunterID  int32
	Hardlevel uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	HunterPicPath  StrWithPrefix16
	HunterName     Hash
	HunterDes1     Hash
	HunterDes2     Hash
	WeaknessIDList []uint16
}

type OWHuntActivityMachineEventMetaData struct {
	// Fields
	MachineID int32

	// Properties
	EventID int32
	MapID   int32
}

type OWHuntActivityMachineMetaData struct {
	// Fields
	MachineID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	MachineIcon StrWithPrefix16
	MachineName Hash
	MachineDes  Hash
}

type OWHuntActivityMapDataMetaData struct {
	// Fields
	MapID int32

	// Properties
	PreMapId           int32
	SectionID          int32
	FloorsNum          uint8
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	OpenPremise        uint8
	StrongholdMaxTimes uint8
	QuestDailyTimes    uint8
	QuestMaxTimes      uint8
	Addr4              uint32 `json:"-"`
	ForceAreaEventID   int32
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	ForceAreaFloor     uint8
	BornPointFloor     uint8
	HunterID           int32
	HunterEventID      int32
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`

	// Objects
	OpenWorldName1     Hash
	OpenWorldName2     Hash
	OpenWorldPicPath   StrWithPrefix16
	QuestIDList        []int32
	ForceAreaPicPath   StrWithPrefix16
	ForceAreaFuncParam []StrWithPrefix16
	MachineShowIDList  []int32
	InitialTalentList  []OWHuntActivityMapDataMetaData_InitialTalentData
}

type OWHuntActivityMapDataMetaData_InitialTalentData struct {
	// Fields
	TalentLevel uint8
	TalentID    int32
}

type OWHuntActivityMonsterMetaData struct {
	// Fields
	MonsterShowID int32
	Hardlevel     uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	MonsterWeather      Hash
	MonsterDes          Hash
	MonsterDetailIDList []int32
}

type OWHuntActivityProgressMetaData struct {
	// Fields
	EventID int32

	// Properties
	Progress float32
	MapID    int32
	Floor    uint8
}

type OWHuntActivityQuestMetaData struct {
	// Fields
	QuestID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	FuncParam []StrWithPrefix16
}

type OWHuntActivityScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	ActivityID int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
}

type OWHuntActivitySHLevelMetaData struct {
	// Fields
	StrongholdLevel     uint8
	StrongholdLevelType uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	HardDes Hash
}

type OWHuntActivityStrongholdsMetaData struct {
	// Fields
	StrongholdID int32

	// Properties
	MapID               int32
	EventID             int32
	BornPointFloor      uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	StrongholdLevelType uint8
	Addr3               uint32 `json:"-"`

	// Objects
	StrongholdName Hash
	StrongholdDes  Hash
	FuncParam      []StrWithPrefix16
}

type OWHuntActivityTalentMetaData struct {
	// Fields
	TalentID    int32
	TalentLevel uint8

	// Properties
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	MapID                int32
	Addr3                uint32 `json:"-"`
	TalentUpgradePremise uint8
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	Addr6                uint32 `json:"-"`
	Addr7                uint32 `json:"-"`
	Addr8                uint32 `json:"-"`

	// Objects
	TalentIconPath    StrWithPrefix16
	PreTalent         []OWHuntActivityTalentMetaData_PreTalentData
	TalentUpgradeCost []OWHuntActivityTalentMetaData_TalentUpgradeCostData
	TalentName        Hash
	TalentDes         Hash
	NextLevelDes      Hash
	AbilityName       StrWithPrefix16
	AbilityParams     []StrWithPrefix16
}

type OWHuntActivityTalentMetaData_PreTalentData struct {
	// Fields
	TalentLevel uint8
	TalentID    int32
}

type OWHuntActivityTalentMetaData_TalentUpgradeCostData struct {
	// Fields
	TalentUpgradeCostMaterial int32
	TalentUpgradeCostNum      int32
}

type OWHuntActivityTeleportMetaData struct {
	// Fields
	TeleportID int32

	// Properties
	MapID int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	TextmapName    Hash
	TextmapDesc    Hash
	SpawnPointName StrWithPrefix16
	FuncParam      []StrWithPrefix16
}

type OWTalentDataMetaData struct {
	// Fields
	TalentID int32

	// Properties
	MaxLevel uint8
}

type OWTalentLevelDataMetaData struct {
	// Fields
	TalentID int32
	Level    uint8

	// Properties
	IsStoryValid bool
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`

	// Objects
	ConditionList StrWithPrefix16
	Name          StrWithPrefix16
	Desc          StrWithPrefix16
	IconPath      StrWithPrefix16
	AbilityName   StrWithPrefix16
	ParamList     []StrWithPrefix16
}

type OWTeleporterMetaData struct {
	// Fields
	EventID int32

	// Properties
	SectionID int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	TeleporterName Hash
	TeleporterDesc Hash
	SpawnPointList []StrWithPrefix16
	LockHintText   StrWithPrefix16
}

type PanelConversationTriggerMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	Popup_x_default int16
	Popup_y_default int16
}

type PayActivityDisplayScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`
	CGEventID int32

	// Objects
	BeginTime           StrWithPrefix16
	EndTime             StrWithPrefix16
	DisplayTextMap      StrWithPrefix16
	PopupDisplayTextMap StrWithPrefix16
	IconImgPath         StrWithPrefix16
	DisplayImgPathList  []StrWithPrefix16
}

type PayInfoMetaData struct {
	// Fields
	UserType uint8
	DevType  uint8

	// Properties
	CanPay bool
	Addr1  uint32 `json:"-"`

	// Objects
	CannotPayHint Hash
}

type PerformEventMetaData struct {
	// Fields
	EventID uint16

	// Properties
	EventType uint8
	Addr1     uint32 `json:"-"`

	// Objects
	ParamsVar []int32
}

type PhoneEntranceAcountOverrideMetaData struct {
	// Fields
	EntranceId  int32
	AccountType int32

	// Properties
	Addr1  uint32 `json:"-"`
	IsOpen bool
	Addr2  uint32 `json:"-"`

	// Objects
	EntranceIconPath StrWithPrefix16
	LinkUrl          StrWithPrefix16
}

type PhoneEntranceMetaData struct {
	// Fields
	EntranceId int32

	// Properties
	BtnType            int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Weight             int32
	UrlType            int32
	Addr3              uint32 `json:"-"`
	LinkType           int32
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	RedHintType        int32
	RedHintUiType      int32
	RedHintMainPage    int32
	MinLv              int32
	MaxLv              int32
	ChannelSettingType int32
	Addr7              uint32 `json:"-"`

	// Objects
	EntranceBeginTime  StrWithPrefix16
	EntranceEndTime    StrWithPrefix16
	LinkUrl            StrWithPrefix16
	LinkParams         StrWithPrefix16
	EntranceText       Hash
	EntranceIconPath   StrWithPrefix16
	UIOverrideItemList []uint32
}

type PhoneNoticeDataMetaData struct {
	// Fields
	Noticeid int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	NoticeAppearCD   int32
	NoticeAppearTime int32
	Addr3            uint32 `json:"-"`

	// Objects
	NoticeBeginTime RemoteTime
	NoticeEndTime   RemoteTime
	NoticeVideoID   []int32
}

type PicBGListdMeta struct {
	// Fields
	ID int32

	// Properties
	PicBGType int32
}

type PictureStepMetaData struct {
	// Fields
	StepID   int32
	ChoiceID int32

	// Properties
	KeyID     int32
	Addr1     uint32 `json:"-"`
	Reward    int32
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`
	NeedScore int32

	// Objects
	ButtonPrefab         StrWithPrefix16
	Name_Textmap         Hash
	ChoiceBefore_Textmap []Hash
	ChoiceAfter_Textmap  []Hash
	Pic                  StrWithPrefix16
}

type PicTutorialMetaData struct {
	// Fields
	ActivityID int32
	ID         int32

	// Properties
	Type            uint32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	UnlockCondition int32
	Weight          uint8
	Param           int32

	// Objects
	Title      StrWithPrefix16
	StepIDList []int32
}

type PicTutorialStepDataMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Title   StrWithPrefix16
	Desc    StrWithPrefix16
	ImgPath StrWithPrefix16
}

type PlayerLevelLockMetaData struct {
	// Fields
	LockID int32

	// Properties
	MaxLevel        uint8
	MaxLayUpExp     int32
	Addr1           uint32 `json:"-"`
	EventID         int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	TutorialVideoID int32

	// Objects
	MissionList        []int32
	PlayerLockHintText Hash
	LinkButton         StrWithPrefix16
	LevelGiftPackBG    StrWithPrefix16
	LevelGiftPackFrame StrWithPrefix16
}

type PlayerLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp               int32
	Stamina           int32
	NumFriends        int32
	AvatarLevelLimit  int32
	StaminaBonus      int32
	SweepStaminaLimit int32
}

type PlayerLevelShopGoodsMetaData struct {
	// Fields
	Level int32

	// Properties
	Addr1  uint32 `json:"-"`
	LockID int32

	// Objects
	GoodsIDList []int32
}

type PlayerPrivilegeScheduleMetaData struct {
	// Fields
	PrivilegeID int32

	// Properties
	PrivilegeType int32
	ScheduleID    int32
	Addr1         uint32 `json:"-"`

	// Objects
	RightConfigList []int32
}

type PlotAvatarMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Addr1       uint32 `json:"-"`
	DisplayType int32
	Addr2       uint32 `json:"-"`

	// Objects
	Name StrWithPrefix16
	Path StrWithPrefix16
}

type PlotLineCgDataMetaData struct {
	// Fields
	PlotId int32

	// Properties
	CGID int32
}

type PlotlineConditionDataMetaData struct {
	// Fields
	ID uint32

	// Properties
	ConditionType uint32
	Addr1         uint32 `json:"-"`

	// Objects
	Param StrWithPrefix16
}

type PlotLineDataMetaData struct {
	// Fields
	PlotId uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	PlotlineJsonPath StrWithPrefix16
}

type PlotMetaData struct {
	// Fields
	PlotID int32

	// Properties
	LevelID       int32
	StartDialogID int32
	EndDialogID   int32
}

type PowerTypeMetaData struct {
	// Fields
	Type int32

	// Properties
	PowerConf float32
}

type PrayGachaMetaData struct {
	// Fields
	GachaID int32

	// Properties
	AvatarID int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`
	Addr6    uint32 `json:"-"`
	Addr7    uint32 `json:"-"`

	// Objects
	LuckAddDesc     Hash
	CurrentLuckDesc Hash
	RecommendDesc   Hash
	SmallTipsDesc   Hash
	HelpDesc        Hash
	OneGachaImg     StrWithPrefix16
	TenGachaImg     StrWithPrefix16
}

type PredownloadAsbMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	Version    StrWithPrefix16
	EnableTime RemoteTime
}

type PredownloadAudioMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	Version    StrWithPrefix16
	EnableTime RemoteTime
}

type PredownloadAudioPackageMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2   uint32 `json:"-"`
	PckType uint8

	// Objects
	PackageName StrWithPrefix16
	Version     StrWithPrefix16
}

type PredownloadVideoFileMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2    uint32 `json:"-"`
	PckType  uint8
	FileSize int32

	// Objects
	VideoName StrWithPrefix16
	Version   StrWithPrefix16
}

type PredownloadVideoMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	Version    StrWithPrefix16
	EnableTime RemoteTime
}

type ProductRecommendMetaData struct {
	// Fields
	RecommandID int32

	// Properties
	ShowWeight     int32
	CancelType     int32
	DialogType     uint8
	TrickEventType int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	ImgBGType      uint8
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`

	// Objects
	ShowTagList             []int32
	ImgBG                   StrWithPrefix16
	ShopGoodsScheduleIDList []uint32
	TitleText               Hash
	SubTitleText            Hash
	RecommendTipsText       Hash
	BeginTime               RemoteTime
	EndTime                 RemoteTime
	TipsLink                StrWithPrefix16
}

type ProtectTypeMetaData struct {
	// Fields
	Protectypeid int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ProtectDisplayDesc Hash
	ProtectInfoTitle   Hash
	ProtectInfoDesc    Hash
}

type PVloginMetaData struct {
	// Fields
	ActivityId uint16

	// Properties
	IsHideCountDownInfo bool
	IsHideGotoLiveBtn   bool
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	Addr9               uint32 `json:"-"`

	// Objects
	JumpLiveSpecifiedTimeBefore             RemoteTime
	JumpLiveSpecifiedTimeAfter              RemoteTime
	JumpInGameLink1                         StrWithPrefix16
	JumpInGameLink2                         StrWithPrefix16
	JumpInGameLink3                         StrWithPrefix16
	CL_ActivityPvLoginPanel_LastingTime     Hash
	CL_ActivityPvLoginPanel_Label_Countdown Hash
	CL_ActivityPvLoginPanel_Dialog_NotStart Hash
	CL_ActivityPvLoginPanel_Dialog_HasEnded Hash
}

type PVZActivityMetaData struct {
	// Fields
	ActivityID uint32

	// Properties
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	MedalID int32

	// Objects
	QavatarIDList  []int32
	JsonConfigPath StrWithPrefix16
}

type PVZMonsterMetaData struct {
	// Fields
	MonsterID int32

	// Properties
	LinkedActivityID int32
	Model            int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	OffsetY          float32

	// Objects
	QmonsterName         Hash
	QmonsterSkill        Hash
	Solution             Hash
	RecommendQavatarList []int32
	QmonsterIcon         StrWithPrefix16
}

type PVZQavatarMetaData struct {
	// Fields
	QavatarID uint32
	Level     uint32

	// Properties
	QavatarLevelUpMaterialID int32
	NextlevelCost            int32
	SpecialMaterialID        int32
	EntityID                 int32
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	QavatarCD                float32
	Addr3                    uint32 `json:"-"`
	CGID                     int32
	QavatarCost              int32
	NextLevelConditionType   uint8
	NextLevelConditionParam  int32
	QavatarATKShow           int32
	QavatarHPShow            int32
	Addr4                    uint32 `json:"-"`
	Addr5                    uint32 `json:"-"`
	QavatarType              uint8

	// Objects
	QavatarName             Hash
	QavatarDesc             Hash
	ChibiIcons              StrWithPrefix16
	TargetArea              []PVZQavatarMetaData_CellAreaConfig
	TargetAreaWithOrnaments []PVZQavatarMetaData_CellAreaConfig
}

type PVZQavatarMetaData_CellAreaConfig struct {
	// Fields
	Type  int32
	Param float32
}

type PVZSiteMetaData struct {
	// Fields
	RpgSiteID int32

	// Properties
	SiteType   int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	NextSiteID int32
	MaxFloor   int32
	Addr3      uint32 `json:"-"`
	MaxWaves   int32

	// Objects
	TicketCost        []PVZSiteMetaData_MatPair
	DropMaterialLimit []PVZSiteMetaData_MatPair
	SiteTitle         Hash
}

type PVZSiteMetaData_MatPair struct {
	// Fields
	ID  int32
	Num int32
}

type PVZTileMetaData struct {
	// Fields
	TowerID int32
	FloorID int32

	// Properties
	LimitQAvatarNumber    int32
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	RecommendQavatarLevel int32
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`

	// Objects
	InitialQavatarList   []uint32
	MonsterShowList      []int32
	RecommendQavatarList []uint32
	StageDesc            Hash
	DropMaterial         []PVZTileMetaData_MatPair
}

type PVZTileMetaData_MatPair struct {
	// Fields
	ID  int32
	Num int32
}

type QAvatarBattleBroadCastMetaData struct {
	// Fields
	BroadCastID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	BroadCastDescription Hash
}

type QAvatarBattleExpressionMetaData struct {
	// Fields
	EmojiFaceID int32
}

type QAvatarBattleGadgetMetaData struct {
	// Fields
	GadgetID int32

	// Properties
	ContactId       int32
	ParaInt         int32
	GadgetHP        float32
	IsShowHPBar     bool
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	GadgetIconPos   float32
	GadgetIconScale float32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`

	// Objects
	GadgetPrefab                 StrWithPrefix16
	GadgetIcon                   StrWithPrefix16
	GadgetBornEffectPattern      StrWithPrefix16
	GadgetDisappearEffectPattern StrWithPrefix16
}

type QAvatarBattleGadgetRefreshMetaData struct {
	// Fields
	ID int32

	// Properties
	ForcastID   int32
	BroadCastID int32
}

type QAvatarBattleKillingStreakMetaData struct {
	// Fields
	Num  uint16
	Type int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	Name          Hash
	AnimName      StrWithPrefix16
	AudioPattern1 StrWithPrefix16
	AudioPattern2 StrWithPrefix16
}

type QAvatarBattleSpawnZoneMetaData struct {
	// Fields
	SpawnPointID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	SpawnPointTextmap Hash
}

type QAvatarBattleTextIdMapMetaData struct {
	// Fields
	ID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TextString StrWithPrefix16
}

type QAvatarMissionMetaData struct {
	// Fields
	MissionID int32
	AvatarID  int32

	// Properties
	ProcessReward int32
	Addr1         uint32 `json:"-"`

	// Objects
	MissionTypeText Hash
}

type QAvatarPVPMetaData struct {
	// Fields
	QAvatarPVPID int32

	// Properties
	PairAvatarID int32
	WeaponID     int32
	Addr1        uint32 `json:"-"`

	// Objects
	QAvatarName Hash
}

type QAvatarPVPWeaponMetaData struct {
	// Fields
	WeaponID int32

	// Properties
	Type              int32
	Addr1             uint32 `json:"-"`
	AimDistance       float32
	SkillID           int32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	BulletNum         int32
	BulletRecoverRate float32
	ReloadTime        float32
	BulletWarningNum  int32

	// Objects
	WeaponTypeIcon                StrWithPrefix16
	WeaponAssetPath               StrWithPrefix16
	WeaponAnimationPartsAssetPath []StrWithPrefix16
	WeaponAttachRoot              StrWithPrefix16
	WeaponAttachPosition          []float32
	WeaponAttachRotation          []float32
}

type QCandyActivityDataMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	UnlockItemID         int32
	CoinMaterialID       int32
	DailyLimit           int32
	ActivityMedalID      int32
	InviteStageID        int32
	InLevelTrialAvatarID int32
	MedalUnlockRank      int32
	MedalHighestRank     int32
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	PVPRequireSite       int32
	DefaultAvatar        int32

	// Objects
	MapPoolList  []int32
	PVPCloseTime RemoteTime
}

type QCandyAvatarMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	InitFace      bool
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	ItemCost      int32
	RankRequest   int32
	IsShownLocked bool

	// Objects
	AvatarRegistryKey  StrWithPrefix16
	UIAvatarPrefabPath StrWithPrefix16
	UIControllerPath   StrWithPrefix16
	FullName           Hash
	ChibiIconPath      StrWithPrefix16
}

type QCandyPvpMapMetaData struct {
	// Fields
	ID int32

	// Properties
	Type         uint8
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	StageID      int32
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Lives        uint8
	CountDown    int32
	Score        int32
	Addr6        uint32 `json:"-"`
	TimeLineTime float32
	Addr7        uint32 `json:"-"`

	// Objects
	Name         Hash
	Desc         Hash
	Image        StrWithPrefix16
	RecordImage  StrWithPrefix16
	SkillList    []int32
	JsonName     StrWithPrefix16
	TimeLineName StrWithPrefix16
}

type QCandyPvpMapPoolMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MapPool []int32
}

type QCandyPvpMapSkillMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	CD          int32
	MaxUseTimes int32
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`

	// Objects
	Icon        StrWithPrefix16
	Name        Hash
	Desc        Hash
	AbilityName StrWithPrefix16
	ParamList   []float32
}

type QCandyRankSectionMetaData struct {
	// Fields
	SectionID uint8

	// Properties
	RankScoreBegin int32
	RankScoreEnd   int32
}

type QCandySettleConfigMetaData struct {
	// Fields
	Rank uint8

	// Properties
	Addr1          uint32 `json:"-"`
	AddCurrencyNum int32

	// Objects
	RankScoreDelta []QCandySettleConfigMetaData_Setting
}

type QCandySettleConfigMetaData_Setting struct {
	// Fields
	SectionID  uint8
	ScoreDelta int32
}

type QTEndlessMonsterData struct {
	// Fields
	MonsterID uint16

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	IsElite          bool
	UniqueID         int32
	Score            int32
	Nature           uint8
	Addr3            uint32 `json:"-"`
	MonsterRank      uint8
	Weight           uint8
	Addr4            uint32 `json:"-"`
	MonsterDisplayID int32
	Addr5            uint32 `json:"-"`

	// Objects
	MonsterName     StrWithPrefix16
	ConfigType      StrWithPrefix16
	MonsterTagList  []StrWithPrefix16
	GroupLevelRange []int32
	Tags            StrWithPrefix16
}

type QTEndlessMonsterWaveMetaData struct {
	// Fields
	WaveID uint16

	// Properties
	Addr1            uint32 `json:"-"`
	MonsterGroupType uint8
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	HPRatio          float32
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`

	// Objects
	RankList    []StrWithPrefix16
	TagDistRule StrWithPrefix16
	TagFilter   StrWithPrefix16
	DistRule    []StrWithPrefix16
	RankDist    []int32
}

type QTEndlessScheduleMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	RewardConfig uint32

	// Objects
	StartTime RemoteTime
}

type QuestionBankMetaData struct {
	// Fields
	QuestionBankID uint32

	// Properties
	Type        int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	RightReward int32
	WrongReward int32

	// Objects
	QuestionText   Hash
	AnswerTextList []QuestionBankMetaData_Answer
}

type QuestionBankMetaData_Answer struct {
	// Fields
	AnswerID uint32
	Addr1    uint32 `json:"-"`

	// Objects
	AnswerTextID Hash
}

type QuestionConfigMetaData struct {
	// Fields
	QuestionID uint32

	// Properties
	AskAvatarID int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`

	// Objects
	StartText           Hash
	FinishText          Hash
	RightAnswerTextList []Hash
	WrongAnserTextList  []Hash
}

type QuestionScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	QuestionID uint32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	OpenTime  StrWithPrefix16
	CloseTime StrWithPrefix16
}

type QuickEntryScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	EntryType int32
	Priority  int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`
	MinLevel  int32
	MaxLevel  int32

	// Objects
	PrefabPath   StrWithPrefix16
	TitleTextMap StrWithPrefix16
	ImagePath    StrWithPrefix16
	Link         StrWithPrefix16
	StartTime    StrWithPrefix16
	EndTime      StrWithPrefix16
}

type QuickSellDataMetaData struct {
	// Fields
	SellUcId int32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	IsDisplayVitality int32
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`

	// Objects
	SellDialogCrtNumDesc        Hash
	SellDialogSellNumDesc       Hash
	SellDialogUnitPriceDesc     Hash
	SellDialogObtainNumDesc     Hash
	SellDialogSellButtonDesc    Hash
	SellDialogInfoTitle         Hash
	SellDialogInfoContent       Hash
	AdvSellDialogSellButtonDesc Hash
}

type RaffleClientMetaData struct {
	// Fields
	RaffleId uint32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	RewardLevel uint8
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	FirstShow   uint32
	SecondShow  uint32
	ThirdShow   uint32
	HitFaceCD   int32
	HaveLink    bool
	Addr5       uint32 `json:"-"`
	Addr6       uint32 `json:"-"`

	// Objects
	RafflePic       StrWithPrefix16
	FirstRewardType StrWithPrefix16
	FaceBeginTime   StrWithPrefix16
	FaceEndTime     StrWithPrefix16
	LinkButton      StrWithPrefix16
	RaffleColor     StrWithPrefix16
}

type RaffleConfigMetaData struct {
	// Fields
	RaffleId uint32

	// Properties
	Addr1             uint32 `json:"-"`
	MaxFirstRewardNum uint32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	MaxDrawTimes      uint32
	RaffleGroup       int32
	Addr4             uint32 `json:"-"`

	// Objects
	FirstRewardNumberList  []uint32
	SecondRewardNumberList []uint32
	CostMaterialList       []RaffleConfigMetaData_CostMaterial
	TextmapRule            Hash
}

type RaffleConfigMetaData_CostMaterial struct {
	// Fields
	Count      int32
	MaterialID int32
}

type RaffleRewardConfigMetaData struct {
	// Fields
	RaffleId uint32

	// Properties
	FirstRewardId  uint32
	SecondRewardId uint32
	ThirdRewardId  uint32
}

type RaffleScheduleConfigMetaData struct {
	// Fields
	ScheduleId uint32

	// Properties
	Type                   uint8
	ExchangeCostMaterialId int32
	Addr1                  uint32 `json:"-"`

	// Objects
	ExchangeGetMaterialList []int32
}

type RaidActivateScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	DailyMPActivityList []int32
	RaidActivityList    []int32
}

type RanchAreaConfigMetaData struct {
	// Fields
	RanchAreaID int32

	// Properties
	AddSlot      int32
	ProductLimit int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	ProductDisplayList []int32
	UnlockTextmap      Hash
}

type RanchAreaMetaData struct {
	// Fields
	AreaID int32

	// Properties
	Addr1           uint32 `json:"-"`
	ExitRanchSiteID int32
	Addr2           uint32 `json:"-"`
	IsChallenge     bool
	Addr3           uint32 `json:"-"`
	Boundary        float32
	BackGroundType  uint8

	// Objects
	ContentList []int32
	Collection  []RanchAreaMetaData_DropDisplay
	Location    []float32
}

type RanchAreaMetaData_DropDisplay struct {
	// Fields
	IsOvert   uint8
	MonsterID int32
}

type RanchChallengeSiteMetaData struct {
	// Fields
	RanchSiteID int32

	// Properties
	Addr1          uint32 `json:"-"`
	ModelMonsterID int32
	ShowRank       bool

	// Objects
	MissionList []int32
}

type RanchContentMetaData struct {
	// Fields
	ContentID int32

	// Properties
	Type  uint8
	Para  int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Icon          StrWithPrefix16
	NoticeTextMap Hash
}

type RanchDataMetaData struct {
	// Fields
	RanchID int32

	// Properties
	ReturnRatio           float32
	ReturnMaterialID      int32
	MonsterSyntheticRatio float32
	LockSkillNumOverallID int32
	WarehouseLimit        int32
}

type RanchMonsterExpMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp        int32
	Addr1      uint32 `json:"-"`
	UnlockSite int32

	// Objects
	UnlockTime RemoteTime
}

type RanchMonsterMetaData struct {
	// Fields
	MonsterID int32

	// Properties
	ActiveSkill       int32
	Addr1             uint32 `json:"-"`
	MaxLevel          uint8
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	CanUpgrade        bool
	CanBeConsumed     bool
	Addr6             uint32 `json:"-"`
	BasicExp          int32
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	GrazingEfficiency int32
	IsAvatar          bool
	UniqueID          int32
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	ObtainType        uint8
	DropParam         int32
	CollectOrder      int32

	// Objects
	MonsterName         Hash
	BattlePicPath       StrWithPrefix16
	BattleStagePicPath  StrWithPrefix16
	HeadPicPath         StrWithPrefix16
	RanchHeadPicPath    StrWithPrefix16
	EasterEggSite       []int32
	MonsterDesc         Hash
	UnlockTime          RemoteTime
	ActiveSkillCD       []RanchMonsterMetaData_CDRange
	SkillTypeMapDisplay []int32
}

type RanchMonsterMetaData_CDRange struct {
	// Fields
	MinLevel uint8
	CoolDown int32
}

type RanchMonsterSkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	BasicType uint8
	SkillType uint16
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Rarity    int32

	// Objects
	SkillName StrWithPrefix16
	Name      Hash
	Icon      StrWithPrefix16
}

type RanchSiteDataMetaData struct {
	// Fields
	RanchSiteID int32

	// Properties
	UnlockCost       int32
	PassedCost       int32
	Addr1            uint32 `json:"-"`
	BuffDrop         int32
	IsSiteDialogShow uint8
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`

	// Objects
	RanchMonsterDropList []int32
	UnlockFlagList       StrWithPrefix16
	Location             []float32
}

type RanchSyntheticMetaData struct {
	// Fields
	SyntheticID int32

	// Properties
	FormulaMaterialID  int32
	MainRanchMonster   int32
	AssistRanchMonster int32
	SyntheticProduct   int32
}

type RandomDialogActionMetaData struct {
	// Fields
	PlotId   int32
	DialogId int32

	// Properties
	ActionType int32
}

type RandomDialogCGRawMetaData struct {
	// Fields
	CGRawID int32

	// Properties
	CGPosType uint8
	Addr1     uint32 `json:"-"`

	// Objects
	Pos RandomDialogCGRawMetaData_CGRawPos
}

type RandomDialogCGRawMetaData_CGRawPos struct {
	// Fields
	X float32
	Y float32
}

type RandomDialogInputMetaData struct {
	// Fields
	InputID int32

	// Properties
	InputType   int32
	InputLength int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	TipText          StrWithPrefix16
	InputNullTipText StrWithPrefix16
	PlaceholderText  StrWithPrefix16
}

type RandomDialogQuestionMetaData struct {
	// Fields
	Question_dialogId int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	Talker_name      StrWithPrefix16
	Question_content StrWithPrefix16
}

type RandomEffectLevelMetaData struct {
	// Fields
	EffectID    int32
	EffectLevel int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	EffectIcon StrWithPrefix16
	EffectText StrWithPrefix16
	EffectName StrWithPrefix16
}

type RandomEffectMetaData struct {
	// Fields
	EffectID int32

	// Properties
	EffectType     int32
	EffectQuality  int32
	EffectMaxLevel int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`

	// Objects
	EffectIcon      StrWithPrefix16
	EffectText      Hash
	EffectName      Hash
	StageEffectIcon StrWithPrefix16
	EffectColor     StrWithPrefix16
}

type RandomHideBranchTimeMetaData struct {
	// Fields
	HideBranchID int32

	// Properties
	HideTime int32
}

type RandomPlotDataMetaData struct {
	// Fields
	PlotId int32

	// Properties
	StartDialogId  int32
	Addr1          uint32 `json:"-"`
	BgmDown        float32
	BlackFade      bool
	UserCanControl bool
	UseSoundBank   bool
	CanSkipOption  bool
	Addr2          uint32 `json:"-"`

	// Objects
	FinishDialogIdList []int32
	ExternalEvtName    StrWithPrefix16
}

type RandomPlotNPCMetaData struct {
	// Fields
	NPCID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Path StrWithPrefix16
}

type RandomSubSelectContentDataMetaData struct {
	// Fields
	DialogID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	SatusBasedContents []RandomSubSelectContentDataMetaData_SubSelectContentData
}

type RandomSubSelectContentDataMetaData_SubSelectContentData struct {
	// Fields
	Status uint32
	Addr1  uint32 `json:"-"`

	// Objects
	Content StrWithPrefix16
}

type RandomSubSelectStyleConfigMetaData struct {
	// Fields
	DialogID uint32

	// Properties
	StyleType uint32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	UiStyle   bool

	// Objects
	IntParams []uint32
	Icon      StrWithPrefix16
}

type RareaffixMetaData struct {
	// Fields
	CombinationID int32

	// Properties
	AffixAttributetype1 int32
	Attributepercent1   int32
	AffixAttributetype2 int32
	Attributepercent2   int32
	Addr1               uint32 `json:"-"`

	// Objects
	Reconfirm_textmap Hash
}

type ReclaimConfigMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	OpenCgId int32
	Addr5    uint32 `json:"-"`

	// Objects
	CloseTime           StrWithPrefix16
	EnterIconPrefabPath StrWithPrefix16
	EnterIconImgPath    StrWithPrefix16
	BgPrefabPath        StrWithPrefix16
	MissionList         []int32
}

type ReclaimEventListMetaData struct {
	// Fields
	Id int32

	// Properties
	ConfigId      int32
	IconPos       int32
	ShowFinishWay int32
	Addr1         uint32 `json:"-"`
	UnLockLevel   int32
	ActivityId    int32

	// Objects
	SideMark StrWithPrefix16
}

type ReclaimEventStageListMetaData struct {
	// Fields
	StageId int32

	// Properties
	UnlockLevel int32
}

type ReclaimEventVirtualAvatarMetaData struct {
	// Fields
	Id int32

	// Properties
	AvatarId   int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type ReclaimEventVirtualStigmataMetaData struct {
	// Fields
	Id int32

	// Properties
	StigmataId int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type ReclaimEventVirtualWeaponMetaData struct {
	// Fields
	Id int32

	// Properties
	WeaponId   int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type ReclaimLevelMetaData struct {
	// Fields
	Id int32

	// Properties
	ConfigId     int32
	Level        int32
	NeedEventExp int32
	RewardId     int32
}

type ReclaimRankingRewardData struct {
	// Fields
	Id int32

	// Properties
	Ranking  int32
	RewardId int32
}

type ReclaimScoreRewardData struct {
	// Fields
	Id int32

	// Properties
	RankingScore int32
	RewardId     int32
}

type ReclaimVirtualGachaChestMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	ChestIconID  int32
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	ChestRarity  int32
	ShowRewardId int32

	// Objects
	GachaPoolList      []ReclaimVirtualGachaChestMetaData_Materials
	SelectChestImgPath StrWithPrefix16
	ChestName          Hash
	ChestDesc          Hash
}

type ReclaimVirtualGachaChestMetaData_Materials struct {
	// Fields
	MaterialID int32
	Unused     int32
}

type RecommendEquipMetaData struct {
	// Fields
	ID      int32
	EquipID int32

	// Properties
	GroupAdvsID   int32
	Addr1         uint32 `json:"-"`
	RecommendStar int32

	// Objects
	RecommendText Hash
}

type RedEnvelopeMetaData struct {
	// Fields
	ID int32

	// Properties
	Num   int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	Channellist         []int32
	DefaultText         StrWithPrefix16
	ChatTakeIcon        StrWithPrefix16
	OpenPageTitleText   Hash
	ClickIcon           StrWithPrefix16
	BackgroundImg       StrWithPrefix16
	RecordBackgroundImg StrWithPrefix16
}

type RegionInfoMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	Region StrWithPrefix16
	Name   StrWithPrefix16
}

type RelationActivityScheduleMetaData struct {
	// Fields
	Id int32

	// Properties
	ShopType   int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`
	MaterialID int32
	LimitNum   int32
	Addr6      uint32 `json:"-"`

	// Objects
	InfoTitle        Hash
	InfoContent      Hash
	ActivityTimeDesc Hash
	BeginTime        StrWithPrefix16
	EndTime          StrWithPrefix16
	WikiLink         StrWithPrefix16
}

type RelationStageBonusDropMetaData struct {
	// Fields
	StageId int32
}

type ReplayLobbyFilterMetaData struct {
	// Fields
	ID int32

	// Properties
	LobbyType int32
	KeyID     int32
	Priority  int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
	Name     StrWithPrefix16
}

type ReplayLobbyPageMetaData struct {
	// Fields
	ID int32

	// Properties
	Type     int32
	Addr1    uint32 `json:"-"`
	Priority int32
	Addr2    uint32 `json:"-"`
	IsShow   uint8

	// Objects
	TabName  StrWithPrefix16
	IconPath StrWithPrefix16
}

type ReplayLobbyScoreMetaData struct {
	// Fields
	ID int32

	// Properties
	MaxDisplayNum int32
}

type ReplayLobbySubPageMetaData struct {
	// Fields
	ID       int32
	FatherID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TabName StrWithPrefix16
}

type ReplayLobbyUploadMetaData struct {
	// Fields
	UserType int32

	// Properties
	MinLevel        int32
	MaxLevel        int32
	ExBossUploadNum int32
}

type ReportReasonMetaData struct {
	// Fields
	ReportReason int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	ReasonName   Hash
	ReasonDetail Hash
}

type ResourceRetrieveBoxMetaData struct {
	// Fields
	BoxID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	RewardID    int32
	CostStamina int32
	CostHcoin   int32
	Addr3       uint32 `json:"-"`

	// Objects
	BoxIcon          StrWithPrefix16
	BoxDes           Hash
	CostMaterialList []ResourceRetrieveBoxMetaData_CostMaterial
}

type ResourceRetrieveBoxMetaData_CostMaterial struct {
	// Fields
	CostId  int32
	CostNum int32
}

type ResourceRetrieveModuleMetaData struct {
	// Fields
	ModuleID int32

	// Properties
	NormalBoxID int32
	SuperBoxID  int32
	Addr1       uint32 `json:"-"`

	// Objects
	ModuleDes Hash
}

type ResourceRetrieveScheduleMetaData struct {
	// Fields
	ScheduleID int32
}

type RestaurantActionZoneMetaData struct {
	// Fields
	RoomType uint8

	// Properties
	MoveSpeed   float32
	IntervalMin float32
	IntervalMax float32
	BoundaryMin float32
	CollisionX  float32
	CollisionY  float32
}

type RestaurantAvatarMetaData struct {
	// Fields
	VirtualAvatarID int32

	// Properties
	Rank          uint8
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	CookBehaviour uint8
	Addr8         uint32 `json:"-"`
	Addr9         uint32 `json:"-"`

	// Objects
	RestaurantSkillList []int32
	AvatarName          Hash
	AvatarDesc          Hash
	AvatarChibiIcon     StrWithPrefix16
	AvatarImgPath       StrWithPrefix16
	PrefabPath          StrWithPrefix16
	RandomBubbleID      []Hash
	FirstBubbleDuration []float32
	AfterBubbleDuration []float32
}

type RestaurantIngredientsMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	StockLimit []RestaurantIngredientsMetaData_LvLimit
	ImagePath  StrWithPrefix16
}

type RestaurantIngredientsMetaData_LvLimit struct {
	// Fields
	Level uint8
	Limit int32
}

type RestaurantLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	UpgradeCost    int32
	Addr1          uint32 `json:"-"`
	MaxAvatar      int32
	MaxOrder       int32
	UpgradedPlotID int32
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`

	// Objects
	RequireMissionIDList []int32
	UnlockDesc           Hash
	UnlockIngredientList []int32
	UnlockRecipeList     []int32
	PreviewImgPath       StrWithPrefix16
}

type RestaurantQuestMetaData struct {
	// Fields
	QuestID int32

	// Properties
	MissionID int32
	SubType   uint8
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	RequireMaterialList []RestaurantQuestMetaData_MatCost
	AvatarIconPath      StrWithPrefix16
}

type RestaurantQuestMetaData_MatCost struct {
	// Fields
	MaterialID int32
	Number     int32
}

type RestaurantRecipeMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Tag       uint8
	Addr1     uint32 `json:"-"`
	SellPrice int32
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Bundle    int32

	// Objects
	Ingredient []RestaurantRecipeMetaData_Source
	StockLimit []RestaurantRecipeMetaData_LvLimit
	CookSpeed  []RestaurantRecipeMetaData_LvSpeed
	SellSpeed  []RestaurantRecipeMetaData_LvSpeed
}

type RestaurantRecipeMetaData_LvLimit struct {
	// Fields
	Level uint8
	Limit int32
}

type RestaurantRecipeMetaData_Source struct {
	// Fields
	MaterialID int32
	Number     int32
}

type RestaurantRecipeMetaData_LvSpeed struct {
	// Fields
	Level uint8
	Speed float32
}

type RestaurantRoomMetaData struct {
	// Fields
	RoomID int32

	// Properties
	Type  uint8
	Addr1 uint32 `json:"-"`

	// Objects
	NodePath StrWithPrefix16
}

type RestaurantSkillMetaData struct {
	// Fields
	SkillID int32

	// Properties
	SkillType  uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	SkillLevel uint8

	// Objects
	SkillName      Hash
	SkillDesc      Hash
	SkillIconPath  StrWithPrefix16
	SkillParameter []int32
}

type RestaurantWeatherMetaData struct {
	// Fields
	WeatherID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	SkillID          []int32
	WeatherDesc      Hash
	WeatherEffectDes Hash
	BeginTime        RemoteTime
	EndTime          RemoteTime
	ShowRoleImage    StrWithPrefix16
	ShowMascot       StrWithPrefix16
}

type RestrictExtend_StageMetaData struct {
	// Fields
	StageID int32

	// Properties
	RestrictID uint32
}

type ReunionCookBookMetaData struct {
	// Fields
	CookGroupID int32
	CookID      int32

	// Properties
	CookType       int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	CookPoint      int32
	CookProduce    int32
	CookDailyTimes int32
	CookTotalTimes int32

	// Objects
	CookName         StrWithPrefix16
	CookDesc         StrWithPrefix16
	CookConsumeList1 []ReunionCookBookMetaData_CookConsumerData
}

type ReunionCookBookMetaData_CookConsumerData struct {
	// Fields
	Count int32
	ID    int32
}

type ReunionCookRewardMetaData struct {
	// Fields
	ScoreRewardGroup int32
	ID               uint8

	// Properties
	NeedScore int32
	RewardID  int32
}

type ReunionCookStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Rate          int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	RewardDisplay int32
	Addr5         uint32 `json:"-"`

	// Objects
	BGpicPath    StrWithPrefix16
	PlotLuaPath  StrWithPrefix16
	Position     []int32
	TablePicPath StrWithPrefix16
	RewardDesc   StrWithPrefix16
}

type ReunionDisplayScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	CollectBeginTime RemoteTime
	CollectEndTime   RemoteTime
	RewardBeginTime  RemoteTime
	RewardEndTime    RemoteTime
	OpenDayTime      RemoteTimeSpan
	EndDayTime       RemoteTimeSpan
}

type ReviveCostTypeMetaData struct {
	// Fields
	ReviveTimes int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Type1 []int32
}

type ReviveEffectMetaData struct {
	// Fields
	ReviveEffectID int32

	// Properties
	ReviveHp int32
	ReviveSp int32
	Addr1    uint32 `json:"-"`

	// Objects
	SpecialAbilityName StrWithPrefix16
}

type ReviveUseMetaData struct {
	// Fields
	ReviveUseID uint16

	// Properties
	MaterialID   int32
	IsExclude    bool
	Addr1        uint32 `json:"-"`
	ReviveEffect int32
	Addr2        uint32 `json:"-"`

	// Objects
	CostCoinList     []int32
	ReviveEffectDesc Hash
}

type RewardData struct {
	// Fields
	RewardID int32

	// Properties
	RewardExp         uint16
	RewardHCoin       uint16
	RewardStamina     uint16
	RewardFriendPoint uint16
	RewardDutyPoint   uint8
	RewardItem1ID     int32
	RewardItem1Level  uint8
	RewardItem1Num    int32
	RewardItem2ID     int32
	RewardItem2Level  uint8
	RewardItem2Num    int32
	RewardItem3ID     int32
	RewardItem3Level  uint8
	RewardItem3Num    int32
	RewardItem4ID     int32
	RewardItem4Level  uint8
	RewardItem4Num    int32
	RewardItem5ID     int32
	RewardItem5Level  uint8
	RewardItem5Num    int32
	RewardItem6ID     int32
	RewardItem6Level  uint8
	RewardItem6Num    int32
}

type RewardLineMetaData struct {
	// Fields
	RewardLineID uint16

	// Properties
	Addr1             uint32 `json:"-"`
	ScoreLinkMaterial int32

	// Objects
	RankIDList []uint16
}

type RewardLineRankMetaData struct {
	// Fields
	RankID uint16

	// Properties
	Rank          uint8
	RewardID      int32
	RequiredScore uint32
	IsKeyReward   bool
}

type RewardOverviewPanelMetaData struct {
	// Fields
	ID int32

	// Properties
	ItemID   int32
	Addr1    uint32 `json:"-"`
	Mission  int32
	LinkType int32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	ImagePath   StrWithPrefix16
	LinkParams  []int32
	DisplayTime StrWithPrefix16
}

type RichAreaMetaData struct {
	// Fields
	RichAreaID int32

	// Properties
	RichID         int32
	TileTowerID    int32
	ThemeID        int32
	AreaCurrencyID int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	QAvartarID     int32
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	SubToArea      int32
	IsMainArea     int32

	// Objects
	AreaName      StrWithPrefix16
	BuffList      []int32
	ItemList      []int32
	AreaSubTitle  StrWithPrefix16
	AreaNumber    StrWithPrefix16
	AreaTabBGpath StrWithPrefix16
}

type RichBuffMetaData struct {
	// Fields
	BuffID int32
	Level  int32

	// Properties
	Value1 float32
	Value2 float32
	Value3 float32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`

	// Objects
	Name   StrWithPrefix16
	Desc   StrWithPrefix16
	Icon   StrWithPrefix16
	SpDesc StrWithPrefix16
}

type RichBuildingMetaData struct {
	// Fields
	BuildingType  int32
	BuildingLevel int32

	// Properties
	UpgradeCurrency int32
	UpgradeNum      int32
	FuncType        int32
	Addr1           uint32 `json:"-"`
	Model           int32
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`

	// Objects
	Para         StrWithPrefix16
	Name         StrWithPrefix16
	BuildingDesc StrWithPrefix16
	UpgradeDesc  StrWithPrefix16
	Scale        []float32
	BuildingIcon StrWithPrefix16
}

type RichChallengeBossMetaData struct {
	// Fields
	BossID int32

	// Properties
	AreaID        int32
	UnlockFloorID int32
	StageID       int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	Addr8         uint32 `json:"-"`

	// Objects
	BossImg               StrWithPrefix16
	BossExploreUnlockIcon StrWithPrefix16
	BossExploreLockedIcon StrWithPrefix16
	BossExplorePassedIcon StrWithPrefix16
	BossName              StrWithPrefix16
	RecommendBuff         []RichChallengeBossMetaData_MonopolyBuff
	BossDesc              StrWithPrefix16
	BossUnlockText        StrWithPrefix16
}

type RichChallengeBossMetaData_MonopolyBuff struct {
	// Fields
	BuffID int32
	BuffLV int32
}

type RichDiceMetaData struct {
	// Fields
	DiceID int32

	// Properties
	AnimSkipPer float32
}

type RichFloorInfoMetaData struct {
	// Fields
	Area    int32
	FloorID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	Addr7        uint32 `json:"-"`
	Addr8        uint32 `json:"-"`
	IsResetable  int32
	Addr9        uint32 `json:"-"`
	OverallID    int32
	RpgCoinLimit int32

	// Objects
	LocationList    StrWithPrefix16
	BuildingLimit   []RichFloorInfoMetaData_BuildLimit
	MissionList     []int32
	RichFloorName   StrWithPrefix16
	UnlockTime      RemoteTime
	UnlockText      StrWithPrefix16
	UnlockText_Time StrWithPrefix16
	ShopList        StrWithPrefix16
	RecommendBuff   []RichFloorInfoMetaData_BuffInfo
}

type RichFloorInfoMetaData_BuildLimit struct {
	// Fields
	BuildingMaxLevel  int32
	BuildingMaxNumber int32
	BuildingType      int32
}

type RichFloorInfoMetaData_BuffInfo struct {
	// Fields
	BuffID    int32
	BuffLevel int32
}

type RichItemMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Type  int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	ParaList    []int32
	UseHintText StrWithPrefix16
}

type RichmanMetaData struct {
	// Fields
	RichID int32

	// Properties
	DiceID int32
}

type RichMapInfoMetaData struct {
	// Fields
	ID int32

	// Properties
	RichAreaID int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`

	// Objects
	InfoTitle    StrWithPrefix16
	InfoText     StrWithPrefix16
	InfoIconPath StrWithPrefix16
}

type RichMonsterInfoMetaData struct {
	// Fields
	MonsterID int32

	// Properties
	StageID int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`

	// Objects
	MonsterName   StrWithPrefix16
	MonsterDesc   StrWithPrefix16
	RecommendBuff []RichMonsterInfoMetaData_MonopolyBuff
}

type RichMonsterInfoMetaData_MonopolyBuff struct {
	// Fields
	BuffID int32
	BuffLV int32
}

type RichShopMetaData struct {
	// Fields
	GoodsID int32

	// Properties
	ShopID        int32
	Item          int32
	PurchaseLimit int32
	CostItem      int32
	CostNum       int32
}

type RogueActiveCardMetaData struct {
	// Fields
	ActiveId int32

	// Properties
	EntityClass int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	ChargeCount int32
	Cd          float32
	Param1      float32
	Param1Add   float32
	Param2      float32
	Param2Add   float32
	Param3      float32
	Param3Add   float32

	// Objects
	Name     Hash
	Desc     Hash
	Model    StrWithPrefix16
	IconPath StrWithPrefix16
}

type RogueBuffPoolMetaData struct {
	// Fields
	PoolID uint32

	// Properties
	IsExtra bool
	Addr1   uint32 `json:"-"`

	// Objects
	BuffWeight []RogueBuffPoolMetaData_Pool
}

type RogueBuffPoolMetaData_Pool struct {
	// Fields
	ID     int32
	Weight int32
}

type RogueEliteAbilityMetaData struct {
	// Fields
	EliteAbilityID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	AbilityName        StrWithPrefix16
	AbilityDescription Hash
	IconPath           StrWithPrefix16
}

type RogueEliteAbilityTypeMetaData struct {
	// Fields
	EliteAbilityType int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	EliteAbilityIDList []int32
}

type RogueGoodsMetaData struct {
	// Fields
	GoodID int32

	// Properties
	Type     int32
	SubID    int32
	Price    int32
	Addr1    uint32 `json:"-"`
	Rare     int32
	Category int32

	// Objects
	StoreType []int32
}

type RoguePriceMetaData struct {
	// Fields
	PriceID int32

	// Properties
	Addr1  uint32 `json:"-"`
	Price  int32
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Param1 float32

	// Objects
	Dec       Hash
	StoreType []int32
	Ability   StrWithPrefix16
}

type RogueStageAvatarPosMetaData struct {
	// Fields
	Position int32

	// Properties
	UnlockLevel int32
}

type RogueStageDebuffMetaData struct {
	// Fields
	DebuffID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	DebuffDisplayText Hash
}

type RogueStageHardLevelMetaData struct {
	// Fields
	RogueHard int32

	// Properties
	UnlockNextHardLevel  int32
	RecommandPlayerLevel int32
}

type RogueStageLuaMetaData struct {
	// Fields
	LuaID int32

	// Properties
	Addr1         uint32 `json:"-"`
	EnableStageID int32
	Weight        int32

	// Objects
	Path StrWithPrefix16
}

type RogueStageMetaData struct {
	// Fields
	RogueHard uint8
	Level     uint8

	// Properties
	NpcHardLevel       uint16
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	RepeatRatio        uint8
	EliteAbilityTypeID uint8

	// Objects
	DisplayDropList []int32
	LuaIDList       []int32
}

type RogueStoreMetaData struct {
	// Fields
	ShopID int32

	// Properties
	StoreType     int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	StoreDialog   int32
	SuccessDialog int32
	FailDialog    int32
	CoinPos       float32

	// Objects
	ShopName     Hash
	BaseProp     StrWithPrefix16
	TrueProp     StrWithPrefix16
	BeginPlotIDs []int32
}

type RogueTowerDataMetaData struct {
	// Fields
	ID uint32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	SiteLockTipsParam int32
	BaseHeal          int32
	Addr3             uint32 `json:"-"`
	RogueTowerCoin    int32
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	MaxSaveDataNum    int32

	// Objects
	PoolIDList      []uint32
	ReduceMinScore  []RogueTowerDataMetaData_MaterialReduceMinScore
	ExtraHeal       []RogueTowerDataMetaData_ExtraHealth
	BuffReset       []RogueTowerDataMetaData_ResetMaterialCost
	BuffResetCost   []int32
	BuffLevelupCost []int32
}

type RogueTowerDataMetaData_MaterialReduceMinScore struct {
	// Fields
	AddPercent int32
	MaterialID int32
	MinCount   int32
}

type RogueTowerDataMetaData_ExtraHealth struct {
	// Fields
	AddPercent int32
	CostNum    int32
	MaterialID int32
}

type RogueTowerDataMetaData_ResetMaterialCost struct {
	// Fields
	CostNum      int32
	MaterialID   int32
	MaxResetTime int32
}

type RogueTowerEndSiteMetaData struct {
	// Fields
	AreaID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	EndSiteList []uint32
}

type RogueTowerStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ScoreRate []RogueTowerStageMetaData_ScorePercent
}

type RogueTowerStageMetaData_ScorePercent struct {
	// Fields
	MinScore int32
	Percent  int32
}

type RogueTriggerMetaData struct {
	// Fields
	TriggerID int32

	// Properties
	TriggerType int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Param       float32

	// Objects
	Text           Hash
	TriggerAbility StrWithPrefix16
	DropType       StrWithPrefix16
}

type RoutineDataMetaData struct {
	// Fields
	MissionId int32

	// Properties
	RoutineValue int32
	CountMax     int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	RefreshText      Hash
	MissionImagePath StrWithPrefix16
}

type RoutineRewardMetaData struct {
	// Fields
	WeeklyRewardId int32

	// Properties
	RoutineValue int32
	RewardId     int32
}

type RoutineScheduleMetaData struct {
	// Fields
	ScheduleId int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	WeeklyRewardList []int32
	IntroImagePath   StrWithPrefix16
}

type RpgAreaLineMetaData struct {
	// Fields
	PreAreaID int32
	AreaID    int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	LineNode StrWithPrefix16
}

type RpgAreaMetaData struct {
	// Fields
	AreaID int32

	// Properties
	Addr1             uint32 `json:"-"`
	ZoneID            int32
	Addr2             uint32 `json:"-"`
	AreaRewardPreview int32
	ThemeID           int32
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`
	Addr12            uint32 `json:"-"`
	Addr13            uint32 `json:"-"`

	// Objects
	AreaLinkName      StrWithPrefix16
	SiteList          []int32
	AreaPanelNode     StrWithPrefix16
	AreaTitle         Hash
	AreaDesc          Hash
	AreaPicPath       StrWithPrefix16
	AreaShowPicPath   StrWithPrefix16
	AreaPassedPicPath StrWithPrefix16
	AreaDialogBGPath  StrWithPrefix16
	AreaInfo          []Hash
	AreaLockInfo      Hash
	AreaBuffList      []int32
	AreaMissionList   []int32
}

type RpgBuffDataMetaData struct {
	// Fields
	ID    int32
	Level int32

	// Properties
	MaxLevel       int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Hp             int32
	Atk            int32
	Def            int32
	AllDamageRatio int32
	Addr9          uint32 `json:"-"`
	Addr10         uint32 `json:"-"`
	EffectParam    int32
	CGID           int32

	// Objects
	BuffSuit       RpgBuffDataMetaData_SuitBuffData
	WeaponTypeList []int32
	Name           Hash
	Desc           Hash
	Icon           StrWithPrefix16
	EffectIcon     StrWithPrefix16
	AbilityName    StrWithPrefix16
	ConfigType     StrWithPrefix16
	ParamList      []StrWithPrefix16
	BGColor        StrWithPrefix16
}

type RpgBuffDataMetaData_SuitBuffData struct {
	// Fields
	SuitBuffID int32
	SuitID     int32
	SuitNum    int32
}

type RpgBuffLimitMetaData struct {
	// Fields
	ID    int32
	Level int32

	// Properties
	TaleID int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`

	// Objects
	UnlockFlagList StrWithPrefix16
	CostContent    []RpgBuffLimitMetaData_RpgBuffCostMaterial
}

type RpgBuffLimitMetaData_RpgBuffCostMaterial struct {
	// Fields
	Count      int32
	MaterialID int32
}

type RpgBuffSuitClientInfoMetaData struct {
	// Fields
	BuffDataID int32
	Level      int32

	// Properties
	RpgID  int32
	Addr1  uint32 `json:"-"`
	SuitID int32

	// Objects
	BGColorDisplay StrWithPrefix16
}

type RpgClientMainDataMetaData struct {
	// Fields
	TaleID int32

	// Properties
	TaleType      uint8
	BeginEvent    int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	FuncLinkType  int32
	Addr7         uint32 `json:"-"`
	ShopLink      int32
	MuseumID      int32
	Addr8         uint32 `json:"-"`
	Addr9         uint32 `json:"-"`
	Addr10        uint32 `json:"-"`
	Addr11        uint32 `json:"-"`
	ShowPriority  int32
	Addr12        uint32 `json:"-"`
	Addr13        uint32 `json:"-"`
	EntryShowType int32

	// Objects
	TaleCurrency         []int32
	BuffMaterialList     []int32
	GachaLink            StrWithPrefix16
	TaleWebLink          StrWithPrefix16
	TaleTutorial         []RpgClientMainDataMetaData_Tutorial
	TaleInfo             []StrWithPrefix16
	FuncParam            []int32
	GenCfgPath           StrWithPrefix16
	ContextName          StrWithPrefix16
	WeaponList           []int32
	ShowUnlockSiteList   []int32
	BookPageShowPageList []int32
	MenuShowList         []StrWithPrefix16
}

type RpgClientMainDataMetaData_Tutorial struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Content StrWithPrefix16
	Sprite  StrWithPrefix16
	Title   StrWithPrefix16
}

type RpgCollectionRewardMetaData struct {
	// Fields
	TaleID     int32
	PositionID int32

	// Properties
	CollectionCount int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	RewardID        int32
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	IsSpecialReward int32

	// Objects
	PositionShowIcon            StrWithPrefix16
	LimitTime                   StrWithPrefix16
	ImagePath                   StrWithPrefix16
	RewardProgressUIPrefabPath  StrWithPrefix16
	RewardPointUIPrefabPath     StrWithPrefix16
	RewardPointCurrentIconPath  StrWithPrefix16
	RewardPointLockedIconPath   StrWithPrefix16
	RewardPointFinishedIconPath StrWithPrefix16
}

type RpgDungeonAdventureQuestMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	QuestDetail     StrWithPrefix16
	QuestDetailPic  StrWithPrefix16
	RequiredTagList []int32
	RewardDisplay   []RpgDungeonAdventureQuestMetaData_RewardMaterial
}

type RpgDungeonAdventureQuestMetaData_RewardMaterial struct {
	// Fields
	MaterialID  int32
	MaterialNum int32
}

type RpgEventDialogMetaData struct {
	// Fields
	DialogId int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	UnlockFlagList StrWithPrefix16
}

type RpgEventTextMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	DetailDesc     StrWithPrefix16
	EventName      StrWithPrefix16
	EventDesc      StrWithPrefix16
	EventImagePath StrWithPrefix16
	RewardDesc     StrWithPrefix16
}

type RpgFilesMetaData struct {
	// Fields
	FileID int32

	// Properties
	TaleID        int32
	Addr1         uint32 `json:"-"`
	ImageShowType int32
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`

	// Objects
	FileTitle    Hash
	FileImages   []StrWithPrefix16
	FileText     Hash
	FileLockText Hash
	FileLockHint Hash
	AreaType     StrWithPrefix16
}

type RpgLevelProgressData struct {
	// Fields
	RpgLevel int32

	// Properties
	Progress int32
}

type RpgMaterialMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	TaleID    int32
	ShowType  uint8
	Addr1     uint32 `json:"-"`
	Order     int32
	RpgBuffId int32

	// Objects
	UnlockItemDesc StrWithPrefix16
}

type RpgMissionCategoryMetaData struct {
	// Fields
	TaleID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MissionCategoryList []int32
}

type RpgMissionMetaData struct {
	// Fields
	MissionID int32

	// Properties
	TaleID          int32
	ShowType        int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	IsMainMission   bool
	IsBranchMission bool

	// Objects
	CustomData                 StrWithPrefix16
	MissionIcon                StrWithPrefix16
	RpgTaleMissionBGDesc       Hash
	RpgTaleMissionLocationDesc Hash
	MissionConditionText       Hash
}

type RpgNavSiteMetaData struct {
	// Fields
	SiteID uint16

	// Properties
	Addr1            uint32 `json:"-"`
	NavRefreshSiteID int32
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`

	// Objects
	UICameraID      []float32
	NavNextSiteList []int32
	FogofWar        []int32
}

type RpgOverallMetaData struct {
	// Fields
	ParameterID int32

	// Properties
	IsClientSet bool
	Min         int32
	Max         int32
}

type RpgQAvatarBattleCollideMetaData struct {
	// Fields
	ContactType int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	CollideSet []int32
}

type RpgQAvatarBattleDivisionMetaData struct {
	// Fields
	DivisionId uint16

	// Properties
	Addr1                 uint32 `json:"-"`
	NeedBattleScore       uint32
	FirstDivisionUpReward int32

	// Objects
	DivisionName Hash
}

type RpgQAvatarBattleSceneMetaData struct {
	// Fields
	SceneId int32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	SceneGroundHeight float32

	// Objects
	StagePath     StrWithPrefix16
	ClientGridMap StrWithPrefix16
}

type RpgQAvatarBattleSiteMetaData struct {
	// Fields
	SiteId int32

	// Properties
	Addr1            uint32 `json:"-"`
	BattleTimeMax    uint32
	WinCondition     int32
	FinishBattleKill int16
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	BeatKillNum      int32
	BeatRemainTime   int32

	// Objects
	BGM            StrWithPrefix16
	DieEffect      StrWithPrefix16
	GadgetBuffName StrWithPrefix16
	GadgetBuffIcon StrWithPrefix16
}

type RpgRestaurantMetaData struct {
	// Fields
	RestaurantID int32

	// Properties
	CoinMaterialID     int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	AccelerateTicketID int32

	// Objects
	IngredientList []int32
	RecipeList     []int32
	QuestList      []int32
}

type RpgRoleMetaData struct {
	// Fields
	RPGRoleID int32

	// Properties
	PairAvatarID int32
	RoleType     int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`
	Addr7        uint32 `json:"-"`
	Addr8        uint32 `json:"-"`
	Addr9        uint32 `json:"-"`
	Addr10       uint32 `json:"-"`
	Addr11       uint32 `json:"-"`
	Addr12       uint32 `json:"-"`
	Addr13       uint32 `json:"-"`
	Addr14       uint32 `json:"-"`
	Addr15       uint32 `json:"-"`
	Addr16       uint32 `json:"-"`
	Addr17       uint32 `json:"-"`
	Addr18       uint32 `json:"-"`
	Addr19       uint32 `json:"-"`
	Addr20       uint32 `json:"-"`
	Addr21       uint32 `json:"-"`
	Addr22       uint32 `json:"-"`
	Addr23       uint32 `json:"-"`

	// Objects
	Fake_Avatar_Registry StrWithPrefix16
	Fake_Avatar_Config   StrWithPrefix16
	Color                StrWithPrefix16
	AtkIcon              StrWithPrefix16
	AtkText              Hash
	HpIcon               StrWithPrefix16
	HpText               Hash
	SsIcon               StrWithPrefix16
	SsText               Hash
	SsBase               Hash
	ScIcon               StrWithPrefix16
	ScText               Hash
	ScBase               Hash
	SurvivalNameText     Hash
	SurvivalDescText     Hash
	SurvivalWeaponName   Hash
	SurvivalWeaponIcon   StrWithPrefix16
	SurvivalWeaponDesc   Hash
	QImgPath             StrWithPrefix16
	QDrawPath            StrWithPrefix16
	QPicPath             StrWithPrefix16
	SurvivalGetText      StrWithPrefix16
	QavatarProductTip    StrWithPrefix16
}

type RpgScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	TaleIDList []int32
	BeginTime  StrWithPrefix16
	EndTime    StrWithPrefix16
}

type RpgShootPlayTriggerMetaData struct {
	// Fields
	ID int32

	// Properties
	EventType          int32
	StageID            int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	InLevelTriggerType int32
	Addr3              uint32 `json:"-"`

	// Objects
	Prefab              StrWithPrefix16
	PrefabSpawn         StrWithPrefix16
	InLevelTriggerParam []float32
}

type RpgSimpleBuffMetaData struct {
	// Fields
	BuffLevel      uint16
	BuffMaterialID int32

	// Properties
	Hp             int32
	Atk            uint16
	Def            uint16
	CostNum        uint16
	AllDamageRatio uint16
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`

	// Objects
	BuffEffectDesc StrWithPrefix16
	BuffDesc       StrWithPrefix16
}

type RpgSiteLineMetaData struct {
	// Fields
	PreSiteID int32
	SiteID    int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	LineNode StrWithPrefix16
}

type RpgSiteLockTipsMetaData struct {
	// Fields
	Value int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	SiteLockTips Hash
}

type RpgSiteMetaData struct {
	// Fields
	SiteID int32

	// Properties
	SiteType            uint8
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	Addr9               uint32 `json:"-"`
	StageID             int32
	EnterCoinType       int32
	CoinNum             int32
	Addr10              uint32 `json:"-"`
	Addr11              uint32 `json:"-"`
	Addr12              uint32 `json:"-"`
	Addr13              uint32 `json:"-"`
	Addr14              uint32 `json:"-"`
	SiteUnlockPanelTime float32
	SpecialReward       int32
	Addr15              uint32 `json:"-"`
	Addr16              uint32 `json:"-"`
	Addr17              uint32 `json:"-"`
	ContentID           int32
	LimitID             int32

	// Objects
	SiteNodeName         StrWithPrefix16
	HidePic              StrWithPrefix16
	ShowPic              StrWithPrefix16
	UnlockPic            StrWithPrefix16
	CoolDownPic          StrWithPrefix16
	PassedPic            StrWithPrefix16
	ClosePic             StrWithPrefix16
	TipsPic              StrWithPrefix16
	ShadowPic            StrWithPrefix16
	SiteProgressConfig   []int32
	SiteTitle            Hash
	SiteDes              Hash
	SiteUnlockPanelTitle Hash
	SiteUnlockPanelDes   Hash
	SiteIndexDes         Hash
	SiteLockTips         Hash
	SiteRewardPreview    StrWithPrefix16
}

type RpgSitePackMetaData struct {
	// Fields
	PackID uint16

	// Properties
	AreaID int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`

	// Objects
	UnlockTime RemoteTime
	SiteList   []int32
	PackName   StrWithPrefix16
	PackImage  StrWithPrefix16
}

type RpgSiteProgressConfigMetaData struct {
	// Fields
	SiteProgressID int32

	// Properties
	ProgressType     uint8
	ProgressMaxNum   int32
	Addr1            uint32 `json:"-"`
	ProgressParamInt int32
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`

	// Objects
	ProgressEventList []int32
	ProgressIcon      StrWithPrefix16
	ProgressDesc      StrWithPrefix16
}

type RpgSkillDisplayMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	UniqueTag int32
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`
	Addr6     uint32 `json:"-"`

	// Objects
	TypeName    StrWithPrefix16
	PrefabPath  StrWithPrefix16
	UIPointName StrWithPrefix16
	UILineName  []StrWithPrefix16
	UnlockDesc  Hash
	SkillDesc   Hash
}

type RpgStageDropMetaData struct {
	// Fields
	StageID        int32
	DropMaterialID int32

	// Properties
	Addr1          uint32 `json:"-"`
	MaxCurrencyNum int32

	// Objects
	DailyDropLimit []int32
}

type RpgStageScoreMetaData struct {
	// Fields
	ID int32

	// Properties
	StageID        int32
	MinScore       int32
	CurrencyID     int32
	MaxCurrencyNum int32
	CurrencyNum    int32
}

type RpgStigmataBookMetaData struct {
	// Fields
	StigmataID int32

	// Properties
	GroupID uint32
	Order   uint8
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	IsMain  bool

	// Objects
	UnlockPic      StrWithPrefix16
	UnlockItemDesc StrWithPrefix16
}

type RpgSuddenEventMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	EventName      StrWithPrefix16
	EventDesc      StrWithPrefix16
	EventImagePath StrWithPrefix16
}

type RpgSurvivalRoleMetaData struct {
	// Fields
	SurvivalRolID int32

	// Properties
	RPGRoleID      int32
	QAvatarPVPID   int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	WeaponId       int32
	Star           int32
	AtkBase        int32
	HpBase         int32
	Prop1          int32
	Prop1RPGRoleID int32
	Prop1Param1    float32
	Prop1Param2    float32
	Prop2          int32
	Prop2RPGRoleID int32
	Prop2Param1    float32
	Prop2Param2    float32
	Prop3          int32
	Prop3RPGRoleID int32
	Prop3Param1    float32
	Prop3Param2    float32
	Tspeed         int32
	Mspeed         int32
	TaleID         int32
	RoleBaseScore  int32

	// Objects
	QAvatarRoleName StrWithPrefix16
	AttrIcon        StrWithPrefix16
	AttrImg         StrWithPrefix16
}

type RpgSurvivalTraitMetaData struct {
	// Fields
	TraitID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	TitleText    Hash
	TraitText    Hash
	ItemIcon     StrWithPrefix16
	AbilityName  StrWithPrefix16
	PVPSkillName StrWithPrefix16
}

type RpgSurvivalWeaponMetaData struct {
	// Fields
	SurvivalWeaponID int32

	// Properties
	WeaponType int32
	AttackPara int32
	LockRange  int32
	Range      int32
	AtkSpeed   int32
}

type RpgTaleMetaData struct {
	// Fields
	TaleID int32

	// Properties
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	MinLevel                uint16
	MaxLevel                uint16
	CollectionRewardType    uint8
	Addr3                   uint32 `json:"-"`
	GenActivityRewardShowID uint16
	Addr4                   uint32 `json:"-"`
	VirtualGroupID          int32
	Addr5                   uint32 `json:"-"`
	Addr6                   uint32 `json:"-"`
	RichID                  int32
	Addr7                   uint32 `json:"-"`
	LinkedActivityType      int32
	LinkedActivityID        uint32

	// Objects
	TicketIDList            []int32
	AreaList                []int32
	CollectionRewardParam   []int32
	QAvatarPvpMaterialLimit RpgTaleMetaData_DailyMaterialLimit
	StageList               []int32
	TileTowerList           []int32
	DailyDropLimit          StrWithPrefix16
}

type RpgTaleMetaData_DailyMaterialLimit struct {
	// Fields
	LimitNum   int32
	MaterialID int32
}

type RpgTaleRefreshMetaData struct {
	// Fields
	TaleID int32

	// Properties
	CostHcoinNum int32
	Addr1        uint32 `json:"-"`
	RefreshTimes int32

	// Objects
	CostMaterial []RpgTaleRefreshMetaData_MaterialCost
}

type RpgTaleRefreshMetaData_MaterialCost struct {
	// Fields
	ID  int32
	Num int32
}

type RpgTaleStageEnterTimesMetaData struct {
	// Fields
	LimitID int32

	// Properties
	TaleID     int32
	Addr1      uint32 `json:"-"`
	EnterTimes int32

	// Objects
	StageList []int32
}

type RpgTicketMetaData struct {
	// Fields
	TicketID int32

	// Properties
	MaterialID int32
	MaxNum     int32
}

type RpgTowerV2MetaData struct {
	// Fields
	Stage int32
	Floor int16

	// Properties
	RewardID   int32
	Score      int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	RandomType uint8
	SaveFloor  int16

	// Objects
	DisplayText        Hash
	MonsterDisplayList []int32
	RandomBuffList     []int32
	FixedBuffList      []int32
}

type RpgUIConfigMetaData struct {
	// Fields
	TaleID uint16

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`
	Addr11 uint32 `json:"-"`
	Addr12 uint32 `json:"-"`
	Addr13 uint32 `json:"-"`
	Addr14 uint32 `json:"-"`
	Addr15 uint32 `json:"-"`
	Addr16 uint32 `json:"-"`
	Addr17 uint32 `json:"-"`

	// Objects
	UIManager              StrWithPrefix16
	AlbumPage              StrWithPrefix16
	BookPage               StrWithPrefix16
	PreparePage            StrWithPrefix16
	AvatarUnlockDialog     StrWithPrefix16
	ChapterUnlockDialog    StrWithPrefix16
	RewardDialog           StrWithPrefix16
	SubStageUnlockDialog   StrWithPrefix16
	CollectionRewardDialog StrWithPrefix16
	RpgSiteSelectDialog    StrWithPrefix16
	RpgFileInfoDialog      StrWithPrefix16
	MissionGotDialog       StrWithPrefix16
	EntrancePrefab         StrWithPrefix16
	CoinPrefab             StrWithPrefix16
	BuffPrefab             StrWithPrefix16
	MatrixPage             StrWithPrefix16
	StigmataPage           StrWithPrefix16
}

type RpgVirtualAvatarLimitMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	ConditionList StrWithPrefix16
	ConditionTips Hash
}

type RpgZoneMetaData struct {
	// Fields
	ZoneID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ZoneNode      StrWithPrefix16
	ZoneName      Hash
	UnlockTextMap Hash
}

type SanctuaryBuildingActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	MaxSanctuaryLevel int32
	ShopType          int32
	CurrencyID        int32
	GetTimeCD         int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`

	// Objects
	BgPath       StrWithPrefix16
	AvatarBgPath StrWithPrefix16
}

type SanctuaryBuildingLevelMetaData struct {
	// Fields
	SanctuaryLevel      int32
	SanctuaryActivityID int32

	// Properties
	LevelExp       int32
	Addr1          uint32 `json:"-"`
	ProductOutput  int32
	ProductStorage int32
	Addr2          uint32 `json:"-"`

	// Objects
	BuildingIcon  StrWithPrefix16
	DescTextmapID Hash
}

type SanctuaryPlayerGroupMetaData struct {
	// Fields
	PlayerGroupID int32
	ActivityID    int32

	// Properties
	MinPlayerLevel int32
	MaxPlayerLevel int32
}

type SanctuaryStageConfigMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	RewardExpMin int32
	RewardExpMax int32

	// Objects
	StageRarityIcon       StrWithPrefix16
	StageRarityFrameColor StrWithPrefix16
	StageTypeIcon         StrWithPrefix16
	StageBGIcon           StrWithPrefix16
}

type ScDLCAchievementMetaData struct {
	// Fields
	ID uint32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	SequenceID       int32
	IsShowProgress   bool
	AchieveMissionID int32
	IsHide           bool

	// Objects
	Title Hash
	Desc  Hash
	Icon  StrWithPrefix16
}

type ScDLCAvatarLevelMetaData struct {
	// Fields
	AvatarID int32
	Level    int32

	// Properties
	Addr1           uint32 `json:"-"`
	HPAdd           int32
	ATKAdd          int32
	DEFAdd          int32
	LevelReward     int32
	FeverLevelLimit int32

	// Objects
	LevelUpCost []ScDLCAvatarLevelMetaData_LevelUpRequire
}

type ScDLCAvatarLevelMetaData_LevelUpRequire struct {
	// Fields
	MatID int32
	Num   int32
}

type ScDLCChallengeMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	UnlockTime StrWithPrefix16
}

type ScDLCChallengeRewardMetaData struct {
	// Fields
	ChallengeProgress int32

	// Properties
	StarNum  int32
	RewardID int32
}

type ScDLCExplorePointMaterialMetaData struct {
	// Fields
	ExploreID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MaterialByEvent StrWithPrefix16
}

type ScDLCFeverAbilityMetaData struct {
	// Fields
	ID           int32
	AbilityLevel int32

	// Properties
	AbilityGroupID int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Param          float32
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	UnlockLevel    int32
	ShowLevel      int32
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`

	// Objects
	Ability          StrWithPrefix16
	Icon             StrWithPrefix16
	Desc             Hash
	CurrentLevelDesc Hash
	NextLevelDesc    Hash
	RemainDesc       Hash
	Title            Hash
	ShortName        Hash
}

type ScDLCFurnitureMetaData struct {
	// Fields
	ID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	PrefabPath    StrWithPrefix16
	ConditionList []ScDLCFurnitureMetaData_ConditionStatePair
	NPCHeadMark   []ScDLCFurnitureMetaData_ConditionStatePair
}

type ScDLCFurnitureMetaData_ConditionStatePair struct {
	// Fields
	ConditionID int32
	StateID     int32
}

type ScDLCGachaMetaData struct {
	// Fields
	ID int32

	// Properties
	ActivityID int32
}

type ScDLCImageReplaceMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2     uint32 `json:"-"`
	Condition int32

	// Objects
	SourceImg int32
	TargetImg StrWithPrefix16
}

type ScDLCInteractActionMetaData struct {
	// Fields
	ID int32

	// Properties
	Type  uint8
	Addr1 uint32 `json:"-"`

	// Objects
	Param []StrWithPrefix16
}

type ScDLCInteractStateMetaData struct {
	// Fields
	State int32

	// Properties
	IsShow         bool
	Interactive    bool
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	InteractRadius float32
	InteractAngle  float32
	NPCID          int32
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`

	// Objects
	InteractName            StrWithPrefix16
	InteractIcon            StrWithPrefix16
	InteractEnterActionList []int32
	InteractExitActionList  []int32
}

type ScDLCLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	Exp         uint16
	Addr1       uint32 `json:"-"`
	FeverReward int32
	LevelReward int32
	HighLight   bool
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	NextLevelDesc              Hash
	OwHardlv                   StrWithPrefix16
	MaterialDropLimitAddedList []ScDLCLevelMetaData_ItemLimit
}

type ScDLCLevelMetaData_ItemLimit struct {
	// Fields
	ItemID int32
	Limit  int32
}

type ScDLCMapProgressMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1         uint32 `json:"-"`
	AreaID        int32
	Addr2         uint32 `json:"-"`
	ShowCondition int32

	// Objects
	Desc      StrWithPrefix16
	EventList []int32
}

type ScDLCMissionRewardMetaData struct {
	// Fields
	MissionID int32

	// Properties
	ID            uint16
	TypeID        int32
	Addr1         uint32 `json:"-"`
	ShowCondition int32

	// Objects
	TypeName Hash
}

type ScDLCMPStageDisplayMetaData struct {
	// Fields
	Stage int32

	// Properties
	FirstRewardDisplay int32
	RewardDisplay      int32
}

type ScDLCNPCMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	PrefabScale float32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	PrefabPath    StrWithPrefix16
	ConditionList []ScDLCNPCMetaData_ConditionStatePair
	NPCHeadMark   []ScDLCNPCMetaData_ConditionStatePair
}

type ScDLCNPCMetaData_ConditionStatePair struct {
	// Fields
	ConditionID int32
	StateID     int32
}

type ScDLCNPCPositionGroupMetaData struct {
	// Fields
	Priority int32

	// Properties
	ConditionID int32
	Addr1       uint32 `json:"-"`

	// Objects
	PositionIDList []int32
}

type ScDLCNPCPositionMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1       uint32 `json:"-"`
	ConditionID int32

	// Objects
	NPCPosList []ScDLCNPCPositionMetaData_Positions
}

type ScDLCNPCPositionMetaData_Positions struct {
	// Fields
	AnimID int32
	NpcID  int32
	Addr1  uint32 `json:"-"`

	// Objects
	PostionName StrWithPrefix16
}

type ScDLCOpenFunctionMetaData struct {
	// Fields
	Function uint16

	// Properties
	ConditionID    int32
	Addr1          uint32 `json:"-"`
	ShowUnlockTips bool
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`

	// Objects
	BtnPath    StrWithPrefix16
	Title      Hash
	Desc       Hash
	TitleLabel Hash
	IconPath   StrWithPrefix16
}

type ScDLCShowMapConditionMetaData struct {
	// Fields
	AreaID int32

	// Properties
	ConditionID int32
	Addr1       uint32 `json:"-"`

	// Objects
	MaskPath StrWithPrefix16
}

type ScDLCSkillChipMetaData struct {
	// Fields
	ItemID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Cost      int32
	Addr3     uint32 `json:"-"`
	DisplayID int32
	Addr4     uint32 `json:"-"`
	Addr5     uint32 `json:"-"`

	// Objects
	Icon                 StrWithPrefix16
	Desc                 Hash
	Ability              StrWithPrefix16
	MatchFeverAbilityIDs []int32
	Title                Hash
}

type ScDLCSpecialExplorePointMetaData struct {
	// Fields
	ExploreID int32

	// Properties
	ConditionID   int32
	MapID         int32
	AreaID        int32
	TraceStoryID  int32
	AllowTransfer bool
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`

	// Objects
	RequiredAvatarList []int32
	RequiredAvatarHint Hash
}

type ScDLCSpecialItemMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	RemindType uint8
}

type ScDLCStoryMetaData struct {
	// Fields
	ID    uint32
	State uint8

	// Properties
	Action uint8
}

type ScDLCTeachStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	ShowCondition     int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	SupportAvatar     int32
	SupportAvatarType int32

	// Objects
	AvatarList     []int32
	AvatarListType []int32
}

type ScDLCTowerBonusMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BuffList []int32
	Name     StrWithPrefix16
	Desc     StrWithPrefix16
	IconPath StrWithPrefix16
}

type ScDLCTowerBuffMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	BuffType int32

	// Objects
	BuffName    StrWithPrefix16
	OverrideMap []StrWithPrefix16
}

type ScDLCTowerFloorMetaData struct {
	// Fields
	Floor      int32
	ConfigType int32

	// Properties
	RecommendLevel int32
	WarningLevel   int32
	HardLevel      int32
	SpawnType      int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	TimeBonus      int32
	IsCheckPoint   int32
	Reward         int32
	Addr4          uint32 `json:"-"`
	HardLevelGroup int32
	MaxCoin        int32
	MaxScore       int32
	UnlockStoryID  int32
	MaxTime        int32

	// Objects
	SpawnTypeDesc StrWithPrefix16
	Parameter     []int32
	WaveListPool  []int32
	StageSection  StrWithPrefix16
}

type ScDLCTowerMonsterMetaData struct {
	// Fields
	UniqueID int32

	// Properties
	AddTime int32
	AddCoin int32
}

type ScDLCTowerScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	NextID     int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	ConfigType int32
	LevelID    int32
	MaxFloor   int32
	Weather    int32
	Addr4      uint32 `json:"-"`

	// Objects
	StartTime  StrWithPrefix16
	EndTime    StrWithPrefix16
	SettleTime StrWithPrefix16
	Bonus      []int32
}

type ScDLCTowerScoreRewardMetaData struct {
	// Fields
	Score int32

	// Properties
	Reward int32
}

type ScDLCTowerWaveMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MonsterConfig []int32
}

type ScDLCTowerWeatherMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BuffList []int32
	Name     StrWithPrefix16
	Desc     StrWithPrefix16
	IconPath StrWithPrefix16
}

type ScratchTicketDataMetaData struct {
	// Fields
	PanelID int32

	// Properties
	Addr1         uint32 `json:"-"`
	PlateLevelNum int32
	TicketID      int32
	ResetItemID   int32
	Addr2         uint32 `json:"-"`
	MaxRoundNum   int32
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`

	// Objects
	TicketGoodsIDList   []int32
	ResetItemNumList    []int32
	TicketOutlineImg    StrWithPrefix16
	TicketGoodsImgPath1 StrWithPrefix16
	TicketGoodsImgPath2 StrWithPrefix16
	TicketGoodsImgPath3 StrWithPrefix16
}

type ScratchTicketItemDataMetaData struct {
	// Fields
	PlateID       int32
	PlateDetailID int32

	// Properties
	ItemID     int32
	ItemLv     int32
	ItemNum    int32
	IsCoreItem bool
	Rarity     int32
}

type ScratchTicketPlateDataMetaData struct {
	// Fields
	ScratchPlateID int32

	// Properties
	IsCoreGacha bool
	Addr1       uint32 `json:"-"`

	// Objects
	CardBackImgPath StrWithPrefix16
}

type SecondAnniversaryMemoryMetaData struct {
	// Fields
	MemoryID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	MemoryName  Hash
	MemoryText  Hash
	MemoryImage StrWithPrefix16
}

type SecondAnniversaryMissionMetaData struct {
	// Fields
	MissionID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MissionImage StrWithPrefix16
}

type SeriesMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1         uint32 `json:"-"`
	EnableScore   int32
	SweepMaxTimes uint8

	// Objects
	Title Hash
}

type ServantDialogMetaData struct {
	// Fields
	DialogID int32
	MapId    int32

	// Properties
	Condition int32
	Weight    int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Position  int32
	Time      float32

	// Objects
	Content Hash
	AudioID StrWithPrefix16
	Emotion StrWithPrefix16
}

type ServantTouchMetaData struct {
	// Fields
	Id    int32
	MapID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	SpecialTimes int32
	Addr4        uint32 `json:"-"`
	Addr5        uint32 `json:"-"`
	Addr6        uint32 `json:"-"`

	// Objects
	BodyStateName            StrWithPrefix16
	FaceStateNameList        []StrWithPrefix16
	FaceEffectName           StrWithPrefix16
	SpecialBodyStateName     StrWithPrefix16
	SpecialFaceStateNameList []StrWithPrefix16
	SpecialFaceEffectName    StrWithPrefix16
}

type SettingAudioVolumeMetaData struct {
	// Fields
	Addr1            uint32 `json:"-"`
	AudioSettingType uint8
	ParamIndex       uint8

	// Properties
	Volume float32

	// Objects
	AudioSettingOption StrWithPrefix16
}

type SettingGraphicsDeviceLimitMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Properties
	Addr3 uint32 `json:"-"`

	// Objects
	Device                  StrWithPrefix16
	GraphicsSettingItemName StrWithPrefix16
	ParamBlackList          []uint8
}

type SettingGraphicsItemLineMetaData struct {
	// Fields
	ID        int32
	Hierarchy uint8

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	ItemType         uint8
	OccupiedCols     uint8
	ChildrenColCount uint8
	Addr3            uint32 `json:"-"`

	// Objects
	GraphicsSettingItemText Hash
	GraphicsSettingItemName StrWithPrefix16
	ChildrenLineList        []int32
}

type SettingGraphicsTitleLineMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	ItemColCount uint8

	// Objects
	SettingGraphicsItemLineText Hash
	SettingGraphicsItemLineList []int32
}

type ShareChannelContentMetaData struct {
	// Fields
	ShareContentID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	Title   Hash
	URL     Hash
	Text    Hash
	Topic   Hash
	Section Hash
	Image   StrWithPrefix16
}

type ShareChannelMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Battle          int32
	AvatarData      int32
	Main            int32
	Gacha           int32
	Openworld       int32
	Dormitory       int32
	FeaturedContent int32
	AllowCustom     uint8
	SequenceAndroid int32
	SequenceIOS     int32
	Activity        int32
	Addr2           uint32 `json:"-"`
	AllowWebLink    bool

	// Objects
	AppName      StrWithPrefix16
	ShareBtnIcon StrWithPrefix16
}

type ShareConfigMetaData struct {
	// Fields
	Id int32

	// Properties
	ShareType int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`

	// Objects
	ShareImgPath StrWithPrefix16
	WebviewPath  StrWithPrefix16
	ShareTips    Hash
}

type SharePakDataMetaData struct {
	// Fields
	DownloadType int32

	// Properties
	SubPakId int32
}

type ShareShowDetailConfigMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	ContextName int32
	Paths       []StrWithPrefix16
}

type ShareSpecialMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	PrefabPath StrWithPrefix16
}

type ShareSwitchMetaData struct {
	// Fields
	ChannelId  int32
	DeviceType int16

	// Properties
	MainSwitch        bool
	QrCodeSwitch      bool
	LogoSwitch        bool
	WechatShare       bool
	QqShare           bool
	QzoneShare        bool
	BlogShare         bool
	MiYouShe          bool
	Baidu             bool
	BilibiliShare     bool
	WechatMomentShare bool
	TikTok            bool
}

type ShopAboutToOnlineItem struct {
	// Fields
	ShopID int32
	ItemID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	BeginDate StrWithPrefix16
	EndDate   StrWithPrefix16
}

type ShopDiscountMetaData struct {
	// Fields
	Shopgoodsid int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Discountlist StrWithPrefix16
}

type ShopGoodsClassifyMetaData struct {
	// Fields
	GoodsClassId uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	GoodsClassTitle Hash
}

type ShopGoodsDisplayDataMetaData struct {
	// Fields
	ShopID uint32

	// Properties
	GoodsID            uint32
	Addr1              uint32 `json:"-"`
	ShowMaterialNumber uint32
	DisplayMaterialID  int32
	DisplayMissionID1  int32
	DisplayMissionID2  int32

	// Objects
	ContentDisplay Hash
}

type ShopGoodsJumpMetaData struct {
	// Fields
	ShopGoodsID uint32

	// Properties
	LinkType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`

	// Objects
	LinkParams   []int32
	LinkParamStr StrWithPrefix16
	TextMap      Hash
}

type ShopGoodsMetaData struct {
	// Fields
	ID int32

	// Properties
	ItemID            int32
	ItemLevel         uint8
	ItemNum           int32
	MCoinCost         int32
	HCoinCost         int32
	CostItemId1       int32
	CostItemNum1      int32
	CostItemId2       int32
	CostItemNum2      int32
	CostItemId3       int32
	CostItemNum3      int32
	CostItemId4       int32
	CostItemNum4      int32
	CostItemId5       int32
	CostItemNum5      int32
	MaxBuyTimes       int32
	PriceRateID       uint8
	Discount          int32
	IsSuperWorth      uint8
	IsNew             bool
	CouponType        uint8
	RedDot            bool
	IsBuyMultiple     bool
	Addr1             uint32 `json:"-"`
	GoodsSortId       uint32
	IsAutoOpen        uint8
	BuyMultipleMax    int32
	GlobalMaxBuyTimes int32
	ChooseDialogType  uint8

	// Objects
	GoodsClassIdList []uint16
}

type ShopGoodsPriceRateMetaData struct {
	// Fields
	BuyTimes int32

	// Properties
	PriceType1  int32
	PriceType2  int32
	PriceType3  int32
	PriceType4  int32
	PriceType5  int32
	PriceType6  int32
	PriceType7  int32
	PriceType8  int32
	PriceType9  int32
	PriceType10 int32
}

type ShopGoodsRefreshTimeTypeData struct {
	// Fields
	TypeID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	GiftPackLimitDetailTextmap Hash
	QuotaDescTextmap           Hash
}

type ShopGoodsScheduleMetaData struct {
	// Fields
	Id int32

	// Properties
	GoodsID   int32
	ValidTime uint32
}

type ShopGoodsTagMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	LabelImage StrWithPrefix16
	LabelName  StrWithPrefix16
}

type ShopTabsMetaData struct {
	// Fields
	ShopTabId uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ShopTabTitle Hash
	ShopTypeList []uint32
	TabIconPath  StrWithPrefix16
}

type SignInRewardMetaData struct {
	// Fields
	Month int32
	Day   int32

	// Properties
	RewardItemID int32
	Display      int32
}

type SimulateRankRewardMetaData struct {
	// Fields
	ActivityId  int32
	RankPercent int32

	// Properties
	RewardId int32
}

type SinDemonExBossLevelMetaData struct {
	// Fields
	MonsterLevel int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	LightColor    StrWithPrefix16
	LevelIconPath StrWithPrefix16
	LevelName     Hash
}

type SinDemonExConfigMetaData struct {
	// Fields
	ID int32

	// Properties
	LevelLimitMin   int32
	LevelLimitMax   int32
	ChallengeTimes  int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	ExbossScoreLine int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	CanReChallenge  bool
	IsSweepOn       bool

	// Objects
	RankName            Hash
	RankIcon            StrWithPrefix16
	ActivityEntranceTag Hash
	ActivityTagTitle    Hash
	ActivityDescTitle   Hash
	ActivityDesc        Hash
}

type SinDemonExMonsterMetaData struct {
	// Fields
	BossId int32

	// Properties
	BossGroupId            int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	MonsterId              int32
	HardLevel              int32
	HardLevelGroup         int32
	MonsterHp              int32
	MonsterLevel           uint8
	MonsterBaseScore       int32
	Addr3                  uint32 `json:"-"`
	BossNextId             int32
	BossAttribute          uint8
	Addr4                  uint32 `json:"-"`
	DefaultShowSkillDetail bool
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`
	Addr7                  uint32 `json:"-"`
	Addr8                  uint32 `json:"-"`
	TimesScore             int32
	Addr9                  uint32 `json:"-"`
	Addr10                 uint32 `json:"-"`
	Addr11                 uint32 `json:"-"`
	ExtraTimeScore         int32

	// Objects
	BossName          StrWithPrefix16
	BossPrefabPath    StrWithPrefix16
	SceneName         StrWithPrefix16
	BossSkillTipsList []int32
	BossDesc          Hash
	ImagePath         StrWithPrefix16
	RestrictList      []uint8
	EventMark         Hash
	CornerMarkPath    StrWithPrefix16
	UpTagList         []SinDemonExMonsterMetaData_Tag
	DownTagList       []SinDemonExMonsterMetaData_Tag
}

type SinDemonExMonsterMetaData_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type SinDemonExMonsterScheduleMetaData struct {
	// Fields
	ID       int32
	ConfigId int32

	// Properties
	Addr1        uint32 `json:"-"`
	IsTimesScore bool
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	IsActivity   bool

	// Objects
	BossIdList             []int32
	BossMonsterWeatherList []SinDemonExMonsterScheduleMetaData_BossMonsterWeatherData
	BossScoreRewardList    []SinDemonExMonsterScheduleMetaData_BossMonsterRewardData
	BossTurboModeList      []SinDemonExMonsterScheduleMetaData_BossMonsterTurboData
}

type SinDemonExMonsterScheduleMetaData_BossMonsterWeatherData struct {
	// Fields
	EffectLevel uint8
	Index       int32
	WeatherID   int32
}

type SinDemonExMonsterScheduleMetaData_BossMonsterRewardData struct {
	// Fields
	EffectLevel    uint8
	Index          int32
	RewardConfigID int32
}

type SinDemonExMonsterScheduleMetaData_BossMonsterTurboData struct {
	// Fields
	IsTurboFeatureOn bool
	EffectLevel      uint8
	Index            int32
}

type SinDemonExRankRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	ConfigId   int32
	Rank       int32
	RewardId   int32
	ScheduleId int32
	Addr1      uint32 `json:"-"`

	// Objects
	DivisionName StrWithPrefix16
}

type SinDemonExScoreRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	ConfigId   int32
	Score      int32
	RewardId   int32
	ScheduleId int32
	IsEvent    int32
	TimesScore int32
}

type SinDemonExSeasonBossMetaData struct {
	// Fields
	ConfigID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	ScheduleIDList    []int32
	UpBossList        []int32
	NormalBossList    []int32
	SeasonPageImgPath StrWithPrefix16
}

type SinDemonExSkillTipsMetaData struct {
	// Fields
	SkillTipsID int32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	ShowDetail bool
	ShowCG     bool
	Addr5      uint32 `json:"-"`
	CgID       int32
	Addr6      uint32 `json:"-"`

	// Objects
	SkillNameText      Hash
	SkillNameBriefText Hash
	SkillTipText       Hash
	SkillTipBriefText  Hash
	DetailPicPath      StrWithPrefix16
	BgColor            StrWithPrefix16
}

type SingleRaidActivityMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	StepStageEnterTimes int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	RewardShow          int32

	// Objects
	RewardMissionList []int32
	RewardTitle       Hash
}

type SingleRaidStepStageMetaData struct {
	// Fields
	ChallengeStageID int32

	// Properties
	StepStagePoint int32
	Difficulty     int32
}

type SingleWantedActivityMetaData struct {
	// Fields
	ActivityID uint32

	// Properties
	TicketRefreshNumMax int32
	Addr1               uint32 `json:"-"`
	ResourceMaterialID  uint32
	TicketMaterialID    uint32
	ShopLink            int32
	Addr2               uint32 `json:"-"`

	// Objects
	MpAssistMaterialLimitList []SingleWantedActivityMetaData_DropItem
	StigmataConfigIDList      []uint32
}

type SingleWantedActivityMetaData_DropItem struct {
	// Fields
	ItemID int32
	Num    int32
}

type SingleWantedEquipShowConfigMetaData struct {
	// Fields
	ConfigID uint32

	// Properties
	EquipmentID int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`

	// Objects
	BeginTime          RemoteTime
	EndTime            RemoteTime
	OtherStigmataInfo  SingleWantedEquipShowConfigMetaData_StigmataInfo
	OtherStigmataInfo1 SingleWantedEquipShowConfigMetaData_StigmataInfo
	OtherStigmataInfo2 SingleWantedEquipShowConfigMetaData_StigmataInfo
}

type SingleWantedEquipShowConfigMetaData_StigmataInfo struct {
	// Fields
	Index      int32
	StigmataID int32
}

type SingleWantedMaterialTierMetaData struct {
	// Fields
	MaterialID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	TagDisplay StrWithPrefix16
}

type SingleWantedProgressMetaData struct {
	// Fields
	ThemeID  uint32
	Progress uint32

	// Properties
	Addr1                uint32 `json:"-"`
	FirstCostMaterialNum int32
	CostMaterialNum      int32
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	Addr6                uint32 `json:"-"`
	Addr7                uint32 `json:"-"`

	// Objects
	MPSubTabName         Hash
	ExtraRewardList      []SingleWantedProgressMetaData_MaterialItem
	ExtraDropCostList    []SingleWantedProgressMetaData_MaterialItem
	MpAssistMaterialList []SingleWantedProgressMetaData_MaterialItem
	FirstDropDisplayList []SingleWantedProgressMetaData_MaterialItem
	DropDisplayList      []SingleWantedProgressMetaData_MaterialItem
	WebLink              StrWithPrefix16
}

type SingleWantedProgressMetaData_MaterialItem struct {
	// Fields
	MaterialID int32
	Num        int32
}

type SingleWantedScheduelMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	MinLevel               int32
	SingleWantedActivityID uint32
}

type SingleWantedStageGroupMetaData struct {
	// Fields
	StageGroupID uint32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	StageGroupThemeID uint32
	Addr3             uint32 `json:"-"`

	// Objects
	StageIDList    []int32
	MPStageIDList  []int32
	StageGroupDesc Hash
}

type SingleWantedThemeConfigMetaData struct {
	// Fields
	ThemeID uint32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	MPListDataConfigID int32

	// Objects
	ThemeName          Hash
	ThemeBriefPicPath  StrWithPrefix16
	ThemeDetailPicPath StrWithPrefix16
}

type SkillTutorialStepData struct {
	// Fields
	Id int32

	// Properties
	TutorialID   int32
	Addr1        uint32 `json:"-"`
	StepIndex    int32
	Destroy      int32
	DestroyDelay float32

	// Objects
	SkillStepStr StrWithPrefix16
}

type SkillVideoMetaData struct {
	// Fields
	ID int32

	// Properties
	AvatarID        int32
	SkillID         int32
	SubSkillID      int32
	SkillType       int32
	TutorialVideoID int32
	TutoriallevelID int32
	Addr1           uint32 `json:"-"`

	// Objects
	TutorialDesc Hash
}

type SLGBattlePointMetaData struct {
	// Fields
	PointID int32
	MapID   int32

	// Properties
	Addr1                      uint32 `json:"-"`
	Addr2                      uint32 `json:"-"`
	Stage                      int32
	PointLevel                 int32
	IsEmpty                    bool
	CanAttack                  bool
	Building                   int32
	Addr3                      uint32 `json:"-"`
	MinOccupyThreshold         int32
	FirstOccupyLoyalty         int32
	PointLoyalty               int32
	CanQuickPass               bool
	Addr4                      uint32 `json:"-"`
	Addr5                      uint32 `json:"-"`
	QuickPassAPCost            int32
	BattleAPCost               int32
	QuickPassReward            int32
	QuickPassScore             int32
	Addr6                      uint32 `json:"-"`
	Addr7                      uint32 `json:"-"`
	QuickPassSystemStaminaCost int32

	// Objects
	Name                 StrWithPrefix16
	Desc                 StrWithPrefix16
	NearPointList        []int32
	QuickPassCostList    []SLGBattlePointMetaData_QuickPassMaterial
	BattleCostList       []SLGBattlePointMetaData_BattleMaterial
	BuildingOverrideName StrWithPrefix16
	BuildingOverrideDesc StrWithPrefix16
}

type SLGBattlePointMetaData_QuickPassMaterial struct {
	// Fields
	CostNum    int32
	MaterialID int32
}

type SLGBattlePointMetaData_BattleMaterial struct {
	// Fields
	CostNum    int32
	MaterialID int32
}

type SLGBigBossPointMetaData struct {
	// Fields
	BossScheduleID int32
	BossID         int32

	// Properties
	BossStage int32
	Loyalty   int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
}

type SLGBroadCastMetaData struct {
	// Fields
	ID int32

	// Properties
	ConfigID int32
	Type     int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	ShowTime float32
	IsShow   bool

	// Objects
	Parameter StrWithPrefix16
	Desc      Hash
}

type SLGBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	Icon              StrWithPrefix16
	Desc              Hash
	BuffRatioList     []int32
	BuffRatioDescList []Hash
}

type SLGBuildingMetaData struct {
	// Fields
	ID int32

	// Properties
	Type      int32
	Parameter int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	Model        StrWithPrefix16
	SmallIcon    StrWithPrefix16
	Name         StrWithPrefix16
	BuildingDesc StrWithPrefix16
}

type SLGCountrySetMetaData struct {
	// Fields
	ID int16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	Color              StrWithPrefix16
	Name               Hash
	PanelPic           StrWithPrefix16
	BarColor           StrWithPrefix16
	CharacterSmallIcon StrWithPrefix16
	CountryModel       StrWithPrefix16
	ScoreBarColor      StrWithPrefix16
}

type SLGRankRewardMetaData struct {
	// Fields
	ScheduleID int32
	Rank       int32

	// Properties
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Reward int32

	// Objects
	Name StrWithPrefix16
	Desc StrWithPrefix16
}

type SLGScheduleMetaData struct {
	// Fields
	ID int16

	// Properties
	MaxAP               int32
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	MinLevel            int32
	MaxLevel            int32
	APMaterial          int32
	APRecover           int32
	MainRefreshCD       int32
	Addr3               uint32 `json:"-"`
	Shop                int32
	ShopMaterial        int32
	BuffShopMaterial    int32
	BeginEntryPerformID int32
	EndEntryMission     int32
	EndEntryPerformID   int32
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	QuickPassTicket     int32
	ShopShowID          int32

	// Objects
	SingleBattleList []int32
	MissionList      []int32
	StoryList        []int32
	JsonConfig       StrWithPrefix16
	PhotoOpenTime    RemoteTime
	TaleTutorial     []SLGScheduleMetaData_Tutorial
}

type SLGScheduleMetaData_Tutorial struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Content StrWithPrefix16
	Sprite  StrWithPrefix16
	Title   StrWithPrefix16
}

type SLGScoreRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	ScheduleID int32
	Addr1      uint32 `json:"-"`
	Type       int32
	Score      int32
	Reward     int32

	// Objects
	Desc StrWithPrefix16
}

type SLGShopBuffMetaData struct {
	// Fields
	ID int32

	// Properties
	MaterialID  int32
	MaterialNum int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	AbilityName   StrWithPrefix16
	ShopBuffTitle StrWithPrefix16
	ShopBuffDesc  StrWithPrefix16
}

type SLGSingleBattleMetaData struct {
	// Fields
	ID int16

	// Properties
	MapID               uint16
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	BattleReportCD      int32
	Addr6               uint32 `json:"-"`
	CenterPoint         int32
	PointScoreAccountCD int32
	BattleFog           bool
	FogDistance         int32
	Addr7               uint32 `json:"-"`

	// Objects
	BeginTime        RemoteTime
	EndTime          RemoteTime
	BattleStartTime  RemoteTime
	OpenTime         RemoteTimeSpan
	CloseTime        RemoteTimeSpan
	CountryList      []SLGSingleBattleMetaData_CountryPoint
	BattlePreviewPic StrWithPrefix16
}

type SLGSingleBattleMetaData_CountryPoint struct {
	// Fields
	CountryID int32
	PointID   int32
}

type SLGSmallBossPointMetaData struct {
	// Fields
	ID int32

	// Properties
	BossStage          int32
	BossHP             int32
	FirstOccupyLoyalty int32
}

type SLGStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Reward []SLGStageMetaData_ScoreRewardItem
}

type SLGStageMetaData_ScoreRewardItem struct {
	// Fields
	RewardID int32
	Score    int32
}

type SLGStoryMetaData struct {
	// Fields
	ID int16

	// Properties
	PhotoID         int32
	UnlockMissionID int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	IsShow          bool
	Addr3           uint32 `json:"-"`

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
	Desc      Hash
}

type SLGVoteMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	DanmakuSlotID int32
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`

	// Objects
	Pic                StrWithPrefix16
	StoryID            []int32
	DanmakuColor       StrWithPrefix16
	Desc               StrWithPrefix16
	CharacterSmallIcon StrWithPrefix16
	Name               StrWithPrefix16
}

type SLGVoteScheduleMetaData struct {
	// Fields
	ID int16

	// Properties
	Addr1        uint32 `json:"-"`
	DanmakuID    int32
	VoteMaterial int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`

	// Objects
	VoteIDList    []int32
	VoteBeginTime RemoteTime
	VoteEndTime   RemoteTime
}

type SlotMachineBoxitemMetaData struct {
	// Fields
	BoxItemID int32

	// Properties
	RewardID int32
	Addr1    uint32 `json:"-"`

	// Objects
	CombinationIDList []int32
}

type SlotMachineCombinationMetaData struct {
	// Fields
	CombinationID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	EffectLevel int32

	// Objects
	Combination        []SlotMachineCombinationMetaData_CombElement
	CombinationNumList []SlotMachineCombinationMetaData_CombTier
}

type SlotMachineCombinationMetaData_CombElement struct {
	// Fields
	Count  int32
	IconID int32
}

type SlotMachineCombinationMetaData_CombTier struct {
	// Fields
	IconID int32
	Tier1  int32
	Tier2  int32
	Tier3  int32
}

type SlotMachineIconMetaData struct {
	// Fields
	IconID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	IconImgPath StrWithPrefix16
}

type SlotMachineProgressRewardMetaData struct {
	// Fields
	PanelID          int32
	ProgressRewardID int32

	// Properties
	Progress int32
	RewardID int32
}

type SlotMachineScheduleMetaData struct {
	// Fields
	PanelID int32

	// Properties
	NormalMaterialID  int32
	NormalMaterialNum int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`

	// Objects
	IconIDList      []int32
	SuperIconIDList []int32
}

type SpareShareConfigMetaData struct {
	// Fields
	ConfigID uint8

	// Properties
	DeviceType uint8
	Addr1      uint32 `json:"-"`
	Type       uint8

	// Objects
	ShareChannelID StrWithPrefix16
}

type SpecialAffixDataMetaData struct {
	// Fields
	Id int32

	// Properties
	AffixTreeNodeID int32
	NodeNumRequire  int32
	TriggerMaterial int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`

	// Objects
	RefineCost       []SpecialAffixDataMetaData_MaterialNeed
	PreviewDesc      Hash
	RefineButtonDesc Hash
}

type SpecialAffixDataMetaData_MaterialNeed struct {
	// Fields
	MaterialID  int32
	MaterialNum int32
}

type SpecialConfigMetaData struct {
	// Fields
	ThemeID int32
	Addr1   uint32 `json:"-"`

	// Properties
	ReplaceType int32
	Addr2       uint32 `json:"-"`

	// Objects
	DefaultResPath StrWithPrefix16
	ThemeResPath   StrWithPrefix16
}

type StageDamageStatisticsMetaData struct {
	// Fields
	StageID int32

	// Properties
	ShowElf bool
}

type StageDetailAbilityMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1  uint32 `json:"-"`
	Threat int32
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`

	// Objects
	DisplayTitle  Hash
	SkillTypeList []int32
	SkillInfo     Hash
	DisplayGuide  Hash
}

type StageDetailAccumuatorGaugeMetaData struct {
	// Fields
	ID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	TagID uint16

	// Objects
	Desc Hash
}

type StageDetailEffectMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	EffectText Hash
	EffectIcon StrWithPrefix16
}

type StageDetailLayoutMetaData struct {
	// Fields
	LayoutID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	GroupDisplaySequence1 []int32
	GroupDisplaySequence2 []int32
}

type StageDetailMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1               uint32 `json:"-"`
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`
	Addr6               uint32 `json:"-"`
	Addr7               uint32 `json:"-"`
	Addr8               uint32 `json:"-"`
	StageTagCheckType   uint8
	StageTagCheckNum    uint8
	StageDetailLayoutID int32
	StageRecommendType  uint8

	// Objects
	DisplayMission         Hash
	EffectList             []int32
	EnemyList              []int32
	DisplayGuide           Hash
	RecommendAvatar        []int32
	SummerRanchRecommended []int32
	UpTagList              []StageDetailMetaData_Tag
	DownTagList            []StageDetailMetaData_Tag
}

type StageDetailMetaData_Tag struct {
	// Fields
	TagID int32
	Addr1 uint32 `json:"-"`

	// Objects
	TagComment Hash
}

type StageDetailMonsterMetaData struct {
	// Fields
	ID int32

	// Properties
	Nature            int32
	Addr1             uint32 `json:"-"`
	DisplayDescID     int32
	BossType          int32
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	ModelScale        float32
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`
	Addr8             uint32 `json:"-"`
	AccumuatorGaugeID uint16
	Addr9             uint32 `json:"-"`
	Addr10            uint32 `json:"-"`
	Addr11            uint32 `json:"-"`

	// Objects
	DisplayName    Hash
	AbilitiesList  []int32
	IconPath       StrWithPrefix16
	PrefabPath     StrWithPrefix16
	ControllerPath StrWithPrefix16
	ModelPosition  []float32
	ModelRotation  []float32
	ResistList     []int32
	RecommendTag   []uint16
	UnRecommendTag []uint16
	PageLayout     []uint8
}

type StageDetailMonsterResistMetaData struct {
	// Fields
	ResistID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	StarNum       uint8
	PositiveState uint8

	// Objects
	IconPath   StrWithPrefix16
	ResistName Hash
}

type StageDetailNatureMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type StageDetailSkillTypeMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	DisplayName Hash
	IconPath    StrWithPrefix16
}

type StageDialogMetaData struct {
	// Fields
	DialogID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	AvatarID int32
	Addr4    uint32 `json:"-"`
	Position int32
	Time     float32

	// Objects
	Name    Hash
	Content Hash
	AudioID StrWithPrefix16
	Emotion StrWithPrefix16
}

type StageDifficultySelectMetaData struct {
	// Fields
	LevelID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	DifficultyList []uint8
}

type StageDropItemDataMetaData struct {
	// Fields
	DropID uint32

	// Properties
	ItemID  int32
	ItemNum int32
}

type StageEffectCarryOnMetaData struct {
	// Fields
	StageID int32
}

type StageEnhanceMetaData struct {
	// Fields
	LevelID int32

	// Properties
	Leveltype       int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	AvatarBuffType  int32
	AvatarBuffRatio int32
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	Addr5           uint32 `json:"-"`
	Addr6           uint32 `json:"-"`
	CountDown       int32

	// Objects
	TypePara      StrWithPrefix16
	AvatarList    StrWithPrefix16
	HpRatio       StrWithPrefix16
	AtkRatio      StrWithPrefix16
	DefRatio      StrWithPrefix16
	HardLevelPara StrWithPrefix16
}

type StageLiveRecommendMetaData struct {
	// Fields
	LevelId int32

	// Properties
	Addr1            uint32 `json:"-"`
	MaxRecNum        int32
	MaxRecDetailNum  int32
	PriorityShowType int32

	// Objects
	RecommandTypeParam []int32
}

type StageMaxScoreMetaData struct {
	// Fields
	StageID int32

	// Properties
	MaxScore int32
}

type StageMessageMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Value int32

	// Objects
	KeyName StrWithPrefix16
}

type StageMultiTeamMetaData struct {
	// Fields
	StageID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	A_Position     []int32
	A_MonsterTypes []int32
	B_Position     []int32
	B_MonsterTypes []int32
	C_Position     []int32
	C_MonsterTypes []int32
}

type StageNeedMuteChallengePopupMetaData struct {
	// Fields
	LevelID int32

	// Properties
	NeedMuteChallengePopup bool
}

type StagePhotoMetaData struct {
	// Fields
	StageID int32

	// Properties
	PhotoID int32
}

type StageQAvatarMetaData struct {
	// Fields
	LevelId int32

	// Properties
	BossHP              int32
	RhythmLevel         int32
	RecommendQAvatarNum int32
}

type StageRandomEffectMetaData struct {
	// Fields
	StageID int32

	// Properties
	RandomEffectType int32
	IsCarryOn        bool
}

type StageRankConfigMetaData struct {
	// Fields
	RankID uint32

	// Properties
	MinLevel uint32
	MaxLevel uint32
}

type StageRankGroupConfigMetaData struct {
	// Fields
	RankGroupID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	RankIDList []uint32
}

type StageRechallengeMetaData struct {
	// Fields
	LevelID int32

	// Properties
	Rechallenge int32
}

type StageResistanceMetaData struct {
	// Fields
	LevelID int32

	// Properties
	ResistanceNeed float32
	LevelType      uint8
	AvatarType     uint8
}

type StageScoreRewardDisplayMetaData struct {
	// Fields
	StageID int32

	// Properties
	RewardMaterial int32
	BestTime       uint16
	MaxScore       int32
}

type StageScoreRewardMetaData struct {
	// Fields
	StageID  uint32
	MinScore uint32

	// Properties
	Addr1               uint32 `json:"-"`
	LevelProgressReward uint16
	RemainedTimeReward  uint16

	// Objects
	RewardMaterialList StageScoreRewardMetaData_RewardMaterialData
}

type StageScoreRewardMetaData_RewardMaterialData struct {
	// Fields
	MaterialID  int32
	MaterialNum int32
}

type StageStoryMapMetaData struct {
	// Fields
	LevelID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	QuestID []int32
}

type StageStoryMetaData struct {
	// Fields
	StoryID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Sub      int32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	MaxCount int32
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`

	// Objects
	Name           Hash
	Description    Hash
	Target         Hash
	LuaPath        StrWithPrefix16
	TargetLocation StrWithPrefix16
}

type StageTypeMetaData struct {
	// Fields
	StageTypeID uint8

	// Properties
	StageDetailLayoutID   int32
	ApplicationFocusPause uint8
}

type StepMissionCompensateMetaData struct {
	// Fields
	CompensationID uint8

	// Properties
	UnlockLevel uint8
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	MainLineStepIDList     []int32
	NewChallengeStepIDList []int32
	OldChallengeStepIDList []int32
}

type StigmataAffixMetaData struct {
	// Fields
	AffixID int32

	// Properties
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	PropID     int32
	PropParam1 float32
	PropParam2 float32
	PropParam3 float32
	UIType     int32
	UIValue    float32
	UINature   int32
	UIClass    int32

	// Objects
	NameMono Hash
	NameDual Hash
}

type StigmataBriefMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	SetID  int32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`
	Addr11 uint32 `json:"-"`

	// Objects
	Height     Hash
	Weight     Hash
	HomePlace  Hash
	ActiveSite Hash
	Property   Hash
	Sector     Hash
	Story01    Hash
	Story02    Hash
	Story03    Hash
	Story04    Hash
	Story05    Hash
}

type StigmataMemoirMetaData struct {
	// Fields
	AvatarID int32

	// Properties
	SetID       int32
	OnePieceSet int32
	Unlock      int32
	UiOrder     int32
	ThemeGroup  int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	MaxProcess  int32
	Addr3       uint32 `json:"-"`

	// Objects
	StigmataQNameText Hash
	StigmataQIcon     StrWithPrefix16
	RewardGroupList   []int32
}

type StigmataPositionMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	RootWidth      float32
	RootHeight     float32
	BackEnable     bool
	BackPosX       float32
	BackPosY       float32
	BackWidth      float32
	BackHeight     float32
	Addr2          uint32 `json:"-"`
	ImageWidth     float32
	ImageHeight    float32
	Addr3          uint32 `json:"-"`
	FrontEnable    bool
	FrontPosX      float32
	FrontPosY      float32
	FrontWidth     float32
	FrontHeight    float32
	FrontRotationX float32
	FrontRotationY float32
	FrontRotationZ float32
	Addr4          uint32 `json:"-"`

	// Objects
	Name      StrWithPrefix16
	BackPath  StrWithPrefix16
	ImagePath StrWithPrefix16
	FrontPath StrWithPrefix16
}

type StigmataRuneAffixMetaData struct {
	// Fields
	AffixID int32

	// Properties
	Level          int32
	PropID         int32
	ValueMin       float32
	ValueMax       float32
	ValueStep      float32
	BaseValue      float32
	AddPerPercent  float32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	PropParam2     float32
	PropParam3     float32
	UIType         int32
	UINature       int32
	Addr3          uint32 `json:"-"`
	AffixSkillType int32

	// Objects
	NameMono Hash
	NameDual Hash
	UIClass  []int32
}

type StigmataTagMetaData struct {
	// Fields
	StigmataID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	SpecificAvatarBonus        []StigmataTagMetaData_RatioItem
	SpecificAvatarBonusSupport []StigmataTagMetaData_RatioItem
}

type StigmataTagMetaData_RatioItem struct {
	// Fields
	Pred  int32
	Ratio float32
}

type StigmataThemeMetaData struct {
	// Fields
	ThemeID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ThemeNameText Hash
}

type StigmataVirtualSetMetaData struct {
	// Fields
	SetID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	StigmataMainID   []int32
	SetModify        []StigmataVirtualSetMetaData_RatioItem
	SetModifySupport []StigmataVirtualSetMetaData_RatioItem
}

type StigmataVirtualSetMetaData_RatioItem struct {
	// Fields
	Pred  int32
	Ratio float32
}

type StorySweepGroupInfo struct {
	// Fields
	GroupID int32

	// Properties
	Material int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`

	// Objects
	SweepName Hash
	SweepIcon StrWithPrefix16
}

type StorySweepMetaData struct {
	// Fields
	ID int32

	// Properties
	StageID     int32
	GroupType   int32
	GroupID     int32
	TimeCost    int32
	Addr1       uint32 `json:"-"`
	SelectOrder int32

	// Objects
	RestrictList []int32
}

type SubMallTypeDataMetaData struct {
	// Fields
	ID int32

	// Properties
	FatherID int32
	Addr1    uint32 `json:"-"`
	SortID   int32
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`
	Addr5    uint32 `json:"-"`

	// Objects
	Name                      Hash
	AddOnEventMissionList     []int32
	AddOnEventMissionIconPath StrWithPrefix16
	BeginTime                 StrWithPrefix16
	EndTime                   StrWithPrefix16
}

type SubPakDataMetaData struct {
	// Fields
	SubPakId int32

	// Properties
	DownloadType uint8
	GroupID      int32
}

type SubPakGroupMetaData struct {
	// Fields
	GroupID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	DownloadPriority float32

	// Objects
	GroupName StrWithPrefix16
	GroupDesc StrWithPrefix16
}

type SubsBenefitDataMetaData struct {
	// Fields
	BenefitID int32

	// Properties
	RewardID int32
}

type SubsDataMetaData struct {
	// Fields
	SubsID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	EntryName   StrWithPrefix16
	EntryImg    StrWithPrefix16
	EntryDesc   StrWithPrefix16
	DetailTitle StrWithPrefix16
	DetailImg   StrWithPrefix16
	BenefitList []int32
}

type SubsidiaryEventMetaData struct {
	// Fields
	SubsidiaryEventID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	FaceDelayTime float32
	Addr5         uint32 `json:"-"`

	// Objects
	AvatarOffset           SubsidiaryEventMetaData_Vector3Float
	Path                   []SubsidiaryEventMetaData_DressPath
	SubsidiaryMotion       SubsidiaryEventMetaData_OverrideClip
	Face                   StrWithPrefix16
	AnimatorControllerPath StrWithPrefix16
}

type SubsidiaryEventMetaData_Vector3Float struct {
	// Fields
	X float32
	Y float32
	Z float32
}

type SubsidiaryEventMetaData_OverrideClip struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	ClipPath StrWithPrefix16
	Name     StrWithPrefix16
}

type SubsidiaryEventMetaData_DressPath struct {
	// Fields
	DressID int32
	Addr1   uint32 `json:"-"`

	// Objects
	Path StrWithPrefix16
}

type SubsWareDataMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	SubsID  int32
	SubsDur int32
	Sort    int32
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`
	Addr4   uint32 `json:"-"`
	Addr5   uint32 `json:"-"`
	Addr6   uint32 `json:"-"`

	// Objects
	ProductName StrWithPrefix16
	Promo       StrWithPrefix16
	Renewal     StrWithPrefix16
	BGImg       StrWithPrefix16
	Name        StrWithPrefix16
	Desc        StrWithPrefix16
}

type SupportActivityScheduleMetaData struct {
	// Fields
	SeasonScheduleID int32

	// Properties
	MinLevel                    int32
	MaxLevel                    int32
	Addr1                       uint32 `json:"-"`
	SupportConfigID             int32
	SupportMissionConfigID      int32
	GlobalSupportRewardConfigID int32
	Addr2                       uint32 `json:"-"`
	LinkShopGoodsID             int32
	AvatarID                    int32

	// Objects
	EndTime StrWithPrefix16
	WebLink StrWithPrefix16
}

type SupportAvatarInitialMetaData struct {
	// Fields
	ActivityType uint16
	ActivityID   uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	InitialSupportAvatarIDs []uint16
}

type SupportAvatarLevelMetaData struct {
	// Fields
	SupportAvatarID    uint16
	SupportAvatarLevel uint16

	// Properties
	BuffID uint32
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	CD     int32
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Addr6  uint32 `json:"-"`

	// Objects
	SkillFunctionName StrWithPrefix16
	AbilityName       StrWithPrefix16
	ParamList         []StrWithPrefix16
	SkillIcon         StrWithPrefix16
	SkillName         Hash
	SkillDesc         Hash
}

type SupportAvatarMetaData struct {
	// Fields
	SupportAvatarID uint16

	// Properties
	AvatarID       uint16
	DefaultDressID int32
	ActivityType   uint16
	ActivityID     uint32
	SkillMaxLevel  uint16
	MaterialID     uint32
}

type SupportConfigMetaData struct {
	// Fields
	SupportConfigID int32

	// Properties
	Addr1                    uint32 `json:"-"`
	Addr2                    uint32 `json:"-"`
	Addr3                    uint32 `json:"-"`
	Addr4                    uint32 `json:"-"`
	Addr5                    uint32 `json:"-"`
	Addr6                    uint32 `json:"-"`
	MaxShowNum               int32
	AvatarSupportContextType int32
	UiFormat                 uint8
	ShareRewardID            int32

	// Objects
	BeginTime                        RemoteTime
	EndTime                          RemoteTime
	CountDownTime                    RemoteTime
	AvatarSupportPrefabPath          StrWithPrefix16
	AvatarSupportSharePagePrefabPath StrWithPrefix16
	AvatarSupportMissionPrefabPath   StrWithPrefix16
}

type SupportMissionConfigMetaData struct {
	// Fields
	SupportMissionConfigID int32
	Index                  int32

	// Properties
	MissionID       int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	CgVideoID       int32
	ButtonDisappear bool
	Addr3           uint32 `json:"-"`

	// Objects
	TitleTextMap Hash
	StoryTextMap Hash
	StarIntro    StrWithPrefix16
}

type SurvivalWeaponShowMetaData struct {
	// Fields
	WeaponID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	// Objects
	IconPath             StrWithPrefix16
	WeaponAssetPath      StrWithPrefix16
	WeaponAttachRoot     StrWithPrefix16
	WeaponAttachPosition []float32
	WeaponAttachRotation []float32
	WeaponAnimation      StrWithPrefix16
}

type TeamAssaultBossMetaData struct {
	// Fields
	BossID int32

	// Properties
	BaseScore       int32
	MaxScore        int32
	HardLevel       int32
	HpRatio         int32
	ScoreRatio      int32
	Addr1           uint32 `json:"-"`
	TimeLimit       int32
	MonsterDetailID int32
	HardShow        int32
	Addr2           uint32 `json:"-"`

	// Objects
	RestrictList []int32
	BossNameText Hash
}

type TeamAssaultMetaData struct {
	// Fields
	ActivityID int32

	// Properties
	ActivityEnterTimes int32
	Addr1              uint32 `json:"-"`
	StageID            int32
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`

	// Objects
	BossList   []int32
	UiIconPath StrWithPrefix16
	UiBgPath   StrWithPrefix16
}

type TeamBuffMetaData struct {
	// Fields
	BuffId int32

	// Properties
	Addr1        uint32 `json:"-"`
	Type         int32
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	MaxSuperpose int32
	ParamBase_1  float32
	ParamAdd_1   float32
	ParamBase_2  float32
	ParamAdd_2   float32
	ParamBase_3  float32
	ParamAdd_3   float32

	// Objects
	Name     Hash
	Desc     Hash
	Model    StrWithPrefix16
	IconPath StrWithPrefix16
}

type TextIDReplaceMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2          uint32 `json:"-"`
	ValidMissionID int32

	// Objects
	OriginalTextID Hash
	ReplaceTextID  Hash
}

type TextMapMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	ID   Hash
	Text StrWithPrefix16
}

type ThemeActivityMetaData struct {
	// Fields
	ActivityData int32

	// Properties
	DailyUpCount       int32
	ActivityBonusScore int32
	ActivityBonusRatio float32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`

	// Objects
	ActivityImg   StrWithPrefix16
	ActivityTitle Hash
	ActivityDes   Hash
	BuffList      []int32
	ActivityTips  []Hash
}

type ThemeAvatarMetaData struct {
	// Fields
	AvatarData int32

	// Properties
	Addr1                 uint32 `json:"-"`
	ClientAvatarBonusType int32
	AvatarBonusScore      int32
	AvatarBonusRatio      float32
	Addr2                 uint32 `json:"-"`
	Addr3                 uint32 `json:"-"`
	Addr4                 uint32 `json:"-"`
	Addr5                 uint32 `json:"-"`
	Addr6                 uint32 `json:"-"`

	// Objects
	AvatarIDList      []int32
	AvatarTitle       Hash
	AvatarDes         Hash
	AvatarDisplayList []int32
	BuffList          []int32
	AvatarTips        []Hash
}

type ThemeBonusDataMetaData struct {
	// Fields
	BonusThemeID int32

	// Properties
	BonusThemeType      int32
	Addr1               uint32 `json:"-"`
	DisplayDropHideType int32
	BonusThemeMax       int32
	BonusThemeSingleMax int32
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`
	Addr4               uint32 `json:"-"`
	Addr5               uint32 `json:"-"`

	// Objects
	DisplayDroplist []ThemeBonusDataMetaData_DropItem
	BonusList       []ThemeBonusDataMetaData_BonusItem
	BonusDescTitle  Hash
	BonusDescText   Hash
	BonusTips       []Hash
}

type ThemeBonusDataMetaData_DropItem struct {
	// Fields
	ItemID int32
	Num    int32
}

type ThemeBonusDataMetaData_BonusItem struct {
	// Fields
	ID    int32
	Type  int32
	Value int32
}

type ThemeBuffData struct {
	// Fields
	BuffID int32

	// Properties
	BuffType int32
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Addr3    uint32 `json:"-"`
	Addr4    uint32 `json:"-"`

	// Objects
	BuffNameStrId StrWithPrefix16
	BuffNameText  Hash
	BuffDescStrId Hash
	BuffIcon      StrWithPrefix16
}

type ThemeBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	IsScoreBuff int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	Addr6       uint32 `json:"-"`
	ParamBase_1 float32
	ParamAdd_1  float32
	ParamBase_2 float32
	ParamAdd_2  float32
	ParamBase_3 float32
	ParamAdd_3  float32
	InLevelType int32

	// Objects
	BuffName         Hash
	BuffDes          Hash
	BuffIcon         StrWithPrefix16
	AbilityName      StrWithPrefix16
	ParamEmptyString StrWithPrefix16
	ParamString      StrWithPrefix16
}

type ThemeEquipMetaData struct {
	// Fields
	EquipData int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	WeaponMainIDList   []int32
	StigmataMainIDList []int32
	EquipTitle         Hash
	EquipDes           Hash
	EquipDisplayList   []int32
	BuffList           []int32
	EquipTips          []Hash
}

type ThemeEventImgMetaData struct {
	// Fields
	ImgID uint16

	// Properties
	CanSwitch bool
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	Addr4     uint32 `json:"-"`

	// Objects
	BeginTime RemoteTime
	EndTime   RemoteTime
	ImgPath   []StrWithPrefix16
	GoToLink  []StrWithPrefix16
}

type ThemeEventMetaData struct {
	// Fields
	ActivityID uint16

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	UseNewPanel bool
	Addr4       uint32 `json:"-"`

	// Objects
	BGPrefabPath     StrWithPrefix16
	UIPrefabPath     StrWithPrefix16
	AvatarPrefabPath StrWithPrefix16
	MissionIDList    []int32
}

type ThemeScheduleMetaData struct {
	// Fields
	ThemeID int32

	// Properties
	ThemeType    int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	Addr3        uint32 `json:"-"`
	BonusThemeID int32
	SeasonData   int32
	ActivityData int32
	AvatarData   int32
	EquipData    int32

	// Objects
	ThemeTypePara []int32
	BeginTime     StrWithPrefix16
	EndTime       StrWithPrefix16
}

type ThemeSeasonMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	ThemeStageList []int32
	ThemeTitle     Hash
	ThemeDesc      Hash
	ThemeBGPath    StrWithPrefix16
}

type ThemeSpecialConfigMetaData struct {
	// Fields
	SpecialID int32

	// Properties
	ReplaceType uint8
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`

	// Objects
	DefaultResPath StrWithPrefix16
	ThemeResPath   StrWithPrefix16
}

type ThemeStageMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ThemeStageName   Hash
	ThemeStageBGPath StrWithPrefix16
	ThemeStageDesc   Hash
}

type ThemeStigmataGroup struct {
	// Fields
	StigmataGroupID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	StigmataGroupName     Hash
	StigmataGroupShowList []int32
	StigmataIDList        []int32
}

type ThemeUIConfigMetaData struct {
	// Fields
	ThemeUIID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	MainPage          StrWithPrefix16
	PlayerStatusPanel StrWithPrefix16
	ChatDialogs       StrWithPrefix16
}

type ThemeWeatherAtmosphereMetaData struct {
	// Fields
	AtmosphereConfigID uint8

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	Path StrWithPrefix16
}

type ThemeWeatherConfigMetaData struct {
	// Fields
	WeatherID uint8

	// Properties
	AtmosphereConfigID    uint8
	ReplacePrefabConfigID uint8
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`

	// Objects
	DisplayName Hash
	IconPath    StrWithPrefix16
}

type ThemeWeatherGroupConfigMetaData struct {
	// Fields
	WeatherGroupID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	WeatherIDList        []uint8
	AllLevelsClearedRate []int32
	ChooseRate           []int32
}

type ThemeWeatherReplacePrefabMetaData struct {
	// Fields
	ReplacePrefabConfigID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	ReplacePrefabPath        StrWithPrefix16
	TimeCondition            []ThemeWeatherReplacePrefabMetaData_DayTimeInterval
	RealTimeWeatherCondition []uint8
}

type ThemeWeatherReplacePrefabMetaData_DayTimeInterval struct {
	// Fields
	EndDayTime   uint8
	StartDayTime uint8
}

type ThreadCollectionFileDataMetaData struct {
	// Fields
	ThreadCollectionFileID uint32

	// Properties
	PhotoId            uint32
	HidePhotoID        uint32
	HideWhenLocked     bool
	HideInArchive      bool
	CollectionSiteType int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	PlotID             int32
	Addr3              uint32 `json:"-"`

	// Objects
	Name    StrWithPrefix16
	Desc    StrWithPrefix16
	LuaPath StrWithPrefix16
}

type TileEntityDataMetaData struct {
	// Fields
	EntityID int32

	// Properties
	ModelID int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`

	// Objects
	JsonFile         StrWithPrefix16
	MonopolyJsonFile StrWithPrefix16
}

type TileMapInfoMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	// Objects
	BgMod           StrWithPrefix16
	AudioBank       []StrWithPrefix16
	FloorNameText   Hash
	FloorPanelTitle Hash
	FloorPanelDesc  Hash
	AudioNameEnter  StrWithPrefix16
	AudioNameExit   StrWithPrefix16
}

type TileMapMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Entrance int32
	Addr3    uint32 `json:"-"`

	// Objects
	LuaFile            StrWithPrefix16
	ValueTriggerIDList []int32
	ForwardFork        StrWithPrefix16
}

type TileModelMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	ModelType uint8

	// Objects
	ModelPath  StrWithPrefix16
	ModelScale []float32
}

type TileTerrainDataMetaData struct {
	// Fields
	TerrainID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	OnEnterValueModIDList  []int32
	OnExitValueModIDList   []int32
	OnPassbyValueModIDList []int32
}

type TileValueKeyMetaData struct {
	// Fields
	MapID int32
	Key   int32

	// Properties
	RpgID         int32
	RpgOverallKey int32
}

type TileValueModifierMetaData struct {
	// Fields
	ValueModID int32

	// Properties
	Key                 int32
	Operation           uint16
	Param               int32
	EnablePlayerTrigger bool
	Addr1               uint32 `json:"-"`

	// Objects
	EntityIDList []int32
}

type TileValueTriggerMetaData struct {
	// Fields
	ValueTriggerID int32

	// Properties
	Key                      int32
	ConditionType            uint16
	Param                    int32
	AutoDisablePlayerControl bool
}

type TimePocketMoneyDataMetaData struct {
	// Fields
	ID int32

	// Properties
	RewardId  int32
	TotalTime int32
}

type TouchBuffMetaData struct {
	// Fields
	BuffId int32

	// Properties
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Param1    float32
	Param2    float32
	Param3    float32
	Param1Add float32
	Param2Add float32
	Param3Add float32

	// Objects
	Effect     StrWithPrefix16
	BuffDetail Hash
}

type TouchLevelMetaData struct {
	// Fields
	Level int32

	// Properties
	TouchExp             int32
	Prop                 float32
	Rate                 float32
	BattleGain           int32
	AddGoodfeelPerMinute int32
}

type TouchMetaData struct {
	// Fields
	TouchId int32

	// Properties
	Point int32
	Addr1 uint32 `json:"-"`

	// Objects
	Buff []int32
}

type TowerBuffConfigMetaData struct {
	// Fields
	TowerBuffConfigID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	TowerBuffIcon StrWithPrefix16
	TowerBuffDes  Hash
}

type TowerGrowBuffConfigMetaData struct {
	// Fields
	BuffLevelID int32
	ActivityID  int32

	// Properties
	HP             int32
	ATK            int32
	DEF            int32
	AllDamageRatio int32
}

type TowerRaidBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	BuffIDGroup int32
	Addr1       uint32 `json:"-"`
	StagePos    int32
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`

	// Objects
	BuffTag  StrWithPrefix16
	Textmap  Hash
	BuffIcon StrWithPrefix16
}

type TowerRaidConfigMetaData struct {
	// Fields
	TowerRaidID int32

	// Properties
	MinDifficulty    int32
	MaxDifficulty    int32
	TowerRaidGroupID int32
}

type TowerRaidDifficultyMetaData struct {
	// Fields
	ID int32

	// Properties
	AddHardLevel      int32
	AvatarPlayTimes   int32
	UpAvatarPlayTimes int32
	BuffSwitch        bool
	Reward_1          int32
	Reward_2          int32
	Reward_3          int32
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	Addr5             uint32 `json:"-"`
	Addr6             uint32 `json:"-"`
	Addr7             uint32 `json:"-"`

	// Objects
	FinalDisplayDropList   StrWithPrefix16
	DifficultyDesc         Hash
	DifficultyName         Hash
	RewardDesc_1           Hash
	RewardDesc_2           Hash
	RewardDesc_3           Hash
	StageDropDisplayList_1 StrWithPrefix16
}

type TowerRaidScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1          uint32 `json:"-"`
	ThemeAvatar    int32
	Addr2          uint32 `json:"-"`
	DailyPlayTimes int32

	// Objects
	EndTime     StrWithPrefix16
	BuffGroupID []int32
}

type TowerRaidStageMetaData struct {
	// Fields
	TowerRaidID int32

	// Properties
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`
	MaxStageCount int32
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`

	// Objects
	StageIDList          []int32
	StageNameList        []int32
	RewardCheckPointList []int32
	StageBGPrefabPath    StrWithPrefix16
	StageDecorationPath  StrWithPrefix16
}

type TowerStageConfigMetaData struct {
	// Fields
	StageConfigID int32

	// Properties
	Stageid                int32
	PreStage               int32
	Reward                 int32
	RewardShow             int32
	RecommendGrowBuffLevel int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`

	// Objects
	CgExtraKey                 StrWithPrefix16
	TowerBuffConfig            []int32
	DLCChallengeSpawnPointName StrWithPrefix16
	StageUIPath                StrWithPrefix16
}

type TradingCardActivityScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Reward      int32
	CanTradeNum int32

	// Objects
	OpenTime           RemoteTime
	CloseTime          RemoteTime
	NeedMaterialIDList []TradingCardActivityScheduleMetaData_MaterialNeeded
}

type TradingCardActivityScheduleMetaData_MaterialNeeded struct {
	// Fields
	MaterialID int32
	Num        int32
}

type TrainingLevelAchieveMetaData struct {
	// Fields
	MissionID int32

	// Properties
	AvatarID int32
}

type TrainingLevelData struct {
	// Fields
	LevelID int32

	// Properties
	AvatarID    int32
	AvatarLevel int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`

	// Objects
	AvatarSubskillList []int32
	HintMessage        Hash
}

type TutorialData struct {
	// Fields
	Id int32

	// Properties
	Index               int32
	TriggerMinLevel     int32
	TriggerMaxLevel     int32
	TriggerMissionID    int32
	TriggerOnDoing      bool
	TriggerOnFinish     bool
	TriggerOnClose      bool
	Addr1               uint32 `json:"-"`
	TriggerSpecial      int32
	Addr2               uint32 `json:"-"`
	StartStepID         int32
	Reward              int32
	SkipGroup           int32
	SkipFinishMissionID int32

	// Objects
	TriggerUIContextName StrWithPrefix16
	TriggerSpecialParam  StrWithPrefix16
}

type TutorialGraphicMetaData struct {
	// Fields
	Addr1  uint32 `json:"-"`
	StepId uint8

	// Properties
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	// Objects
	TutorialKey  StrWithPrefix16
	ImagePath    StrWithPrefix16
	TextTitle    Hash
	TextContent  Hash
	TextContent2 Hash
}

type TutorialRestrictMetaData struct {
	// Fields
	TutorialID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ExcludeTutorialList []int32
}

type TutorialStepData struct {
	// Fields
	Id int32

	// Properties
	StepType               uint8
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	HandIconPosType        uint8
	PlayEffect             bool
	Addr4                  uint32 `json:"-"`
	OverrideLink           int32
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`
	BubblePosType          uint8
	NextStepID             int32
	DelayTime              float32
	Addr7                  uint32 `json:"-"`
	Addr8                  uint32 `json:"-"`
	Addr9                  uint32 `json:"-"`
	Addr10                 uint32 `json:"-"`
	Addr11                 uint32 `json:"-"`
	Addr12                 uint32 `json:"-"`
	Addr13                 uint32 `json:"-"`
	DisableHighlightInvoke uint8
	Addr14                 uint32 `json:"-"`
	Addr15                 uint32 `json:"-"`

	// Objects
	RootPath              StrWithPrefix16
	TargetUIPath          StrWithPrefix16
	TargetButtonPath      StrWithPrefix16
	ScrollUIPath          StrWithPrefix16
	LinkParams            []int32
	GuideDesc             Hash
	FinishOnStart         []int32
	FinishOnEnd           []int32
	FacePath              StrWithPrefix16
	TargetCamera          StrWithPrefix16
	TutorialSoundBankName StrWithPrefix16
	TutorialVoiceEvent    StrWithPrefix16
	AnimationTutorialPath StrWithPrefix16
	TargetUIManager       StrWithPrefix16
	TargetCamPosition     []float32
}

type TutorialStepScrollerData struct {
	// Fields
	StepID int32

	// Properties
	DelayTime   float32
	TargetIndex int32
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	UseTextID   bool

	// Objects
	TargetSubPathList []StrWithPrefix16
	ComparePath       StrWithPrefix16
	CompareType       StrWithPrefix16
	CompareField      StrWithPrefix16
	CompareValue      StrWithPrefix16
}

type TutorialVideoMetaData struct {
	// Fields
	TutorialVideoID int32

	// Properties
	GroupID int32
	StepId  int32
	CgID    int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`

	// Objects
	TextTitle Hash
	ImagePath StrWithPrefix16
	TextMAPID Hash
}

type TutorialWeeklyScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Day        int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	TutorialID int32

	// Objects
	Begin StrWithPrefix16
	End   StrWithPrefix16
}

type TVTAvatarConfigMetaData struct {
	// Fields
	AvatarID int32
	DressID  int32

	// Properties
	OffsetX float32
	OffsetY float32
	Scale   float32
}

type TVTCardLevelMetaData struct {
	// Fields
	ID    int32
	Level uint8

	// Properties
	AffectRange uint8
	CardCost    uint16
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	Addr3       uint32 `json:"-"`
	Addr4       uint32 `json:"-"`
	Addr5       uint32 `json:"-"`
	Addr6       uint32 `json:"-"`

	// Objects
	Description        Hash
	ShortDescription   Hash
	EffectFunctionName StrWithPrefix16
	ParamList          []StrWithPrefix16
	UniqueMaterialList []StrWithPrefix16
	CommonMaterialList []StrWithPrefix16
}

type TVTCardMetaData struct {
	// Fields
	ID int32

	// Properties
	CardType   uint8
	MaxLevel   uint8
	Rarity     uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	UnlockRank uint8

	// Objects
	CardName Hash
	IconPath StrWithPrefix16
}

type TVTCardSlotMetaData struct {
	// Fields
	ID         uint16
	IsUnlocked bool
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Properties
	Addr3            uint32 `json:"-"`
	UnlockDivisionLv uint16

	// Objects
	UnlockConditionHint StrWithPrefix16
	UnlockCondition     StrWithPrefix16
	UnlockTime          StrWithPrefix16
}

type TVTConstValuesMetaData struct {
	// Fields
	ScheduleType uint32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	TypeName           Hash
	TypeInfo           Hash
	TypeTutorialIDList []int32
}

type TVTDivisionMetaData struct {
	// Fields
	DivisionID uint32
	SeasonID   uint32

	// Properties
	Addr1                   uint32 `json:"-"`
	Addr2                   uint32 `json:"-"`
	Addr3                   uint32 `json:"-"`
	DivisionNeedStar        uint32
	FirstDivisionUpReward   int32
	DivisionProtectMaterial int32
	DivisionProtectUseID    uint8
	SeasonOffReward         int32

	// Objects
	ChildDivisionName Hash
	DivisionName      Hash
	DivisionIconPath  StrWithPrefix16
}

type TVTDivisionProtectUseMetaData struct {
	// Fields
	ProtectUseID uint8

	// Properties
	DivisionProtectCost      uint16
	DivisionProtectMaxNum    uint32
	DivisionProtectInitValue uint16
}

type TVTInLevelInteractionMetaData struct {
	// Fields
	ID int32

	// Properties
	UIPositionID int32
	Type         int32
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	ShowDescription StrWithPrefix16
	Detail          StrWithPrefix16
}

type TVTMaterialTransformMetaData struct {
	// Fields
	MaterialID      int32
	TransformCardID uint16
}

type TVTMessageConfigMetaData struct {
	// Fields
	MessageID      uint32
	HardLevelGroup uint32
	HardLevel      uint32

	// Properties
	MessagePara uint32
}

type TVTMessageMetaData struct {
	// Fields
	MessageID uint32

	// Properties
	MessageType          uint8
	MessagePara          uint32
	StageID              int32
	IsBattleWinCondition bool
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`

	// Objects
	BossName Hash
	BossIcon StrWithPrefix16
}

type TVTMissionMetaData struct {
	// Fields
	MissionId int32

	// Properties
	MissionType  uint8
	NeedDivision uint32
	Addr1        uint32 `json:"-"`

	// Objects
	IconPath StrWithPrefix16
}

type TVTRobotDataMetaData struct {
	// Fields
	RobotID uint32

	// Properties
	Addr1   uint32 `json:"-"`
	HeadID  int32
	StageID int32
	Addr2   uint32 `json:"-"`
	ElfID   int32
	Addr3   uint32 `json:"-"`

	// Objects
	RobotName    Hash
	AvatadIDList []int32
	CardIDList   []int32
}

type TVTScheduleMetaData struct {
	// Fields
	ScheduleID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ScheduleMissionList []int32
}

type TVTSeasonMetaData struct {
	// Fields
	SeasonID uint32

	// Properties
	Addr1                uint32 `json:"-"`
	Addr2                uint32 `json:"-"`
	Addr3                uint32 `json:"-"`
	ShopId               uint32
	RewardShowID         int32
	Addr4                uint32 `json:"-"`
	Addr5                uint32 `json:"-"`
	CardSystemUnlockType uint16

	// Objects
	SeasonName   Hash
	TutorialList []int32
	WebLink      StrWithPrefix16
	CurrencyID   []int32
	CardSlotList []int32
}

type TVTStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	MaxEnergyDrop      int32
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	TVTTutorialStageID int32

	// Objects
	MessageIDList     []int32
	TVTTutorialIDList []int32
}

type TVTTutorialMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	PicPath    StrWithPrefix16
	DescTitle  StrWithPrefix16
	DescDetail StrWithPrefix16
}

type UIOverrideItemMetaData struct {
	// Fields
	ID uint32

	// Properties
	OverrideType uint8
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`

	// Objects
	NodePath      StrWithPrefix16
	OverrideParam StrWithPrefix16
}

type UIOverrideScheduleMetaData struct {
	// Fields
	ID uint32

	// Properties
	UIType uint16
	Addr1  uint32 `json:"-"`
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`

	// Objects
	BeginTime      RemoteTime
	EndTime        RemoteTime
	UIOverrideList []uint32
}

type UltraEndlessBaseRewardMetaData struct {
	// Fields
	ID uint8

	// Properties
	GroupLevel uint8
	MinScore   int32
	RewardID   int32
}

type UltraEndlessBattleConfigMetaData struct {
	// Fields
	BattleConfigID uint32

	// Properties
	Addr1                 uint32 `json:"-"`
	Addr2                 uint32 `json:"-"`
	Weeklyreport_BossInfo int32
	Weeklyreport_BuffInfo int32

	// Objects
	SiteMapName   StrWithPrefix16
	MissionIDList []int32
}

type UltraEndlessBuffMetaData struct {
	// Fields
	BuffID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	BuffName     StrWithPrefix16
	IconPath     StrWithPrefix16
	BuffNameText Hash
	BuffDescText Hash
}

type UltraEndlessDisplayMetaData struct {
	// Fields
	ID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	UltraEndlessTitle StrWithPrefix16
	GroupIcon         StrWithPrefix16
	GroupSmallIcon    StrWithPrefix16
	GroupIconColor    StrWithPrefix16
}

type UltraEndlessFloorMetaData struct {
	// Fields
	FloorID uint8
	StageID int32

	// Properties
	NeedScore uint16
	MaxScore  uint16
}

type UltraEndlessGroupMetaData struct {
	// Fields
	GroupLevel uint8

	// Properties
	GroupID    uint8
	Addr1      uint32 `json:"-"`
	NeedCupNum uint16

	// Objects
	GroupName StrWithPrefix16
}

type UltraEndlessRewardMetaData struct {
	// Fields
	BeginScheduleId uint32
	GroupLevel      uint8

	// Properties
	EndScheduleId uint32
	SettleReward  uint32
}

type UltraEndlessScheduleMetaData struct {
	// Fields
	ScheduleID uint16

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	StartTime        RemoteTime
	BattleConfigList []UltraEndlessScheduleMetaData_GroupLevelBattleConfig
}

type UltraEndlessScheduleMetaData_GroupLevelBattleConfig struct {
	// Fields
	GroupLevel     int32
	BattleConfigID uint32
}

type UltraEndlessSettleCupMetaData struct {
	// Fields
	GroupLevel       uint8
	PlayerNumInGroup uint8

	// Properties
	OnePlayerNoAttendCup int16
	Addr1                uint32 `json:"-"`

	// Objects
	CupChangeList []int16
}

type UltraEndlessSiteMetaData struct {
	// Fields
	SiteID int32

	// Properties
	StageID        int32
	BuffID         int32
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	Addr4          uint32 `json:"-"`
	HardLevelGroup int32
	Addr5          uint32 `json:"-"`

	// Objects
	SiteNodeName    StrWithPrefix16
	PreSiteList     []int32
	SiteName        StrWithPrefix16
	LevelEndWebLink StrWithPrefix16
	WebLink         StrWithPrefix16
}

type UltraEndlessStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	TutorialStageID int32
}

type UltraEndlessTopRankShowMetaData struct {
	// Fields
	ID uint8

	// Properties
	RankMin  uint16
	RankMax  uint16
	RewardId int32
}

type UniqueGachaPoolMetaData struct {
	// Fields
	PoolId uint16
	ItemId uint32
}

type UniqueGachaReplaceItemMetaData struct {
	// Fields
	GachaId uint16

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	CandidatePoolList []UniqueGachaReplaceItemMetaData_CandidatePool
}

type UniqueGachaReplaceItemMetaData_CandidatePool struct {
	// Fields
	SelectNum uint8
	PoolID    uint16
}

type UniqueGoodsItemMetaData struct {
	// Fields
	ItemID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ObjectName StrWithPrefix16
}

type UniqueMonsterMetaData struct {
	// Fields
	UniqueID uint32

	// Properties
	Addr1          uint32 `json:"-"`
	Addr2          uint32 `json:"-"`
	Addr3          uint32 `json:"-"`
	HPRatio        float32
	AttackRatio    float32
	DefenseRatio   float32
	MoveSpeedRatio float32
	Addr4          uint32 `json:"-"`
	Addr5          uint32 `json:"-"`
	Addr6          uint32 `json:"-"`
	Addr7          uint32 `json:"-"`
	Addr8          uint32 `json:"-"`
	Addr9          uint32 `json:"-"`
	HPSegmentNum   int32
	Addr10         uint32 `json:"-"`
	BossRank       int32
	HandBookId     int32
	MonsterRank    int32

	// Objects
	Name          Hash
	MonsterName   StrWithPrefix16
	TypeName      StrWithPrefix16
	ATKRatios     []float32
	ConfigType    StrWithPrefix16
	AIName        StrWithPrefix16
	AttackCDNames []StrWithPrefix16
	AttackCDs     []float32
	Abilities     StrWithPrefix16
	Scale         []float32
}

type UnlockUIData struct {
	// Fields
	Id int32

	// Properties
	UnlockByMission  int32
	OnDoing          int32
	OnFinish         int32
	OnClose          int32
	UnlockByTutorial int32
}

type VibrateConfigsMetaData struct {
	// Fields
	Addr1 uint32 `json:"-"`

	// Properties
	Addr2 uint32 `json:"-"`

	// Objects
	AvatarRegistryKey StrWithPrefix16
	ConfigPath        StrWithPrefix16
}

type VirtualAvatarMetaData struct {
	// Fields
	Id int32

	// Properties
	AvatarType    int32
	AvatarId      int32
	PvpAvatarId   int32
	ArtifactLevel int32
	BaseRarity    int32
	BaseLevel     int32
	DefaultWeapon int32
	DefaultDress  int32
	PreId         int32
	Addr1         uint32 `json:"-"`
	Addr2         uint32 `json:"-"`
	Addr3         uint32 `json:"-"`

	// Objects
	EvoMaterialList        []VirtualAvatarMetaData_EvoMaterialItem
	UnlockMaterialList     []VirtualAvatarMetaData_Condition
	UnlockStigmataSlotList StrWithPrefix16
}

type VirtualAvatarMetaData_EvoMaterialItem struct {
	// Fields
	Amount int32
	ItemID int32
}

type VirtualAvatarMetaData_Condition struct {
	// Fields
	Count      int32
	MaterialID int32
}

type VirtualBuffMetaData struct {
	// Fields
	Id int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	BuffId   StrWithPrefix16
	BuffIcon StrWithPrefix16
}

type VirtualCustomLevelDataMetaData struct {
	// Fields
	CustomID int32
	Level    int32

	// Properties
	CostMaterialID int32
	Cost           int32
	Addr1          uint32 `json:"-"`
	UnlockStageID  int32
	Addr2          uint32 `json:"-"`

	// Objects
	VirtualItemList []StrWithPrefix16
	PreCustomLevel  StrWithPrefix16
}

type VirtualCustomMetaData struct {
	// Fields
	CustomID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`
	Addr8 uint32 `json:"-"`

	// Objects
	CustomName            Hash
	IconPath_1            StrWithPrefix16
	IconPath_2            StrWithPrefix16
	IconPath_3            StrWithPrefix16
	IconPath_4            StrWithPrefix16
	VirtualItemNameList   []Hash
	VirtualItemDescList   []Hash
	VirtualItemLockedList []Hash
}

type VirtualGachaChestMetaData struct {
	// Fields
	Id int32

	// Properties
	GroupID          int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	IsOpenAfterStage bool

	// Objects
	GachaPoolList      []VirtualGachaChestMetaData_GachaPool
	SelectChestImgPath StrWithPrefix16
	ChestName          Hash
}

type VirtualGachaChestMetaData_GachaPool struct {
	// Fields
	GachaPoolID int32
	MaterialID  int32
}

type VirtualGachaItemMetaData struct {
	// Fields
	Id int32

	// Properties
	GachaPoolId int32
	DropId      int32
}

type VirtualGroupConfigMetaData struct {
	// Fields
	GroupID int32

	// Properties
	GroupTypeParam     int32
	Addr1              uint32 `json:"-"`
	MaxTeamMemberCount int32

	// Objects
	DefaultVirtualAvatar []int32
}

type VirtualRoleMetaData struct {
	// Fields
	TrainRoleID int32

	// Properties
	RoleID                int32
	VirtualGroupID        int32
	RoleSectionID         int32
	UnlockVirtualAvatarId int32
	Addr1                 uint32 `json:"-"`

	// Objects
	UnlockTips Hash
}

type VirtualStigmataDetailMetaData struct {
	// Fields
	Id int32

	// Properties
	Star            uint8
	Rarity          int32
	Addr1           uint32 `json:"-"`
	Addr2           uint32 `json:"-"`
	Addr3           uint32 `json:"-"`
	Addr4           uint32 `json:"-"`
	PveHp           int32
	PveAttack       int32
	Addr5           uint32 `json:"-"`
	PveAbilityLevel int32
	PvpHp           int32
	PvpAttack       int32
	Addr6           uint32 `json:"-"`
	Addr7           uint32 `json:"-"`
	Addr8           uint32 `json:"-"`
	Addr9           uint32 `json:"-"`
	Addr10          uint32 `json:"-"`
	Addr11          uint32 `json:"-"`

	// Objects
	DisplayTitle       StrWithPrefix16
	BagIconPath        StrWithPrefix16
	BagDetailIconPath  StrWithPrefix16
	SlotIconPath       StrWithPrefix16
	PveAbilityName     StrWithPrefix16
	PveBuffDes         Hash
	PvpBuffDes         Hash
	PveBuffShortDes    Hash
	PvpBuffShortDes    Hash
	CommonBuffShortDes Hash
	PveAbilityParam    []float32
}

type VirtualStigmataMetaData struct {
	// Fields
	Id int32

	// Properties
	StigmataId        int32
	BaseRarity        int32
	BaseLevel         int32
	IsAllowMultiDress bool
	PreId             int32
	Type              uint8
}

type VirtualTrainStageDropMetaData struct {
	// Fields
	StageId  int32
	EndFloor int16

	// Properties
	DropTimes int32
}

type VirtualWeaponMetaData struct {
	// Fields
	Id int32

	// Properties
	WeaponId   int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type WarehouseRequireData struct {
	// Fields
	ID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Addr2        uint32 `json:"-"`
	ItemID       int32
	ItemNum      int32
	CostSC       int32
	CostItem     int32
	CostItemNum  int32
	RequireLevel int32
	CD           int32
	Reward       int32
	Obsolete     int32

	// Objects
	BeginTime StrWithPrefix16
	EndTime   StrWithPrefix16
}

type WaveRushBuffDataMetaData struct {
	// Fields
	BuffID int32

	// Properties
	Addr1        uint32 `json:"-"`
	Type         uint8
	Addr2        uint32 `json:"-"`
	BuffSect     uint8
	Addr3        uint32 `json:"-"`
	Addr4        uint32 `json:"-"`
	IsShowDialog bool

	// Objects
	BuffName     StrWithPrefix16
	ActivateCore []WaveRushBuffDataMetaData_SectTuple
	Name         Hash
	Icon         StrWithPrefix16
}

type WaveRushBuffDataMetaData_SectTuple struct {
	// Fields
	Type uint8
	Num  int32
}

type WaveRushBuffLevelGroupMetaData struct {
	// Fields
	BuffGroupConfigID int32
	BuffID            int32
	Level             int32

	// Properties
	ShopIndex              int32
	LevelUpCostMaterialNum int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`
	LimitConditionType     uint8
	Addr5                  uint32 `json:"-"`
	Addr6                  uint32 `json:"-"`

	// Objects
	ParamList           []float32
	BasePropertyList    []int32
	Description         Hash
	SimpleDescription   Hash
	LimitConditionParam []StrWithPrefix16
	LockedTips          Hash
}

type WaveRushFuncOpenMetaData struct {
	// Fields
	FuncID uint8

	// Properties
	ShowConditionType   uint8
	Addr1               uint32 `json:"-"`
	UnlockConditionType uint8
	Addr2               uint32 `json:"-"`
	Addr3               uint32 `json:"-"`

	// Objects
	ShowConditionParam   []StrWithPrefix16
	UnlockConditionParam []StrWithPrefix16
	LockedTips           Hash
}

type WaveRushScheduleMetaData struct {
	// Fields
	ScheduleID int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	MinLevel           int32
	UIConfigID         int32
	StageGroupConfigID int32
	BuffGroupConfigID  int32
	Addr3              uint32 `json:"-"`
	MaxCoreBuffNum     uint8
	Addr4              uint32 `json:"-"`
	RewardLineID       int32
	Addr5              uint32 `json:"-"`
	ActivityMaterialID int32
	Addr6              uint32 `json:"-"`
	EntryShowType      uint8
	TutorialStage      int32

	// Objects
	BeginTime          StrWithPrefix16
	EndTime            StrWithPrefix16
	CoreBuff           []int32
	SlotNumPreSiteList []WaveRushScheduleMetaData_LimitGroup
	Mission            []int32
	MaterialID         []int32
}

type WaveRushScheduleMetaData_LimitGroup struct {
	// Fields
	BuffLimit uint8
	SiteID    int32
}

type WaveRushStageGroupMetaData struct {
	// Fields
	StageGroupConfigID int32
	StageID            int32

	// Properties
	SiteID             int32
	Addr1              uint32 `json:"-"`
	Area               int32
	Type               uint8
	SiteType           uint8
	PreSiteID          int32
	BeginTimeOffSet    int32
	Addr2              uint32 `json:"-"`
	SiteIndex          int32
	ShowDialog         uint8
	Addr3              uint32 `json:"-"`
	IsNeedRank         bool
	IsNeedRankEntrance bool
	IsShowSiteReward   bool

	// Objects
	SiteName         Hash
	UnlockBuffIDList []int32
	DialogDesc       []Hash
}

type WaveRushStageMetaData struct {
	// Fields
	StageID int32

	// Properties
	Type             uint8
	MaxScore         int32
	IsInRewardLine   bool
	StageTime        int32
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	IsOnlySimple     bool
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	HardlevelGroupID int32
	Hardlevel        int32

	// Objects
	WeatherIDList     []int32
	SimpleStageDetail StrWithPrefix16
	StageScene        StrWithPrefix16
	WaveIDList        []int32
}

type WaveRushStageScoreDropMetaData struct {
	// Fields
	StageID      int32
	NeedMinScore int32

	// Properties
	DropMaterialNum int32
}

type WaveRushStageWaveMetaData struct {
	// Fields
	WaveID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Addr4            uint32 `json:"-"`
	Addr5            uint32 `json:"-"`
	WaveDistribution uint8

	// Objects
	UIDList     StrWithPrefix16
	WaveBossUID []int32
	FeverAdd    StrWithPrefix16
	TimeAdd     StrWithPrefix16
	ScoreAdd    StrWithPrefix16
}

type WaveRushUIConfigMetaData struct {
	// Fields
	UIConfigID int32

	// Properties
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	Addr4             uint32 `json:"-"`
	MissionMaxShowNum int32

	// Objects
	UIManager   StrWithPrefix16
	AreaMapList []WaveRushUIConfigMetaData_AreaTuple
	WebLink     StrWithPrefix16
	Tutorial    []WaveRushUIConfigMetaData_TutorialTuple
}

type WaveRushUIConfigMetaData_AreaTuple struct {
	// Fields
	ID    int32
	Addr1 uint32 `json:"-"`

	// Objects
	PrefabPath StrWithPrefix16
}

type WaveRushUIConfigMetaData_TutorialTuple struct {
	// Fields
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	Content StrWithPrefix16
	Sprite  StrWithPrefix16
	Title   StrWithPrefix16
}

type WeaponBaseTypeMetaData struct {
	// Fields
	BaseType int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	// Objects
	CapsulesList []int32
	TypeIcon     StrWithPrefix16
	TypeName     Hash
	TypeNameEn   Hash
}

type WeaponCustomDisplayMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	AvatarIDList []int32
	BodyMod      StrWithPrefix16
}

type WeaponLowResOverrideMetaData struct {
	// Fields
	AvatarID uint16
	WeaponID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	BodyMod StrWithPrefix16
}

type WeaponMainIDMetaData struct {
	// Fields
	WeaponMainID int32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	ReforgeTargetIDList []int32
}

type WeaponShowOverrideMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	ForbidAutoRotate bool

	// Objects
	BodyMod  StrWithPrefix16
	IconPath StrWithPrefix16
}

type WeaponTagMetaData struct {
	// Fields
	WeaponID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	// Objects
	SpecificAvatarBonus        []WeaponTagMetaData_RatioItem
	SpecificAvatarBonusSupport []WeaponTagMetaData_RatioItem
}

type WeaponTagMetaData_RatioItem struct {
	// Fields
	Pred  int32
	Ratio float32
}

type WebLinkAvatarMetaData struct {
	// Fields
	AvatarID         int32
	IsAvatarArtifact bool

	// Properties
	Addr1     uint32 `json:"-"`
	IsWebView bool

	// Objects
	WebLinks StrWithPrefix16
}

type WebLinkElfMetaData struct {
	// Fields
	ElfID int32

	// Properties
	Addr1     uint32 `json:"-"`
	IsWebView bool

	// Objects
	WebLinks StrWithPrefix16
}

type WebLinkEquipMetaData struct {
	// Fields
	EquipId int32

	// Properties
	Addr1     uint32 `json:"-"`
	IsWebView bool

	// Objects
	WebLinks StrWithPrefix16
}

type WebLinkExBossMetaData struct {
	// Fields
	BossId int32

	// Properties
	Addr1     uint32 `json:"-"`
	IsWebView bool

	// Objects
	WebLinks StrWithPrefix16
}

type WeekDayActivityMetaData struct {
	// Fields
	WeekDayActivityID int32

	// Properties
	SeriesID                    int32
	ActivityType                uint8
	ModeSwitch                  uint8
	Addr1                       uint32 `json:"-"`
	EnableCornerName            bool
	Addr2                       uint32 `json:"-"`
	Addr3                       uint32 `json:"-"`
	Addr4                       uint32 `json:"-"`
	Addr5                       uint32 `json:"-"`
	Addr6                       uint32 `json:"-"`
	Addr7                       uint32 `json:"-"`
	Addr8                       uint32 `json:"-"`
	Addr9                       uint32 `json:"-"`
	ShowEnterImg                bool
	ShowRedHint                 bool
	Addr10                      uint32 `json:"-"`
	Addr11                      uint32 `json:"-"`
	Addr12                      uint32 `json:"-"`
	Addr13                      uint32 `json:"-"`
	MaxEnterTimes               uint16
	LevelLockID                 int32
	MinPlayerLevel              uint8
	MaxPlayerLevel              int8
	MinDisplayLevel             uint8
	IsSegmentActivity           bool
	DisplayTodayEndTime         bool
	DisplayTodayBeginTime       bool
	DisplayLeftTime             bool
	ShowActivityShopEntry       uint16
	ActivityCoinEnable          bool
	ActivityCoinCost            int32
	ActivityCoinID              uint16
	ShowInMP                    bool
	MemberMaxTimesForCurrentAct uint8
	MPCostTimes                 uint8
	DisplayVRID                 uint8
	Addr14                      uint32 `json:"-"`
	IsAutoEnterMode             bool
	OpenScoinCost               uint16
	SeasonID                    uint8
	AvatarSampleSwitch          bool
	RemoveTimeoutDisplay        bool
	TeamListID                  uint16

	// Objects
	RefIDSwitch         []int32
	TextMapCornerName   Hash
	Title               Hash
	SubTitle            Hash
	Desc                Hash
	DescLock            Hash
	SmallImgPath        StrWithPrefix16
	BgImgPath           StrWithPrefix16
	EnterImgPath        StrWithPrefix16
	LevelPanelPath      StrWithPrefix16
	LevelIDList         []int32
	DisplayDropList     []int32
	StageEventList      []int32
	ActivityLinkContent StrWithPrefix16
}

type WeekDayActivityScheduleMetaData struct {
	// Fields
	Id int32

	// Properties
	ActivityId int32
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`

	// Objects
	BeginDate       StrWithPrefix16
	VersionTagTitle Hash
}

type WeeklyReportMedalMetaData struct {
	// Fields
	ID uint32

	// Properties
	Addr1 uint32 `json:"-"`

	// Objects
	MedalList []uint32
}

type WeeklyReportMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Priority int32
	Levelmax int32
	Levelmin int32
	Addr2    uint32 `json:"-"`

	// Objects
	TabName   StrWithPrefix16
	BeginTime StrWithPrefix16
}

type WeeklyReportSubPageDataMetaData struct {
	// Fields
	ID       int32
	FatherID int32

	// Properties
	Addr1    uint32 `json:"-"`
	Addr2    uint32 `json:"-"`
	Levelmax int32
	Levelmin int32

	// Objects
	WebLink         StrWithPrefix16
	Ishave_extraweb []int32
}

type WelfareMissionMetaData struct {
	// Fields
	MissionID int32

	// Properties
	Addr1       uint32 `json:"-"`
	Addr2       uint32 `json:"-"`
	PlotID      int32
	DisplayType int32

	// Objects
	Comment     Hash
	MissionIcon StrWithPrefix16
}

type WikiActivityCatalogMetaData struct {
	// Fields
	CatalogID uint8

	// Properties
	ActivityID uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`

	// Objects
	Title       StrWithPrefix16
	Description StrWithPrefix16
	Author      StrWithPrefix16
	AuthorPic   StrWithPrefix16
}

type WikiActivityChapterMetaData struct {
	// Fields
	ChapterID uint16

	// Properties
	ActivityID uint8
	Addr1      uint32 `json:"-"`
	Addr2      uint32 `json:"-"`
	Addr3      uint32 `json:"-"`
	Addr4      uint32 `json:"-"`
	Addr5      uint32 `json:"-"`

	// Objects
	Title       StrWithPrefix16
	Description StrWithPrefix16
	Cover       StrWithPrefix16
	Author      StrWithPrefix16
	AuthorQIcon StrWithPrefix16
}

type WikiActivityEventMetaData struct {
	// Fields
	EventID uint16

	// Properties
	ChapterID uint16
	Addr1     uint32 `json:"-"`
	Plot      uint16

	// Objects
	Title StrWithPrefix16
}

type WikiActivityMetaData struct {
	// Fields
	ActivityID uint8

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`
	Addr8 uint32 `json:"-"`
	Addr9 uint32 `json:"-"`

	// Objects
	Name        StrWithPrefix16
	Time        StrWithPrefix16
	Year        StrWithPrefix16
	Season      StrWithPrefix16
	SeasonPic   StrWithPrefix16
	SeasonColor StrWithPrefix16
	Description StrWithPrefix16
	Cover       StrWithPrefix16
	Thumbnails  StrWithPrefix16
}

type WikiCollectionRankMetaData struct {
	// Fields
	RankID int32

	// Properties
	RankLevel int32
	Addr1     uint32 `json:"-"`
	Addr2     uint32 `json:"-"`
	Addr3     uint32 `json:"-"`
	RankScore int32
	Reward    int32

	// Objects
	RankName   Hash
	RankNameEn Hash
	RankIcon   StrWithPrefix16
}

type WikiTypeMetaData struct {
	// Fields
	WikiType int32

	// Properties
	Addr1              uint32 `json:"-"`
	Addr2              uint32 `json:"-"`
	Addr3              uint32 `json:"-"`
	Addr4              uint32 `json:"-"`
	Addr5              uint32 `json:"-"`
	Addr6              uint32 `json:"-"`
	Addr7              uint32 `json:"-"`
	Addr8              uint32 `json:"-"`
	Addr9              uint32 `json:"-"`
	Lock               int32
	UnLockLevel        int32
	CrownCoef          int32
	WikiCollectionCoef int32

	// Objects
	WikiTypeName          Hash
	WikiTypeNameEn        Hash
	WikiTypeColor         StrWithPrefix16
	WikiTypeIconTypeColor StrWithPrefix16
	WikiTypeLineColor     StrWithPrefix16
	WikiTypeIcon          StrWithPrefix16
	WikiTypeNameIcon      StrWithPrefix16
	WikiTypeShareIcon     StrWithPrefix16
	WikiTypeImg           StrWithPrefix16
}

type WorldBossGroupMetaData struct {
	// Fields
	GroupID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	RankingTitle    Hash
	RankingSubTitle Hash
	RankingIconPath StrWithPrefix16
}

type WorldBossRankRewardMetaData struct {
	// Fields
	ID int32

	// Properties
	WorldBossRank int32
	Reward        int32
}

type WorldBossScoreRewardMetaData struct {
	// Fields
	Id int32

	// Properties
	WorldBossScore int32
	Reward         int32
}

type WorldBossSkillMetaData struct {
	// Fields
	ID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	IconPath  StrWithPrefix16
	SkillName Hash
	SkillDesc Hash
}

type WorldMapActivityInfoMetaData struct {
	// Fields
	InfoID int32

	// Properties
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	// Objects
	IconPath    StrWithPrefix16
	Desc        Hash
	LinkContent StrWithPrefix16
}

type WorldMapCoverInfoMetaData struct {
	// Fields
	InfoID int32

	// Properties
	EntryID int32
	Addr1   uint32 `json:"-"`
	Addr2   uint32 `json:"-"`
	Addr3   uint32 `json:"-"`

	// Objects
	ContentIDList []int32
	Title         Hash
	Desc          Hash
}

type WorldMapEntryMetaData struct {
	// Fields
	ID int32

	// Properties
	EntryID                int32
	Addr1                  uint32 `json:"-"`
	Addr2                  uint32 `json:"-"`
	Addr3                  uint32 `json:"-"`
	Addr4                  uint32 `json:"-"`
	Addr5                  uint32 `json:"-"`
	EntryCG                int32
	UnlockLevel            int32
	UnlockHintLevel        int32
	ForceShowContent       int32
	InfoID                 int32
	ShopType               int32
	Addr6                  uint32 `json:"-"`
	Addr7                  uint32 `json:"-"`
	EntryRewardID          int32
	VersionTag             bool
	VersionTagMinLv        int32
	VersionTagMaxLv        int32
	IsShowActivitySchedule bool

	// Objects
	FeatureType          []int32
	ContentIDList        []int32
	EntryTitle           Hash
	EntryImagePath       StrWithPrefix16
	EntryTitleImgPath    StrWithPrefix16
	EntryDesc            Hash
	EntryDetailImagePath StrWithPrefix16
}

type WorldMapScheduleMetaData struct {
	// Fields
	ID int32

	// Properties
	WorldMapID int32
	Addr1      uint32 `json:"-"`

	// Objects
	AdvanceNote Hash
}

type AbilitySpecialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []AbilitySpecialMetaData
}

type AccountBuffEffectReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AccountBuffEffect
}

type AccumulatorUIConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []AccumulatorUIConfigMetaData
}

type AchievementTagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AchievementTagMetaData
}

type ActChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ActChallengeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ActChallengeMetaData
}

type ActChallengeMetaDataReader_StructKey struct {
	// Fields
	ActId      int32
	Difficulty int32
}

type ActiveCollectionItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActiveCollectionItemMetaData
}

type ActiveCollectionScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActiveCollectionScheduleMetaData
}

type ActivityBbqMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityBbqMetaData
}

type ActivityBingoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityBingoMetaData
}

type ActivityBossRushMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityBossRushMetaData
}

type ActivityCardDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityCardDataMetaData
}

type ActivityChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityChallengeMetaData
}

type ActivityChargeLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ActivityChargeLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ActivityChargeLevelMetaData
}

type ActivityChargeLevelMetaDataReader_StructKey struct {
	// Fields
	ChargeID    uint8
	AcitivityID int32
}

type ActivityDropLimitDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ActivityDropLimitData
}

type ActivityDropLimitScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ActivityDropLimitScheduleMetaData
}

type ActivityLoginBonusMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ActivityLoginBonusMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ActivityLoginBonusMetaData
}

type ActivityLoginBonusMetaDataReader_StructKey struct {
	// Fields
	Days int32
	ID   int32
}

type ActivityLoginMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityLoginMetaData
}

type ActivityMosaicMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityMosaicMetaData
}

type ActivityPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ActivityPanelMetaData
}

type ActivityPanelScoreDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ActivityPanelScoreDataReader_StructKey
	ItemOffsets []int32

	Data []ActivityPanelScoreData
}

type ActivityPanelScoreDataReader_StructKey struct {
	// Fields
	MissionID int32
	PanelID   int32
}

type ActivityPanelScoreRewardDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ActivityPanelScoreRewardDataReader_StructKey
	ItemOffsets []int32

	Data []ActivityPanelScoreRewardData
}

type ActivityPanelScoreRewardDataReader_StructKey struct {
	// Fields
	PanelID  int32
	Progross int32
}

type ActivityPictureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityPictureMetaData
}

type ActivityQuestionnaireMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityQuestionnaireMetaData
}

type ActivityRandomEffectChargeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ActivityRandomEffectChargeMetaData
}

type ActivityRandomEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityRandomEffectMetaData
}

type ActivityRewardShowMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ActivityRewardShowMetaData
}

type ActivityScheduleDisplayDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityScheduleDisplayData
}

type ActivitySchedulePageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []ActivitySchedulePageMetaData
}

type ActivityStageRankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityStageRankMetaData
}

type ActivityTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActivityTowerMetaData
}

type ActivityTypeManagementMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ActivityTypeManagementMetaData
}

type ActMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ActMetaData
}

type ActRitaRankingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ActRitaRankingMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ActRitaRankingMetaData
}

type ActRitaRankingMetaDataReader_StructKey struct {
	// Fields
	LetterRankID int32
	RankLevel    int32
}

type AddParamGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AddParamGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AddParamGroupMetaData
}

type AddParamGroupMetaDataReader_StructKey struct {
	// Fields
	GroupID uint8
	SkillID int32
}

type AdventureAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureAvatarMetaData
}

type AdventureAvatarMissionListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureAvatarMissionListMetaData
}

type AdventureAvatarSkillLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureAvatarSkillLevelMetaData
}

type AdventureAvatarSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureAvatarSkillMetaData
}

type AdventureAvatarUnlockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureAvatarUnlockMetaData
}

type AdventureDecorateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []AdventureDecorateMetaData
}

type AdventureElfMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureElfMetaData
}

type AdventureElfUnlockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureElfUnlockMetaData
}

type AdventureLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureLevelDataMetaData
}

type AdventureQuestConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureQuestConfigMetaData
}

type AdventureQuestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureQuestMetaData
}

type AdventureQuestPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureQuestPoolMetaData
}

type AdventureWelfareMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AdventureWelfareMetaData
}

type AffixEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixEffectMetaData
}

type AffixRecycleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixRecycleMetaData
}

type AffixTitleEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixTitleEffectMetaData
}

type AffixTitleEffectTemplateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixTitleEffectTemplateMetaData
}

type AffixTitleLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixTitleLevelMetaData
}

type AffixTreeNodeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixTreeNodeMetaData
}

type AffixTypeConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AffixTypeConfigMetaData
}

type AiCyberActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []AiCyberActivityScheduleMetaData
}

type AiCyberAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberAreaMetaData
}

type AiCyberDailyStageDropLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AiCyberDailyStageDropLimitMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AiCyberDailyStageDropLimitMetaData
}

type AiCyberDailyStageDropLimitMetaDataReader_StructKey struct {
	// Fields
	Priority           uint8
	ActivityScheduleID uint32
}

type AiCyberDailyStageScoreDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AiCyberDailyStageScoreDropMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AiCyberDailyStageScoreDropMetaData
}

type AiCyberDailyStageScoreDropMetaDataReader_StructKey struct {
	// Fields
	MinScore     int32
	StageGroupID uint32
}

type AiCyberGalTouchMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberGalTouchMetaData
}

type AiCyberHyperionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberHyperionMetaData
}

type AiCyberMainUIRepairMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberMainUIRepairMetaData
}

type AiCyberProgressMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberProgressMetaData
}

type AiCyberPuzzleEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberPuzzleEventMetaData
}

type AiCyberQavatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberQavatarMetaData
}

type AiCyberQavatarSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberQavatarSkillMetaData
}

type AiCyberStageGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberStageGroupMetaData
}

type AiCyberStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AiCyberStageMetaData
}

type AiCyberStoryDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AiCyberStoryDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AiCyberStoryDataMetaData
}

type AiCyberStoryDataMetaDataReader_StructKey struct {
	// Fields
	Index        uint8
	StageGroupID int32
}

type AnniversaryFifthBoxDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []AnniversaryFifthBoxDataMetaData
}

type AnniversaryFifthDetailItemsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []AnniversaryFifthDetailItemsMetaData
}

type AnniversaryFifthShowItemsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []AnniversaryFifthShowItemsMetaData
}

type AnniversaryIconintegrateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []AnniversaryIconintegrateMetaData
}

type AppointmentDownloadScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []AppointmentDownloadScheduleMetaData
}

type ArmadaBossActivityDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaBossActivityDataMetaData
}

type ArmadaBossActivityScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ArmadaBossActivityScoreMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ArmadaBossActivityScoreMetaData
}

type ArmadaBossActivityScoreMetaDataReader_StructKey struct {
	// Fields
	ActivityID    int32
	ActivityScore int32
	ScoreType     int32
}

type ArmadaBossDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaBossData
}

type ArmadaContactDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaContactData
}

type ArmadaExchequerDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaExchequerData
}

type ArmadaGridDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaGridData
}

type ArmadaHangarDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaHangarData
}

type ArmadaLinerRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaLinerRewardMetaData
}

type ArmadaMainDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaMainData
}

type ArmadaPersonalizedLabelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaPersonalizedLabelMetaData
}

type ArmadaRecruitTipsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ArmadaRecruitTipsMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ArmadaRecruitTipsMetaData
}

type ArmadaRecruitTipsMetaDataReader_StructKey struct {
	// Fields
	ArmadaLevel int32
	TYPE        int32
}

type ArmadaReunionLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ArmadaReunionLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ArmadaReunionLevelMetaData
}

type ArmadaReunionLevelMetaDataReader_StructKey struct {
	// Fields
	Level      int32
	ScheduleID int32
}

type ArmadaReunionMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ArmadaReunionMissionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ArmadaReunionMissionMetaData
}

type ArmadaReunionMissionMetaDataReader_StructKey struct {
	// Fields
	ArmadaReunionMissionID int32
	ScheduleID             int32
}

type ArmadaReunionPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaReunionPointMetaData
}

type ArmadaReunionScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaReunionScheduleMetaData
}

type ArmadaStageBossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaStageBossMetaData
}

type ArmadaStageHardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaStageHardMetaData
}

type ArmadaWelcomeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaWelcomeMetaData
}

type ArmadaWorkshopDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArmadaWorkshopData
}

type ArtifactSkillNameOverrideMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ArtifactSkillNameOverrideMetaData
}

type ArtifactTryEnterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ArtifactTryEnterMetaData
}

type AudioPackageDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AudioPackageData
}

type AvatarArtifactLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarArtifactLevelMetaData
}

type AvatarArtifactMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarArtifactMetaData
}

type AvatarAttackPunishMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarAttackPunishMetaData
}

type AvatarAttackSpeedParamMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarAttackSpeedParamMetaData
}

type AvatarBriefMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarBriefMetaData
}

type AvatarCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarCardMetaData
}

type AvatarDefencePunishMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarDefencePunishMetaData
}

type AvatarEffectInfoDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarEffectInfoData
}

type AvatarExMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarExMetaData
}

type AvatarFilterDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarFilterMetaData
}

type AvatarForgeRecommendMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarForgeRecommendMetaData
}

type AvatarForgeWhiteListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarForgeWhiteListMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarForgeWhiteListMetaData
}

type AvatarForgeWhiteListMetaDataReader_StructKey struct {
	// Fields
	AvatarID     int32
	LevelgroupID int32
}

type AvatarFragmentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarFragmentMetaData
}

type AvatarGachaDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarGachaDisplayMetaData
}

type AvatarGoodFeelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarGoodFeelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarGoodFeelMetaData
}

type AvatarGoodFeelMetaDataReader_StructKey struct {
	// Fields
	AvatarId      int32
	GoodFeelLevel int32
}

type AvatarLevelGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarLevelGroupMetaData
}

type AvatarLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarLevelMetaData
}

type AvatarMemoirMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarMemoirMetaData
}

type AvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarMetaData
}

type AvatarMissionActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarMissionActivityMetaData
}

type AvatarMissionActivityRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarMissionActivityRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarMissionActivityRewardMetaData
}

type AvatarMissionActivityRewardMetaDataReader_StructKey struct {
	// Fields
	AccumulatedDays int32
	ID              int32
}

type AvatarModelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarModelMetaData
}

type AvatarNewbieMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarNewbieMetaData
}

type AvatarPracticeMainMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarPracticeMainMetaData
}

type AvatarPracticeStepMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarPracticeStepMetaData
}

type AvatarRelayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarRelayMetaData
}

type AvatarRewardGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarRewardGroupMetaData
}

type AvatarRoleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarRoleMetaData
}

type AvatarSampleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarSampleMetaData
}

type AvatarSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarSkillMetaData
}

type AvatarStarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarStarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarStarMetaData
}

type AvatarStarMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	Star     int32
	SubStar  int32
}

type AvatarStarTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarStarTypeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarStarTypeMetaData
}

type AvatarStarTypeMetaDataReader_StructKey struct {
	// Fields
	AvatarStarUpType uint8
	AvatarType       uint8
	Star             uint8
	SubStar          uint8
}

type AvatarStarUpSubSkillDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarStarUpSubSkillDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarStarUpSubSkillData
}

type AvatarStarUpSubSkillDataReader_StructKey struct {
	// Fields
	Star    uint8
	SubStar uint8
	SkillID int32
}

type AvatarSubSkillLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []AvatarSubSkillLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []AvatarSubSkillLevelMetaData
}

type AvatarSubSkillLevelMetaDataReader_StructKey struct {
	// Fields
	ItemType   int32
	SkillLevel int32
}

type AvatarSubSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []AvatarSubSkillMetaData
}

type AvatarTagUnLockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarTagUnLockMetaData
}

type AvatarTrialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []AvatarTrialMetaData
}

type AvatarTryEnterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []AvatarTryEnterMetaData
}

type AvatarTutorialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarTutorialMetaData
}

type AvatarTutorialSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []AvatarTutorialSiteMetaData
}

type AvatarTwinsDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarTwinsMetaData
}

type AvatarVideoPopupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []AvatarVideoPopupMetaData
}

type BattleDisplayDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BattleDisplayData
}

type BattlePassBannerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []BattlePassBannerMetaData
}

type BattlePassChallengeDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BattlePassChallengeDataMetaData
}

type BattlePassScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []BattlePassScheduleMetaData
}

type BattlePassSeasonLevelConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []BattlePassSeasonLevelConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []BattlePassSeasonLevelConfigMetaData
}

type BattlePassSeasonLevelConfigMetaDataReader_StructKey struct {
	// Fields
	SeasonLevel         uint8
	SeasonLevelComfigID uint8
}

type BattlePassTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BattlePassTypeMetaData
}

type BattleTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BattleTypeMetaData
}

type BbqBonusSetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BbqBonusSetMetaData
}

type BbqLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BbqLevelMetaData
}

type BeastStageDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []BeastStageDisplayMetaDataReader_StructKey
	ItemOffsets []int32

	Data []BeastStageDisplayMetaData
}

type BeastStageDisplayMetaDataReader_StructKey struct {
	// Fields
	BeastStageID int32
	Level        int32
}

type BeastStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BeastStageMetaData
}

type BeastTreasureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BeastTreasureMetaData
}

type BossChallengeActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BossChallengeActivityMetaData
}

type BossChallengeRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []BossChallengeRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []BossChallengeRewardMetaData
}

type BossChallengeRewardMetaDataReader_StructKey struct {
	// Fields
	Index          uint16
	RewardConfigID int32
}

type BossChallengeStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BossChallengeStageMetaData
}

type BossPointScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []BossPointScoreRewardMetaData
}

type BossRushBuffConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []BossRushBuffConfigMetaData
}

type BossRushBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BossRushBuffMetaData
}

type BossRushBuffPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BossRushBuffPoolMetaData
}

type BossRushStageScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []BossRushStageScheduleMetaData
}

type BoxGachaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BoxGachaMetaData
}

type BridgeLinkMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []BridgeLinkMetaData
}

type BubbleBridgeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BubbleBridgeMetaData
}

type BuffAssistanceActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BuffAssistanceActivityMetaData
}

type BuffAssistanceBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BuffAssistanceBuffMetaData
}

type BuffAssistanceLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []BuffAssistanceLevelMetaData
}

type BuffAssistanceNPCRobotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BuffAssistanceNPCRobotMetaData
}

type BuffAssistanceScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BuffAssistanceScheduleMetaData
}

type BuffAssistanceStageGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BuffAssistanceStageGroupMetaData
}

type BuffAssistanceWordMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BuffAssistanceWordMetaData
}

type BulletinEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BulletinEffectMetaData
}

type BurdenAlleviationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BurdenAlleviationMetaData
}

type BurialPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []BurialPointMetaData
}

type CgCategoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CgCategoryMetaData
}

type CgGroupDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CgGroupDataMetaData
}

type CgMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CgMetaData
}

type CGPackageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CGPackageDataMetaData
}

type ChallengeMissionLinkReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeMissionLink
}

type ChallengeMissionPanelReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeMissionPanel
}

type ChallengeMissionStepDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeMissionStepData
}

type ChallengeWarAchieveMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeWarAchieveMissionMetaData
}

type ChallengeWarAreaInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeWarAreaInfoMetaData
}

type ChallengeWarBuffDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeWarBuffDataMetaData
}

type ChallengeWarSeasonMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeWarSeasonMetaData
}

type ChallengeWarStageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChallengeWarStageDataMetaData
}

type ChallengeWarStageFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChallengeWarStageFloorMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChallengeWarStageFloorMetaData
}

type ChallengeWarStageFloorMetaDataReader_StructKey struct {
	// Fields
	Floor   int16
	StageID int32
}

type ChallengeWarTagEffectPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChallengeWarTagEffectPoolMetaData
}

type ChanllengeWarAvatarSubStarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChanllengeWarAvatarSubStarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChanllengeWarAvatarSubStarMetaData
}

type ChanllengeWarAvatarSubStarMetaDataReader_StructKey struct {
	// Fields
	AvatarUnlockStar uint8
	Star             uint8
	SubStar          uint8
}

type ChapterActivityRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterActivityRewardMetaData
}

type ChapterActivityRpgAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterActivityRpgAvatarMetaData
}

type ChapterActivitySectionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterActivitySectionMetaData
}

type ChapterActivityStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterActivityStageMetaData
}

type ChapterActivityZoneBuffDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterActivityZoneBuffDataMetaData
}

type ChapterActivityZoneDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterActivityZoneDataMetaData
}

type ChapterBuffDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterBuffDisplayMetaData
}

type ChapterChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterChallengeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterChallengeMetaData
}

type ChapterChallengeMetaDataReader_StructKey struct {
	// Fields
	ChallengeNum int32
	ChapterID    int32
}

type ChapterCultivateAttrMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterCultivateAttrMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterCultivateAttrMetaData
}

type ChapterCultivateAttrMetaDataReader_StructKey struct {
	// Fields
	CultivateID    uint32
	CultivateLevel uint32
}

type ChapterCultivateBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterCultivateBuffMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterCultivateBuffMetaData
}

type ChapterCultivateBuffMetaDataReader_StructKey struct {
	// Fields
	ID    uint32
	Level uint32
}

type ChapterCultivateBuffShopMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterCultivateBuffShopMetaData
}

type ChapterCultivateConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterCultivateConfigMetaData
}

type ChapterCultivateJumpMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterCultivateJumpMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterCultivateJumpMetaData
}

type ChapterCultivateJumpMetaDataReader_StructKey struct {
	// Fields
	ChapterID      int32
	CultivateLevel uint32
}

type ChapterCultivateLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterCultivateLevelDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterCultivateLevelDataMetaData
}

type ChapterCultivateLevelDataMetaDataReader_StructKey struct {
	// Fields
	CultivateUseID uint32
	Level          uint32
}

type ChapterCultivateMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterCultivateMonsterMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterCultivateMonsterMetaData
}

type ChapterCultivateMonsterMetaDataReader_StructKey struct {
	// Fields
	SkillID    uint32
	SkillLevel uint32
}

type ChapterCultivateSiteEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterCultivateSiteEffectMetaData
}

type ChapterCultivateSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterCultivateSkillMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterCultivateSkillMetaData
}

type ChapterCultivateSkillMetaDataReader_StructKey struct {
	// Fields
	CultivateID    uint32
	CultivateLevel uint32
}

type ChapterCultivateUseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterCultivateUseMetaData
}

type ChapterGroupConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ChapterGroupConfigMetaData
}

type ChapterGroupFuncButtonMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterGroupFuncButtonMetaData
}

type ChapterGroupSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ChapterGroupSiteMetaData
}

type ChapterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterMetaData
}

type ChapterOW31DialogueMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOW31DialogueMetaData
}

type ChapterOWAchievementMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWAchievementMetaData
}

type ChapterOWActivityPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWActivityPanelMetaData
}

type ChapterOWAntiGravityGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWAntiGravityGroupMetaData
}

type ChapterOWAntiGravityLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWAntiGravityLevelMetaData
}

type ChapterOWBagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWBagMetaData
}

type ChapterOWBuildingLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWBuildingLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWBuildingLevelMetaData
}

type ChapterOWBuildingLevelMetaDataReader_StructKey struct {
	// Fields
	Level  uint32
	MainID uint32
}

type ChapterOWBuildingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWBuildingMetaData
}

type ChapterOWBuildingStateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWBuildingStateMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWBuildingStateMetaData
}

type ChapterOWBuildingStateMetaDataReader_StructKey struct {
	// Fields
	StateType uint8
	MainID    int32
}

type ChapterOWChallengeBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWChallengeBuffMetaData
}

type ChapterOWChallengeDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWChallengeDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWChallengeDataMetaData
}

type ChapterOWChallengeDataMetaDataReader_StructKey struct {
	// Fields
	ChallengeIndex uint32
	ConfigID       uint32
}

type ChapterOWChallengeGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWChallengeGroupMetaData
}

type ChapterOWChallengeSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWChallengeSiteMetaData
}

type ChapterOWCollectionTabMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWCollectionTabMetaData
}

type ChapterOWConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWConfigMetaData
}

type ChapterOWDailyStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWDailyStageMetaData
}

type ChapterOWDigsiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWDigsiteMetaData
}

type ChapterOWDigsiteProgramMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWDigsiteProgramMetaData
}

type ChapterOWEndlessChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWEndlessChallengeMetaData
}

type ChapterOWEntryPlotLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWEntryPlotLineMetaData
}

type ChapterOWEquipmentBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWEquipmentBuffMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWEquipmentBuffMetaData
}

type ChapterOWEquipmentBuffMetaDataReader_StructKey struct {
	// Fields
	ID    int32
	Level int32
}

type ChapterOWEquipmentPartMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWEquipmentPartMetaData
}

type ChapterOWEquipmentSlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWEquipmentSlotMetaData
}

type ChapterOWEventTabConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWEventTabConfigMetaData
}

type ChapterOWFameMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWFameMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWFameMetaData
}

type ChapterOWFameMetaDataReader_StructKey struct {
	// Fields
	FameLevel uint8
	MapID     int32
}

type ChapterOWFortDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWFortDataMetaData
}

type ChapterOWFurnaceFormulaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWFurnaceFormulaMetaData
}

type ChapterOWFurnitureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ChapterOWFurnitureMetaData
}

type ChapterOWHeroCardLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWHeroCardLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWHeroCardLevelMetaData
}

type ChapterOWHeroCardLevelMetaDataReader_StructKey struct {
	// Fields
	Level  uint8
	CardID int32
}

type ChapterOWHeroCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWHeroCardMetaData
}

type ChapterOWHeroCardRarityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWHeroCardRarityMetaData
}

type ChapterOWHeroDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWHeroDisplayMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWHeroDisplayMetaData
}

type ChapterOWHeroDisplayMetaDataReader_StructKey struct {
	// Fields
	HeroStatus uint8
	ID         uint32
}

type ChapterOWHeroLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWHeroLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWHeroLevelMetaData
}

type ChapterOWHeroLevelMetaDataReader_StructKey struct {
	// Fields
	Level  uint8
	HeroID int32
}

type ChapterOWHeroMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWHeroMetaData
}

type ChapterOWHeroSpMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWHeroSpMetaData
}

type ChapterOWInteractActionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWInteractActionMetaData
}

type ChapterOWInteractStateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWInteractStateMetaData
}

type ChapterOWLandMarkMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWLandMarkMetaData
}

type ChapterOWLevelSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWLevelSiteMetaData
}

type ChapterOWMainLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWMainLineMetaData
}

type ChapterOWMemoryContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWMemoryContentMetaData
}

type ChapterOWMemoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWMemoryMetaData
}

type ChapterOWMemorySiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWMemorySiteMetaData
}

type ChapterOWMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWMissionMetaData
}

type ChapterOWMissionTabMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWMissionTabMetaData
}

type ChapterOWMoonConditionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWMoonConditionMetaData
}

type ChapterOWMoonTowerFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWMoonTowerFloorMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWMoonTowerFloorMetaData
}

type ChapterOWMoonTowerFloorMetaDataReader_StructKey struct {
	// Fields
	FloorConfigID uint16
	FloorOrder    uint16
}

type ChapterOWMoonTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWMoonTowerMetaData
}

type ChapterOWMoonTowerScoreDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWMoonTowerScoreDropMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWMoonTowerScoreDropMetaData
}

type ChapterOWMoonTowerScoreDropMetaDataReader_StructKey struct {
	// Fields
	ScoreDropConfigID uint16
	Score             uint32
}

type ChapterOWNPCMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWNPCMetaData
}

type ChapterOWNPCPositionGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWNPCPositionGroupMetaData
}

type ChapterOWNPCPositionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWNPCPositionMetaData
}

type ChapterOWPhaseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWPhaseMetaData
}

type ChapterOWPlotSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWPlotSiteMetaData
}

type ChapterOWQTECircleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWQTECircleMetaData
}

type ChapterOWQTECircleRangeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWQTECircleRangeMetaData
}

type ChapterOWQTELastRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWQTELastRewardMetaData
}

type ChapterOWQTEMapGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWQTEMapGroupMetaData
}

type ChapterOWQTEMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWQTEMapMetaData
}

type ChapterOWQuestDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWQuestDataMetaData
}

type ChapterOWQuestLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWQuestLevelMetaData
}

type ChapterOWRandomStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWRandomStageMetaData
}

type ChapterOWRangePlotGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWRangePlotGroupMetaData
}

type ChapterOWRefiningLevelUPMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWRefiningLevelUPMetaData
}

type ChapterOWRelicsEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterOWRelicsEventMetaData
}

type ChapterOWRelicsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWRelicsMetaData
}

type ChapterOWRelicsSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChapterOWRelicsSiteMetaData
}

type ChapterOWRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWRewardMetaData
}

type ChapterOWRewardMetaDataReader_StructKey struct {
	// Fields
	CountNumber int32
	MapID       int32
}

type ChapterOWShopGoodsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWShopGoodsMetaData
}

type ChapterOWShopLinkMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWShopLinkMetaData
}

type ChapterOWShopMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWShopMetaData
}

type ChapterOWSpecialStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWSpecialStoryMetaData
}

type ChapterOWTalentBuffLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ChapterOWTalentBuffLevelMetaData
}

type ChapterOWTalentDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWTalentDataMetaData
}

type ChapterOWTalentLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWTalentLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWTalentLevelMetaData
}

type ChapterOWTalentLevelMetaDataReader_StructKey struct {
	// Fields
	TalentLevel uint8
	TalentID    int32
}

type ChapterOWTerminalLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChapterOWTerminalLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChapterOWTerminalLevelMetaData
}

type ChapterOWTerminalLevelMetaDataReader_StructKey struct {
	// Fields
	Level uint8
	MapID int32
}

type ChapterOWTerminalTipsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWTerminalTipsMetaData
}

type ChapterOWTipsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterOWTipsMetaData
}

type ChapterRpgAvatarRewardLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterRpgAvatarRewardLineMetaData
}

type ChapterRpgBuffLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterRpgBuffLevelMetaData
}

type ChapterStageCompensationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChapterStageCompensationMetaData
}

type ChapterSwitchAnimConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ChapterSwitchAnimConfigMetaData
}

type ChatGroupIconMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatGroupIconMetaData
}

type ChatLobbyActionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []ChatLobbyActionMetaData
}

type ChatLobbyActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyActivityMetaData
}

type ChatLobbyActivityNoticeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyActivityNoticeMetaData
}

type ChatLobbyActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyActivityScheduleMetaData
}

type ChatLobbyAnnouncementMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyAnnouncementMetaData
}

type ChatLobbyBeastMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyBeastMetaData
}

type ChatLobbyBoxActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyBoxActivityMetaData
}

type ChatLobbyBoxSeriesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyBoxSeriesMetaData
}

type ChatLobbyDishLimitedRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyDishLimitedRewardMetaData
}

type ChatLobbyDishMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChatLobbyDishMetaData
}

type ChatLobbyDishRateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ChatLobbyDishRateMetaData
}

type ChatLobbyFightExpressionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyFightExpressionMetaData
}

type ChatLobbyFishActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyFishActivityMetaData
}

type ChatLobbyFishMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyFishMetaData
}

type ChatLobbyFishShowMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyFishShowMetaData
}

type ChatLobbyItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyItemMetaData
}

type ChatLobbyKillingStreakMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyKillingStreakMetaData
}

type ChatLobbyLanternDifficultyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyLanternDifficultyMetaData
}

type ChatLobbyLayoutMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyLayoutMetaData
}

type ChatLobbyMainPageNoticeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyMainPageNoticeMetaData
}

type ChatLobbyMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyMissionMetaData
}

type ChatLobbyNPCMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyNPCMetaData
}

type ChatLobbyNPCSystemFunctionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ChatLobbyNPCSystemFunctionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ChatLobbyNPCSystemFunctionMetaData
}

type ChatLobbyNPCSystemFunctionMetaDataReader_StructKey struct {
	// Fields
	NPCID       int32
	Type        int32
	UIPostionID int32
}

type ChatLobbyObjectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyObjectMetaData
}

type ChatLobbyPrayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyPrayMetaData
}

type ChatLobbyQuestionActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChatLobbyQuestionActivityMetaData
}

type ChatLobbyQuestionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ChatLobbyQuestionMetaData
}

type ChatLobbyRoomMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyRoomMetaData
}

type ChatLobbySceneItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbySceneItemMetaData
}

type ChatLobbySkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbySkillMetaData
}

type ChatLobbyStanceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyStanceMetaData
}

type ChatLobbyTreasureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyTreasureMetaData
}

type ChatLobbyUsableItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyUsableItemMetaData
}

type ChatLobbyVoiceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []ChatLobbyVoiceMetaData
}

type ChatLobbyWayPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatLobbyWayPointMetaData
}

type ChatMessageContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatMessageContentMetaData
}

type ChatMessageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatMessageDataMetaData
}

type ChatReportMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ChatReportMetaData
}

type CityEventPhotoDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CityEventPhotoMetaData
}

type ClickDialogBGCloseBlackListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []ClickDialogBGCloseBlackListMetaData
}

type ClientLogDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ClientLogDataMetaData
}

type CollaborationScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollaborationScheduleMetaData
}

type CollaborationStigmataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollaborationStigmataMetaData
}

type CollaborationWeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollaborationWeaponMetaData
}

type CollectionDialogueMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionDialogueMetaData
}

type CollectionEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionEventMetaData
}

type CollectionFileMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionFileMetaData
}

type CollectionItemDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionItemDataMetaData
}

type CollectionMonsterDetailInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionMonsterDetailInfoMetaData
}

type CollectionMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionMonsterMetaData
}

type CollectionPictureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionPictureMetaData
}

type CollectionTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CollectionTypeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CollectionTypeMetaData
}

type CollectionTypeMetaDataReader_StructKey struct {
	// Fields
	CollectionType uint8
	SysType        uint8
}

type CollectionVisualNovelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CollectionVisualNovelDataMetaData
}

type CompensationOfDormitoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CompensationOfDormitoryMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CompensationOfDormitoryMetaData
}

type CompensationOfDormitoryMetaDataReader_StructKey struct {
	// Fields
	FacilityType int32
	Level        int32
}

type CompensationOfElfBreakMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CompensationOfElfBreakMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CompensationOfElfBreakMetaData
}

type CompensationOfElfBreakMetaDataReader_StructKey struct {
	// Fields
	BreakStep int32
	ElfID     int32
}

type CompensationOfElfSlotUnlockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CompensationOfElfSlotUnlockMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CompensationOfElfSlotUnlockMetaData
}

type CompensationOfElfSlotUnlockMetaDataReader_StructKey struct {
	// Fields
	ElfID   int32
	SlotNum int32
}

type CompensationOfElfTalentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CompensationOfElfTalentMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CompensationOfElfTalentMetaData
}

type CompensationOfElfTalentMetaDataReader_StructKey struct {
	// Fields
	TalentID    int32
	TalentLevel int32
}

type CompensationOfExtendGradeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CompensationOfExtendGradeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CompensationOfExtendGradeMetaData
}

type CompensationOfExtendGradeMetaDataReader_StructKey struct {
	// Fields
	CabinType   int32
	ExtendGrade int32
}

type CompensationOfIslandMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CompensationOfIslandMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CompensationOfIslandMetaData
}

type CompensationOfIslandMetaDataReader_StructKey struct {
	// Fields
	CabinType int32
	Level     int32
}

type ConstValueMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []Hash
	ItemOffsets []int32

	Data []ConstValueMetaData
}

type CoupleTowerRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CoupleTowerRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CoupleTowerRewardMetaData
}

type CoupleTowerRewardMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	MinFloor   int32
}

type CoupleTowerScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []CoupleTowerScoreMetaDataReader_StructKey
	ItemOffsets []int32

	Data []CoupleTowerScoreMetaData
}

type CoupleTowerScoreMetaDataReader_StructKey struct {
	// Fields
	Floor   int32
	StageID int32
}

type CoupleTowerStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CoupleTowerStageMetaData
}

type CouponDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CouponMetaData
}

type CreditRankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CreditRankMetaData
}

type CreditRegeditMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CreditRegeditMetaData
}

type CrisisModeActivityConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CrisisModeActivityConfigMetaData
}

type CrisisModeStageConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CrisisModeStageConfigMetaData
}

type CrisisModeStageGroupConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CrisisModeStageGroupConfigMetaData
}

type CurrencyIconMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CurrencyIconMetaData
}

type CurrencyQuickExchangeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CurrencyQuickExchangeMetaData
}

type CustomGachaExFragmentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []CustomGachaExFragmentMetaData
}

type CustomHeadMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CustomHeadMetaData
}

type CustomHeadPageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CustomHeadPageMetaData
}

type CycleMissionManagementMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []CycleMissionManagementMetaData
}

type DailyMissionMaterialIconMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DailyMissionMaterialIconMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DailyMissionMaterialIconMetaData
}

type DailyMissionMaterialIconMetaDataReader_StructKey struct {
	// Fields
	MaterialType int16
	MaterialID   int32
}

type DailyMPStageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DailyMPStageDataMetaData
}

type DailyRechargeRewardGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DailyRechargeRewardGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DailyRechargeRewardGroupMetaData
}

type DailyRechargeRewardGroupMetaDataReader_StructKey struct {
	// Fields
	GroupID  int32
	Progress int32
}

type DailyRecommendMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []DailyRecommendMeta
}

type DanmakuMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []DanmakuMetaData
}

type DanmakuSlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []DanmakuSlotMetaData
}

type DeviceFPSLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DeviceFPSLimitMetaData
}

type DialogImageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DialogImageDataMetaData
}

type DialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DialogMetaData
}

type DiceyDungeonActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonActivityMetaData
}

type DiceyDungeonBubbleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DiceyDungeonBubbleMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DiceyDungeonBubbleMetaData
}

type DiceyDungeonBubbleMetaDataReader_StructKey struct {
	// Fields
	BubbleID uint8
	RoleID   int32
}

type DiceyDungeonBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []DiceyDungeonBuffMetaData
}

type DiceyDungeonDailyScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonDailyScheduleMetaData
}

type DiceyDungeonDailyScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonDailyScoreMetaData
}

type DiceyDungeonDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonDataMetaData
}

type DiceyDungeonGachaPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonGachaPoolMetaData
}

type DiceyDungeonMonsterBehaviorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonMonsterBehaviorMetaData
}

type DiceyDungeonMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonMonsterMetaData
}

type DiceyDungeonOrnamentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DiceyDungeonOrnamentMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DiceyDungeonOrnamentMetaData
}

type DiceyDungeonOrnamentMetaDataReader_StructKey struct {
	// Fields
	Level      int32
	OrnamentID int32
}

type DiceyDungeonPassiveSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonPassiveSkillMetaData
}

type DiceyDungeonRoleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DiceyDungeonRoleMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DiceyDungeonRoleMetaData
}

type DiceyDungeonRoleMetaDataReader_StructKey struct {
	// Fields
	Level  int32
	RoleID int32
}

type DiceyDungeonRoomMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonRoomMetaData
}

type DiceyDungeonSiteDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonSiteDisplayMetaData
}

type DiceyDungeonSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonSkillMetaData
}

type DiceyDungeonTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DiceyDungeonTowerMetaData
}

type DiceyDungeonTutorialDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DiceyDungeonTutorialDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DiceyDungeonTutorialDataMetaData
}

type DiceyDungeonTutorialDataMetaDataReader_StructKey struct {
	// Fields
	TriggerType   uint8
	DungeonRoomID int32
}

type DiceyDungeonWeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DiceyDungeonWeaponMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DiceyDungeonWeaponMetaData
}

type DiceyDungeonWeaponMetaDataReader_StructKey struct {
	// Fields
	Level    int32
	WeaponID int32
}

type DLC2ConditionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLC2ConditionMetaData
}

type DLC2DailyQuestInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLC2DailyQuestInfoMetaData
}

type DLC2EntryPlotLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLC2EntryPlotLineMetaData
}

type DLC2PlotUIConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLC2PlotUIConfigMetaData
}

type DLC2SupportGiftMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLC2SupportGiftMetaData
}

type DLCAvatarCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCAvatarCardMetaData
}

type DLCAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCAvatarMetaData
}

type DLCGachaTabMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCGachaTabMetaData
}

type DLCLevelDiffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []DLCLevelDiffMetaData
}

type DLCLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCLevelMetaData
}

type DLCNPCMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCNPCMetaData
}

type DLCRecommendLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCRecommendLevelMetaData
}

type DLCReviveCostMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DLCReviveCostMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DLCReviveCostMetaData
}

type DLCReviveCostMetaDataReader_StructKey struct {
	// Fields
	ReviveTimes int32
	ReviveType  int32
}

type DLCStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCStoryMetaData
}

type DLCStoryPreviewMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCStoryPreviewMetaData
}

type DLCSupportDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DLCSupportDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DLCSupportDataMetaData
}

type DLCSupportDataMetaDataReader_StructKey struct {
	// Fields
	SupportType  uint8
	SupportParam int32
}

type DLCSupportLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []DLCSupportLevelDataMetaData
}

type DLCSupportRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DLCSupportRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DLCSupportRewardMetaData
}

type DLCSupportRewardMetaDataReader_StructKey struct {
	// Fields
	NPCID        uint8
	SupportLevel uint8
}

type DLCTalentAffixMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []DLCTalentAffixMetaData
}

type DLCTalentAffixSetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []DLCTalentAffixSetMetaData
}

type DLCTalentLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DLCTalentLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DLCTalentLevelMetaData
}

type DLCTalentLevelMetaDataReader_StructKey struct {
	// Fields
	DLCTalentID int32
	TalentLevel int32
}

type DLCTalentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTalentMetaData
}

type DLCTalentMultiDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []DLCTalentMultiDisplayMetaData
}

type DLCTalentTagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTalentTagMetaData
}

type DLCTowerBonusMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerBonusMetaData
}

type DLCTowerBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerBuffMetaData
}

type DLCTowerFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DLCTowerFloorMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DLCTowerFloorMetaData
}

type DLCTowerFloorMetaDataReader_StructKey struct {
	// Fields
	ConfigType int32
	Floor      int32
}

type DLCTowerMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerMonsterMetaData
}

type DLCTowerRankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerRankMetaData
}

type DLCTowerScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerScheduleMetaData
}

type DLCTowerStyleBonusMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerStyleBonusMetaData
}

type DLCTowerWaveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerWaveMetaData
}

type DLCTowerWeatherMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DLCTowerWeatherMetaData
}

type DormitoryAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryAvatarMetaData
}

type DormitoryComfortMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryComfortMetaData
}

type DormitoryDecorationEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryDecorationEffectMetaData
}

type DormitoryDecorationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []DormitoryDecorationMetaDataReader_StructKey
	ItemOffsets []int32

	Data []DormitoryDecorationMetaData
}

type DormitoryDecorationMetaDataReader_StructKey struct {
	// Fields
	DecorationLevel int32
	TargetHouse     int32
}

type DormitoryDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryDialogMetaData
}

type DormitoryEventDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryEventDialogMetaData
}

type DormitoryEventSequenceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryEventSequenceMetaData
}

type DormitoryEventWeightMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryEventWeightMetaData
}

type DormitoryFacilityBonusDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryFacilityBonusDropMetaData
}

type DormitoryFacilityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryFacilityMetaData
}

type DormitoryFurnitureCollectRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryFurnitureCollectRewardMetaData
}

type DormitoryFurnitureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryFurnitureMetaData
}

type DormitoryFurnitureSuitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryFurnitureSuitMetaData
}

type DormitoryHouseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryHouseMetaData
}

type DormitoryInteractEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryInteractEventMetaData
}

type DormitoryRoomMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryRoomMetaData
}

type DormitoryTalkEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryTalkEventMetaData
}

type DormitoryVoiceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []DormitoryVoiceMetaData
}

type DormitoryVoteActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryVoteActivityMetaData
}

type DormitoryVotePrizeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryVotePrizeMetaData
}

type DormitoryVoteProductionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DormitoryVoteProductionMetaData
}

type DownLoadConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DownLoadConfigMetaData
}

type DownloadInterfaceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DownloadInterfaceMetaData
}

type DreamMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DreamMetaData
}

type DreamMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DreamMissionMetaData
}

type DreamRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DreamRewardMetaData
}

type DreamUnlockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DreamUnlockMetaData
}

type DressMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DressMetaData
}

type DressTagDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DressTagDataMetaData
}

type DungeonsMirrorRecoveryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DungeonsMirrorRecoveryMetaData
}

type DutyDailyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DutyDailyMetaData
}

type DutyWeeklyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DutyWeeklyMetaData
}

type DynamicHardLvMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []DynamicHardLvMetaData
}

type ElevatorConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElevatorConfigMetaData
}

type ElfBreakMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ElfBreakMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ElfBreakMetaData
}

type ElfBreakMetaDataReader_StructKey struct {
	// Fields
	BreakStep int32
	ElfID     int32
}

type ElfCaptainSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfCaptainSkillMetaData
}

type ElfCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfCardMetaData
}

type ElfFragmentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfFragmentMetaData
}

type ElfGalEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfGalEventMetaData
}

type ElfLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfLevelMetaData
}

type ElfMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfMetaData
}

type ElfRarityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfRarityMetaData
}

type ElfSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfSkillMetaData
}

type ElfSkillTreeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ElfSkillTreeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ElfSkillTreeMetaData
}

type ElfSkillTreeMetaDataReader_StructKey struct {
	// Fields
	ElfSkillID int32
	ElfSkillLv int32
}

type ElfStarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ElfStarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ElfStarMetaData
}

type ElfStarMetaDataReader_StructKey struct {
	// Fields
	ElfID int32
	Star  int32
}

type ElfStoryActMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ElfStoryActMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ElfStoryActMetaData
}

type ElfStoryActMetaDataReader_StructKey struct {
	// Fields
	StoryId     int32
	TriggerTime int32
}

type ElfStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfStoryMetaData
}

type ElfTalentLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ElfTalentLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ElfTalentLevelMetaData
}

type ElfTalentLevelMetaDataReader_StructKey struct {
	// Fields
	TalentID    int32
	TalentLevel int32
}

type ElfTalentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ElfTalentMetaData
}

type ElfTrialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ElfTrialMetaData
}

type EliteAbilityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []EliteAbilityMetaData
}

type EliteChapterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EliteChapterMetaData
}

type EliteStageCompensationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EliteStageCompensationMetaData
}

type EliteStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EliteStageMetaData
}

type EmojiFaceDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EmojiFaceDataMetaData
}

type EmojiFacePageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EmojiFacePageMetaData
}

type EntryPerformMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []EntryPerformMetaData
}

type EntryThemeDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EntryThemeDataMetaData
}

type EntryThemeItemDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EntryThemeItemDataMetaData
}

type EntryThemeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []EntryThemeMetaData
}

type EntryThemeSampleDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EntryThemeSampleDataMetaData
}

type EntryThemeScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []EntryThemeScheduleMetaData
}

type EntryThemeTagDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []EntryThemeTagDataMetaData
}

type EquipFilterTagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipFilterTagMetaData
}

type EquipForgeExchangeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipForgeExchangeMetaData
}

type EquipForgeGenerationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipForgeGenerationMetaData
}

type EquipForgeKeepAffixMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []EquipForgeKeepAffixMetaData
}

type EquipForgeMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipForgeMeta
}

type EquipForgeRemindReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipForgeRemindMeta
}

type EquipForgeShadowMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipForgeShadowMetaData
}

type EquipFragMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipFragMetaData
}

type EquipmentActivitySetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipmentActivitySetMetaData
}

type EquipmentLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipmentLevelMetaData
}

type EquipmentSetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipmentSetMetaData
}

type EquipSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EquipSkillMetaData
}

type EquipTypeDetailMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []EquipTypeDetailMetaDataReader_StructKey
	ItemOffsets []int32

	Data []EquipTypeDetailMetaData
}

type EquipTypeDetailMetaDataReader_StructKey struct {
	// Fields
	EquipmentType int32
	SortType      int32
}

type EvaluateDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EvaluateDialogMetaData
}

type EvaluateIntroMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EvaluateIntroMetaData
}

type EventCollectionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EventCollectionMetaData
}

type EventDialogDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []EventDialogDataMetaData
}

type ExaminationAnswerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExaminationAnswerMetaData
}

type ExaminationQuestionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExaminationQuestionMetaData
}

type ExaminationRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ExaminationRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ExaminationRewardMetaData
}

type ExaminationRewardMetaDataReader_StructKey struct {
	// Fields
	ConfigID                 int32
	ExaminationQuestionIndex int32
}

type ExaminationScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExaminationScheduleMetaData
}

type ExBossMonsterLevelConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ExBossMonsterLevelConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ExBossMonsterLevelConfigMetaData
}

type ExBossMonsterLevelConfigMetaDataReader_StructKey struct {
	// Fields
	MonsterLevel uint8
	ID           int32
}

type ExBossMonsterScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ExBossMonsterScoreRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ExBossMonsterScoreRewardMetaData
}

type ExBossMonsterScoreRewardMetaDataReader_StructKey struct {
	// Fields
	RewardConfigID int32
	ScoreLineNode  int32
}

type ExBossMonsterWeatherMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExBossMonsterWeatherMetaData
}

type ExBossTransferInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExBossTransferInfoMetaData
}

type ExhibitionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ExhibitionMetaData
}

type ExposureRateDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExposureRateDataMetaData
}

type ExScheduleTextmapListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []ExScheduleTextmapListMetaData
}

type ExScheduleTextmapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExScheduleTextmapMetaData
}

type ExtraShortStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraShortStoryMetaData
}

type ExtraStoryAchieveDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryAchieveDisplayMetaData
}

type ExtraStoryAchieveGroupyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryAchieveGroupyMetaData
}

type ExtraStoryAchieveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryAchieveMetaData
}

type ExtraStoryActMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryActMetaData
}

type ExtraStoryAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryAreaMetaData
}

type ExtraStoryChallengeModeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ExtraStoryChallengeModeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ExtraStoryChallengeModeMetaData
}

type ExtraStoryChallengeModeMetaDataReader_StructKey struct {
	// Fields
	ChapterId  int32
	Difficulty int32
}

type ExtraStoryChapterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryChapterMetaData
}

type ExtraStoryLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryLineMetaData
}

type ExtraStoryThemeDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryThemeData
}

type ExtraStoryThemeScheduleReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ExtraStoryThemeSchedule
}

type FakeAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FakeAvatarMetaData
}

type FarmBuffScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmBuffScheduleMetaData
}

type FarmEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmEventMetaData
}

type FarmLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmLevelMetaData
}

type FarmMaterialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmMaterialMetaData
}

type FarmScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmScheduleMetaData
}

type FarmSlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmSlotMetaData
}

type FarmSpeedUpMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmSpeedUpMetaData
}

type FarmUISkinMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FarmUISkinMetaData
}

type FastPassLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FastPassLevelMetaData
}

type FastPassMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []FastPassMetaData
}

type FeaturedContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FeaturedContentMetaData
}

type FeatureSubPakConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []FeatureSubPakConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []FeatureSubPakConfigMetaData
}

type FeatureSubPakConfigMetaDataReader_StructKey struct {
	// Fields
	ContextName StrWithPrefix16
	Tag         StrWithPrefix16
}

type FixedPlotUIConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []FixedPlotUIConfigMetaData
}

type FlopActivityPanelReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FlopActivityPanel
}

type FoundationRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []FoundationRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []FoundationRewardMetaData
}

type FoundationRewardMetaDataReader_StructKey struct {
	// Fields
	Level        int32
	Product_name StrWithPrefix16
}

type FoundationTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []FoundationTypeMetaData
}

type FOWEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []FOWEffectMetaData
}

type FrameDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FrameDataMetaData
}

type FrontEndlessBattleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FrontEndlessBattleConfigMetaData
}

type FrontEndlessFloorConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FrontEndlessFloorConfigMetaData
}

type FrontEndlessStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FrontEndlessStageMetaData
}

type FurnitureTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []FurnitureTypeMetaData
}

type GachaAdventureDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GachaAdventureDisplayMetaData
}

type GachaBoxConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GachaBoxConfigMetaData
}

type GachaEntranceManageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GachaEntranceManageMetaData
}

type GachaFirstDiscountDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GachaFirstDiscountData
}

type GachaGroupManageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GachaGroupManageMetaData
}

type GachaPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GachaPanelMetaData
}

type GalAvatarStandByMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GalAvatarStandByMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GalAvatarStandByMetaData
}

type GalAvatarStandByMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	DressID  int32
}

type GalEventEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GalEventEffectMetaData
}

type GalEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GalEventMetaData
}

type GalEventTwinsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GalEventTwinsMetaData
}

type GardenCropDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GardenCropDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GardenCropDataMetaData
}

type GardenCropDataMetaDataReader_StructKey struct {
	// Fields
	CropID   int32
	GardenID int32
}

type GardenScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GardenScheduleMetaData
}

type GardenWeatherDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GardenWeatherDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GardenWeatherDataMetaData
}

type GardenWeatherDataMetaDataReader_StructKey struct {
	// Fields
	GardenID  int32
	WeatherID int32
}

type GenActivityRewardScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GenActivityRewardScheduleMetaData
}

type GenActivityRewardShowItemsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GenActivityRewardShowItemsMetaData
}

type GeneralActivityActUIMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityActUIMetaData
}

type GeneralActivityDisplayDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityDisplayDataMetaData
}

type GeneralActivityLinkDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityLinkDataMetaData
}

type GeneralActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityMetaData
}

type GeneralActivityOptionalBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityOptionalBuffMetaData
}

type GeneralActivityRankDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GeneralActivityRankDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GeneralActivityRankDataMetaData
}

type GeneralActivityRankDataMetaDataReader_StructKey struct {
	// Fields
	Rank     int32
	RankData int32
}

type GeneralActivityScoreDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityScoreDataMetaData
}

type GeneralActivityStageGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityStageGroupMetaData
}

type GeneralActivityStageHomuMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityStageHomuMetaData
}

type GeneralActivityStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityStageMetaData
}

type GeneralActivityStageRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralActivityStageRewardMetaData
}

type GeneralActivityTicketMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GeneralActivityTicketMetaData
}

type GeneralRarityDisplayConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralRarityDisplayConfigMetaData
}

type GeneralScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GeneralScheduleMetaData
}

type GeniusCommonTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []GeniusCommonTowerMetaData
}

type GeniusEndlessTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []GeniusEndlessTowerMetaData
}

type GeniusEndlessTowerRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []GeniusEndlessTowerRewardMetaData
}

type GeniusTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []GeniusTowerMetaData
}

type GerenalRulesEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GerenalRulesEffectMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GerenalRulesEffectMetaData
}

type GerenalRulesEffectMetaDataReader_StructKey struct {
	// Fields
	RuleID int32
	Value  int32
}

type GiftPackAdditionalDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GiftPackAdditionalDisplayMetaData
}

type GiftPackMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GiftPackMetaData
}

type GlobalArmadaReunionLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GlobalArmadaReunionLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GlobalArmadaReunionLevelMetaData
}

type GlobalArmadaReunionLevelMetaDataReader_StructKey struct {
	// Fields
	Level      int32
	ScheduleID int32
}

type GlobalArmadaReunionScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalArmadaReunionScheduleMetaData
}

type GlobalExploreActionDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreActionDisplayMetaData
}

type GlobalExploreActionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreActionMetaData
}

type GlobalExploreAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreAreaMetaData
}

type GlobalExploreBuffDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GlobalExploreBuffDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GlobalExploreBuffDataMetaData
}

type GlobalExploreBuffDataMetaDataReader_StructKey struct {
	// Fields
	AreaBuffID int32
	Level      int32
}

type GlobalExploreDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreDataMetaData
}

type GlobalExploreEntityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GlobalExploreEntityMetaData
}

type GlobalExploreEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreEventMetaData
}

type GlobalExploreFlagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreFlagMetaData
}

type GlobalExploreGridMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalExploreGridMapMetaData
}

type GlobalExploreKingdomMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreKingdomMetaData
}

type GlobalExploreLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreLevelMetaData
}

type GlobalExploreMapPathMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GlobalExploreMapPathMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GlobalExploreMapPathMetaData
}

type GlobalExploreMapPathMetaDataReader_StructKey struct {
	// Fields
	PosX int32
	PosY int32
}

type GlobalExploreMessageBoardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreMessageBoardMetaData
}

type GlobalExplorePlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExplorePlotMetaData
}

type GlobalExploreQuestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreQuestMetaData
}

type GlobalExploreStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalExploreStageMetaData
}

type GlobalExploreStageScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []GlobalExploreStageScoreMetaData
}

type GlobalSupportRewardConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalSupportRewardConfigMetaData
}

type GlobalWarAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []GlobalWarAreaMetaData
}

type GlobalWarAvatarBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalWarAvatarBuffMetaData
}

type GlobalWarAvatarCollectionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalWarAvatarCollectionMetaData
}

type GlobalWarBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalWarBuffMetaData
}

type GlobalWarDamageTextMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalWarDamageTextMetaData
}

type GlobalWarExchangeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalWarExchangeMetaData
}

type GlobalWarIsolatePointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalWarIsolatePointMetaData
}

type GlobalWarPhotoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalWarPhotoMetaData
}

type GlobalWarPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GlobalWarPointMetaData
}

type GlobalWarScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalWarScheduleMetaData
}

type GlobalWarScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalWarScoreMetaData
}

type GlobalWarSpecialBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GlobalWarSpecialBuffMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GlobalWarSpecialBuffMetaData
}

type GlobalWarSpecialBuffMetaDataReader_StructKey struct {
	// Fields
	SpecialBuffLevel uint16
	SpecialBuffType  uint16
}

type GlobalWarStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GlobalWarStageMetaData
}

type GoBackFreeSelectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackFreeSelectMetaData
}

type GoBackFundMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackFundMetaData
}

type GoBackFundRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GoBackFundRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GoBackFundRewardMetaData
}

type GoBackFundRewardMetaDataReader_StructKey struct {
	// Fields
	FundID int32
	Sort   int32
}

type GoBackGachaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GoBackGachaMetaData
}

type GoBackGamePlayConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GoBackGamePlayConfigMetaData
}

type GoBackGrowUpActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GoBackGrowUpActivityMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GoBackGrowUpActivityMetaData
}

type GoBackGrowUpActivityMetaDataReader_StructKey struct {
	// Fields
	GrowUpConfigID uint8
	Rank           uint8
}

type GoBackLoginImgMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackLoginImgMetaData
}

type GoBackMissionDayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GoBackMissionDayMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GoBackMissionDayMetaData
}

type GoBackMissionDayMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	TabId      int32
}

type GoBackMissionScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackMissionScoreMetaData
}

type GoBackScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackScheduleMetaData
}

type GoBackScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackScoreRewardMetaData
}

type GoBackTabMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GoBackTabMetaData
}

type GoBackWebMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GoBackWebMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GoBackWebMetaData
}

type GoBackWebMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	TabID      int32
}

type GodWarAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarAreaMetaData
}

type GodWarAvatarLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarAvatarLevelDataMetaData
}

type GodWarAvatarTaleScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarAvatarTaleScheduleMetaData
}

type GodWarAvatarUpScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarAvatarUpScheduleMetaData
}

type GodWarBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarBuffMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarBuffMetaData
}

type GodWarBuffMetaDataReader_StructKey struct {
	// Fields
	BuffID    uint32
	BuffLevel uint32
}

type GodWarBuffPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarBuffPoolMetaData
}

type GodWarBuffSuitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarBuffSuitMetaData
}

type GodWarChallengeBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarChallengeBuffMetaData
}

type GodWarChallengeLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarChallengeLevelMetaData
}

type GodWarChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarChallengeMetaData
}

type GodWarChallengeRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarChallengeRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarChallengeRewardMetaData
}

type GodWarChallengeRewardMetaDataReader_StructKey struct {
	// Fields
	Step   int32
	TaleID int32
}

type GodWarChapterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarChapterMetaData
}

type GodWarChar2DMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarChar2DMetaData
}

type GodWarClientDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarClientDataMetaData
}

type GodWarCollectionDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarCollectionDataMetaData
}

type GodWarCollectionSuitDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarCollectionSuitDataMetaData
}

type GodWarEventFlagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarEventFlagMetaData
}

type GodWarEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarEventMetaData
}

type GodWarEventPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarEventPanelMetaData
}

type GodWarEventPlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarEventPlotMetaData
}

type GodWarEventPropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarEventPropMetaData
}

type GodWarEventScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarEventScoreMetaData
}

type GodWarEventStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarEventStageMetaData
}

type GodWarExtraItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarExtraItemMetaData
}

type GodWarExtraItemTextMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarExtraItemTextMetaData
}

type GodWarHardLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarHardLevelMetaData
}

type GodWarHardLevelScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarHardLevelScheduleMetaData
}

type GodWarLobbyAreaTriggerDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarLobbyAreaTriggerDataMetaData
}

type GodWarLobbyAvatarDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarLobbyAvatarDataMetaData
}

type GodWarLobbyButtonMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarLobbyButtonMetaData
}

type GodWarLobbyDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarLobbyDataMetaData
}

type GodWarLobbyInteractActionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarLobbyInteractActionMetaData
}

type GodWarLobbyInteractPropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarLobbyInteractPropMetaData
}

type GodWarMainAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarMainAvatarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarMainAvatarMetaData
}

type GodWarMainAvatarMetaDataReader_StructKey struct {
	// Fields
	GodWarID     uint16
	MainAvatarID uint32
}

type GodWarMainMissionNodeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarMainMissionNodeMetaData
}

type GodWarMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarMapMetaData
}

type GodWarMaterialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarMaterialMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarMaterialMetaData
}

type GodWarMaterialMetaDataReader_StructKey struct {
	// Fields
	GodWarID   int32
	MaterialID int32
}

type GodWarMissionDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarMissionDataMetaData
}

type GodWarNodeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarNodeMetaData
}

type GodWarParentTabDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarParentTabDataMetaData
}

type GodWarPlotStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarPlotStageMetaData
}

type GodWarPunishBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarPunishBuffMetaData
}

type GodWarPunishRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarPunishRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarPunishRewardMetaData
}

type GodWarPunishRewardMetaDataReader_StructKey struct {
	// Fields
	PunishLevel uint16
	PunishScore uint32
}

type GodWarPunishScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarPunishScoreMetaData
}

type GodWarPunishStepMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarPunishStepMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarPunishStepMetaData
}

type GodWarPunishStepMetaDataReader_StructKey struct {
	// Fields
	PunishStep         uint8
	PunishStepConfigID uint32
}

type GodWarRelationDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarRelationDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarRelationDataMetaData
}

type GodWarRelationDataMetaDataReader_StructKey struct {
	// Fields
	GodWarID uint16
	Level    uint16
	RoleID   uint32
}

type GodWarRoleSkillDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarRoleSkillDataMetaData
}

type GodWarRoleStoryDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarRoleStoryDataMetaData
}

type GodWarSceneConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarSceneConfigMetaData
}

type GodWarScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarScheduleMetaData
}

type GodWarSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarSiteMetaData
}

type GodWarStageHintMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarStageHintMetaData
}

type GodWarStageSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarStageSkillMetaData
}

type GodWarSubTabDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarSubTabDataMetaData
}

type GodWarSupportAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarSupportAvatarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarSupportAvatarMetaData
}

type GodWarSupportAvatarMetaDataReader_StructKey struct {
	// Fields
	GodWarID        uint16
	SupportAvatarID uint32
}

type GodWarSupportLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarSupportLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarSupportLevelMetaData
}

type GodWarSupportLevelMetaDataReader_StructKey struct {
	// Fields
	GodWarID uint16
	Level    uint32
}

type GodWarTaleClientDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GodWarTaleClientDataMetaData
}

type GodWarTaleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarTaleMetaData
}

type GodWarTalentClientDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarTalentClientDataMetaData
}

type GodWarTalentDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarTalentDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarTalentDataMetaData
}

type GodWarTalentDataMetaDataReader_StructKey struct {
	// Fields
	GodWarID uint16
	TalentID uint32
}

type GodWarTalentLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarTalentLevelDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarTalentLevelDataMetaData
}

type GodWarTalentLevelDataMetaDataReader_StructKey struct {
	// Fields
	TalentLevel uint16
	TalentID    uint32
}

type GodWarTaleScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarTaleScheduleMetaData
}

type GodWarTeleportMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []GodWarTeleportMetaData
}

type GodWarTicketMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GodWarTicketMetaData
}

type GodWarUseAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GodWarUseAvatarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GodWarUseAvatarMetaData
}

type GodWarUseAvatarMetaDataReader_StructKey struct {
	// Fields
	ChapterID int32
	Step      int32
}

type GrandKeyBuffActiveInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GrandKeyBuffActiveInfoMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GrandKeyBuffActiveInfoMetaData
}

type GrandKeyBuffActiveInfoMetaDataReader_StructKey struct {
	// Fields
	Order               int32
	UnlockGrandKeyLevel int32
}

type GrandKeyBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GrandKeyBuffMetaData
}

type GrandKeyLevelReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GrandKeyLevelReader_StructKey
	ItemOffsets []int32

	Data []GrandKeyLevel
}

type GrandKeyLevelReader_StructKey struct {
	// Fields
	GrandKeyID int32
	Level      int32
}

type GrandKeyMainWeaponReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GrandKeyMainWeapon
}

type GrandKeyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GrandKeyMetaData
}

type GrandKeySkillLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GrandKeySkillLimitMetaData
}

type GrandKeyWeaponContrastReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GrandKeyWeaponContrast
}

type GrandKeyWeaponRaidSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GrandKeyWeaponRaidSkillMetaData
}

type GratuityMessageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GratuityMessageMetaData
}

type GratuityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []GratuityScheduleMetaData
}

type GratuityStageClassMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GratuityStageClassMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GratuityStageClassMetaData
}

type GratuityStageClassMetaDataReader_StructKey struct {
	// Fields
	GratuityScheduleID int32
	StageClassID       int32
}

type GratuityStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GratuityStageMetaData
}

type GreedyEndlessBattleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GreedyEndlessBattleConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GreedyEndlessBattleConfigMetaData
}

type GreedyEndlessBattleConfigMetaDataReader_StructKey struct {
	// Fields
	BattleConfig int32
	Floor        int32
	GroupLevel   int32
}

type GreedyEndlessBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GreedyEndlessBuffMetaData
}

type GreedyEndlessFloorConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GreedyEndlessFloorConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GreedyEndlessFloorConfigMetaData
}

type GreedyEndlessFloorConfigMetaDataReader_StructKey struct {
	// Fields
	Floor         int32
	FloorConfigID int32
	GroupLevel    int32
}

type GreedyEndlessGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GreedyEndlessGroupMetaData
}

type GreedyEndlessMechanismMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GreedyEndlessMechanismMetaData
}

type GreedyEndlessPlayerGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GreedyEndlessPlayerGroupMetaData
}

type GreedyEndlessRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GreedyEndlessRankRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GreedyEndlessRankRewardMetaData
}

type GreedyEndlessRankRewardMetaDataReader_StructKey struct {
	// Fields
	GroupLevel  int32
	RankPercent int32
}

type GreedyEndlessScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RemoteTime
	ItemOffsets []int32

	Data []GreedyEndlessScheduleMetaData
}

type GreedyEndlessSettleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []GreedyEndlessSettleConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []GreedyEndlessSettleConfigMetaData
}

type GreedyEndlessSettleConfigMetaDataReader_StructKey struct {
	// Fields
	GroupLevel           int32
	SettleRewardConfigID int32
}

type GreedyEndlessWeatherMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []GreedyEndlessWeatherMetaData
}

type HalfBalanceModeAttrMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HalfBalanceModeAttrMetaData
}

type HalfBalanceModeDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HalfBalanceModeDataMetaData
}

type HoukaiTownBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownBuffMetaData
}

type HoukaiTownBuildingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownBuildingMetaData
}

type HoukaiTownChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownChallengeMetaData
}

type HoukaiTownEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownEventMetaData
}

type HoukaiTownItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownItemMetaData
}

type HoukaiTownLuckLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownLuckLevelMetaData
}

type HoukaiTownMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownMapMetaData
}

type HoukaiTownPathMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []HoukaiTownPathMetaDataReader_StructKey
	ItemOffsets []int32

	Data []HoukaiTownPathMetaData
}

type HoukaiTownPathMetaDataReader_StructKey struct {
	// Fields
	PrePosition int32
	TowerID     int32
}

type HoukaiTownQAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownQAvatarMetaData
}

type HoukaiTownQBattleSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownQBattleSkillMetaData
}

type HoukaiTownQbossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownQbossMetaData
}

type HoukaiTownQMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownQMonsterMetaData
}

type HoukaiTownRefreshMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownRefreshMetaData
}

type HoukaiTownSpeedLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownSpeedLevelMetaData
}

type HoukaiTownStrengthLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []HoukaiTownStrengthLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []HoukaiTownStrengthLevelMetaData
}

type HoukaiTownStrengthLevelMetaDataReader_StructKey struct {
	// Fields
	LevelType uint8
	Strength  int32
}

type HoukaiTownTerrainSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownTerrainSkillMetaData
}

type HoukaiTownTutorialStepMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HoukaiTownTutorialStepMetaData
}

type HybridRelayPhaseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []HybridRelayPhaseMetaData
}

type HybridSiteBossDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []HybridSiteBossDataMetaData
}

type HybridSiteCameraMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []HybridSiteCameraMetaDataReader_StructKey
	ItemOffsets []int32

	Data []HybridSiteCameraMetaData
}

type HybridSiteCameraMetaDataReader_StructKey struct {
	// Fields
	ChapterID    uint32
	HybridSiteID uint32
}

type HybridSiteContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []HybridSiteContentMetaData
}

type HybridSiteDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []HybridSiteDataMetaData
}

type HybridSiteRemindMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []HybridSiteRemindMetaData
}

type ImgPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ImgPanelMetaData
}

type InControlActionMapDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []InControlActionMapMetaData
}

type InControlActionSubTypeInfoDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []InControlActionSubTypeInfoMetaData
}

type InControlBattleOverrideMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []InControlBattleOverrideMetaData
}

type InControlControlTypeInfoDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []InControlControlTypeInfoMetaData
}

type InControlKeyInfoDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []InControlKeyInfoMetaData
}

type InControlUIButtonConfigReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []InControlUIButtonConfigReader_StructKey
	ItemOffsets []int32

	Data []InControlUIButtonConfigMetaData
}

type InControlUIButtonConfigReader_StructKey struct {
	// Fields
	ContextName      StrWithPrefix16
	PlayerActionName StrWithPrefix16
}

type InLevelBuffUIMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []InLevelBuffUIMetaData
}

type InLevelDialogReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []InLevelDialogMetaData
}

type InviteeActivityGobackConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []InviteeActivityGobackConfigMetaData
}

type InviteeActivityNewbieConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []InviteeActivityNewbieConfigMetaData
}

type InviteeActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []InviteeActivityScheduleMetaData
}

type InviterActivityRewardConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []InviterActivityRewardConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []InviterActivityRewardConfigMetaData
}

type InviterActivityRewardConfigMetaDataReader_StructKey struct {
	// Fields
	ConfigID   int32
	InviteeNum int32
}

type InviterActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []InviterActivityScheduleMetaData
}

type ItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ItemMetaData
}

type JigsawGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []JigsawGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []JigsawGroupMetaData
}

type JigsawGroupMetaDataReader_StructKey struct {
	// Fields
	GroupID  int32
	JigsawID int32
}

type JigsawMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []JigsawMetaData
}

type KingdomsAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []KingdomsAreaMetaData
}

type KingdomsBossPointScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []KingdomsBossPointScheduleMetaData
}

type KingdomsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsMetaData
}

type KingdomsPhaseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsPhaseMetaData
}

type KingdomsPhaseTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []KingdomsPhaseTypeMetaData
}

type KingdomsPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []KingdomsPointMetaData
}

type KingdomsStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []KingdomsStageMetaData
}

type KingdomsWarRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsWarRankRewardMetaData
}

type KingdomsWarScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsWarScheduleMetaData
}

type KingdomsWarScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsWarScoreRewardMetaData
}

type KingdomsWarStageCurrencyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []KingdomsWarStageCurrencyMetaData
}

type KingdomsWarStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsWarStoryMetaData
}

type KingdomsWarVoteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []KingdomsWarVoteMetaData
}

type KingdomsWarWaveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []KingdomsWarWaveMetaDataReader_StructKey
	ItemOffsets []int32

	Data []KingdomsWarWaveMetaData
}

type KingdomsWarWaveMetaDataReader_StructKey struct {
	// Fields
	GroupID int32
	Wave    int32
}

type LevelChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelChallengeMetaData
}

type LevelMatrixArmadaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixArmadaMetaData
}

type LevelMatrixBlockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixBlockMetaData
}

type LevelMatrixBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixBuffMetaData
}

type LevelMatrixEventDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixEventDialogMetaData
}

type LevelMatrixFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixFloorMetaData
}

type LevelMatrixFloorRequirementMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixFloorRequirementMetaData
}

type LevelMatrixGoodsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixGoodsMetaData
}

type LevelMatrixGridIndexMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixGridIndexMetaData
}

type LevelMatrixLevelLogicMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixLevelLogicMetaData
}

type LevelMatrixMapInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixMapInfoMetaData
}

type LevelMatrixMessageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixMessageMetaData
}

type LevelMatrixMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixMonsterMetaData
}

type LevelMatrixPresetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixPresetMetaData
}

type LevelMatrixScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixScheduleMetaData
}

type LevelMatrixScoreBonusMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixScoreBonusMetaData
}

type LevelMatrixSuddenEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixSuddenEventMetaData
}

type LevelMatrixThemeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixThemeMetaData
}

type LevelMatrixTreasureBoxMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixTreasureBoxMetaData
}

type LevelMatrixUseItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMatrixUseItemMetaData
}

type LevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelMetaData
}

type LevelResetCostMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelResetCostMetaData
}

type LevelSaveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelSaveMetaData
}

type LevelStatisticsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelStatisticsMetaData
}

type LevelTrialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []LevelTrialMetaData
}

type LevelTutorialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LevelTutorialMetaData
}

type LinearMissionDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LinearMissionData
}

type LinkDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []LinkDataMetaData
}

type LoadingTipMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LoadingTipMetaData
}

type LoadingTipStringMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []LoadingTipStringMetaData
}

type LoginWishActivityLinkMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []LoginWishActivityLinkMetaData
}

type LoginWishActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []LoginWishActivityMetaData
}

type LoginWishActivityRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []LoginWishActivityRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []LoginWishActivityRewardMetaData
}

type LoginWishActivityRewardMetaDataReader_StructKey struct {
	// Fields
	LoginDay   uint8
	ActivityId uint16
}

type LoginWishActivityWishMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []LoginWishActivityWishMetaData
}

type MailDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MailDataMetaData
}

type MailStyleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MailStyleMetaData
}

type MainEventPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MainEventPanelMetaData
}

type MainlineStepMissionDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MainlineStepMissionDataMetaData
}

type MainlineStepMissionDisplayDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MainlineStepMissionDisplayDataMetaData
}

type MainPageActivityEntryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MainPageActivityEntryMetaData
}

type MainPageSequenceDialogInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []MainPageSequenceDialogInfoMetaData
}

type MainstageEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MainstageEventMetaData
}

type MallEntranceDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MallEntranceDataMetaData
}

type MapGrowthMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MapGrowthMetaData
}

type MapLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MapLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MapLevelMetaData
}

type MapLevelMetaDataReader_StructKey struct {
	// Fields
	Level int32
	MapId int32
}

type MapSiteQualityConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []MapSiteQualityConfigMetaData
}

type MapSiteScheduleDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MapSiteScheduleDataMetaData
}

type MapSiteStageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MapSiteStageDataMetaData
}

type MapSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MapSkillMetaData
}

type MapSkillTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MapSkillTypeMetaData
}

type MassiveWarDamageNeedMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []MassiveWarDamageNeedMetaData
}

type MassiveWarDamageRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MassiveWarDamageRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MassiveWarDamageRewardMetaData
}

type MassiveWarDamageRewardMetaDataReader_StructKey struct {
	// Fields
	DamageRank uint8
	ScheduleID uint8
}

type MassiveWarMonsterDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []MassiveWarMonsterDataMetaData
}

type MassiveWarPlayerGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []MassiveWarPlayerGroupMetaData
}

type MassiveWarRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MassiveWarRankRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MassiveWarRankRewardMetaData
}

type MassiveWarRankRewardMetaDataReader_StructKey struct {
	// Fields
	RankId     uint8
	ScheduleID uint8
}

type MassiveWarRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MassiveWarRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MassiveWarRewardMetaData
}

type MassiveWarRewardMetaDataReader_StructKey struct {
	// Fields
	PlayerGroupID uint16
	ScoreLevel    uint16
}

type MassiveWarScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []MassiveWarScheduleMetaData
}

type MassiveWarStageInfoReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MassiveWarStageInfoMetaData
}

type MassiveWarStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MassiveWarStageMetaData
}

type MaterialDeleteDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MaterialDeleteData
}

type MaterialProtectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MaterialProtectMetaData
}

type MaterialProvideElfExpDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MaterialProvideElfExpMetaData
}

type MaterialUseConversionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MaterialUseConversionMetaData
}

type MaterialUseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MaterialUseMetaData
}

type MatrixFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MatrixFloorMetaData
}

type MatrixGridIndexMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MatrixGridIndexMetaData
}

type MatrixMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MatrixMapMetaData
}

type MatrixPanelLinkMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MatrixPanelLinkMetaData
}

type MatrixSuddenEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MatrixSuddenEventMetaData
}

type MatrixThemeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MatrixThemeMetaData
}

type MechMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MechMetaData
}

type MechPaperMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MechPaperMetaData
}

type MechTDDiffcultyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MechTDDiffcultyMetaData
}

type MechTDRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MechTDRewardMetaData
}

type MechTDWeatherMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MechTDWeatherMetaData
}

type MedalMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MedalMetaData
}

type MentorCompanyMissionRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorCompanyMissionRewardMetaData
}

type MentorConstReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []MentorConst
}

type MentorDailyMissionRewardReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorDailyMissionReward
}

type MentorEvaluateReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorEvaluate
}

type MentorExamDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorExamData
}

type MentorFameLevelDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorFameLevelData
}

type MentorLobbyTagDisplayReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorLobbyTagDisplay
}

type MentorMainSceneNoticeReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorMainSceneNotice
}

type MentorRewardAttenuationReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MentorRewardAttenuation
}

type MiniBossExamExplainInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MiniBossExamExplainInfoMetaData
}

type MiniBossExamHardStageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MiniBossExamHardStageDataMetaData
}

type MiniMonopolyActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MiniMonopolyActivityMetaData
}

type MiniMonopolyGridMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MiniMonopolyGridMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MiniMonopolyGridMetaData
}

type MiniMonopolyGridMetaDataReader_StructKey struct {
	// Fields
	GridId uint8
	MapId  uint16
}

type MiniMonopolyItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MiniMonopolyItemMetaData
}

type MiniMonopolyMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MiniMonopolyMapMetaData
}

type MiniMonopolySectorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MiniMonopolySectorMetaData
}

type MiniSkyFireInfoDesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MiniSkyFireInfoDesMetaData
}

type MiniSkyFireInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MiniSkyFireInfoMetaData
}

type MiniSkyFireStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MiniSkyFireStageMetaData
}

type MinuteChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MinuteChallengeMetaData
}

type MissionCategoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MissionCategoryMetaData
}

type MissionDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MissionData
}

type MissionExtraRewardReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MissionExtraRewardMetaData
}

type MissionGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MissionGroupMetaData
}

type MissionPanelConversationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MissionPanelConversationMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MissionPanelConversationMetaData
}

type MissionPanelConversationMetaDataReader_StructKey struct {
	// Fields
	ConversationID uint16
	ActivityID     int32
}

type MissionThemeReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MissionThemeMetaData
}

type MonopolyBuffInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MonopolyBuffInfoMetaData
}

type MonopolyBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MonopolyBuffMetaData
}

type MonopolyDiceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MonopolyDiceMetaData
}

type MonopolyGuessMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonopolyGuessMetaData
}

type MonopolyLevelInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MonopolyLevelInfoMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MonopolyLevelInfoMetaData
}

type MonopolyLevelInfoMetaDataReader_StructKey struct {
	// Fields
	BuffLevelID uint16
	MonopolyID  uint16
}

type MonopolyLotteryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MonopolyLotteryMetaData
}

type MonsterCardActivityDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MonsterCardActivityDataMetaData
}

type MonsterCardBuffDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []MonsterCardBuffDisplayMetaData
}

type MonsterCardLevelLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MonsterCardLevelLimitMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MonsterCardLevelLimitMetaData
}

type MonsterCardLevelLimitMetaDataReader_StructKey struct {
	// Fields
	Level      uint8
	ActivityID uint32
}

type MonsterCardLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MonsterCardLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MonsterCardLevelMetaData
}

type MonsterCardLevelMetaDataReader_StructKey struct {
	// Fields
	Level    uint8
	UniqueID uint32
}

type MonsterCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MonsterCardMetaData
}

type MonsterCardSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MonsterCardSkillMetaData
}

type MonsterCardStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MonsterCardStageMetaData
}

type MonsterCardStarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MonsterCardStarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MonsterCardStarMetaData
}

type MonsterCardStarMetaDataReader_StructKey struct {
	// Fields
	Star     uint8
	UniqueID uint32
}

type MonsterCardTalentBookMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonsterCardTalentBookMetaData
}

type MonsterCardTowerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonsterCardTowerMetaData
}

type MonsterConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MonsterConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MonsterConfigMetaData
}

type MonsterConfigMetaDataReader_StructKey struct {
	// Fields
	MonsterName StrWithPrefix16
	TypeName    StrWithPrefix16
}

type MonsterResistanceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MonsterResistanceMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MonsterResistanceMetaData
}

type MonsterResistanceMetaDataReader_StructKey struct {
	// Fields
	UniqueMonsterID uint32
	MonsterName     StrWithPrefix16
	TypeName        StrWithPrefix16
}

type MonsterWikiDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonsterWikiDataMetaData
}

type MonsterWikiDescConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MonsterWikiDescConfigMetaData
}

type MonsterWikiPageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []MonsterWikiPageMetaData
}

type MPBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MPBuffMetaData
}

type MPLevelDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPLevelData
}

type MPListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPListMetaData
}

type MPPveActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPPveActivityMetaData
}

type MPPveActivityStageGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPPveActivityStageGroupMetaData
}

type MPRaidActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPRaidActivityMetaData
}

type MPRaidActMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPRaidActMetaData
}

type MPRaidRankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPRaidRankMetaData
}

type MPRaidRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPRaidRankRewardMetaData
}

type MPRaidScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPRaidScoreRewardMetaData
}

type MPRaidSeriesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPRaidSeriesMetaData
}

type MPSkillsDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPSkillsData
}

type MPStageBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MPStageBuffMetaData
}

type MPStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPStageMetaData
}

type MPStagePlayerLevelDropDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MPStagePlayerLevelDropDataReader_StructKey
	ItemOffsets []int32

	Data []MPStagePlayerLevelDropMetaData
}

type MPStagePlayerLevelDropDataReader_StructKey struct {
	// Fields
	MinPlayerLevel int32
	StageID        int32
}

type MPTagDetectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []MPTagDetectMetaData
}

type MPTeamSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MPTeamSkillMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MPTeamSkillMetaData
}

type MPTeamSkillMetaDataReader_StructKey struct {
	// Fields
	AvatarId        int32
	MPTeamSkillType int32
}

type MPTrophyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []MPTrophyMetaData
}

type MultiPlatforms_UserInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []MultiPlatforms_UserInfoMetaDataReader_StructKey
	ItemOffsets []int32

	Data []MultiPlatforms_UserInfoMetaData
}

type MultiPlatforms_UserInfoMetaDataReader_StructKey struct {
	// Fields
	AccountType uint8
	UserType    uint8
}

type MuseumMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []MuseumMetaData
}

type NatureCounterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NatureCounterMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NatureCounterMetaData
}

type NatureCounterMetaDataReader_StructKey struct {
	// Fields
	MonsterNature    int32
	PlayerNatureType int32
}

type NetworkErrCodeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NetworkErrCodeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NetworkErrCodeMetaData
}

type NetworkErrCodeMetaDataReader_StructKey struct {
	// Fields
	ErrType StrWithPrefix16
	RetCode StrWithPrefix16
}

type NewbieActivityPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NewbieActivityPanelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NewbieActivityPanelMetaData
}

type NewbieActivityPanelMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	ScheduleID int32
}

type NewbieActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieActivityScheduleMetaData
}

type NewbieActivityStageMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NewbieActivityStageMissionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NewbieActivityStageMissionMetaData
}

type NewbieActivityStageMissionMetaDataReader_StructKey struct {
	// Fields
	ID         int32
	StageLevel int32
}

type NewbieAvatarGuideDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NewbieAvatarGuideDataReader_StructKey
	ItemOffsets []int32

	Data []NewbieAvatarGuideData
}

type NewbieAvatarGuideDataReader_StructKey struct {
	// Fields
	IsAvatarArtifact bool
	AvatarID         int32
}

type NewbieBattleBuffReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NewbieBattleBuffReader_StructKey
	ItemOffsets []int32

	Data []NewbieBattleBuff
}

type NewbieBattleBuffReader_StructKey struct {
	// Fields
	MaxLevel int32
	MinLevel int32
}

type NewbieCumulativeLotteryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []NewbieCumulativeLotteryMetaData
}

type NewbieDevelopRewardPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []NewbieDevelopRewardPanelMetaData
}

type NewbieDormQAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieDormQAvatarMetaData
}

type NewbieEquipAdvDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieEquipAdvData
}

type NewbieEquipTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieEquipTypeMetaData
}

type NewbieFirstPaymentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []NewbieFirstPaymentMetaData
}

type NewbieFirstPaymentRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieFirstPaymentRewardMetaData
}

type NewbieGlobalBuffReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGlobalBuff
}

type NewbieGroupAdvDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGroupAdvData
}

type NewbieGrowUpMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NewbieGrowUpMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NewbieGrowUpMetaData
}

type NewbieGrowUpMetaDataReader_StructKey struct {
	// Fields
	Rank       uint8
	ScheduleID uint32
}

type NewbieGuideDialogueDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGuideDialogueData
}

type NewbieGuideFaqDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGuideFaqData
}

type NewbieGuideFaqListDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGuideFaqListData
}

type NewbieGuideGrowthQuestDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGuideGrowthQuestData
}

type NewbieGuideTutorialLevelDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieGuideTutorialLevelData
}

type NewbieLevelRushMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieLevelRushMetaData
}

type NewbieLevelRushRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NewbieLevelRushRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NewbieLevelRushRewardMetaData
}

type NewbieLevelRushRewardMetaDataReader_StructKey struct {
	// Fields
	UnlockLevel    uint8
	RewardConfigID int32
}

type NewbieLoginPostMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []NewbieLoginPostMetaData
}

type NewbieRecurrenceLoginRewardDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieRecurrenceLoginRewardData
}

type NewbieWebsitesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewbieWebsitesMetaData
}

type NewFeatureGuideDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NewFeatureGuideData
}

type NinjaAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []NinjaAreaMetaData
}

type NinjaAttrMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []NinjaAttrMetaData
}

type NinjaMissionCGMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []NinjaMissionCGMetaData
}

type NinjaScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []NinjaScheduleMetaData
}

type NinjaSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []NinjaSiteMetaData
}

type NinjaSlotLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NinjaSlotLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NinjaSlotLevelMetaData
}

type NinjaSlotLevelMetaDataReader_StructKey struct {
	// Fields
	SlotID uint8
	Level  int32
}

type NinjaSlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []NinjaSlotMetaData
}

type NinjaSpEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []NinjaSpEffectMetaData
}

type NPCLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []NPCLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []NPCLevelMetaData
}

type NPCLevelMetaDataReader_StructKey struct {
	// Fields
	HardLevel      int32
	HardLevelGroup int32
}

type OpenWorldAreaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldArea
}

type OpenWorldBoss030MetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldBoss030MetaData
}

type OpenWorldContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldContentMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldContentMetaData
}

type OpenWorldContentMetaDataReader_StructKey struct {
	// Fields
	ID    int32
	MapID int32
}

type OpenWorldCookDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldCookData
}

type OpenWorldCycleActivityMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldCycleActivityMetaData
}

type OpenworldCycleDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenworldCycleData
}

type OpenWorldDestinyLevelDiffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldDestinyLevelDiffMetaData
}

type OpenWorldDestinyLineLinkMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldDestinyLineLinkMetaData
}

type OpenWorldDynamicHardLvMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldDynamicHardLvMetaData
}

type OpenWorldEventActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldEventActivityMetaData
}

type OpenWorldEventDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldEventData
}

type OpenWorldEventExtraDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldEventExtraDropMetaData
}

type OpenWorldExplorePointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldExplorePointMetaData
}

type OpenWorldFunctionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldFunctionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldFunctionMetaData
}

type OpenWorldFunctionMetaDataReader_StructKey struct {
	// Fields
	Function int32
	MapId    int32
}

type OpenWorldLocationReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldLocation
}

type OpenWorldMapReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldMap
}

type OpenWorldOverViewDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldOverViewData
}

type OpenWorldQuestActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldQuestActivityMetaData
}

type OpenworldQuestBuffDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenworldQuestBuffData
}

type OpenworldQuestChallengeDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenworldQuestChallengeData
}

type OpenworldQuestDataCfgReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenworldQuestDataCfg
}

type OpenWorldQuestDeleteRuleReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldQuestDeleteRuleReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldQuestDeleteRule
}

type OpenWorldQuestDeleteRuleReader_StructKey struct {
	// Fields
	FinishWay  int32
	IsActivate int32
}

type OpenWorldQuestJudgeDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldQuestJudgeDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldQuestJudgeMetaData
}

type OpenWorldQuestJudgeDataReader_StructKey struct {
	// Fields
	JudgeLv int32
	MapId   int32
	QuestLv int32
}

type OpenworldQuestLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenworldQuestLevelDataMetaData
}

type OpenWorldQuestMapLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldQuestMapLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldQuestMapLevelMetaData
}

type OpenWorldQuestMapLevelMetaDataReader_StructKey struct {
	// Fields
	MapId      int32
	QuestLevel int32
}

type OpenWorldQuestMonsterPowerReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldQuestMonsterPowerData
}

type OpenWorldQuestRarityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldQuestRarityMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldQuestRarityMetaData
}

type OpenWorldQuestRarityMetaDataReader_StructKey struct {
	// Fields
	MapId      int32
	QuestLevel int32
}

type OpenWorldQuestRarityRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldQuestRarityRewardMetaData
}

type OpenWorldQuestSlotDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldQuestSlotData
}

type OpenWorldQuestThemeDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldQuestTheme
}

type OpenWorldQuestThemeScheduleDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldQuestThemeScheduleDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldQuestThemeSchedule
}

type OpenWorldQuestThemeScheduleDataReader_StructKey struct {
	// Fields
	AreaID int32
	Step   int32
}

type OpenWorldRewardUpListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OpenWorldRewardUpListMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OpenWorldRewardUpListMetaData
}

type OpenWorldRewardUpListMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	QuestLevel int32
}

type OpenWorldSectionReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldSection
}

type OpenWorldStoryDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldStoryData
}

type OpenWorldStoryDetailMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []OpenWorldStoryDetailMetaData
}

type OpenWorldTaskConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldTaskConfigMetaData
}

type OpenWorldTimerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OpenWorldTimerMetaData
}

type OptionalBuffListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OptionalBuffListMetaData
}

type OptionalBuffPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OptionalBuffPoolMetaData
}

type OptionalGachaDetailDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OptionalGachaDetailDataMetaData
}

type OrderClientMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OrderClientMetaData
}

type OrderItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OrderItemMetaData
}

type OverlappingCatMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OverlappingCatMetaData
}

type OverlappingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OverlappingMetaData
}

type OverlappingRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OverlappingRewardMetaData
}

type OWActivityBossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWActivityBossMetaData
}

type OWActivityBossRatingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWActivityBossRatingMetaData
}

type OWActivityClueMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWActivityClueMetaData
}

type OWActivityLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWActivityLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWActivityLevelMetaData
}

type OWActivityLevelMetaDataReader_StructKey struct {
	// Fields
	Level        int32
	OWActivityID int32
}

type OWActivityPhaseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWActivityPhaseMetaData
}

type OWActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWActivityScheduleMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWActivityScheduleMetaData
}

type OWActivityScheduleMetaDataReader_StructKey struct {
	// Fields
	MapID        int32
	OWActivityID int32
}

type OWAvatarActivityDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWAvatarActivityDataMetaData
}

type OWAvatarActivityFilesDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []OWAvatarActivityFilesData
}

type OWAvatarActivityLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWAvatarActivityLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWAvatarActivityLevelMetaData
}

type OWAvatarActivityLevelMetaDataReader_StructKey struct {
	// Fields
	ActivityLevel uint8
	ActivityID    uint16
}

type OWAvatarActivityQuestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWAvatarActivityQuestMetaData
}

type OWAvatarActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWAvatarActivityScheduleMetaData
}

type OWAvatarCultivateDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWAvatarCultivateDataMetaData
}

type OWAvatarCultivateLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWAvatarCultivateLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWAvatarCultivateLevelMetaData
}

type OWAvatarCultivateLevelMetaDataReader_StructKey struct {
	// Fields
	Level       uint8
	CultivateID int32
}

type OWAvatarUnlockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWAvatarUnlockMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWAvatarUnlockMetaData
}

type OWAvatarUnlockMetaDataReader_StructKey struct {
	// Fields
	ActivityID uint16
	AvatarID   uint16
}

type OWEndlessAreaDetailMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessAreaDetailMetaData
}

type OWEndlessAreaScoreConfigReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessAreaScoreConfig
}

type OWEndlessBattleAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessBattleAreaMetaData
}

type OWEndlessBattleConfigReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessBattleConfigReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessBattleConfig
}

type OWEndlessBattleConfigReader_StructKey struct {
	// Fields
	EndlessConfigType uint8
	BattleConfigID    int32
}

type OWEndlessBossDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessBossData
}

type OWEndlessBuffReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessBuff
}

type OWEndlessEnvironmentMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessEnvironmentMeta
}

type OWEndlessGroupMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessGroupMetaReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessGroupMeta
}

type OWEndlessGroupMetaReader_StructKey struct {
	// Fields
	Type       uint8
	GroupLevel int32
}

type OWEndlessInvadeBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessInvadeBuffMetaData
}

type OWEndlessInvadeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessInvadeMetaData
}

type OWEndlessItemEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessItemEffectMetaData
}

type OWEndlessItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessItemMetaData
}

type OWEndlessMonsterDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessMonsterData
}

type OWEndlessMonsterGroupScoreReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessMonsterGroupScore
}

type OWEndlessMonsterInitTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessMonsterInitTypeMetaData
}

type OWEndlessPlayerBaseRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessPlayerBaseRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessPlayerBaseRewardMetaData
}

type OWEndlessPlayerBaseRewardMetaDataReader_StructKey struct {
	// Fields
	Type       uint8
	GroupLevel int32
	MinScore   int32
}

type OWEndlessPlayerGroupMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessPlayerGroupMetaReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessPlayerGroupMeta
}

type OWEndlessPlayerGroupMetaReader_StructKey struct {
	// Fields
	Type        uint8
	PlayerGroup int32
}

type OWEndlessRewardConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessRewardConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessRewardConfigMetaData
}

type OWEndlessRewardConfigMetaDataReader_StructKey struct {
	// Fields
	ConfigID   uint32
	GroupLevel uint32
}

type OWEndlessScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RemoteTime
	ItemOffsets []int32

	Data []OWEndlessScheduleMetaData
}

type OWEndlessSingleFloorDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessSingleFloorDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessSingleFloorDataMetaData
}

type OWEndlessSingleFloorDataMetaDataReader_StructKey struct {
	// Fields
	Floor      uint8
	ActivityID int32
	StageID    int32
}

type OWEndlessSingleModeStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEndlessSingleModeStageMetaData
}

type OWEndlessSingleMonsterGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWEndlessSingleMonsterGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWEndlessSingleMonsterGroupMetaData
}

type OWEndlessSingleMonsterGroupMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	BattleID   int32
}

type OWEventDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWEventDisplayMetaData
}

type OWHuntActivityDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityDataMetaData
}

type OWHuntActivityFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []OWHuntActivityFloorMetaData
}

type OWHuntActivityHunterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWHuntActivityHunterMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWHuntActivityHunterMetaData
}

type OWHuntActivityHunterMetaDataReader_StructKey struct {
	// Fields
	Hardlevel uint8
	HunterID  int32
}

type OWHuntActivityMachineEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityMachineEventMetaData
}

type OWHuntActivityMachineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityMachineMetaData
}

type OWHuntActivityMapDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityMapDataMetaData
}

type OWHuntActivityMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWHuntActivityMonsterMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWHuntActivityMonsterMetaData
}

type OWHuntActivityMonsterMetaDataReader_StructKey struct {
	// Fields
	Hardlevel     uint8
	MonsterShowID int32
}

type OWHuntActivityProgressMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityProgressMetaData
}

type OWHuntActivityQuestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityQuestMetaData
}

type OWHuntActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityScheduleMetaData
}

type OWHuntActivitySHLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWHuntActivitySHLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWHuntActivitySHLevelMetaData
}

type OWHuntActivitySHLevelMetaDataReader_StructKey struct {
	// Fields
	StrongholdLevel     uint8
	StrongholdLevelType uint8
}

type OWHuntActivityStrongholdsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityStrongholdsMetaData
}

type OWHuntActivityTalentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWHuntActivityTalentMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWHuntActivityTalentMetaData
}

type OWHuntActivityTalentMetaDataReader_StructKey struct {
	// Fields
	TalentLevel uint8
	TalentID    int32
}

type OWHuntActivityTeleportMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWHuntActivityTeleportMetaData
}

type OWTalentDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWTalentDataMetaData
}

type OWTalentLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []OWTalentLevelDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []OWTalentLevelDataMetaData
}

type OWTalentLevelDataMetaDataReader_StructKey struct {
	// Fields
	Level    uint8
	TalentID int32
}

type OWTeleporterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []OWTeleporterMetaData
}

type PanelConversationTriggerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PanelConversationTriggerMetaData
}

type PayActivityDisplayScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PayActivityDisplayScheduleMetaData
}

type PayInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []PayInfoMetaDataReader_StructKey
	ItemOffsets []int32

	Data []PayInfoMetaData
}

type PayInfoMetaDataReader_StructKey struct {
	// Fields
	DevType  uint8
	UserType uint8
}

type PerformEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []PerformEventMetaData
}

type PhoneEntranceAcountOverrideMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []PhoneEntranceAcountOverrideMetaDataReader_StructKey
	ItemOffsets []int32

	Data []PhoneEntranceAcountOverrideMetaData
}

type PhoneEntranceAcountOverrideMetaDataReader_StructKey struct {
	// Fields
	AccountType int32
	EntranceId  int32
}

type PhoneEntranceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PhoneEntranceMetaData
}

type PhoneNoticeDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PhoneNoticeDataMetaData
}

type PhonePendantMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PhonePendantMetaData
}

type PicBGListMetaReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PicBGListdMeta
}

type PictureStepMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []PictureStepMetaDataReader_StructKey
	ItemOffsets []int32

	Data []PictureStepMetaData
}

type PictureStepMetaDataReader_StructKey struct {
	// Fields
	ChoiceID int32
	StepID   int32
}

type PicTutorialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []PicTutorialMetaDataReader_StructKey
	ItemOffsets []int32

	Data []PicTutorialMetaData
}

type PicTutorialMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	ID         int32
}

type PicTutorialStepDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PicTutorialStepDataMetaData
}

type PlayerLevelLockMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlayerLevelLockMetaData
}

type PlayerLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlayerLevelMetaData
}

type PlayerLevelShopGoodsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlayerLevelShopGoodsMetaData
}

type PlayerPrivilegeScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlayerPrivilegeScheduleMetaData
}

type PlotAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlotAvatarMetaData
}

type PlotLineCgDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlotLineCgDataMetaData
}

type PlotlineConditionDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []PlotlineConditionDataMetaData
}

type PlotLineDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []PlotLineDataMetaData
}

type PlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PlotMetaData
}

type PowerTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PowerTypeMetaData
}

type PrayGachaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PrayGachaMetaData
}

type PredownloadAsbMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []PredownloadAsbMetaData
}

type PredownloadAudioMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []PredownloadAudioMetaData
}

type PredownloadAudioPackageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []PredownloadAudioPackageMetaData
}

type PredownloadVideoFileMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []PredownloadVideoFileMetaData
}

type PredownloadVideoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []PredownloadVideoMetaData
}

type ProductRecommendMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ProductRecommendMetaData
}

type ProtectTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ProtectTypeMetaData
}

type PVloginMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []PVloginMetaData
}

type PVZActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []PVZActivityMetaData
}

type PVZMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PVZMonsterMetaData
}

type PVZQavatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []PVZQavatarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []PVZQavatarMetaData
}

type PVZQavatarMetaDataReader_StructKey struct {
	// Fields
	Level     uint32
	QavatarID uint32
}

type PVZSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []PVZSiteMetaData
}

type PVZTileMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []PVZTileMetaDataReader_StructKey
	ItemOffsets []int32

	Data []PVZTileMetaData
}

type PVZTileMetaDataReader_StructKey struct {
	// Fields
	FloorID int32
	TowerID int32
}

type QAvatarBattleBroadCastMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarBattleBroadCastMetaData
}

type QAvatarBattleExpressionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarBattleExpressionMetaData
}

type QAvatarBattleGadgetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarBattleGadgetMetaData
}

type QAvatarBattleGadgetRefreshMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarBattleGadgetRefreshMetaData
}

type QAvatarBattleKillingStreakMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []QAvatarBattleKillingStreakMetaDataReader_StructKey
	ItemOffsets []int32

	Data []QAvatarBattleKillingStreakMetaData
}

type QAvatarBattleKillingStreakMetaDataReader_StructKey struct {
	// Fields
	Num  uint16
	Type int32
}

type QAvatarBattleSpawnZoneMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarBattleSpawnZoneMetaData
}

type QAvatarBattleTextIdMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []QAvatarBattleTextIdMapMetaData
}

type QAvatarMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []QAvatarMissionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []QAvatarMissionMetaData
}

type QAvatarMissionMetaDataReader_StructKey struct {
	// Fields
	AvatarID  int32
	MissionID int32
}

type QAvatarPVPMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarPVPMetaData
}

type QAvatarPVPWeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QAvatarPVPWeaponMetaData
}

type QCandyActivityDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QCandyActivityDataMetaData
}

type QCandyAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QCandyAvatarMetaData
}

type QCandyPvpMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QCandyPvpMapMetaData
}

type QCandyPvpMapPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QCandyPvpMapPoolMetaData
}

type QCandyPvpMapSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QCandyPvpMapSkillMetaData
}

type QCandyRankSectionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []QCandyRankSectionMetaData
}

type QCandySettleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []QCandySettleConfigMetaData
}

type QTEndlessMonsterDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []QTEndlessMonsterData
}

type QTEndlessMonsterWaveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []QTEndlessMonsterWaveMetaData
}

type QTEndlessScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RemoteTime
	ItemOffsets []int32

	Data []QTEndlessScheduleMetaData
}

type QuestionBankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []QuestionBankMetaData
}

type QuestionConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []QuestionConfigMetaData
}

type QuestionScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []QuestionScheduleMetaData
}

type QuickEntryScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QuickEntryScheduleMetaData
}

type QuickSellDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []QuickSellDataMetaData
}

type RaffleClientMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RaffleClientMetaData
}

type RaffleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RaffleConfigMetaData
}

type RaffleRewardConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RaffleRewardConfigMetaData
}

type RaffleScheduleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RaffleScheduleConfigMetaData
}

type RaidActivateScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RaidActivateScheduleMetaData
}

type RanchAreaConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchAreaConfigMetaData
}

type RanchAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchAreaMetaData
}

type RanchChallengeSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchChallengeSiteMetaData
}

type RanchContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchContentMetaData
}

type RanchDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchDataMetaData
}

type RanchMonsterExpMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchMonsterExpMetaData
}

type RanchMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchMonsterMetaData
}

type RanchMonsterSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchMonsterSkillMetaData
}

type RanchSiteDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchSiteDataMetaData
}

type RanchSyntheticMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RanchSyntheticMetaData
}

type RandomDialogActionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RandomDialogActionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RandomDialogActionMetaData
}

type RandomDialogActionMetaDataReader_StructKey struct {
	// Fields
	DialogId int32
	PlotId   int32
}

type RandomDialogCGRawMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomDialogCGRawMetaData
}

type RandomDialogInputMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomDialogInputMetaData
}

type RandomDialogQuestionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomDialogQuestionMetaData
}

type RandomEffectLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RandomEffectLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RandomEffectLevelMetaData
}

type RandomEffectLevelMetaDataReader_StructKey struct {
	// Fields
	EffectID    int32
	EffectLevel int32
}

type RandomEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomEffectMetaData
}

type RandomHideBranchTimeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomHideBranchTimeMetaData
}

type RandomPlotDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomPlotDataMetaData
}

type RandomPlotNPCMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RandomPlotNPCMetaData
}

type RandomSubSelectContentDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RandomSubSelectContentDataMetaData
}

type RandomSubSelectStyleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RandomSubSelectStyleConfigMetaData
}

type RareaffixMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RareaffixMetaData
}

type ReclaimConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimConfigMetaData
}

type ReclaimEventListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimEventListMetaData
}

type ReclaimEventStageListMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimEventStageListMetaData
}

type ReclaimEventVirtualAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimEventVirtualAvatarMetaData
}

type ReclaimEventVirtualStigmataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimEventVirtualStigmataMetaData
}

type ReclaimEventVirtualWeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimEventVirtualWeaponMetaData
}

type ReclaimLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimLevelMetaData
}

type ReclaimRankingRewardDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimRankingRewardData
}

type ReclaimScoreRewardDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimScoreRewardData
}

type ReclaimVirtualGachaChestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReclaimVirtualGachaChestMetaData
}

type RecommendEquipMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RecommendEquipMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RecommendEquipMetaData
}

type RecommendEquipMetaDataReader_StructKey struct {
	// Fields
	EquipID int32
	ID      int32
}

type RedEnvelopeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RedEnvelopeMetaData
}

type RegionInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []RegionInfoMetaData
}

type RelationActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RelationActivityScheduleMetaData
}

type RelationStageBonusDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RelationStageBonusDropMetaData
}

type ReplayLobbyFilterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReplayLobbyFilterMetaData
}

type ReplayLobbyPageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReplayLobbyPageMetaData
}

type ReplayLobbyScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReplayLobbyScoreMetaData
}

type ReplayLobbySubPageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ReplayLobbySubPageMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ReplayLobbySubPageMetaData
}

type ReplayLobbySubPageMetaDataReader_StructKey struct {
	// Fields
	FatherID int32
	ID       int32
}

type ReplayLobbyUploadMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReplayLobbyUploadMetaData
}

type ReportReasonMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReportReasonMetaData
}

type ResourceRetrieveBoxMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ResourceRetrieveBoxMetaData
}

type ResourceRetrieveModuleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ResourceRetrieveModuleMetaData
}

type ResourceRetrieveScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ResourceRetrieveScheduleMetaData
}

type RestaurantActionZoneMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []RestaurantActionZoneMetaData
}

type RestaurantAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantAvatarMetaData
}

type RestaurantIngredientsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantIngredientsMetaData
}

type RestaurantLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantLevelMetaData
}

type RestaurantQuestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantQuestMetaData
}

type RestaurantRecipeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantRecipeMetaData
}

type RestaurantRoomMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantRoomMetaData
}

type RestaurantSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantSkillMetaData
}

type RestaurantWeatherMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestaurantWeatherMetaData
}

type RestrictExtend_StageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RestrictExtend_StageMetaData
}

type ReunionCookBookMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ReunionCookBookMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ReunionCookBookMetaData
}

type ReunionCookBookMetaDataReader_StructKey struct {
	// Fields
	CookGroupID int32
	CookID      int32
}

type ReunionCookRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ReunionCookRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ReunionCookRewardMetaData
}

type ReunionCookRewardMetaDataReader_StructKey struct {
	// Fields
	ID               uint8
	ScoreRewardGroup int32
}

type ReunionCookStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReunionCookStageMetaData
}

type ReunionDisplayScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReunionDisplayScheduleMetaData
}

type ReviveCostTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReviveCostTypeMetaData
}

type ReviveEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ReviveEffectMetaData
}

type ReviveUseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ReviveUseMetaData
}

type RewardDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RewardData
}

type RewardLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []RewardLineMetaData
}

type RewardLineRankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []RewardLineRankMetaData
}

type RewardOverviewPanelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RewardOverviewPanelMetaData
}

type RichAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichAreaMetaData
}

type RichBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RichBuffMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RichBuffMetaData
}

type RichBuffMetaDataReader_StructKey struct {
	// Fields
	BuffID int32
	Level  int32
}

type RichBuildingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RichBuildingMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RichBuildingMetaData
}

type RichBuildingMetaDataReader_StructKey struct {
	// Fields
	BuildingLevel int32
	BuildingType  int32
}

type RichChallengeBossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichChallengeBossMetaData
}

type RichDiceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichDiceMetaData
}

type RichFloorInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RichFloorInfoMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RichFloorInfoMetaData
}

type RichFloorInfoMetaDataReader_StructKey struct {
	// Fields
	Area    int32
	FloorID int32
}

type RichItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichItemMetaData
}

type RichmanMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichmanMetaData
}

type RichMapInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichMapInfoMetaData
}

type RichMonsterInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichMonsterInfoMetaData
}

type RichShopMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RichShopMetaData
}

type RogueActiveCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueActiveCardMetaData
}

type RogueBuffPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RogueBuffPoolMetaData
}

type RogueEliteAbilityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueEliteAbilityMetaData
}

type RogueEliteAbilityTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueEliteAbilityTypeMetaData
}

type RogueGoodsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueGoodsMetaData
}

type RoguePriceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RoguePriceMetaData
}

type RogueStageAvatarPosetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueStageAvatarPosMetaData
}

type RogueStageDebuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueStageDebuffMetaData
}

type RogueStageHardLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueStageHardLevelMetaData
}

type RogueStageLuaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueStageLuaMetaData
}

type RogueStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RogueStageMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RogueStageMetaData
}

type RogueStageMetaDataReader_StructKey struct {
	// Fields
	Level     uint8
	RogueHard uint8
}

type RogueStoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueStoreMetaData
}

type RogueTowerDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RogueTowerDataMetaData
}

type RogueTowerEndSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []RogueTowerEndSiteMetaData
}

type RogueTowerStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueTowerStageMetaData
}

type RogueTriggerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RogueTriggerMetaData
}

type RoutineDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RoutineDataMetaData
}

type RoutineRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RoutineRewardMetaData
}

type RoutineScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RoutineScheduleMetaData
}

type RpgAreaLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgAreaLineMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgAreaLineMetaData
}

type RpgAreaLineMetaDataReader_StructKey struct {
	// Fields
	AreaID    int32
	PreAreaID int32
}

type RpgAreaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgAreaMetaData
}

type RpgBuffDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgBuffDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgBuffDataMetaData
}

type RpgBuffDataMetaDataReader_StructKey struct {
	// Fields
	ID    int32
	Level int32
}

type RpgBuffLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgBuffLimitMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgBuffLimitMetaData
}

type RpgBuffLimitMetaDataReader_StructKey struct {
	// Fields
	ID    int32
	Level int32
}

type RpgBuffSuitClientInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgBuffSuitClientInfoMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgBuffSuitClientInfoMetaData
}

type RpgBuffSuitClientInfoMetaDataReader_StructKey struct {
	// Fields
	BuffDataID int32
	Level      int32
}

type RpgClientMainDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgClientMainDataMetaData
}

type RpgCollectionRewardDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgCollectionRewardDataReader_StructKey
	ItemOffsets []int32

	Data []RpgCollectionRewardMetaData
}

type RpgCollectionRewardDataReader_StructKey struct {
	// Fields
	PositionID int32
	TaleID     int32
}

type RpgDungeonAdventureQuestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgDungeonAdventureQuestMetaData
}

type RpgEventDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgEventDialogMetaData
}

type RpgEventTextMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgEventTextMetaData
}

type RpgFilesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgFilesMetaData
}

type RpgLevelProgressDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgLevelProgressData
}

type RpgMaterialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgMaterialMetaData
}

type RpgMissionCategoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgMissionCategoryMetaData
}

type RpgMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgMissionMetaData
}

type RpgNavSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []RpgNavSiteMetaData
}

type RpgOverallMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgOverallMetaData
}

type RpgQAvatarBattleCollideMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgQAvatarBattleCollideMetaData
}

type RpgQAvatarBattleDivisionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []RpgQAvatarBattleDivisionMetaData
}

type RpgQAvatarBattleSceneMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgQAvatarBattleSceneMetaData
}

type RpgQAvatarBattleSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgQAvatarBattleSiteMetaData
}

type RpgRestaurantMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgRestaurantMetaData
}

type RpgRoleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgRoleMetaData
}

type RpgScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgScheduleMetaData
}

type RpgShootPlayTriggerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgShootPlayTriggerMetaData
}

type RpgSimpleBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgSimpleBuffMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgSimpleBuffMetaData
}

type RpgSimpleBuffMetaDataReader_StructKey struct {
	// Fields
	BuffLevel      uint16
	BuffMaterialID int32
}

type RpgSiteLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgSiteLineMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgSiteLineMetaData
}

type RpgSiteLineMetaDataReader_StructKey struct {
	// Fields
	PreSiteID int32
	SiteID    int32
}

type RpgSiteLockTipsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSiteLockTipsMetaData
}

type RpgSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSiteMetaData
}

type RpgSitePackMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []RpgSitePackMetaData
}

type RpgSiteProgressConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSiteProgressConfigMetaData
}

type RpgSkillDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSkillDisplayMetaData
}

type RpgStageDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgStageDropMetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgStageDropMetaData
}

type RpgStageDropMetaDataReader_StructKey struct {
	// Fields
	DropMaterialID int32
	StageID        int32
}

type RpgStageScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgStageScoreMetaData
}

type RpgStigmataBookMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgStigmataBookMetaData
}

type RpgSuddenEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSuddenEventMetaData
}

type RpgSurvivalRoleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSurvivalRoleMetaData
}

type RpgSurvivalTraitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSurvivalTraitMetaData
}

type RpgSurvivalWeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgSurvivalWeaponMetaData
}

type RpgTaleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgTaleMetaData
}

type RpgTaleRefreshMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgTaleRefreshMetaData
}

type RpgTaleStageEnterTimesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgTaleStageEnterTimesMetaData
}

type RpgTicketMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgTicketMetaData
}

type RpgTowerV2MetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []RpgTowerV2MetaDataReader_StructKey
	ItemOffsets []int32

	Data []RpgTowerV2MetaData
}

type RpgTowerV2MetaDataReader_StructKey struct {
	// Fields
	Floor int16
	Stage int32
}

type RpgUIConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []RpgUIConfigMetaData
}

type RpgVirtualAvatarLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgVirtualAvatarLimitMetaData
}

type RpgZoneMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []RpgZoneMetaData
}

type SanctuaryBuildingActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SanctuaryBuildingActivityMetaData
}

type SanctuaryBuildingLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SanctuaryBuildingLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SanctuaryBuildingLevelMetaData
}

type SanctuaryBuildingLevelMetaDataReader_StructKey struct {
	// Fields
	SanctuaryActivityID int32
	SanctuaryLevel      int32
}

type SanctuaryPlayerGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SanctuaryPlayerGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SanctuaryPlayerGroupMetaData
}

type SanctuaryPlayerGroupMetaDataReader_StructKey struct {
	// Fields
	ActivityID    int32
	PlayerGroupID int32
}

type SanctuaryStageConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SanctuaryStageConfigMetaData
}

type ScDLCAchievementMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ScDLCAchievementMetaData
}

type ScDLCAvatarLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ScDLCAvatarLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ScDLCAvatarLevelMetaData
}

type ScDLCAvatarLevelMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	Level    int32
}

type ScDLCAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCAvatarMetaData
}

type ScDLCChallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCChallengeMetaData
}

type ScDLCChallengeRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCChallengeRewardMetaData
}

type ScDLCCombineTalentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCCombineTalentMetaData
}

type ScDLCExplorePointMaterialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCExplorePointMaterialMetaData
}

type ScDLCFeverAbilityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ScDLCFeverAbilityMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ScDLCFeverAbilityMetaData
}

type ScDLCFeverAbilityMetaDataReader_StructKey struct {
	// Fields
	AbilityLevel int32
	ID           int32
}

type ScDLCFurnitureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ScDLCFurnitureMetaData
}

type ScDLCGachaMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCGachaMetaData
}

type ScDLCImageReplaceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCImageReplaceMetaData
}

type ScDLCInteractActionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCInteractActionMetaData
}

type ScDLCInteractStateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCInteractStateMetaData
}

type ScDLCLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCLevelMetaData
}

type ScDLCMapProgressMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCMapProgressMetaData
}

type ScDLCMissionRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCMissionRewardMetaData
}

type ScDLCMPStageDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCMPStageDisplayMetaData
}

type ScDLCNPCMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCNPCMetaData
}

type ScDLCNPCPositionGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCNPCPositionGroupMetaData
}

type ScDLCNPCPositionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCNPCPositionMetaData
}

type ScDLCOpenFunctionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ScDLCOpenFunctionMetaData
}

type ScDLCShowMapConditionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCShowMapConditionMetaData
}

type ScDLCSkillChipMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCSkillChipMetaData
}

type ScDLCSpecialExplorePointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCSpecialExplorePointMetaData
}

type ScDLCSpecialItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCSpecialItemMetaData
}

type ScDLCStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ScDLCStoryMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ScDLCStoryMetaData
}

type ScDLCStoryMetaDataReader_StructKey struct {
	// Fields
	State uint8
	ID    uint32
}

type ScDLCTalentDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTalentDataMetaData
}

type ScDLCTalentLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ScDLCTalentLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ScDLCTalentLevelMetaData
}

type ScDLCTalentLevelMetaDataReader_StructKey struct {
	// Fields
	TalentID    int32
	TalentLevel int32
}

type ScDLCTeachStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTeachStageMetaData
}

type ScDLCTowerBonusMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerBonusMetaData
}

type ScDLCTowerBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerBuffMetaData
}

type ScDLCTowerFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ScDLCTowerFloorMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ScDLCTowerFloorMetaData
}

type ScDLCTowerFloorMetaDataReader_StructKey struct {
	// Fields
	ConfigType int32
	Floor      int32
}

type ScDLCTowerMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerMonsterMetaData
}

type ScDLCTowerScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerScheduleMetaData
}

type ScDLCTowerScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerScoreRewardMetaData
}

type ScDLCTowerWaveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerWaveMetaData
}

type ScDLCTowerWeatherMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScDLCTowerWeatherMetaData
}

type ScratchTicketDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScratchTicketDataMetaData
}

type ScratchTicketItemDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ScratchTicketItemDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ScratchTicketItemDataMetaData
}

type ScratchTicketItemDataMetaDataReader_StructKey struct {
	// Fields
	PlateDetailID int32
	PlateID       int32
}

type ScratchTicketPlateDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ScratchTicketPlateDataMetaData
}

type SecondAnniversaryMemoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SecondAnniversaryMemoryMetaData
}

type SecondAnniversaryMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SecondAnniversaryMissionMetaData
}

type SeriesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SeriesMetaData
}

type ServantDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ServantDialogMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ServantDialogMetaData
}

type ServantDialogMetaDataReader_StructKey struct {
	// Fields
	DialogID int32
	MapId    int32
}

type ServantTouchMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ServantTouchMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ServantTouchMetaData
}

type ServantTouchMetaDataReader_StructKey struct {
	// Fields
	Id    int32
	MapID int32
}

type SettingAudioVolumeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SettingAudioVolumeMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SettingAudioVolumeMetaData
}

type SettingAudioVolumeMetaDataReader_StructKey struct {
	// Fields
	AudioSettingType   uint8
	ParamIndex         uint8
	AudioSettingOption StrWithPrefix16
}

type SettingGraphicsDeviceLimitMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SettingGraphicsDeviceLimitMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SettingGraphicsDeviceLimitMetaData
}

type SettingGraphicsDeviceLimitMetaDataReader_StructKey struct {
	// Fields
	Device                  StrWithPrefix16
	GraphicsSettingItemName StrWithPrefix16
}

type SettingGraphicsItemLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SettingGraphicsItemLineMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SettingGraphicsItemLineMetaData
}

type SettingGraphicsItemLineMetaDataReader_StructKey struct {
	// Fields
	Hierarchy uint8
	ID        int32
}

type SettingGraphicsTitleLineMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SettingGraphicsTitleLineMetaData
}

type ShareChannelContentMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShareChannelContentMetaData
}

type ShareChannelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []ShareChannelMetaData
}

type ShareConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShareConfigMetaData
}

type SharePakDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SharePakDataMetaData
}

type ShareShowDetailConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShareShowDetailConfigMetaData
}

type ShareSpecialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShareSpecialMetaData
}

type ShareSwitchMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ShareSwitchMetaDataReader_StructKey
	ItemOffsets []int32

	Data []ShareSwitchMetaData
}

type ShareSwitchMetaDataReader_StructKey struct {
	// Fields
	DeviceType int16
	ChannelId  int32
}

type ShopAboutToOnlineItemReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []ShopAboutToOnlineItemReader_StructKey
	ItemOffsets []int32

	Data []ShopAboutToOnlineItem
}

type ShopAboutToOnlineItemReader_StructKey struct {
	// Fields
	ItemID int32
	ShopID int32
}

type ShopDiscountMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShopDiscountMetaData
}

type ShopGoodsClassifyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ShopGoodsClassifyMetaData
}

type ShopGoodsDisplayDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ShopGoodsDisplayDataMetaData
}

type ShopGoodsJumpMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ShopGoodsJumpMetaData
}

type ShopGoodsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShopGoodsMetaData
}

type ShopGoodsPriceRateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShopGoodsPriceRateMetaData
}

type ShopGoodsRefreshTimeTypeDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShopGoodsRefreshTimeTypeData
}

type ShopGoodsScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShopGoodsScheduleMetaData
}

type ShopGoodsTagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ShopGoodsTagMetaData
}

type ShopTabsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ShopTabsMetaData
}

type SignInRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SignInRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SignInRewardMetaData
}

type SignInRewardMetaDataReader_StructKey struct {
	// Fields
	Day   int32
	Month int32
}

type SimulateRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SimulateRankRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SimulateRankRewardMetaData
}

type SimulateRankRewardMetaDataReader_StructKey struct {
	// Fields
	ActivityId  int32
	RankPercent int32
}

type SinDemonExBossLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExBossLevelMetaData
}

type SinDemonExConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExConfigMetaData
}

type SinDemonExMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExMonsterMetaData
}

type SinDemonExMonsterScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SinDemonExMonsterScheduleMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SinDemonExMonsterScheduleMetaData
}

type SinDemonExMonsterScheduleMetaDataReader_StructKey struct {
	// Fields
	ConfigId int32
	ID       int32
}

type SinDemonExRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExRankRewardMetaData
}

type SinDemonExScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExScoreRewardMetaData
}

type SinDemonExSeasonBossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExSeasonBossMetaData
}

type SinDemonExSkillTipsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SinDemonExSkillTipsMetaData
}

type SingleRaidActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SingleRaidActivityMetaData
}

type SingleRaidStepStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SingleRaidStepStageMetaData
}

type SingleWantedActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []SingleWantedActivityMetaData
}

type SingleWantedEquipShowConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []SingleWantedEquipShowConfigMetaData
}

type SingleWantedMaterialTierMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SingleWantedMaterialTierMetaData
}

type SingleWantedProgressMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SingleWantedProgressMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SingleWantedProgressMetaData
}

type SingleWantedProgressMetaDataReader_StructKey struct {
	// Fields
	Progress uint32
	ThemeID  uint32
}

type SingleWantedScheduelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []SingleWantedScheduelMetaData
}

type SingleWantedStageGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []SingleWantedStageGroupMetaData
}

type SingleWantedThemeConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []SingleWantedThemeConfigMetaData
}

type SkillTutorialStepDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SkillTutorialStepData
}

type SkillVideoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SkillVideoMetaData
}

type SLGBattlePointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SLGBattlePointMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SLGBattlePointMetaData
}

type SLGBattlePointMetaDataReader_StructKey struct {
	// Fields
	MapID   int32
	PointID int32
}

type SLGBigBossPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SLGBigBossPointMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SLGBigBossPointMetaData
}

type SLGBigBossPointMetaDataReader_StructKey struct {
	// Fields
	BossID         int32
	BossScheduleID int32
}

type SLGBroadCastMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGBroadCastMetaData
}

type SLGBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGBuffMetaData
}

type SLGBuildingMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGBuildingMetaData
}

type SLGCountrySetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []SLGCountrySetMetaData
}

type SLGRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SLGRankRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SLGRankRewardMetaData
}

type SLGRankRewardMetaDataReader_StructKey struct {
	// Fields
	Rank       int32
	ScheduleID int32
}

type SLGScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []SLGScheduleMetaData
}

type SLGScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGScoreRewardMetaData
}

type SLGShopBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGShopBuffMetaData
}

type SLGSingleBattleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []SLGSingleBattleMetaData
}

type SLGSmallBossPointMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGSmallBossPointMetaData
}

type SLGStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGStageMetaData
}

type SLGStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []SLGStoryMetaData
}

type SLGVoteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SLGVoteMetaData
}

type SLGVoteScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int16
	ItemOffsets []int32

	Data []SLGVoteScheduleMetaData
}

type SlotMachineBoxitemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SlotMachineBoxitemMetaData
}

type SlotMachineCombinationMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SlotMachineCombinationMetaData
}

type SlotMachineIconMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SlotMachineIconMetaData
}

type SlotMachineProgressRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SlotMachineProgressRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SlotMachineProgressRewardMetaData
}

type SlotMachineProgressRewardMetaDataReader_StructKey struct {
	// Fields
	PanelID          int32
	ProgressRewardID int32
}

type SlotMachineScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SlotMachineScheduleMetaData
}

type SpaceShipConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SpaceShipConfigMetaData
}

type SpareShareConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []SpareShareConfigMetaData
}

type SpecialAffixDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SpecialAffixDataMetaData
}

type SpecialAvatarDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SpecialAvatarDataMetaData
}

type SpecialConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SpecialConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SpecialConfigMetaData
}

type SpecialConfigMetaDataReader_StructKey struct {
	// Fields
	ThemeID        int32
	DefaultResPath StrWithPrefix16
}

type StageDamageStatisticsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDamageStatisticsMetaData
}

type StageDetailAbilityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailAbilityMetaData
}

type StageDetailAccumuatorGaugeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []StageDetailAccumuatorGaugeMetaData
}

type StageDetailEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailEffectMetaData
}

type StageDetailLayoutMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailLayoutMetaData
}

type StageDetailMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailMetaData
}

type StageDetailMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailMonsterMetaData
}

type StageDetailMonsterResistMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailMonsterResistMetaData
}

type StageDetailNatureMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailNatureMetaData
}

type StageDetailSkillTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDetailSkillTypeMetaData
}

type StageDialogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDialogMetaData
}

type StageDifficultySelectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageDifficultySelectMetaData
}

type StageDropItemDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []StageDropItemDataMetaData
}

type StageEffectCarryOnMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageEffectCarryOnMetaData
}

type StageEnhanceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageEnhanceMetaData
}

type StageLiveRecommendMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageLiveRecommendMetaData
}

type StageMaxScoreMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageMaxScoreMetaData
}

type StageMessageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []StageMessageMetaData
}

type StageMultiTeamMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageMultiTeamMetaData
}

type StageNeedMuteChallengePopupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageNeedMuteChallengePopupMetaData
}

type StagePhotoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StagePhotoMetaData
}

type StageQAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageQAvatarMetaData
}

type StageRandomEffectMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageRandomEffectMetaData
}

type StageRankConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []StageRankConfigMetaData
}

type StageRankGroupConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []StageRankGroupConfigMetaData
}

type StageRechallengeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageRechallengeMetaData
}

type StageResistanceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageResistanceMetaData
}

type StageRestrictExtendMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []StageRestrictExtendMetaData
}

type StageRestrictMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageRestrictMetaData
}

type StageScoreRewardDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageScoreRewardDisplayMetaData
}

type StageScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StageScoreRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []StageScoreRewardMetaData
}

type StageScoreRewardMetaDataReader_StructKey struct {
	// Fields
	MinScore uint32
	StageID  uint32
}

type StageStoryMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageStoryMapMetaData
}

type StageStoryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StageStoryMetaData
}

type StageTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []StageTypeMetaData
}

type StepMissionCompensateMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []StepMissionCompensateMetaData
}

type StigmataAffixMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataAffixMetaData
}

type StigmataBriefMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataBriefMetaData
}

type StigmataMemoirMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataMemoirMetaData
}

type StigmataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataMetaData
}

type StigmataPositionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []StigmataPositionMetaData
}

type StigmataRuneAffixMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataRuneAffixMetaData
}

type StigmataTagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataTagMetaData
}

type StigmataThemeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataThemeMetaData
}

type StigmataVirtualSetMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StigmataVirtualSetMetaData
}

type StorySweepGroupInfoReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StorySweepGroupInfo
}

type StorySweepMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []StorySweepMetaData
}

type SubMallTypeDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SubMallTypeDataMetaData
}

type SubPakDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SubPakDataMetaData
}

type SubPakGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SubPakGroupMetaData
}

type SubsBenefitDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SubsBenefitDataMetaData
}

type SubsDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SubsDataMetaData
}

type SubsidiaryEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SubsidiaryEventMetaData
}

type SubsWareDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []SubsWareDataMetaData
}

type SupportActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SupportActivityScheduleMetaData
}

type SupportAvatarInitialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SupportAvatarInitialMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SupportAvatarInitialMetaData
}

type SupportAvatarInitialMetaDataReader_StructKey struct {
	// Fields
	ActivityType uint16
	ActivityID   uint32
}

type SupportAvatarLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SupportAvatarLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SupportAvatarLevelMetaData
}

type SupportAvatarLevelMetaDataReader_StructKey struct {
	// Fields
	SupportAvatarID    uint16
	SupportAvatarLevel uint16
}

type SupportAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []SupportAvatarMetaData
}

type SupportConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SupportConfigMetaData
}

type SupportMissionConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []SupportMissionConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []SupportMissionConfigMetaData
}

type SupportMissionConfigMetaDataReader_StructKey struct {
	// Fields
	Index                  int32
	SupportMissionConfigID int32
}

type SurvivalWeaponShowMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []SurvivalWeaponShowMetaData
}

type TagListDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []TagListDataMetaData
}

type TeamAssaultBossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TeamAssaultBossMetaData
}

type TeamAssaultMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TeamAssaultMetaData
}

type TeamBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TeamBuffMetaData
}

type TextIDReplaceMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []Hash
	ItemOffsets []int32

	Data []TextIDReplaceMetaData
}

type TextMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []Hash
	ItemOffsets []int32

	Data []TextMapMetaData
}

type TextMapMultiLangMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []Hash
	ItemOffsets []int32

	Data []TextMapMetaData
}

type ThemeActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeActivityMetaData
}

type ThemeAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeAvatarMetaData
}

type ThemeBGMConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeBGMConfigMetaData
}

type ThemeBonusDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeBonusDataMetaData
}

type ThemeBuffDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeBuffData
}

type ThemeBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeBuffMetaData
}

type ThemeEquipMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeEquipMetaData
}

type ThemeEventImgMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ThemeEventImgMetaData
}

type ThemeEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []ThemeEventMetaData
}

type ThemeScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeScheduleMetaData
}

type ThemeSeasonMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeSeasonMetaData
}

type ThemeSpecialConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeSpecialConfigMetaData
}

type ThemeStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeStageMetaData
}

type ThemeStigmataGroupReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeStigmataGroup
}

type ThemeUIConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []ThemeUIConfigMetaData
}

type ThemeWeatherAtmosphereMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ThemeWeatherAtmosphereMetaData
}

type ThemeWeatherConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ThemeWeatherConfigMetaData
}

type ThemeWeatherGroupConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ThemeWeatherGroupConfigMetaData
}

type ThemeWeatherReplacePrefabMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []ThemeWeatherReplacePrefabMetaData
}

type ThreadCollectionFileDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []ThreadCollectionFileDataMetaData
}

type TileEntityDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileEntityDataMetaData
}

type TileMapInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileMapInfoMetaData
}

type TileMapMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileMapMetaData
}

type TileModelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileModelMetaData
}

type TileTerrainDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileTerrainDataMetaData
}

type TileValueKeyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TileValueKeyMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TileValueKeyMetaData
}

type TileValueKeyMetaDataReader_StructKey struct {
	// Fields
	Key   int32
	MapID int32
}

type TileValueModifierMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileValueModifierMetaData
}

type TileValueTriggerMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TileValueTriggerMetaData
}

type TimePocketMoneyDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TimePocketMoneyDataMetaData
}

type TouchBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TouchBuffMetaData
}

type TouchLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TouchLevelMetaData
}

type TouchMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TouchMetaData
}

type TowerBuffConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerBuffConfigMetaData
}

type TowerGrowBuffConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TowerGrowBuffConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TowerGrowBuffConfigMetaData
}

type TowerGrowBuffConfigMetaDataReader_StructKey struct {
	// Fields
	ActivityID  int32
	BuffLevelID int32
}

type TowerRaidBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerRaidBuffMetaData
}

type TowerRaidConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerRaidConfigMetaData
}

type TowerRaidDifficultyMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerRaidDifficultyMetaData
}

type TowerRaidScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerRaidScheduleMetaData
}

type TowerRaidStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerRaidStageMetaData
}

type TowerStageConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TowerStageConfigMetaData
}

type TradingCardActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TradingCardActivityScheduleMetaData
}

type TrainingLevelAchieveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TrainingLevelAchieveMetaData
}

type TrainingLevelDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TrainingLevelData
}

type TutorialDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TutorialData
}

type TutorialGraphicMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TutorialGraphicMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TutorialGraphicMetaData
}

type TutorialGraphicMetaDataReader_StructKey struct {
	// Fields
	StepId      uint8
	TutorialKey StrWithPrefix16
}

type TutorialRestrictMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TutorialRestrictMetaData
}

type TutorialStepDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TutorialStepData
}

type TutorialStepScrollerDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TutorialStepScrollerData
}

type TutorialVideoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TutorialVideoMetaData
}

type TutorialWeeklyScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TutorialWeeklyScheduleMetaData
}

type TVTAvatarConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TVTAvatarConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TVTAvatarConfigMetaData
}

type TVTAvatarConfigMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	DressID  int32
}

type TVTCardLevelMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TVTCardLevelMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TVTCardLevelMetaData
}

type TVTCardLevelMetaDataReader_StructKey struct {
	// Fields
	Level uint8
	ID    int32
}

type TVTCardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TVTCardMetaData
}

type TVTCardSlotMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []TVTCardSlotMetaData
}

type TVTConstValuesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []TVTConstValuesMetaData
}

type TVTDivisionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TVTDivisionMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TVTDivisionMetaData
}

type TVTDivisionMetaDataReader_StructKey struct {
	// Fields
	DivisionID uint32
	SeasonID   uint32
}

type TVTDivisionProtectUseMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []TVTDivisionProtectUseMetaData
}

type TVTInLevelInteractionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TVTInLevelInteractionMetaData
}

type TVTMaterialTransformMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TVTMaterialTransformMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TVTMaterialTransformMetaData
}

type TVTMaterialTransformMetaDataReader_StructKey struct {
	// Fields
	TransformCardID uint16
	MaterialID      int32
}

type TVTMessageConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []TVTMessageConfigMetaDataReader_StructKey
	ItemOffsets []int32

	Data []TVTMessageConfigMetaData
}

type TVTMessageConfigMetaDataReader_StructKey struct {
	// Fields
	HardLevel      uint32
	HardLevelGroup uint32
	MessageID      uint32
}

type TVTMessageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []TVTMessageMetaData
}

type TVTMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TVTMissionMetaData
}

type TVTRobotDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []TVTRobotDataMetaData
}

type TVTScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []TVTScheduleMetaData
}

type TVTSeasonMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []TVTSeasonMetaData
}

type TVTStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TVTStageMetaData
}

type TVTTutorialMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []TVTTutorialMetaData
}

type UIOverrideItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []UIOverrideItemMetaData
}

type UIOverrideScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []UIOverrideScheduleMetaData
}

type UltraEndlessBaseRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []UltraEndlessBaseRewardMetaData
}

type UltraEndlessBattleConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []UltraEndlessBattleConfigMetaData
}

type UltraEndlessBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []UltraEndlessBuffMetaData
}

type UltraEndlessDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []UltraEndlessDisplayMetaData
}

type UltraEndlessFloorMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []UltraEndlessFloorMetaDataReader_StructKey
	ItemOffsets []int32

	Data []UltraEndlessFloorMetaData
}

type UltraEndlessFloorMetaDataReader_StructKey struct {
	// Fields
	FloorID uint8
	StageID int32
}

type UltraEndlessGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []UltraEndlessGroupMetaData
}

type UltraEndlessRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []UltraEndlessRewardMetaDataReader_StructKey
	ItemOffsets []int32

	Data []UltraEndlessRewardMetaData
}

type UltraEndlessRewardMetaDataReader_StructKey struct {
	// Fields
	GroupLevel      uint8
	BeginScheduleId uint32
}

type UltraEndlessScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []UltraEndlessScheduleMetaData
}

type UltraEndlessSettleCupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []UltraEndlessSettleCupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []UltraEndlessSettleCupMetaData
}

type UltraEndlessSettleCupMetaDataReader_StructKey struct {
	// Fields
	GroupLevel       uint8
	PlayerNumInGroup uint8
}

type UltraEndlessSiteMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []UltraEndlessSiteMetaData
}

type UltraEndlessStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []UltraEndlessStageMetaData
}

type UltraEndlessTopRankShowMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []UltraEndlessTopRankShowMetaData
}

type UniqueGachaPoolMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []UniqueGachaPoolMetaDataReader_StructKey
	ItemOffsets []int32

	Data []UniqueGachaPoolMetaData
}

type UniqueGachaPoolMetaDataReader_StructKey struct {
	// Fields
	PoolId uint16
	ItemId uint32
}

type UniqueGachaReplaceItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []UniqueGachaReplaceItemMetaData
}

type UniqueGoodsItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []UniqueGoodsItemMetaData
}

type UniqueMonsterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []UniqueMonsterMetaData
}

type UnlockUIDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []UnlockUIData
}

type VibrateConfigsMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []StrWithPrefix16
	ItemOffsets []int32

	Data []VibrateConfigsMetaData
}

type VirtualAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualAvatarMetaData
}

type VirtualBuffMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualBuffMetaData
}

type VirtualCustomLevelDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []VirtualCustomLevelDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []VirtualCustomLevelDataMetaData
}

type VirtualCustomLevelDataMetaDataReader_StructKey struct {
	// Fields
	CustomID int32
	Level    int32
}

type VirtualCustomMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualCustomMetaData
}

type VirtualGachaChestMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualGachaChestMetaData
}

type VirtualGachaItemMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualGachaItemMetaData
}

type VirtualGroupConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualGroupConfigMetaData
}

type VirtualResourcesMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualResourcesMetaData
}

type VirtualRoleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualRoleMetaData
}

type VirtualStigmataDetailMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualStigmataDetailMetaData
}

type VirtualStigmataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualStigmataMetaData
}

type VirtualTrainStageDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []VirtualTrainStageDropMetaDataReader_StructKey
	ItemOffsets []int32

	Data []VirtualTrainStageDropMetaData
}

type VirtualTrainStageDropMetaDataReader_StructKey struct {
	// Fields
	EndFloor int16
	StageId  int32
}

type VirtualWeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []VirtualWeaponMetaData
}

type WarehouseRequireDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WarehouseRequireData
}

type WaveRushBuffDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WaveRushBuffDataMetaData
}

type WaveRushBuffLevelGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []WaveRushBuffLevelGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []WaveRushBuffLevelGroupMetaData
}

type WaveRushBuffLevelGroupMetaDataReader_StructKey struct {
	// Fields
	BuffGroupConfigID int32
	BuffID            int32
	Level             int32
}

type WaveRushFuncOpenMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []WaveRushFuncOpenMetaData
}

type WaveRushScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WaveRushScheduleMetaData
}

type WaveRushStageGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []WaveRushStageGroupMetaDataReader_StructKey
	ItemOffsets []int32

	Data []WaveRushStageGroupMetaData
}

type WaveRushStageGroupMetaDataReader_StructKey struct {
	// Fields
	StageGroupConfigID int32
	StageID            int32
}

type WaveRushStageMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WaveRushStageMetaData
}

type WaveRushStageScoreDropMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []WaveRushStageScoreDropMetaDataReader_StructKey
	ItemOffsets []int32

	Data []WaveRushStageScoreDropMetaData
}

type WaveRushStageScoreDropMetaDataReader_StructKey struct {
	// Fields
	NeedMinScore int32
	StageID      int32
}

type WaveRushStageWaveMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WaveRushStageWaveMetaData
}

type WaveRushUIConfigMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WaveRushUIConfigMetaData
}

type WeaponBaseTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeaponBaseTypeMetaData
}

type WeaponCustomDisplayMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeaponCustomDisplayMetaData
}

type WeaponLowResOverrideMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []WeaponLowResOverrideMetaDataReader_StructKey
	ItemOffsets []int32

	Data []WeaponLowResOverrideMetaData
}

type WeaponLowResOverrideMetaDataReader_StructKey struct {
	// Fields
	AvatarID uint16
	WeaponID int32
}

type WeaponMainIDMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeaponMainIDMetaData
}

type WeaponMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeaponMetaData
}

type WeaponShowOverrideMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeaponShowOverrideMetaData
}

type WeaponTagMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeaponTagMetaData
}

type WebLinkAvatarMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []WebLinkAvatarMetaDataReader_StructKey
	ItemOffsets []int32

	Data []WebLinkAvatarMetaData
}

type WebLinkAvatarMetaDataReader_StructKey struct {
	// Fields
	IsAvatarArtifact bool
	AvatarID         int32
}

type WebLinkElfMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WebLinkElfMetaData
}

type WebLinkEquipMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WebLinkEquipMetaData
}

type WebLinkExBossMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WebLinkExBossMetaData
}

type WeekDayActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeekDayActivityMetaData
}

type WeekDayActivityScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeekDayActivityScheduleMetaData
}

type WeeklyReportMedalMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint32
	ItemOffsets []int32

	Data []WeeklyReportMedalMetaData
}

type WeeklyReportMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WeeklyReportMetaData
}

type WeeklyReportSubPageDataMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []WeeklyReportSubPageDataMetaDataReader_StructKey
	ItemOffsets []int32

	Data []WeeklyReportSubPageDataMetaData
}

type WeeklyReportSubPageDataMetaDataReader_StructKey struct {
	// Fields
	FatherID int32
	ID       int32
}

type WelfareMissionMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WelfareMissionMetaData
}

type WikiActivityCatalogMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []WikiActivityCatalogMetaData
}

type WikiActivityChapterMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []WikiActivityChapterMetaData
}

type WikiActivityEventMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint16
	ItemOffsets []int32

	Data []WikiActivityEventMetaData
}

type WikiActivityMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []uint8
	ItemOffsets []int32

	Data []WikiActivityMetaData
}

type WikiCollectionRankMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WikiCollectionRankMetaData
}

type WikiTypeMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WikiTypeMetaData
}

type WorldBossGroupMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldBossGroupMetaData
}

type WorldBossRankRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldBossRankRewardMetaData
}

type WorldBossScoreRewardMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldBossScoreRewardMetaData
}

type WorldBossSkillMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldBossSkillMetaData
}

type WorldMapActivityInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldMapActivityInfoMetaData
}

type WorldMapEntryMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldMapEntryMetaData
}

type WorldMapScheduleMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldMapScheduleMetaData
}

type WWorldMapCoverInfoMetaDataReader struct {
	Filesize    uint32
	EntryCount  uint32 `bin:"sizeof=Keys,ItemOffsets,Data"`
	Keys        []int32
	ItemOffsets []int32

	Data []WorldMapCoverInfoMetaData
}
