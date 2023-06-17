package dump

type ArtifactSkillNameOverrideMetaData struct {
	// Properties
	OverrideName AddrTo[Hash]
}

type AvatarArtifactLevelMetaData struct {
	// Properties
	Level             int32
	UnlockAvatarStar  int32
	UnlockAvatarLevel int32
	UnlockMaterial    AddrTo[[]AvatarArtifactLevelMetaData_UpgradeMaterialItem]
}

type AvatarArtifactLevelMetaData_UpgradeMaterialItem struct {
	// Fields
	Amount int32
	ItemID int32
}

type AvatarArtifactMetaData struct {
	// Properties
	ShowPlayLevel       int32
	UnlockAvatarLevel   int32
	ArtifactName        AddrTo[Hash]
	ArtifactDes         AddrTo[Hash]
	ArtifactPrefab      AddrTo[StrWithPrefix16]
	LevelList           AddrTo[[]int32]
	ReplaceSkillList    AddrTo[[]AvatarArtifactMetaData_ReplaceSkillPair]
	DressBanner         AddrTo[StrWithPrefix16]
	DressBannerDes1     AddrTo[Hash]
	DressBannerDes2     AddrTo[Hash]
	BannerAvatarImgPath AddrTo[StrWithPrefix16]
	BannerTitleImgPath  AddrTo[StrWithPrefix16]
	ShowSkillList       AddrTo[[]AvatarArtifactMetaData_LockedDisplaySkillInfo]
	TagUnlockList       AddrTo[[]int32]
}

type AvatarArtifactMetaData_ReplaceSkillPair struct {
	// Fields
	NewSkillID int32
	OldSkillID int32
}

type AvatarArtifactMetaData_LockedDisplaySkillInfo struct {
	// Fields
	SkillID int32

	// Objects
	SkillDescriptionTextID Hash
}

type AvatarCardMetaData struct {
	// Properties
	Rarity                     int32
	MaxRarity                  int32
	Cost                       int32
	MaxLv                      int32
	SellPriceBase              float32
	SellPriceAdd               float32
	GearExpProvideBase         float32
	GearExpPorvideAdd          float32
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	CharacterExpProvide        float32
	GachaDisplayRarity         int32
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
}

type AvatarFragmentMetaData struct {
	// Properties
	Rarity                     int32
	MaxRarity                  int32
	Cost                       int32
	MaxLv                      int32
	SellPriceBase              float32
	SellPriceAdd               float32
	GearExpProvideBase         float32
	GearExpPorvideAdd          float32
	BaseType                   int32
	SortID                     uint8
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	CharacterExpProvide        float32
	QuantityLimit              int32
	LinkIDList                 AddrTo[[]uint32]
	ReturnMaterialID           int32
	ReturnNum                  int32
	TagType                    uint8
	GachaDisplayRarity         int32
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
}

type AvatarLevelMetaData struct {
	// Properties
	Exp              int32
	Cost             int32
	AvatarAssistConf float32
	SubSkillScoin    int32
}

type AvatarMetaData struct {
	// Properties
	ClassID                int32
	RoleID                 int32
	AvatarType             uint8
	FullName               AddrTo[Hash]
	ShortName              AddrTo[Hash]
	RomaName               AddrTo[Hash]
	Desc                   AddrTo[Hash]
	AvatarRegistryKey      AddrTo[StrWithPrefix16]
	WeaponBaseTypeList     AddrTo[[]int32]
	UnlockStar             int32
	SkillList              AddrTo[[]int32]
	Attribute              int32
	InitialWeapon          int32
	AvatarCardID           int32
	AvatarFragmentID       int32
	ArtifactFragmentID     int32
	UltraSkillID           int32
	CaptainSkillID         int32
	SKL01SP                float32
	SKL01SPNeed            float32
	SKL01Charges           int32
	SKL01CD                float32
	SKL02SP                float32
	SKL02SPNeed            float32
	SKL02Charges           int32
	SKL02CD                float32
	SKL03SP                float32
	SKL03SPNeed            float32
	SKL03Charges           int32
	SKL03CD                float32
	SKL02ArtifactCD        float32
	SKL02ArtifactSP        float32
	SKL02ArtifactSPNeed    float32
	BaseAvatarID           int32
	FirstName              AddrTo[Hash]
	LastName               AddrTo[Hash]
	EnFirstName            AddrTo[Hash]
	EnLastName             AddrTo[Hash]
	UISelectVoice          AddrTo[StrWithPrefix16]
	UILevelUpVoice         AddrTo[StrWithPrefix16]
	DA_Name                AddrTo[StrWithPrefix16]
	DA_Type                AddrTo[StrWithPrefix16]
	ArtifactID             int32
	IsEasterner            bool
	FaceAnimationGroupName AddrTo[StrWithPrefix16]
	AvatarEffects          AddrTo[[]int32]
	TagUnlockList          AddrTo[[]int32]
	DefaultDressId         int32
	AvatarStarUpType       uint8
	AvatarStarSourceID     AddrTo[[]int32]
	IsCollaboration        bool
	StarUpBG               AddrTo[StrWithPrefix16]
}

type AvatarModelMetaData struct {
	// Properties
	ModelID          int32
	AvatarID         uint16
	AvatarLV         uint8
	AvatarStar       uint8
	AvatarSubSkills  AddrTo[[]AvatarModelMetaData_SkillInfoItem]
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
}

type AvatarModelMetaData_SkillInfoItem struct {
	// Fields
	SkillID int32
	SkillLV int32
}

type AvatarSampleMetaData struct {
	// Properties
	ModelID         int32
	AvatarID        int32
	MaterialID      int32
	Type            int32
	ForceLevel      int32
	StageTypeList   AddrTo[[]uint8]
	StageTypeList2  AddrTo[[]uint8]
	StageTypeEffect int32
	StageIDList     AddrTo[[]int32]
	SampleDesc      AddrTo[StrWithPrefix16]
}

type AvatarSkillMetaData struct {
	// Properties
	Name            AddrTo[Hash]
	Info            AddrTo[Hash]
	ShowOrder       int32
	UnlockLv        int32
	UnlockStar      int32
	SkillStep       AddrTo[Hash]
	IconPath        AddrTo[StrWithPrefix16]
	IconPathInLevel AddrTo[StrWithPrefix16]
	ButtonName      AddrTo[StrWithPrefix16]
	ParamBase_1     float32
	ParamLogic_1    uint16
	ParamSubID_1    int32
	ParamSubIndex_1 int32
	ParamBase_2     float32
	ParamLogic_2    uint16
	ParamSubID_2    int32
	ParamSubIndex_2 int32
	ParamBase_3     float32
	ParamLogic_3    uint16
	ParamSubID_3    int32
	ParamSubIndex_3 int32
	CanTry          int32
	TagList         AddrTo[[]AvatarSkillMetaData_Tag]
	UnlockItemList  AddrTo[[]AvatarArtifactLevelMetaData_UpgradeMaterialItem]
}

type AvatarSkillMetaData_Tag struct {
	// Fields
	Strength uint8
	TagID    int32

	// Objects
	TagComment Hash
}

type AvatarStarMetaData struct {
	// Properties
	IconPath              AddrTo[StrWithPrefix16]
	IconPathInLevel       AddrTo[StrWithPrefix16]
	FigurePath            AddrTo[StrWithPrefix16]
	HpBase                float32
	HpAdd                 float32
	SpBase                float32
	SpAdd                 float32
	AtkBase               float32
	AtkAdd                float32
	DfsBase               float32
	DfsAdd                float32
	CrtBase               float32
	CrtAdd                float32
	ChibiIconPath         AddrTo[StrWithPrefix16]
	SkillID               AddrTo[[]int32]
	SkillUnlockType       AddrTo[[]int32]
	SkillNodeName         AddrTo[Hash]
	SkillNodeDesc         AddrTo[Hash]
	SkillUnlockSketchList AddrTo[[]Hash]
	SkillUnlockDescList   AddrTo[[]Hash]
}

type AvatarStarTypeMetaData struct {
	// Properties
	Upgrade  uint32
	IconPath AddrTo[StrWithPrefix16]
}

type AvatarStarUpSubSkillData struct {
	// Properties
	ParamBase_1 float32
	ParamBase_2 float32
	ParamBase_3 float32
	ParamAdd_1  float32
	ParamAdd_2  float32
	ParamAdd_3  float32
}

type AvatarSubSkillMetaData struct {
	// Properties
	Name                    AddrTo[Hash]
	Info                    AddrTo[Hash]
	Brief                   AddrTo[Hash]
	ShowOrder               uint8
	SkillId                 int32
	IgnoreLeader            bool
	IconPath                AddrTo[StrWithPrefix16]
	UnlockStar              uint8
	UnlockSubStar           uint8
	UnlockLv                uint8
	UnlockLvAdd             uint8
	MaxLv                   uint8
	UpLevelSubStarNeedList  AddrTo[[]AvatarSubSkillMetaData_UpLevelStarNeed]
	ScoinCalc               bool
	UnlockScoin             int32
	ScoinLvAdd              uint8
	ItemType                uint8
	SkillToggle             bool
	ParamBase_1             float32
	ParamAdd_1              float32
	ParamBase_2             float32
	ParamAdd_2              float32
	ParamBase_3             float32
	ParamAdd_3              float32
	CanTry                  bool
	ArtifactSkillID         uint16
	UpLevelArtifactNeedList AddrTo[[]AvatarSubSkillMetaData_ArtifactSubSkillLevelUpCondition]
	TagList                 AddrTo[[]AvatarSkillMetaData_Tag]
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

type AvatarTagUnLockMetaData struct {
	// Properties
	TagID                int32
	TagSecondDesc        AddrTo[Hash]
	UnlockSkillIDList    AddrTo[[]int32]
	UnlockSubSkillIDList AddrTo[[]int32]
}

type CollaborationStigmataMetaData struct {
	// Properties
	OptionalStigmataList      AddrTo[[]int32]
	CollaborationStigmataList AddrTo[[]int32]
	StigmataSetName           AddrTo[Hash]
	StigmataSetIcon           AddrTo[StrWithPrefix16]
	OptionalStigmataTextmap   AddrTo[Hash]
	OptionalStigmataShowTime  AddrTo[RemoteTime]
}

type CustomHeadMetaData struct {
	// Properties
	IndexID         int32
	Page            int32
	HeadTitle       AddrTo[StrWithPrefix16]
	HeadDescription AddrTo[StrWithPrefix16]
	IconPath        AddrTo[StrWithPrefix16]
	ImagePath       AddrTo[StrWithPrefix16]
	Show            bool
	HeadType        int32
	HeadParaInt     int32
	TimeType        int32
	During          int32
	EndTime         AddrTo[StrWithPrefix16]
	BgColorPath     AddrTo[StrWithPrefix16]
	Rarity          int32
}

type DLCAvatarCardMetaData struct {
	// Properties
	Rarity             int32
	DisplayTitle       AddrTo[Hash]
	DisplayDescription AddrTo[Hash]
	IconPath           AddrTo[StrWithPrefix16]
	ImagePath          AddrTo[StrWithPrefix16]
}

type DLCAvatarMetaData struct {
	// Properties
	Index             int32
	StoryBGImg        AddrTo[StrWithPrefix16]
	StoryDesc         AddrTo[Hash]
	UnlockHintStoryID int32
	RotationFront     float32
	RotationSide      float32
	TalentTattooImg   AddrTo[StrWithPrefix16]
}

type DormitoryFurnitureMetaData struct {
	// Properties
	Rarity                  uint8
	FurnitureTag            AddrTo[[]uint8]
	Type                    uint8
	EditType                uint8
	InteractiveType         uint8
	InitialUnLock           uint8
	SuitId                  uint8
	Size                    AddrTo[[]uint8]
	ManualScale             float32
	ComfortType             uint8
	ComfortValue            uint8
	HcoinCost               uint16
	HcoinUnlock             uint16
	CostItemList            AddrTo[[]DormitoryFurnitureMetaData_Material]
	UnlockItemList          AddrTo[[]DormitoryFurnitureMetaData_Material]
	UnlockAdd               bool
	FurnitureModPath        AddrTo[StrWithPrefix16]
	FurnitureNameText       AddrTo[Hash]
	FurnitureIconPath       AddrTo[StrWithPrefix16]
	FurnitureDescText       AddrTo[Hash]
	FurnitureSourceDescText AddrTo[Hash]
	KeepLimit               uint8
	DeleteMaterialList      AddrTo[[]DormitoryFurnitureMetaData_Material]
	VerandaPath             AddrTo[StrWithPrefix16]
	CollectionType          uint8
}

type DormitoryFurnitureMetaData_Material struct {
	// Fields
	MaterialID     int32
	MaterialNumber int32
}

type DressMetaData struct {
	// Properties
	Name                       AddrTo[Hash]
	Rarity                     int32
	AvatarIDList               AddrTo[[]int32]
	RoleID                     int32
	IconPath                   AddrTo[StrWithPrefix16]
	DressType                  int32
	GetWay                     AddrTo[Hash]
	Desc                       AddrTo[Hash]
	Coin                       int32
	DressResource              AddrTo[StrWithPrefix16]
	CardPath                   AddrTo[StrWithPrefix16]
	TachiePath                 AddrTo[StrWithPrefix16]
	AvatarIconPath             AddrTo[StrWithPrefix16]
	AvatarIconSidePath         AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	LinkIDList                 AddrTo[[]uint32]
	AvatarGachaFigure          AddrTo[StrWithPrefix16]
	Show                       int32
	PreviewStageID             int32
	DressTagList               AddrTo[[]int32]
	RechargeShowName           AddrTo[Hash]
	RechargeShowAvatarName     AddrTo[Hash]
	RechargeShowPicPath        AddrTo[StrWithPrefix16]
	ArtifactID                 int32
	MVPVoicePattern            AddrTo[StrWithPrefix16]
	UISelectVoice              AddrTo[StrWithPrefix16]
	UILevelupVoice             AddrTo[StrWithPrefix16]
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
	IsCollaboration            bool
}

type ElfBreakMetaData struct {
	// Properties
	MaxLevel              int32
	BreakCostMaterialList AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
	NeedPlayerLevel       int32
}

type ElfBreakMetaData_ElfBreakMaterial struct {
	// Fields
	MaterialID int32
	Number     int32
}

type ElfCaptainSkillMetaData struct {
	// Properties
	Restricts      AddrTo[[]int32]
	RestrictExtend int32
}

type ElfCardMetaData struct {
	// Properties
	ElfID                      int32
	Rarity                     int32
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	GachaDisplayRarity         int32
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
}

type ElfFragmentMetaData struct {
	// Properties
	ElfID                      int32
	Rarity                     int32
	BaseType                   int32
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	QuantityLimit              int32
	LinkIDList                 AddrTo[[]uint32]
	ReturnMaterialID           int32
	ReturnNum                  int32
	ReturnMinStar              int32
	TagType                    uint8
	GachaDisplayRarity         int32
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
}

type ElfGalEventMetaData struct {
	// Properties
	ElfID                  int32
	TriggerName            AddrTo[StrWithPrefix16]
	Audio                  AddrTo[StrWithPrefix16]
	AudioDelayTime         int32
	SelfCD                 int32
	Priority               int32
	TriggerEffectList      AddrTo[[]ElfGalEventMetaData_ElfUIEffectPattern]
	FaceAnimationKey       AddrTo[StrWithPrefix16]
	FaceAnimationDelayTime int32
}

type ElfGalEventMetaData_ElfUIEffectPattern struct {
	// Objects
	AttachPoint StrWithPrefix16
	EffectPath  StrWithPrefix16
	BeginTime   float32
}

type ElfMetaData struct {
	// Properties
	Index                  int32
	FullName               AddrTo[Hash]
	ElfRegistryKey         AddrTo[StrWithPrefix16]
	ElfUIModelPath         AddrTo[StrWithPrefix16]
	ElfUIModelPosition     AddrTo[ElfMetaData_VectorFloat3]
	ElfUIModelRotation     AddrTo[ElfMetaData_VectorFloat3]
	IconPath               AddrTo[StrWithPrefix16]
	ElfChibiIcon           AddrTo[StrWithPrefix16]
	PortraitCommonIcon     AddrTo[StrWithPrefix16]
	StoryBGImg             AddrTo[StrWithPrefix16]
	PresentBGImg           AddrTo[StrWithPrefix16]
	AIName                 AddrTo[StrWithPrefix16]
	IsFlight               int32
	UnlockStar             int32
	Rarity                 int32
	ElfCardID              int32
	ElfFragmentID          int32
	CarryPlayerLevel       int32
	UISoundbankName        AddrTo[StrWithPrefix16]
	StoryDesc              AddrTo[Hash]
	RGBCGID                int32
	AlphaCGID              int32
	UltraSkillCD           int32
	UltraSkillSPCost       int32
	AttackSkillCD          float32
	SkillTabBGImg          AddrTo[StrWithPrefix16]
	TalentTabBGImg         AddrTo[StrWithPrefix16]
	SelectAudio            AddrTo[StrWithPrefix16]
	LevelUpAudio           AddrTo[StrWithPrefix16]
	CombatPointRarity      float32
	TagList                AddrTo[[]ElfMetaData_Tag]
	FeatureBrief           AddrTo[Hash]
	TrialStageActivityList AddrTo[[]int32]
	CaptainSkillIDs        AddrTo[[]int32]
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

	// Objects
	TagComment Hash
}

type ElfRarityMetaData struct {
	// Properties
	RarityImg AddrTo[StrWithPrefix16]
}

type ElfSkillMetaData struct {
	// Properties
	ElfID                   int32
	Name                    AddrTo[Hash]
	Info                    AddrTo[Hash]
	SkillTypeTag            AddrTo[Hash]
	MaxLv                   int32
	SkillType               int32
	IconPath                AddrTo[StrWithPrefix16]
	IconType                int32
	UnlockStar              int32
	UIPointRow              int32
	UIPointColumn           int32
	TagList                 AddrTo[[]ElfMetaData_Tag]
	IconSpecial             int32
	AbilityParamBase_1      float32
	AbilityParamAdd_1       float32
	AbilityParamBase_2      float32
	AbilityParamAdd_2       float32
	AbilityParamBase_3      float32
	AbilityParamAdd_3       float32
	HasNoRestrictionAbility int32
}

type ElfSkillTreeMetaData struct {
	// Properties
	LevelUpPreSkill     AddrTo[[]ElfSkillTreeMetaData_ElfPreSkill]
	LevelUpStar         int32
	NeedElfLevel        int32
	LevelUpMaterialList AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
}

type ElfSkillTreeMetaData_ElfPreSkill struct {
	// Fields
	Lv      int32
	SkillID int32
}

type ElfStarMetaData struct {
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
	// Properties
	ElfId       int32
	StoryTitle  AddrTo[Hash]
	StoryDetail AddrTo[Hash]
	Audio       AddrTo[StrWithPrefix16]
	SelfCD      float32
}

type EntryThemeItemDataMetaData struct {
	// Properties
	UnlockType   uint8
	UnlockPartID int32
	Rarity       int32
	ItemName     AddrTo[StrWithPrefix16]
	ItemDesc     AddrTo[StrWithPrefix16]
	IconPath     AddrTo[StrWithPrefix16]
	ImagePath    AddrTo[StrWithPrefix16]
}

type EntryThemeSampleDataMetaData struct {
	// Properties
	SpaceShipConfigID int32
	Rarity            uint8
	IconPath          AddrTo[StrWithPrefix16]
	ImagePath         AddrTo[StrWithPrefix16]
	SampleName        AddrTo[Hash]
	SampleDesc        AddrTo[Hash]
}

type EquipFragMetaData struct {
	// Properties
	Cost      int32
	Equipment int32
}

type EquipmentSetMetaData struct {
	// Properties
	SetName         AddrTo[Hash]
	SetDesc         AddrTo[Hash]
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
}

type EquipSkillMetaData struct {
	// Properties
	SkillName       AddrTo[Hash]
	SkillDisplay    AddrTo[Hash]
	SkillIconPath   AddrTo[StrWithPrefix16]
	SkillCD         int32
	SPCost          float32
	SPNeed          float32
	MaxChargesCount int32
	TagList         AddrTo[[]ElfMetaData_Tag]
}

type FrameDataMetaData struct {
	// Properties
	Name        AddrTo[Hash]
	Rarity      int32
	Desc        AddrTo[Hash]
	IconPath    AddrTo[StrWithPrefix16]
	ImagePath   AddrTo[StrWithPrefix16]
	IsShow      bool
	UIShowOrder int32
	Type        uint8
}

type GrandKeyBuffMetaData struct {
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
	BuffNameText        AddrTo[Hash]
	BuffDescStrId       AddrTo[Hash]
	BuffIcon            AddrTo[StrWithPrefix16]
	BuffTypeIcon        AddrTo[StrWithPrefix16]
}

type GrandKeySkillLimitMetaData struct {
	// Properties
	Minlevel         int32
	GrandKeySkillNum int32
	SynchronLimit    int32
}

type HalfBalanceModeAttrMetaData struct {
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
	// Properties
	ExceptFunctionId AddrTo[[]int32]
	AttrPoolId       AddrTo[[]HalfBalanceModeDataMetaData_AvatarPoolKV]
	DescInfo         AddrTo[Hash]
}

type HalfBalanceModeDataMetaData_AvatarPoolKV struct {
	// Fields
	AvatarId int32
	PoolId   int32
}

type ItemMetaData struct {
	// Properties
	Rarity                     uint8
	MaxRarity                  uint8
	Cost                       uint8
	MaxLv                      uint8
	SellPriceBase_1            int32
	SellPriceAdd               uint16
	ServantExpProvide          uint16
	GearExpProvideBase         uint16
	GearExpPorvideAdd          uint16
	ItemType                   AddrTo[StrWithPrefix16]
	UseID                      int32
	BaseType                   uint8
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	CharacterExpProvide        uint16
	LinkIDList                 AddrTo[[]uint32]
	ShopUseList                AddrTo[[]StrWithPrefix16]
	DisplayBGDescription       AddrTo[Hash]
	QuantityLimit              int32
	SortID                     uint8
	AffixTrainType             uint8
	AffixRandomValueIncress    uint8
	QuickBuyType               uint8
	ShopType                   uint8
	IdShopGoods                int32
	QuickBuyConfirm            bool
	HideInInventory            bool
	HideNumInTips              bool
	SellPriceID_1              uint8
	CostVitality               uint16
	EnableQuickSell            bool
	AdvSellBonusNum            uint16
	TagType                    uint8
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
	AlwaysShowPopUp            bool
}

type MaterialDeleteData struct {
	// Properties
	CallbackTime     AddrTo[RemoteTime]
	ReturnList       AddrTo[StrWithPrefix16]
	ReturnMaterialID int32
	ReturnNum        float32
	RoundType        int32
}

type MechMetaData struct {
	// Properties
	MonsterName           AddrTo[StrWithPrefix16]
	MonsterTypeName       AddrTo[StrWithPrefix16]
	Name_TextMap          AddrTo[Hash]
	Type                  int32
	SubType               int32
	RandomType_TextMap    AddrTo[Hash]
	CanRide               int32
	Rarity                int32
	IconPath              AddrTo[StrWithPrefix16]
	TowerStandby_IconPath AddrTo[StrWithPrefix16]
	HPGrade               AddrTo[StrWithPrefix16]
	AttackGrade           AddrTo[StrWithPrefix16]
	DefenceGrade          AddrTo[StrWithPrefix16]
	CriticalGrade         AddrTo[StrWithPrefix16]
	SPGrade               AddrTo[StrWithPrefix16]
	HP                    int32
	Attack                int32
	Defence               int32
	Critical              int32
	SP                    int32
	HP_Tower              int32
	Attack_Tower          int32
	Defence_Tower         int32
	Critical_Tower        int32
	EffectParam           int32
	DisjoinScoinCost      int32
	DisjoinAddMaterial    AddrTo[[]MechMetaData_DisjoinOutputItem]
	Story                 AddrTo[Hash]
	BestWeather           AddrTo[[]int32]
	ModeRenderYOffset     float32
	ModeRenderScale       float32
}

type MechMetaData_DisjoinOutputItem struct {
	// Fields
	ID  int32
	Num int32
}

type MedalMetaData struct {
	// Properties
	Rarity            int32
	TimeType          int32
	During            int32
	IndexID           int32
	MedalTitle        AddrTo[Hash]
	MedalDescription  AddrTo[Hash]
	MedalGet          AddrTo[Hash]
	IconPath          AddrTo[StrWithPrefix16]
	ImagePath         AddrTo[StrWithPrefix16]
	SmallPath         AddrTo[StrWithPrefix16]
	NumPositionOffset AddrTo[[]float32]
	HP                int32
	SP                int32
	Attack            int32
	Defence           int32
	Critical          int32
	Prop1ID           int32
	Prop1param1       float32
	Prop1param2       float32
	Prop1param3       float32
	Prop2ID           int32
	Prop2param1       float32
	Prop2param2       float32
	Prop2param3       float32
	Prop3ID           int32
	Prop3param1       float32
	Prop3param2       float32
	Prop3param3       float32
	Show              int32
	EndTime           AddrTo[StrWithPrefix16]
	ShowMission       int32
	ShowTime          AddrTo[RemoteTime]
	IsCollaboration   bool
}

type PhonePendantMetaData struct {
	// Properties
	Rarity                int32
	TimeType              int32
	TimeDuring            int32
	ScoinReward           int32
	PendantModPath        AddrTo[StrWithPrefix16]
	PendantIconPath       AddrTo[StrWithPrefix16]
	PendantImagePath      AddrTo[StrWithPrefix16]
	PendantNameText       AddrTo[Hash]
	PendantDescText       AddrTo[Hash]
	PendantSourceDescText AddrTo[Hash]
	PendantUnlockDescText AddrTo[Hash]
}

type ScDLCAvatarMetaData struct {
	// Properties
	IsFake            bool
	AvatarRegistryKey AddrTo[StrWithPrefix16]
	FullName          AddrTo[Hash]
	Desc              AddrTo[Hash]
	IconPath          AddrTo[StrWithPrefix16]
	ImagePath         AddrTo[StrWithPrefix16]
	HpBase            int32
	SpBase            int32
	AtkBase           int32
	DfsBase           int32
	TalentPoint       int32
	NPCID             int32
	AvatarHead        AddrTo[StrWithPrefix16]
	UIOffset          AddrTo[[]float32]
}

type ScDLCCombineTalentMetaData struct {
	// Properties
	CombineTalentID     int32
	CombineTalentSource AddrTo[[]int32]
	SubIcon             AddrTo[StrWithPrefix16]
	SubTalentTitle      AddrTo[Hash]
	SubTalentDesc       AddrTo[Hash]
}

type ScDLCTalentDataMetaData struct {
	// Properties
	AvatarID       int32
	TalentType     int32
	ParentTalent   int32
	UniqueTag      int32
	ForceLock      bool
	OWStoryID      int32
	UIPointName    AddrTo[StrWithPrefix16]
	UILineName     AddrTo[[]StrWithPrefix16]
	TalentName     AddrTo[Hash]
	IconPath       AddrTo[StrWithPrefix16]
	CGID           AddrTo[[]int32]
	IsSupport      bool
	SupportDerived AddrTo[[]int32]
}

type ScDLCTalentLevelMetaData struct {
	// Properties
	TalentDesc         AddrTo[Hash]
	NextLevelDesc      AddrTo[Hash]
	SubTalentLevelDesc AddrTo[[]StrWithPrefix16]
	PreTalent          AddrTo[[]ScDLCTalentLevelMetaData_LevelUpTalentRequire]
	FeverLevelLimit    int32
	LevelUpCost        AddrTo[[]ScDLCTalentLevelMetaData_LevelUpMatRequire]
	AbilityParamBase_1 float32
	AbilityParamBase_2 float32
	AbilityParamBase_3 float32
	IsCoopTalent       bool
	CoopAbilityName    AddrTo[StrWithPrefix16]
	CoopAbilityParam   AddrTo[[]StrWithPrefix16]
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
	// Properties
	ResPath      AddrTo[StrWithPrefix16]
	IsSkyBoxShow bool
}

type SpecialAvatarDataMetaData struct {
	// Properties
	SpecialAvatarPrefab AddrTo[StrWithPrefix16]
}

type StageRestrictExtendMetaData struct {
	// Properties
	AvatartNumMax    uint8
	RequestAvatarNum uint8
	ForceSelect      bool
	RequestType1     uint8
	RequestParam1    AddrTo[[]int32]
	RequsetTitle1    AddrTo[Hash]
	RequestDesc1     AddrTo[Hash]
	RequestPosName1  AddrTo[Hash]
	RequestType2     uint8
	RequestParam2    AddrTo[[]int32]
	RequestTitle_2   AddrTo[Hash]
	RequestDesc2     AddrTo[Hash]
	RequestPosName2  AddrTo[Hash]
	RequestType3     uint8
	RequestParam3    AddrTo[[]int32]
	RequestTitle3    AddrTo[Hash]
	RequestDesc3     AddrTo[Hash]
	RequestPosName3  AddrTo[Hash]
}

type StageRestrictMetaData struct {
	// Properties
	DisplayNameID     AddrTo[Hash]
	RequestType1      int32
	Argument1Type1    int32
	Argument2Type1    int32
	ArgumentListType1 AddrTo[[]int32]
	RequestType2      int32
	Argument1Type2    int32
	Argument2Type2    int32
	ArgumentListType2 AddrTo[[]int32]
	RequestType3      int32
	Argument1Type3    int32
	Argument2Type3    int32
	ArgumentListType3 AddrTo[[]int32]
}

type StigmataMetaData struct {
	// Properties
	Rarity                     uint8
	MaxRarity                  uint8
	SubRarity                  uint8
	SubMaxRarity               uint8
	Cost                       uint8
	PowerType                  uint8
	MaxLv                      uint8
	ExpType                    uint8
	SellPriceBase              float32
	SellPriceAdd               float32
	GearExpProvideBase         float32
	GearExpPorvideAdd          float32
	BaseType                   uint8
	LabelPath                  AddrTo[StrWithPrefix16]
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	DisplayNumber              int32
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	HPBase                     float32
	HPAdd                      float32
	SPBase                     float32
	SPAdd                      float32
	AttackBase                 float32
	AttackAdd                  float32
	DefenceBase                float32
	DefenceAdd                 float32
	CriticalBase               float32
	CriticalAdd                float32
	DurabilityMax              float32
	EvoMaterial                AddrTo[[]MechMetaData_DisjoinOutputItem]
	EvoID                      int32
	Prop1ID                    int32
	Prop1Param1                float32
	Prop1Param2                float32
	Prop1Param3                float32
	Prop1Param1Add             float32
	Prop1Param2Add             float32
	Prop1Param3Add             float32
	Prop2ID                    int32
	Prop2Param1                float32
	Prop2Param2                float32
	Prop2Param3                float32
	Prop2Param1Add             float32
	Prop2Param2Add             float32
	Prop2Param3Add             float32
	Prop3ID                    int32
	Prop3Param1                float32
	Prop3Param2                float32
	Prop3Param3                float32
	Prop3Param1Add             float32
	Prop3Param2Add             float32
	Prop3Param3Add             float32
	Protect                    bool
	SetID                      int32
	SmallIcon                  AddrTo[StrWithPrefix16]
	TattooPath                 AddrTo[StrWithPrefix16]
	OffsetX                    float32
	OffsetY                    float32
	Scale                      float32
	AffixTreeId                uint8
	CanRefine                  bool
	RecycleID                  uint8
	DisjoinScoinCost           int32
	DisjoinAddMaterial         AddrTo[[]MechMetaData_DisjoinOutputItem]
	LinkIDList                 AddrTo[[]uint32]
	Quality                    uint8
	StigmataMainID             int32
	ShortName                  AddrTo[Hash]
	SellPriceID                uint8
	Transcendent               bool
	Target                     int32
	IsSecurityProtect          bool
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
	StigmataFilterList         AddrTo[[]int32]
	CollaborationSetID         int32
}

type TagListDataMetaData struct {
	// Properties
	TagName             AddrTo[Hash]
	TagDec              AddrTo[Hash]
	Icon                AddrTo[StrWithPrefix16]
	Type                uint8
	ShowInScreen        bool
	DisplayWeaponList   AddrTo[[]int32]
	DisplayStigmataList AddrTo[[]int32]
}

type ThemeBGMConfigMetaData struct {
	// Properties
	StateName       AddrTo[StrWithPrefix16]
	SwitchName      AddrTo[StrWithPrefix16]
	SwitchGroupName AddrTo[StrWithPrefix16]
	SourceDesc      AddrTo[Hash]
	ShowName        AddrTo[Hash]
}

type VirtualResourcesMetaData struct {
	// Properties
	DisplayTitle         AddrTo[Hash]
	DisplayRarity        int32
	DisplayDescription   AddrTo[Hash]
	IconPath             AddrTo[StrWithPrefix16]
	ImagePath            AddrTo[StrWithPrefix16]
	ProtoID              int32
	LinkIDList           AddrTo[[]uint32]
	DisplayBGDescription AddrTo[Hash]
	CurrencyIconPath     AddrTo[StrWithPrefix16]
}

type WeaponMetaData struct {
	// Properties
	Rarity                     uint8
	MaxRarity                  uint8
	SubRarity                  uint8
	SubMaxRarity               uint8
	Cost                       uint8
	PowerType                  uint8
	MaxLv                      uint8
	ExpType                    uint8
	SellPriceBase              float32
	SellPriceAdd               float32
	GearExpProvideBase         float32
	GearExpPorvideAdd          float32
	BaseType                   uint8
	BodyMod                    AddrTo[StrWithPrefix16]
	DisplayTitle               AddrTo[Hash]
	DisplayDescription         AddrTo[Hash]
	IconPath                   AddrTo[StrWithPrefix16]
	ImagePath                  AddrTo[StrWithPrefix16]
	HPBase                     float32
	HPAdd                      float32
	SPBase                     float32
	SPAdd                      float32
	AttackBase                 float32
	AttackAdd                  float32
	DefenceBase                float32
	DefenceAdd                 float32
	CriticalBase               float32
	CriticalAdd                float32
	ResistanceBase             float32
	ResistanceAdd              float32
	EvoMaterial                AddrTo[[]MechMetaData_DisjoinOutputItem]
	EvoPlayerLevel             int32
	EvoID                      int32
	Prop1ID                    int32
	Prop1Param1                float32
	Prop1Param2                float32
	Prop1Param3                float32
	Prop1Param1Add             float32
	Prop1Param2Add             float32
	Prop1Param3Add             float32
	Prop2ID                    int32
	Prop2Param1                float32
	Prop2Param2                float32
	Prop2Param3                float32
	Prop2Param1Add             float32
	Prop2Param2Add             float32
	Prop2Param3Add             float32
	Prop3ID                    int32
	Prop3Param1                float32
	Prop3Param2                float32
	Prop3Param3                float32
	Prop3Param1Add             float32
	Prop3Param2Add             float32
	Prop3Param3Add             float32
	Protect                    bool
	ExDisjoinCurrencyCost      int32
	ExDisjoinAddMaterial       AddrTo[[]MechMetaData_DisjoinOutputItem]
	DisjoinScoinCost           int32
	DisjoinAddMaterial         AddrTo[[]MechMetaData_DisjoinOutputItem]
	WeaponMainID               int32
	LinkIDList                 AddrTo[[]uint32]
	WeaponQuality              uint8
	SellPriceID                uint8
	Transcendent               bool
	Target                     int32
	GachaMainDropDisplayConfig AddrTo[[]float32]
	GachaGiftDropDisplayConfig AddrTo[[]float32]
	PreloadEffectFolderPath    AddrTo[StrWithPrefix16]
	WeaponFilterList           AddrTo[[]int32]
	SoundBanks                 AddrTo[[]StrWithPrefix16]
	CollaborationWeaponID      int32
	AvatarCustomDisplayID      int32
}

type AbilitySpecialMetaData struct {
	// Properties
	ParamList AddrTo[[]AbilitySpecialMetaData_ParamEntry]
}

type AbilitySpecialMetaData_ParamEntry struct {
	// Objects
	Paramkey   StrWithPrefix16
	ParamValue float32
}

type AccountBuffEffect struct {
	// Properties
	EffectDisplayIcon AddrTo[StrWithPrefix16]
	EffectNameText    AddrTo[Hash]
	EffectDisplayText AddrTo[Hash]
	LimitType         int32
	LimitNum          int32
}

type AccumulatorUIConfigMetaData struct {
	// Properties
	AccDebuff           AddrTo[StrWithPrefix16]
	IconPath            AddrTo[StrWithPrefix16]
	IconColor           AddrTo[StrWithPrefix16]
	CycleColor          AddrTo[StrWithPrefix16]
	RootColor           AddrTo[StrWithPrefix16]
	ResWaveColor        AddrTo[StrWithPrefix16]
	StarColor           AddrTo[StrWithPrefix16]
	GlowColor           AddrTo[StrWithPrefix16]
	MaskSliderFillColor AddrTo[StrWithPrefix16]
	LayerTextColor      AddrTo[StrWithPrefix16]
	MaxLayerTextColor   AddrTo[StrWithPrefix16]
	MaskColor           AddrTo[StrWithPrefix16]
	BaseColor           AddrTo[StrWithPrefix16]
	FillColor           AddrTo[StrWithPrefix16]
	FlashColor          AddrTo[StrWithPrefix16]
	TipsName            AddrTo[Hash]
	TipsDesc            AddrTo[Hash]
}

type AchievementTagMetaData struct {
	// Properties
	TagName    AddrTo[Hash]
	TagColor   AddrTo[StrWithPrefix16]
	LineColor  AddrTo[StrWithPrefix16]
	TypeIcon   AddrTo[StrWithPrefix16]
	Type       int32
	TypeName   AddrTo[StrWithPrefix16]
	TypeNameEn AddrTo[StrWithPrefix16]
}

type ActChallengeMetaData struct {
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
	// Properties
	ButtonType      uint8
	ButtonTextMap   AddrTo[Hash]
	ButtonImagePath AddrTo[StrWithPrefix16]
	Link            AddrTo[StrWithPrefix16]
	RewardIDList    AddrTo[[]int32]
	StartTime       AddrTo[RemoteTime]
	EndTime         AddrTo[RemoteTime]
	MinLevel        int32
	MaxLevel        int32
}

type ActiveCollectionScheduleMetaData struct {
	// Properties
	BGPath             AddrTo[StrWithPrefix16]
	TitleTextMap       AddrTo[StrWithPrefix16]
	TitleIconImagePath AddrTo[StrWithPrefix16]
	CurrencyIDList     AddrTo[[]int32]
	ActiveItemList     AddrTo[[]int32]
	StartTime          AddrTo[RemoteTime]
	EndTime            AddrTo[RemoteTime]
	MinLevel           int32
	MaxLevel           int32
}

type ActivityBbqMetaData struct {
	// Properties
	CostItemID           int32
	CostItemNum          int32
	GratuityTriggerCount int32
}

type ActivityBingoMetaData struct {
	// Properties
	PreBingo             int32
	BingoSize            int32
	BingoNum             int32
	ResetTimes           int32
	ResetItemID          int32
	ResetItemNum         int32
	CostItemID           int32
	CostItemNum          int32
	UnlockTime           AddrTo[StrWithPrefix16]
	HighlightTimeUp      AddrTo[StrWithPrefix16]
	HighlightTimeDown    AddrTo[StrWithPrefix16]
	BingoColorTheme      int32
	TagText              AddrTo[Hash]
	TagImage             AddrTo[StrWithPrefix16]
	BingoCardFrontCircle AddrTo[StrWithPrefix16]
	BingoCardFrontCross  AddrTo[StrWithPrefix16]
	BingoCardBack        AddrTo[StrWithPrefix16]
	TitleText            AddrTo[Hash]
	TipText              AddrTo[Hash]
	BingoPromotText      AddrTo[StrWithPrefix16]
	BingoPromotImage     AddrTo[StrWithPrefix16]
	BingoPromotAvatar    AddrTo[StrWithPrefix16]
	BingoLinkText1       AddrTo[Hash]
	BingoLinkImage1      AddrTo[StrWithPrefix16]
	BingoLinkType1       int32
	BingoLinkParams1     AddrTo[[]int32]
	BingoLinkParamStr1   AddrTo[StrWithPrefix16]
	BingoLinkText2       AddrTo[Hash]
	BingoLinkImage2      AddrTo[StrWithPrefix16]
	BingoLinkType2       int32
	BingoLinkParams2     AddrTo[[]int32]
	BingoLinkParamStr2   AddrTo[StrWithPrefix16]
	BingoRulesImage      AddrTo[StrWithPrefix16]
	BingoWebImage        AddrTo[StrWithPrefix16]
	BingoWebLink         AddrTo[StrWithPrefix16]
}

type ActivityBossRushMetaData struct {
	// Properties
	AvatarTrial          AddrTo[[]int32]
	AvatarInitial        int32
	AvatarUnlockTime     int32
	RushDailyMission     AddrTo[[]int32]
	RushChallengeMission AddrTo[[]int32]
	BossRushImg          AddrTo[[]StrWithPrefix16]
}

type ActivityCardDataMetaData struct {
	// Properties
	MissionIDList      AddrTo[[]int32]
	TriggerMission     int32
	DisplayMission     int32
	MissionListName    AddrTo[Hash]
	MissionListDisplay AddrTo[Hash]
	MissionListReward  AddrTo[[]int32]
	ActivePic          AddrTo[StrWithPrefix16]
	ActivePic1         AddrTo[StrWithPrefix16]
	TextmapRuleTitle   AddrTo[Hash]
	TextmapRule        AddrTo[Hash]
	Xcoordinate        int32
	Ycoordinate        int32
}

type ActivityChallengeMetaData struct {
	// Properties
	ChallengeNum1 int32
	RewardID1     int32
	ChallengeNum2 int32
	ChallengeNum3 int32
}

type ActivityChargeLevelMetaData struct {
	// Properties
	ChargeAttrList AddrTo[[]int32]
}

type ActivityDropLimitData struct {
	// Properties
	MaterialId int32
	Amount     uint16
	Type       uint8
	TypeParam  AddrTo[[]int32]
}

type ActivityDropLimitScheduleMetaData struct {
	// Properties
	DataIdList AddrTo[[]int32]
}

type ActivityLoginBonusMetaData struct {
	// Properties
	ImagePath AddrTo[StrWithPrefix16]
	BonusDesc AddrTo[Hash]
}

type ActivityLoginMetaData struct {
	// Properties
	Type            int32
	TabTextMapID    AddrTo[Hash]
	ActivityType    int32
	EndTime         AddrTo[StrWithPrefix16]
	PlayerLevel     int32
	RewardList      AddrTo[[]int32]
	KeyRewardList   AddrTo[[]int32]
	LoginTitleText  AddrTo[Hash]
	LoginRewardText AddrTo[Hash]
	ImgAvatar       AddrTo[StrWithPrefix16]
	ImgTheme        AddrTo[StrWithPrefix16]
	ImgTitle        AddrTo[StrWithPrefix16]
	ImgBG           AddrTo[StrWithPrefix16]
	ImgBGLight      AddrTo[StrWithPrefix16]
	AvatarPosX      int32
	AvatarPosY      int32
	WebLink         AddrTo[StrWithPrefix16]
	CanSharePost    bool
}

type ActivityMosaicMetaData struct {
	// Properties
	MissionIDList AddrTo[[]int32]
	MosaicNum     int32
	MosaicPic     AddrTo[StrWithPrefix16]
	TextmapRule   AddrTo[StrWithPrefix16]
	MosaicReward  int32
	Xcoordinate   int32
	Ycoordinate   int32
}

type ActivityPanelMetaData struct {
	// Properties
	GroupID        int32
	Type           uint8
	TypeParam      AddrTo[[]int32]
	Description    AddrTo[Hash]
	ImagePath      AddrTo[StrWithPrefix16]
	TextImagePath  AddrTo[StrWithPrefix16]
	SuperTitleName AddrTo[StrWithPrefix16]
	TimeTitleName  AddrTo[Hash]
	BGPath         AddrTo[StrWithPrefix16]
	PanelColor     AddrTo[StrWithPrefix16]
	TabLabelIcon   AddrTo[StrWithPrefix16]
	CompleteDelete bool
	CurrencyHide   bool
}

type ActivityPanelScoreData struct {
	// Properties
	Score int32
}

type ActivityPanelScoreRewardData struct {
	// Properties
	TotalScore int32
	RewardID   int32
	StageID    int32
	LinkType   int32
	LinkParams AddrTo[[]int32]
}

type ActivityPictureMetaData struct {
	// Properties
	MaxStep         int32
	MaxScore        int32
	Explain_Textmap AddrTo[Hash]
	JumpActivity    int32
	BackgroundPic   AddrTo[StrWithPrefix16]
	ContentPrefab   AddrTo[StrWithPrefix16]
}

type ActivityQuestionnaireMetaData struct {
	// Properties
	LinkType   int32
	Website    AddrTo[StrWithPrefix16]
	ImgTheme   AddrTo[StrWithPrefix16]
	ImgBG      AddrTo[StrWithPrefix16]
	ImgBGLight AddrTo[StrWithPrefix16]
}

type ActivityRandomEffectChargeMetaData struct {
	// Properties
	ChargeValue uint8
}

type ActivityRandomEffectMetaData struct {
	// Properties
	BuffChargeList  AddrTo[[]int32]
	DisplayBuffList AddrTo[[]int32]
}

type ActivityRewardShowMetaData struct {
	// Properties
	ShopType        uint32
	ShopID          uint32
	ShowGoodsIDList AddrTo[[]int32]
}

type ActivityRewardShowPanelMetaData struct {
	// Properties
	PopupType          uint8
	ImportantMaterials AddrTo[[]int32]
	Image1Path         AddrTo[StrWithPrefix16]
	Image2Path         AddrTo[StrWithPrefix16]
	Image3Path         AddrTo[StrWithPrefix16]
}

type ActivityScheduleDisplayData struct {
	// Properties
	ItemDisplay AddrTo[[]ActivityScheduleDisplayData_ItemDisplay]
	BeginTime   AddrTo[StrWithPrefix16]
	EndTime     AddrTo[StrWithPrefix16]
	ImagePath   AddrTo[StrWithPrefix16]
	Text        AddrTo[Hash]
	Link        AddrTo[StrWithPrefix16]
}

type ActivityScheduleDisplayData_ItemDisplay struct {
	// Fields
	ItemID  int32
	ItemNum int32
}

type ActivitySchedulePageMetaData struct {
	// Properties
	ActivityTitle AddrTo[Hash]
	TitleImgPath  AddrTo[StrWithPrefix16]
	ActivityDesc  AddrTo[Hash]
	BgImgPath     AddrTo[StrWithPrefix16]
	ShowReward    int32
	EndTime       AddrTo[StrWithPrefix16]
	DescBgImgPath AddrTo[StrWithPrefix16]
}

type ActivityStageRankMetaData struct {
	// Properties
	IsDisplay   bool
	RankGroupID uint32
}

type ActivityTowerMetaData struct {
	// Properties
	ActivityType        int32
	EnterImg            AddrTo[StrWithPrefix16]
	EnterImgColor       AddrTo[StrWithPrefix16]
	UnlockStage         int32
	MaterialID          int32
	LinkChapter         int32
	StageConfigList     AddrTo[[]int32]
	DailyCount          int32
	ClearReward         int32
	SpecialMissionPanel AddrTo[StrWithPrefix16]
}

type ActivityTypeManagementMetaData struct {
	// Properties
	ManagementTypeList AddrTo[[]int32]
}

type ActMetaData struct {
	// Properties
	ChapterId           int32
	NumInChapter        int32
	ActName             AddrTo[Hash]
	ActType             int32
	LevelPannelPath     AddrTo[StrWithPrefix16]
	HardLevelPannelPath AddrTo[StrWithPrefix16]
	HellLevelPannelPath AddrTo[StrWithPrefix16]
	SmallImgPath        AddrTo[StrWithPrefix16]
	BgImgPath           AddrTo[StrWithPrefix16]
	BgPrefabPath        AddrTo[StrWithPrefix16]
	OpenTime            AddrTo[StrWithPrefix16]
	EventGroup          AddrTo[[]ActMetaData_ActEventData]
}

type ActMetaData_ActEventData struct {
	// Fields
	EventID    int32
	PreLevelID int32
}

type ActRitaRankingMetaData struct {
	// Properties
	Rankscore int32
	ImagePath AddrTo[StrWithPrefix16]
	RankName  AddrTo[Hash]
}

type AddParamGroupMetaData struct {
	// Properties
	LevelMin   uint8
	LevelMax   uint8
	BasicList  AddrTo[[]float32]
	AddList    AddrTo[[]float32]
	Desc       AddrTo[Hash]
	DetailDesc AddrTo[Hash]
}

type AdventureAvatarMetaData struct {
	// Properties
	GrowupTestID      int32
	TechnicTestID     AddrTo[[]int32]
	Mastery           AddrTo[[]int32]
	AvatarQuestSucess AddrTo[[]int32]
}

type AdventureAvatarMissionListMetaData struct {
	// Properties
	MissionId   int32
	AchieveType int32
	AvatarID    int32
	BadgeReward int32
}

type AdventureAvatarSkillLevelMetaData struct {
	// Properties
	AvatarSkillID int32
	Level         int32
	Par           float32
}

type AdventureAvatarSkillMetaData struct {
	// Properties
	Type                  int32
	SkillName             AddrTo[Hash]
	SkillDesc             AddrTo[Hash]
	DisplayAdventureQuest AddrTo[Hash]
	Icon                  AddrTo[StrWithPrefix16]
}

type AdventureAvatarUnlockMetaData struct {
	// Properties
	AvatarID int32
	Unlock   int32
	Skill    int32
}

type AdventureDecorateMetaData struct {
	// Properties
	CollectionTypeList AddrTo[[]uint8]
}

type AdventureElfMetaData struct {
	// Properties
	ElfQuestSucess AddrTo[[]int32]
}

type AdventureElfUnlockMetaData struct {
	// Properties
	ElfID  int32
	Unlock int32
	Skill  int32
}

type AdventureLevelDataMetaData struct {
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
	// Properties
	DailyTimes            int32
	AdventureMaxGrain     int32
	AdventureGrainRecover int32
	QuestRefreshTimes     int32
}

type AdventureQuestMetaData struct {
	// Properties
	Rarity           int32
	QuestRare        uint8
	GrainCost        int32
	TimeCost         int32
	NeedTraitList    AddrTo[[]AdventureQuestMetaData_ADQuestSkill]
	BonusDrop        int32
	QuestDropShow    int32
	BonusDropShow    int32
	FailReward       int32
	MinMember        int32
	MaxMember        int32
	QuestNameText    AddrTo[Hash]
	TraitSuccessMax  int32
	AvatarSuccessMax int32
}

type AdventureQuestMetaData_ADQuestSkill struct {
	// Fields
	SkillID       int32
	SkillStrength int32
}

type AdventureQuestPoolMetaData struct {
	// Properties
	MapPosition       int32
	QuestPoolBgPath   AddrTo[StrWithPrefix16]
	QuestPoolIconPath AddrTo[StrWithPrefix16]
	QuestPoolDescText AddrTo[Hash]
}

type AdventureWelfareMetaData struct {
	// Properties
	AdventureLevel int32
	Type           int32
	Par1           int32
	Par2           int32
	Par3           int32
	Title          AddrTo[Hash]
	Dec            AddrTo[Hash]
	Icon           AddrTo[StrWithPrefix16]
}

type AffixEffectMetaData struct {
	// Properties
	Attributetype int32
}

type AffixRecycleMetaData struct {
	// Properties
	MaterialList AddrTo[[]AffixRecycleMetaData_IntIntPair]
}

type AffixRecycleMetaData_IntIntPair struct {
	// Fields
	Item1 int32
	Item2 int32
}

type AffixTreeNodeMetaData struct {
	// Properties
	RefineCost_0to1  AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
	RefineCost_1     AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
	RefineCost_2     AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
	RefineCost_1to2  AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
	RefineCost_core  AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
	RefineCost_lock1 AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
}

type AffixTreeNodeMetaData_MaterialNeed struct {
	// Fields
	MaterialID  int32
	MaterialNum int32
}

type AffixTypeConfigMetaData struct {
	// Properties
	TypeName AddrTo[Hash]
	Icon     AddrTo[StrWithPrefix16]
}

type AiCyberActivityScheduleMetaData struct {
	// Properties
	BeginTime              AddrTo[RemoteTime]
	EndTime                AddrTo[RemoteTime]
	MapID                  int32
	AreaIDList             AddrTo[[]int32]
	MainMissionIDList      AddrTo[[]int32]
	DailyMissionIDList     AddrTo[[]int32]
	ChallengeMissionIDList AddrTo[[]int32]
	TicketID               int32
	CurrencyID             int32
	MaxTicketNum           int32
	RepairMaterialID       int32
	RewardScheduleID       int32
	PerformID              int32
}

type AiCyberAreaMetaData struct {
	// Properties
	AreaName              AddrTo[Hash]
	UnlockConditionList   AddrTo[StrWithPrefix16]
	MainStageGroupID      int32
	DailyStageGroupID     int32
	ChallengeStageGroupID int32
	QAvatarID             AddrTo[[]int32]
	QAvatarSkillID        AddrTo[[]int32]
	PuzzleEventIDList     AddrTo[[]int32]
	PuzzleReward          int32
}

type AiCyberDailyStageDropLimitMetaData struct {
	// Properties
	UnlockConditionList   AddrTo[StrWithPrefix16]
	MaterialDropLimitList AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
}

type AiCyberDailyStageScoreDropMetaData struct {
	// Properties
	DropMaterialList AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
}

type AiCyberGalTouchMetaData struct {
	// Properties
	Dialogue    AddrTo[Hash]
	StartMotion AddrTo[StrWithPrefix16]
}

type AiCyberHyperionMetaData struct {
	// Properties
	SpaceshipPath       AddrTo[StrWithPrefix16]
	AvatarModelPath     AddrTo[StrWithPrefix16]
	UnlockConditionList AddrTo[StrWithPrefix16]
	GalEvents           AddrTo[[]int32]
}

type AiCyberMainUIRepairMetaData struct {
	// Properties
	ID            int32
	ProgressLimit int32
}

type AiCyberProgressMetaData struct {
	// Properties
	SlotType              uint8
	SlotTypeParam         int32
	CostRepairMaterialNum int32
	SlotName              AddrTo[Hash]
	SlotDec               AddrTo[Hash]
	SlotIcon              AddrTo[StrWithPrefix16]
	UnlockConditionList   AddrTo[StrWithPrefix16]
	IsCritical            bool
	IsShow                bool
}

type AiCyberPuzzleEventMetaData struct {
	// Properties
	PuzzleDes                 AddrTo[Hash]
	RewardUnlockConditionList AddrTo[StrWithPrefix16]
	PuzzleSpawnPoint          int32
}

type AiCyberQavatarMetaData struct {
	// Properties
	QAvatarName           AddrTo[Hash]
	QAvatarIcon           AddrTo[StrWithPrefix16]
	QAvatarModel          AddrTo[StrWithPrefix16]
	QAvatarExploreSkillID AddrTo[[]int32]
	QAvatarAttackSkillID  AddrTo[[]int32]
	QAvatarReaction       AddrTo[[]AiCyberQavatarMetaData_Reaction]
	QAvatarColor          AddrTo[StrWithPrefix16]
	UnlockConditionList   AddrTo[StrWithPrefix16]
}

type AiCyberQavatarMetaData_Reaction struct {
	// Objects
	Bubble  Hash
	AnimKey StrWithPrefix16
}

type AiCyberQavatarSkillMetaData struct {
	// Properties
	SkillType    uint8
	SkillName    AddrTo[Hash]
	SkillDec     AddrTo[Hash]
	SkillIcon    AddrTo[StrWithPrefix16]
	SkillAbility AddrTo[StrWithPrefix16]
}

type AiCyberStageGroupMetaData struct {
	// Properties
	StageIDList AddrTo[[]int32]
}

type AiCyberStageMetaData struct {
	// Properties
	UnlockConditionList AddrTo[StrWithPrefix16]
	QAvatarID           AddrTo[[]int32]
	QAvatarSkillID      AddrTo[[]int32]
	AvatarTrialID       AddrTo[[]int32]
}

type AiCyberStoryDataMetaData struct {
	// Properties
	StageID             int32
	UnlockConditionList AddrTo[StrWithPrefix16]
	QAvatarID           AddrTo[[]int32]
	QAvatarSkillID      AddrTo[[]int32]
}

type AnniversaryFifthBoxDataMetaData struct {
	// Properties
	InProgressImagePath AddrTo[StrWithPrefix16]
	FinishedImagePath   AddrTo[StrWithPrefix16]
	MaterialID          uint32
	BoxType             uint32
	BoxStartTime        AddrTo[RemoteTime]
	BoxEndTime          AddrTo[RemoteTime]
	IntParams           AddrTo[[]uint32]
}

type AnniversaryFifthDetailItemsMetaData struct {
	// Properties
	RewardImagePath AddrTo[StrWithPrefix16]
	MinLevel        uint32
	StartTime       AddrTo[RemoteTime]
	EndTime         AddrTo[RemoteTime]
	IntParam        AddrTo[[]uint32]
	ProgressDesc    AddrTo[Hash]
}

type AnniversaryFifthShowItemsMetaData struct {
	// Properties
	RewardImagePath AddrTo[StrWithPrefix16]
	MinLevel        uint32
	StartTime       AddrTo[RemoteTime]
	EndTime         AddrTo[RemoteTime]
	ProgressDesc    AddrTo[Hash]
}

type AnniversaryIconintegrateMetaData struct {
	// Properties
	IconX          int16
	IconY          int16
	IconTitle      AddrTo[Hash]
	StartTime      AddrTo[RemoteTime]
	EndTime        AddrTo[RemoteTime]
	IconGo         AddrTo[StrWithPrefix16]
	IconImage      AddrTo[StrWithPrefix16]
	IconScheduleID int32
	IconRule       AddrTo[Hash]
	MinPlayerLevel uint8
	MaxPlayerLevel uint8
}

type AppointmentDownloadScheduleMetaData struct {
	// Properties
	DownloadLimitTime AddrTo[StrWithPrefix16]
	NoticeTextMapId   AddrTo[Hash]
	ErrorTextMapId    AddrTo[Hash]
}

type ArmadaBossActivityDataMetaData struct {
	// Properties
	EndTime              AddrTo[RemoteTime]
	ShopID               int32
	LevelPanelPrefabPath AddrTo[StrWithPrefix16]
	StageIDList          AddrTo[StrWithPrefix16]
}

type ArmadaBossActivityScoreMetaData struct {
	// Properties
	Reward int32
}

type ArmadaBossData struct {
	// Properties
	ArmadaLevel int32
	Prefab      AddrTo[StrWithPrefix16]
	CostFund    int32
	CostTime    int32
	ApCap       int32
}

type ArmadaContactData struct {
	// Properties
	ArmadaLevel int32
	Prefab      AddrTo[StrWithPrefix16]
	CostFund    int32
	CostTime    int32
	Capacity    int32
	PoplCap     int32
}

type ArmadaExchequerData struct {
	// Properties
	ArmadaLevel int32
	Prefab      AddrTo[StrWithPrefix16]
	CostFund    int32
	CostTime    int32
	FundCap     int32
}

type ArmadaGridData struct {
	// Properties
	X int32
	Y int32
}

type ArmadaLinerRewardMetaData struct {
	// Properties
	NeedScore int32
	RewardID  int32
}

type ArmadaMainData struct {
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
	// Properties
	Type      int32
	TextMap   AddrTo[StrWithPrefix16]
	MutexList AddrTo[[]int32]
}

type ArmadaRecruitTipsMetaData struct {
	// Properties
	Value1       int32
	Value2       int32
	VisibleRange int32
	Priority     int32
	TextMap_Tip  AddrTo[StrWithPrefix16]
}

type ArmadaReunionLevelMetaData struct {
	// Properties
	Demand               AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
	RewardList           AddrTo[[]int32]
	Displayreward1       int32
	Displayreward2       int32
	Displayreward3       int32
	PicPath              AddrTo[StrWithPrefix16]
	TextMap_Level        AddrTo[StrWithPrefix16]
	OriginalPaintingPath AddrTo[StrWithPrefix16]
	ActivePaintingPath   AddrTo[StrWithPrefix16]
}

type ArmadaReunionMissionMetaData struct {
	// Properties
	FinishTime        AddrTo[StrWithPrefix16]
	CountMax          int32
	MissionIntroImage AddrTo[StrWithPrefix16]
	RewardDisplay     AddrTo[[]ArmadaReunionMissionMetaData_MissionRewardDisplay]
}

type ArmadaReunionMissionMetaData_MissionRewardDisplay struct {
	// Fields
	MaterialID int32
	Num        int32
}

type ArmadaReunionPointMetaData struct {
	// Properties
	Point int32
}

type ArmadaReunionScheduleMetaData struct {
	// Properties
	CurrencyID      int32
	OpenTime        AddrTo[StrWithPrefix16]
	RewardTime      AddrTo[StrWithPrefix16]
	CoolingTime     int32
	MaxAwardsTimes  int32
	RewardEndTime   AddrTo[StrWithPrefix16]
	DailyOpenTime   AddrTo[RemoteTimeSpan]
	CloseTime       AddrTo[StrWithPrefix16]
	PlayInstruction AddrTo[StrWithPrefix16]
}

type ArmadaStageBossMetaData struct {
	// Properties
	Name      AddrTo[Hash]
	UniqueID  int32
	ImagePath AddrTo[StrWithPrefix16]
}

type ArmadaStageHardMetaData struct {
	// Properties
	HardText       AddrTo[Hash]
	BossLevel      int32
	SummonPop      int32
	KillReward     AddrTo[[]int32]
	Stage1RewardID int32
	Stage2RewardID int32
	Stage3RewardID int32
}

type ArmadaWelcomeMetaData struct {
	// Properties
	TextMap AddrTo[StrWithPrefix16]
}

type ArmadaWorkshopData struct {
	// Properties
	ArmadaLevel int32
	Prefab      AddrTo[StrWithPrefix16]
	CostFund    int32
	CostTime    int32
	TechLevel   int32
}

type ArtifactTryEnterMetaData struct {
	// Properties
	Level_id          int32
	UnlockLevel       int32
	IsGeneralActivity bool
	GeneralActivityID int32
}

type AudioPackageData struct {
	// Properties
	PackageName                   AddrTo[StrWithPrefix16]
	Version                       AddrTo[StrWithPrefix16]
	PckType                       uint8
	AppointmentDownloadScheduleID uint32
}

type AvatarAttackPunishMetaData struct {
	// Properties
	DamageReduceRate float32
}

type AvatarAttackSpeedParamMetaData struct {
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
	// Properties
	Birthday     AddrTo[Hash]
	Sex          AddrTo[Hash]
	Organization AddrTo[Hash]
	Height       AddrTo[Hash]
	Weight       AddrTo[Hash]
	HomePlace    AddrTo[Hash]
	Story01      AddrTo[Hash]
	Story02      AddrTo[Hash]
	Story03      AddrTo[Hash]
	Dialogue01   AddrTo[Hash]
	Dialogue02   AddrTo[Hash]
	Dialogue03   AddrTo[Hash]
}

type AvatarDefencePunishMetaData struct {
	// Properties
	DamageIncreaseRate float32
}

type AvatarEffectInfoData struct {
	// Properties
	EffectIcon AddrTo[StrWithPrefix16]
}

type AvatarExMetaData struct {
	// Properties
	AvatarIDList       AddrTo[[]int32]
	CommonSubSkillList AddrTo[[]AvatarExMetaData_CommonSkill]
}

type AvatarExMetaData_CommonSkill struct {
	// Objects
	Skill1 StrWithPrefix16
	Skill2 StrWithPrefix16
}

type AvatarFilterMetaData struct {
	// Properties
	Type            int32
	Value           int32
	OrderDec        AddrTo[Hash]
	IsCollaboration bool
}

type AvatarForgeRecommendMetaData struct {
	// Properties
	DefaultType uint8
	HeadPicPath AddrTo[StrWithPrefix16]
}

type AvatarForgeWhiteListMetaData struct {
	// Properties
	ForgeList AddrTo[[]int32]
	IsEffect  bool
}

type AvatarGachaDisplayMetaData struct {
	// Properties
	AvatarUIPath                 AddrTo[StrWithPrefix16]
	AvatarWinAnim                AddrTo[StrWithPrefix16]
	AvatarIdleAnim               AddrTo[StrWithPrefix16]
	AvatarAddWinAnim             AddrTo[StrWithPrefix16]
	AvatarAddIdleAnim            AddrTo[StrWithPrefix16]
	CameraAnimPath               AddrTo[StrWithPrefix16]
	StagePath                    AddrTo[StrWithPrefix16]
	AvatarAnimatorControllerPath AddrTo[StrWithPrefix16]
	WeaponID                     int32
	AudioPattern                 AddrTo[[]AvatarGachaDisplayMetaData_AudioPatternEvent]
	FaceAnimation                AddrTo[[]AvatarGachaDisplayMetaData_FaceAnimationEvent]
	TextPositionType             int32
	TextShowDelay                float32
	StageOffset                  AddrTo[[]float32]
}

type AvatarGachaDisplayMetaData_AudioPatternEvent struct {
	// Fields
	Time float32

	// Objects
	AudioPattern StrWithPrefix16
}

type AvatarGachaDisplayMetaData_FaceAnimationEvent struct {
	// Fields
	Time float32

	// Objects
	FaceAnimation StrWithPrefix16
}

type AvatarGoodFeelMetaData struct {
}

type AvatarLevelGroupMetaData struct {
	// Properties
	LevelMin int32
	LevelMax int32
}

type AvatarMemoirMetaData struct {
	// Properties
	Unlock int32
}

type AvatarMissionActivityMetaData struct {
	// Properties
	BeginTime                AddrTo[StrWithPrefix16]
	EndTime                  AddrTo[StrWithPrefix16]
	AvatarSampleIDList       AddrTo[[]int32]
	DailyContractPoint       int32
	DailyContractPointReward int32
	MissionList              AddrTo[[]int32]
}

type AvatarMissionActivityRewardMetaData struct {
	// Properties
	RewardIDList AddrTo[[]AvatarMissionActivityRewardMetaData_SampleReward]
}

type AvatarMissionActivityRewardMetaData_SampleReward struct {
	// Fields
	RewardID int32
	SampleID int32
}

type AvatarNewbieMetaData struct {
	// Properties
	TextMapID AddrTo[Hash]
}

type AvatarPracticeMainMetaData struct {
	// Properties
	StepIDs AddrTo[[]int32]
}

type AvatarPracticeStepMetaData struct {
	// Properties
	Type              int32
	Text              AddrTo[Hash]
	Param             AddrTo[[]StrWithPrefix16]
	PrefabType        AddrTo[StrWithPrefix16]
	SkillPrefabType   AddrTo[StrWithPrefix16]
	PrefabImagePath   AddrTo[StrWithPrefix16]
	UseDuration       bool
	Duration          float32
	HoldDuration      float32
	AnimatorStateName AddrTo[StrWithPrefix16]
	TimeOffset        float32
	UseSyncBling      bool
	SKLButton         AddrTo[StrWithPrefix16]
}

type AvatarRelayMetaData struct {
	// Properties
	RelayPhaseList AddrTo[[]uint16]
	StageImgPath   AddrTo[StrWithPrefix16]
	StageDesc      AddrTo[Hash]
}

type AvatarRewardGroupMetaData struct {
	// Properties
	NeedProcess int32
	RewardType  int32
	RewardID    int32
}

type AvatarRoleMetaData struct {
	// Properties
	AvatarRoleIcon   AddrTo[StrWithPrefix16]
	AvatarRoleName   AddrTo[Hash]
	AvatarRoleNameEn AddrTo[Hash]
}

type AvatarSubSkillLevelMetaData struct {
	// Properties
	ItemList_1 AddrTo[[]AvatarSubSkillLevelMetaData_SkillUpLevelNeedItem]
}

type AvatarSubSkillLevelMetaData_SkillUpLevelNeedItem struct {
	// Fields
	ItemMetaID int32
	ItemNum    int32
}

type AvatarTrialMetaData struct {
	// Properties
	AvatarID         uint16
	AvatarLV         uint8
	AvatarStar       uint8
	UseGrandKey      bool
	UseMedal         bool
	UseIsland        bool
	UseNewIslandBuff bool
	AvatarSubSkills  AddrTo[[]AvatarModelMetaData_SkillInfoItem]
	DLCAvatarTalents AddrTo[[]AvatarTrialMetaData_TalentInfoItem]
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
}

type AvatarTrialMetaData_TalentInfoItem struct {
	// Fields
	Equiped  int32
	TalentID uint32
	TalentLV uint32
}

type AvatarTryEnterMetaData struct {
	// Properties
	Level_id          int32
	AvatarTry_textmap AddrTo[Hash]
	UnlockLevel       int32
	IsGeneralActivity bool
	GeneralActivityID int32
}

type AvatarTutorialMetaData struct {
	// Properties
	AvatarID        int32
	ArtifactPath    AddrTo[StrWithPrefix16]
	TitleName       AddrTo[Hash]
	SiteList        AddrTo[[]uint32]
	AvatarImgOffset AddrTo[[]float32]
	MissionList     AddrTo[[]int32]
	Width           int32
	GachatypeList   AddrTo[[]int32]
	EntryWeight     uint32
	IsActivityEntry bool
	TagBeginTime    AddrTo[StrWithPrefix16]
	TagEndTime      AddrTo[StrWithPrefix16]
	InfoTextMap     AddrTo[Hash]
}

type AvatarTutorialSiteMetaData struct {
	// Properties
	Site          uint8
	SiteName      AddrTo[Hash]
	SiteStage     int32
	ShowSiteLevel int32
	VideoLink     AddrTo[StrWithPrefix16]
}

type AvatarTwinsMetaData struct {
	// Properties
	StandByMotion1     AddrTo[StrWithPrefix16]
	SecondAvatarID     int32
	StandByMotion2     AddrTo[StrWithPrefix16]
	ExcludeDressIDList AddrTo[[]int32]
}

type AvatarVideoPopupMetaData struct {
	// Properties
	MiddleImageSource  AddrTo[StrWithPrefix16]
	Video1Icon         AddrTo[StrWithPrefix16]
	Video1Link         AddrTo[StrWithPrefix16]
	Video1LinkOpenType uint8
	Video1Position     AddrTo[[]int32]
	Video2Icon         AddrTo[StrWithPrefix16]
	Video2Link         AddrTo[StrWithPrefix16]
	Video2LinkOpenType uint8
	Video2Position     AddrTo[[]int32]
	MaxTimes           int32
}

type BattlePassBannerMetaData struct {
	// Properties
	ScheduleID       uint16
	MinLevel         uint16
	MaxLevel         uint16
	BPIdentification AddrTo[[]uint8]
	BannerImgPath    AddrTo[StrWithPrefix16]
	LinkDataID       uint32
	Description      AddrTo[StrWithPrefix16]
	Description2     AddrTo[StrWithPrefix16]
	EndTime          AddrTo[RemoteTime]
}

type BattlePassChallengeDataMetaData struct {
	// Properties
	ThemeExpReward    uint16
	SeasonExpReward   uint16
	CountMax          uint8
	SortID            uint8
	MissionIntroImage AddrTo[StrWithPrefix16]
	BattlePassType    uint8
}

type BattlePassScheduleMetaData struct {
	// Properties
	BeginTime             AddrTo[RemoteTime]
	EndTime               AddrTo[RemoteTime]
	SeasonLevelConfigID   uint8
	BasicTypeID           int32
	AdvancedTypeID        int32
	LuxuryTypeID          int32
	ExtraRewardID         int32
	ShowRewardID          int32
	ShopEnterIcon         AddrTo[StrWithPrefix16]
	ExpHcoinPrice         uint16
	SeasonExpMaterialID   int32
	IntroImgPathList      AddrTo[[]StrWithPrefix16]
	IntroImageAnimTime    uint16
	PurchaseBeginTime     AddrTo[StrWithPrefix16]
	PurchaseEndTime       AddrTo[StrWithPrefix16]
	DisplayBeginTime      AddrTo[RemoteTime]
	AdvancedRewardFormat  AddrTo[[]BattlePassScheduleMetaData_Format]
	LuxuryRewardFormat    AddrTo[[]BattlePassScheduleMetaData_Format]
	BpCoinID              int32
	CrystalID             int32
	AvatarSprite          AddrTo[StrWithPrefix16]
	LastingDesc           AddrTo[Hash]
	RewardnoticeFreePic   AddrTo[StrWithPrefix16]
	MaxSeasonLevel        uint16
	MaxDisplaySeasonLevel uint16
	DefaultPic            AddrTo[StrWithPrefix16]
}

type BattlePassScheduleMetaData_Format struct {
	// Fields
	Type uint8

	// Objects
	Text  Hash
	Param StrWithPrefix16
}

type BattlePassSeasonLevelConfigMetaData struct {
	// Properties
	NextLevelExp         uint16
	IsNotice             bool
	RewardNoticeImgPath  AddrTo[StrWithPrefix16]
	BasicRewardID        int32
	AdvancedRewardID     int32
	EliteRewardID        int32
	IsConvenientPurchase bool
}

type BattlePassTypeMetaData struct {
	// Properties
	OriginalPrice        int32
	DiscountPrice        int32
	DisplayIcon          AddrTo[StrWithPrefix16]
	DisplayImage         AddrTo[StrWithPrefix16]
	DisplayExpBoxIcon    AddrTo[StrWithPrefix16]
	DisplayExpBoxTextMap AddrTo[StrWithPrefix16]
}

type BattleTypeMetaData struct {
	// Properties
	IconPath AddrTo[StrWithPrefix16]
}

type BbqBonusSetMetaData struct {
	// Properties
	Bonus_1   AddrTo[[]float32]
	Bonus_2   AddrTo[[]float32]
	MovingSpd int32
}

type BbqLevelMetaData struct {
	// Properties
	LevelUpExp      int32
	LevelUpRewardID int32
}

type BeastStageDisplayMetaData struct {
	// Properties
	LevelDesc            AddrTo[StrWithPrefix16]
	LevelDropDisplayList AddrTo[[]BeastStageDisplayMetaData_LevelDropDisplay]
}

type BeastStageDisplayMetaData_LevelDropDisplay struct {
	// Fields
	Count      int32
	MaterialID int32
}

type BeastStageMetaData struct {
	// Properties
	StageHpTotal int32
}

type BeastTreasureMetaData struct {
	// Properties
	TreasurePath AddrTo[StrWithPrefix16]
}

type BossBattleNameListMetaData struct {
	// Properties
	Name AddrTo[Hash]
	Type uint8
}

type BossChallengeActivityMetaData struct {
	// Properties
	StageIDList AddrTo[[]int32]
}

type BossChallengeRewardMetaData struct {
	// Properties
	Time     int32
	RewardID int32
}

type BossChallengeStageMetaData struct {
	// Properties
	Index               int32
	UnlockTime          AddrTo[RemoteTime]
	CollectionMonsterID int32
	RewardConfigID      int32
	MaxFloor            uint16
	IsShowEffect        bool
	BossName            AddrTo[Hash]
}

type BossPointScoreRewardMetaData struct {
	// Properties
	BossPointScheduleID uint32
	Score               uint64
	RewardID            uint32
}

type BossRushBuffConfigMetaData struct {
	// Properties
	BuffGridUnlockTimeOffsets AddrTo[[]uint32]
	BuffPoolIDs               AddrTo[[]uint32]
}

type BossRushBuffMetaData struct {
	// Properties
	BossRushBuffName    AddrTo[Hash]
	BossRushBuffDes     AddrTo[Hash]
	BossRushAbilityName AddrTo[StrWithPrefix16]
	IconPath            AddrTo[StrWithPrefix16]
	BuffType            int32
}

type BossRushBuffPoolMetaData struct {
	// Properties
	BossRushBuffID AddrTo[[]int32]
}

type BossRushStageScheduleMetaData struct {
	// Properties
	StartTimeOffset      uint32
	EndTimeOffset        uint32
	DisplayedStageGroups AddrTo[[]int32]
	BuffConfigID         uint32
	CoinMaterialID       int32
}

type BoxGachaMetaData struct {
	// Properties
	IsHide bool
}

type BridgeLinkMetaData struct {
	// Properties
	BeginTime     AddrTo[StrWithPrefix16]
	EndTime       AddrTo[StrWithPrefix16]
	LinkMinLevel  uint8
	LinkMaxLevel  uint8
	LinkParam     AddrTo[StrWithPrefix16]
	BridgeLinkPic AddrTo[StrWithPrefix16]
}

type BubbleBridgeMetaData struct {
	// Properties
	TriggerMissionID int32
	AvatarID         int32
	Type             int32
	GalEventID       int32
}

type BuffAssistanceActivityMetaData struct {
	// Properties
	ActivityName            AddrTo[Hash]
	ActivityIconPath        AddrTo[StrWithPrefix16]
	UnlockConditionList     AddrTo[StrWithPrefix16]
	LockTips                AddrTo[Hash]
	BuffPoolMaxRefreshTimes uint8
	MissionIDList           AddrTo[[]int32]
	BeginTime               AddrTo[RemoteTime]
	EndTime                 AddrTo[RemoteTime]
}

type BuffAssistanceBuffMetaData struct {
	// Properties
	BuffType              uint8
	BuffName              AddrTo[Hash]
	BuffDescription       AddrTo[Hash]
	BuffSimpleDescription AddrTo[Hash]
	BuffIconPath          AddrTo[StrWithPrefix16]
	BuffStringName        AddrTo[StrWithPrefix16]
	BuffParamList         AddrTo[[]StrWithPrefix16]
}

type BuffAssistanceLevelMetaData struct {
	// Properties
	Title AddrTo[Hash]
}

type BuffAssistanceNPCRobotMetaData struct {
	// Properties
	NPCIconPath    AddrTo[StrWithPrefix16]
	NPCPicturePath AddrTo[StrWithPrefix16]
	BuffList       AddrTo[[]int32]
}

type BuffAssistanceScheduleMetaData struct {
	// Properties
	ScheduleName AddrTo[Hash]
	DeBuffNum    uint8
	EndTime      AddrTo[RemoteTime]
}

type BuffAssistanceStageGroupMetaData struct {
	// Properties
	NormalStageIDList       AddrTo[[]int32]
	ChallengeStageIDList    AddrTo[[]int32]
	BuffSlotPreStageList    AddrTo[[]BuffAssistanceStageGroupMetaData_BuffAssistanceStageBuffNum]
	DebuffSlotPreStageList  AddrTo[[]BuffAssistanceStageGroupMetaData_BuffAssistanceStageBuffNum]
	StageRewardNumRangeList AddrTo[[]int32]
}

type BuffAssistanceStageGroupMetaData_BuffAssistanceStageBuffNum struct {
	// Fields
	BuffNum uint8
	StageID int32
}

type BuffAssistanceWordMetaData struct {
	// Properties
	AssistWordTextmapID AddrTo[Hash]
	CanBeFirst          uint8
	ContinueWordID      AddrTo[[]int32]
	CanBeLast           uint8
}

type BulletinEffectMetaData struct {
	// Properties
	LoadAnimeLastTime uint16
}

type BurdenAlleviationMetaData struct {
	// Properties
	Type               int32
	TypeParam          int32
	TypeParamStrClient AddrTo[[]int32]
	MaxNum             int32
	Priority           int32
	EndTime            AddrTo[RemoteTime]
	BeginTime          AddrTo[RemoteTime]
	BGPath             AddrTo[StrWithPrefix16]
	DisplayInfo        AddrTo[StrWithPrefix16]
	RewardDisplay      int32
	NotOpenInfoText    AddrTo[StrWithPrefix16]
	CloseInfoText      AddrTo[StrWithPrefix16]
	BurdenDialogTitle  AddrTo[StrWithPrefix16]
	BurdenDialogDesc   AddrTo[StrWithPrefix16]
	ShowPopup          int32
}

type BurialPointMetaData struct {
	// Properties
	TutoriaStepId int32
}

type CgCategoryMetaData struct {
	// Properties
	CgCategory_Name   AddrTo[Hash]
	CgCategory_NameEn AddrTo[Hash]
	CgCategory_Icon   AddrTo[StrWithPrefix16]
}

type CgGroupDataMetaData struct {
	// Properties
	PackageID      int32
	CgDownLoadMode int16
}

type CgMetaData struct {
	// Properties
	UnlockType                    uint8
	UnlockCondition               uint32
	LevelIDBegin                  int32
	LevelIDEnd                    int32
	CgCategory                    int32
	CgSubCategory                 int32
	CgGroupID                     AddrTo[[]int32]
	WikiCgScore                   int32
	InitialUnlock                 bool
	CgPath                        AddrTo[StrWithPrefix16]
	CgIconSpritePath              AddrTo[StrWithPrefix16]
	CgLockHint                    AddrTo[Hash]
	InStreamingAssets             int32
	CgPlayMode                    int32
	CgExtraKey                    AddrTo[StrWithPrefix16]
	FileSize                      int32
	PckType                       uint8
	DownloadLimitTime             AddrTo[StrWithPrefix16]
	AppointmentDownloadScheduleID uint32
}

type CGPackageDataMetaData struct {
	// Properties
	Title AddrTo[Hash]
	Desc  AddrTo[Hash]
}

type ChallengeMissionLink struct {
	// Properties
	DisplayIconPath AddrTo[StrWithPrefix16]
	ComeFromList    AddrTo[[]StrWithPrefix16]
}

type ChallengeMissionPanel struct {
	// Properties
	TypePanelTexture AddrTo[StrWithPrefix16]
}

type ChallengeMissionStepData struct {
	// Properties
	TypeID                     int32
	ChallengeMissionList       AddrTo[[]int32]
	ChallengeMissionPreStep    int32
	ChallengeMissionStepReward int32
	StepUnlockLevel            int32
	ChallengeMissionStepBonus  int32
}

type ChapterActivityPhaseActCfgMetaData struct {
	// Properties
	ActID             uint32
	ActLevelPanelPath AddrTo[StrWithPrefix16]
}

type ChapterActivityPhaseMetaData struct {
	// Properties
	CoverImagePath          AddrTo[StrWithPrefix16]
	MainStoryEntryImagePath AddrTo[StrWithPrefix16]
	UnlockConditionList     AddrTo[StrWithPrefix16]
	ProgressText            AddrTo[StrWithPrefix16]
	BGPanelPath             AddrTo[StrWithPrefix16]
	BGMStateName            AddrTo[StrWithPrefix16]
	ActLevelPanelConfigIDs  AddrTo[[]int32]
}

type ChapterActivityPlotActionMetaData struct {
	// Properties
	Type  uint8
	Param AddrTo[[]StrWithPrefix16]
}

type ChapterActivityRewardLineMetaData struct {
	// Properties
	LimitRewardID    int32
	ResidentRewardID int32
	RewardBeginTime  AddrTo[RemoteTime]
	IsPhasesReward   bool
	IsKeyReward      bool
	KeyRewardName    AddrTo[StrWithPrefix16]
	IsAvatarReward   bool
	KeyRewardImage   AddrTo[StrWithPrefix16]
}

type ChapterActivityRewardMetaData struct {
	// Properties
	ChapterID    int32
	EventLevel   int32
	NeedEventExp int32
	RewardID     int32
	UnlockTime   AddrTo[StrWithPrefix16]
}

type ChapterActivityRpgAvatarMetaData struct {
	// Properties
	NormalStageRowPrefabPath    AddrTo[StrWithPrefix16]
	ChallengeStageRowPrefabPath AddrTo[StrWithPrefix16]
	AvatarType                  uint8
	PassiveTalentIDList         AddrTo[[]int32]
	InitiativeTalentID          int32
	ChallengeSiteTyle           uint8
	AvatarImgPath               AddrTo[StrWithPrefix16]
	AvatarThumbnail             AddrTo[StrWithPrefix16]
	EntryAvatarBGPath           AddrTo[StrWithPrefix16]
	AvatarName                  AddrTo[Hash]
	AvatarRestrictDisplayList   AddrTo[[]int32]
	ThemeAvatarID               int32
	CurrencyDisplay             AddrTo[[]int32]
	AvatarTutorialID            int32
	WebLink                     AddrTo[StrWithPrefix16]
	SweepLevelDropLimit         int32
}

type ChapterActivitySectionMetaData struct {
	// Properties
	ChapterID            int32
	IconPos              int32
	SPActName            AddrTo[Hash]
	SPActDesc            AddrTo[Hash]
	UnlockDesc           AddrTo[Hash]
	EntryBG              AddrTo[[]StrWithPrefix16]
	SPActivityType       int32
	ActList              AddrTo[[]int32]
	ActivityScheduleList AddrTo[[]int32]
	UnlockLink           AddrTo[StrWithPrefix16]
	UnlockLevel          int32
	ShowLevel            int32
	EndTime              AddrTo[StrWithPrefix16]
	PreLevelID           int32
}

type ChapterActivityStageMetaData struct {
	// Properties
	IsAllowNormalAvatar        bool
	StageEventList             AddrTo[[]StrWithPrefix16]
	StageEventHideTextList     AddrTo[[]StrWithPrefix16]
	StageEventCompleteTextList AddrTo[[]StrWithPrefix16]
	StageEventDescTextList     AddrTo[[]StrWithPrefix16]
	StageEventIconList         AddrTo[[]StrWithPrefix16]
	StageMark                  AddrTo[StrWithPrefix16]
}

type ChapterActivityZoneBuffDataMetaData struct {
	// Properties
	BuffDesc AddrTo[StrWithPrefix16]
}

type ChapterActivityZoneDataMetaData struct {
	// Properties
	ChapterID       int32
	NumInChapter    int32
	PreZone         int32
	OpenTime        AddrTo[RemoteTime]
	DropLimit       int32
	RepeatStages    AddrTo[[]int32]
	ChallengeStages AddrTo[[]int32]
	BuffRequireList AddrTo[[]int32]
	BuffList        AddrTo[[]int32]
	BuffTitle       AddrTo[StrWithPrefix16]
	BuffDesc        AddrTo[StrWithPrefix16]
	BuffDescShort   AddrTo[StrWithPrefix16]
	BuffIcon        AddrTo[StrWithPrefix16]
	ZoneOpenHint    AddrTo[StrWithPrefix16]
	ZoneRefuseHint  AddrTo[StrWithPrefix16]
}

type ChapterBuffDisplayMetaData struct {
	// Properties
	EnterTime   uint16
	Buff1       uint16
	Buff2       uint16
	Buff3       uint16
	Desc        AddrTo[Hash]
	UpgradeDesc AddrTo[Hash]
}

type ChapterBWFogConfigMetaData struct {
	// Properties
	PreStage int32
	ZoneName AddrTo[StrWithPrefix16]
}

type ChapterBWSiteConfigMetaData struct {
	// Properties
	FilterChapterID        int32
	WorldType              uint8
	PivotID                int32
	PrefabID               int32
	TargetFocusSiteID      int32
	CameraFocusLocalPos    AddrTo[Vector3]
	IsNeedCameraRecover    bool
	ShowSitePreMissionList AddrTo[[]int32]
	PopupIcon              AddrTo[StrWithPrefix16]
	ThreadCollectionID     uint32
}

type ChapterBWSitePivotMetaData struct {
	// Properties
	Pos AddrTo[Vector3]
}

type ChapterBWSitePrefabCfgMetaData struct {
	// Properties
	PrefabPath AddrTo[StrWithPrefix16]
}

type ChapterBWSystemUnlockTipMetaData struct {
	// Properties
	Type    int32
	Param   int32
	Title   AddrTo[StrWithPrefix16]
	Desc    AddrTo[StrWithPrefix16]
	ImgPath AddrTo[StrWithPrefix16]
}

type ChapterBWWorldBuffDataMetaData struct {
	// Properties
	Name        AddrTo[Hash]
	Desc        AddrTo[Hash]
	Icon        AddrTo[StrWithPrefix16]
	BuffIcon    AddrTo[StrWithPrefix16]
	AbilityName AddrTo[StrWithPrefix16]
	ParamList   AddrTo[[]StrWithPrefix16]
}

type ChapterBWWorldBuffSetMetaData struct {
	// Properties
	BuffIDList AddrTo[[]int32]
}

type ChapterBWWorldConfigMetaData struct {
	// Properties
	RuneList                 AddrTo[[]int32]
	MaxEquipRuneNum          int32
	MaxRuneNum               int32
	MaxRuneAffixNum          int32
	UpgradeConsumeMaxRuneNum int32
	RewardLineMaterial       int32
	RewardLineConfigID       int32
	RewardDisplayScheduleID  int32
	LimitRewardEndTime       AddrTo[RemoteTime]
	ExpMaterialList          AddrTo[[]ChapterBWWorldConfigMetaData_ExpMaterial]
	MaxExpMaterialNum        int32
	RuneExpConvertPercent    int32
	RuneConsumeNoticeLevel   int32
	FilterChapterIDCount     int32
	DailyDifficultyCount     int32
	RuneSystemUnlock         int32
	RuneLevelupUnlock        int32
	RuneFusionUnlock         int32
	UnlockMissionNoteLevelID int32
	UnlockBlackWorldLevelID  int32
}

type ChapterBWWorldConfigMetaData_ExpMaterial struct {
	// Fields
	MaterialExp int32
	MaterialID  int32
}

type ChapterBWWorldRuneAffixMetaData struct {
	// Properties
	RuneSkillQuality         uint8
	RuneSkillQualityIconPath AddrTo[StrWithPrefix16]
	RuneSkillDesc            AddrTo[StrWithPrefix16]
	PassiveSkillAbility      AddrTo[StrWithPrefix16]
	PassiveSkillParam        AddrTo[[]float32]
}

type ChapterBWWorldRuneAffixPoolMetaData struct {
	// Properties
	AffixWeightMap AddrTo[StrWithPrefix16]
}

type ChapterBWWorldRuneLevelMetaData struct {
	// Properties
	AtkBase          float32
	DamageBase       float32
	UnlockTime       AddrTo[RemoteTime]
	PreStageID       int32
	MaxAffixNum      int32
	NeedExp          int32
	ActiveSkillDesc  AddrTo[StrWithPrefix16]
	ActiveSkillParam AddrTo[[]float32]
}

type ChapterBWWorldRuneMetaData struct {
	// Properties
	RuneType             uint8
	LevelConfigID        int32
	RuneMainType         uint8
	ActiveSkillAbility   AddrTo[StrWithPrefix16]
	RuneAffixPoolList    AddrTo[[]int32]
	RuneName             AddrTo[StrWithPrefix16]
	RuneIconPath         AddrTo[StrWithPrefix16]
	RuneCurrencyPath     AddrTo[StrWithPrefix16]
	BasicExp             int32
	RuneAffixUnlockLevel AddrTo[[]int32]
}

type ChapterBWWorldRuneSuitMetaData struct {
	// Properties
	RuneMainTypeList AddrTo[[]int32]
	RuneSuitTitle    AddrTo[StrWithPrefix16]
	RuneSuitDesc     AddrTo[StrWithPrefix16]
}

type ChapterBWWorldRuneTrialMetaData struct {
	// Properties
	RuneID    int32
	Level     int32
	AffixList AddrTo[[]int32]
}

type ChapterBWWorldSageTowerCfgMetaData struct {
	// Properties
	SiteIDList AddrTo[[]int32]
	BGImage    AddrTo[StrWithPrefix16]
	Title      AddrTo[StrWithPrefix16]
	VideoTitle AddrTo[StrWithPrefix16]
	VideoDesc  AddrTo[StrWithPrefix16]
	VideoCgID  int32
}

type ChapterBWWorldStageMetaData struct {
	// Properties
	ChapterID               int32
	PreMissionList          AddrTo[[]int32]
	RewardLineID            int32
	IsNeedRank              uint8
	NoteChallengeBGColor    int32
	RestrictRuneTypeList    AddrTo[[]int32]
	RestrictRuneLevelList   AddrTo[[]int32]
	TrialRuneList           AddrTo[[]int32]
	FirstDropExpMaterialMap AddrTo[[]ChapterBWWorldStageMetaData_DropExpMaterial]
	DropExpMaterialMap      AddrTo[[]ChapterBWWorldStageMetaData_DropExpMaterial]
	DropRuneAffixShowList   AddrTo[[]int32]
	NoteDisplayImg          AddrTo[StrWithPrefix16]
	WaveIDList              AddrTo[[]int32]
	LevelWeatherAbility     AddrTo[StrWithPrefix16]
	StageName               AddrTo[StrWithPrefix16]
}

type ChapterBWWorldStageMetaData_DropExpMaterial struct {
	// Fields
	MaterialID  int32
	MaterialNum int32
}

type ChapterBWWorldStageWaveMetaData struct {
	// Properties
	UIDList     AddrTo[[]int32]
	WaveBossUID AddrTo[[]int32]
}

type ChapterBWWorldTowerStageMetaData struct {
	// Properties
	IsSaveFloor      bool
	PassBaseScore    int32
	DisplayRewardNum int32
	RecommendLevel   int32
}

type ChapterChallengeMetaData struct {
	// Properties
	RewardID int32
}

type ChapterCultivateAttrMetaData struct {
	// Properties
	HP             uint32
	ATK            uint32
	DEF            uint32
	AllDamageRatio uint32
}

type ChapterCultivateBuffMetaData struct {
	// Properties
	ActID                 uint32
	BuffGoodsID           uint32
	RequireBuffMaterialID uint32
	IconPath              AddrTo[StrWithPrefix16]
	Name                  AddrTo[Hash]
	Desc                  AddrTo[Hash]
	RowGroupID            uint8
	Phase                 uint8
}

type ChapterCultivateBuffShopMetaData struct {
	// Properties
	ShopID         uint32
	ShopCurrencyID uint32
	Titles         AddrTo[[]Hash]
}

type ChapterCultivateConfigMetaData struct {
	// Properties
	Title    AddrTo[StrWithPrefix16]
	SubTitle AddrTo[StrWithPrefix16]
	Desc     AddrTo[StrWithPrefix16]
}

type ChapterCultivateJumpMetaData struct {
	// Properties
	CultivateSiteID uint32
}

type ChapterCultivateLevelDataMetaData struct {
	// Properties
	CostMaterialID uint32
	CostNum        uint32
}

type ChapterCultivateMonsterMetaData struct {
	// Properties
	MonsterID uint32
}

type ChapterCultivateSiteEffectMetaData struct {
	// Properties
	LevelJumpDisplay      AddrTo[[]ChapterCultivateSiteEffectMetaData_LevelJump]
	MaxFloor              int32
	LevelEffectList       AddrTo[[]ChapterCultivateSiteEffectMetaData_LevelEffect]
	EnterTimeEffectList   AddrTo[StrWithPrefix16]
	RandomEffectList      AddrTo[[]int32]
	BaseEffectList        AddrTo[[]int32]
	UpAvatarList          AddrTo[[]int32]
	UpAvatarConditionNum  int32
	UpAvatarConditionIcon AddrTo[StrWithPrefix16]
	UpAvatarConditionText AddrTo[Hash]
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
	// Properties
	SkillType  uint8
	SkillID    uint32
	SkillLevel uint32
	IconPath   AddrTo[StrWithPrefix16]
	Name       AddrTo[StrWithPrefix16]
	Desc       AddrTo[StrWithPrefix16]
}

type ChapterCultivateUseMetaData struct {
	// Properties
	ChapterID          uint32
	UseType            uint8
	CultivateUseIDList AddrTo[[]uint32]
	CultivateIDList    AddrTo[[]uint32]
}

type ChapterGroupConfigMetaData struct {
	// Properties
	GroupType          uint8
	BeginShowTime      AddrTo[RemoteTime]
	BeginTime          AddrTo[RemoteTime]
	BeginShowLevel     uint8
	SiteList           AddrTo[[]uint8]
	UnlockLevel        uint8
	Title              AddrTo[Hash]
	SubTitle           AddrTo[Hash]
	BGPath             AddrTo[StrWithPrefix16]
	LeftFuncBtn        uint16
	FuncBtnList        AddrTo[[]uint16]
	AutoShowPreviously bool
	PreviouslyImgPath  AddrTo[StrWithPrefix16]
	PreviouslyDesc     AddrTo[Hash]
}

type ChapterGroupFuncButtonMetaData struct {
	// Properties
	Type      uint8
	ImagePath AddrTo[StrWithPrefix16]
	Name      AddrTo[Hash]
	FuncParam AddrTo[StrWithPrefix16]
}

type ChapterGroupSiteMetaData struct {
	// Properties
	SiteType    uint8
	SiteName    AddrTo[Hash]
	SiteChapter int32
	UIConfig    AddrTo[StrWithPrefix16]
	PopUpType   uint8
	IsInterval  bool
}

type ChapterMetaData struct {
	// Properties
	Title                  AddrTo[Hash]
	ChapterDetail          AddrTo[Hash]
	MonsterList            AddrTo[[]int32]
	ChapterType            int32
	VirtualGroupID         int32
	ChapterMaxLevelID      int32
	DefaultSection         int32
	CoverPicture           AddrTo[StrWithPrefix16]
	UnlockDisplay          AddrTo[Hash]
	ActiviyDesc            AddrTo[Hash]
	EventBGPrefabPath      AddrTo[StrWithPrefix16]
	EventSectionPrefabPath AddrTo[StrWithPrefix16]
	BgImgPath              AddrTo[StrWithPrefix16]
	CoverChapter           AddrTo[StrWithPrefix16]
	CoverChapterColor      AddrTo[StrWithPrefix16]
	VersionTagTitle        AddrTo[Hash]
	MissionList            AddrTo[[]int32]
	LoginDayLimit          uint8
}

type ChapterOW31DialogueMetaData struct {
	// Properties
	TipsTextMap AddrTo[StrWithPrefix16]
	Duration    float32
}

type ChapterOWAchievementMetaData struct {
	// Properties
	MapID          uint32
	Name           AddrTo[Hash]
	LockedName     AddrTo[Hash]
	Desc           AddrTo[Hash]
	LockedDesc     AddrTo[Hash]
	ImgPath        AddrTo[StrWithPrefix16]
	IsHide         bool
	IsShowProgress bool
	SequenceID     int32
}

type ChapterOWActivityPanelMetaData struct {
	// Properties
	MapID           int32
	ItemType        uint8
	MaterialType    uint8
	MonitorMaterial int32
	NumberList      AddrTo[[]int32]
	ItemTilte       AddrTo[Hash]
	ItemIconPath    AddrTo[StrWithPrefix16]
	ScoreDisplay    AddrTo[Hash]
	LinkParam       AddrTo[StrWithPrefix16]
}

type ChapterOWAntiGravityGroupMetaData struct {
	// Properties
	MapID               int32
	LevelIDList         AddrTo[[]int32]
	TalentIDList        AddrTo[[]int32]
	MaterialID          AddrTo[[]int32]
	LocationID          int32
	DisplayCondition    int32
	UnlockConditionTips AddrTo[Hash]
	Title               AddrTo[Hash]
}

type ChapterOWAntiGravityLevelMetaData struct {
	// Properties
	MapID                  int32
	GroupID                uint8
	StageID                int32
	LevelType              uint8
	ChallengeShowList      AddrTo[[]ChapterOWAntiGravityLevelMetaData_ChallengeShowCondition]
	SpecialChallengeIDList AddrTo[[]int32]
	SeniorCoinIndexList    AddrTo[[]int32]
	LowerCoinIndexList     AddrTo[[]int32]
	LowerCoinMaterial      AddrTo[ChapterOWAntiGravityLevelMetaData_CoinMaterialData]
	SeniorCoinMaterial     AddrTo[ChapterOWAntiGravityLevelMetaData_CoinMaterialData]
	UnlockConditionList    AddrTo[StrWithPrefix16]
	UnlockConditionTips    AddrTo[[]StrWithPrefix16]
	ImagePath              AddrTo[StrWithPrefix16]
	Title                  AddrTo[Hash]
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
	// Properties
	Type  uint8
	Param uint32
}

type ChapterOWBuildingLevelMetaData struct {
	// Properties
	BuildingType         uint8
	Tips                 AddrTo[Hash]
	CostTime             uint32
	NeedTerminalLevel    uint32
	CostMaterialList     AddrTo[[]ChapterOWBuildingLevelMetaData_CostMaterialData]
	PrefabPath           AddrTo[StrWithPrefix16]
	UpgradeAddPrefabList AddrTo[[]StrWithPrefix16]
	UpgradeAddPathList   AddrTo[[]StrWithPrefix16]
	UpgradeOpenPathList  AddrTo[[]StrWithPrefix16]
	UpgradeClosePathList AddrTo[[]StrWithPrefix16]
	AnimPathList         AddrTo[[]StrWithPrefix16]
	IconPath             AddrTo[StrWithPrefix16]
	RemindTipsList       AddrTo[[]ChapterOWBuildingLevelMetaData_RemindTipData]
	ProgressPanelRadius  float32
	ProgressPanelHeight  float32
}

type ChapterOWBuildingLevelMetaData_CostMaterialData struct {
	// Fields
	CostNum    int32
	MaterialID int32
}

type ChapterOWBuildingLevelMetaData_RemindTipData struct {
	// Fields
	ConditionID int32

	// Objects
	ConditionTextID Hash
}

type ChapterOWBuildingMetaData struct {
	// Properties
	MapID                     int32
	BuildingType              uint8
	IsLobbyBuilding           uint8
	BuildingName              AddrTo[Hash]
	RepairDesc                AddrTo[Hash]
	RepairSuccDesc            AddrTo[Hash]
	StateTypeList             AddrTo[[]int32]
	ConditionList             AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	InitLevel                 uint32
	MaxLevel                  uint32
	MaterialList              AddrTo[[]int32]
	IsShow                    bool
	PrefabPath                AddrTo[StrWithPrefix16]
	IconPath                  AddrTo[StrWithPrefix16]
	MainStoryIDList           AddrTo[[]int32]
	BranchStoryIDList         AddrTo[[]int32]
	NPCHeadMarkOffsetPosition AddrTo[[]float32]
}

type ChapterOWBuildingMetaData_ConditionStatePair struct {
	// Fields
	ConditionID int32
	StateID     int32
}

type ChapterOWBuildingStateMetaData struct {
	// Properties
	Prefabpath      AddrTo[StrWithPrefix16]
	PopType         uint8
	TerminalTextmap AddrTo[Hash]
	ShowType        uint8
}

type ChapterOWChallengeBuffMetaData struct {
	// Properties
	Type uint8
	Icon AddrTo[StrWithPrefix16]
	Desc AddrTo[Hash]
}

type ChapterOWChallengeDataMetaData struct {
	// Properties
	ShowRewardID  uint32
	RewardID      uint32
	Title         AddrTo[Hash]
	MonsterIDList AddrTo[[]int32]
	BuffList      AddrTo[[]int32]
	RecommendBP   int32
	FailedNote    AddrTo[[]Hash]
}

type ChapterOWChallengeGroupMetaData struct {
	// Properties
	LocationID            uint32
	ChallengeConfigID     uint32
	UnlockConditionList   AddrTo[StrWithPrefix16]
	OpenTimeScheduleText  AddrTo[Hash]
	Title                 AddrTo[Hash]
	RecommendHero         uint32
	HighlightMaterialList AddrTo[[]uint32]
}

type ChapterOWChallengeSiteMetaData struct {
	// Properties
	MapID               int32
	EventID             int32
	ChallengeHardLevel  int32
	UnlockConditionList AddrTo[StrWithPrefix16]
	MonsterImgPath      AddrTo[StrWithPrefix16]
	RecommendBuffList   AddrTo[[]int32]
	ChallengeRewardID   int32
	Title               AddrTo[Hash]
	Icon                AddrTo[StrWithPrefix16]
	Desc                AddrTo[Hash]
	MapDisplayTitle     AddrTo[Hash]
}

type ChapterOWCollectionTabMetaData struct {
	// Properties
	MapID          uint32
	CollectionType uint8
	TabName        AddrTo[StrWithPrefix16]
	TabDesc        AddrTo[StrWithPrefix16]
}

type ChapterOWConfigMetaData struct {
	// Properties
	ChapterId                int32
	Type                     int32
	PerformID                int32
	GroupTypeList            AddrTo[[]int32]
	PhaseIDList              AddrTo[[]int32]
	FortIDList               AddrTo[[]int32]
	MainstageEventId         int32
	EquipmentPartIDList      AddrTo[[]int32]
	RewardCountMaterial      int32
	RewardLimitEndTime       AddrTo[RemoteTime]
	AchievementMissionList   AddrTo[[]int32]
	UIManager                AddrTo[StrWithPrefix16]
	GeniusTowerList          AddrTo[[]uint8]
	GeniusTowerConutMaterial int32
	JsonPath                 AddrTo[StrWithPrefix16]
	CostMaterial             int32
	FameMaterial             int32
	FameShopMaterial         int32
	PhotoIdList              AddrTo[[]int32]
	DailyChallengeNumber     int32
	PlotGroupList            AddrTo[[]int32]
	MemoryIDList             AddrTo[[]int32]
	BagDataIDList            AddrTo[[]int32]
	DailyChallengeTowerTimes int32
}

type ChapterOWDailyStageMetaData struct {
	// Properties
	Multiple                 int32
	Group                    int32
	LocationID               int32
	UnlockConditionList      AddrTo[StrWithPrefix16]
	HighLightDropDisplayList AddrTo[[]int32]
	MonsterList              AddrTo[[]int32]
	StageInfoText            AddrTo[StrWithPrefix16]
}

type ChapterOWDigsiteMetaData struct {
	// Properties
	DisplayRewardList AddrTo[[]int32]
	EventID           int32
	ProgramList       AddrTo[[]uint16]
}

type ChapterOWDigsiteProgramMetaData struct {
	// Properties
	ProgramName     AddrTo[Hash]
	ProduceMaterial uint32
	ProduceNumber   uint16
	CostTime        uint32
	LimitTimes      uint16
}

type ChapterOWEndlessChallengeMetaData struct {
	// Properties
	MapID                  int32
	Group                  int32
	PreChallengeID         int32
	FirstReward            int32
	Reward                 int32
	DisplayFirstReward     int32
	DisplayReward          int32
	MonsterIDList          AddrTo[[]int32]
	CostMaterial           AddrTo[StrWithPrefix16]
	MaxScore               int32
	Difficult              int32
	ChallengeDesc          AddrTo[StrWithPrefix16]
	ChallengeHeroRecommend int32
	LockTips               AddrTo[StrWithPrefix16]
	RecommendBP            int32
	GroupName              AddrTo[StrWithPrefix16]
	FailedNote             AddrTo[[]Hash]
}

type ChapterOWEntryPlotLineMetaData struct {
	// Properties
	ConditionID  int32
	PlotLinePath AddrTo[StrWithPrefix16]
}

type ChapterOWEquipmentBuffMetaData struct {
	// Properties
	MapID       int32
	Type        uint16
	Cost        int32
	BuffWeight  int32
	IconPath    AddrTo[StrWithPrefix16]
	SuitID      int32
	Name        AddrTo[Hash]
	Desc        AddrTo[Hash]
	AbilityName AddrTo[StrWithPrefix16]
	ConfigType  AddrTo[StrWithPrefix16]
	ParamList   AddrTo[[]StrWithPrefix16]
	MaterialId  int32
	SourceDesc  AddrTo[Hash]
}

type ChapterOWEquipmentPartMetaData struct {
	// Properties
	UnlockConditionList AddrTo[StrWithPrefix16]
	SlotIDList          AddrTo[[]int32]
	InitialCost         int32
	UINodePath          AddrTo[StrWithPrefix16]
}

type ChapterOWEquipmentSlotMetaData struct {
	// Properties
	UnlockConditionList AddrTo[StrWithPrefix16]
	Type                uint16
	UINodePath          AddrTo[StrWithPrefix16]
}

type ChapterOWEventTabConfigMetaData struct {
	// Properties
	MapID         uint32
	EventIDs      AddrTo[[]uint32]
	TabIconPath   AddrTo[StrWithPrefix16]
	TabDescTextID AddrTo[Hash]
}

type ChapterOWFameMetaData struct {
	// Properties
	FameNumber int32
}

type ChapterOWFortDataMetaData struct {
	// Properties
	QuestIDList  AddrTo[[]int32]
	FortBuffList AddrTo[[]int32]
}

type ChapterOWFurnaceFormulaMetaData struct {
	// Properties
	FormulaName     AddrTo[Hash]
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
}

type ChapterOWFurnitureMetaData struct {
	// Properties
	PrefabPath                AddrTo[StrWithPrefix16]
	ConditionList             AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	NPCHeadMark               AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	NPCHeadMarkScale          float32
	NPCHeadMarkOffsetPosition AddrTo[[]float32]
	MainStoryIDList           AddrTo[[]int32]
	BranchStoryIDList         AddrTo[[]int32]
}

type ChapterOWHeroCardLevelMetaData struct {
	// Properties
	NeedExp     int32
	Desc        AddrTo[StrWithPrefix16]
	AbilityName AddrTo[StrWithPrefix16]
	ParamList   AddrTo[[]float32]
	BattlePoint int32
}

type ChapterOWHeroCardMetaData struct {
	// Properties
	MapID                       int32
	HeroID                      int32
	MaxLevel                    uint8
	Name                        AddrTo[StrWithPrefix16]
	CardContentTextID           AddrTo[Hash]
	Rarity                      int32
	ShowConditionList           AddrTo[StrWithPrefix16]
	OverflowConvertMaterialList AddrTo[StrWithPrefix16]
	SymbiosisCardsList          AddrTo[[]int32]
	CounterCardsList            AddrTo[[]int32]
	ActivationDesc              AddrTo[StrWithPrefix16]
}

type ChapterOWHeroCardRarityMetaData struct {
	// Properties
	RarityBasePath   AddrTo[StrWithPrefix16]
	IconOutLineColor AddrTo[StrWithPrefix16]
	ProgressBarColor AddrTo[StrWithPrefix16]
}

type ChapterOWHeroDisplayMetaData struct {
	// Properties
	UnlockPhaseRange    AddrTo[ChapterOWHeroDisplayMetaData_PhaseRangeData]
	HeroName            AddrTo[StrWithPrefix16]
	HeroDesc            AddrTo[StrWithPrefix16]
	HeroOrderDesc       AddrTo[StrWithPrefix16]
	HeroImagePath       AddrTo[StrWithPrefix16]
	HeroImageForDisplay AddrTo[StrWithPrefix16]
	HeroIconPath        AddrTo[StrWithPrefix16]
}

type ChapterOWHeroDisplayMetaData_PhaseRangeData struct {
	// Fields
	BeginPhaseID int32
	EndPhaseID   int32
}

type ChapterOWHeroLevelMetaData struct {
	// Properties
	UnlockSlotIndex            AddrTo[[]int32]
	SpecialSlotLimitHeroList   AddrTo[[]int32]
	SpecialCardSlotUpgrade     uint8
	LevelUpCostMaterialList    AddrTo[StrWithPrefix16]
	ActiveSkillName            AddrTo[StrWithPrefix16]
	ActiveSkillNextLevelDesc   AddrTo[StrWithPrefix16]
	ActiveSkillDesc            AddrTo[StrWithPrefix16]
	AttributeDisplayDesc       AddrTo[StrWithPrefix16]
	BuildingSkillDesc          AddrTo[StrWithPrefix16]
	BuildingSkillNextLevelDesc AddrTo[StrWithPrefix16]
	ActiveSkill                AddrTo[StrWithPrefix16]
	ActiveSkillParam           AddrTo[[]float32]
	ActiveSkillCD              float32
	AttributeSkill             AddrTo[StrWithPrefix16]
	AttributeSkillParam        AddrTo[[]int32]
	BattlePoint                int32
}

type ChapterOWHeroMetaData struct {
	// Properties
	Type                 uint8
	MapID                int32
	MaterialID           int32
	Name                 AddrTo[StrWithPrefix16]
	BuffName             AddrTo[StrWithPrefix16]
	HeroImagePath        AddrTo[StrWithPrefix16]
	HeroHeadImagePath    AddrTo[StrWithPrefix16]
	ShadowImagePath      AddrTo[StrWithPrefix16]
	HeroIconPath         AddrTo[StrWithPrefix16]
	HeroCardIconPath     AddrTo[StrWithPrefix16]
	IconColor            AddrTo[StrWithPrefix16]
	IconLightColor       AddrTo[StrWithPrefix16]
	HeroMaxLevel         uint8
	SpecialCardSlotIndex AddrTo[[]int32]
	LockHeroName         AddrTo[StrWithPrefix16]
	LockHeroDesc         AddrTo[StrWithPrefix16]
	LockHeroComfrom      AddrTo[StrWithPrefix16]
}

type ChapterOWHeroSpMetaData struct {
	// Properties
	ShowConditionList   AddrTo[StrWithPrefix16]
	UnlockConditionList AddrTo[StrWithPrefix16]
	Reward              int32
	SubLineStegeList    AddrTo[[]int32]
}

type ChapterOWInteractStateMetaData struct {
	// Properties
	NPCID                   int32
	IsShow                  bool
	Interactive             bool
	InteractName            AddrTo[Hash]
	InteractIcon            AddrTo[StrWithPrefix16]
	InteractRadius          float32
	InteractBelowHeight     float32
	InteractAboveHeight     float32
	InteractAngle           float32
	InteractEnterActionList AddrTo[[]int32]
	InteractExitActionList  AddrTo[[]int32]
}

type ChapterOWLandMarkMetaData struct {
	// Properties
	MapID            int32
	DisplayType      uint16
	AreaID           int32
	Icon             AddrTo[StrWithPrefix16]
	Title            AddrTo[Hash]
	Desc             AddrTo[Hash]
	DisplayCondition AddrTo[StrWithPrefix16]
	UnlockCondition  int32
	LocationID       int32
}

type ChapterOWLevelSiteMetaData struct {
	// Properties
	MapID           int32
	UINodePath      AddrTo[StrWithPrefix16]
	LevelId         int32
	ShowCondition   AddrTo[StrWithPrefix16]
	UnlockCondition AddrTo[StrWithPrefix16]
	CloseCondition  AddrTo[StrWithPrefix16]
}

type ChapterOWMainLineMetaData struct {
	// Properties
	PhaseID             uint32
	Type                uint8
	Param               AddrTo[[]uint32]
	UnlockConditionList AddrTo[StrWithPrefix16]
	EntryImagePath      AddrTo[StrWithPrefix16]
	BackgroundIconPath  AddrTo[StrWithPrefix16]
	ChapterDisplay      AddrTo[StrWithPrefix16]
	OrderDisplay        AddrTo[StrWithPrefix16]
}

type ChapterOWMemoryContentMetaData struct {
	// Properties
	ContentType   int32
	ContentParams AddrTo[[]int32]
	Title         AddrTo[Hash]
	Desc          AddrTo[Hash]
	ImagePath     AddrTo[StrWithPrefix16]
}

type ChapterOWMemoryMetaData struct {
	// Properties
	MemoryGroup       int32
	MemoryType        int32
	UIPanelPrefabPath AddrTo[StrWithPrefix16]
}

type ChapterOWMemorySiteMetaData struct {
	// Properties
	MemoryID        int32
	SiteType        int32
	PreSiteList     AddrTo[[]int32]
	ContentID       int32
	UnlockCondition AddrTo[StrWithPrefix16]
	Title           AddrTo[Hash]
	ImagePath       AddrTo[StrWithPrefix16]
}

type ChapterOWMissionMetaData struct {
	// Properties
	Tab      uint32
	Priority int32
}

type ChapterOWMissionTabMetaData struct {
	// Properties
	TabName    AddrTo[Hash]
	UpdateDesc AddrTo[Hash]
}

type ChapterOWMoonConditionMetaData struct {
	// Properties
	ConditionType   uint8
	Parameter       AddrTo[[]int32]
	StringParameter AddrTo[[]StrWithPrefix16]
}

type ChapterOWMoonTowerFloorMetaData struct {
	// Properties
	FirstDropMaterialID  int32
	FirstDropMaterialNum uint16
	FirstDropTerminalExp uint32
	IsSaveFloor          bool
	MonsterID            int32
	Title                AddrTo[Hash]
	RecommendTalentLevel uint16
	MaxScore             uint32
	IsDifficult          bool
	MonsterNatureList    AddrTo[[]int32]
}

type ChapterOWMoonTowerMetaData struct {
	// Properties
	FloorConfigID       uint16
	MapID               int32
	LocationID          uint32
	DisplayCondition    int32
	OpenTime            AddrTo[RemoteTime]
	ScoreDropConfigID   uint16
	DropLimitConfigID   uint16
	RecommendTalentList AddrTo[[]ChapterOWMoonTowerMetaData_RecommendTalent]
	SweepUnlockFloor    uint16
	DescList            AddrTo[[]Hash]
	Title               AddrTo[Hash]
	DisplayMaterialList AddrTo[[]int32]
}

type ChapterOWMoonTowerMetaData_RecommendTalent struct {
	// Fields
	TalentTag uint8
	TalentID  int32
}

type ChapterOWMoonTowerScoreDropMetaData struct {
	// Properties
	DropMaterialID  int32
	DropMaterialNum uint32
}

type ChapterOWNPCMetaData struct {
	// Properties
	PrefabPath                AddrTo[StrWithPrefix16]
	PrefabScale               float32
	IsNeedLookAtPosition      uint8
	ConditionList             AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	NPCHeadMark               AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	NPCHeadMarkOffsetPosition AddrTo[[]float32]
	MainStoryIDList           AddrTo[[]int32]
	BranchStoryIDList         AddrTo[[]int32]
}

type ChapterOWNPCPositionGroupMetaData struct {
	// Properties
	ConditionID    int32
	PositionIDList AddrTo[[]int32]
}

type ChapterOWNPCPositionMetaData struct {
	// Properties
	NPCPosList  AddrTo[[]ChapterOWNPCPositionMetaData_Positions]
	ConditionID int32
}

type ChapterOWNPCPositionMetaData_Positions struct {
	// Fields
	NpcID int32

	// Objects
	PostionName StrWithPrefix16
	AnimID      int32
}

type ChapterOWPhaseMetaData struct {
	// Properties
	UnlockConditionList     AddrTo[StrWithPrefix16]
	ProgressText            AddrTo[Hash]
	CoverImagePath          AddrTo[StrWithPrefix16]
	MainStoryEntryImagePath AddrTo[StrWithPrefix16]
	ActivityEntryImagePath  AddrTo[StrWithPrefix16]
	BGMStateName            AddrTo[StrWithPrefix16]
	PhaseNumber             int32
}

type ChapterOWPlotSiteMetaData struct {
	// Properties
	MapID           int32
	UINodePath      AddrTo[StrWithPrefix16]
	PlotID          int32
	SiteTitle       AddrTo[Hash]
	SiteSubTitle    AddrTo[Hash]
	WindowImgPath   AddrTo[StrWithPrefix16]
	WindowDesc      AddrTo[Hash]
	WindowBtnText   AddrTo[Hash]
	ShowCondition   AddrTo[StrWithPrefix16]
	UnlockCondition AddrTo[StrWithPrefix16]
	PassedCondition AddrTo[StrWithPrefix16]
	CloseCondition  AddrTo[StrWithPrefix16]
}

type ChapterOWQTECircleMetaData struct {
	// Properties
	CircleRangeIDList AddrTo[[]uint16]
	StartAngle        uint16
	CircleModeList    AddrTo[[]ChapterOWQTECircleMetaData_CircleMode]
}

type ChapterOWQTECircleMetaData_CircleMode struct {
	// Fields
	Angle int16
	Speed float32
}

type ChapterOWQTECircleRangeMetaData struct {
	// Properties
	MinAngle         float32
	MaxAngle         float32
	MinPerfectAngle  float32
	MaxPerfectAngle  float32
	BasicRewardNum   int32
	PerfectRewardNum int32
}

type ChapterOWQTELastRewardMetaData struct {
	// Properties
	PullTime      float32
	QTEReduceTime float32
	EnergyCost    float32
	MinInterval   float32
}

type ChapterOWQTEMapGroupMetaData struct {
	// Properties
	GroupType uint8
	MapList   AddrTo[[]ChapterOWQTEMapGroupMetaData_Map]
}

type ChapterOWQTEMapGroupMetaData_Map struct {
	// Fields
	ID     uint16
	Weight uint16
}

type ChapterOWQTEMapMetaData struct {
	// Properties
	MapType              uint8
	MapDepth             uint16
	NormalDropMaterialID int32
	MaxNormalDropNum     int32
	EnergyCost           uint16
	CircleIDList         AddrTo[[]ChapterOWQTEMapMetaData_Circle]
	LastRewardID         uint16
	LastRewardMaterialID int32
	LastRewardNum        int32
	ChallengeTextmap     AddrTo[Hash]
	RecommendEnergy      uint16
}

type ChapterOWQTEMapMetaData_Circle struct {
	// Fields
	Depth uint16
	ID    uint16
}

type ChapterOWQuestDataMetaData struct {
	// Properties
	IsSpecialQuest  bool
	QuestLevel      uint32
	QuestConsigner  AddrTo[StrWithPrefix16]
	QuestInfo       AddrTo[StrWithPrefix16]
	QuestLocationID int32
	RarityTips      AddrTo[StrWithPrefix16]
}

type ChapterOWQuestLevelMetaData struct {
	// Properties
	MapID                int32
	SpecialQuestNum      uint32
	MaxDailyAcceptTimes  uint32
	MaxSlotNum           uint32
	MaxDailyRefreshTimes uint32
}

type ChapterOWRandomStageMetaData struct {
	// Properties
	HardLevelPointRange AddrTo[[]int32]
	DisplayIcon         AddrTo[StrWithPrefix16]
	MainDropCardDisplay int32
	ShowWeight          float32
	MonsterIDList       AddrTo[[]int32]
}

type ChapterOWRangePlotGroupMetaData struct {
	// Properties
	PlotGroupID uint32
	RangeType   uint32
	Range       AddrTo[ChapterOWRangePlotGroupMetaData_PlotGroupRangeData]
}

type ChapterOWRangePlotGroupMetaData_PlotGroupRangeData struct {
	// Fields
	Max uint32
	Min uint32
}

type ChapterOWRefiningLevelUPMetaData struct {
	// Properties
	FormulaList   AddrTo[[]int32]
	FieldNumber   int32
	ReducePercent int32
}

type ChapterOWRelicsEventMetaData struct {
	// Properties
	Reward int32
}

type ChapterOWRelicsMetaData struct {
	// Properties
	MapID                    int32
	RelicsMaterialID         int32
	RelicsFragmentMaterialID int32
	RelicsChangeNumber       uint16
	SiteIDList               AddrTo[[]uint16]
	UnlockTime               AddrTo[RemoteTime]
	ShowCondition            int32
	RelicsReward             int32
	Title                    AddrTo[Hash]
}

type ChapterOWRelicsSiteMetaData struct {
	// Properties
	IsSeniorSite           bool
	SectionID              int32
	UINodeName             AddrTo[StrWithPrefix16]
	EventIDList            AddrTo[[]uint32]
	AntiGravityLevelIDList AddrTo[[]int32]
	LocationID             int32
	Title                  AddrTo[Hash]
	SiteDesc               AddrTo[Hash]
}

type ChapterOWRewardMetaData struct {
	// Properties
	LimitRewardID    int32
	ResidentRewardID int32
	RewardBeginTime  AddrTo[RemoteTime]
	IsPhasesReward   bool
	IsKeyReward      bool
	KeyRewardName    AddrTo[StrWithPrefix16]
	IsAvatarReward   bool
	KeyRewardImage   AddrTo[StrWithPrefix16]
}

type ChapterOWShopGoodsMetaData struct {
	// Properties
	RewardID   int32
	ShopID     int32
	CostNumber int32
	NeedFame   int32
}

type ChapterOWShopLinkMetaData struct {
	// Properties
	Type             int32
	MapID            int32
	SellerImagePath  AddrTo[StrWithPrefix16]
	ChangePhase      int32
	ChangePhaseV2    int32
	TabDesc          AddrTo[StrWithPrefix16]
	CurrencyShowList AddrTo[[]int32]
}

type ChapterOWShopMetaData struct {
	// Properties
	MapID               int32
	EventID             int32
	UnlockConditionList AddrTo[StrWithPrefix16]
	Title               AddrTo[Hash]
}

type ChapterOWSpecialStoryMetaData struct {
	// Properties
	MapID        int32
	SubTitle     AddrTo[StrWithPrefix16]
	Introduction AddrTo[StrWithPrefix16]
	CanExplore   bool
}

type ChapterOWTalentBuffLevelMetaData struct {
	// Properties
	LevelRange          uint8
	MapID               int32
	DisplayDesc         AddrTo[StrWithPrefix16]
	AbilityName         AddrTo[StrWithPrefix16]
	AbilityParamList    AddrTo[[]float32]
	RobotPropState      uint8
	RobotPrefabPath     AddrTo[StrWithPrefix16]
	RobotScale          float32
	RobotPosition       AddrTo[[]float32]
	RobotRotation       AddrTo[[]float32]
	RobotImgPath        AddrTo[StrWithPrefix16]
	RobotUpgradeImgPath AddrTo[StrWithPrefix16]
	RobotOutlineBGPath  AddrTo[StrWithPrefix16]
}

type ChapterOWTalentDataMetaData struct {
	// Properties
	MapID               int32
	TalentType          uint8
	IsDefaultActivate   uint8
	AutoActivate        int32
	UniqueTag           uint8
	MaxLevel            uint8
	RelateTalentIDList  AddrTo[[]int32]
	UIChangePath        AddrTo[StrWithPrefix16]
	UIPointPath         AddrTo[StrWithPrefix16]
	UILinePath          AddrTo[[]StrWithPrefix16]
	TalentName          AddrTo[Hash]
	TalentDesc          AddrTo[Hash]
	IconPath            AddrTo[StrWithPrefix16]
	CGID                AddrTo[[]int32]
	UnlockConditionList AddrTo[StrWithPrefix16]
}

type ChapterOWTalentLevelMetaData struct {
	// Properties
	TalentDesc              AddrTo[Hash]
	NextLevelDesc           AddrTo[Hash]
	PreTalent               AddrTo[[]ChapterOWTalentLevelMetaData_PreTalentRequirement]
	NeedWorkShopLevel       uint8
	CostMaterialList        AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	AbilityName             AddrTo[StrWithPrefix16]
	AbilityParamList        AddrTo[[]float32]
	AbilityParamDisplayList AddrTo[[]float32]
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
	// Properties
	LevelExp               int32
	IsBreak                uint8
	BattleTalentTotalLevel uint8
	ConsumeMaterialList    AddrTo[[]ChapterOWTerminalLevelMetaData_ConsumeMaterial]
	PreStoryList           AddrTo[[]int32]
	PlotlinePath           AddrTo[StrWithPrefix16]
	HardLevel              uint8
	PrefabPath             AddrTo[StrWithPrefix16]
	IconPath               AddrTo[StrWithPrefix16]
}

type ChapterOWTerminalLevelMetaData_ConsumeMaterial struct {
	// Fields
	MatrialID int32
	Num       int32
}

type ChapterOWTerminalTipsMetaData struct {
	// Properties
	MapID       int32
	Type        uint8
	Title       AddrTo[Hash]
	Desc        AddrTo[Hash]
	ConditionID int32
	Weight      int32
}

type ChapterOWTipsMetaData struct {
	// Properties
	MapID               int32
	Type                int16
	PopTiming           int16
	UnlockConditionList AddrTo[StrWithPrefix16]
	Icon                AddrTo[StrWithPrefix16]
	Title               AddrTo[Hash]
	Desc                AddrTo[Hash]
	TutorialID          int32
	MaterialID          int32
	Params              AddrTo[[]int32]
}

type ChapterRewardLineCfgMetaData struct {
	// Properties
	RewardLineConfigID       int32
	RewardLineMatID          int32
	LimitedRewardEndTime     AddrTo[RemoteTime]
	MissionIDs               AddrTo[[]int32]
	AllRewardReleasedTime    AddrTo[RemoteTime]
	AllMissionReleasedTime   AddrTo[RemoteTime]
	DefaultKeyImgPath        AddrTo[StrWithPrefix16]
	IsDefaultKeyImgAvatar    bool
	MissionRewardPrefabPath  AddrTo[StrWithPrefix16]
	HideGotNumMaterialIDList AddrTo[[]int32]
}

type ChapterRpgAvatarRewardLineMetaData struct {
	// Properties
	StageID       int32
	RequiredScore int32
	RewardID      int32
}

type ChapterRpgBuffLevelMetaData struct {
	// Properties
	LevelRange       int32
	AbilityName      AddrTo[StrWithPrefix16]
	AbilityParamList AddrTo[[]StrWithPrefix16]
}

type ChapterStageCompensationMetaData struct {
	// Properties
	FinishReward   int32
	UnfinishReward int32
}

type ChapterSwitchAnimConfigMetaData struct {
	// Properties
	ChapterImage         AddrTo[StrWithPrefix16]
	ChapterBgImage       AddrTo[StrWithPrefix16]
	ChapterTitleTextID   AddrTo[Hash]
	ChapterContentTextID AddrTo[Hash]
}

type ChatGroupIconMetaData struct {
	// Properties
	IconPath AddrTo[StrWithPrefix16]
}

type ChatLobbyActionMetaData struct {
	// Properties
	Name        AddrTo[Hash]
	Type        uint8
	ActionIndex int16
	Icon        AddrTo[StrWithPrefix16]
	Loop        bool
}

type ChatLobbyActivityMetaData struct {
	// Properties
	Type int32
	Para int32
}

type ChatLobbyActivityNoticeMetaData struct {
	// Properties
	OpeningShowText  AddrTo[Hash]
	ActivityIcon     AddrTo[StrWithPrefix16]
	ActivityNameText AddrTo[Hash]
	ActivityTypeText AddrTo[Hash]
	LevelText        AddrTo[Hash]
	TimeText         AddrTo[Hash]
	MainRewardList   AddrTo[[]int32]
	ActivityDescText AddrTo[Hash]
	RewardDetailList AddrTo[[]int32]
}

type ChatLobbyActivityScheduleMetaData struct {
	// Properties
	ActivityId int32
}

type ChatLobbyAnnouncementMetaData struct {
	// Properties
	AnnouncementType int32
	BeginTime        AddrTo[StrWithPrefix16]
	Title            AddrTo[Hash]
	Desc             AddrTo[Hash]
	ShowTimes        int16
	ShowSeconds      int16
	EndTime          AddrTo[StrWithPrefix16]
}

type ChatLobbyBeastMetaData struct {
	// Properties
	BeastModelPath    AddrTo[StrWithPrefix16]
	BeastMaterialPath AddrTo[StrWithPrefix16]
	Slogan            AddrTo[StrWithPrefix16]
	HeadPicPath       AddrTo[StrWithPrefix16]
}

type ChatLobbyBoxActivityMetaData struct {
	// Properties
	BoxSeriesID          AddrTo[[]int32]
	PlayerDailyOpenLimit int32
}

type ChatLobbyBoxSeriesMetaData struct {
	// Properties
	BoxObject      int32
	BoxTriggerType int32
	Para02         AddrTo[Hash]
}

type ChatLobbyBubbleDialogMetaData struct {
	// Properties
	Weight         float32
	DialogText     AddrTo[Hash]
	DialogDuration uint8
	DialogAudio    AddrTo[StrWithPrefix16]
}

type ChatLobbyCinemaMetaData struct {
	// Properties
	CgID        int32
	CGDuration  float32
	CGFrameRate uint8
	Cover       AddrTo[StrWithPrefix16]
	NameText    AddrTo[Hash]
}

type ChatLobbyCinemaTimeGroupMetaData struct {
	// Properties
	StartTime           AddrTo[RemoteTime]
	EndTime             AddrTo[RemoteTime]
	Order               AddrTo[[]int32]
	IsDownloadable      bool
	NotDownloadableText AddrTo[Hash]
}

type ChatLobbyDishLimitedRewardMetaData struct {
	// Properties
	BeginTime              AddrTo[StrWithPrefix16]
	EndTime                AddrTo[StrWithPrefix16]
	DishDailyRewardDisplay AddrTo[[]int32]
}

type ChatLobbyDishMetaData struct {
	// Properties
	RewardIDList      AddrTo[[]ChatLobbyDishMetaData_RewardTurn]
	RewardGetLimit    uint8
	RewardGetColdDown uint16
}

type ChatLobbyDishMetaData_RewardTurn struct {
	// Fields
	RewardID1st int32
	RewardID2nd int32
}

type ChatLobbyDishRateMetaData struct {
	// Properties
	TotalStartTime AddrTo[StrWithPrefix16]
	TotalEndTime   AddrTo[StrWithPrefix16]
	Rate           uint8
}

type ChatLobbyFightExpressionMetaData struct {
}

type ChatLobbyFishActivityMetaData struct {
	// Properties
	FishCurrencyId int32
	FishSpotList   AddrTo[[]int32]
}

type ChatLobbyFishMetaData struct {
	// Properties
	FishGroup      uint8
	Weight         uint8
	FishPicID      AddrTo[StrWithPrefix16]
	Textmap_Name   AddrTo[Hash]
	Textmap_Desc   AddrTo[Hash]
	IsRare         uint8
	FishType       uint8
	AddCurrencyID  int32
	AddCurrencyNum int32
}

type ChatLobbyFishShowMetaData struct {
	// Properties
	FishType uint8
	FishPic  AddrTo[StrWithPrefix16]
	FishName AddrTo[Hash]
	FishDesc AddrTo[Hash]
}

type ChatLobbyItemMetaData struct {
	// Properties
	Type                 uint8
	Effect               AddrTo[StrWithPrefix16]
	IsUsed               uint8
	IsShowUseTips        bool
	TargetType           uint8
	UseInfo              AddrTo[Hash]
	IsShowBg             bool
	ItemUseBroadcastType uint8
	ItemUseText          AddrTo[Hash]
}

type ChatLobbyKillingStreakMetaData struct {
	// Properties
	Name         AddrTo[Hash]
	Icon         AddrTo[StrWithPrefix16]
	AudioPattern AddrTo[StrWithPrefix16]
}

type ChatLobbyLanternDifficultyMetaData struct {
	// Properties
	LanternSignSpeed    float32
	LanternFailedScale  float32
	LanternGoodScale    float32
	LanternPerfectScale float32
}

type ChatLobbyLayoutMetaData struct {
	// Properties
	RoomID     int32
	ObjectType uint8
	ObjectID   int32
	X          float32
	Y          float32
	Z          float32
	Angle      float32
	StartTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
}

type ChatLobbyMainPageNoticeMetaData struct {
	// Properties
	NotificationIcon AddrTo[StrWithPrefix16]
	NotificationText AddrTo[Hash]
}

type ChatLobbyMissionMetaData struct {
	// Properties
	NavigateX      float32
	NavigateY      float32
	NavigateZ      float32
	NavigateRadius float32
}

type ChatLobbyNPCMetaData struct {
	// Properties
	QAvatarID               int32
	Name                    AddrTo[Hash]
	NameOffset              float32
	IsNameShown             bool
	HeadIcon                AddrTo[StrWithPrefix16]
	Interactable            bool
	IdleList                AddrTo[[]ChatLobbyNPCMetaData_Idle]
	IdleIntervalMin         uint8
	IdleIntervalMax         uint8
	DialogOffset            float32
	NPCDialog               AddrTo[[]int32]
	NPCBubbleDialog         AddrTo[[]int32]
	BubbleDialogIntervalMin uint8
	BubbleDialogIntervalMax uint8
}

type ChatLobbyNPCMetaData_Idle struct {
	// Objects
	Path   StrWithPrefix16
	Weight float32
}

type ChatLobbyNPCSystemFunctionMetaData struct {
	// Properties
	LinkType     int32
	LinkParams   AddrTo[[]int32]
	LinkParamStr AddrTo[StrWithPrefix16]
}

type ChatLobbyObjectMetaData struct {
	// Properties
	Name                   AddrTo[Hash]
	NameHight              float32
	IsNameShown            bool
	AssetPath              AddrTo[StrWithPrefix16]
	InteractiveType        uint8
	InteractiveRange       float32
	InteractiveButtonRange float32
	InteractiveButtonIcon  AddrTo[StrWithPrefix16]
	NoticeEffect           int32
	NoticeEffectRange      float32
	NoticeEffectPath       AddrTo[StrWithPrefix16]
	ChatLobbyPlotID        AddrTo[[]int32]
}

type ChatLobbyPlotMetaData struct {
	// Properties
	PlotID           int32
	TriggerType      uint8
	Param1           int32
	TutorialKey      AddrTo[StrWithPrefix16]
	TriggerStartTime AddrTo[RemoteTime]
	TriggerEndTime   AddrTo[RemoteTime]
}

type ChatLobbyPrayMetaData struct {
	// Properties
	Type          uint16
	PrayMainTitle AddrTo[Hash]
	PrayText1     AddrTo[Hash]
	PrayText2     AddrTo[Hash]
	IconPath      AddrTo[StrWithPrefix16]
	RewardID      int32
}

type ChatLobbyQuestionActivityMetaData struct {
	// Properties
	RightBuff            uint16
	WrongBuff            uint16
	QuestionActivityType uint16
}

type ChatLobbyQuestionMetaData struct {
	// Properties
	Question_Text AddrTo[StrWithPrefix16]
	AnswerList    AddrTo[[]ChatLobbyQuestionMetaData_AnswerItem]
}

type ChatLobbyQuestionMetaData_AnswerItem struct {
	// Fields
	ID int32

	// Objects
	Text StrWithPrefix16
}

type ChatLobbyRoomMetaData struct {
	// Properties
	Name                 AddrTo[Hash]
	Type                 uint8
	LeftBoardPageList    AddrTo[[]int32]
	ShowMissionList      AddrTo[[]int32]
	QAvatarList          AddrTo[[]int32]
	PlayerSpawnPoint     AddrTo[[]float32]
	PlayerRotation       float32
	CameraRotation       float32
	IsCameraLocked       bool
	CameraRotationReset  bool
	CameraRotationMin    float32
	CameraRotationMax    float32
	StagePath            AddrTo[StrWithPrefix16]
	Weather              AddrTo[StrWithPrefix16]
	ChatLobbySupportJump int32
	ChatLobbyMainUIPath  AddrTo[StrWithPrefix16]
}

type ChatLobbySceneItemMetaData struct {
	// Properties
	Type                 int32
	DestroyEffectPattern AddrTo[StrWithPrefix16]
	InteractRange        float32
	SceneItemNameText    AddrTo[Hash]
	IsNameShow           bool
	SceneItemModel       AddrTo[StrWithPrefix16]
}

type ChatLobbySkillMetaData struct {
	// Properties
	CDTime        float32
	CostBulletNum int32
}

type ChatLobbyStanceMetaData struct {
	// Properties
	StanceName                 AddrTo[Hash]
	StanceModel                AddrTo[StrWithPrefix16]
	StanceOccupiedAnnouncement int32
	ProcessBarOffset           float32
	StanceRange                float32
}

type ChatLobbyTreasureMetaData struct {
	// Properties
	TreasureNameText    AddrTo[Hash]
	ProcessOffset       float32
	TreasureModel       AddrTo[StrWithPrefix16]
	CostMaterialList    AddrTo[[]BeastStageDisplayMetaData_LevelDropDisplay]
	OpenNum             int32
	DropDisplayItemList AddrTo[[]ChatLobbyTreasureMetaData_DisplayItem]
}

type ChatLobbyTreasureMetaData_DisplayItem struct {
	// Fields
	Count  int32
	ItemID int32
}

type ChatLobbyUsableItemMetaData struct {
	// Properties
	UsableItemIcon AddrTo[StrWithPrefix16]
}

type ChatLobbyVoiceMetaData struct {
	// Properties
	AudioID      AddrTo[StrWithPrefix16]
	RoleID       int32
	FromHour     uint8
	ToHour       uint8
	AudioPattern AddrTo[StrWithPrefix16]
	Weight       float32
}

type ChatLobbyWayPointMetaData struct {
	// Properties
	Width   float32
	Depth   float32
	Poslist AddrTo[[]ChatLobbyWayPointMetaData_FishWayPoint]
}

type ChatLobbyWayPointMetaData_FishWayPoint struct {
	// Fields
	X float32
	Z float32
}

type ChatLobbyWoodenBoyMetaData struct {
	// Properties
	ScheduleID                  int32
	MaxCredit                   float32
	CreditIncreaseRate          float32
	CreditDecrease              float32
	TimeBetweenRound            float32
	TimePerRound                float32
	OpenEyesToleranceTime       float32
	PauseCreditIncrease         float32
	MaxPlayerNum                uint16
	GameBgm                     AddrTo[StrWithPrefix16]
	GamePlayLayoutID            int32
	NPCDialogOpenEyes           AddrTo[[]uint16]
	NPCDialogCloseEyes          AddrTo[[]uint16]
	NPCDialogBeforeStart        uint16
	NPCDialogSpotPlayerMovement AddrTo[[]uint16]
	NPCDialogPlayerWinText      AddrTo[[]uint16]
	NPCDialogPlayerLoseText     AddrTo[[]uint16]
	PlayerAvatarQWinAction      int16
	PlayerAvatarQLoseAction     int16
	PlayerAvatarQPunishAction   int16
	PlayerAvatarQActionGroup    AddrTo[[]int16]
	NPCAvatarQWinAction         int16
	NPCAvatarQLoseAction        int16
	NPCAvatarQInactiveAction    int16
	NPCactiveAction             int16
}

type ChatMessageContentMetaData struct {
	// Properties
	ChatContent AddrTo[StrWithPrefix16]
}

type ChatMessageDataMetaData struct {
	// Properties
	ChatMessageList AddrTo[[]int32]
}

type ChatReportMetaData struct {
	// Properties
	SortWeight     int32
	ReportTypeName AddrTo[Hash]
}

type CityEventPhotoMetaData struct {
	// Properties
	PlotID     int32
	IsDefault  bool
	PhotoType  int32
	ActivityID int32
	PlotName   AddrTo[Hash]
	UnlockTips AddrTo[Hash]
	Type       int32
	Role       int32
	RolePic    AddrTo[StrWithPrefix16]
	Face       AddrTo[StrWithPrefix16]
	BackGround AddrTo[StrWithPrefix16]
}

type ClickDialogBGCloseBlackListMetaData struct {
	// Properties
	Hierachies AddrTo[[]StrWithPrefix16]
}

type ClientLogDataMetaData struct {
	// Properties
	ContextName  AddrTo[StrWithPrefix16]
	FeatureKey   AddrTo[StrWithPrefix16]
	NeedEntrance bool
}

type CollaborationScheduleMetaData struct {
	// Properties
	CollaborationStigmataIDList AddrTo[[]int32]
	CollaborationWeaponIDList   AddrTo[[]int32]
	AvatarIDList                AddrTo[[]int32]
	RoleIDList                  AddrTo[[]int32]
	DressIDList                 AddrTo[[]int32]
	MedalIDList                 AddrTo[[]int32]
}

type CollaborationWeaponMetaData struct {
	// Properties
	ImagePath AddrTo[StrWithPrefix16]
	Priority  int32
}

type CollectionDialogueMetaData struct {
	// Properties
	Name      AddrTo[Hash]
	NameBg    AddrTo[StrWithPrefix16]
	BankName  AddrTo[StrWithPrefix16]
	AudioName AddrTo[StrWithPrefix16]
}

type CollectionEventMetaData struct {
	// Properties
	Name                  AddrTo[Hash]
	NameBg                AddrTo[StrWithPrefix16]
	PlotID                int32
	HideContentWhenLocked bool
	LockedText            AddrTo[Hash]
	Desc                  AddrTo[Hash]
}

type CollectionFileMetaData struct {
	// Properties
	Name                  AddrTo[Hash]
	NameBg                AddrTo[StrWithPrefix16]
	Desc                  AddrTo[Hash]
	HideContentWhenLocked bool
	LockedText            AddrTo[Hash]
}

type CollectionItemDataMetaData struct {
	// Properties
	MaterialID int32
}

type CollectionMonsterDetailInfoMetaData struct {
	// Properties
	Desc       AddrTo[Hash]
	Mission    int32
	LockedDesc AddrTo[Hash]
}

type CollectionMonsterMetaData struct {
	// Properties
	UniqueMonsterID       AddrTo[[]int32]
	AllowNormal           bool
	Name                  AddrTo[Hash]
	TypeName              AddrTo[Hash]
	Desc                  AddrTo[Hash]
	Icon                  AddrTo[StrWithPrefix16]
	PrefabPath            AddrTo[StrWithPrefix16]
	Scale                 float32
	Position              AddrTo[[]float32]
	Rotation              AddrTo[[]float32]
	Health                int32
	Attack                int32
	AbilitiesList         AddrTo[[]int32]
	DetailInfoIDs         AddrTo[[]int32]
	HideContentWhenLocked bool
	LockedText            AddrTo[Hash]
}

type CollectionPictureMetaData struct {
	// Properties
	Name                  AddrTo[Hash]
	Icon                  AddrTo[StrWithPrefix16]
	ShowImg               AddrTo[StrWithPrefix16]
	HideContentWhenLocked bool
	LockedText            AddrTo[Hash]
}

type CollectionTypeMetaData struct {
	// Properties
	Title AddrTo[Hash]
	Desc  AddrTo[Hash]
}

type CollectionVisualNovelDataMetaData struct {
	// Properties
	LuaPath AddrTo[StrWithPrefix16]
}

type CompensationOfDormitoryMetaData struct {
	// Properties
	FacilityName AddrTo[Hash]
	ItemList     AddrTo[[]CompensationOfDormitoryMetaData_CompensationOfDormitoryRewardMetaData]
}

type CompensationOfDormitoryMetaData_CompensationOfDormitoryRewardMetaData struct {
	// Fields
	Num      int32
	RewardID int32
}

type CompensationOfElfBreakMetaData struct {
	// Properties
	MaxLevel int32
	ItemList AddrTo[[]ElfMetaData_MatCost]
}

type CompensationOfElfSlotUnlockMetaData struct {
	// Properties
	ItemList AddrTo[[]ElfMetaData_MatCost]
}

type CompensationOfElfTalentMetaData struct {
	// Properties
	ItemList AddrTo[[]ElfMetaData_MatCost]
}

type CompensationOfExtendGradeMetaData struct {
	// Properties
	ItemList AddrTo[[]CompensationOfDormitoryMetaData_CompensationOfDormitoryRewardMetaData]
}

type CompensationOfIslandMetaData struct {
	// Properties
	CabinName AddrTo[Hash]
	ItemList  AddrTo[[]CompensationOfDormitoryMetaData_CompensationOfDormitoryRewardMetaData]
}

type ConstValueMetaData struct {
	// Properties
	Type  AddrTo[StrWithPrefix16]
	Value AddrTo[StrWithPrefix16]
}

type CoupleTowerRewardMetaData struct {
	// Properties
	MaxFloor int32
	RewardID int32
}

type CoupleTowerScoreMetaData struct {
	// Properties
	ScoreType int32
	Score     int32
}

type CouponMetaData struct {
	// Properties
	CouponTypeList  AddrTo[[]int32]
	CouponNeedPrice int32
	CouponValue     int32
	BeginTime       AddrTo[StrWithPrefix16]
	EndTime         AddrTo[StrWithPrefix16]
	SortId          int32
	IsShow          int32
}

type CreditRankMetaData struct {
	// Properties
	RankName        AddrTo[StrWithPrefix16]
	RankColor       AddrTo[StrWithPrefix16]
	ShowWarningHint int32
	WarningHint     AddrTo[StrWithPrefix16]
	MinScore        int32
	MaxScore        int32
}

type CreditRegeditMetaData struct {
	// Properties
	TypeList  AddrTo[[]int32]
	Title     AddrTo[StrWithPrefix16]
	RuleTitle AddrTo[StrWithPrefix16]
	RuleInfo  AddrTo[StrWithPrefix16]
}

type CrisisModeActivityConfigMetaData struct {
	// Properties
	PreLevelID int32
	AvatarPool AddrTo[[]int32]
}

type CrisisModeStageConfigMetaData struct {
	// Properties
	Reward int32
}

type CrisisModeStageGroupConfigMetaData struct {
	// Properties
	TagList AddrTo[[]int32]
}

type CurrencyIconMetaData struct {
	// Properties
	IconPath AddrTo[StrWithPrefix16]
}

type CurrencyQuickExchangeMetaData struct {
	// Properties
	DialogTitle              AddrTo[Hash]
	ExchangeMethodDesc       AddrTo[Hash]
	CurrentNumberDesc        AddrTo[Hash]
	ExchangeNumberDesc       AddrTo[Hash]
	ObtainNumberDesc         AddrTo[Hash]
	CancelDesc               AddrTo[Hash]
	OkDesc                   AddrTo[Hash]
	TipsDesc                 AddrTo[Hash]
	EnableDiamondExchange    int32
	EnablePurpleJadeExchange int32
	EnableMCoinExchange      int32
	MaxBuyNum                int32
	UnitCost                 int32
}

type CustomGachaExFragmentMetaData struct {
	// Properties
	MaterialAmountRequest int32
	FragmentAmountByOnce  int32
	BonusGachaTimes       int32
	BonusRate             int32
}

type CustomHeadPageMetaData struct {
	// Properties
	PageDec AddrTo[StrWithPrefix16]
}

type CycleMissionManagementMetaData struct {
	// Properties
	HaveLink                bool
	PreDisplayCycleIDList   AddrTo[[]int32]
	RewardUnlockCycleIDList AddrTo[[]int32]
	RewardUnlockTime        AddrTo[StrWithPrefix16]
}

type CycleMissionMetaData struct {
	// Properties
	PreShowLevel int32
	MinLevel     int32
}

type DailyMissionMaterialIconMetaData struct {
	// Properties
	IconPath AddrTo[StrWithPrefix16]
}

type DailyMPStageDataMetaData struct {
	// Properties
	CoverBGPath   AddrTo[StrWithPrefix16]
	RewardDisplay int32
}

type DailyRechargeRewardGroupMetaData struct {
	// Properties
	RewardID    int32
	GroupDetail AddrTo[Hash]
}

type DailyRecommendMeta struct {
	// Properties
	TextMapTitle          AddrTo[StrWithPrefix16]
	SortID                uint8
	HangMissionList       AddrTo[[]int32]
	DefaultMissionDisplay int32
	InfoButtonIcon        AddrTo[StrWithPrefix16]
	InfoButtonFigure      AddrTo[StrWithPrefix16]
	TextMapTag            AddrTo[StrWithPrefix16]
	TagColor              AddrTo[StrWithPrefix16]
	LinkType              uint16
	LinkParams            AddrTo[[]int32]
	LinkParamStr          AddrTo[StrWithPrefix16]
	TextMapActivityTime   AddrTo[StrWithPrefix16]
	TextMapActivityType   AddrTo[StrWithPrefix16]
	TextMapLevelLimit     AddrTo[StrWithPrefix16]
	DisplayRewardItemList AddrTo[[]int32]
	DisplayMaterialList   AddrTo[[]StrWithPrefix16]
	MentoringRelationship uint8
	IsRetainGoBtn         bool
	ShowEndCountdown      bool
}

type DanmakuMetaData struct {
	// Properties
	MinLevel   uint16
	BeginTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
	SlotIDList AddrTo[[]uint16]
}

type DanmakuSlotMetaData struct {
	// Properties
	ChannelType uint8
	ChannelID   uint16
}

type DeviceFPSLimitMetaData struct {
	// Properties
	DisplayTextID           AddrTo[Hash]
	DeviceWhiteList         AddrTo[[]StrWithPrefix16]
	GraphicsDeviceWhiteList AddrTo[[]StrWithPrefix16]
	DeviceBlackList         AddrTo[[]StrWithPrefix16]
}

type DialogImageDataMetaData struct {
	// Properties
	Path      AddrTo[StrWithPrefix16]
	X         float32
	Y         float32
	Width     float32
	Height    float32
	Rotationy float32
	Title     AddrTo[Hash]
}

type DialogMetaData struct {
	// Properties
	// PostDialogIdList AddrTo[[]int32]
	PreDialogIDList AddrTo[[]int32]
	DialogType      uint8
	AvatarID        uint16
	ScreenSide      uint8
	Source          AddrTo[StrWithPrefix16]
	Content         AddrTo[[]DialogMetaData_PlotChatNode]
	AudioID         AddrTo[StrWithPrefix16]
	Emotion         AddrTo[StrWithPrefix16]
	LipMotion       AddrTo[StrWithPrefix16]
}

type DialogMetaData_PlotChatNode struct {
	// Objects
	ChatContent  Hash
	ChatDuration float32
}

type DiceyDungeonActivityMetaData struct {
	// Properties
	WeaponGachaMaterial    AddrTo[MechMetaData_DisjoinOutputItem]
	DiceyDungeonMedalID    int32
	DailyDungeonUnlockSite int32
}

type DiceyDungeonBubbleMetaData struct {
	// Properties
	TriggerCondition   uint8
	ConditionParameter AddrTo[[]int32]
	TriggerTimes       uint8
	TriggerPR          uint8
	Weight             uint16
	DialogTextList     AddrTo[[]Hash]
}

type DiceyDungeonBuffMetaData struct {
	// Properties
	IsHide   bool
	BuffIcon AddrTo[StrWithPrefix16]
	BuffName AddrTo[Hash]
	BuffDesc AddrTo[Hash]
}

type DiceyDungeonDailyScheduleMetaData struct {
	// Properties
	StartTime                    AddrTo[RemoteTime]
	EndTime                      AddrTo[RemoteTime]
	DungeonID                    int32
	MaxOrnamentNum               uint8
	MaxPassiveOrnamentNum        uint8
	DailyDungeonMaterialNumLimit AddrTo[[]DiceyDungeonDailyScheduleMetaData_MatLimitInfo]
}

type DiceyDungeonDailyScheduleMetaData_MatLimitInfo struct {
	// Fields
	ID       int32
	LimitNum int32
}

type DiceyDungeonDailyScoreMetaData struct {
	// Properties
	DungeonID            int32
	DailyDungeonMinScore int32
	DailyDungeonMaterial AddrTo[[]MechMetaData_DisjoinOutputItem]
}

type DiceyDungeonDataMetaData struct {
	// Properties
	DungeonName                AddrTo[Hash]
	DungeonType                uint8
	BeginDungeonRoomID         int32
	RoleRestrictList           AddrTo[[]int32]
	RoleRecommendList          AddrTo[[]int32]
	EnemyDisplayList           AddrTo[[]int32]
	OrnamentDisplayList        AddrTo[[]int32]
	TipDisplayList             AddrTo[[]Hash]
	OrnamentPoolID             int32
	FristDropRewardDisplayList int32
	RewardDisplayList          int32
	UnlockRoleList             AddrTo[[]DiceyDungeonDataMetaData_RoleInfo]
	UnlockWeaponList           AddrTo[[]DiceyDungeonDataMetaData_RoleInfo]
	RecoveryRoomHP             float32
	CheckEventHP               float32
	OrnamentChoiceHP           float32
	BgPath                     AddrTo[StrWithPrefix16]
	OrnamentRefreshNum         uint8
}

type DiceyDungeonDataMetaData_RoleInfo struct {
	// Fields
	ID    int32
	Level int32
}

type DiceyDungeonGachaPoolMetaData struct {
	// Properties
	GachaPoolType     uint8
	GachaPoolItemList AddrTo[[]DiceyDungeonGachaPoolMetaData_WeaponInfo]
	UnlockSite        int32
}

type DiceyDungeonGachaPoolMetaData_WeaponInfo struct {
	// Fields
	ID       int32
	Priority int32
}

type DiceyDungeonMonsterBehaviorMetaData struct {
	// Properties
	SkillID      int32
	BeginTurn    uint8
	CastTurn     uint8
	CastPriority uint8
	CastLimit    AddrTo[StrWithPrefix16]
}

type DiceyDungeonMonsterMetaData struct {
	// Properties
	MonsterName             AddrTo[Hash]
	MonsterIcon             AddrTo[StrWithPrefix16]
	IsBoss                  bool
	MonsterDetailText       AddrTo[Hash]
	MonsterFigurePath       AddrTo[StrWithPrefix16]
	MonsterPrefabPath       AddrTo[StrWithPrefix16]
	MonsterHP               int32
	BehaviorIDList          AddrTo[[]int32]
	MonsterPassiveSKillList AddrTo[[]int32]
	SkillDisplayType1       AddrTo[StrWithPrefix16]
	SkillDisplayType2       AddrTo[StrWithPrefix16]
	SkillDisplayType3       AddrTo[StrWithPrefix16]
}

type DiceyDungeonOrnamentMetaData struct {
	// Properties
	MaxLevel       int32
	DungeonSkillID AddrTo[[]int32]
	UnlockSite     int32
	BuffBriefDes   AddrTo[[]Hash]
}

type DiceyDungeonPassiveSkillMetaData struct {
	// Properties
	PassiveskillType uint8
	PassiveskillTag  uint8
	PassiveskillName AddrTo[Hash]
	PassiveskillDes  AddrTo[Hash]
	TriggerCondition AddrTo[StrWithPrefix16]
	TriggerParam     AddrTo[StrWithPrefix16]
	ResultOnExecute  AddrTo[[]StrWithPrefix16]
	IconPath         AddrTo[StrWithPrefix16]
	ShowDetailDes    bool
}

type DiceyDungeonRoleMetaData struct {
	// Properties
	RoleName                 AddrTo[Hash]
	RoleChibiIcon            AddrTo[StrWithPrefix16]
	RoleFigurePath           AddrTo[StrWithPrefix16]
	RolePrefabPath           AddrTo[StrWithPrefix16]
	MaxLevel                 int32
	LevelUpCostMaterial      AddrTo[MechMetaData_DisjoinOutputItem]
	Health                   int32
	Strength                 int32
	Agility                  int32
	Intelligence             int32
	DungeonSkillID           AddrTo[[]int32]
	SupportAbilityName       AddrTo[StrWithPrefix16]
	SupportParamList         AddrTo[[]float32]
	SupportSkillDesc         AddrTo[Hash]
	RecommendWeaponIDList    AddrTo[[]int32]
	SkillDisplayType1        AddrTo[StrWithPrefix16]
	SkillDisplayType2        AddrTo[StrWithPrefix16]
	SkillDisplayType3        AddrTo[StrWithPrefix16]
	DamageIncrease           float32
	BattleDicePool           AddrTo[[]MechMetaData_DisjoinOutputItem]
	OrnamentInitiativePoolID int32
	OrnamentPassivePoolID    int32
	UnlockSiteID             int32
	DailyRecommendList       AddrTo[[]int32]
}

type DiceyDungeonRoomMetaData struct {
	// Properties
	Floor                    uint8
	Position                 uint8
	NextRoomList             AddrTo[[]int32]
	DungeonRoomType          uint8
	DungeonRoomTypeParameter int32
	CheckEventType           uint8
	CheckEventParameter      int32
	EnterPlotID              int32
	PassPlotID               int32
	PassReward               int32
	DisplayDialogTitle       AddrTo[Hash]
	DisplayDialogContent     AddrTo[Hash]
	FristDropReward          int32
	CheckEventDesc           AddrTo[Hash]
	CheckEventOption         AddrTo[Hash]
}

type DiceyDungeonSiteDisplayMetaData struct {
	// Properties
	RoleID   int32
	RewardID int32
}

type DiceyDungeonSkillMetaData struct {
	// Properties
	SkillType         uint8
	SkillName         AddrTo[Hash]
	SkillDes          AddrTo[Hash]
	ResultOnExecute   AddrTo[[]StrWithPrefix16]
	SlotType          uint8
	SlotTypeParameter uint8
	UseTimeType       uint8
	UseTimeParameter  uint8
	DisplayType       uint8
	ShowDetailDes     bool
}

type DiceyDungeonTowerMetaData struct {
	// Properties
	FloorID         int16
	Medal           int16
	ShowTowerReward bool
}

type DiceyDungeonTutorialDataMetaData struct {
	// Properties
	TutorialDataID int32
}

type DiceyDungeonWeaponMetaData struct {
	// Properties
	WeaponIcon     AddrTo[StrWithPrefix16]
	MaxLevel       int32
	DungeonSkillID AddrTo[[]int32]
}

type DLC2DailyQuestInfoMetaData struct {
	// Properties
	QuestConsigner   AddrTo[Hash]
	QuestInfo        AddrTo[Hash]
	SpecialRewardID  int32
	SpecialQuestTime int32
}

type DLC2PlotUIConfigMetaData struct {
	// Properties
	ShowHistoryBtn bool
	ShowHideBtn    bool
	ShowAutoBtn    bool
	ShowSkipBtn    bool
}

type DLC2SupportGiftMetaData struct {
	// Properties
	SupportExp AddrTo[[]DLC2SupportGiftMetaData_SupportExpPair]
	UpNPC      AddrTo[[]uint8]
}

type DLC2SupportGiftMetaData_SupportExpPair struct {
	// Fields
	NpcID uint8
	Exp   int32
}

type DLCGachaTabMetaData struct {
	// Properties
	ActivityID int32
}

type DLCLevelDiffMetaData struct {
	// Properties
	AbilityName AddrTo[StrWithPrefix16]
	IconBg      AddrTo[StrWithPrefix16]
}

type DLCLevelMetaData struct {
	// Properties
	Exp                        int32
	Reward                     int32
	UnlockSysTips              AddrTo[[]Hash]
	UnlockTalentTips           AddrTo[[]DLCLevelMetaData_TalentTip]
	IsHighlight                bool
	MaterialDropLimitAddedList AddrTo[[]DLCLevelMetaData_ItemLimit]
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
	// Properties
	IsHidden        bool
	UnlockStoryID   int32
	HiddenList      AddrTo[[]int32]
	Name            AddrTo[Hash]
	Portrait        AddrTo[StrWithPrefix16]
	AvatarIcon      AddrTo[StrWithPrefix16]
	ParentTransName AddrTo[StrWithPrefix16]
	NPCPerfabPath   AddrTo[StrWithPrefix16]
	ActionName      AddrTo[StrWithPrefix16]
	Audio           AddrTo[StrWithPrefix16]
	AudioDelayTime  int32
	Lines           AddrTo[Hash]
	DialogStoryList AddrTo[[]int32]
	NPCSupportDesc  AddrTo[Hash]
	NPCSupportInfo  AddrTo[Hash]
}

type DLCRecommendLevelMetaData struct {
	// Properties
	RecommendLevel int16
	WarningLevel   int16
}

type DLCReviveCostMetaData struct {
	// Properties
	CostCoin int32
}

type DLCStoryMetaData struct {
	// Properties
	NPCID            int32
	PlotID           int32
	MissionTag       uint8
	UnlockDLCStoryID int32
	BlockDLCStoryID  int32
}

type DLCStoryPreviewMetaData struct {
	// Properties
	PreviewStoryID int32
}

type DLCSupportDataMetaData struct {
	// Properties
	NPCID        uint8
	SupportPoint uint16
	// SupportTypeEnum int32

}

type DLCSupportLevelDataMetaData struct {
	// Properties
	UpgradeSupportPoint int16
	IconPath            AddrTo[StrWithPrefix16]
}

type DLCSupportRewardMetaData struct {
	// Properties
	PlotID int32
}

type DLCTalentAffixMetaData struct {
	// Properties
	PropID    uint16
	ValueMin  float32
	ValueMax  float32
	ValueStep float32
}

type DLCTalentAffixSetMetaData struct {
	// Properties
	Name            AddrTo[Hash]
	Ability1Name    AddrTo[StrWithPrefix16]
	Prop1Desc       AddrTo[Hash]
	Prop1ParamName1 AddrTo[StrWithPrefix16]
	Prop1Param1     float32
	Prop1ParamName2 AddrTo[StrWithPrefix16]
	Prop1Param2     float32
	Prop1ParamName3 AddrTo[StrWithPrefix16]
	Prop1Param3     float32
	Ability2Name    AddrTo[StrWithPrefix16]
	Prop2Desc       AddrTo[Hash]
	Prop2ParamName1 AddrTo[StrWithPrefix16]
	Prop2Param1     float32
	Prop2ParamName2 AddrTo[StrWithPrefix16]
	Prop2Param2     float32
	Prop2ParamName3 AddrTo[StrWithPrefix16]
	Prop2Param3     float32
}

type DLCTalentLevelMetaData struct {
	// Properties
	LevelUpPreTalent    AddrTo[[]DiceyDungeonDataMetaData_RoleInfo]
	LevelUpDLCLv        int32
	LevelUpMaterialList AddrTo[[]ElfMetaData_MatCost]
	RefreshMaterialList AddrTo[[]ElfMetaData_MatCost]
	HPAdd               float32
	SPAdd               float32
	ATKAdd              float32
	AbilityParamBase_1  float32
	AbilityParamBase_2  float32
	AbilityParamBase_3  float32
	RelateAvatarList    AddrTo[[]DLCTalentLevelMetaData_RelateAvatar]
	TalentInfo          AddrTo[Hash]
	TalentInfo2         AddrTo[Hash]
}

type DLCTalentLevelMetaData_RelateAvatar struct {
	// Fields
	AvatarID int32
	Star     int32
}

type DLCTalentMetaData struct {
	// Properties
	AvatarID                 int32
	TalentType               int32
	UniqueTag                int32
	TalentName               AddrTo[Hash]
	TalentShortName          AddrTo[Hash]
	IconPath                 AddrTo[StrWithPrefix16]
	UIPointName              AddrTo[StrWithPrefix16]
	UILineName               AddrTo[[]StrWithPrefix16]
	CGID                     int32
	IsNeedSwitch             bool
	UpgradeDisplayType       uint8
	UpgradeAdditionalDisplay AddrTo[Hash]
	IsMultiDisplay           bool
	DisplayPriority          uint8
}

type DLCTalentMultiDisplayMetaData struct {
	// Properties
	DLCTalentID              int32
	TalentLevel              uint8
	CGID                     int32
	DisplayPriority          uint8
	TalentInfo               AddrTo[Hash]
	TalentInfo2              AddrTo[Hash]
	UpgradeAdditionalDisplay AddrTo[Hash]
}

type DLCTalentTagMetaData struct {
	// Properties
	Weight int32
}

type DLCTowerBonusMetaData struct {
	// Properties
	BuffList AddrTo[[]int32]
	Name     AddrTo[StrWithPrefix16]
	Desc     AddrTo[StrWithPrefix16]
	IconPath AddrTo[StrWithPrefix16]
}

type DLCTowerBuffMetaData struct {
	// Properties
	BuffName    AddrTo[StrWithPrefix16]
	OverrideMap AddrTo[[]StrWithPrefix16]
	BuffType    int32
}

type DLCTowerFloorMetaData struct {
	// Properties
	RecommendLevel int32
	WarningLevel   int32
	HardLevel      int32
	SpawnType      int32
	SpawnTypeDesc  AddrTo[StrWithPrefix16]
	Parameter      AddrTo[[]int32]
	WaveListPool   AddrTo[[]StrWithPrefix16]
	TimeBonus      int32
	IsCheckPoint   int32
	Reward         int32
	StageSection   AddrTo[StrWithPrefix16]
	HardLevelGroup int32
	MaxCoin        int32
}

type DLCTowerMonsterMetaData struct {
	// Properties
	AddTime int32
	AddCoin int32
}

type DLCTowerRankMetaData struct {
	// Properties
	ScheduleReward AddrTo[[]DLCTowerRankMetaData_RewardData]
}

type DLCTowerRankMetaData_RewardData struct {
	// Fields
	ConfigType int32
	RewaredID  int32
}

type DLCTowerScheduleMetaData struct {
	// Properties
	NextID     int32
	StartTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
	SettleTime AddrTo[StrWithPrefix16]
	ConfigType int32
	LevelID    int32
	MaxFloor   int32
	Weather    int32
	Bonus      AddrTo[[]int32]
}

type DLCTowerStyleBonusMetaData struct {
	// Properties
	Rate float32
}

type DLCTowerWaveMetaData struct {
	// Properties
	MonsterConfig AddrTo[[]int32]
}

type DormitoryAvatarMetaData struct {
	// Properties
	RoleID               int32
	Name                 AddrTo[StrWithPrefix16]
	QAvatarIcon          AddrTo[StrWithPrefix16]
	NameText             AddrTo[Hash]
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
}

type DormitoryComfortMetaData struct {
	// Properties
	NeedComfortValues          int32
	PowerCostBonus             int32
	ComfortTitleText           AddrTo[Hash]
	ComfortIcon                AddrTo[StrWithPrefix16]
	ComfortIconColor           AddrTo[StrWithPrefix16]
	EnergyCost                 int32
	AdventureGrainRecoverBonus int32
	StaminaReward              int32
}

type DormitoryDecorationEffectMetaData struct {
	// Properties
	EffectType    int32
	EffectPara    int32
	EffectSubPara int32
}

type DormitoryDecorationMetaData struct {
	// Properties
	TargetRoom           int32
	DecorationPhase      int32
	DecorationStepShow   int32
	DecorationMaxStep    int32
	NeedComfort          int32
	DecorationEffectList AddrTo[[]int32]
	NeedTime             int32
	MaterialList         AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
}

type DormitoryDialogMetaData struct {
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
	DialogText    AddrTo[Hash]
}

type DormitoryEventDialogMetaData struct {
	// Properties
	EventKey int32
	RoleID   int32
	AvatarID int32
	TextMap  AddrTo[Hash]
	Weight   float32
}

type DormitoryEventSequenceMetaData struct {
	// Properties
	Type             AddrTo[StrWithPrefix16]
	Avatar           int8
	Wait             float32
	WaitRandomAdd    float32
	TalkPop          AddrTo[StrWithPrefix16]
	TalkTxt          AddrTo[StrWithPrefix16]
	TalkToAvatar     int8
	FaceAnimStop     bool
	FaceAnimType     AddrTo[StrWithPrefix16]
	TriggerAction    AddrTo[StrWithPrefix16]
	TriggerSubAction AddrTo[StrWithPrefix16]
}

type DormitoryEventWeightMetaData struct {
	// Properties
	RoleID   int32
	AvatarID int32
	EventID  int32
	Weight   float32
}

type DormitoryFacilityBonusDropMetaData struct {
	// Properties
	BeginTime     AddrTo[StrWithPrefix16]
	EndTime       AddrTo[StrWithPrefix16]
	DisplayList   AddrTo[[]int32]
	FacilityType  int32
	FacilityLevel int32
}

type DormitoryFacilityMetaData struct {
	// Properties
	FacilityType     int32
	HouseID          int32
	FacilityModPath  AddrTo[StrWithPrefix16]
	SwitchOrder      int32
	FacilityNameText AddrTo[Hash]
}

type DormitoryFurnitureCollectRewardMetaData struct {
}

type DormitoryFurnitureSuitMetaData struct {
	// Properties
	Rarity            int32
	SuitType          int32
	UIOrder           int32
	SuitNameText      AddrTo[Hash]
	SuitNameEn        AddrTo[Hash]
	SuitVersionText   AddrTo[Hash]
	SuitIcon          AddrTo[StrWithPrefix16]
	SmallSuitIcon     AddrTo[StrWithPrefix16]
	RewardMissionList AddrTo[[]int32]
}

type DormitoryHouseMetaData struct {
	// Properties
	TotalRoomList          AddrTo[[]int32]
	HouseAssetPath         AddrTo[StrWithPrefix16]
	HouseFacilityList      AddrTo[[]DormitoryHouseMetaData_Facility]
	UnlockLv               int32
	UnlockMaterial         int32
	HouseNameText          AddrTo[Hash]
	HouseTypeText          AddrTo[Hash]
	HouseEnterImgPath      AddrTo[StrWithPrefix16]
	HouseEnterImgLineColor AddrTo[StrWithPrefix16]
	HouseIconPath          AddrTo[StrWithPrefix16]
	HouseDescText          AddrTo[Hash]
	HouseIntro             AddrTo[StrWithPrefix16]
}

type DormitoryHouseMetaData_Facility struct {
	// Fields
	Col int32
	Dir int32
	Id  int32
	Row int32
}

type DormitoryInteractEventMetaData struct {
	// Properties
	InteractEvent int32
	TouchExp      int32
}

type DormitoryRoomMetaData struct {
	// Properties
	RoomType         int32
	SizeType         int32
	BornPositionX    float32
	BornPositionY    float32
	BornPositionZ    float32
	RoomSquare       int32
	AvatarLimit      int32
	RoomNameText     AddrTo[Hash]
	RoomUnlockText   AddrTo[Hash]
	InitialLimitList AddrTo[[]int32]
}

type DormitoryTalkEventMetaData struct {
	// Properties
	StartDialogId      int32
	FinishDialogIdList AddrTo[[]int32]
	Reward             int32
	TouchExp           int32
}

type DormitoryVoiceMetaData struct {
	// Properties
	AudioID      AddrTo[StrWithPrefix16]
	RoleID       int32
	FromHour     uint8
	ToHour       uint8
	AudioPattern AddrTo[StrWithPrefix16]
	Weight       uint8
}

type DownLoadConfigMetaData struct {
	// Properties
	AccountType   AddrTo[[]uint8]
	Icontype      uint8
	Icon          AddrTo[StrWithPrefix16]
	Icontext      AddrTo[Hash]
	LinkUrl       AddrTo[StrWithPrefix16]
	Isshow        bool
	SupplyImgPath AddrTo[StrWithPrefix16]
	Rect          AddrTo[DownLoadConfigMetaData_RectTransformData]
	HaveRect      bool
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
	// Properties
	BackGroundType   uint8
	BackGroundInfo   AddrTo[StrWithPrefix16]
	ProgressBarColor AddrTo[StrWithPrefix16]
	ActionsRect      AddrTo[DownLoadConfigMetaData_RectTransformData]
	ProgressBarRect  AddrTo[DownLoadConfigMetaData_RectTransformData]
}

type DreamMetaData struct {
	// Properties
	DreamRewardList AddrTo[[]int32]
	ScoreMaterial   int32
}

type DreamMissionMetaData struct {
	// Properties
	Score             int32
	MinLevel          int32
	MaxLevel          int32
	BeginTime         AddrTo[RemoteTime]
	EndTime           AddrTo[RemoteTime]
	DisplayPreMission int32
}

type DreamRewardMetaData struct {
	// Properties
	TitleName        AddrTo[Hash]
	Description      AddrTo[Hash]
	ImagePath        AddrTo[StrWithPrefix16]
	BgPath           AddrTo[StrWithPrefix16]
	RewardID         int32
	MissionList      AddrTo[[]int32]
	Score            int32
	ExchangeOpen     bool
	ExchangeOpentime AddrTo[StrWithPrefix16]
	ExchangeEndtime  AddrTo[StrWithPrefix16]
	ExchangeScore    int32
	ExchangeCost     AddrTo[[]int32]
	RewardShowNum    uint8
}

type DreamUnlockMetaData struct {
	// Properties
	UnlockHint AddrTo[Hash]
}

type DressTagDataMetaData struct {
	// Properties
	TagName AddrTo[Hash]
	TagDesc AddrTo[Hash]
	TagIcon AddrTo[StrWithPrefix16]
}

type DungeonsMirrorRecoveryMetaData struct {
	// Properties
	Rarity    int32
	TicketID  int32
	TicketNum float32
}

type DutyDailyMetaData struct {
	// Properties
	NeedDuty int32
	UnlockLv int32
	RewardId int32
}

type DynamicHardLvMetaData struct {
	// Properties
	HardLv int32
}

type ElevatorConfigMetaData struct {
	// Properties
	ResPath AddrTo[StrWithPrefix16]
}

type ElfLevelMetaData struct {
	// Properties
	Exp int32
}

type ElfStoryActMetaData struct {
	// Properties
	TriggerName            AddrTo[StrWithPrefix16]
	TriggerEffectList      AddrTo[[]ElfGalEventMetaData_ElfUIEffectPattern]
	FaceAnimationKey       AddrTo[StrWithPrefix16]
	FaceAnimationDelayTime int32
}

type ElfTalentLevelMetaData struct {
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
	// Properties
	Name        AddrTo[Hash]
	Info        AddrTo[Hash]
	TalentGroup int32
	TalentCost  int32
	MaxLv       int32
	IconPath    AddrTo[StrWithPrefix16]
	SortWeight  int32
}

type ElfTrialMetaData struct {
	// Properties
	ElfID          uint8
	ElfLV          uint8
	ElfStar        uint8
	EquipSkillList AddrTo[[]ElfTrialMetaData_SkillData]
}

type ElfTrialMetaData_SkillData struct {
	// Fields
	ID int32
	LV int32
}

type EliteAbilityMetaData struct {
	// Properties
	EliteAbilityDesc AddrTo[Hash]
	EliteAbilityIcon AddrTo[StrWithPrefix16]
	EliteAbilityText AddrTo[Hash]
	EliteAbilityTag  AddrTo[StrWithPrefix16]
}

type EliteChapterMetaData struct {
	// Properties
	SortID              int32
	ChapterType         int32
	ContentID           int32
	UnlockLevel         int32
	PreviewLevel        int32
	ChaperImgPath       AddrTo[StrWithPrefix16]
	StageImgPath        AddrTo[StrWithPrefix16]
	RewardShow          int32
	EliteChapterName    AddrTo[StrWithPrefix16]
	KeyItemIDList       AddrTo[[]int32]
	KeyResourceTypeList AddrTo[[]int32]
}

type EliteStageCompensationMetaData struct {
	// Properties
	FirstDropRewardID int32
}

type EliteStageMetaData struct {
	// Properties
	StageName         AddrTo[StrWithPrefix16]
	SortID            int32
	EliteChapterID    int32
	PreStageID        int32
	UnlockLevel       int32
	UnlockMainStoryID int32
}

type EmojiFaceDataMetaData struct {
	// Properties
	FaceDec    AddrTo[Hash]
	Page       int32
	Pic        AddrTo[StrWithPrefix16]
	UnlockType int32
	UnlockDec  AddrTo[Hash]
	Show       int32
}

type EmojiFacePageMetaData struct {
	// Properties
	ShowType   AddrTo[[]uint8]
	PageDec    AddrTo[Hash]
	PageNameEn AddrTo[Hash]
	Pic        AddrTo[StrWithPrefix16]
}

type EntryPerformMetaData struct {
	// Properties
	EventList       AddrTo[[]uint16]
	EventListFollow AddrTo[[]uint16]
}

type EntryThemeDataMetaData struct {
	// Properties
	Display                   bool
	SpecialTag                bool
	ThemeUIConfig             int32
	RandomThemeBGM            int32
	SpecialConfigList         AddrTo[[]int32]
	InitialThemeBGMConfigList AddrTo[[]int32]
	FullThemeBGMConfigList    AddrTo[[]int32]
	WeatherGroupConfig        uint8
	ThemeName                 AddrTo[Hash]
	ThemeDesc                 AddrTo[Hash]
	ThemeGetDesc              AddrTo[Hash]
	ThemeBGSmall              AddrTo[StrWithPrefix16]
	ThemeBGLarge              AddrTo[StrWithPrefix16]
	ThemeTagList              AddrTo[[]uint16]
	GiftDesc                  AddrTo[Hash]
	GiftBGPath                AddrTo[StrWithPrefix16]
}

type EntryThemeMetaData struct {
	// Properties
	ElevatorConfig  int32
	SpaceShipConfig int32
	ThemeUIConfig   int32
	ThemeBGMConfig  int32
}

type EntryThemeScheduleMetaData struct {
	// Properties
	BeginTime AddrTo[RemoteTime]
	EndTime   AddrTo[RemoteTime]
	ThemeID   uint8
}

type EntryThemeTagDataMetaData struct {
	// Properties
	ThemeTagName AddrTo[Hash]
	ThemeTagDesc AddrTo[Hash]
	ThemeTagIcon AddrTo[StrWithPrefix16]
}

type EquipFilterTagMetaData struct {
	// Properties
	TagName AddrTo[Hash]
	Icon    AddrTo[StrWithPrefix16]
}

type EquipForgeExchangeMetaData struct {
	// Properties
	GoodsID int32
}

type EquipForgeGenerationMetaData struct {
	// Properties
	Priority        int32
	IconPath        AddrTo[StrWithPrefix16]
	NameTextMap     AddrTo[Hash]
	BackGroundColor AddrTo[StrWithPrefix16]
	TagDisplay      AddrTo[StrWithPrefix16]
}

type EquipForgeKeepAffixMetaData struct {
	// Properties
	ConsumeNumList AddrTo[[]AvatarSubSkillLevelMetaData_SkillUpLevelNeedItem]
}

type EquipForgeMeta struct {
	// Properties
	TargetId            int32
	TargetInitLevel     int32
	CurrencyShow        AddrTo[[]int32]
	CurrencyList        AddrTo[[]EquipForgeMeta_ItemMaterial]
	ForgeParaStr        AddrTo[[]EquipForgeMeta_EFCostMatItem]
	PreEquipmentList    AddrTo[[]EquipForgeMeta_WeaponMaterial]
	MaxForgeType        int32
	MaxBuyTimes         int32
	EquipmentType       int32
	SortType            int32
	Evolution           bool
	SortPriority        int32
	RefreshTimeType     int32
	DisplayMinLevel     int32
	MinLevel            int32
	MaxLevel            int32
	UnlockMainIDList    AddrTo[[]int32]
	Generation          int32
	InferiorEquip       int32
	ForgeParaStrDisplay AddrTo[[]int32]
	IsProgressShow      bool
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
	// Properties
	MinLevel  int32
	MaxLevel  int32
	ForgeList AddrTo[[]int32]
}

type EquipForgeShadowMetaData struct {
	// Properties
	TargetMainID               int32
	TypeIconStigmataIDRelation int32
}

type EquipmentActivitySetMetaData struct {
	// Properties
	SetName          AddrTo[Hash]
	SetNum           int32
	DisPlayList      AddrTo[[]int32]
	IsShow           bool
	BeginShowingTime AddrTo[StrWithPrefix16]
	SetIcon          AddrTo[StrWithPrefix16]
	SourceTextMap    AddrTo[Hash]
	RewardType       int32
	Reward           int32
}

type EquipmentLevelMetaData struct {
	// Properties
	Type1               AddrTo[[]int32]
	WeaponUpgradeCost   int32
	WeaponEvoCost       int32
	StigmataUpgradeCost int32
	StigmataEvoCost     int32
}

type EquipTypeDetailMetaData struct {
	// Properties
	PicPath AddrTo[StrWithPrefix16]
}

type EvaluateDialogMetaData struct {
	// Properties
	DialogTitle   AddrTo[Hash]
	DialogContent AddrTo[Hash]
	DialogAgree   AddrTo[Hash]
	DialogRefuse  AddrTo[Hash]
	DialogClose   AddrTo[Hash]
	DialogPic     AddrTo[StrWithPrefix16]
}

type EvaluateIntroMetaData struct {
	// Properties
	IsEnable         int32
	ChannelList      AddrTo[[]StrWithPrefix16]
	Param1           AddrTo[[]int32]
	Param2           AddrTo[[]int32]
	Param3           AddrTo[[]int32]
	EvaluateDialogID int32
}

type EventCollectionMetaData struct {
	// Properties
	SysType         uint8
	CollectionType  uint8
	CollectionSubID int32
	UnlockType      uint8
	UnlockParam     int32
}

type EventDialogDataMetaData struct {
	// Properties
	PreDialogIdList      AddrTo[[]int32]
	JumpID               int32
	LeafIDList           AddrTo[[]int32]
	DialogType           uint8
	InputID              int32
	CGRawPos             int32
	AvatarName           AddrTo[StrWithPrefix16]
	AvatarId             int32
	DressId              int32
	AvatarViceKey        int32
	ScreenSide           uint8
	Face                 uint8
	FaceAnimation        AddrTo[StrWithPrefix16]
	FaceExpression       AddrTo[StrWithPrefix16]
	FaceEffect           AddrTo[StrWithPrefix16]
	AnimID               uint8
	Distortion           uint8
	Transparency         float32
	BgmCover             AddrTo[StrWithPrefix16]
	BGMVolumeControlList AddrTo[[]StrWithPrefix16]
	CgId                 AddrTo[StrWithPrefix16]
	ImageId              int32
	Content              AddrTo[[]EventDialogDataMetaData_PlotChatNode]
	Background           AddrTo[StrWithPrefix16]
	BackgroundCG         AddrTo[StrWithPrefix16]
	Backgroundeffect     uint8
	EnterEffect          uint8
	AudioId              AddrTo[StrWithPrefix16]
	LaterForAudio        float32
	LipMotion            AddrTo[StrWithPrefix16]
	FaceVersion          uint8
	ScreenEffect         AddrTo[StrWithPrefix16]
	DialogueSubTitle     AddrTo[StrWithPrefix16]
	ImagePathList        AddrTo[[]StrWithPrefix16]
	ImageDescList        AddrTo[[]StrWithPrefix16]
	DialogueDesignType   int32
	ImageChar2DID        uint32
	PostDialogIdList     AddrTo[[]int32]
	TalkerName           AddrTo[StrWithPrefix16]
	QuestionContent      AddrTo[StrWithPrefix16]
}

type EventDialogDataMetaData_PlotChatNode struct {
	// Objects
	ChatContent  StrWithPrefix16
	ChatDuration float32
}

type ExaminationAnswerMetaData struct {
	// Properties
	TextMap AddrTo[Hash]
}

type ExaminationQuestionMetaData struct {
	// Properties
	QuestionTextMap AddrTo[Hash]
	WrongTipTextMap AddrTo[Hash]
	CorrectAnswerID int32
	HeadImgPath     AddrTo[StrWithPrefix16]
	LinkType        int32
	LinkPram        AddrTo[[]int32]
	LinkParamStr    AddrTo[StrWithPrefix16]
}

type ExaminationRewardMetaData struct {
	// Properties
	RewardID int32
}

type ExaminationScheduleMetaData struct {
	// Properties
	ExaminationRewardConfigID int32
}

type ExBossMonsterLevelConfigMetaData struct {
	// Properties
	TurboModeBuffName AddrTo[StrWithPrefix16]
	TimeScoreK        int32
	BaseScoreK        int32
	TurboModeDesc     AddrTo[StrWithPrefix16]
	ContinueChallenge bool
	UnlockWindow      bool
}

type ExBossMonsterScoreRewardMetaData struct {
	// Properties
	Score      int32
	RewardID   int32
	RewardDesc AddrTo[Hash]
}

type ExBossMonsterWeatherMetaData struct {
	// Properties
	IconPath     AddrTo[StrWithPrefix16]
	WeatherTitle AddrTo[Hash]
	WeatherDesc  AddrTo[Hash]
}

type ExBossTransferInfoMetaData struct {
	// Properties
	BossID int32
}

type ExhibitionMetaData struct {
	// Properties
	MuseumID       uint8
	ItemType       uint8
	ItemID         int32
	LinkType       uint16
	LinkParams     AddrTo[[]int32]
	LinkParamStr   AddrTo[StrWithPrefix16]
	LinkColor      AddrTo[StrWithPrefix16]
	SpecialTextmap AddrTo[StrWithPrefix16]
}

type ExposureRateDataMetaData struct {
	// Properties
	ExposureAreaRate int32
	ShopGoodsList    AddrTo[[]int32]
}

type ExScheduleTextmapListMetaData struct {
	// Properties
	TextmapList AddrTo[[]int32]
}

type ExScheduleTextmapMetaData struct {
	// Properties
	Begintime  AddrTo[RemoteTime]
	Endtime    AddrTo[RemoteTime]
	Newtextmap AddrTo[StrWithPrefix16]
}

type ExtraShortStoryMetaData struct {
	// Properties
	StoryModeChapter      int32
	ChallangeModeChapter  int32
	ShortStoryDescTitle   AddrTo[Hash]
	ShortStoryDescContent AddrTo[Hash]
}

type ExtraStoryAchieveDisplayMetaData struct {
	// Properties
	DisplayName    AddrTo[Hash]
	DisplayNameEng AddrTo[Hash]
	DisplayPicPath AddrTo[StrWithPrefix16]
	AreaID         int32
	Type           int32
	ChapterID      int32
	GroupIDList    AddrTo[[]int32]
	PreAreaID      int32
}

type ExtraStoryAchieveGroupyMetaData struct {
	// Properties
	RewardID   int32
	GroupType  int32
	AreaID     int32
	Difficulty int32
	ChapterID  int32
}

type ExtraStoryAchieveMetaData struct {
	// Properties
	GroupId         int32
	AchieveLocation int32
}

type ExtraStoryActMetaData struct {
	// Properties
	ActType             int32
	AreaId              int32
	StageIdList         AddrTo[[]int32]
	Location            int32
	NameTextId          AddrTo[Hash]
	CgID                int32
	EnhanceOn           int32
	EnhanceReward       int32
	EnhanceRestrictList AddrTo[[]int32]
}

type ExtraStoryAreaMetaData struct {
	// Properties
	ChapterId         int32
	WithMap           int32
	NameTextId        AddrTo[Hash]
	RogueEntryPicPath AddrTo[StrWithPrefix16]
}

type ExtraStoryChallengeModeMetaData struct {
	// Properties
	DifficultyText         AddrTo[Hash]
	DifficultyUnlockText   AddrTo[Hash]
	RewardPreview          int32
	ActRewardShow          int32
	UnlockTeamLv           int32
	FirstAreaID            int32
	AreaList               AddrTo[[]int32]
	RecTeamLv              AddrTo[[]int32]
	RecGroupText           AddrTo[Hash]
	MaxAvatarNumberPerArea int32
}

type ExtraStoryChapterMetaData struct {
	// Properties
	MinPlayerLevel     int32
	CurrenceIDList     AddrTo[[]int32]
	NameTextId         AddrTo[Hash]
	EntryPicPath       AddrTo[StrWithPrefix16]
	InfoTextID         AddrTo[Hash]
	ChapterType        int32
	ShopType           int32
	IsFeatured         int32
	AvatarSampleSwitch bool
}

type ExtraStoryLineMetaData struct {
	// Properties
	StoryChapterId      int32
	Weight              int32
	ActivityWeight      int32
	ActivityBeginTime   AddrTo[StrWithPrefix16]
	ActivityEndTime     AddrTo[StrWithPrefix16]
	ActivityText        AddrTo[Hash]
	ChallengeChapterId  int32
	UnlockStageId       int32
	CGGroup             AddrTo[[]int32]
	StoryLineName       AddrTo[Hash]
	StoryFinshReward    int32
	RewardShow          int32
	ChallengeRewardShow int32
	UnlockText          AddrTo[Hash]
	StoryLineDesc       AddrTo[Hash]
	CgStoryBg           AddrTo[StrWithPrefix16]
	CgStoryLinkType     int32
	CgStoryLink         AddrTo[StrWithPrefix16]
	EnhanceUnlockLevel  int32
}

type ExtraStoryThemeData struct {
	// Properties
	ThemeName           AddrTo[Hash]
	ThemeBuffList       AddrTo[[]int32]
	StigmataGroupList   AddrTo[[]int32]
	StigmataGroupBuff1  int32
	StigmataGroupBuff2  int32
	StigmataGroupBuff3  int32
	StigmataGroup2Buff1 int32
	StigmataGroup2Buff2 int32
	StigmataGroup2Buff3 int32
	ThemeAvatarList     AddrTo[[]int32]
	ThemeRoleList       AddrTo[[]int32]
	ThemeAvatarBuff     int32
	ThemeAvatarTitle    AddrTo[Hash]
	ThemeImg            AddrTo[StrWithPrefix16]
	ThemeDisplay        AddrTo[Hash]
}

type ExtraStoryThemeSchedule struct {
	// Properties
	BeginTime   AddrTo[StrWithPrefix16]
	EndTime     AddrTo[StrWithPrefix16]
	ChapterList AddrTo[[]int32]
	ThemeID     int32
}

type FakeAvatarMetaData struct {
	// Properties
	FakeAvatarName AddrTo[StrWithPrefix16]
	IconPath       AddrTo[StrWithPrefix16]
	DesTextMap     AddrTo[Hash]
}

type FarmBuffScheduleMetaData struct {
	// Properties
	BuffIcon        AddrTo[StrWithPrefix16]
	BeginTime       AddrTo[StrWithPrefix16]
	BuffSpeedUpTime int32
}

type FarmEventMetaData struct {
	// Properties
	UnlockEventExp int32
}

type FarmLevelMetaData struct {
	// Properties
	RewardId      int32
	UnlockFarmExp int32
	MaxSlotNum    int32
}

type FarmMaterialMetaData struct {
	// Properties
	EveryNum        int32
	LimitTimes      int32
	CostTime        int32
	UnlockFarmLevel int32
	RewardFarmExp   int32
	FoodIcon        AddrTo[StrWithPrefix16]
	FoodDesc        AddrTo[Hash]
	FoodNameEn      AddrTo[Hash]
}

type FarmScheduleMetaData struct {
	// Properties
	BeginTime        AddrTo[StrWithPrefix16]
	ProduceEndTime   AddrTo[StrWithPrefix16]
	EndTime          AddrTo[StrWithPrefix16]
	DisplayEndTime   AddrTo[StrWithPrefix16]
	BeginDayTime     AddrTo[StrWithPrefix16]
	EndDayTime       AddrTo[StrWithPrefix16]
	MissionList      AddrTo[[]int32]
	MinLevel         int32
	MaxLevel         int32
	StagePath        AddrTo[StrWithPrefix16]
	CGList           AddrTo[[]StrWithPrefix16]
	SpeedUpMaterLink AddrTo[StrWithPrefix16]
	PointRewardID    int32
	PointRewardLevel int32
	SfxActionStart   AddrTo[StrWithPrefix16]
	SfxActionLoop    AddrTo[StrWithPrefix16]
	UISkinList       AddrTo[[]int32]
	CameraPosition   AddrTo[[]int32]
	CameraRotation   AddrTo[[]int32]
}

type FarmSlotMetaData struct {
	// Properties
	CostMaterialStr AddrTo[StrWithPrefix16]
	UnlockFarmLevel int32
	Position        AddrTo[[]float32]
	Rotation        AddrTo[[]float32]
	Scale           float32
}

type FarmSpeedUpMetaData struct {
	// Properties
	SpeedUpTime int32
}

type FarmUISkinMetaData struct {
	// Properties
	UIType  int32
	UIPath  AddrTo[StrWithPrefix16]
	ResPath AddrTo[StrWithPrefix16]
	ResType int32
	Color   AddrTo[StrWithPrefix16]
}

type FastPassLevelMetaData struct {
	// Properties
	ID int32
}

type FastPassMetaData struct {
	// Properties
	MaterialCost uint8
	MaterialID   uint16
	BeginTime    AddrTo[StrWithPrefix16]
	EndTime      AddrTo[StrWithPrefix16]
	LimitType    uint8
	LimitNum     uint8
	TextID       AddrTo[Hash]
	MaterialLink AddrTo[StrWithPrefix16]
}

type FeaturedContentMetaData struct {
	// Properties
	Title                 AddrTo[Hash]
	Picture               AddrTo[StrWithPrefix16]
	PreviewType           uint8
	PreviewContent        AddrTo[StrWithPrefix16]
	Type                  uint8
	ShareContent          AddrTo[StrWithPrefix16]
	SpareType             uint8
	SpareShareContent     AddrTo[StrWithPrefix16]
	ShareChannelContentID int32
	BeginTime             AddrTo[RemoteTime]
	EndTime               AddrTo[RemoteTime]
	Period                uint8
	MinLevel              uint8
	MaxLevel              uint8
	DisabledChannel       AddrTo[[]int32]
}

type FeatureSubPakConfigMetaData struct {
	// Properties
	SubPakList AddrTo[[]int32]
}

type FixedPlotUIConfigMetaData struct {
	// Properties
	ShowHistoryBtn bool
	ShowHideBtn    bool
	ShowAutoBtn    bool
	ShowSkipBtn    bool
}

type FlopActivityPanel struct {
	// Properties
	CostItemID          int32
	NoCostTimes         int32
	CostItemNum         int32
	DailyMaxFlopTimes   int32
	AssitentActiveTimes int32
	RewardList          AddrTo[[]int32]
	ImagePath           AddrTo[[]StrWithPrefix16]
	ImagePathBackground AddrTo[StrWithPrefix16]
}

type FoundationRewardMetaData struct {
	// Properties
	Reward  int32
	Textmap AddrTo[Hash]
}

type FoundationTypeMetaData struct {
	// Properties
	BGpic     AddrTo[StrWithPrefix16]
	BGTitle   AddrTo[Hash]
	Smallpic  AddrTo[StrWithPrefix16]
	Smalltext AddrTo[Hash]
	Reference AddrTo[Hash]
	Minlevel  int32
	Maxlevel  int32
}

type FOWEffectMetaData struct {
	// Properties
	HeightRange       float32
	XSize             int32
	ZSize             int32
	TexWidth          int32
	TexHeigth         int32
	PositionWS        AddrTo[[]float32]
	RotationWS        AddrTo[[]float32]
	CutOff            float32
	Luminance         float32
	StepScale         float32
	StepMinValue      float32
	ImagePath         AddrTo[StrWithPrefix16]
	PlayerSectorAngle float32
	PlayerRadius      float32
	FogValue          float32
}

type FrontEndlessBattleConfigMetaData struct {
	// Properties
	WeatherID           int32
	MechanismID         int32
	MonsterTable        AddrTo[[]int32]
	StageMonsterDisplay AddrTo[[]int32]
	TimeLimit           int32
	ChallengeTime       int32
	FloorType           int32
}

type FrontEndlessFloorConfigMetaData struct {
	// Properties
	BaseScore               int32
	FloorDisplayType        uint8
	FrontEndlessFloorReward int32
	FloorName               AddrTo[Hash]
	FloorEntrancePath       AddrTo[StrWithPrefix16]
	ChallengeRewardDisplay  int32
	ChallengeFloorName      AddrTo[Hash]
	RankIconPath            AddrTo[StrWithPrefix16]
	RankTitle               AddrTo[Hash]
	RankDesc                AddrTo[Hash]
	IsAutoContinue          uint8
}

type FrontEndlessStageMetaData struct {
	// Properties
	StageID int32
}

type FurnitureTypeMetaData struct {
	// Properties
	ListType     int32
	UIOrder      int32
	TypeNameText AddrTo[Hash]
	TypeIcon     AddrTo[StrWithPrefix16]
}

type GachaAdventureDisplayMetaData struct {
	// Properties
	AvatarImg    AddrTo[StrWithPrefix16]
	AvatarWeight int32
}

type GachaBoxConfigMetaData struct {
	// Properties
	BoxModel                AddrTo[StrWithPrefix16]
	BoxModelMaterialList    AddrTo[[]StrWithPrefix16]
	StageModel              AddrTo[StrWithPrefix16]
	StageOffset             AddrTo[[]float32]
	StageModelMaterialList  AddrTo[[]StrWithPrefix16]
	CameraAnim              AddrTo[StrWithPrefix16]
	BoxModelPathInfoOneDraw AddrTo[StrWithPrefix16]
	BoxModelPathInfoTenDraw AddrTo[StrWithPrefix16]
	BoxPicPathInfo          AddrTo[StrWithPrefix16]
	BoxAudioPattern         AddrTo[StrWithPrefix16]
	BoxEffectPath           AddrTo[StrWithPrefix16]
}

type GachaEntranceManageMetaData struct {
	// Properties
	Gacha1Title                AddrTo[Hash]
	Gacha1Desc                 AddrTo[Hash]
	Gacha1ButtonText           AddrTo[Hash]
	Gacha10Title               AddrTo[Hash]
	Gacha10Desc                AddrTo[Hash]
	Gacha10ButtonText          AddrTo[Hash]
	GachaNameText              AddrTo[StrWithPrefix16]
	GachaTenDrawTips           AddrTo[StrWithPrefix16]
	Gacha10SpecificPic         AddrTo[StrWithPrefix16]
	UnlockLevel                int32
	TitleGachaContent_Avatar   AddrTo[Hash]
	TitleGachaContent_Weapon   AddrTo[Hash]
	TitleGachaContent_Stigmata AddrTo[Hash]
	TitleGachaContent_Material AddrTo[Hash]
	TitleGachaContent_Elf      AddrTo[Hash]
	GachaMainPageIsShow        bool
	PriorityValue              int32
	SortID                     int32
	GachaDefaultBoxType        int32
}

type GachaFirstDiscountData struct {
	// Properties
	IsEnableRedHint bool
	IsHideTimer     bool
}

type GachaGroupManageMetaData struct {
	// Properties
	GachaTypeList      AddrTo[[]int32]
	GachaGroupName     AddrTo[StrWithPrefix16]
	GachaGroupIconPath AddrTo[StrWithPrefix16]
}

type GachaPanelMetaData struct {
	// Properties
	GachaType int32
	BeginTime AddrTo[StrWithPrefix16]
	EndTime   AddrTo[StrWithPrefix16]
	Postion   int32
	ImagePath AddrTo[StrWithPrefix16]
	Text      AddrTo[Hash]
}

type GalAvatarStandByMetaData struct {
	// Properties
	StandByMotion AddrTo[StrWithPrefix16]
}

type GalEventEffectMetaData struct {
	// Properties
	TriggerPath   AddrTo[StrWithPrefix16]
	TriggerEffect AddrTo[StrWithPrefix16]
	Duration      float32
}

type GalEventMetaData struct {
	// Properties
	GalEventType      uint8
	Dialogue          AddrTo[Hash]
	RoleID            uint16
	AvatarID          AddrTo[[]int32]
	DressID           AddrTo[[]int32]
	StartMotion       AddrTo[StrWithPrefix16]
	StartMotionSpeed  float32
	LoopMotion        AddrTo[StrWithPrefix16]
	LoopMotionSpeed   float32
	LoopTimes         uint8
	EndMotion         AddrTo[StrWithPrefix16]
	EndMotionSpeed    float32
	AudioType         uint8
	Audio             AddrTo[StrWithPrefix16]
	AudioDelayTime    float32
	Effect            AddrTo[StrWithPrefix16]
	EffectDelayTime   float32
	Face              AddrTo[StrWithPrefix16]
	FaceDelayTime     float32
	Priority          uint16
	BubbleSwitch      uint8
	Bubble            AddrTo[StrWithPrefix16]
	BubbleCoordinate  AddrTo[StrWithPrefix16]
	TouchExp          uint8
	Condition1        uint8
	Parameter1        AddrTo[StrWithPrefix16]
	Condition2        uint8
	Parameter2        AddrTo[StrWithPrefix16]
	Condition3        uint8
	Parameter3        AddrTo[StrWithPrefix16]
	SubsidiaryEventID int32
	PortraitActive    bool
}

type GalEventMetaData_RawData struct {
	// Fields
	GalEventID        int32
	GalEventType      uint8
	RoleID            uint16
	StartMotionSpeed  float32
	LoopMotionSpeed   float32
	LoopTimes         uint8
	EndMotionSpeed    float32
	AudioType         uint8
	AudioDelayTime    float32
	EffectDelayTime   float32
	FaceDelayTime     float32
	Priority          uint16
	BubbleSwitch      uint8
	TouchExp          uint8
	Condition1        uint8
	Condition2        uint8
	Condition3        uint8
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
	// Properties
	GalEventID1 int32
	GalEventID2 int32
}

type GardenCropDataMetaData struct {
	// Properties
	WeatherList       AddrTo[StrWithPrefix16]
	UnlockCropNumList AddrTo[StrWithPrefix16]
	Delay             int32
	DelayIcon         AddrTo[StrWithPrefix16]
	Icon              AddrTo[StrWithPrefix16]
	Name              AddrTo[Hash]
	Desc              AddrTo[Hash]
	PlayPlot          int32
}

type GardenScheduleMetaData struct {
	// Properties
	BeginTime        AddrTo[StrWithPrefix16]
	EndTime          AddrTo[StrWithPrefix16]
	MinLevel         int32
	MaxLevel         int32
	DefaultWeather   int32
	MissionList      AddrTo[StrWithPrefix16]
	ProductCD        int32
	ProductMax       int32
	DailyMax         int32
	GrowUpMaterialID AddrTo[StrWithPrefix16]
}

type GardenWeatherDataMetaData struct {
	// Properties
	IconPath      AddrTo[StrWithPrefix16]
	SmallIconPath AddrTo[StrWithPrefix16]
	WeatherName   AddrTo[StrWithPrefix16]
	ShadowColor   AddrTo[StrWithPrefix16]
	NeedCropNum   int32
	PlayPlot      int32
}

type GenActivityRewardScheduleMetaData struct {
	// Properties
	ShowItemList AddrTo[[]ElfMetaData_MatCost]
}

type GenActivityRewardShowItemsMetaData struct {
	// Properties
	ItemID       int32
	ShowRarity   int32
	ComeFromList AddrTo[[]StrWithPrefix16]
}

type GeneralActivityActUIMetaData struct {
	// Properties
	ConfigPath   AddrTo[StrWithPrefix16]
	LetterRankID int32
}

type GeneralActivityDisplayDataMetaData struct {
	// Properties
	ActivityName        AddrTo[Hash]
	AcitivityDes        AddrTo[Hash]
	AcitivityDetailDes  AddrTo[Hash]
	AcitivityBG         AddrTo[StrWithPrefix16]
	MainReward          int32
	RankTitle           AddrTo[Hash]
	RankGroup           AddrTo[Hash]
	IsRankGroupHidden   bool
	RankIcon            AddrTo[StrWithPrefix16]
	RankDes             AddrTo[Hash]
	RankDetailDes       AddrTo[Hash]
	EnterImg            AddrTo[StrWithPrefix16]
	IsMP                bool
	DisplayLeftTime     bool
	DisplayScoreOnEntry bool
	SpecialImgBtnPath   AddrTo[StrWithPrefix16]
	SpecialImgPath      AddrTo[StrWithPrefix16]
	IsInfoDialogHidden  bool
}

type GeneralActivityLinkDataMetaData struct {
	// Properties
	MissionList     AddrTo[[]int32]
	GachaList       AddrTo[[]int32]
	ShopList        AddrTo[[]int32]
	ActivityPanelID int32
	PlotID          int32
	WebLink         AddrTo[StrWithPrefix16]
	BulletId        int32
	CgExtraKey      AddrTo[StrWithPrefix16]
}

type GeneralActivityMetaData struct {
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
	PreCondParaStr AddrTo[StrWithPrefix16]
	PreUnlockHint  AddrTo[Hash]
	ActivityBuffID int32
	TicketIDList   AddrTo[[]uint32]
}

type GeneralActivityOptionalBuffMetaData struct {
	// Properties
	StageID      int32
	BuffPool     AddrTo[[]int32]
	OriginScore  int32
	WarningScore int32
	BGPath       AddrTo[StrWithPrefix16]
}

type GeneralActivityRankDataMetaData struct {
	// Properties
	Reward int32
}

type GeneralActivityScoreDataMetaData struct {
	// Properties
	RewardType      int32
	ScoreData       int32
	Score           int32
	Reward          int32
	IsSpecialReward int32
	RewardShow      int32
}

type GeneralActivityStageGroupMetaData struct {
	// Properties
	AcitivityID          int32
	StageIDList          AddrTo[[]int32]
	StepStageIDList      AddrTo[[]int32]
	ChallengeNum         int32
	ChallengeReward      int32
	LevelPanelPrefabPath AddrTo[StrWithPrefix16]
	GroupImagePath       AddrTo[StrWithPrefix16]
	GroupTitle           AddrTo[Hash]
	GroupDescTitle       AddrTo[Hash]
	GroupDesc            AddrTo[Hash]
	GroupBackImagePath   AddrTo[StrWithPrefix16]
}

type GeneralActivityStageHomuMetaData struct {
	// Properties
	HomuPanel    AddrTo[StrWithPrefix16]
	BriefPicPath AddrTo[StrWithPrefix16]
}

type GeneralActivityStageMetaData struct {
	// Properties
	ActivityPanelType       int32
	PanelSizeRatio          float32
	ActivityDesBlockPic     AddrTo[StrWithPrefix16]
	ActivityTitlePic        AddrTo[StrWithPrefix16]
	ActivityLevelBlockPanel AddrTo[StrWithPrefix16]
	ActivityLevelBlockPic   AddrTo[StrWithPrefix16]
	ShowAllGroup            bool
	LevelPanelDragDir       int32
	DailyCount              int32
	Token                   int32
	TokenType               int32
	TokenCount              int32
	ShowReward              int32
	ShowRewardTips          AddrTo[Hash]
	LinkIcon                AddrTo[StrWithPrefix16]
	LinkName                AddrTo[StrWithPrefix16]
	LinkType                int32
	LinkParams              AddrTo[[]int32]
	LinkParamStr            AddrTo[StrWithPrefix16]
}

type GeneralActivityStageRewardMetaData struct {
}

type GeneralActivityTicketMetaData struct {
	// Properties
	MaterialID int32
	MaxNum     int32
}

type GeneralRarityDisplayConfigMetaData struct {
	// Properties
	GachaBoxPlateEffectPath AddrTo[StrWithPrefix16]
	GachaUIColorType        AddrTo[StrWithPrefix16]
	GachaDropNameTextColor  AddrTo[StrWithPrefix16]
	GachaDropIconBgPath     AddrTo[StrWithPrefix16]
}

type GeneralScheduleMetaData struct {
	// Properties
	VersionTagTitle AddrTo[Hash]
}

type GeniusCommonTowerMetaData struct {
	// Properties
	TowerID           uint8
	Type              uint8
	SiteContent       int32
	PreSite           int16
	SiteTitle         AddrTo[StrWithPrefix16]
	MonsterIDList     AddrTo[[]int32]
	RecommendCost     int32
	RecommendBuffList AddrTo[[]int32]
	HardTag           bool
	IsBossLevel       bool
}

type GeniusEndlessTowerMetaData struct {
	// Properties
	TowerID               uint8
	StageID               int32
	RecommendAvatarIDList AddrTo[[]int32]
	HardTag               bool
}

type GeniusEndlessTowerRewardMetaData struct {
	// Properties
	SiteID     int16
	TotalScore int32
	RewardID   int32
}

type GeniusTowerMetaData struct {
	// Properties
	MapID                uint16
	Type                 uint8
	TowerTitle           AddrTo[StrWithPrefix16]
	TowerSubTitle        AddrTo[StrWithPrefix16]
	TowerDifficultyTitle AddrTo[StrWithPrefix16]
	ImagePath            AddrTo[StrWithPrefix16]
	BGImagePath          AddrTo[StrWithPrefix16]
	UnlockConditionList  AddrTo[StrWithPrefix16]
	UnlockConditionTips  AddrTo[[]StrWithPrefix16]
}

type GerenalRulesEffectMetaData struct {
}

type GiftPackAdditionalDisplayMetaData struct {
	// Properties
	HorizontalBg AddrTo[StrWithPrefix16]
}

type GiftPackMetaData struct {
	// Properties
	GiftPackName          AddrTo[Hash]
	GiftPackTips          AddrTo[Hash]
	GiftPackPicPath       AddrTo[StrWithPrefix16]
	Rarity                int32
	DisplaySpecialEffect  int32
	ImmediateReward       int32
	ExtraRewardDays       int32
	ExtraRewardID         int32
	MainRewarMatrixList   AddrTo[[]int32]
	SelectableRewardList  AddrTo[[]int32]
	SelectableDisplayType uint8
}

type GlobalArmadaReunionLevelMetaData struct {
	// Properties
	TextMap_Level        AddrTo[StrWithPrefix16]
	Percent              float32
	Displayreward1       int32
	Displayreward2       int32
	Displayreward3       int32
	OriginalPaintingPath AddrTo[StrWithPrefix16]
	ActivePaintingPath   AddrTo[StrWithPrefix16]
}

type GlobalArmadaReunionScheduleMetaData struct {
	// Properties
	OpenTime        AddrTo[StrWithPrefix16]
	WaitingTime     AddrTo[StrWithPrefix16]
	RewardTime      AddrTo[StrWithPrefix16]
	RewardEndTime   AddrTo[StrWithPrefix16]
	DailyOpenTime   AddrTo[RemoteTimeSpan]
	CloseTime       AddrTo[StrWithPrefix16]
	PlayInstruction AddrTo[StrWithPrefix16]
}

type GlobalExploreActionDisplayMetaData struct {
	// Properties
	Name          AddrTo[Hash]
	PicPath       AddrTo[StrWithPrefix16]
	DicePointList AddrTo[[]int32]
}

type GlobalExploreActionMetaData struct {
	// Properties
	ActionType        uint8
	ActionName        AddrTo[Hash]
	FlagIDList        AddrTo[[]int32]
	EventIDList       AddrTo[[]int32]
	ActionDisplayList AddrTo[[]int32]
}

type GlobalExploreAreaMetaData struct {
	// Properties
	AreaBuffID         int32
	AreaName           AddrTo[Hash]
	ShowQuestList      AddrTo[[]int32]
	PreAreaList        AddrTo[[]int32]
	ThumbnailDeviation AddrTo[[]float32]
	ThumbnailScale     float32
	ThumbnailPath      AddrTo[StrWithPrefix16]
	AreaTpFrontMission int32
	DropLimitID        int32
	AreaIconPath       AddrTo[StrWithPrefix16]
	AudioSwitchEvent   AddrTo[StrWithPrefix16]
}

type GlobalExploreBuffDataMetaData struct {
	// Properties
	AbilityNameID AddrTo[StrWithPrefix16]
	AreaBuffName  AddrTo[Hash]
	AreaBuffDesc  AddrTo[Hash]
	FlagID        int32
	Params        AddrTo[[]float32]
}

type GlobalExploreDataMetaData struct {
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
	SettleTime             AddrTo[RemoteTime]
}

type GlobalExploreEntityMetaData struct {
	// Properties
	EntityName                   AddrTo[Hash]
	EntityDesc                   AddrTo[Hash]
	EntityType                   int32
	ActionIDList                 AddrTo[[]uint32]
	InteractActionIDList         AddrTo[[]uint32]
	ActiveProgress               int32
	IsGlobalProgress             bool
	IsBuffGrow                   bool
	ActiveRewardTypeList         AddrTo[[]uint8]
	ActiveEventIDList            AddrTo[[]uint32]
	GlobalExploreBloodID         uint8
	QuestID                      int32
	EntityTriggerType            uint8
	EntityFigure                 AddrTo[StrWithPrefix16]
	EntityStrokeFigure           AddrTo[StrWithPrefix16]
	EntityPrefab                 AddrTo[StrWithPrefix16]
	IsCanOverlap                 bool
	ActiveExtraStaminaLimit      int32
	ActiveExtraStaminaRecoverNum int32
	ActiveStaminaItemNum         int32
	RecommendedLevel             int32
	ViewRadius                   uint8
	TipsIcon                     AddrTo[StrWithPrefix16]
	DanmakuID                    int32
}

type GlobalExploreEventMetaData struct {
	// Properties
	HintTitle AddrTo[Hash]
}

type GlobalExploreFlagMetaData struct {
	// Properties
	FlagType  uint8
	FlagDesc  AddrTo[Hash]
	ParamList AddrTo[[]int32]
	ValidNum  int32
}

type GlobalExploreGridMapMetaData struct {
	// Properties
	PosX     int32
	PosY     int32
	EntityID uint32
	AreaID   uint8
	TilePic  AddrTo[StrWithPrefix16]
}

type GlobalExploreKingdomMetaData struct {
	// Properties
	StartPosX          int32
	StartPosY          int32
	AreaIDList         AddrTo[[]int32]
	KingdomName        AddrTo[Hash]
	KingdomIcon        AddrTo[StrWithPrefix16]
	KingdomFigure      AddrTo[StrWithPrefix16]
	MissionDisplayList AddrTo[[]int32]
	DanmakuID          int32
	EntranceEventID    int32
}

type GlobalExploreLevelMetaData struct {
	// Properties
	Exp        int32
	Attributes AddrTo[[]float32]
}

type GlobalExploreMapPathMetaData struct {
	// Properties
	MapPath AddrTo[StrWithPrefix16]
}

type GlobalExploreMessageBoardMetaData struct {
	// Properties
	UnlockArea          int32
	AreaSpotActiveNum   int32
	HeadPath            AddrTo[StrWithPrefix16]
	LayerMaster         AddrTo[StrWithPrefix16]
	MessageBoardTitle   AddrTo[Hash]
	MessageBoardContent AddrTo[Hash]
	Conciliations       AddrTo[StrWithPrefix16]
	CommentList         AddrTo[StrWithPrefix16]
	BGPath              AddrTo[StrWithPrefix16]
	PublishTime         AddrTo[Hash]
	MissionCondition    int32
}

type GlobalExplorePlotMetaData struct {
	// Properties
	CorrectAnswer AddrTo[[]int32]
	CorrectAction int32
	ErrorAction   int32
}

type GlobalExploreQuestMetaData struct {
	// Properties
	QuestType           uint8
	QuestName           AddrTo[Hash]
	QuestDesc           AddrTo[Hash]
	QuestRecommendLevel int32
	Progress            int32
	RewardDiasplay      int32
	FinishTimesLimit    int32
	AssociatedEntity    AddrTo[[]int32]
	Order               int32
}

type GlobalExploreStageMetaData struct {
	// Properties
	StageType          uint8
	EntityID           int32
	AreaBuffID         int32
	StageGroupID       int32
	MonstersUIDList    AddrTo[[]int32]
	MonsterSpawnList   AddrTo[[]StrWithPrefix16]
	StageName          AddrTo[StrWithPrefix16]
	BgmName            AddrTo[StrWithPrefix16]
	TimeLimit          int32
	TimeLimitIncrement AddrTo[[]int32]
}

type GlobalExploreStageScoreMetaData struct {
	// Properties
	MaxScore              int32
	ConvertContributeRate float32
}

type GlobalSupportRewardConfigMetaData struct {
	// Properties
	BeginTime              AddrTo[RemoteTime]
	EndTime                AddrTo[RemoteTime]
	GiftImgPath            AddrTo[StrWithPrefix16]
	ShareImgPath           AddrTo[StrWithPrefix16]
	ShareBackGroundImgPath AddrTo[StrWithPrefix16]
	ShareButtonImgPath     AddrTo[StrWithPrefix16]
}

type GlobalWarAreaMetaData struct {
	// Properties
	ClearPointNum uint8
	AreaName      AddrTo[Hash]
	AreaUnlockTip AddrTo[Hash]
}

type GlobalWarAvatarBuffMetaData struct {
	// Properties
	Hp              int32
	Atk             int32
	Def             int32
	AllDamageRatio  int32
	BuffIcon        AddrTo[StrWithPrefix16]
	BuffName        AddrTo[Hash]
	BuffDesc        AddrTo[Hash]
	BuffEffectDesc  AddrTo[StrWithPrefix16]
	SpecialBuffList AddrTo[[]GlobalWarAvatarBuffMetaData_BuffTypeLevel]
}

type GlobalWarAvatarBuffMetaData_BuffTypeLevel struct {
	// Fields
	BuffLevel uint16
	BuffType  uint16
}

type GlobalWarAvatarCollectionMetaData struct {
	// Properties
	ScheduleId int32
	MissionId  int32
	PointList  AddrTo[[]int32]
}

type GlobalWarBuffMetaData struct {
	// Properties
	BuffDesc AddrTo[Hash]
	BuffIcon AddrTo[StrWithPrefix16]
}

type GlobalWarDamageTextMetaData struct {
	// Properties
	DamageMax int32
	TextID    AddrTo[Hash]
}

type GlobalWarExchangeMetaData struct {
	// Properties
	CostMaterialList AddrTo[[]MechMetaData_DisjoinOutputItem]
	GetCurrencyNum   int32
	CostStamina      int32
}

type GlobalWarIsolatePointMetaData struct {
	// Properties
	UIIndex         int32
	StageList       AddrTo[[]int32]
	ClientPrePlotID int32
	Lockedtext      AddrTo[StrWithPrefix16]
	Finishedtext    AddrTo[StrWithPrefix16]
}

type GlobalWarPhotoMetaData struct {
	// Properties
	PhotoID             int32
	PhotoTag            uint8
	StartPointID        int32
	EndPointID          int32
	StartIsolatePointID int32
	EndIsolatePointID   int32
	PointTypelist       AddrTo[StrWithPrefix16]
	BeginTime           AddrTo[StrWithPrefix16]
	EndTime             AddrTo[StrWithPrefix16]
	IsShow              bool
}

type GlobalWarPointMetaData struct {
	// Properties
	UnlockTime            AddrTo[StrWithPrefix16]
	UIIndex               uint8
	PointType             uint8
	StageList             AddrTo[[]int32]
	StageBuffIcon         AddrTo[StrWithPrefix16]
	ExchangeID            uint16
	PrePointList          AddrTo[[]int32]
	PointTitle            AddrTo[Hash]
	PointDesc             AddrTo[Hash]
	RewardID              int32
	BuffList              AddrTo[[]int32]
	AvatarList            AddrTo[[]int32]
	LockAvatarId          int32
	PointHP               int32
	AreaID                int32
	PointImagePath        AddrTo[StrWithPrefix16]
	PreEventID            int32
	FinishEventID         int32
	DanmakuID             int32
	PointClientType       uint8
	SweepCurrencyCD       int32
	SweepCurrencyNumPerCD int32
	SweepCurrencyMaxNum   int32
}

type GlobalWarScheduleMetaData struct {
	// Properties
	GlobalRewardMinCurrency int32
	TicketID                int32
	TicketDailyMax          int32
	CurrencyID              int32
	CurrencyMax             int32
	ClientExchangeCoinID    int32
	AreaList                AddrTo[[]uint8]
	IsolatePointList        AddrTo[[]uint16]
	PhotoOpenTime           AddrTo[StrWithPrefix16]
	ShopType                int32
	RewardShowID            int32
	MissionList             AddrTo[[]int32]
	DanmakuID               int32
	PerformID               int32
	BackgroundCGID          int32
	CGID                    int32
	PlotID                  int32
	WebLink                 AddrTo[StrWithPrefix16]
	GachaLink               AddrTo[StrWithPrefix16]
	Tutorial                AddrTo[[]GlobalWarScheduleMetaData_ConfigTutorial]
}

type GlobalWarScheduleMetaData_ConfigTutorial struct {
	// Objects
	Content StrWithPrefix16
	Sprite  StrWithPrefix16
	Title   StrWithPrefix16
}

type GlobalWarScoreMetaData struct {
	// Properties
	Currency int32
}

type GlobalWarSpecialBuffMetaData struct {
	// Properties
	StageDetailEffectList AddrTo[[]int32]
}

type GlobalWarStageMetaData struct {
	// Properties
	MinLevel          uint8
	MaxLevel          uint8
	MinCurrency       int32
	MaxCurrency       int32
	StageBuffTypeList AddrTo[[]uint16]
}

type GoBackFreeSelectMetaData struct {
	// Properties
	PreSelectedIcon   AddrTo[StrWithPrefix16]
	PreSelectedRarity int32
	TitleText         AddrTo[Hash]
	DescText          AddrTo[Hash]
	RewardIDList      AddrTo[[]int32]
}

type GoBackFundMetaData struct {
	// Properties
	FundPrice         int32
	FundShowMission   int32
	DiscountText      AddrTo[Hash]
	ShowRewardID      int32
	FundBannerLeftPic AddrTo[StrWithPrefix16]
}

type GoBackFundRewardMetaData struct {
	// Properties
	MissionID int32
	RewardID  int32
}

type GoBackGachaMetaData struct {
	// Properties
	GachaTypeID      uint32
	ShowMaterial     uint32
	DescriptionText  AddrTo[Hash]
	CurrentGachaLink AddrTo[StrWithPrefix16]
}

type GoBackGamePlayConfigMetaData struct {
	// Properties
	PicPath           AddrTo[StrWithPrefix16]
	Title             AddrTo[Hash]
	Name              AddrTo[Hash]
	TagDesc           AddrTo[Hash]
	Goto              AddrTo[StrWithPrefix16]
	ProgressType      uint16
	ProgressParam     AddrTo[[]int32]
	RewardMissionList AddrTo[[]int32]
	IsEveryDayNew     bool
}

type GoBackGrowUpActivityMetaData struct {
	// Properties
	MissionChainDisplay AddrTo[StrWithPrefix16]
	RankMissionGroupID  int32
	EndLevel            uint8
}

type GoBackLoginImgMetaData struct {
	// Properties
	ImgAvatar     AddrTo[StrWithPrefix16]
	ImgTheme      AddrTo[StrWithPrefix16]
	ImgTitle      AddrTo[StrWithPrefix16]
	ImgBG         AddrTo[StrWithPrefix16]
	ImgBGLight    AddrTo[StrWithPrefix16]
	AvatarPosX    int32
	AvatarPosY    int32
	NewLoginTitle AddrTo[Hash]
	NewLoginInfo  AddrTo[Hash]
}

type GoBackMissionDayMetaData struct {
	// Properties
	MissionList AddrTo[[]int32]
	PageText    AddrTo[Hash]
}

type GoBackMissionScoreMetaData struct {
	// Properties
	DisplayPreMissionID int32
	Score               int32
}

type GoBackScheduleMetaData struct {
	// Properties
	MissionList             AddrTo[[]int32]
	MissionTitleName        AddrTo[Hash]
	MissionBanner           AddrTo[StrWithPrefix16]
	ScoreRewardList         AddrTo[[]int32]
	LoginRewardID           AddrTo[[]int32]
	LoginCGID               int32
	ShopType                int32
	MallType                int32
	LoginRewardUIType       int32
	GoBackMissionUIType     int32
	GachaConfigIDList       AddrTo[[]int32]
	MissionShopShowCurrency AddrTo[[]int32]
	FundMallShowCurrency    AddrTo[[]int32]
	GrowUpConfigID          uint8
	GoBackGamePlayConfigID  AddrTo[[]uint32]
}

type GoBackScoreRewardMetaData struct {
	// Properties
	TotalScore int32
	RewardID   int32
}

type GoBackTabMetaData struct {
	// Properties
	SortId           int32
	TabName          AddrTo[StrWithPrefix16]
	TabIcon          AddrTo[StrWithPrefix16]
	Currency         AddrTo[[]int32]
	Params           AddrTo[[]int32]
	OpenScheduleList AddrTo[[]int32]
}

type GoBackWebMetaData struct {
	// Properties
	WebLink      AddrTo[StrWithPrefix16]
	MinGobackDay int32
	MaxGobackDay int32
}

type GodWarAreaMetaData struct {
	// Properties
	SiteIDList        AddrTo[[]int32]
	PreAreaID         int32
	AreaName          AddrTo[Hash]
	AreaImage         AddrTo[StrWithPrefix16]
	AreaLevelDesc     AddrTo[StrWithPrefix16]
	AreaBossList      AddrTo[[]GodWarAreaMetaData_BossDataItem]
	AreaUnlockMission int32
	LockHintText      AddrTo[Hash]
	AreaSelectEvent   int32
	TaleID            int32
	Index             int32
}

type GodWarAreaMetaData_BossDataItem struct {
	// Fields
	BossID      int32
	ShowMission int32
}

type GodWarAvatarAbilityMetaData struct {
	// Properties
	AbilityType             uint8
	AbilityTitle            AddrTo[Hash]
	AbilityDes              AddrTo[Hash]
	AbilityIcon             AddrTo[StrWithPrefix16]
	IsShowInAvatarChallenge bool
}

type GodWarAvatarChallengeMetaData struct {
	// Properties
	StepLevel                 uint8
	AvatarSpecialBuffList     AddrTo[[]uint32]
	CoreBuffList              AddrTo[[]uint32]
	SupportAvatarList         AddrTo[[]int32]
	RewardMaterialMap         AddrTo[[]GodWarAvatarChallengeMetaData_MaterialNumPair]
	AvatarExpMaterialAddNum   uint32
	StageName                 AddrTo[StrWithPrefix16]
	HardLevel                 uint32
	BossName                  AddrTo[StrWithPrefix16]
	Wave1                     AddrTo[[]uint32]
	Wave2                     AddrTo[[]uint32]
	Wave3                     AddrTo[[]uint32]
	Wave4                     AddrTo[[]uint32]
	Wave5                     AddrTo[[]uint32]
	FirstWinExpMaterialAddNum uint32
}

type GodWarAvatarChallengeMetaData_MaterialNumPair struct {
	// Fields
	Material uint32
	Num      uint32
}

type GodWarAvatarLevelDataMetaData struct {
	// Properties
	ATK uint16
	HP  uint16
}

type GodWarAvatarTaleScheduleMetaData struct {
	// Properties
	BeginTime         AddrTo[RemoteTime]
	EndTime           AddrTo[RemoteTime]
	MainAvatarList    AddrTo[[]uint32]
	VirtualGroup      uint32
	SupportAvatarList AddrTo[[]uint32]
	PunishStepID      int32
}

type GodWarAvatarUpScheduleMetaData struct {
	// Properties
	DisplayUpMaterialList AddrTo[[]int32]
}

type GodWarBuffMetaData struct {
	// Properties
	BuffSuit           uint32
	BuffQuality        uint32
	AbilityName        AddrTo[StrWithPrefix16]
	ParamList          AddrTo[[]StrWithPrefix16]
	BuffIcon           AddrTo[StrWithPrefix16]
	BuffName           AddrTo[Hash]
	BuffDesc           AddrTo[Hash]
	BuffUpDesc         AddrTo[Hash]
	SimpleBuffDesc     AddrTo[Hash]
	BuffTagList        AddrTo[[]uint16]
	ShowBuffIconEffect bool
	BuffGroup          uint32
}

type GodWarBuffPoolMetaData struct {
	// Properties
	BuffSuit uint32
}

type GodWarBuffSuitMetaData struct {
	// Properties
	SuitIcon           AddrTo[StrWithPrefix16]
	SuitName           AddrTo[Hash]
	SuitDesc           AddrTo[Hash]
	SuitImage          AddrTo[StrWithPrefix16]
	ShowSuitIconEffect bool
	ShowFilterOption   bool
}

type GodWarChallengeBuffMetaData struct {
	// Properties
	BuffCost      int32
	BuffDesc      AddrTo[Hash]
	BuffAbility   AddrTo[StrWithPrefix16]
	BuffParamList AddrTo[[]StrWithPrefix16]
}

type GodWarChallengeLevelMetaData struct {
	// Properties
	Floor       int32
	Stage       AddrTo[StrWithPrefix16]
	ChallengeID int32
	MissionFlag int32
	ReplaceID   int32
	Wave1       AddrTo[[]int32]
	Wave2       AddrTo[[]int32]
	Wave3       AddrTo[[]int32]
	Wave4       AddrTo[[]int32]
}

type GodWarChallengeMetaData struct {
	// Properties
	BossPreviewMap AddrTo[[]GodWarChallengeMetaData_FloorBossDataItem]
	EffectIDList   AddrTo[[]int32]
	BuffList       AddrTo[[]int32]
	BuffCostLimit  int32
}

type GodWarChallengeMetaData_FloorBossDataItem struct {
	// Fields
	BossID         int32
	DifficultyType int32
	FloorNum       int32
}

type GodWarChallengeRewardMetaData struct {
	// Properties
	Mission int32
	Score   int32
	Icon    AddrTo[StrWithPrefix16]
}

type GodWarChapterMetaData struct {
	// Properties
	ChapterName        AddrTo[Hash]
	ChapterSubName     AddrTo[Hash]
	ChapterEnterImage  AddrTo[StrWithPrefix16]
	LockHintText       AddrTo[Hash]
	BeginTime          AddrTo[RemoteTime]
	EndTime            AddrTo[RemoteTime]
	LobbyID            int32
	GodWarID           int32
	TaleIDList         AddrTo[[]int32]
	KeyNodeMissionList AddrTo[[]int32]
}

type GodWarChar2DMetaData struct {
	// Properties
	CharID            uint32
	FaceType          uint8
	ImageChar2DPath   AddrTo[StrWithPrefix16]
	SpImageChar2DPath AddrTo[StrWithPrefix16]
}

type GodWarClientDataMetaData struct {
	// Properties
	PhotoIDMap              AddrTo[GodWarClientDataMetaData_CollectionData]
	ShopLink                uint32
	EntryIconPath           AddrTo[StrWithPrefix16]
	WebLink                 AddrTo[StrWithPrefix16]
	EntryPerformID          int32
	GodWarShowPopDropIDList AddrTo[[]int32]
	BlockRoleIDList         AddrTo[[]uint32]
	TopRoleIDList           AddrTo[[]uint32]
	ProgressMissionList     AddrTo[[]int32]
	MaterialList            AddrTo[[]int32]
	SkipMissionMaterialID   uint32
	WeeklyRewardPreview     int32
	SpecialStageOverallKey  int32
}

type GodWarClientDataMetaData_CollectionData struct {
	// Fields
	Num          int32
	BeginPhotoID uint32
}

type GodWarCollectionDataMetaData struct {
	// Properties
	GodWarCollectionType  uint8
	GodWarSubType         uint8
	LockName              AddrTo[Hash]
	CollectionSubTitle    AddrTo[Hash]
	LockSubTitle          AddrTo[Hash]
	DetailDesc            AddrTo[Hash]
	BriefDesc             AddrTo[Hash]
	LockDesc              AddrTo[Hash]
	MissionID             int32
	GodWarID              uint32
	CollectionStarNum     uint32
	ShowImgOffset         AddrTo[[]float32]
	GodWarChapterID       int32
	CollectionMissionType int32
	CollectionOrder       int32
	BuffID                uint32
}

type GodWarCollectionSuitDataMetaData struct {
	// Properties
	GodWarID         uint32
	CollectionIDList AddrTo[[]int32]
	SuitType         int32
	SuitTextID       AddrTo[Hash]
	MissionID        int32
}

type GodWarEventFlagMetaData struct {
	// Properties
	FlagType        uint16
	Operator        uint16
	ParamsVar       AddrTo[[]uint32]
	ConditionTextID AddrTo[Hash]
}

type GodWarEventMetaData struct {
	// Properties
	EventType uint8
	ParamsVar AddrTo[[]uint32]
}

type GodWarEventPanelMetaData struct {
	// Properties
	Icon     AddrTo[StrWithPrefix16]
	IconType uint32
	Name     AddrTo[Hash]
	Desc     AddrTo[Hash]
	Star     uint8
}

type GodWarEventPlotMetaData struct {
	// Properties
	EventParams          AddrTo[[]GodWarEventPlotMetaData_GodWarDialogOverData]
	DefaultFollowEventID uint32
}

type GodWarEventPlotMetaData_GodWarDialogOverData struct {
	// Fields
	EventID        int32
	FinishDialogID int32
}

type GodWarEventPropMetaData struct {
	// Properties
	PropName     AddrTo[StrWithPrefix16]
	PropEvent    uint32
	InteractIcon AddrTo[StrWithPrefix16]
	InteractDesc AddrTo[Hash]
}

type GodWarEventScoreMetaData struct {
	// Properties
	ScoreName AddrTo[Hash]
}

type GodWarEventStageMetaData struct {
	// Properties
	StartCoinEventID uint32
	FollowEventID    uint32
	FailEventID      uint32
	Score            int32
}

type GodWarExtraItemMetaData struct {
	// Properties
	GodWarID              uint16
	ExtraItemSkillType    uint16
	AbilityName           AddrTo[StrWithPrefix16]
	ParamList             AddrTo[[]StrWithPrefix16]
	ExtraItemStar         uint16
	ExtraItemName         AddrTo[Hash]
	ExtraItemSkillOverall uint32
	ExtraItemIcon         AddrTo[StrWithPrefix16]
	UnlockExtraItemHint   AddrTo[Hash]
	ExtraItemSkill        AddrTo[Hash]
	ExtraItemType         int32
	ExtraItemSuitID       int32
}

type GodWarExtraItemTextMetaData struct {
	// Properties
	TextMap        AddrTo[[]GodWarExtraItemTextMetaData_OverallTextMap]
	DefaultTextMap AddrTo[Hash]
}

type GodWarExtraItemTextMetaData_OverallTextMap struct {
	// Fields
	OverallValue uint32

	// Objects
	Desc Hash
}

type GodWarFlameChasersMetaData struct {
	// Properties
	RoleID       uint32
	IconPath     AddrTo[StrWithPrefix16]
	BuffSuitName AddrTo[Hash]
	BuffSuitDes  AddrTo[Hash]
	Order        int32
}

type GodWarHardLevelMetaData struct {
	// Properties
	DamageRaito float32
}

type GodWarHardLevelScheduleMetaData struct {
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
	// Properties
	LobbyID        int32
	PrefabPath     AddrTo[StrWithPrefix16]
	OnEnterEventID uint32
	OnExitEventID  uint32
}

type GodWarLobbyAvatarDataMetaData struct {
	// Properties
	AvatarPrefabPath AddrTo[StrWithPrefix16]
	PositionID       uint32
	SubPositionID    uint32
	SceneConfigID    uint32
	AvatarAnimation  int32
	AvatarEventGroup AddrTo[[]uint32]
	RoleID           int32
	NpcName          AddrTo[StrWithPrefix16]
	IconPath         AddrTo[StrWithPrefix16]
}

type GodWarLobbyButtonMetaData struct {
	// Properties
	ButtonUnlockMissionID int32
	TargetUIPath          AddrTo[StrWithPrefix16]
	IsButtonActive        bool
	SpeculativeEventID    uint32
	ButtonType            int32
}

type GodWarLobbyDataMetaData struct {
	// Properties
	SceneConfigIDList AddrTo[[]uint32]
	TaleID            int32
}

type GodWarLobbyInteractActionMetaData struct {
	// Properties
	ActionType       int32
	ActionParam      AddrTo[[]int32]
	ImgPath          AddrTo[StrWithPrefix16]
	InteractPropDesc AddrTo[Hash]
	DisplayMission   int32
	UnlockMission    AddrTo[[]int32]
	DeleteMission    int32
}

type GodWarLobbyInteractPropMetaData struct {
	// Properties
	LobbyID      int32
	Type         int32
	PrefabPath   AddrTo[StrWithPrefix16]
	Radius       float32
	Angle        float32
	ActionIDList AddrTo[[]int32]
	HintOffset   AddrTo[[]float32]
}

type GodWarMainAvatarMetaData struct {
	// Properties
	AbilityName                AddrTo[StrWithPrefix16]
	ParamList                  AddrTo[[]StrWithPrefix16]
	StepUnlockMissionMap       AddrTo[[]GodWarMainAvatarMetaData_StepUnLockMissionDataItem]
	AvatarAbilityID            AddrTo[[]uint32]
	AvatarMissionList          AddrTo[[]int32]
	RecommendLink              AddrTo[StrWithPrefix16]
	AvatarChallengeStageIDList AddrTo[[]int32]
}

type GodWarMainAvatarMetaData_StepUnLockMissionDataItem struct {
	// Fields
	MissionID int32
	StepID    int32
}

type GodWarMainMissionNodeMetaData struct {
	// Properties
	ChapterID             int32
	NodeType              int32
	PreNodeID             int32
	NodeName              AddrTo[Hash]
	AssociatedMissionList AddrTo[[]int32]
	SkipCostMaterialID    int32
	SkipCostMaterialNum   uint32
	IsShowTime            bool
	UnlockTime            AddrTo[RemoteTime]
	Icon                  AddrTo[StrWithPrefix16]
	RewardPreview         int32
}

type GodWarMapMetaData struct {
	// Properties
	AreaList     AddrTo[[]int32]
	PrefabPath   AddrTo[StrWithPrefix16]
	LastAreaList AddrTo[[]int32]
}

type GodWarMaterialMetaData struct {
	// Properties
	MaterialType  uint8
	QuantityLimit int32
}

type GodWarMissionDataMetaData struct {
	// Properties
	GodWarID           uint32
	GodWarChapterID    int32
	ParentTabID        uint32
	SubTabID           uint32
	ParentNode         uint32
	MissionIconPath    AddrTo[StrWithPrefix16]
	ShowTimeRemain     bool
	LobbyCharID        int32
	DisplayMissionType int32
}

type GodWarNodeMetaData struct {
	// Properties
	PreNodeID   uint32
	NodeIcon    AddrTo[StrWithPrefix16]
	ParentTabID uint32
	SubTabID    uint32
}

type GodWarParentTabDataMetaData struct {
	// Properties
	TabIconPath   AddrTo[StrWithPrefix16]
	TabNameTextID AddrTo[Hash]
}

type GodWarPlotStageMetaData struct {
	// Properties
	StageID   int32
	MissionID int32
}

type GodWarPunishBuffMetaData struct {
	// Properties
	PunishLevelCost     int32
	PunishBuffName      AddrTo[Hash]
	PunishBuffDesc      AddrTo[Hash]
	PunishBuffType      uint8
	PunishBuffAbility   AddrTo[StrWithPrefix16]
	PunishBuffParamList AddrTo[[]StrWithPrefix16]
	UnlockPunishLevel   uint8
	UnlockHint          AddrTo[Hash]
	UnlockMissionID     uint32
	PunishBuffIcon      AddrTo[StrWithPrefix16]
}

type GodWarPunishRewardMetaData struct {
	// Properties
	PunishGradeDesc         AddrTo[Hash]
	PunishMaterial          AddrTo[[]GodWarPunishRewardMetaData_MaterialNumPair]
	AvatarExpMaterialAddNum int32
	ShowRewardMission       int32
}

type GodWarPunishRewardMetaData_MaterialNumPair struct {
	// Fields
	Material int32
	Num      int32
}

type GodWarPunishScoreMetaData struct {
	// Properties
	PunishExtraScore     float32
	PunishRecommendLevel AddrTo[[]GodWarPunishScoreMetaData_AvatarLevelRatio]
	MaxPunishMaterial    uint32
}

type GodWarPunishScoreMetaData_AvatarLevelRatio struct {
	// Fields
	AvatarLevel uint32
	Ratio       uint32
}

type GodWarPunishStepMetaData struct {
	// Properties
	PunishStepTitle           AddrTo[Hash]
	PunishStepFirstReward     uint32
	PunishStepInitialBuffList AddrTo[[]uint32]
	PunishStepUnlockHint      AddrTo[Hash]
	PunishStepType            uint8
	PunishParamMissionLink    AddrTo[[]uint32]
	PunishStepMaxLevel        uint32
	PunishStepMaxHardLevel    uint32
	PunishStepPrefabPath      AddrTo[StrWithPrefix16]
	StoryHint                 AddrTo[Hash]
	ChallengeID               int32
	MainMissionNode           int32
	MapID                     int32
	SupportAvatarLimit        AddrTo[[]int32]
}

type GodWarRelationDataMetaData struct {
	// Properties
	AvatarID             uint32
	AvatarLevelID        uint32
	MaxLevel             uint16
	DisplayLevelCap      uint16
	Exp                  uint32
	ExpMaterialID        uint32
	LvUpMaterial         AddrTo[GodWarRelationDataMetaData_UpgradeMaterial]
	LvUpMissionLinkParam AddrTo[[]uint32]
	LockHint             AddrTo[Hash]
	RewardID             uint32
	IconPath             AddrTo[StrWithPrefix16]
	IconName             AddrTo[Hash]
	IconDesc             AddrTo[Hash]
	FigurePath           AddrTo[StrWithPrefix16]
	FigureName           AddrTo[Hash]
	FigureSubName        AddrTo[Hash]
	ArchiveLinkParam     AddrTo[[]uint32]
	IconOffset           AddrTo[[]float32]
	AvatarArtifactType   uint8
	RoleType             uint8
}

type GodWarRelationDataMetaData_UpgradeMaterial struct {
	// Fields
	MaterialID uint32
	Num        uint32
}

type GodWarRoleSkillDataMetaData struct {
	// Properties
	RoleID              uint16
	NeedRelationLevel   uint16
	UnlockAvatarStar    uint8
	UnlockAvatarSubStar uint8
	AbilityName         AddrTo[StrWithPrefix16]
	ParamList           AddrTo[[]StrWithPrefix16]
	SkillName           AddrTo[Hash]
	SkillDesc           AddrTo[Hash]
	ShowSkillDesc       bool
}

type GodWarRoleStoryDataMetaData struct {
	// Properties
	RoleID            uint32
	NeedRelationLevel uint32
	UnlockMissionList AddrTo[[]uint32]
	StoryType         uint8
	TypeParam         uint32
	AutoPlay          bool
	RewardID          uint32
	StoryTitle        AddrTo[Hash]
	StoryDetail       AddrTo[Hash]
	Picture           AddrTo[StrWithPrefix16]
}

type GodWarSceneConfigMetaData struct {
	// Properties
	OpenSound          AddrTo[StrWithPrefix16]
	CloseSound         AddrTo[StrWithPrefix16]
	StartEventGroup    uint32
	ButtonConfigIDList AddrTo[[]uint32]
}

type GodWarScheduleMetaData struct {
	// Properties
	SupportAvatarSlotMaterialList AddrTo[[]uint32]
	BeginTime                     AddrTo[RemoteTime]
	EndTime                       AddrTo[RemoteTime]
	TicketList                    AddrTo[[]uint32]
}

type GodWarSiteMetaData struct {
	// Properties
	SiteName     AddrTo[Hash]
	SiteIndex    uint32
	SiteMaxIndex uint32
	SiteEvent    uint32
	BuffEvent    uint32
	AreaID       int32
}

type GodWarStageHintMetaData struct {
	// Properties
	HintName           AddrTo[Hash]
	HintImage          AddrTo[StrWithPrefix16]
	LevelHintText      AddrTo[Hash]
	LevelChallengeText AddrTo[Hash]
	RewardShow         int32
}

type GodWarStageSkillMetaData struct {
	// Properties
	StageSkillType        uint8
	StageSkillRandomAudio AddrTo[[]StrWithPrefix16]
	StageSkillImage       AddrTo[StrWithPrefix16]
	StageSkillIcon        AddrTo[StrWithPrefix16]
	StageSkillName        AddrTo[Hash]
	StageSkillDesc        AddrTo[Hash]
	OffsetX               float32
	OffsetY               float32
	Scale                 float32
}

type GodWarSubTabDataMetaData struct {
	// Properties
	TabIconPath   AddrTo[StrWithPrefix16]
	TabNameTextID AddrTo[Hash]
	ImageType     uint32
}

type GodWarSupportAvatarMetaData struct {
	// Properties
	SkillFunctionName       AddrTo[StrWithPrefix16]
	AbilityName             AddrTo[StrWithPrefix16]
	ParamList               AddrTo[[]StrWithPrefix16]
	CD                      uint32
	SkillName               AddrTo[Hash]
	SkillDesc               AddrTo[Hash]
	SkillIcon               AddrTo[StrWithPrefix16]
	UnlockSupportAvatarHint AddrTo[Hash]
	IsSkillChange           bool
}

type GodWarSupportLevelMetaData struct {
	// Properties
	TalentPointExp uint32
	SupportType    uint16
	IconPath       AddrTo[StrWithPrefix16]
	Name           AddrTo[Hash]
	Desc           AddrTo[Hash]
}

type GodWarTaleClientDataMetaData struct {
	// Properties
	TaleName       AddrTo[Hash]
	TaleEnterImage AddrTo[StrWithPrefix16]
	RewardPreview  int32
	LockHintText   AddrTo[StrWithPrefix16]
}

type GodWarTaleMetaData struct {
	// Properties
	RogueCoin                 uint32
	StartEvent                uint32
	TeleportResetNum          uint32
	CloseVirtualAvatarMission uint32
	StageList                 AddrTo[[]int32]
	TaleType                  int32
}

type GodWarTalentClientDataMetaData struct {
	// Properties
	GodWarID        uint16
	TalentArea      uint8
	TalentCol       uint16
	TalentLineList  AddrTo[[]uint16]
	TalentIndexList AddrTo[[]GodWarTalentClientDataMetaData_TalentIndex]
}

type GodWarTalentClientDataMetaData_TalentIndex struct {
	// Fields
	Index    uint16
	TalentID uint32
}

type GodWarTalentDataMetaData struct {
	// Properties
	TaleIDList       AddrTo[[]uint16]
	TalentEffectType uint16
	MaxLevel         uint16
	TalentName       AddrTo[Hash]
	IconPath         AddrTo[StrWithPrefix16]
}

type GodWarTalentLevelDataMetaData struct {
	// Properties
	TalentDesc          AddrTo[Hash]
	LvUpFlagList        AddrTo[[]uint32]
	UpGradeMaterialList AddrTo[[]GodWarTalentLevelDataMetaData_UpGradeMaterial]
	TalentEffect        AddrTo[StrWithPrefix16]
	EffectParamList     AddrTo[[]StrWithPrefix16]
	TalentPoint         uint16
}

type GodWarTalentLevelDataMetaData_UpGradeMaterial struct {
	// Fields
	Cost       uint32
	MaterialID uint32
}

type GodWarTaleScheduleMetaData struct {
	// Properties
	TaleIDList AddrTo[[]uint16]
	BeginTime  AddrTo[RemoteTime]
	EndTime    AddrTo[RemoteTime]
}

type GodWarTeleportMetaData struct {
	// Properties
	TeleportType    uint8
	TeleportProp    AddrTo[StrWithPrefix16]
	TeleportIcon    AddrTo[StrWithPrefix16]
	TeleportEvent   uint32
	TeleportTitle   AddrTo[Hash]
	TeleportContent AddrTo[Hash]
	TeleportDesc    AddrTo[Hash]
	TeleportAbility AddrTo[StrWithPrefix16]
	InteractIcon    AddrTo[StrWithPrefix16]
	InteractDesc    AddrTo[Hash]
}

type GodWarTypewriterMetaData struct {
	// Properties
	TypewriterFileID uint32
	Order            uint16
	FileTitle        AddrTo[Hash]
	FileContent      AddrTo[Hash]
	UnlockMissionID  uint32
	UnlockFlagDes    AddrTo[Hash]
	StartEvent       uint32
	FinishEvent      uint32
	UIPath           AddrTo[StrWithPrefix16]
}

type GodWarUseAvatarMetaData struct {
	// Properties
	StepUnlockMission     int32
	AvatarID              int32
	LobbyAvatarPrefabPath AddrTo[StrWithPrefix16]
}

type GrandKeyBuffActiveInfoMetaData struct {
	// Properties
	Duration  int32
	Cost      AddrTo[[]StrWithPrefix16]
	CostLabel AddrTo[Hash]
}

type GrandKeyLevel struct {
	// Properties
	UnlockWeaponLevel int32
	PlayerLevel       int32
	UpgradeMaterial   AddrTo[[]MechMetaData_DisjoinOutputItem]
}

type GrandKeyMainWeapon struct {
	// Properties
	WeaponType       int32
	EffectPath       AddrTo[StrWithPrefix16]
	Open             int32
	PositionCorrect1 float32
	PositionCorrect2 float32
	PositionCorrect3 float32
	StrokeColor      AddrTo[StrWithPrefix16]
	StrokeThickness  float32
	Material         AddrTo[StrWithPrefix16]
}

type GrandKeyMetaData struct {
	// Properties
	MaxGrandKeyLv       int32
	PositionNum         int32
	GrandkeyName        AddrTo[Hash]
	MainWeaponIdList    AddrTo[[]int32]
	SubWeaponmainIDlist AddrTo[[]int32]
	BreachIconList      AddrTo[[]StrWithPrefix16]
	WebLinks            AddrTo[StrWithPrefix16]
}

type GrandKeyWeaponContrast struct {
	// Properties
	WeaponMainIDBefore int32
	WeaponIDBefore     int32
	WeaponMainIDAfter  int32
	WeaponIDAfter      int32
}

type GrandKeyWeaponRaidSkillMetaData struct {
	// Properties
	PropRaidID     int32
	PropRaidParam1 float32
	PropRaidParam2 float32
	PropRaidParam3 float32
}

type GratuityMessageMetaData struct {
	// Properties
	Initial uint32
}

type GratuityScheduleMetaData struct {
	// Properties
	CurrencyID       int32
	TeamListID       int32
	RewardLinkType   int32
	RewardLinkParams AddrTo[[]int32]
}

type GratuityStageClassMetaData struct {
	// Properties
	StageClassDec       AddrTo[Hash]
	ReciveStageMinLevel int32
	Colors              AddrTo[[]StrWithPrefix16]
	TeamSubList         int32
	TeamSubListTitle    AddrTo[Hash]
	DefeatRewardNum     int32
	StaminaCoef         int32
	DamageCoef          float32
	BonusStaminaRatio   float32
	BonusDamageRatio    float32
}

type GratuityStageMetaData struct {
	// Properties
	MonsterWaveList AddrTo[[]GratuityStageMetaData_MessageIDWaveData]
	SpecialBuffText AddrTo[Hash]
	EnemyNumType    uint8
}

type GratuityStageMetaData_MessageIDWaveData struct {
	// Fields
	Wave            uint16
	MessageID       int32
	UniqueMonsterID int32
}

type GreedyEndlessBattleConfigMetaData struct {
	// Properties
	StageID             int32
	WeatherID           int32
	MechanismID         int32
	MonsterTable        AddrTo[[]int32]
	StageMonsterDisplay AddrTo[[]int32]
	TimeLimit           int32
	ChallengeTime       int32
	FloorType           int32
}

type GreedyEndlessBuffMetaData struct {
	// Properties
	BuffName        AddrTo[StrWithPrefix16]
	BuffType        int32
	InLevelIconPath AddrTo[StrWithPrefix16]
	Effect          AddrTo[StrWithPrefix16]
}

type GreedyEndlessFloorConfigMetaData struct {
	// Properties
	FloorDisplayType       int16
	FloorReward            int32
	FloorName              AddrTo[Hash]
	RankIconPath           AddrTo[StrWithPrefix16]
	RankTitle              AddrTo[Hash]
	RankDesc               AddrTo[Hash]
	ChallengeRewardDisplay int32
	ChallengeFloorName     AddrTo[Hash]
	Challenge              int32
	BaseScore              int32
	ExtraScore             int32
	FloorEntrancePath      AddrTo[StrWithPrefix16]
	ChangeType             uint8
	IsAutoContinue         uint8
}

type GreedyEndlessGroupMetaData struct {
	// Properties
	PlayerGroup      int32
	IsInactiveGroup  bool
	GroupName        AddrTo[Hash]
	GroupDesc        AddrTo[Hash]
	GroupIcon        AddrTo[StrWithPrefix16]
	BgColor          AddrTo[StrWithPrefix16]
	SelectPrefabPath AddrTo[StrWithPrefix16]
}

type GreedyEndlessMechanismMetaData struct {
	// Properties
	BuffID            int32
	MechanismNameText AddrTo[Hash]
	MechanismDescText AddrTo[Hash]
	MechanismIconPath AddrTo[StrWithPrefix16]
}

type GreedyEndlessPlayerGroupMetaData struct {
	// Properties
	MinPlayerLevel int32
	MaxPlayerLevel int32
	Icon           AddrTo[StrWithPrefix16]
	Name           AddrTo[Hash]
}

type GreedyEndlessRankRewardMetaData struct {
	// Properties
	RankReward int32
}

type GreedyEndlessScheduleMetaData struct {
	// Properties
	SettleRewardConfig int32
}

type GreedyEndlessSettleConfigMetaData struct {
	// Properties
	PrototeRewardID    int32
	NormalRewardID     int32
	DemoteRewardID     int32
	GroupRewardDispaly int32
}

type GreedyEndlessWeatherMetaData struct {
	// Properties
	BuffID          int32
	WeatherNameText AddrTo[Hash]
	WeatherDescText AddrTo[Hash]
	WeatherIconPath AddrTo[StrWithPrefix16]
}

type HoukaiTownBuffMetaData struct {
	// Properties
	BuffType      uint8
	ParamList     AddrTo[[]int32]
	DurationType  uint8
	DurationNum   uint8
	BuffName      AddrTo[Hash]
	BuffDes       AddrTo[Hash]
	BuffImagePath AddrTo[StrWithPrefix16]
	IsShowBuffDes bool
}

type HoukaiTownBuildingMetaData struct {
	// Properties
	ModelPath       int32
	ModelType       uint8
	Blueprint       AddrTo[StrWithPrefix16]
	UItype          uint8
	IsCanSell       bool
	Price           int32
	IconPath        AddrTo[StrWithPrefix16]
	ImagePath       AddrTo[StrWithPrefix16]
	BuildingName    AddrTo[Hash]
	BuildingDes     AddrTo[Hash]
	EffectType      uint8
	TriggerHint     AddrTo[[]int32]
	Hp              int32
	BuildingLimit   int32
	BackgroundColor uint8
}

type HoukaiTownChallengeMetaData struct {
	// Properties
	ChallengeFinishWay uint8
	CompareType        uint8
	ChallengeParam     int32
	TargetValue        int32
	Textmap            AddrTo[Hash]
	ChallengeHint      AddrTo[[]int32]
	HintTextmap        AddrTo[Hash]
}

type HoukaiTownEventMetaData struct {
	// Properties
	EventName   AddrTo[Hash]
	EventDes    AddrTo[Hash]
	ImagePath   AddrTo[StrWithPrefix16]
	QavatarFace AddrTo[StrWithPrefix16]
}

type HoukaiTownItemMetaData struct {
	// Properties
	Type                uint8
	UseHintText         AddrTo[Hash]
	PurchaseLimitText   AddrTo[Hash]
	Param               AddrTo[[]int32]
	IsForbidOnFullBrick bool
	AutoUse             bool
}

type HoukaiTownLuckLevelMetaData struct {
	// Properties
	GetCoinRatio int32
	CriticalRate int32
}

type HoukaiTownMapMetaData struct {
	// Properties
	HonkaiTownsID         int32
	FloorID               int32
	QavatarList           AddrTo[[]int32]
	BuildingList          AddrTo[StrWithPrefix16]
	BossList              AddrTo[[]int32]
	BossBornLocation      int32
	InitialBuildingList   AddrTo[StrWithPrefix16]
	ChallengeList         AddrTo[[]HoukaiTownMapMetaData_Int2]
	BrickRefreshTime      int32
	BrickLimitNum         int32
	CoinMaterial          int32
	EnterTitle            AddrTo[Hash]
	EnterDialog           AddrTo[Hash]
	StageName             AddrTo[Hash]
	StageDes              AddrTo[Hash]
	StageImagePath        AddrTo[StrWithPrefix16]
	TalentPointRewardShow AddrTo[[]int32]
	TutorialStepID        int32
	FirstTimePlotID       int32
	WinPlotID             int32
	RpgSiteID             int32
	RecommendStrength     int32
	RecommendAttribute    int32
	RecommendTarget       AddrTo[[]int32]
	WebLink               AddrTo[StrWithPrefix16]
	StartScaleStrength    int32
	MaxScaleStrength      int32
	TalentMaterial        int32
	RefreshBrickCostList  AddrTo[[]int32]
	IsCanRefreshBrick     bool
}

type HoukaiTownMapMetaData_Int2 struct {
	// Fields
	Item1 int32
	Item2 int32
}

type HoukaiTownPathMetaData struct {
	// Properties
	DefaultNextPosition     int32
	AlterNextPosition       int32
	AlterCondition          int32
	AlterConditionParamList AddrTo[[]int32]
	ConditionText           AddrTo[Hash]
}

type HoukaiTownQAvatarMetaData struct {
	// Properties
	DormitoryAvatarID    int32
	Skill                AddrTo[[]int32]
	Strength             int32
	Speed                int32
	Luck                 int32
	Attribute            uint8
	HealTime             int32
	HealPricePerSecond   int32
	HealPriceUpgradeList AddrTo[[]int32]
	LevelType            uint8
	QavatarName          AddrTo[Hash]
	QavatarDesc          AddrTo[Hash]
	ImagePath            AddrTo[StrWithPrefix16]
	CardPicPath          AddrTo[StrWithPrefix16]
	ChibiIcons           AddrTo[StrWithPrefix16]
	L2DPath              AddrTo[StrWithPrefix16]
}

type HoukaiTownQBattleSkillMetaData struct {
	// Properties
	SkillPredicates      AddrTo[StrWithPrefix16]
	PredicatesPara       int32
	PredicatesTarget     AddrTo[StrWithPrefix16]
	PredicatesTargetPara AddrTo[[]int32]
	ActionTarget         AddrTo[StrWithPrefix16]
	ActionTargetPara     AddrTo[[]int32]
	SkillAction          AddrTo[StrWithPrefix16]
	ActionPara           int32
	IsOccupyARound       bool
	StateName            AddrTo[StrWithPrefix16]
	SkillName            AddrTo[Hash]
	SkillDes             AddrTo[Hash]
	SkillImagePath       AddrTo[StrWithPrefix16]
	IsShowBuffDes        bool
}

type HoukaiTownQbossMetaData struct {
	// Properties
	AttackBuildingDamage int32
	AttackInterval       int32
	BornTime             int32
	StateName            AddrTo[StrWithPrefix16]
}

type HoukaiTownQMonsterMetaData struct {
	// Properties
	ModelPath       int32
	ImagePath       AddrTo[StrWithPrefix16]
	Skill           AddrTo[[]int32]
	Strength        int32
	Attribute       uint8
	DisplayDropList AddrTo[[]int32]
	QmonsterName    AddrTo[Hash]
	LevelType       uint8
	L2DPath         AddrTo[StrWithPrefix16]
	ImagePath2      AddrTo[StrWithPrefix16]
}

type HoukaiTownRefreshMetaData struct {
	// Properties
	Type            uint8
	RefreshInterval int32
	TowerID         int32
}

type HoukaiTownSpeedLevelMetaData struct {
	// Properties
	MoveSpeed      int32
	CriticalDamage int32
}

type HoukaiTownStrengthLevelMetaData struct {
	// Properties
	HP  int32
	ATK int32
}

type HoukaiTownTerrainSkillMetaData struct {
	// Properties
	TerrainName    AddrTo[Hash]
	TerrainDes     AddrTo[Hash]
	ImagePath      AddrTo[StrWithPrefix16]
	QavatarSkillID AddrTo[[]int32]
}

type HoukaiTownTutorialStepMetaData struct {
	// Properties
	TriggerRound    int32
	TutorialType    uint8
	TutorialID      int32
	FinishWay       uint8
	FinishParam     AddrTo[[]int32]
	NotStopTutorial bool
	NextStepID      int32
	TutorialTextmap AddrTo[Hash]
	FailedTextmap   AddrTo[Hash]
	StartLocation   AddrTo[[]HoukaiTownTutorialStepMetaData_Pointer]
	TargetLocation  AddrTo[[]HoukaiTownTutorialStepMetaData_Pointer]
	TargetPrefab    AddrTo[[]HoukaiTownTutorialStepMetaData_Pointer]
}

type HoukaiTownTutorialStepMetaData_Pointer struct {
	// Objects
	Path StrWithPrefix16
	ID   int32
}

type HybridRelayPhaseMetaData struct {
	// Properties
	UpAvatarTagList AddrTo[[]uint16]
	MainEnemyList   AddrTo[[]uint8]
	Name            AddrTo[Hash]
	Desc            AddrTo[Hash]
}

type HybridSiteBossDataMetaData struct {
	// Properties
	StageID         uint32
	MinLevel        int32
	MaxLevel        int32
	StageHpAll      int32
	ChallengeOnceHp int32
}

type HybridSiteCameraMetaData struct {
	// Properties
	CameraSite uint32
	OffsetX    float32
}

type HybridSiteContentMetaData struct {
	// Properties
	Priority            uint32
	OpenTime            AddrTo[StrWithPrefix16]
	UnlockTime          AddrTo[StrWithPrefix16]
	CloseTime           AddrTo[StrWithPrefix16]
	RelationContentID   uint32
	ExclusiveGroupID    uint32
	Difficulty          int32
	HardLevelPicPath    AddrTo[StrWithPrefix16]
	LevelRecommendation int32
	DropLimitID         uint16
}

type HybridSiteDataMetaData struct {
	// Properties
	ChapterId           uint32
	ActID               uint32
	SiteType            uint32
	PreSite             AddrTo[[]uint32]
	PreSiteTower        AddrTo[[]HybridSiteDataMetaData_PreTower]
	SiteFirstContentID  uint32
	SiteContentID       AddrTo[[]uint32]
	SiteNameTextmapID   AddrTo[Hash]
	SiteBackgroundPath  AddrTo[StrWithPrefix16]
	SiteBackgroundPath2 AddrTo[StrWithPrefix16]
	RewardDisplayType   uint8
	RewardDisplay       uint32
	DisplayTimeLimit    AddrTo[StrWithPrefix16]
}

type HybridSiteDataMetaData_PreTower struct {
	// Fields
	PreSiteID    uint32
	PreSiteScore uint32
}

type HybridSiteRemindMetaData struct {
	// Properties
	RemindStart    AddrTo[StrWithPrefix16]
	RemindClose    AddrTo[StrWithPrefix16]
	LevelRequire   uint32
	RemindPreSite  uint32
	StopRemindSite uint32
}

type ImgPanelMetaData struct {
	// Properties
	SubType             uint8
	ImgIDList_Position1 AddrTo[[]int32]
	ImgIDList_Position2 AddrTo[[]int32]
	ImgIDList_Position3 AddrTo[[]int32]
}

type InControlActionMapMetaData struct {
	// Properties
	ActionGroup           int32
	ActionGroupString     AddrTo[StrWithPrefix16]
	BindingSettingsType   int32
	HideInMobile          bool
	ActionType            int32
	ActionSubType         int32
	IsSettingComboKey     bool
	DefaultKeys           AddrTo[StrWithPrefix16]
	DefaultInControlTypes AddrTo[[]StrWithPrefix16]
}

type InControlActionSubTypeInfoMetaData struct {
	// Properties
	SubTypeString AddrTo[StrWithPrefix16]
}

type InControlBattleOverrideMetaData struct {
	// Properties
	OverrideActionName AddrTo[StrWithPrefix16]
}

type InControlControlTypeInfoMetaData struct {
	// Properties
	IsSettingControlType bool
	PrefixControlList    AddrTo[[]InControlControlTypeInfoMetaData_PrefixControl]
	IconForSony          AddrTo[StrWithPrefix16]
	IconForXBox          AddrTo[StrWithPrefix16]
	IconForSwitch        AddrTo[StrWithPrefix16]
}

type InControlControlTypeInfoMetaData_PrefixControl struct {
	// Objects
	PrefixControlString StrWithPrefix16
	IsSettingKey        bool
}

type InControlKeyInfoMetaData struct {
	// Properties
	IsSettingKey  bool
	PrefixKeyList AddrTo[[]InControlKeyInfoMetaData_PrefixKey]
	KeyString     AddrTo[StrWithPrefix16]
}

type InControlKeyInfoMetaData_PrefixKey struct {
	// Objects
	PrefixKeyString StrWithPrefix16
	IsSettingKey    bool
}

type InControlUIButtonConfigMetaData struct {
	// Properties
	ButtonPath          AddrTo[StrWithPrefix16]
	KeyTipsType         int32
	IsDynamicallyLoaded bool
}

type InLevelBuffUIMetaData struct {
	// Properties
	IconPath       AddrTo[StrWithPrefix16]
	BGColor        int32
	Description    AddrTo[Hash]
	BuffNameDesc   AddrTo[Hash]
	DisplayType    int32
	DisplayTypeTwo int32
	IsActive       int32
	ShowStacking   bool
}

type InLevelDialogMetaData struct {
	// Properties
	PostDialogID  int32
	Duration      float32
	TxtDuration   float32
	AvatarName    AddrTo[Hash]
	AvatarID      int32
	DressID       int32
	FaceAnimation AddrTo[StrWithPrefix16]
	FaceEffect    AddrTo[StrWithPrefix16]
	AnimID        int32
	LipMotion     AddrTo[StrWithPrefix16]
	Content       AddrTo[Hash]
	AudioID       AddrTo[StrWithPrefix16]
	ScreenEffect  int32
	ScreenPic     AddrTo[StrWithPrefix16]
	ModelPosition AddrTo[[]float32]
	ModelRotation AddrTo[[]float32]
}

type InviteeActivityGobackConfigMetaData struct {
	// Properties
	InviteeReward       int32
	InviteeMissionGroup int32
}

type InviteeActivityNewbieConfigMetaData struct {
	// Properties
	InviteeReward       int32
	InviteeMissionGroup int32
}

type InviteeActivityScheduleMetaData struct {
	// Properties
	NewbieInviteeConfigID uint32
	GobackInviteeConfigID uint32
	EndTime               AddrTo[RemoteTime]
}

type InviterActivityRewardConfigMetaData struct {
	// Properties
	RewardID int32
}

type InviterActivityScheduleMetaData struct {
	// Properties
	InviterRewardConfigID int32
	InviteType            uint8
	EndTime               AddrTo[RemoteTime]
}

type JigsawGroupMetaData struct {
	// Properties
	PiecesID       AddrTo[[]int32]
	GroupRewardID  int32
	GroupRewardPic AddrTo[StrWithPrefix16]
	ShowReward     bool
}

type JigsawMetaData struct {
	// Properties
	Width          int32
	Height         int32
	PreJigsaw      int32
	UnlockTime     AddrTo[StrWithPrefix16]
	CostItemID     int32
	CostItemNum    int32
	FinishedReward int32
	JigsawBgPic    AddrTo[StrWithPrefix16]
	BackGroundPic  AddrTo[StrWithPrefix16]
	PuzzlePic      AddrTo[StrWithPrefix16]
	TagTextID      AddrTo[StrWithPrefix16]
	TagPic         AddrTo[StrWithPrefix16]
	TagMaterialID  int32
	RulesTextID    AddrTo[StrWithPrefix16]
}

type KingdomsAreaMetaData struct {
	// Properties
	PhaseID     int16
	RewardID    uint32
	BuffID      uint16
	PointList   AddrTo[[]uint32]
	BonusRenown uint16
	UIIndex     AddrTo[StrWithPrefix16]
}

type KingdomsBossPointScheduleMetaData struct {
	// Properties
	PointID       uint32
	AdvanceTime   uint32
	BeginTime     AddrTo[StrWithPrefix16]
	EndTime       AddrTo[StrWithPrefix16]
	Belief        AddrTo[[]uint16]
	RewardDisplay AddrTo[[]uint32]
	StageList     AddrTo[[]uint32]
}

type KingdomsMetaData struct {
	// Properties
	IconPath        AddrTo[StrWithPrefix16]
	FlagIconPath    AddrTo[StrWithPrefix16]
	QavatarIconPath AddrTo[StrWithPrefix16]
	Color1          AddrTo[StrWithPrefix16]
	Color2          AddrTo[StrWithPrefix16]
	ColorName       AddrTo[StrWithPrefix16]
	Qavatar         int32
	KingdomsName    AddrTo[Hash]
	PhotoTitle      AddrTo[Hash]
	PhotoDesc       AddrTo[StrWithPrefix16]
}

type KingdomsPhaseMetaData struct {
	// Properties
	PhaseEndTime     AddrTo[StrWithPrefix16]
	ChangeTime       AddrTo[StrWithPrefix16]
	TagChangeNum     uint16
	PhaseType        int16
	ServerRefreshCD  uint16
	StepNum          int32
	MinScore         int32
	RewardList       AddrTo[StrWithPrefix16]
	VoteGroup        int16
	VoteAddScore     int32
	RedPiontRenown   int16
	ShowPlayerNumMin uint32
	EnterTims        int32
	AddBelief        AddrTo[[]KingdomsPhaseMetaData_RankAddBelief]
}

type KingdomsPhaseMetaData_RankAddBelief struct {
	// Fields
	BeliefAddon uint16
	Rank        uint16
}

type KingdomsPhaseTypeMetaData struct {
	// Properties
	MapType          uint8
	IsMove           bool
	IsQavatarShow    bool
	RealTimeReportCD uint32
}

type KingdomsPointMetaData struct {
	// Properties
	UIIndex                      AddrTo[StrWithPrefix16]
	PointType                    uint8
	ParaList                     AddrTo[[]int32]
	NearPointList                AddrTo[[]int32]
	Renown                       uint16
	KingdomsID                   int16
	BeVoted                      bool
	PointTitle                   AddrTo[Hash]
	PointDesc                    AddrTo[Hash]
	ExchangeCostMaterialList     AddrTo[[]MechMetaData_DisjoinOutputItem]
	ExchangeStaminaNum           uint16
	ExchangeKindomsWarStaminaNum uint16
	ExchangeScore                uint32
	OccupyMinScore               uint32
	DisplayDifficult             uint8
	DanmakuID                    uint16
}

type KingdomsStageMetaData struct {
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
	// Properties
	GroupID  int16
	Type     int32
	Rank     int32
	RewardID int32
}

type KingdomsWarScheduleMetaData struct {
	// Properties
	PhotoOpenTime             AddrTo[StrWithPrefix16]
	MinLevel                  uint8
	MaxLevel                  uint8
	KingdomsList              AddrTo[[]int32]
	ScoreRewardGroup          int16
	RankRewardGroup           int16
	ScoreRewardMark           int32
	OpenEventID               uint32
	ShopType                  int16
	KingdomStaminaRecoverTime uint16
	KingdomStaminaLimit       uint8
	StaminaMaterialID         uint32
	DailyRewardMinScore       uint32
	MissionList               AddrTo[[]uint32]
	RefreshCD                 uint16
	WebLink                   AddrTo[StrWithPrefix16]
	IsPictureTutorialShow     bool
	MinReportCurrencyNum      uint32
	RewardShowID              int32
}

type KingdomsWarScoreRewardMetaData struct {
	// Properties
	GroupID  int16
	Score    int32
	RewardID int32
}

type KingdomsWarStageCurrencyMetaData struct {
	// Properties
	PhaseID     uint32
	MinScore    uint32
	CurrencyID  uint32
	CurrencyNum uint32
}

type KingdomsWarStoryMetaData struct {
	// Properties
	PhotoType       int16
	PhotoID         int32
	PrePhotoID      int32
	PointID         uint32
	KingdomsID      int16
	NeedScore       int32
	NeedCurrencyID  int32
	NeedCurrencyNum int32
	BeginTime       AddrTo[StrWithPrefix16]
	EndTime         AddrTo[StrWithPrefix16]
	LockText        AddrTo[Hash]
	IsShow          int32
}

type KingdomsWarVoteMetaData struct {
	// Properties
	Score    uint32
	Addition int16
	GroupID  int16
}

type KingdomsWarWaveMetaData struct {
	// Properties
	CostItem    int32
	Num         int32
	StaminaCost int32
	AddScore    int32
	DisplayID   AddrTo[StrWithPrefix16]
}

type LevelChallengeMetaData struct {
	// Properties
	ConditionId   int32
	ParamList     AddrTo[[]int32]
	HintPeriod    int32
	DiaplayTarget AddrTo[Hash]
	Explanation   AddrTo[Hash]
	Difficulty    int32
}

type LevelChallengeMetaData_RawData struct {
	// Fields
	ChallengeId int32
	ConditionId int32
	HintPeriod  int32
	Difficulty  int32

	// Objects
	ParamList     []int32
	DiaplayTarget Hash
	Explanation   Hash
}

type LevelMetaData struct {
	// Properties
	Name                       AddrTo[Hash]
	ChapterId                  uint8
	ActId                      uint16
	SectionId                  uint8
	Difficulty                 uint8
	Type                       uint8
	Tag                        AddrTo[[]int32]
	BattleType                 uint8
	EnterTimes                 uint16
	ResetType                  uint8
	ResetCoinID                int32
	ResetCostType              uint8
	ResetTimes                 uint8
	StaminaCost                uint8
	AvatarExpReward            uint16
	AvatarExpInside            float32
	ScoinReward                int32
	ScoinInside                float32
	MaxScoinReward             int32
	MaxProgress                uint8
	HighlightDisplayDropIdList AddrTo[[]int32]
	DropList                   AddrTo[[]int32]
	RecommendPlayerLevel       uint8
	UnlockPlayerLevel          uint8
	UnlockStarNum              uint8
	PreLevelID                 AddrTo[[]int32]
	DisplayTitle               AddrTo[Hash]
	DisplayDetail              AddrTo[Hash]
	BriefPicPath               AddrTo[StrWithPrefix16]
	DetailPicPath              AddrTo[StrWithPrefix16]
	LuaFile                    AddrTo[StrWithPrefix16]
	ChallengeList              AddrTo[[]LevelMetaData_LevelChallengeMetaNode]
	IsActChallenge             bool
	FastBonusTime              uint16
	SonicBonusTime             uint16
	HardLevel                  uint16
	HardLevelGroup             uint16
	ReviveTimes                uint16
	ReviveCostType             uint8
	ReviveUseTypeList          AddrTo[[]uint16]
	TeamNum                    uint8
	MaxNumList                 AddrTo[[]uint8]
	RestrictList               AddrTo[[]uint16]
	LoseDescList               AddrTo[[]Hash]
	RecordLevelType            uint8
	UseDynamicHardLv           uint8
	IsTrunk                    bool
	MonsterAttrShow            AddrTo[[]uint8]
	PlayerGetAllDrops          bool
	HardCoeff                  uint16
	EnterTimesType             uint8
	IsEnterWithElf             uint8
	PreMissionList             AddrTo[[]int32]
	LockedText                 AddrTo[Hash]
	PreMissionLink             uint16
	PreMissionLinkParams       AddrTo[[]int32]
	PreMissionLinkParamStr     AddrTo[StrWithPrefix16]
	UnlockedText               AddrTo[Hash]
	UnlockedLink               uint16
	UnlockedLinkParams         AddrTo[[]int32]
	UnlockedLinkParamStr       AddrTo[StrWithPrefix16]
	CostMaterialId             int32
	CostMaterialNum            int32
	FirstCostMaterialNum       int32
	BalanceModeType            int32
	StageEntryNameList         AddrTo[[]StrWithPrefix16]
	FloatDrop                  AddrTo[[]LevelMetaData_FloatDropParam]
	IsBattleYLevel             bool
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
	// Properties
	Type1 AddrTo[[]int32]
}

type LevelSaveMetaData struct {
	// Properties
	Icon        AddrTo[StrWithPrefix16]
	Description AddrTo[Hash]
}

type LevelStatisticsMetaData struct {
	// Properties
	ModuleID      int32
	TypeInt       int32
	HintPeriod    int32
	HintRank      int32
	DisplayTarget AddrTo[Hash]
	OverrideList  AddrTo[[]LevelStatisticsMetaData_IntIntPair]
}

type LevelStatisticsMetaData_IntIntPair struct {
	// Objects
	Item1 StrWithPrefix16
	Item2 float32
}

type LevelTrialMetaData struct {
	// Properties
	StageID             int32
	AvatarList          AddrTo[[]int32]
	AvatarTrialGroup    AddrTo[[]LevelTrialMetaData_AvatarTrialLimit]
	AvatarTrialType     int32
	ElfTrialID          uint8
	FakeAvatarID        int32
	AllowedPlayerAvatar uint8
	AvatarDress         AddrTo[[]LevelTrialMetaData_DressInfo]
	IsAutoSelect        bool
	IsSaveTeam          bool
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
	// Properties
	LevelId       int32
	ParamList     AddrTo[[]float32]
	DiaplayTarget AddrTo[[]Hash]
}

type LinearMissionData struct {
	// Properties
	PreMissionId         int32
	AchievementTag       uint8
	AchievementType      uint8
	WikiAchievementScore uint8
}

type LinkDataMetaData struct {
	// Properties
	IsShow      bool
	LinkType    uint32
	LinkParams  AddrTo[StrWithPrefix16]
	URL         AddrTo[StrWithPrefix16]
	LinkTextMap AddrTo[StrWithPrefix16]
}

type LoadingTipMetaData struct {
	// Properties
	ImgPath                AddrTo[StrWithPrefix16]
	ImgPath_2              AddrTo[StrWithPrefix16]
	MinLevel               int32
	MaxLevel               int32
	LoadingTitle           AddrTo[Hash]
	LoadingTagList         AddrTo[[]int32]
	LoadingTipsList        AddrTo[[]int32]
	Weight                 int32
	BeginTime              AddrTo[StrWithPrefix16]
	EndTime                AddrTo[StrWithPrefix16]
	IsIgnoreZeroTag        int32
	ExtraTriggerTypeList   AddrTo[[]int32]
	ExtraTriggerParamsList AddrTo[[]int32]
}

type LoadingTipStringMetaData struct {
	// Properties
	TipTextID AddrTo[Hash]
}

type LoginWishActivityLinkMetaData struct {
	// Properties
	BeginTime    AddrTo[RemoteTime]
	EndTime      AddrTo[RemoteTime]
	LinkMinLevel uint8
	LinkMaxLevel uint8
	LinkPic      AddrTo[StrWithPrefix16]
}

type LoginWishActivityMetaData struct {
	// Properties
	UnselectPicPath        AddrTo[StrWithPrefix16]
	SpecialUnselectPicPath AddrTo[StrWithPrefix16]
	TextmapRule            AddrTo[StrWithPrefix16]
	PanelBGColor           AddrTo[StrWithPrefix16]
	PaneldecoratingColor   AddrTo[StrWithPrefix16]
	Logindayisopen         bool
}

type LoginWishActivityRewardMetaData struct {
	// Properties
	LoginReward   int32
	SpecialReward int32
	WishId        uint8
	LoginPic      AddrTo[StrWithPrefix16]
}

type LoginWishActivityWishMetaData struct {
	// Properties
	SpecialWish         bool
	CandidateRewardList AddrTo[[]int32]
	RewardDayOffset     uint8
}

type MailDataMetaData struct {
	// Properties
	MailSender    AddrTo[StrWithPrefix16]
	MailTitle     AddrTo[StrWithPrefix16]
	MailContent   AddrTo[StrWithPrefix16]
	ImgPath       AddrTo[StrWithPrefix16]
	MailStyle     int32
	MailLabelPic  AddrTo[StrWithPrefix16]
	MailLabelText AddrTo[Hash]
}

type MailStyleMetaData struct {
	// Properties
	BaseColor       AddrTo[StrWithPrefix16]
	SelectColor     AddrTo[StrWithPrefix16]
	QuickDelete     bool
	DeleteConfirm   int32
	BanRewardGetAll bool
}

type MainEventPanelMetaData struct {
	// Properties
	ImgIDList                   AddrTo[[]int32]
	GoToLink_Raffle             AddrTo[StrWithPrefix16]
	ShopLink                    AddrTo[[]int32]
	GenActivityRewardScheduleID uint16
}

type MainlineStepMissionDataMetaData struct {
	// Properties
	MissionIDList    AddrTo[[]int32]
	ForceClosedLevel int32
}

type MainlineStepMissionDisplayDataMetaData struct {
	// Properties
	SectionType            int32
	UnlockStepID           int32
	SectionDisplayReward   int32
	SectionDisplayLevel    int32
	SectionDisplayStr      AddrTo[Hash]
	SectionDisplayHonor    AddrTo[Hash]
	SectionDisplayIconPath AddrTo[StrWithPrefix16]
	SectionDescription     AddrTo[Hash]
}

type MainPageActivityEntryMetaData struct {
	// Properties
	EntryPrefab AddrTo[StrWithPrefix16]
	BeginTime   AddrTo[RemoteTime]
	EndTime     AddrTo[RemoteTime]
	MinLevel    uint8
	MaxLevel    uint8
	Priority    uint8
	Type        uint8
	Param1      AddrTo[[]int32]
}

type MainPageSequenceDialogInfoMetaData struct {
	// Properties
	Priority    int32
	CheckType   AddrTo[[]int32]
	ShowType    int32
	Review      bool
	NewbieGuide bool
}

type MainstageEventMetaData struct {
	// Properties
	ChapterID           int32
	BeginShowTime       AddrTo[RemoteTime]
	EndTime             AddrTo[RemoteTime]
	UnlockLevel         int32
	EnterIconPrefabPath AddrTo[StrWithPrefix16]
	RewardBGPath        AddrTo[StrWithPrefix16]
	ShowPointReward     int32
	CurrencyList        AddrTo[[]int32]
	CycleMissionIDList  AddrTo[[]int32]
	Ticket              AddrTo[[]int32]
	TicketDailyNumLimit AddrTo[[]int32]
	ShopLink            AddrTo[[]int32]
}

type MallEntranceDataMetaData struct {
	// Properties
	MallType        int32
	MallTabBGPath   AddrTo[StrWithPrefix16]
	Name            AddrTo[Hash]
	SortID          int32
	CurrencyBarList AddrTo[[]int32]
	MallTypeGroup   uint8
	BeginTime       AddrTo[StrWithPrefix16]
	EndTime         AddrTo[StrWithPrefix16]
}

type MapGrowthMetaData struct {
	// Properties
	MapGrowthName                     AddrTo[Hash]
	MapGrowthInfo                     AddrTo[Hash]
	PrefabPath                        AddrTo[StrWithPrefix16]
	ServantChapterDescription         AddrTo[Hash]
	UpgradeResultDialogTitleIcon      AddrTo[StrWithPrefix16]
	UpgradeResultDialogTitleIconColor AddrTo[StrWithPrefix16]
}

type MapLevelMetaData struct {
	// Properties
	Exp          int32
	Atk          int32
	SkillPoint   int32
	KeyQuestID   int32
	KeyQuestDesc AddrTo[Hash]
}

type MapSiteQualityConfigMetaData struct {
	// Properties
	QualityIcon       AddrTo[StrWithPrefix16]
	QualityFrameColor AddrTo[StrWithPrefix16]
	MaterialID        int32
	MaterialDrop      int32
	ChargeIMaterial   int32
	WarningLevel      int32
}

type MapSiteScheduleDataMetaData struct {
	// Properties
	IsDailyRaid    bool
	ScoreType      int32
	ScoreTypePara  int32
	ActivityType   int32
	DailyTimesList AddrTo[[]int32]
	StageGroupList AddrTo[[]int32]
}

type MapSiteStageDataMetaData struct {
	// Properties
	PreStageShowList   AddrTo[[]int32]
	PreStageList       AddrTo[[]int32]
	PreQuality         bool
	UnlockDec          AddrTo[StrWithPrefix16]
	MapSitePosition    AddrTo[[]int32]
	SiteQuality        AddrTo[[]int32]
	RandomEffectCharge int32
	MapSiteType        int32
	SiteBg             AddrTo[StrWithPrefix16]
}

type MapSkillMetaData struct {
	// Properties
	MapID         int32
	MapXY         AddrTo[[]int32]
	PreSkillID    int32
	UnlockMapLv   int32
	UnlockStoryId int32
	UnlockType    int32
	UnlockCost    int32
	SkillType     int32
	SkillPar      AddrTo[StrWithPrefix16]
	SkillIconPath AddrTo[StrWithPrefix16]
	SkillName     AddrTo[Hash]
	SkillInfo     AddrTo[Hash]
}

type MapSkillTypeMetaData struct {
	// Properties
	SkillTypeSeries int32
	SkillTypeName   AddrTo[Hash]
	SkillTypeImg    AddrTo[StrWithPrefix16]
	SkillTypeIcon   AddrTo[StrWithPrefix16]
	SkillTypeColor  AddrTo[StrWithPrefix16]
}

type MassiveWarDamageNeedMetaData struct {
	// Properties
	DamageBasic    int32
	DamageNeedList AddrTo[[]MassiveWarDamageNeedMetaData_DamageNeed]
}

type MassiveWarDamageNeedMetaData_DamageNeed struct {
	// Fields
	Rank   uint8
	Damage int32
}

type MassiveWarDamageRewardMetaData struct {
	// Properties
	DamageRankIconPath AddrTo[StrWithPrefix16]
	DamageRankText     AddrTo[Hash]
	ScoreDamageList    AddrTo[[]MassiveWarDamageRewardMetaData_RewardScore]
}

type MassiveWarDamageRewardMetaData_RewardScore struct {
	// Fields
	PlayerGroupID uint8
	Score         int32
}

type MassiveWarMonsterDataMetaData struct {
	// Properties
	BossName             AddrTo[StrWithPrefix16]
	BossPrefabPath       AddrTo[StrWithPrefix16]
	MessageId            uint8
	StageDetailMonsterId int32
	BossDetail           AddrTo[StrWithPrefix16]
}

type MassiveWarPlayerGroupMetaData struct {
	// Properties
	MinPlayerLevel uint8
	MaxPlayerLevel uint8
}

type MassiveWarRankRewardMetaData struct {
	// Properties
	MaxRankRatio  uint16
	ScoreRankList AddrTo[[]MassiveWarDamageRewardMetaData_RewardScore]
}

type MassiveWarRewardMetaData struct {
	// Properties
	NeedScore int32
	RewardID  int32
}

type MassiveWarScheduleMetaData struct {
	// Properties
	MaxEnterTimes       uint8
	ScoreBasicShowList  AddrTo[[]MassiveWarScheduleMetaData_ScoreRange]
	ScoreRankShowList   AddrTo[[]MassiveWarScheduleMetaData_ScoreRange]
	ScoreDamageShowList AddrTo[[]MassiveWarScheduleMetaData_ScoreRange]
	RewardLinkType      int32
	RewardLinkParams    int32
	LinkBtnImgPath      AddrTo[StrWithPrefix16]
}

type MassiveWarScheduleMetaData_ScoreRange struct {
	// Fields
	PlayerGroupID uint8
	MaxScore      int32
	MinScore      int32
}

type MassiveWarStageInfoMetaData struct {
	// Properties
	IconPath        AddrTo[StrWithPrefix16]
	DescriptionText AddrTo[StrWithPrefix16]
}

type MassiveWarStageMetaData struct {
	// Properties
	DamageNeedRatio           float32
	MonsterWaveList           AddrTo[[]GratuityStageMetaData_MessageIDWaveData]
	BossIdList                AddrTo[[]uint8]
	StageImagePath            AddrTo[StrWithPrefix16]
	StageImagePathRankingPage AddrTo[StrWithPrefix16]
	StageInfoIdList           AddrTo[[]int32]
}

type MatchThreeActivityMetaData struct {
	// Properties
	InitialRoleIDList    AddrTo[[]MatchThreeActivityMetaData_InitialRoleInfo]
	InitialItemIDList    AddrTo[[]int32]
	GachaMaterialList    AddrTo[MatchThreeActivityMetaData_GachaMaterialCosts]
	ExperienceMaterialID int32
	DailyUnlockSite      int32
	GachaUnlockSite      int32
	TowerUnlockSite      int32
	MedalID              int32
	GachaPoolList        AddrTo[[]uint8]
	LimitItemNum         uint8
	NormGem              uint16
	NormGrid             uint16
	RatingScore          AddrTo[[]int32]
	RoundScore           AddrTo[[]int32]
	WeightLimit          uint16
	TutorialGemState1    uint16
	TutorialGemState2    uint16
	TutorialObstacle     uint16
}

type MatchThreeActivityMetaData_GachaMaterialCosts struct {
	// Fields
	MateriaId  int32
	MateriaNum int32
}

type MatchThreeActivityMetaData_InitialRoleInfo struct {
	// Fields
	RoleID    int32
	RoleLevel int32
}

type MatchThreeBuffMetaData struct {
	// Properties
	Priority              int32
	IsHide                bool
	BuffIcon              AddrTo[StrWithPrefix16]
	BuffName              AddrTo[Hash]
	BuffDesc              AddrTo[Hash]
	ShowType              uint8
	IsDisplayed1          bool
	BuffSustainEffectPath AddrTo[StrWithPrefix16]
	NodeName              AddrTo[StrWithPrefix16]
}

type MatchThreeDailyScheduleMetaData struct {
	// Properties
	UnlockSiteID          int32
	StartTime             AddrTo[RemoteTime]
	EndTime               AddrTo[RemoteTime]
	RoomID                int32
	DailyMaterialNumLimit AddrTo[[]ChapterBWWorldStageMetaData_DropExpMaterial]
}

type MatchThreeDailyScoreMetaData struct {
	// Properties
	RoomID               int32
	DailyMinScore        int32
	DailyMinScorePercent int32
	DailyMinScoreRank    uint8
	DailyMaterialReward  AddrTo[[]MatchThreeDailyScoreMetaData_DailyRewardItem]
}

type MatchThreeDailyScoreMetaData_DailyRewardItem struct {
	// Fields
	MateriaID  int32
	MateriaNum int32
}

type MatchThreeDynamicDifficultyMetaData struct {
	// Properties
	IncreaseTakeDamage int32
	ReduceGetDamage    int32
	ReducePowerWeight  int32
}

type MatchThreeEndConditionMetaData struct {
	// Properties
	ResultType         uint8
	ResultText         AddrTo[Hash]
	MonsterDie         bool
	RoleDie            bool
	ReachMaxRound      bool
	ReachSpecificRound int32
	ReachMaxShuffle    bool
	CleanStateList     AddrTo[[]MatchThreeEndConditionMetaData_CleanState]
	GoalStateList      AddrTo[[]MatchThreeEndConditionMetaData_GoalState]
}

type MatchThreeEndConditionMetaData_CleanState struct {
	// Fields
	StateID   int32
	StateTime int32
}

type MatchThreeEndConditionMetaData_GoalState struct {
	// Fields
	StateID  int32
	StateNum int32
}

type MatchThreeGachaPoolMetaData struct {
	// Properties
	UnlockSiteID      int32
	GachaPoolItemList AddrTo[[]MatchThreeGachaPoolMetaData_GemGachaPoolWeight]
	PoolName          AddrTo[Hash]
}

type MatchThreeGachaPoolMetaData_GemGachaPoolWeight struct {
	// Fields
	GemID     uint8
	GemWeight uint8
}

type MatchThreeGemStateMetaData struct {
	// Properties
	StateType                uint8
	StateName                AddrTo[Hash]
	StateDes                 AddrTo[Hash]
	StateIcon                AddrTo[StrWithPrefix16]
	IsUnique                 bool
	CanDestroy               bool
	CleanTime                uint16
	Countdown                uint16
	CleanMethod              uint8
	NextState                uint16
	IsMovable                bool
	IsDroppable              bool
	TriggerSkill             AddrTo[[]MatchThreeGemStateMetaData_GemStateSkillTriggers]
	EffectEnterAnimationName AddrTo[StrWithPrefix16]
	EffectExitAnimationName  AddrTo[StrWithPrefix16]
	IsPlayRootExitAnimation  bool
	ExitEffectPath           AddrTo[StrWithPrefix16]
	EffectPath               AddrTo[[]StrWithPrefix16]
	FlyableEffectPath        AddrTo[StrWithPrefix16]
	ShowInfo                 bool
	ShowNewInfo              bool
	ExplodePath              AddrTo[StrWithPrefix16]
	EnterAudioPath           AddrTo[StrWithPrefix16]
	ExitAudioPath            AddrTo[StrWithPrefix16]
}

type MatchThreeGemStateMetaData_GemStateSkillTriggers struct {
	// Fields
	TriggerType uint8
	SkillId     int32
}

type MatchThreeGridStateMetaData struct {
	// Properties
	StateName    AddrTo[Hash]
	StateDes     AddrTo[Hash]
	StateIcon    AddrTo[StrWithPrefix16]
	IsChangeable bool
	ChangeTime   uint16
	NextState    uint16
}

type MatchThreeItemMetaData struct {
	// Properties
	ItemIcon       AddrTo[StrWithPrefix16]
	ItemFigure     AddrTo[StrWithPrefix16]
	ItemName       AddrTo[Hash]
	ItemDes        AddrTo[Hash]
	ItemCD         uint8
	ItemCost       int32
	ItemSkill      int32
	UseType        uint8
	UnlockSiteID   int32
	UseAudioPath   AddrTo[StrWithPrefix16]
	EffectPath     AddrTo[StrWithPrefix16]
	TargetPosition uint8
}

type MatchThreeMonsterMetaData struct {
	// Properties
	MonsterName       AddrTo[Hash]
	MonsterIcon       AddrTo[StrWithPrefix16]
	MonsterFigurePath AddrTo[StrWithPrefix16]
	MonsterPrefabPath AddrTo[StrWithPrefix16]
	MonsterLevel      int32
	Hp                int32
	Attack            int32
	Defence           int32
	Resist            int32
	Stars             AddrTo[[]int32]
	SkillList         AddrTo[[]int32]
	SkillSequence     AddrTo[[]int32]
	IsBig             bool
	IsMirror          bool
	ShieldOffset      int32
}

type MatchThreeRatingMetaData struct {
	// Properties
	AdditionPercent int32
}

type MatchThreeRoleMetaData struct {
	// Properties
	RoleName       AddrTo[Hash]
	RoleChibiIcon  AddrTo[StrWithPrefix16]
	RoleFigurePath AddrTo[StrWithPrefix16]
	RolePrefabPath AddrTo[StrWithPrefix16]
	PassiveName    AddrTo[Hash]
	PassiveDes     AddrTo[Hash]
	MaxLevel       int32
	UnlockSiteID   int32
	Experience     int32
	Health         int32
	Energy         int32
	HealthMax      int32
	EnergyMax      int32
	Attack         int32
	Element        int32
	GemList        AddrTo[[]uint8]
	BeginSkillID   AddrTo[[]int32]
}

type MatchThreeRoleUltraMetaData struct {
	// Properties
	UltraPath        AddrTo[StrWithPrefix16]
	UltraObject      uint8
	FinalGemMaskPath AddrTo[StrWithPrefix16]
	NodeName         AddrTo[StrWithPrefix16]
	ShakeDelay       float32
}

type MatchThreeRoomMetaData struct {
	// Properties
	NextRoomId        int32
	RoomType          uint8
	RoleTrialID       uint8
	RoleRecommendList AddrTo[[]int32]
	LeastDynamicLevel int32
	RoleRestrictList  AddrTo[[]int32]
	ItemRestrictList  AddrTo[[]int32]
	MonsterList       AddrTo[[]int32]
	MaxRound          int32
	UnlockRoleList1   AddrTo[[]MatchThreeRoomMetaData_UnlockRoleInfo]
	UnlockItemList    AddrTo[[]int32]
	EndCondition      AddrTo[[]uint8]
	EnvironmentSkill  AddrTo[[]int32]
	RoomTitle         AddrTo[Hash]
	RoomDes           AddrTo[Hash]
	TipsTitle         AddrTo[Hash]
	TipsDes           AddrTo[Hash]
	BgPath            AddrTo[StrWithPrefix16]
	FristDropReward   int32
	RpgSiteID         int32
	BeginPlotID       int32
	PlayBgm           AddrTo[StrWithPrefix16]
}

type MatchThreeRoomMetaData_UnlockRoleInfo struct {
	// Fields
	Value1 uint8
	Value2 int32
}

type MatchThreeRoomTutorialMetaData struct {
	// Properties
	TutorialDataID    int32
	SkillOnRoundStart int32
}

type MatchThreeShakeLevelMetaData struct {
	// Properties
	Magnitude   float32
	Speed       float32
	Duration    float32
	RandomRange float32
}

type MatchThreeSkillGemMetaData struct {
	// Properties
	SkillGemIcon   AddrTo[StrWithPrefix16]
	SkillGemFigure AddrTo[StrWithPrefix16]
	SkillGemName   AddrTo[Hash]
	SkillGemType   uint8
	SkillGemDesc   AddrTo[Hash]
	BasicGemIcon   AddrTo[StrWithPrefix16]
	BasicGemName   AddrTo[Hash]
	BasicGemDesc   AddrTo[Hash]
	RoleID         uint8
	MaxLevel       uint8
	SkillIDList    AddrTo[[]int32]
}

type MatchThreeSkillMetaData struct {
	// Properties
	SkillName       AddrTo[Hash]
	SkillDes        AddrTo[Hash]
	SkillIcon       AddrTo[StrWithPrefix16]
	IsNormalAttack  bool
	ResultOnExecute AddrTo[[]StrWithPrefix16]
	DisplayType     uint8
	ActionTrigger   AddrTo[StrWithPrefix16]
	ShakeLevel      int32
}

type MatchThreeTowerMetaData struct {
	// Properties
	FloorID             uint8
	EnemyID             int32
	Medal               int32
	FloorMaterialReward int32
}

type MatchThreeTrialMetaData struct {
	// Properties
	RoleID   uint8
	Level    uint8
	Gemskill AddrTo[[]int32]
}

type MaterialProtectMetaData struct {
}

type MaterialProvideElfExpMetaData struct {
	// Properties
	ProvideExp int32
	ScoinCost  int32
}

type MaterialUseConversionMetaData struct {
	// Properties
	ItemID   int32
	CostNum  int32
	RewardID int32
}

type MaterialUseMetaData struct {
	// Properties
	UseType                int32
	MultiUse               int32
	MaterialUseConfirmType uint8
	ParaStr                AddrTo[[]StrWithPrefix16]
	UseTip                 AddrTo[Hash]
	UseMin                 int32
	UseMax                 int32
	AdditionDesc           AddrTo[StrWithPrefix16]
	EventID                int32
}

type MatrixFloorMetaData struct {
	// Properties
	ThemeID        int32
	MapDataID      int32
	FloorNameText  AddrTo[Hash]
	CfgName        AddrTo[StrWithPrefix16]
	AudioNameEnter AddrTo[StrWithPrefix16]
	AudioNameExit  AddrTo[StrWithPrefix16]
}

type MatrixGridIndexMetaData struct {
	// Properties
	ModPath           AddrTo[StrWithPrefix16]
	EffectPatternName AddrTo[StrWithPrefix16]
	HaveFloor         bool
	ClientParam       int32
	ModScale          AddrTo[Vector3]
	PlaceType         int32
	PlaceParam        float32
	CanRotate         int32
}

type MatrixMapMetaData struct {
	// Properties
	ForkList AddrTo[StrWithPrefix16]
}

type MatrixPanelLinkMetaData struct {
	// Properties
	LinkType     int32
	LinkParams   AddrTo[[]int32]
	LinkParamStr AddrTo[[]StrWithPrefix16]
	TextMapList  AddrTo[[]Hash]
}

type MatrixSuddenEventMetaData struct {
	// Properties
	EventType      uint8
	ParamsVar      AddrTo[[]int32]
	EventName      AddrTo[Hash]
	EventDesc      AddrTo[Hash]
	EventImagePath AddrTo[StrWithPrefix16]
}

type MatrixThemeMetaData struct {
	// Properties
	FloorMod    AddrTo[[]StrWithPrefix16]
	BgMod       AddrTo[StrWithPrefix16]
	WallEdge1   AddrTo[[]StrWithPrefix16]
	WallEdge2   AddrTo[[]StrWithPrefix16]
	WallEdge3   AddrTo[[]StrWithPrefix16]
	WallEdge4   AddrTo[[]StrWithPrefix16]
	WallCorner1 AddrTo[StrWithPrefix16]
	WallCorner2 AddrTo[StrWithPrefix16]
	WallCorner3 AddrTo[[]StrWithPrefix16]
	WallCorner4 AddrTo[StrWithPrefix16]
}

type MechPaperMetaData struct {
	// Properties
	InfoTextMap  AddrTo[Hash]
	UnlockLevel  int32
	PaperIcon    AddrTo[StrWithPrefix16]
	NeedMaterial AddrTo[[]MechMetaData_DisjoinOutputItem]
}

type MechTDDiffcultyMetaData struct {
	// Properties
	Difficulty int32
	Weather    int32
}

type MechTDRewardMetaData struct {
	// Properties
	Stage         int32
	Difficulty    int32
	Icon          AddrTo[StrWithPrefix16]
	Desc          AddrTo[Hash]
	TimeDisplayed AddrTo[Hash]
	Reward        int32
}

type MechTDWeatherMetaData struct {
	// Properties
	WeatherName AddrTo[Hash]
	WeatherIcon AddrTo[StrWithPrefix16]
	WeatherDesc AddrTo[Hash]
}

type MemoirsAudioataMetaData struct {
	// Properties
	AudioPath     AddrTo[StrWithPrefix16]
	AudioDesc     AddrTo[Hash]
	AudioInterval float32
}

type MemoirsCatalogMetaData struct {
	// Properties
	IsHideCata bool
	Type       uint8
	PageIDList AddrTo[[]uint16]
	Title      AddrTo[Hash]
}

type MemoirsPageMetaData struct {
	// Properties
	Type                     uint8
	PageTitle                AddrTo[Hash]
	BeginTime                AddrTo[RemoteTime]
	Endtime                  AddrTo[RemoteTime]
	NewAvatarIDList          AddrTo[[]int32]
	NewArtifactDressIDList   AddrTo[[]int32]
	BGMConfigName            AddrTo[StrWithPrefix16]
	Model1Title              AddrTo[Hash]
	Model1Desc               AddrTo[Hash]
	Model1ImagePath          AddrTo[StrWithPrefix16]
	Model1BGPath             AddrTo[StrWithPrefix16]
	Model1MemoirsGroupIDList AddrTo[[]uint16]
	Model1AudioIconPath      AddrTo[StrWithPrefix16]
	Model1AudioIDList        AddrTo[[]uint16]
	Model1CGID               int32
	Model2Title              AddrTo[Hash]
	Model2Desc               AddrTo[Hash]
	Model2ImagePath          AddrTo[StrWithPrefix16]
	Model2BGPath             AddrTo[StrWithPrefix16]
	Model2MemoirsDescIDList  AddrTo[[]uint16]
	Model2AudioIconPath      AddrTo[StrWithPrefix16]
	Model2AudioList          AddrTo[[]uint16]
	Model2CGID               int32
}

type MemoirsSpecialDataGroupMetaData struct {
	// Properties
	SpecialDataIDList AddrTo[[]uint16]
	ConditionType     uint8
	Condition         AddrTo[StrWithPrefix16]
	Desc              AddrTo[Hash]
}

type MentorCompanyMissionRewardMetaData struct {
	// Properties
	StudentRewardID int32
	MentorRewardID  int32
}

type MentorConst struct {
	// Properties
	Type  AddrTo[StrWithPrefix16]
	Value AddrTo[StrWithPrefix16]
}

type MentorDailyMissionReward struct {
	// Properties
	IdRewardStudent int32
	IdRewardMentor  int32
}

type MentorEvaluate struct {
	// Properties
	EvaluateTextmap AddrTo[Hash]
}

type MentorExamData struct {
	// Properties
	NextExamID        int32
	NameExam          AddrTo[Hash]
	IdRewardStudent   int32
	IdRewardMentor    int32
	RefLevel          int32
	IdExamMissionList AddrTo[[]int32]
}

type MentorFameLevelData struct {
	// Properties
	FameForNextLevel int32
	RewardObtain     int32
	LevelTitle       AddrTo[Hash]
	LevelDesc        AddrTo[Hash]
}

type MentorLobbyTagDisplay struct {
	// Properties
	PathIconStudent AddrTo[StrWithPrefix16]
	PathIconMentor  AddrTo[StrWithPrefix16]
	IsDisplay       int32
}

type MentorMainSceneNotice struct {
	// Properties
	NoticeTextMap AddrTo[Hash]
}

type MentorRewardAttenuation struct {
	// Properties
	RatioExamStudent       float32
	RatioExamMentor        float32
	RatioDailyStudent      float32
	RatioDailyMentor       float32
	AttenuationType        int32
	AttenuationTipsTextmap AddrTo[Hash]
}

type MiniBossExamExplainInfoMetaData struct {
	// Properties
	OpenDate AddrTo[StrWithPrefix16]
	Desc     AddrTo[StrWithPrefix16]
}

type MiniBossExamHardStageDataMetaData struct {
	// Properties
	TagList       AddrTo[[]MiniBossExamHardStageDataMetaData_Tag]
	MissionIDList AddrTo[[]int32]
}

type MiniBossExamHardStageDataMetaData_Tag struct {
	// Fields
	TagID int32

	// Objects
	TagDesc Hash
	TagName Hash
	TagIcon StrWithPrefix16
}

type MiniMonopolyActivityMetaData struct {
	// Properties
	ActivityTitle       AddrTo[Hash]
	ActivityDescription AddrTo[Hash]
	MapPool             AddrTo[[]uint16]
	StartMapId          uint16
	CreditMaterialID    int32
	RewardLineID        int32
	ThrowDiceScore      uint16
	DiceMaterialID      int32
	DiceSpeed           uint16
	SectorIdList        AddrTo[[]uint16]
	ItemIDList          AddrTo[[]uint16]
	MissionList         AddrTo[[]int32]
	AvatarImagePath     AddrTo[StrWithPrefix16]
}

type MiniMonopolyGridMetaData struct {
	// Properties
	GridType       uint8
	RewardId       int32
	GridMaterialID int32
	IsKeyReward    bool
	AvatarOrient   uint8
}

type MiniMonopolyItemMetaData struct {
	// Properties
	ItemType   uint8
	MaterialID int32
}

type MiniMonopolyMapMetaData struct {
	// Properties
	NextMapId     uint16
	MapPrefabPath AddrTo[StrWithPrefix16]
	GridList      AddrTo[[]uint8]
}

type MiniMonopolySectorMetaData struct {
	// Properties
	NumberList AddrTo[[]uint8]
	SectorText AddrTo[Hash]
}

type MiniSkyFireInfoDesMetaData struct {
	// Properties
	Date    AddrTo[StrWithPrefix16]
	TextMap AddrTo[Hash]
}

type MiniSkyFireInfoMetaData struct {
	// Properties
	InfoIconShow        bool
	ShowShopIconTime    AddrTo[RemoteTime]
	RewardShowTextMap   AddrTo[Hash]
	InfoDesList         AddrTo[[]int32]
	TicketIconShow      int32
	RewardShowTpye      int32
	GenActivityRewardID int32
}

type MiniSkyFireStageMetaData struct {
	// Properties
	Index        int32
	HardType     int32
	EnterTimes   int32
	MinScore     int32
	MaxScore     int32
	CoverTextmap AddrTo[Hash]
}

type MirageAvatarMetaData struct {
	// Properties
	SpecialReviseList  AddrTo[[]int32]
	ExclusiveSkillID   AddrTo[[]int32]
	CoreSkillList      AddrTo[[]int32]
	CoreSkillSelectNum int32
	IsRecommend        AddrTo[[]int32]
	NotRecommend       AddrTo[[]int32]
}

type MirageBuffMetaData struct {
	// Properties
	BuffQuality        uint8
	AbilityName        AddrTo[StrWithPrefix16]
	ParamList          AddrTo[[]StrWithPrefix16]
	BuffIcon           AddrTo[StrWithPrefix16]
	ShowBuffIconEffect bool
	BuffName           AddrTo[StrWithPrefix16]
	BuffName_Short     AddrTo[StrWithPrefix16]
	BuffDesc           AddrTo[StrWithPrefix16]
	SimpleBuffDesc     AddrTo[StrWithPrefix16]
	BuffTagList        AddrTo[[]int32]
	BuffSuit           uint8
}

type MirageChallengeLevelMetaData struct {
	// Properties
	HardLevel int32
	BossName  AddrTo[StrWithPrefix16]
	Wave1     AddrTo[[]int32]
	Wave2     AddrTo[[]int32]
	Wave3     AddrTo[[]int32]
	Wave4     AddrTo[[]int32]
	Wave5     AddrTo[[]int32]
}

type MirageMainAvatarMetaData struct {
	// Properties
	AbilityName AddrTo[StrWithPrefix16]
	ParamList   AddrTo[[]StrWithPrefix16]
	SkillName   AddrTo[Hash]
	SkillDesc   AddrTo[Hash]
	SkillIcon   AddrTo[StrWithPrefix16]
}

type MirageMedalMetaData struct {
	// Properties
	MedalRewardType uint8
	GetMedalReward  int32
}

type MirageScheduleMetaData struct {
	// Properties
	BeginTime                       AddrTo[StrWithPrefix16]
	EndTime                         AddrTo[StrWithPrefix16]
	MinLevel                        int32
	MaxLevel                        int32
	BattleStageID                   AddrTo[[]int32]
	DailyStageID                    AddrTo[StrWithPrefix16]
	RotateDays                      int32
	DailyDropMaterialID             int32
	RewardLineActivityID            int32
	MailID                          int32
	ActivityMedalID                 int32
	MedalUnlockStageScore           int32
	MedalMaxScore                   int32
	Mission                         AddrTo[[]int32]
	MissionMaxShowNum               int32
	DailyMaterialDropLimitV2        AddrTo[[]StrWithPrefix16]
	DailyDropLimitResetCostHCoinNum int32
	DailyDropLimitResetTimesLimit   int32
	DailyDropLimitResetOpenTime     AddrTo[StrWithPrefix16]
}

type MirageSkillMetaData struct {
	// Properties
	CommonSkillList  AddrTo[[]int32]
	AmplifySkillList AddrTo[[]int32]
	CoreSkillType    uint8
}

type MirageStageDifficultyMetaData struct {
	// Properties
	StageDifficultyDisplay        AddrTo[StrWithPrefix16]
	StageDifficultyDescript       AddrTo[StrWithPrefix16]
	DifficultyRewardRangeDescript AddrTo[StrWithPrefix16]
}

type MirageStageMetaData struct {
	// Properties
	StageType              uint8
	ThemeID                int32
	StageDifficultyList    AddrTo[[]int32]
	UnlockFlagList         AddrTo[StrWithPrefix16]
	TalentID               int32
	CanUseAvatarList       AddrTo[[]int32]
	CanUseSampleAvatarList AddrTo[[]int32]
	CanUseHelpAvatarList   AddrTo[[]int32]
	MonsterPicture         AddrTo[StrWithPrefix16]
	MonsterName            AddrTo[StrWithPrefix16]
	MonsterModel           AddrTo[StrWithPrefix16]
	MonsterModelScale      float32
	MonsterModelPosition   AddrTo[[]float32]
	MonsterModelRotation   AddrTo[[]float32]
	NewBieGuideStage       int32
	NewBieGuideStageText   AddrTo[StrWithPrefix16]
	BossInfoBtnText        AddrTo[StrWithPrefix16]
	BossInfoTitleText      AddrTo[StrWithPrefix16]
	BossInfoContentText    AddrTo[StrWithPrefix16]
}

type MirageStageScoreMetaData struct {
	// Properties
	CanGetMaterialNum int32
}

type MirageSupportAvatarMetaData struct {
	// Properties
	SkillFunctionName AddrTo[StrWithPrefix16]
	AbilityName       AddrTo[StrWithPrefix16]
	ParamList         AddrTo[[]StrWithPrefix16]
	CD                uint32
	SkillName         AddrTo[Hash]
	SkillDesc         AddrTo[Hash]
	SkillIcon         AddrTo[StrWithPrefix16]
}

type MirageTalentLevelMetaData struct {
	// Properties
	BriefDescription        AddrTo[StrWithPrefix16]
	Iconlist                AddrTo[[]StrWithPrefix16]
	DetailedDescriptionList AddrTo[[]StrWithPrefix16]
	AvatarIDList            AddrTo[[]int32]
	SkillIDList             AddrTo[[]int32]
	MaterialCostNum         int32
}

type MirageTalentMetaData struct {
	// Properties
	EffectName     AddrTo[Hash]
	EffectTxt      AddrTo[StrWithPrefix16]
	EffectIcon     AddrTo[StrWithPrefix16]
	EffectForRange AddrTo[StrWithPrefix16]
	UnlockTime     AddrTo[RemoteTime]
	CostMaterialID int32
	MaxLevel       int32
}

type MirageThemeMetaData struct {
	// Properties
	ThemeType       uint8
	EntrancePicture AddrTo[StrWithPrefix16]
	EntranceName    AddrTo[StrWithPrefix16]
	AwardRange      AddrTo[StrWithPrefix16]
	SceneModel      AddrTo[StrWithPrefix16]
}

type MissionCategoryMetaData struct {
	// Properties
	CategoryTitle AddrTo[StrWithPrefix16]
	CategoryDesc  AddrTo[StrWithPrefix16]
	MissionList   AddrTo[[]int32]
}

type MissionData struct {
	// Properties
	Type          uint8
	SubType       uint16
	Title         AddrTo[Hash]
	Description   AddrTo[Hash]
	Thumb         AddrTo[StrWithPrefix16]
	FinishWay     uint16
	FinishParaInt int32
	FinishParaStr AddrTo[StrWithPrefix16]
	TotalProgress int32
	RewardId      int32
	LinkType      uint16
	LinkParams    AddrTo[[]int32]
	TextmapTag    AddrTo[Hash]
	LinkParamStr  AddrTo[StrWithPrefix16]
	PreviewTime   int32
	NoDisplay     bool
	IsNeedDisplay bool
	Priority      int32
}

type MissionExtraRewardMetaData struct {
	// Properties
	ExtraRewardID int32
}

type MissionGroupMetaData struct {
	// Properties
	MissionIDList   AddrTo[[]int32]
	RPGAreaList     AddrTo[[]int32]
	Reward          int32
	AreaDisplayText AddrTo[Hash]
	NeedMissionNum  int32
}

type MissionPanelConversationMetaData struct {
	// Properties
	Popup_x                 int16
	Popup_y                 int16
	ConversationType        uint16
	ConversationTypeParaInt AddrTo[[]int32]
	ConversationTextMapList AddrTo[[]Hash]
	CharacterImagePathList  AddrTo[[]StrWithPrefix16]
}

type MissionThemeMetaData struct {
	// Properties
	EndTime                 AddrTo[RemoteTime]
	ThemeTitle              AddrTo[Hash]
	BGPath                  AddrTo[StrWithPrefix16]
	AvatarID                int32
	DressID                 int32
	CameraPos               AddrTo[[]float32]
	ImagePath               AddrTo[StrWithPrefix16]
	IconPath                AddrTo[StrWithPrefix16]
	LinkType                int32
	LinkParams              AddrTo[[]int32]
	LinkParamStr            AddrTo[StrWithPrefix16]
	MissionIDList           AddrTo[[]int32]
	UpgradeCostHCoin        int32
	UpgradeCostMCoin        int32
	UpgradeCostMaterialList AddrTo[[]BeastStageDisplayMetaData_LevelDropDisplay]
	UpgradeCostProductName  AddrTo[StrWithPrefix16]
}

type MonopolyBuffInfoMetaData struct {
	// Properties
	SpecialBuffText AddrTo[StrWithPrefix16]
}

type MonopolyBuffMetaData struct {
	// Properties
	BuffName     AddrTo[StrWithPrefix16]
	BuffDesc     AddrTo[StrWithPrefix16]
	BuffIcon     AddrTo[StrWithPrefix16]
	Para         AddrTo[StrWithPrefix16]
	PetJoinText  AddrTo[StrWithPrefix16]
	PetLeaveText AddrTo[StrWithPrefix16]
}

type MonopolyDiceMetaData struct {
	// Properties
	DiceItemID   int32
	DiceMaxValue int32
}

type MonopolyGuessMetaData struct {
	// Properties
	RewardID      int32
	DrawRewardID  int32
	LoseRewardID  int32
	ChatTextLeft  AddrTo[[]Hash]
	ChatTextRight AddrTo[[]Hash]
}

type MonopolyLevelInfoMetaData struct {
	// Properties
	SpecialText1 uint16
	HPValue      int32
	SpecialText2 uint16
	ATKValue     int32
}

type MonopolyLotteryMetaData struct {
	// Properties
	LotteryIcon  AddrTo[StrWithPrefix16]
	LotteryTime  int32
	RewardID     int32
	LoseRewardID int32
}

type MonsterCardActivityDataMetaData struct {
	// Properties
	SweepCostMaterialID   int32
	DailyRewardLimitTimes int32
	CardCollectionMission AddrTo[[]int32]
	InitCardCostLimit     int32
	MaxCardCostLimit      int32
	ExtraCardCostMaterial int32
	CardNumLimit          uint8
	SecretAreaID          uint32
	TowerAreaID           uint32
	ActivityMedalID       int32
	PlayerMonsterCardList AddrTo[[]uint32]
}

type MonsterCardBuffDisplayMetaData struct {
	// Properties
	BuffIconPath AddrTo[StrWithPrefix16]
}

type MonsterCardLevelLimitMetaData struct {
	// Properties
	LimitType     uint8
	Param         AddrTo[StrWithPrefix16]
	LimitTextHint AddrTo[Hash]
}

type MonsterCardLevelMetaData struct {
	// Properties
	LevelUpCost         AddrTo[MonsterCardLevelMetaData_Cost]
	Atk                 float32
	Def                 float32
	HP                  float32
	ResistElementAttack float32
	SkillList           AddrTo[[]uint32]
}

type MonsterCardLevelMetaData_Cost struct {
	// Fields
	Amount     int32
	MaterialID int32
}

type MonsterCardMetaData struct {
	// Properties
	CardPicPath     AddrTo[StrWithPrefix16]
	CardLongPicPath AddrTo[StrWithPrefix16]
	PrefabPath      AddrTo[StrWithPrefix16]
	Scale           float32
	Position        AddrTo[[]float32]
	Rotation        AddrTo[[]float32]
	CardCost        int32
	MonsterSize     uint8
	CardSkill       AddrTo[[]uint32]
	InitThreat      float32
	AttackRange     uint8
	Profession      uint8
	ProfessionName  AddrTo[Hash]
	ProfessionIcon  AddrTo[StrWithPrefix16]
	CardInformation AddrTo[Hash]
	Index           int32
}

type MonsterCardSkillMetaData struct {
	// Properties
	SkillType       uint8
	AbilityName     AddrTo[StrWithPrefix16]
	AbilityParaList AddrTo[[]MonsterCardSkillMetaData_SkillParam]
	Icon            AddrTo[StrWithPrefix16]
	IsAddAbility    bool
	Name            AddrTo[Hash]
	Desc            AddrTo[Hash]
	UnlockDesc      AddrTo[Hash]
}

type MonsterCardSkillMetaData_SkillParam struct {
	// Objects
	Name  StrWithPrefix16
	Param float32
}

type MonsterCardStageMetaData struct {
	// Properties
	StageType             uint8
	StageScene            AddrTo[StrWithPrefix16]
	Wave1                 AddrTo[[]MonsterCardStageMetaData_MonsterSlot]
	Wave2                 AddrTo[[]MonsterCardStageMetaData_MonsterSlot]
	Wave3                 AddrTo[[]MonsterCardStageMetaData_MonsterSlot]
	StageTime             float32
	StageEffect           AddrTo[[]MonsterCardStageMetaData_Effect]
	IsSweep               bool
	SweepNeedChallengeNum uint8
	SweepRewardID         int32
	IsNextStage           bool
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
	// Properties
	LevelCanStarUp           uint8
	LevelLimit               uint8
	StarUpCost               uint32
	AtkRatio                 float32
	DefRatio                 float32
	HPRatio                  float32
	ResistElementAttackRatio float32
	SkillList                AddrTo[[]uint32]
}

type MonsterCardTalentBookMetaData struct {
}

type MonsterCardTowerMetaData struct {
	// Properties
	FloorID uint8
	IsHard  bool
	Medal   uint8
}

type MonsterConfigMetaData struct {
	// Properties
	CategoryName AddrTo[StrWithPrefix16]
	SubTypeName  AddrTo[StrWithPrefix16]
	HP           float32
	Attack       float32
	Defense      float32
	Nature       int32
	ConfigFile   AddrTo[StrWithPrefix16]
	ConfigType   AddrTo[StrWithPrefix16]
	AIName       AddrTo[StrWithPrefix16]
	EliteType    int32
	DisplayTitle AddrTo[StrWithPrefix16]
}

type MonsterResistanceMetaData struct {
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
}

type MonsterWikiDataMetaData struct {
	// Properties
	Page            uint16
	MonsterID       int32
	Name            AddrTo[Hash]
	IconPath        AddrTo[StrWithPrefix16]
	UnlockStageList AddrTo[[]int32]
	UnlockDec       AddrTo[Hash]
	UnlockLevel     uint16
}

type MonsterWikiDescConfigMetaData struct {
	// Properties
	DisplayTitle AddrTo[StrWithPrefix16]
	BGPath       AddrTo[StrWithPrefix16]
	IconPath     AddrTo[StrWithPrefix16]
	DisplayDesc  AddrTo[StrWithPrefix16]
	MessageList  AddrTo[[]StrWithPrefix16]
	AuthorList   AddrTo[[]StrWithPrefix16]
}

type MonsterWikiPageMetaData struct {
	// Properties
	PageName AddrTo[Hash]
	Pic      AddrTo[StrWithPrefix16]
	Show     bool
}

type MPBuffMetaData struct {
	// Properties
	BuffType     uint8
	Icon         AddrTo[StrWithPrefix16]
	Name         AddrTo[StrWithPrefix16]
	Desc         AddrTo[StrWithPrefix16]
	DescInDetail AddrTo[StrWithPrefix16]
}

type MPLevelData struct {
	// Properties
	Name  AddrTo[Hash]
	MpExp int32
}

type MPListMetaData struct {
	// Properties
	Title AddrTo[Hash]
}

type MPPveActivityMetaData struct {
	// Properties
	RewardShowID       uint8
	IsRankShow         bool
	IsAllowSelectSkill bool
}

type MPPveActivityStageGroupMetaData struct {
	// Properties
	StageType uint8
}

type MPRaidActivityMetaData struct {
	// Properties
	SpecialReward     int32
	RankID            int32
	ResetSwitch       bool
	ResetTimes        int32
	WeeklyRefreshDay  AddrTo[[]int32]
	RefreshCostType   int32
	RefreshCostBase   int32
	RefreshCostOffset int32
	Difficulty        int32
	UnlockActivityID  int32
	UnlockTeamLv      int32
	UnlockMpLv        int32
	ActRewardShow     int32
	RecTeamLv         AddrTo[[]int32]
	MapBG             AddrTo[StrWithPrefix16]
	BalanceModeType   int32
}

type MPRaidActMetaData struct {
	// Properties
	ActivityID           int32
	LevelList            AddrTo[[]int32]
	ActMission           int32
	LevelPanelPrefabPath AddrTo[StrWithPrefix16]
	ActNameText          AddrTo[Hash]
	ActTitleText         AddrTo[Hash]
}

type MPRaidRankMetaData struct {
	// Properties
	RankBG       AddrTo[StrWithPrefix16]
	RankTab      AddrTo[StrWithPrefix16]
	RankTabTitle AddrTo[Hash]
}

type MPRaidRankRewardMetaData struct {
	// Properties
	RaidRank int32
	RewardID int32
}

type MPRaidScoreRewardMetaData struct {
	// Properties
	RaidScore int32
	RewardID  int32
}

type MPRaidSeriesMetaData struct {
	// Properties
	UiAssetPath AddrTo[StrWithPrefix16]
	CgExtraKey  AddrTo[StrWithPrefix16]
}

type MPSkillsData struct {
	// Properties
	TypeOfSkill   int32
	MpSkillName   AddrTo[Hash]
	MpSkillDesc   AddrTo[Hash]
	UnlockMPLevel int32
	Proportion1   float32
	Proportion2   float32
	Proportion3   float32
	Proportion4   float32
	MaxSkillLevel int32
}

type MPStageBuffMetaData struct {
	// Properties
	MPBuffIDs AddrTo[[]uint32]
}

type MPStageMetaData struct {
	// Properties
	HostExtraStamina           uint8
	HostExtraDisplayDropList   AddrTo[[]int32]
	MemberExtraDisplayDropList AddrTo[[]int32]
	OnlyCostHostTimes          bool
	UnlockMPLevel              uint8
	WildPlayerExitType         uint8
	ExitButtonType             uint8
	StageMissionList           AddrTo[[]int32]
	MatchType                  uint8
	MaxAssistTimes             uint8
	EnterTimeType              uint8
	EnableGrandKeyBuff         bool
	EnableElf                  bool
	ExitSettleType             uint8
	MPTeamSkillType            uint8
	RevivalDuration            int32
	MPReviveTimes              int32
	MemberIdentifyLimit        int32
	MPMode                     uint8
	MinPlayer                  uint8
	MaxPlayer                  uint8
}

type MPStagePlayerLevelDropMetaData struct {
	// Properties
	MaxPlayerLevel             int32
	HostExtraDisplayDropList   AddrTo[[]int32]
	MemberExtraDisplayDropList AddrTo[[]int32]
}

type MPTagDetectMetaData struct {
	// Properties
	TagList    AddrTo[[]int32]
	TagNumList AddrTo[[]int32]
}

type MPTeamSkillMetaData struct {
	// Properties
	MPTeamSkillName AddrTo[Hash]
	MPTeamSkillDesc AddrTo[Hash]
	MPTeamSkillIcon AddrTo[StrWithPrefix16]
}

type MPTrophyMetaData struct {
	// Properties
	TrophyScore AddrTo[[]int32]
}

type MuseumMetaData struct {
	// Properties
	ActivityID      uint8
	ActivityPanelID uint16
	Image           AddrTo[StrWithPrefix16]
	ShopGoodsIdList AddrTo[[]int32]
	ShopId          int32
}

type NatureCounterMetaData struct {
	// Properties
	NatureCounterTitle AddrTo[Hash]
	NatureCounterDesc  AddrTo[Hash]
}

type NetworkErrCodeMetaData struct {
	// Properties
	TextMapID AddrTo[Hash]
}

type NewbieActivityPanelMetaData struct {
	// Properties
	ActivityType    int32
	ActivityPanelID uint16
	GroupID         int32
	TabTitle        AddrTo[Hash]
	SubTabTitle     AddrTo[Hash]
	TitleText       AddrTo[Hash]
	DescText        AddrTo[Hash]
	InfoText        AddrTo[Hash]
	StringParam     AddrTo[[]StrWithPrefix16]
	CurrencyList    AddrTo[[]uint32]
	CurrencyHide    bool
	ShowLevel       int32
	OpenLevel       int32
	SortID          int32
}

type NewbieActivityScheduleMetaData struct {
	// Properties
	StageMissionID              int32
	LoginRewardSchedule         int32
	NewbiePanelSchedule         int32
	GrowUpActivityID            int32
	WorldMapNewbieChallengeHide bool
	CgGroupID                   int32
	NewbieShopTypeList          AddrTo[[]int32]
	NewbieSubMallID             AddrTo[[]int32]
	DevelopRewardLineID         uint16
	NewbieLevelRush             int32
	HighlightMaterialList       AddrTo[[]int32]
	HighlightResourceTypeList   AddrTo[[]uint8]
	NewbieLoginPostIDList       AddrTo[[]uint16]
}

type NewbieActivityStageMissionMetaData struct {
	// Properties
	MissionList   AddrTo[[]int32]
	BannerImgPath AddrTo[StrWithPrefix16]
}

type NewbieAvatarGuideData struct {
	// Properties
	AvatarBriefText      AddrTo[Hash]
	AvatarTraitsText     AddrTo[Hash]
	RadarParams          AddrTo[[]int32]
	GroupAdvsID          AddrTo[[]int32]
	NewbieGuideList      AddrTo[[]int32]
	DefaultRecommandType int32
	GachaVideoLink       AddrTo[StrWithPrefix16]
}

type NewbieBattleBuff struct {
	// Properties
	EvadeEffectCDReduce     float32
	EvadeEffectCDReduceText AddrTo[Hash]
	ExtraEvadePropID        int32
	ExtraEvadePropText      AddrTo[Hash]
	EvadePropParam1         float32
	EvadePropParam2         float32
	EvadePropParam3         float32
	SpBuffPropID            int32
	SpBuffPropText          AddrTo[Hash]
	SpBuffPropParam1        float32
	SpBuffPropParam2        float32
}

type NewbieCumulativeLotteryMetaData struct {
	// Properties
	KeyRewardMission    int32
	KeyRewardImage      AddrTo[StrWithPrefix16]
	NormalMissionIDList AddrTo[[]StrWithPrefix16]
}

type NewbieDevelopRewardPanelMetaData struct {
	// Properties
	OpenLevel         uint8
	SpecialRewardList AddrTo[[]NewbieDevelopRewardPanelMetaData_SpecialReward]
}

type NewbieDevelopRewardPanelMetaData_SpecialReward struct {
	// Fields
	RankID   uint16
	ShowType uint8

	// Objects
	ParamStr  StrWithPrefix16
	ImagePath StrWithPrefix16
	Tag       Hash
}

type NewbieDormQAvatarMetaData struct {
	// Properties
	IsOpen     int32
	Type       int32
	BigImgPath AddrTo[StrWithPrefix16]
}

type NewbieEquipAdvData struct {
	// Properties
	WeaponAdvID     int32
	Type            int32
	Title           AddrTo[Hash]
	CopyName        AddrTo[Hash]
	WeaponAdvText   AddrTo[Hash]
	StigmataAdv1ID  int32
	StigmataAdv2ID  int32
	StigmataAdv3ID  int32
	StigmataAdvText AddrTo[Hash]
}

type NewbieEquipTypeMetaData struct {
	// Properties
	Desc AddrTo[Hash]
}

type NewbieFirstPaymentMetaData struct {
	// Properties
	OpenLevel    uint8
	MissionID    int32
	MaterialID   int32
	RewardIDList AddrTo[[]int32]
}

type NewbieFirstPaymentRewardMetaData struct {
	// Properties
	AvatarCardID int32
	Name         AddrTo[Hash]
	Tag          AddrTo[Hash]
}

type NewbieGlobalBuff struct {
	// Properties
	BuffEffectList AddrTo[[]int32]
}

type NewbieGroupAdvData struct {
	// Properties
	GroupAvatarID      int32
	RecommendStar      int32
	RecommendText      AddrTo[Hash]
	RecommendSkillIcon AddrTo[StrWithPrefix16]
}

type NewbieGrowUpMetaData struct {
	// Properties
	UseOpenLevel        bool
	OpenLevel           uint8
	DesiredLevel        uint8
	IsOld               bool
	RankName            AddrTo[Hash]
	RankMissionGroupID  uint32
	MissionChainDisplay AddrTo[StrWithPrefix16]
	PreDisplayDesc      AddrTo[Hash]
}

type NewbieGuideDialogueData struct {
	// Properties
	Tag          int32
	SubTag       int32
	RemainTime   int32
	DialogueText AddrTo[Hash]
	AiExpression AddrTo[StrWithPrefix16]
}

type NewbieGuideFaqData struct {
	// Properties
	LevelPhase      int32
	UnlockLevel     int32
	TitleText       AddrTo[Hash]
	DescText        AddrTo[Hash]
	Recommend       int32
	LinkType        int32
	LinkParamList   AddrTo[[]int32]
	LinkParamString AddrTo[StrWithPrefix16]
	FaqBGIcon       AddrTo[StrWithPrefix16]
}

type NewbieGuideFaqListData struct {
	// Properties
	TitleText    AddrTo[Hash]
	DescText     AddrTo[Hash]
	RelatedFaqID AddrTo[[]int32]
	FaqBGIcon    AddrTo[StrWithPrefix16]
	Icon         AddrTo[StrWithPrefix16]
}

type NewbieGuideGrowthQuestData struct {
	// Properties
	LevelPhase       int32
	RelatedMissionID int32
	TitleText        AddrTo[Hash]
	DescText         AddrTo[Hash]
}

type NewbieGuideTutorialLevelData struct {
	// Properties
	StageLevelID        int32
	IsNewMechanismEntry bool
}

type NewbieLevelRushMetaData struct {
	// Properties
	OpenLevel         uint8
	PurchaseCostMCoin uint16
	RewardConfigID    int32
	SpecialReward     AddrTo[NewbieLevelRushMetaData_SpecialRewardData]
	Tag               AddrTo[Hash]
}

type NewbieLevelRushMetaData_SpecialRewardData struct {
	// Fields
	IsPurchaseReward bool
	UnlockLevel      uint8
}

type NewbieLevelRushRewardMetaData struct {
	// Properties
	FreeRewardID     int32
	PurchaseRewardID int32
}

type NewbieLoginPostMetaData struct {
	// Properties
	ImgPath             AddrTo[StrWithPrefix16]
	LinkData            AddrTo[StrWithPrefix16]
	BegineConditionList AddrTo[[]StrWithPrefix16]
	EndConditionList    AddrTo[[]StrWithPrefix16]
	Priority            uint8
}

type NewbieRecurrenceLoginRewardData struct {
	// Properties
	LoginDay     int32
	RewardID     int32
	FreeSelectID int32
	IsKeyReward  bool
	Type         int32
}

type NewbieWebsitesMetaData struct {
	// Properties
	IsOpen      int32
	JumpingType int32
	BigImgPath  AddrTo[StrWithPrefix16]
	WebsiteURL  AddrTo[StrWithPrefix16]
}

type NewFeatureGuideData struct {
	// Properties
	Desc             AddrTo[Hash]
	IsActive         bool
	Type             int32
	MinLv            int32
	MaxLv            int32
	Weight           int32
	UrlPath          AddrTo[StrWithPrefix16]
	FeatureLinkType  int32
	LinkParam        AddrTo[[]int32]
	FeatureBeginTime AddrTo[StrWithPrefix16]
	FeatureEndTime   AddrTo[StrWithPrefix16]
}

type NinjaAreaMetaData struct {
	// Properties
	PreArea         uint8
	UIIndex         AddrTo[StrWithPrefix16]
	TextUIIndex     AddrTo[StrWithPrefix16]
	UnlockTime      AddrTo[StrWithPrefix16]
	AreaName        AddrTo[Hash]
	AreaDesc        AddrTo[Hash]
	BossIcon        AddrTo[StrWithPrefix16]
	BossImg         AddrTo[StrWithPrefix16]
	SlotID          uint8
	UICameraID      AddrTo[[]float32]
	RecommendSlotID uint8
	RecommendLv     uint8
}

type NinjaAttrMetaData struct {
	// Properties
	AttrName AddrTo[Hash]
	AttrDesc AddrTo[Hash]
	AttrIcon AddrTo[StrWithPrefix16]
}

type NinjaMissionCGMetaData struct {
	// Properties
	UnlockCGTime AddrTo[StrWithPrefix16]
	EventID      uint16
	SiteID       uint16
}

type NinjaScheduleMetaData struct {
	// Properties
	MinLevel         uint8
	MaxLevel         uint8
	CurrencyID       int32
	AreaList         AddrTo[[]int32]
	MissionList      AddrTo[[]int32]
	CGID             AddrTo[[]int32]
	RewardShowID     uint8
	EntryPerformID   uint16
	EntryEventID     AddrTo[[]int32]
	WebLink          AddrTo[StrWithPrefix16]
	TalkRoomIconShow bool
	DailyLimitShow   bool
}

type NinjaSiteMetaData struct {
	// Properties
	AreaID          uint8
	SiteType        uint8
	UIIndex         AddrTo[StrWithPrefix16]
	LineUIIndex     AddrTo[StrWithPrefix16]
	StageID         int32
	NextStageID     int32
	ChangeText      AddrTo[Hash]
	UnlockSlotLevel uint8
	PreSiteID       uint16
	LockText        AddrTo[Hash]
	PointTitle      AddrTo[Hash]
	PointDesc       AddrTo[Hash]
}

type NinjaSlotLevelMetaData struct {
	// Properties
	StrengthenExp int32
	AttrList      AddrTo[[]int32]
	AttrValue     AddrTo[[]int32]
	SpEffectID    uint8
	SpValue       float32
}

type NinjaSlotMetaData struct {
	// Properties
	CostMaterialID int32
	StrengthenExp  uint8
	OptionalEffect AddrTo[[]int32]
	OptionValue    AddrTo[[]float32]
	SpUnlockLevel  AddrTo[[]int32]
	SlotName       AddrTo[Hash]
	SlotDesc       AddrTo[Hash]
	SlotIcon       AddrTo[StrWithPrefix16]
}

type NinjaSpEffectMetaData struct {
	// Properties
	SpEffectName AddrTo[Hash]
	SpEffectDesc AddrTo[Hash]
	SpEffectIcon AddrTo[StrWithPrefix16]
}

type NPCLevelMetaData struct {
	// Properties
	HPRatio              float32
	ATKRatio             float32
	DEFRatio             float32
	ElementalResistRatio float32
}

type OpenWorldArea struct {
	// Properties
	MapId          int32
	IsStoryOnly    int32
	IsMarkOnly     bool
	AreaList       AddrTo[[]StrWithPrefix16]
	AreaRGB        AddrTo[StrWithPrefix16]
	UnlockStoryId  int32
	NameTextId     AddrTo[Hash]
	AreaIcon       AddrTo[StrWithPrefix16]
	LevelID        int32
	SwitchPoint    AddrTo[StrWithPrefix16]
	UnlockAreaTips AddrTo[Hash]
	UnlockAreaDesc AddrTo[Hash]
}

type OpenWorldBoss030MetaData struct {
	// Properties
	ProgressText  AddrTo[Hash]
	UnlockStoryId int32
	UnlockParts   AddrTo[StrWithPrefix16]
	UnlockTips    AddrTo[Hash]
}

type OpenWorldContentMetaData struct {
	// Properties
	JudgeType          int32
	ContentJudgeWeight int32
	BeginTime          AddrTo[StrWithPrefix16]
	EndTime            AddrTo[StrWithPrefix16]
	ContentTitle       AddrTo[Hash]
	ContentIconPath    AddrTo[StrWithPrefix16]
	ContentRewardShow  int32
	ContentDesc        AddrTo[Hash]
}

type OpenWorldCookData struct {
	// Properties
	UnlockQuestLv     int32
	DisplayCookName   AddrTo[Hash]
	EffectType        int32
	EffectParam       AddrTo[[]int32]
	DisplayStar       int32
	UseType           int32
	CookTab           int32
	DisplayCookDesc   AddrTo[Hash]
	DisplayEffectDesc AddrTo[Hash]
	DailyTimes        int32
	CostMaterialsList AddrTo[[]AffixRecycleMetaData_IntIntPair]
	IconPath          AddrTo[StrWithPrefix16]
	CookBuffIcon      AddrTo[StrWithPrefix16]
	CookBuffEfx       AddrTo[StrWithPrefix16]
}

type OpenWorldCycleActivityMetaData struct {
	// Properties
	MapID             int32
	MinLevel          int32
	MaxLevel          int32
	BeginTime         AddrTo[StrWithPrefix16]
	EndTime           AddrTo[StrWithPrefix16]
	ContentRewardShow int32
}

type OpenworldCycleData struct {
	// Properties
	Level         int32
	Needstory     int32
	Finishreward  int32
	CycleMap      int32
	HardLvKey     AddrTo[StrWithPrefix16]
	CycleName     AddrTo[Hash]
	EntranceScene int32
	EntranceImg1  AddrTo[StrWithPrefix16]
}

type OpenWorldDestinyLevelDiffMetaData struct {
	// Properties
	AbilityName AddrTo[StrWithPrefix16]
	IconBg      AddrTo[StrWithPrefix16]
}

type OpenWorldDestinyLineLinkMetaData struct {
	// Properties
	Type           int32
	EventID        int32
	LineLinkTypeID int32
	OverrideMap    AddrTo[[]StrWithPrefix16]
	HintTitle      AddrTo[Hash]
	HintContent    AddrTo[Hash]
	HintDesc       AddrTo[Hash]
	HintDuration   float32
}

type OpenWorldDynamicHardLvMetaData struct {
	// Properties
	Cycle1HardLv int32
	Cycle2HardLv int32
	Cycle3HardLv int32
	Cycle4HardLv int32
	Map2HardLv   int32
}

type OpenWorldEventActivityMetaData struct {
	// Properties
	BeginTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
	EventList  AddrTo[[]int32]
	TipsBanner AddrTo[StrWithPrefix16]
}

type OpenWorldEventData struct {
	// Properties
	Type             uint8
	Subtype          AddrTo[[]int32]
	JudgeType        uint8
	EventMap         uint16
	EventArea        uint16
	UnlockEvent      int32
	UnlockCycle      uint8
	EventNameTextMap AddrTo[Hash]
	Area             AddrTo[StrWithPrefix16]
	SpawnPointName   AddrTo[StrWithPrefix16]
	PropName         AddrTo[StrWithPrefix16]
	AreaName         AddrTo[StrWithPrefix16]
	EventInitType    uint16
	RefreshType      uint8
	Reward           int32
}

type OpenWorldEventExtraDropMetaData struct {
	// Properties
	BeginTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
	TipsBanner AddrTo[StrWithPrefix16]
}

type OpenWorldExplorePointMetaData struct {
	// Properties
	MapID           int32
	AreaID          int32
	Type            int32
	LocationID      int32
	BornPointName   AddrTo[StrWithPrefix16]
	PointNameTextID AddrTo[Hash]
	PointDescTextID AddrTo[Hash]
	DefaultStatus   int32
	UIPath          AddrTo[[]StrWithPrefix16]
	PrefabPath      AddrTo[[]StrWithPrefix16]
	UnlockLevel     uint8
	Rotation        float32
}

type OpenWorldFunctionMetaData struct {
	// Properties
	UnlockType    int32
	UnlockPara    int32
	UnlockTextmap AddrTo[Hash]
	UnlockTips    AddrTo[Hash]
	FunctionIcon  AddrTo[StrWithPrefix16]
	FunctionName  AddrTo[Hash]
	FunctionDes   AddrTo[Hash]
}

type OpenWorldLocation struct {
	// Properties
	AreaId             int32
	LocationNameTextId AddrTo[Hash]
	LocationDescTextId AddrTo[Hash]
	SpawnPointName     AddrTo[StrWithPrefix16]
	LocationBuffId     int32
	LocationIcon       AddrTo[StrWithPrefix16]
	LocationType       uint8
}

type OpenWorldMap struct {
	// Properties
	MapType              int32
	UnlockLv             int32
	UnlockStoryId        int32
	ShowTime             AddrTo[StrWithPrefix16]
	UnlockTime           AddrTo[StrWithPrefix16]
	QuestSlotNum         int32
	QuestSettleType      int32
	MapNameText          AddrTo[Hash]
	MapContentText       AddrTo[Hash]
	HpRecoverInterval    int32
	HpRecoverPercent     int32
	QuestMapUIManager    AddrTo[StrWithPrefix16]
	SelectDailyQuestPage AddrTo[StrWithPrefix16]
	SettlementPage       AddrTo[StrWithPrefix16]
	ShopOpenWorldPage    AddrTo[StrWithPrefix16]
	ShopTypeList         AddrTo[[]int32]
	MapInfoText          AddrTo[Hash]
	QuestInfoText        AddrTo[Hash]
	MapSelectPath        AddrTo[StrWithPrefix16]
}

type OpenWorldOverViewData struct {
	// Properties
	NameTextID AddrTo[Hash]
	ScoreRatio int32
}

type OpenWorldQuestActivityMetaData struct {
	// Properties
	BeginTime            AddrTo[StrWithPrefix16]
	EndTime              AddrTo[StrWithPrefix16]
	MapId                int32
	QuestList            AddrTo[[]int32]
	MaxQuestInOneRound   int32
	ShowRewardId         int32
	LinkType             int32
	LinkParams           AddrTo[[]int32]
	ActivityTagName      AddrTo[StrWithPrefix16]
	ActivityQuestTagName AddrTo[StrWithPrefix16]
	TipsBanner           AddrTo[StrWithPrefix16]
}

type OpenworldQuestBuffData struct {
	// Properties
	BuffType      int32
	BuffNameStrId AddrTo[StrWithPrefix16]
	BuffDescStrId AddrTo[Hash]
	BuffIcon      AddrTo[StrWithPrefix16]
}

type OpenworldQuestChallengeData struct {
	// Properties
	ProgressType int32
	Score        int32
	ParamList    AddrTo[[]int32]
}

type OpenworldQuestDataCfg struct {
	// Properties
	QuestType       uint8
	QuestSubType    uint16
	QuestMap        int32
	LevelLockNeed   int32
	QuestArea       int32
	QuestNameId     AddrTo[Hash]
	ShowDesc        bool
	QuestDescText   AddrTo[Hash]
	Rarity          uint8
	Difficulty      int32
	RecTeamLv       uint8
	StaminaCost     uint8
	ChallengeList   AddrTo[[]int32]
	RewardId        int32
	RewardStar      int32
	SlotCost        uint8
	IconPath        AddrTo[StrWithPrefix16]
	LuaPara         AddrTo[[]int32]
	LuaFileName     AddrTo[StrWithPrefix16]
	QuestTipsText   AddrTo[Hash]
	LocationID      int32
	Progress        int32
	FinishWay       int32
	FinishParaStr   AddrTo[StrWithPrefix16]
	ActivityScore   int32
	IsActivityQuest bool
	EnemyList       AddrTo[[]int32]
}

type OpenWorldQuestDeleteRule struct {
	// Properties
	CanDelete int32
}

type OpenWorldQuestJudgeMetaData struct {
	// Properties
	NeedScore int32
	JudgeIcon AddrTo[StrWithPrefix16]
}

type OpenworldQuestLevelDataMetaData struct {
	// Properties
	LevelupNeedStar        int32
	Title                  AddrTo[Hash]
	OWQuestLevelEffectMask int32
}

type OpenWorldQuestMapLevelMetaData struct {
	// Properties
	CompleteReward int32
	WeeklyReward   int32
	KeyQuestID     int32
	Title          AddrTo[Hash]
	RecLevel       AddrTo[[]int32]
}

type OpenWorldQuestMonsterPowerData struct {
	// Properties
	MonsterName AddrTo[StrWithPrefix16]
	ConfigType  AddrTo[[]StrWithPrefix16]
	Power       AddrTo[[]float32]
	LimitCount  int32
}

type OpenWorldQuestRarityMetaData struct {
	// Properties
	WeightGroup AddrTo[[]AffixRecycleMetaData_IntIntPair]
}

type OpenWorldQuestRarityRewardMetaData struct {
	// Properties
	Rarity      int32
	RewardID    int32
	StaminaCost int32
}

type OpenWorldQuestSlotData struct {
	// Properties
	StarReward    int32
	UnlockStoryID int32
}

type OpenWorldQuestTheme struct {
	// Properties
	QuestThemeBuffList AddrTo[[]int32]
	ThemeNameText      AddrTo[Hash]
	ThemeDescText      AddrTo[Hash]
	ThemeIconPath      AddrTo[StrWithPrefix16]
}

type OpenWorldQuestThemeSchedule struct {
	// Properties
	ThemeID int32
	MapID   int32
}

type OpenWorldRewardUpListMetaData struct {
	// Properties
	UpReward int32
}

type OpenWorldSection struct {
	// Properties
	AreaId               int32
	IsOutSide            int32
	SpawnPointAssetPath  AddrTo[StrWithPrefix16]
	MapPrefabPath        AddrTo[StrWithPrefix16]
	UiAssetPath          AddrTo[StrWithPrefix16]
	MapSize              AddrTo[[]float32]
	CoordinateOffsetList AddrTo[[]float32]
	CoordinateList       AddrTo[[]float32]
	SitePanelRotation    float32
	RotationOffset       float32
	RotationSign         float32
	AxisMapping          AddrTo[[]uint8]
}

type OpenWorldStoryData struct {
	// Properties
	StorySeriesID         int32
	StorySeriesStep       int32
	StorySeriesTitle      AddrTo[Hash]
	PreStory              AddrTo[[]int32]
	StoryType             int32
	GroupType             int32
	Cycle                 int32
	StoryStartDate        AddrTo[StrWithPrefix16]
	UnlockMapLv           int32
	UnlockQuestLv         int32
	ShowConditionList     AddrTo[StrWithPrefix16]
	IsUseNewCondition     bool
	UnlockConditionList   AddrTo[StrWithPrefix16]
	PreviewRelateSeries   AddrTo[[]uint32]
	UnlockConditionTips   AddrTo[Hash]
	StoryMap              int32
	StoryArea             int32
	Name                  AddrTo[Hash]
	HuntRewardItem        int32
	HuntRewardItemDisplay int32
	Description           AddrTo[Hash]
	Target                AddrTo[Hash]
	MaxCount              int32
	LocationID            int32
	DLCChallengeMode      bool
	IsTaskAnimation       int32
	IsHideUI              bool
	IsTutorial            bool
	PreStage              AddrTo[[]int32]
}

type OpenWorldStoryDetailMetaData struct {
	// Properties
	GroupIcon        AddrTo[StrWithPrefix16]
	InlevelGroupIcon AddrTo[StrWithPrefix16]
	GroupTitle       AddrTo[Hash]
}

type OpenWorldTaskConfigMetaData struct {
	// Properties
	AutoChaseType int32
	BgPath        AddrTo[StrWithPrefix16]
	GroupTypeList AddrTo[[]int32]
}

type OpenWorldTimerMetaData struct {
	// Properties
	IntervalSecond int32
}

type OptionalBuffListMetaData struct {
	// Properties
	BuffType      int32
	RestrictID    int32
	BuffName      AddrTo[Hash]
	BuffDesc      AddrTo[Hash]
	BuffDescBrief AddrTo[Hash]
	BuffIcon      AddrTo[StrWithPrefix16]
	Score         int32
	GroupType     int32
	Level         int32
}

type OptionalBuffPoolMetaData struct {
	// Properties
	PoolName        AddrTo[Hash]
	BuffList        AddrTo[[]int32]
	UnlockScore     AddrTo[[]int32]
	UnlockCountDown AddrTo[[]int32]
	ForceSelect     bool
}

type OptionalGachaDetailDataMetaData struct {
	// Properties
	BonusFragmentEventID int32
}

type OrderClientMetaData struct {
	// Properties
	Name AddrTo[Hash]
}

type OrderItemMetaData struct {
	// Properties
	ItemId       int32
	ItemNum      int32
	OrderCredits int32
}

type OverlappingCatMetaData struct {
	// Properties
	Group       int32
	ImgPath     AddrTo[StrWithPrefix16]
	TextmapName AddrTo[StrWithPrefix16]
}

type OverlappingMetaData struct {
	// Properties
	SummonItemID  int32
	SummonItemNum int32
	CatIDList     AddrTo[[]int32]
	TextMapRules  AddrTo[StrWithPrefix16]
}

type OverlappingRewardMetaData struct {
	// Properties
	CatID         int32
	OverlappingID int32
	RewardID      int32
}

type OWActivityBossMetaData struct {
	// Properties
	MonsterIdList     AddrTo[[]int32]
	Hardlevel         int32
	MonsterHp         int32
	ClueEventIdList   AddrTo[[]int32]
	ClueLocationList  AddrTo[[]int32]
	ClueTimeLimitTips AddrTo[Hash]
	StaminaCost       int32
	ClueProgress      int32
	ClueTimeLimit     int32
	ClueBuffID        int32
	MutationBuffID    int32
	ActivityExp       int32
	BossReward        int32
	BossName          AddrTo[Hash]
	BossDesc          AddrTo[Hash]
	BossLocationList  AddrTo[[]int32]
	ImagePath         AddrTo[StrWithPrefix16]
	StageID           int32
	MPDamageRatio     float32
	StageMonsterID    int32
}

type OWActivityBossRatingMetaData struct {
	// Properties
	TimeLimit          int32
	RatingReward       int32
	RatingName         AddrTo[Hash]
	RewardBoxPath      AddrTo[StrWithPrefix16]
	RewardBoxAnimation AddrTo[StrWithPrefix16]
}

type OWActivityClueMetaData struct {
	// Properties
	ClueProgressReward int32
	ActivityExpReward  int32
}

type OWActivityLevelMetaData struct {
	// Properties
	NeedExp       int32
	LevelReward   int32
	LevelNameText AddrTo[Hash]
	LevelDescText AddrTo[Hash]
	LevelIconPath AddrTo[StrWithPrefix16]
}

type OWActivityPhaseMetaData struct {
	// Properties
	BeginTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
	ParaString AddrTo[[]int32]
}

type OWActivityScheduleMetaData struct {
	// Properties
	BeginTime       AddrTo[StrWithPrefix16]
	EndTime         AddrTo[StrWithPrefix16]
	PhaseIDList     AddrTo[[]int32]
	BigImgPath      AddrTo[StrWithPrefix16]
	RatingIDList    AddrTo[[]int32]
	ThemeAvatarList AddrTo[[]int32]
	ThemeBuffPara   int32
	GroupName       AddrTo[Hash]
	GroupIcon       AddrTo[StrWithPrefix16]
}

type OWAvatarActivityDataMetaData struct {
	// Properties
	MapID                   uint16
	VAGroup                 uint16
	RestrictAvatarID        uint16
	BanSelfAvatarIDList     AddrTo[[]int32]
	UIManager               AddrTo[StrWithPrefix16]
	TutorialActivityID      int32
	ChallengeActivityID     int32
	QuestDailyLimit         uint8
	QuestRefreshLimit       uint8
	QuestCostStamina        uint8
	QuestTotalLimit         uint8
	LevelMaterialID         int32
	CurrencyLimit           AddrTo[[]OWAvatarActivityDataMetaData_CurrencyLimitItem]
	GenActivityRewardShowID uint8
	CultivateImage          AddrTo[StrWithPrefix16]
	ShowMaterialList        AddrTo[[]int32]
	ShopLink                AddrTo[[]int32]
	WebLink                 AddrTo[StrWithPrefix16]
	GachaLink               AddrTo[StrWithPrefix16]
	GachaID                 int32
	ExploreAreaID           uint16
	ImageInfo               AddrTo[[]GlobalWarScheduleMetaData_ConfigTutorial]
	JsonCfg                 AddrTo[StrWithPrefix16]
	PerformID               int16
	MissionList             AddrTo[[]int32]
}

type OWAvatarActivityDataMetaData_CurrencyLimitItem struct {
	// Fields
	CurrencyID int32
	Max        int32
}

type OWAvatarActivityFilesData struct {
	// Properties
	ActivityID   uint32
	FileTitle    AddrTo[Hash]
	FileText     AddrTo[Hash]
	FileLockText AddrTo[Hash]
	FileLockHint AddrTo[Hash]
}

type OWAvatarActivityLevelMetaData struct {
	// Properties
	NeedNum              int32
	ClientType           uint8
	UnlockAvatarShowList AddrTo[[]uint16]
	UnlockTalentShowList AddrTo[[]OWAvatarActivityLevelMetaData_TalentKey]
}

type OWAvatarActivityLevelMetaData_TalentKey struct {
	// Fields
	TalentID    int32
	TalentLevel int32
}

type OWAvatarActivityQuestMetaData struct {
	// Properties
	ActivityID uint16
	UnlockHint AddrTo[Hash]
}

type OWAvatarActivityScheduleMetaData struct {
	// Properties
	ActivityID int32
	BeginTime  AddrTo[RemoteTime]
	EndTime    AddrTo[RemoteTime]
}

type OWAvatarCultivateDataMetaData struct {
	// Properties
	Type       uint8
	MaxLevel   uint8
	ActivityID uint16
}

type OWAvatarCultivateLevelMetaData struct {
	// Properties
	ConditionList     AddrTo[StrWithPrefix16]
	CostContent       AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	AvatarImage       AddrTo[StrWithPrefix16]
	StigmataImage     AddrTo[StrWithPrefix16]
	WeaponImage       AddrTo[StrWithPrefix16]
	AvatarDetailImage AddrTo[StrWithPrefix16]
	WeaponDetailImage AddrTo[StrWithPrefix16]
	VirtualAvatarID   int32
	VAStigmataIDList  AddrTo[[]int32]
	VAWeaponID        int32
}

type OWAvatarUnlockMetaData struct {
	// Properties
	AvatarIconPath AddrTo[StrWithPrefix16]
	Desc           AddrTo[StrWithPrefix16]
	ConditionList  AddrTo[StrWithPrefix16]
}

type OWEndlessAreaDetailMetaData struct {
	// Properties
	AreaNameTextId AddrTo[StrWithPrefix16]
	AreaDescTextId AddrTo[Hash]
	AreaIcon       AddrTo[StrWithPrefix16]
	AreaBuffId     int32
}

type OWEndlessAreaScoreConfig struct {
	// Properties
	MonsterScoreConfig AddrTo[[]StrWithPrefix16]
	BossScoreConfig    AddrTo[[]int32]
	Type               uint8
	// SplitBattles AddrTo[[]OWEndlessAreaScoreConfig_MonsterGroup]

}

type OWEndlessAreaScoreConfig_MonsterGroup struct {
	// Fields
	GroupScoreId int32
	IsBoss       bool
	LayerId      int32
}

type OWEndlessBattleAreaMetaData struct {
	// Properties
	BattleArea     AddrTo[StrWithPrefix16]
	AreaId         int32
	TrapMaxNum     int32
	BattleAreaName AddrTo[Hash]
}

type OWEndlessBattleConfig struct {
	// Properties
	EnvironmentIDList AddrTo[[]int32]
	AreaIDList        AddrTo[[]int32]
	BossAreaIDList    AddrTo[[]int32]
	MapID             int32
	AreaScoreConfigID int32
}

type OWEndlessBossData struct {
	// Properties
	MonsterTypeID     int32
	GroupLevelRange   AddrTo[[]int32]
	BossLevel         int32
	UniqueIDList      AddrTo[[]int32]
	LuaFilePath       AddrTo[StrWithPrefix16]
	BOSSDisplayIDList AddrTo[[]int32]
}

type OWEndlessBuff struct {
	// Properties
	BuffName        AddrTo[StrWithPrefix16]
	BuffType        int32
	Effect          AddrTo[StrWithPrefix16]
	InLevelIconPath AddrTo[StrWithPrefix16]
}

type OWEndlessEnvironmentMeta struct {
	// Properties
	MonsterTypeID       int32
	BuffID              int32
	ExtraBuffID         int32
	EnvironmentNameText AddrTo[Hash]
	EnvironmentDescText AddrTo[Hash]
	EnvironmentIconPath AddrTo[StrWithPrefix16]
	UpTagList           AddrTo[[]ElfMetaData_Tag]
	DownTagList         AddrTo[[]ElfMetaData_Tag]
}

type OWEndlessGroupMeta struct {
	// Properties
	PlayerGroup        int32
	GroupName          AddrTo[Hash]
	GroupHardLevel     int32
	EliteRatio         float32
	ExtraScore         int32
	ExtraScoreTime     int32
	MaxCapacity        int32
	PromoteRank        int32
	DemoteRank         int32
	PromoteDesc        AddrTo[Hash]
	NormalDesc         AddrTo[Hash]
	DemoteDesc         AddrTo[Hash]
	UnSelectColor      AddrTo[StrWithPrefix16]
	BgColor            AddrTo[StrWithPrefix16]
	UnselectPrefabPath AddrTo[StrWithPrefix16]
	SelectPrefabPath   AddrTo[StrWithPrefix16]
	UnopenPrefabPath   AddrTo[StrWithPrefix16]
	MinSkipLevel       uint8
	MaxSkipLevel       uint8
	SkipLevelTime      uint16
}

type OWEndlessInvadeBuffMetaData struct {
	// Properties
	BuffCondition    int32
	ConditionVarInt  int32
	ConditionVarList AddrTo[[]int32]
	BuffName         AddrTo[StrWithPrefix16]
	BuffDesc         AddrTo[Hash]
	BuffIconPath     AddrTo[StrWithPrefix16]
	ConditionDesc    AddrTo[Hash]
}

type OWEndlessInvadeMetaData struct {
	// Properties
	BossConfigIDList AddrTo[[]OWEndlessInvadeMetaData_IDPointData]
	InvadeScore      int32
	BuffList         AddrTo[[]OWEndlessInvadeMetaData_IDPointData]
}

type OWEndlessInvadeMetaData_IDPointData struct {
	// Fields
	ID    int32
	Point int32
}

type OWEndlessItemEffectMetaData struct {
	// Properties
	LimitType    int32
	BagAmount    int32
	EffectAmount int32
}

type OWEndlessItemMetaData struct {
	// Properties
	Name             AddrTo[Hash]
	PropName         AddrTo[StrWithPrefix16]
	IndicatorOffset  float32
	Rarity           int32
	Type             int32
	MaxUseRange      int32
	MinUseRange      int32
	ParamScore       int32
	RelateBuff       int32
	ParamStr         AddrTo[StrWithPrefix16]
	IconPath         AddrTo[StrWithPrefix16]
	SmallIconPath    AddrTo[StrWithPrefix16]
	DescriptionShort AddrTo[Hash]
	Description      AddrTo[Hash]
	BuffDesc         AddrTo[Hash]
	DebuffDesc       AddrTo[Hash]
	ShowHint         AddrTo[StrWithPrefix16]
}

type OWEndlessMonsterData struct {
	// Properties
	MonsterTypeID   uint8
	InitTypeList    AddrTo[[]uint16]
	GroupLevelRange AddrTo[[]uint8]
	Name            AddrTo[StrWithPrefix16]
	ConfigType      AddrTo[StrWithPrefix16]
	UniqueID        uint16
	Nature          uint8
	Score           uint16
	Weight          uint8
	BossRank        uint8
}

type OWEndlessMonsterGroupScore struct {
	// Properties
	Score                 uint16
	MonsterInitType       int32
	ExtraInitTypes        AddrTo[[]StrWithPrefix16]
	RandomType            uint8
	DifficultyDescription AddrTo[Hash]
	// ExtraInitTypesProcessed AddrTo[[]OWEndlessMonsterGroupScore_InitEntry]

}

type OWEndlessMonsterGroupScore_InitEntry struct {
	// Fields
	InitType  int32
	Weight    int32
	WeightSum int32
}

type OWEndlessMonsterInitTypeMetaData struct {
	// Properties
	MonsterScoreRatio float32
	InitTypeDesc      AddrTo[Hash]
	Type              int32
	WaveList          AddrTo[[]int32]
	IsNegate          bool
	EnvironmentIDList AddrTo[[]int32]
}

type OWEndlessPlayerBaseRewardMetaData struct {
	// Properties
	RewardID int32
}

type OWEndlessPlayerGroupMeta struct {
	// Properties
	MinPlayerLevel int32
	MaxPlayerLevel int32
	Icon           AddrTo[StrWithPrefix16]
	Name           AddrTo[Hash]
}

type OWEndlessRewardConfigMetaData struct {
	// Properties
	PrototeRewardID uint32
	NormalRewardID  uint32
	DemoteRewardID  uint32
}

type OWEndlessScheduleMetaData struct {
	// Properties
	RewardConfig uint32
}

type OWEndlessSingleFloorDataMetaData struct {
	// Properties
	Score                 int32
	ShowReset             bool
	ResetFloor            uint8
	MonsterGroupID        int32
	HardLevel             uint16
	ResetCostStamina      uint8
	ResetCostMaterialList AddrTo[[]AffixRecycleMetaData_IntIntPair]
}

type OWEndlessSingleModeStageMetaData struct {
	// Properties
	StageID       int32
	Type          int32
	EnvironmentID AddrTo[[]StrWithPrefix16]
	MapID         int32
	AreaIDList    AddrTo[[]int32]
	MaxScore      uint16
	// EnvironmentIDArrayProcessed AddrTo[[]OWEndlessSingleModeStageMetaData_EnvironmentEntry]

}

type OWEndlessSingleModeStageMetaData_EnvironmentEntry struct {
	// Fields
	EnvironmentID int32
	Weight        int32
	WeightSum     int32
}

type OWEndlessSingleMonsterGroupMetaData struct {
	// Properties
	BattleArea     AddrTo[StrWithPrefix16]
	AreaID         int32
	MonsterGroupID int32
	IsBoss         bool
	MaxScore       int32
}

type OWEventDisplayMetaData struct {
	// Properties
	Type          uint8
	Level         int32
	EventName     AddrTo[StrWithPrefix16]
	EventLockHint AddrTo[Hash]
	EventImage    AddrTo[StrWithPrefix16]
}

type OWHuntActivityDataMetaData struct {
	// Properties
	MapIDList            AddrTo[[]int32]
	MissionIdList        AddrTo[[]int32]
	BGPicPath            AddrTo[StrWithPrefix16]
	TalentBGPicPath      AddrTo[StrWithPrefix16]
	CollectionsBGPicPath AddrTo[StrWithPrefix16]
	RewardScheduleID     int32
}

type OWHuntActivityFloorMetaData struct {
	// Properties
	MapID         int32
	Floor         uint8
	CurFloor      uint8
	NextFloor     uint8
	CurFloorArea  AddrTo[StrWithPrefix16]
	NextFloorArea AddrTo[StrWithPrefix16]
}

type OWHuntActivityHunterMetaData struct {
	// Properties
	HunterPicPath  AddrTo[StrWithPrefix16]
	HunterName     AddrTo[Hash]
	HunterDes1     AddrTo[Hash]
	HunterDes2     AddrTo[Hash]
	WeaknessIDList AddrTo[[]uint16]
}

type OWHuntActivityMachineEventMetaData struct {
	// Properties
	EventID int32
	MapID   int32
}

type OWHuntActivityMachineMetaData struct {
	// Properties
	MachineIcon AddrTo[StrWithPrefix16]
	MachineName AddrTo[Hash]
	MachineDes  AddrTo[Hash]
}

type OWHuntActivityMapDataMetaData struct {
	// Properties
	PreMapId           int32
	SectionID          int32
	FloorsNum          uint8
	OpenWorldName1     AddrTo[Hash]
	OpenWorldName2     AddrTo[Hash]
	OpenWorldPicPath   AddrTo[StrWithPrefix16]
	OpenPremise        uint8
	StrongholdMaxTimes uint8
	QuestDailyTimes    uint8
	QuestMaxTimes      uint8
	QuestIDList        AddrTo[[]int32]
	ForceAreaEventID   int32
	ForceAreaPicPath   AddrTo[StrWithPrefix16]
	ForceAreaFuncParam AddrTo[[]StrWithPrefix16]
	ForceAreaFloor     uint8
	BornPointFloor     uint8
	HunterID           int32
	HunterEventID      int32
	MachineShowIDList  AddrTo[[]int32]
	InitialTalentList  AddrTo[[]ChapterOWTalentLevelMetaData_PreTalentRequirement]
}

type OWHuntActivityMonsterMetaData struct {
	// Properties
	MonsterWeather      AddrTo[Hash]
	MonsterDes          AddrTo[Hash]
	MonsterDetailIDList AddrTo[[]int32]
}

type OWHuntActivityProgressMetaData struct {
	// Properties
	Progress float32
	MapID    int32
	Floor    uint8
}

type OWHuntActivityQuestMetaData struct {
	// Properties
	FuncParam AddrTo[[]StrWithPrefix16]
}

type OWHuntActivitySHLevelMetaData struct {
	// Properties
	HardDes AddrTo[Hash]
}

type OWHuntActivityStrongholdsMetaData struct {
	// Properties
	MapID               int32
	EventID             int32
	BornPointFloor      uint8
	StrongholdName      AddrTo[Hash]
	StrongholdDes       AddrTo[Hash]
	StrongholdLevelType uint8
	FuncParam           AddrTo[[]StrWithPrefix16]
}

type OWHuntActivityTalentMetaData struct {
	// Properties
	TalentIconPath       AddrTo[StrWithPrefix16]
	PreTalent            AddrTo[[]ChapterOWTalentLevelMetaData_PreTalentRequirement]
	MapID                int32
	TalentUpgradeCost    AddrTo[[]OWHuntActivityTalentMetaData_TalentUpgradeCostData]
	TalentUpgradePremise uint8
	TalentName           AddrTo[Hash]
	TalentDes            AddrTo[Hash]
	NextLevelDes         AddrTo[Hash]
	AbilityName          AddrTo[StrWithPrefix16]
	AbilityParams        AddrTo[[]StrWithPrefix16]
}

type OWHuntActivityTalentMetaData_TalentUpgradeCostData struct {
	// Fields
	TalentUpgradeCostMaterial int32
	TalentUpgradeCostNum      int32
}

type OWHuntActivityTeleportMetaData struct {
	// Properties
	MapID          int32
	TextmapName    AddrTo[Hash]
	TextmapDesc    AddrTo[Hash]
	SpawnPointName AddrTo[StrWithPrefix16]
	FuncParam      AddrTo[[]StrWithPrefix16]
}

type OWTalentDataMetaData struct {
	// Properties
	MaxLevel uint8
}

type OWTalentLevelDataMetaData struct {
	// Properties
	IsStoryValid  bool
	ConditionList AddrTo[StrWithPrefix16]
	Name          AddrTo[StrWithPrefix16]
	Desc          AddrTo[StrWithPrefix16]
	IconPath      AddrTo[StrWithPrefix16]
	AbilityName   AddrTo[StrWithPrefix16]
	ParamList     AddrTo[[]StrWithPrefix16]
}

type OWTeleporterMetaData struct {
	// Properties
	SectionID      int32
	TeleporterName AddrTo[Hash]
	TeleporterDesc AddrTo[Hash]
	SpawnPointList AddrTo[[]StrWithPrefix16]
	LockHintText   AddrTo[StrWithPrefix16]
}

type PanelConversationTriggerMetaData struct {
	// Properties
	Popup_x_default int16
	Popup_y_default int16
}

type PayActivityDisplayScheduleMetaData struct {
	// Properties
	BeginTime           AddrTo[StrWithPrefix16]
	EndTime             AddrTo[StrWithPrefix16]
	DisplayTextMap      AddrTo[StrWithPrefix16]
	PopupDisplayTextMap AddrTo[StrWithPrefix16]
	IconImgPath         AddrTo[StrWithPrefix16]
	DisplayImgPathList  AddrTo[[]StrWithPrefix16]
	CGEventID           int32
}

type PayInfoMetaData struct {
	// Properties
	CanPay        bool
	CannotPayHint AddrTo[Hash]
}

type PerformEventMetaData struct {
	// Properties
	EventType uint8
	ParamsVar AddrTo[[]int32]
}

type PhoneEntranceAcountOverrideMetaData struct {
	// Properties
	EntranceIconPath AddrTo[StrWithPrefix16]
	IsOpen           bool
	LinkUrl          AddrTo[StrWithPrefix16]
}

type PhoneEntranceMetaData struct {
	// Properties
	BtnType            int32
	EntranceBeginTime  AddrTo[StrWithPrefix16]
	EntranceEndTime    AddrTo[StrWithPrefix16]
	Weight             int32
	UrlType            int32
	LinkUrl            AddrTo[StrWithPrefix16]
	LinkType           int32
	LinkParams         AddrTo[StrWithPrefix16]
	EntranceText       AddrTo[Hash]
	EntranceIconPath   AddrTo[StrWithPrefix16]
	RedHintType        int32
	RedHintUiType      int32
	RedHintMainPage    int32
	MinLv              int32
	MaxLv              int32
	ChannelSettingType int32
	UIOverrideItemList AddrTo[[]uint32]
}

type PhoneNoticeDataMetaData struct {
	// Properties
	NoticeBeginTime  AddrTo[RemoteTime]
	NoticeEndTime    AddrTo[RemoteTime]
	NoticeAppearCD   int32
	NoticeAppearTime int32
	NoticeVideoID    AddrTo[[]int32]
}

type PicBGListdMeta struct {
	// Properties
	PicBGType int32
}

type PictureStepMetaData struct {
	// Properties
	KeyID                int32
	ButtonPrefab         AddrTo[StrWithPrefix16]
	Reward               int32
	Name_Textmap         AddrTo[Hash]
	ChoiceBefore_Textmap AddrTo[[]Hash]
	ChoiceAfter_Textmap  AddrTo[[]Hash]
	Pic                  AddrTo[StrWithPrefix16]
	NeedScore            int32
}

type PicTutorialMetaData struct {
	// Properties
	Type            uint32
	Title           AddrTo[StrWithPrefix16]
	StepIDList      AddrTo[[]int32]
	UnlockCondition int32
	Weight          uint8
	Param           int32
}

type PicTutorialStepDataMetaData struct {
	// Properties
	Title   AddrTo[StrWithPrefix16]
	Desc    AddrTo[StrWithPrefix16]
	ImgPath AddrTo[StrWithPrefix16]
}

type PlatformShootDailyScheduleMetaData struct {
	// Properties
	EndTime AddrTo[RemoteTime]
	SiteID  int32
}

type PlatformShooterActivityMetaData struct {
	// Properties
	AffixMaterialId          int32
	WeaponPoolMaterialId     int32
	QavatarLevelUpMaterialId int32
	DailyTicketId            int32
	InitialQavatarList       AddrTo[[]PlatformShooterActivityMetaData_InitAvatar]
	QavatarIDList            AddrTo[[]int16]
	WeaponIDList             AddrTo[[]int16]
	MedalID                  int32
	MedalMaxParam            uint8
	WeaponPoolList           AddrTo[[]int16]
	RankStageIDList          AddrTo[[]int32]
	IsShowChallengeRank      bool
	DailyDropIDList          AddrTo[[]uint8]
}

type PlatformShooterActivityMetaData_InitAvatar struct {
	// Fields
	AvatarID int16
	Level    int16
}

type PlatformShooterAffixPoolMetaData struct {
	// Properties
	CostMaterialNum int16
}

type PlatformShooterAvatarTrialMetaData struct {
	// Properties
	AvatarId    int16
	Level       int16
	WeaponLevel int16
	AffixList   AddrTo[[]int32]
}

type PlatformShooterDailyDropMetaData struct {
	// Properties
	BeginTime               AddrTo[RemoteTime]
	PreSite                 int32
	DailyChallengeDropLimit int16
}

type PlatformShooterMonsterMetaData struct {
	// Properties
	ModelPath   AddrTo[StrWithPrefix16]
	OffsetY     float32
	RotationY   float32
	Scale       float32
	Weight      uint8
	MonsterName AddrTo[Hash]
	MonsterDesc AddrTo[Hash]
	Solution    AddrTo[Hash]
	MonsterIcon AddrTo[StrWithPrefix16]
}

type PlatformShooterQavatarMetaData struct {
	// Properties
	LevelUpMaterialNum   int16
	MaxLevel             int16
	AtkBase              float32
	HpBase               float32
	CriticalRate         uint8
	CriticalDamage       int16
	SpBase               int16
	SkillIdList          AddrTo[[]int32]
	NeedPassSiteID       int32
	WeaponID             int16
	ModelPath            AddrTo[StrWithPrefix16]
	AvatarName           AddrTo[Hash]
	AttackSpeed          AddrTo[Hash]
	AttackDistance       AddrTo[Hash]
	AffixDes             AddrTo[Hash]
	IconPath             AddrTo[StrWithPrefix16]
	ImagePath            AddrTo[StrWithPrefix16]
	Fake_Avatar_Registry AddrTo[StrWithPrefix16]
}

type PlatformShooterRoomMetaData struct {
	// Properties
	LuaFile       AddrTo[StrWithPrefix16]
	MonsterIdList AddrTo[[]int16]
}

type PlatformShooterSkillMetaData struct {
	// Properties
	AbilityName AddrTo[StrWithPrefix16]
	SkillType   uint8
	ParamList   AddrTo[[]StrWithPrefix16]
	Icon        AddrTo[StrWithPrefix16]
	UnlockLv    int16
	SkillTitle  AddrTo[Hash]
	SkillDesc   AddrTo[Hash]
	MaxLimit    int16
}

type PlatformShooterStageMetaData struct {
	// Properties
	NextStageId        int32
	StageType          uint8
	AvatarTrialList    AddrTo[[]int16]
	AvatarRestrictList AddrTo[[]int16]
	MaxAvatarNum       uint8
	RelateAvatarID     int16
	RoomList           AddrTo[[]int16]
	RoomNum            uint8
	Param              uint8
}

type PlatformShooterWeaponMetaData struct {
	// Properties
	WeaponTitle  AddrTo[Hash]
	WeaponDesc   AddrTo[Hash]
	Icon         AddrTo[StrWithPrefix16]
	WeaponAtk    float32
	AffixNum     uint8
	AffixPoolMap AddrTo[[]PlatformShooterWeaponMetaData_AffixGachaPool]
}

type PlatformShooterWeaponMetaData_AffixGachaPool struct {
	// Fields
	PoolID     uint8
	RefineType uint8
}

type PlatformShooterWeaponPoolMetaData struct {
	// Properties
	BeginTime       AddrTo[RemoteTime]
	PreSite         int32
	CostMaterialNum int32
	WeaponList      AddrTo[[]int16]
}

type PlayerLevelLockMetaData struct {
	// Properties
	MaxLevel           uint8
	MaxLayUpExp        int32
	MissionList        AddrTo[[]int32]
	EventID            int32
	PlayerLockHintText AddrTo[Hash]
	LinkButton         AddrTo[StrWithPrefix16]
	LevelGiftPackBG    AddrTo[StrWithPrefix16]
	LevelGiftPackFrame AddrTo[StrWithPrefix16]
	TutorialVideoID    int32
}

type PlayerLevelMetaData struct {
	// Properties
	Exp               int32
	Stamina           int32
	NumFriends        int32
	AvatarLevelLimit  int32
	StaminaBonus      int32
	SweepStaminaLimit int32
}

type PlayerLevelShopGoodsMetaData struct {
	// Properties
	GoodsIDList AddrTo[[]int32]
	LockID      int32
}

type PlayerPrivilegeScheduleMetaData struct {
	// Properties
	PrivilegeType   int32
	ScheduleID      int32
	RightConfigList AddrTo[[]int32]
}

type PlotAvatarMetaData struct {
	// Properties
	Name        AddrTo[StrWithPrefix16]
	DisplayType int32
	Path        AddrTo[StrWithPrefix16]
}

type PlotLineCgDataMetaData struct {
	// Properties
	CGID int32
}

type PlotlineConditionDataMetaData struct {
	// Properties
	ConditionType uint32
	Param         AddrTo[StrWithPrefix16]
}

type PlotLineDataMetaData struct {
	// Properties
	PlotlineJsonPath AddrTo[StrWithPrefix16]
}

type PlotMetaData struct {
	// Properties
	LevelID       int32
	StartDialogID int32
	EndDialogID   int32
}

type PonMachineCollectionMetaData struct {
	// Properties
	Icon AddrTo[StrWithPrefix16]
	Name AddrTo[Hash]
	Des  AddrTo[Hash]
}

type PonMachineScheduleMetaData struct {
	// Properties
	PonMachineMaterialID          int32
	PonMachineMaterialCostNumList AddrTo[[]int32]
	PonMachineInfoKey             AddrTo[Hash]
	RewardlineID                  uint32
	PonMachineCollectionIDList    AddrTo[[]uint32]
}

type PowerTypeMetaData struct {
	// Properties
	PowerConf float32
}

type PredownloadAsbMetaData struct {
	// Properties
	EnableTime AddrTo[RemoteTime]
}

type PredownloadAudioPackageMetaData struct {
	// Properties
	Version AddrTo[StrWithPrefix16]
	PckType uint8
}

type PredownloadVideoFileMetaData struct {
	// Properties
	Version  AddrTo[StrWithPrefix16]
	PckType  uint8
	FileSize int32
}

type ProductRecommendMetaData struct {
	// Properties
	ShowWeight              int32
	CancelType              int32
	DialogType              uint8
	TrickEventType          int32
	ShowTagList             AddrTo[[]int32]
	ImgBG                   AddrTo[StrWithPrefix16]
	ShopGoodsScheduleIDList AddrTo[[]uint32]
	TitleText               AddrTo[Hash]
	SubTitleText            AddrTo[Hash]
	RecommendTipsText       AddrTo[Hash]
	ImgBGType               uint8
	BeginTime               AddrTo[RemoteTime]
	EndTime                 AddrTo[RemoteTime]
	TipsLink                AddrTo[StrWithPrefix16]
}

type ProtectTypeMetaData struct {
	// Properties
	ProtectDisplayDesc AddrTo[Hash]
	ProtectInfoTitle   AddrTo[Hash]
	ProtectInfoDesc    AddrTo[Hash]
}

type PVloginMetaData struct {
	// Properties
	IsHideCountDownInfo                     bool
	IsHideGotoLiveBtn                       bool
	JumpLiveSpecifiedTimeBefore             AddrTo[RemoteTime]
	JumpLiveSpecifiedTimeAfter              AddrTo[RemoteTime]
	JumpInGameLink1                         AddrTo[StrWithPrefix16]
	JumpInGameLink2                         AddrTo[StrWithPrefix16]
	JumpInGameLink3                         AddrTo[StrWithPrefix16]
	CL_ActivityPvLoginPanel_LastingTime     AddrTo[Hash]
	CL_ActivityPvLoginPanel_Label_Countdown AddrTo[Hash]
	CL_ActivityPvLoginPanel_Dialog_NotStart AddrTo[Hash]
	CL_ActivityPvLoginPanel_Dialog_HasEnded AddrTo[Hash]
}

type PVZActivityMetaData struct {
	// Properties
	QavatarIDList  AddrTo[[]int32]
	JsonConfigPath AddrTo[StrWithPrefix16]
	MedalID        int32
}

type PVZMonsterMetaData struct {
	// Properties
	LinkedActivityID     int32
	Model                int32
	QmonsterName         AddrTo[Hash]
	QmonsterSkill        AddrTo[Hash]
	Solution             AddrTo[Hash]
	RecommendQavatarList AddrTo[[]int32]
	QmonsterIcon         AddrTo[StrWithPrefix16]
	OffsetY              float32
}

type PVZQavatarMetaData struct {
	// Properties
	QavatarLevelUpMaterialID int32
	NextlevelCost            int32
	SpecialMaterialID        int32
	EntityID                 int32
	QavatarName              AddrTo[Hash]
	QavatarDesc              AddrTo[Hash]
	QavatarCD                float32
	ChibiIcons               AddrTo[StrWithPrefix16]
	CGID                     int32
	QavatarCost              int32
	NextLevelConditionType   uint8
	NextLevelConditionParam  int32
	QavatarATKShow           int32
	QavatarHPShow            int32
	TargetArea               AddrTo[[]PVZQavatarMetaData_CellAreaConfig]
	TargetAreaWithOrnaments  AddrTo[[]PVZQavatarMetaData_CellAreaConfig]
	QavatarType              uint8
}

type PVZQavatarMetaData_CellAreaConfig struct {
	// Fields
	Type  int32
	Param float32
}

type PVZSiteMetaData struct {
	// Properties
	SiteType          int32
	TicketCost        AddrTo[[]MechMetaData_DisjoinOutputItem]
	DropMaterialLimit AddrTo[[]MechMetaData_DisjoinOutputItem]
	NextSiteID        int32
	MaxFloor          int32
	SiteTitle         AddrTo[Hash]
	MaxWaves          int32
}

type PVZTileMetaData struct {
	// Properties
	LimitQAvatarNumber    int32
	InitialQavatarList    AddrTo[[]uint32]
	MonsterShowList       AddrTo[[]int32]
	RecommendQavatarList  AddrTo[[]uint32]
	RecommendQavatarLevel int32
	StageDesc             AddrTo[Hash]
	DropMaterial          AddrTo[[]MechMetaData_DisjoinOutputItem]
}

type QAvatarBattleBroadCastMetaData struct {
	// Properties
	BroadCastDescription AddrTo[Hash]
}

type QAvatarBattleGadgetMetaData struct {
	// Properties
	ContactId                    int32
	ParaInt                      int32
	GadgetHP                     float32
	IsShowHPBar                  bool
	GadgetPrefab                 AddrTo[StrWithPrefix16]
	GadgetIcon                   AddrTo[StrWithPrefix16]
	GadgetIconPos                float32
	GadgetIconScale              float32
	GadgetBornEffectPattern      AddrTo[StrWithPrefix16]
	GadgetDisappearEffectPattern AddrTo[StrWithPrefix16]
}

type QAvatarBattleGadgetRefreshMetaData struct {
	// Properties
	ForcastID   int32
	BroadCastID int32
}

type QAvatarBattleKillingStreakMetaData struct {
	// Properties
	Name          AddrTo[Hash]
	AnimName      AddrTo[StrWithPrefix16]
	AudioPattern1 AddrTo[StrWithPrefix16]
	AudioPattern2 AddrTo[StrWithPrefix16]
}

type QAvatarBattleSpawnZoneMetaData struct {
	// Properties
	SpawnPointTextmap AddrTo[Hash]
}

type QAvatarBattleTextIdMapMetaData struct {
	// Properties
	TextString AddrTo[StrWithPrefix16]
}

type QAvatarMissionMetaData struct {
	// Properties
	ProcessReward   int32
	MissionTypeText AddrTo[Hash]
}

type QAvatarPVPMetaData struct {
	// Properties
	PairAvatarID int32
	WeaponID     int32
	QAvatarName  AddrTo[Hash]
}

type QAvatarPVPWeaponMetaData struct {
	// Properties
	Type                          int32
	WeaponTypeIcon                AddrTo[StrWithPrefix16]
	AimDistance                   float32
	SkillID                       int32
	WeaponAssetPath               AddrTo[StrWithPrefix16]
	WeaponAnimationPartsAssetPath AddrTo[[]StrWithPrefix16]
	WeaponAttachRoot              AddrTo[StrWithPrefix16]
	WeaponAttachPosition          AddrTo[[]float32]
	WeaponAttachRotation          AddrTo[[]float32]
	BulletNum                     int32
	BulletRecoverRate             float32
	ReloadTime                    float32
	BulletWarningNum              int32
}

type QCandyActivityDataMetaData struct {
	// Properties
	UnlockItemID         int32
	CoinMaterialID       int32
	DailyLimit           int32
	ActivityMedalID      int32
	InviteStageID        int32
	InLevelTrialAvatarID int32
	MedalUnlockRank      int32
	MedalHighestRank     int32
	MapPoolList          AddrTo[[]int32]
	PVPCloseTime         AddrTo[RemoteTime]
	PVPRequireSite       int32
	DefaultAvatar        int32
	TitlePath            AddrTo[StrWithPrefix16]
	MainPagePath         AddrTo[StrWithPrefix16]
}

type QCandyAvatarMetaData struct {
	// Properties
	AvatarRegistryKey  AddrTo[StrWithPrefix16]
	UIAvatarPrefabPath AddrTo[StrWithPrefix16]
	UIControllerPath   AddrTo[StrWithPrefix16]
	InitFace           bool
	FullName           AddrTo[Hash]
	ChibiIconPath      AddrTo[StrWithPrefix16]
	ItemCost           int32
	RankRequest        int32
	IsShownLocked      bool
}

type QCandyEmojiMetaData struct {
	// Properties
	IconPath      AddrTo[StrWithPrefix16]
	EmojiPath     AddrTo[StrWithPrefix16]
	EmojiDisplay  uint8
	EmojiCoolDown uint8
}

type QCandyPvpMapMetaData struct {
	// Properties
	Type         uint8
	IsNew        bool
	Name         AddrTo[Hash]
	Desc         AddrTo[Hash]
	StageID      int32
	Image        AddrTo[StrWithPrefix16]
	RecordImage  AddrTo[StrWithPrefix16]
	RotateImage  AddrTo[StrWithPrefix16]
	SkillList    AddrTo[[]int32]
	Lives        uint8
	CountDown    int32
	Score        int32
	JsonName     AddrTo[StrWithPrefix16]
	TimeLineTime float32
	TimeLineName AddrTo[StrWithPrefix16]
}

type QCandyPvpMapPoolMetaData struct {
	// Properties
	MapPool AddrTo[[]int32]
}

type QCandyPvpMapSkillMetaData struct {
	// Properties
	Icon        AddrTo[StrWithPrefix16]
	Name        AddrTo[Hash]
	Desc        AddrTo[Hash]
	CD          int32
	MaxUseTimes int32
	AbilityName AddrTo[StrWithPrefix16]
	ParamList   AddrTo[[]float32]
}

type QCandyRankSectionMetaData struct {
	// Properties
	RankScoreBegin int32
	RankScoreEnd   int32
}

type QCandySettleConfigMetaData struct {
	// Properties
	RankScoreDelta AddrTo[[]QCandySettleConfigMetaData_Setting]
	AddCurrencyNum int32
}

type QCandySettleConfigMetaData_Setting struct {
	// Fields
	SectionID  uint8
	ScoreDelta int32
}

type QTEndlessMonsterData struct {
	// Properties
	MonsterName      AddrTo[StrWithPrefix16]
	ConfigType       AddrTo[StrWithPrefix16]
	IsElite          bool
	UniqueID         int32
	Score            int32
	Nature           uint8
	MonsterTagList   AddrTo[[]StrWithPrefix16]
	MonsterRank      uint8
	Weight           uint8
	GroupLevelRange  AddrTo[[]int32]
	MonsterDisplayID int32
	// Tags AddrTo[[]StrWithPrefix16]

}

type QTEndlessMonsterWaveMetaData struct {
	// Properties
	RankList         AddrTo[[]StrWithPrefix16]
	MonsterGroupType uint8
	TagDistRule      AddrTo[StrWithPrefix16]
	TagFilter        AddrTo[StrWithPrefix16]
	HPRatio          float32
	// DistRule AddrTo[[]StrWithPrefix16]
	// RankDist AddrTo[[]int32]

}

type QuickEntryScheduleMetaData struct {
	// Properties
	EntryType    int32
	Priority     int32
	PrefabPath   AddrTo[StrWithPrefix16]
	TitleTextMap AddrTo[StrWithPrefix16]
	ImagePath    AddrTo[StrWithPrefix16]
	Link         AddrTo[StrWithPrefix16]
	StartTime    AddrTo[StrWithPrefix16]
	EndTime      AddrTo[StrWithPrefix16]
	MinLevel     int32
	MaxLevel     int32
}

type QuickSellDataMetaData struct {
	// Properties
	SellDialogCrtNumDesc        AddrTo[Hash]
	SellDialogSellNumDesc       AddrTo[Hash]
	SellDialogUnitPriceDesc     AddrTo[Hash]
	SellDialogObtainNumDesc     AddrTo[Hash]
	IsDisplayVitality           int32
	SellDialogSellButtonDesc    AddrTo[Hash]
	SellDialogInfoTitle         AddrTo[Hash]
	SellDialogInfoContent       AddrTo[Hash]
	AdvSellDialogSellButtonDesc AddrTo[Hash]
}

type RaffleClientMetaData struct {
	// Properties
	RafflePic       AddrTo[StrWithPrefix16]
	FirstRewardType AddrTo[StrWithPrefix16]
	RewardLevel     uint8
	FaceBeginTime   AddrTo[StrWithPrefix16]
	FaceEndTime     AddrTo[StrWithPrefix16]
	FirstShow       uint32
	SecondShow      uint32
	ThirdShow       uint32
	HitFaceCD       int32
	HaveLink        bool
	LinkButton      AddrTo[StrWithPrefix16]
	RaffleColor     AddrTo[StrWithPrefix16]
}

type RaffleConfigMetaData struct {
	// Properties
	FirstRewardNumberList  AddrTo[[]uint32]
	MaxFirstRewardNum      uint32
	SecondRewardNumberList AddrTo[[]uint32]
	CostMaterialList       AddrTo[[]BeastStageDisplayMetaData_LevelDropDisplay]
	MaxDrawTimes           uint32
	RaffleGroup            int32
	TextmapRule            AddrTo[Hash]
}

type RaffleRewardConfigMetaData struct {
	// Properties
	FirstRewardId  uint32
	SecondRewardId uint32
	ThirdRewardId  uint32
}

type RaffleScheduleConfigMetaData struct {
	// Properties
	Type                    uint8
	ExchangeCostMaterialId  int32
	ExchangeGetMaterialList AddrTo[[]int32]
}

type RaidActivateScheduleMetaData struct {
	// Properties
	DailyMPActivityList AddrTo[[]int32]
	RaidActivityList    AddrTo[[]int32]
}

type RanchAreaConfigMetaData struct {
	// Properties
	AddSlot            int32
	ProductLimit       int32
	ProductDisplayList AddrTo[[]int32]
	UnlockTextmap      AddrTo[Hash]
}

type RanchAreaMetaData struct {
	// Properties
	ContentList     AddrTo[[]int32]
	ExitRanchSiteID int32
	Collection      AddrTo[[]RanchAreaMetaData_DropDisplay]
	IsChallenge     bool
	Location        AddrTo[[]float32]
	Boundary        float32
	BackGroundType  uint8
}

type RanchAreaMetaData_DropDisplay struct {
	// Fields
	IsOvert   uint8
	MonsterID int32
}

type RanchChallengeSiteMetaData struct {
	// Properties
	MissionList    AddrTo[[]int32]
	ModelMonsterID int32
	ShowRank       bool
}

type RanchContentMetaData struct {
	// Properties
	Type          uint8
	Para          int32
	Icon          AddrTo[StrWithPrefix16]
	NoticeTextMap AddrTo[Hash]
}

type RanchDataMetaData struct {
	// Properties
	ReturnRatio           float32
	ReturnMaterialID      int32
	MonsterSyntheticRatio float32
	LockSkillNumOverallID int32
	WarehouseLimit        int32
}

type RanchMonsterExpMetaData struct {
	// Properties
	Exp        int32
	UnlockTime AddrTo[RemoteTime]
	UnlockSite int32
}

type RanchMonsterMetaData struct {
	// Properties
	ActiveSkill         int32
	MonsterName         AddrTo[Hash]
	MaxLevel            uint8
	BattlePicPath       AddrTo[StrWithPrefix16]
	BattleStagePicPath  AddrTo[StrWithPrefix16]
	HeadPicPath         AddrTo[StrWithPrefix16]
	RanchHeadPicPath    AddrTo[StrWithPrefix16]
	CanUpgrade          bool
	CanBeConsumed       bool
	EasterEggSite       AddrTo[[]int32]
	BasicExp            int32
	MonsterDesc         AddrTo[Hash]
	UnlockTime          AddrTo[RemoteTime]
	GrazingEfficiency   int32
	IsAvatar            bool
	UniqueID            int32
	ActiveSkillCD       AddrTo[[]RanchMonsterMetaData_CDRange]
	SkillTypeMapDisplay AddrTo[[]int32]
	ObtainType          uint8
	DropParam           int32
	CollectOrder        int32
}

type RanchMonsterMetaData_CDRange struct {
	// Fields
	MinLevel uint8
	CoolDown int32
}

type RanchMonsterSkillMetaData struct {
	// Properties
	BasicType uint8
	SkillType uint16
	SkillName AddrTo[StrWithPrefix16]
	Name      AddrTo[Hash]
	Icon      AddrTo[StrWithPrefix16]
	Rarity    int32
}

type RanchSiteDataMetaData struct {
	// Properties
	UnlockCost           int32
	PassedCost           int32
	RanchMonsterDropList AddrTo[[]int32]
	BuffDrop             int32
	IsSiteDialogShow     uint8
	UnlockFlagList       AddrTo[StrWithPrefix16]
	Location             AddrTo[[]float32]
}

type RanchSyntheticMetaData struct {
	// Properties
	FormulaMaterialID  int32
	MainRanchMonster   int32
	AssistRanchMonster int32
	SyntheticProduct   int32
}

type RandomDialogActionMetaData struct {
	// Properties
	ActionType int32
}

type RandomDialogCGRawMetaData struct {
	// Properties
	CGPosType uint8
	Pos       AddrTo[RandomDialogCGRawMetaData_CGRawPos]
}

type RandomDialogCGRawMetaData_CGRawPos struct {
	// Fields
	X float32
	Y float32
}

type RandomDialogInputMetaData struct {
	// Properties
	InputType        int32
	InputLength      int32
	TipText          AddrTo[StrWithPrefix16]
	InputNullTipText AddrTo[StrWithPrefix16]
	PlaceholderText  AddrTo[StrWithPrefix16]
}

type RandomDialogQuestionMetaData struct {
	// Properties
	Talker_name      AddrTo[StrWithPrefix16]
	Question_content AddrTo[StrWithPrefix16]
}

type RandomEffectLevelMetaData struct {
	// Properties
	EffectIcon AddrTo[StrWithPrefix16]
	EffectText AddrTo[StrWithPrefix16]
	EffectName AddrTo[StrWithPrefix16]
}

type RandomEffectMetaData struct {
	// Properties
	EffectType      int32
	EffectQuality   int32
	EffectMaxLevel  int32
	EffectIcon      AddrTo[StrWithPrefix16]
	EffectText      AddrTo[Hash]
	EffectName      AddrTo[Hash]
	StageEffectIcon AddrTo[StrWithPrefix16]
	EffectColor     AddrTo[StrWithPrefix16]
}

type RandomHideBranchTimeMetaData struct {
	// Properties
	HideTime int32
}

type RandomPlotDataMetaData struct {
	// Properties
	StartDialogId      int32
	FinishDialogIdList AddrTo[[]int32]
	BgmDown            float32
	BlackFade          bool
	UserCanControl     bool
	UseSoundBank       bool
	CanSkipOption      bool
	ExternalEvtName    AddrTo[StrWithPrefix16]
}

type RandomPlotNPCMetaData struct {
	// Properties
	Path AddrTo[StrWithPrefix16]
}

type RandomSubSelectContentDataMetaData struct {
	// Properties
	SatusBasedContents AddrTo[[]RandomSubSelectContentDataMetaData_SubSelectContentData]
}

type RandomSubSelectContentDataMetaData_SubSelectContentData struct {
	// Fields
	Status uint32

	// Objects
	Content StrWithPrefix16
}

type RandomSubSelectStyleConfigMetaData struct {
	// Properties
	StyleType uint32
	IntParams AddrTo[[]uint32]
	Icon      AddrTo[StrWithPrefix16]
	UiStyle   bool
}

type RareaffixMetaData struct {
	// Properties
	AffixAttributetype1 int32
	Attributepercent1   int32
	AffixAttributetype2 int32
	Attributepercent2   int32
	Reconfirm_textmap   AddrTo[Hash]
}

type ReclaimConfigMetaData struct {
	// Properties
	CloseTime           AddrTo[StrWithPrefix16]
	EnterIconPrefabPath AddrTo[StrWithPrefix16]
	EnterIconImgPath    AddrTo[StrWithPrefix16]
	BgPrefabPath        AddrTo[StrWithPrefix16]
	OpenCgId            int32
	MissionList         AddrTo[[]int32]
}

type ReclaimEventListMetaData struct {
	// Properties
	ConfigId      int32
	IconPos       int32
	ShowFinishWay int32
	SideMark      AddrTo[StrWithPrefix16]
	UnLockLevel   int32
	ActivityId    int32
}

type ReclaimEventStageListMetaData struct {
	// Properties
	UnlockLevel int32
}

type ReclaimEventVirtualAvatarMetaData struct {
	// Properties
	AvatarId   int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type ReclaimEventVirtualStigmataMetaData struct {
	// Properties
	StigmataId int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type ReclaimEventVirtualWeaponMetaData struct {
	// Properties
	WeaponId   int32
	BaseRarity int32
	BaseLevel  int32
	PreId      int32
}

type ReclaimLevelMetaData struct {
	// Properties
	ConfigId     int32
	Level        int32
	NeedEventExp int32
	RewardId     int32
}

type ReclaimRankingRewardData struct {
	// Properties
	Ranking  int32
	RewardId int32
}

type ReclaimScoreRewardData struct {
	// Properties
	RankingScore int32
	RewardId     int32
}

type ReclaimVirtualGachaChestMetaData struct {
	// Properties
	GachaPoolList      AddrTo[[]ReclaimVirtualGachaChestMetaData_Materials]
	SelectChestImgPath AddrTo[StrWithPrefix16]
	ChestIconID        int32
	ChestName          AddrTo[Hash]
	ChestDesc          AddrTo[Hash]
	ChestRarity        int32
	ShowRewardId       int32
}

type ReclaimVirtualGachaChestMetaData_Materials struct {
	// Fields
	MaterialID int32
	Unused     int32
}

type RecommendEquipMetaData struct {
	// Properties
	GroupAdvsID   int32
	RecommendText AddrTo[Hash]
	RecommendStar int32
}

type RedEnvelopeMetaData struct {
	// Properties
	Num                 int32
	Channellist         AddrTo[[]int32]
	DefaultText         AddrTo[StrWithPrefix16]
	ChatTakeIcon        AddrTo[StrWithPrefix16]
	OpenPageTitleText   AddrTo[Hash]
	ClickIcon           AddrTo[StrWithPrefix16]
	BackgroundImg       AddrTo[StrWithPrefix16]
	RecordBackgroundImg AddrTo[StrWithPrefix16]
}

type RegionInfoMetaData struct {
	// Properties
	Name AddrTo[StrWithPrefix16]
}

type RelationActivityScheduleMetaData struct {
	// Properties
	ShopType         int32
	InfoTitle        AddrTo[Hash]
	InfoContent      AddrTo[Hash]
	ActivityTimeDesc AddrTo[Hash]
	BeginTime        AddrTo[StrWithPrefix16]
	EndTime          AddrTo[StrWithPrefix16]
	MaterialID       int32
	LimitNum         int32
	WikiLink         AddrTo[StrWithPrefix16]
}

type RelationStageBonusDropMetaData struct {
}

type ReplayLobbyFilterMetaData struct {
	// Properties
	LobbyType int32
	KeyID     int32
	Priority  int32
	IconPath  AddrTo[StrWithPrefix16]
	Name      AddrTo[StrWithPrefix16]
}

type ReplayLobbyPageMetaData struct {
	// Properties
	Type     int32
	TabName  AddrTo[StrWithPrefix16]
	Priority int32
	IconPath AddrTo[StrWithPrefix16]
	IsShow   uint8
}

type ReplayLobbyScoreMetaData struct {
	// Properties
	MaxDisplayNum int32
}

type ReplayLobbySubPageMetaData struct {
	// Properties
	TabName AddrTo[StrWithPrefix16]
}

type ReplayLobbyUploadMetaData struct {
	// Properties
	MinLevel        int32
	MaxLevel        int32
	ExBossUploadNum int32
}

type ReportReasonMetaData struct {
	// Properties
	ReasonName   AddrTo[Hash]
	ReasonDetail AddrTo[Hash]
}

type ResourceRetrieveBoxMetaData struct {
	// Properties
	BoxIcon          AddrTo[StrWithPrefix16]
	BoxDes           AddrTo[Hash]
	RewardID         int32
	CostStamina      int32
	CostHcoin        int32
	CostMaterialList AddrTo[[]ResourceRetrieveBoxMetaData_CostMaterial]
}

type ResourceRetrieveBoxMetaData_CostMaterial struct {
	// Fields
	CostId  int32
	CostNum int32
}

type ResourceRetrieveModuleMetaData struct {
	// Properties
	NormalBoxID int32
	SuperBoxID  int32
	ModuleDes   AddrTo[Hash]
}

type ResourceRetrieveScheduleMetaData struct {
}

type RestaurantActionZoneMetaData struct {
	// Properties
	MoveSpeed   float32
	IntervalMin float32
	IntervalMax float32
	BoundaryMin float32
	CollisionX  float32
	CollisionY  float32
}

type RestaurantAvatarMetaData struct {
	// Properties
	Rank                uint8
	RestaurantSkillList AddrTo[[]int32]
	AvatarName          AddrTo[Hash]
	AvatarDesc          AddrTo[Hash]
	AvatarChibiIcon     AddrTo[StrWithPrefix16]
	AvatarImgPath       AddrTo[StrWithPrefix16]
	PrefabPath          AddrTo[StrWithPrefix16]
	RandomBubbleID      AddrTo[[]Hash]
	CookBehaviour       uint8
	FirstBubbleDuration AddrTo[[]float32]
	AfterBubbleDuration AddrTo[[]float32]
}

type RestaurantIngredientsMetaData struct {
	// Properties
	StockLimit AddrTo[[]RestaurantIngredientsMetaData_LvLimit]
	ImagePath  AddrTo[StrWithPrefix16]
}

type RestaurantIngredientsMetaData_LvLimit struct {
	// Fields
	Level uint8
	Limit int32
}

type RestaurantLevelMetaData struct {
	// Properties
	UpgradeCost          int32
	RequireMissionIDList AddrTo[[]int32]
	MaxAvatar            int32
	MaxOrder             int32
	UpgradedPlotID       int32
	UnlockDesc           AddrTo[Hash]
	UnlockIngredientList AddrTo[[]int32]
	UnlockRecipeList     AddrTo[[]int32]
	PreviewImgPath       AddrTo[StrWithPrefix16]
}

type RestaurantQuestMetaData struct {
	// Properties
	MissionID           int32
	SubType             uint8
	RequireMaterialList AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
	AvatarIconPath      AddrTo[StrWithPrefix16]
}

type RestaurantRecipeMetaData struct {
	// Properties
	Tag        uint8
	Ingredient AddrTo[[]ElfBreakMetaData_ElfBreakMaterial]
	SellPrice  int32
	StockLimit AddrTo[[]RestaurantIngredientsMetaData_LvLimit]
	CookSpeed  AddrTo[[]RestaurantRecipeMetaData_LvSpeed]
	SellSpeed  AddrTo[[]RestaurantRecipeMetaData_LvSpeed]
	Bundle     int32
}

type RestaurantRecipeMetaData_LvSpeed struct {
	// Fields
	Level uint8
	Speed float32
}

type RestaurantRoomMetaData struct {
	// Properties
	Type     uint8
	NodePath AddrTo[StrWithPrefix16]
}

type RestaurantSkillMetaData struct {
	// Properties
	SkillType      uint8
	SkillName      AddrTo[Hash]
	SkillDesc      AddrTo[Hash]
	SkillIconPath  AddrTo[StrWithPrefix16]
	SkillParameter AddrTo[[]int32]
	SkillLevel     uint8
}

type RestaurantWeatherMetaData struct {
	// Properties
	SkillID          AddrTo[[]int32]
	WeatherDesc      AddrTo[Hash]
	WeatherEffectDes AddrTo[Hash]
	BeginTime        AddrTo[RemoteTime]
	EndTime          AddrTo[RemoteTime]
	ShowRoleImage    AddrTo[StrWithPrefix16]
	ShowMascot       AddrTo[StrWithPrefix16]
}

type RestrictExtend_StageMetaData struct {
	// Properties
	RestrictID uint32
}

type ReunionCookBookMetaData struct {
	// Properties
	CookType         int32
	CookName         AddrTo[StrWithPrefix16]
	CookDesc         AddrTo[StrWithPrefix16]
	CookConsumeList1 AddrTo[[]ReunionCookBookMetaData_CookConsumerData]
	CookPoint        int32
	CookProduce      int32
	CookDailyTimes   int32
	CookTotalTimes   int32
}

type ReunionCookBookMetaData_CookConsumerData struct {
	// Fields
	Count int32
	ID    int32
}

type ReunionCookRewardMetaData struct {
	// Properties
	NeedScore int32
	RewardID  int32
}

type ReunionCookStageMetaData struct {
	// Properties
	Rate          int32
	BGpicPath     AddrTo[StrWithPrefix16]
	PlotLuaPath   AddrTo[StrWithPrefix16]
	Position      AddrTo[[]int32]
	TablePicPath  AddrTo[StrWithPrefix16]
	RewardDisplay int32
	RewardDesc    AddrTo[StrWithPrefix16]
}

type ReunionDisplayScheduleMetaData struct {
	// Properties
	CollectBeginTime AddrTo[RemoteTime]
	CollectEndTime   AddrTo[RemoteTime]
	RewardBeginTime  AddrTo[RemoteTime]
	RewardEndTime    AddrTo[RemoteTime]
	OpenDayTime      AddrTo[RemoteTimeSpan]
	EndDayTime       AddrTo[RemoteTimeSpan]
}

type ReviveCostTypeMetaData struct {
	// Properties
	Type1 AddrTo[[]int32]
}

type ReviveEffectMetaData struct {
	// Properties
	ReviveHp           int32
	ReviveSp           int32
	SpecialAbilityName AddrTo[StrWithPrefix16]
}

type ReviveUseMetaData struct {
	// Properties
	MaterialID       int32
	IsExclude        bool
	CostCoinList     AddrTo[[]int32]
	ReviveEffect     int32
	ReviveEffectDesc AddrTo[Hash]
}

type RewardData struct {
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
	// Properties
	RankIDList        AddrTo[[]uint16]
	ScoreLinkMaterial int32
}

type RewardLineRankMetaData struct {
	// Properties
	Rank          uint8
	RewardID      int32
	RequiredScore uint32
	IsKeyReward   bool
}

type RewardOverviewPanelMetaData struct {
	// Properties
	ItemID      int32
	ImagePath   AddrTo[StrWithPrefix16]
	Mission     int32
	LinkType    int32
	LinkParams  AddrTo[[]int32]
	DisplayTime AddrTo[StrWithPrefix16]
}

type RichAreaMetaData struct {
	// Properties
	RichID         int32
	TileTowerID    int32
	ThemeID        int32
	AreaCurrencyID int32
	AreaName       AddrTo[StrWithPrefix16]
	BuffList       AddrTo[[]int32]
	ItemList       AddrTo[[]int32]
	QAvartarID     int32
	AreaSubTitle   AddrTo[StrWithPrefix16]
	AreaNumber     AddrTo[StrWithPrefix16]
	AreaTabBGpath  AddrTo[StrWithPrefix16]
	SubToArea      int32
	IsMainArea     int32
}

type RichBuffMetaData struct {
	// Properties
	Value1 float32
	Value2 float32
	Value3 float32
	Name   AddrTo[StrWithPrefix16]
	Desc   AddrTo[StrWithPrefix16]
	Icon   AddrTo[StrWithPrefix16]
	SpDesc AddrTo[StrWithPrefix16]
}

type RichBuildingMetaData struct {
	// Properties
	UpgradeCurrency int32
	UpgradeNum      int32
	FuncType        int32
	Para            AddrTo[StrWithPrefix16]
	Model           int32
	Name            AddrTo[StrWithPrefix16]
	BuildingDesc    AddrTo[StrWithPrefix16]
	UpgradeDesc     AddrTo[StrWithPrefix16]
	Scale           AddrTo[[]float32]
	BuildingIcon    AddrTo[StrWithPrefix16]
}

type RichChallengeBossMetaData struct {
	// Properties
	AreaID                int32
	UnlockFloorID         int32
	StageID               int32
	BossImg               AddrTo[StrWithPrefix16]
	BossExploreUnlockIcon AddrTo[StrWithPrefix16]
	BossExploreLockedIcon AddrTo[StrWithPrefix16]
	BossExplorePassedIcon AddrTo[StrWithPrefix16]
	BossName              AddrTo[StrWithPrefix16]
	RecommendBuff         AddrTo[[]RichChallengeBossMetaData_MonopolyBuff]
	BossDesc              AddrTo[StrWithPrefix16]
	BossUnlockText        AddrTo[StrWithPrefix16]
}

type RichChallengeBossMetaData_MonopolyBuff struct {
	// Fields
	BuffID int32
	BuffLV int32
}

type RichDiceMetaData struct {
	// Properties
	AnimSkipPer float32
}

type RichFloorInfoMetaData struct {
	// Properties
	LocationList    AddrTo[StrWithPrefix16]
	BuildingLimit   AddrTo[[]RichFloorInfoMetaData_BuildLimit]
	MissionList     AddrTo[[]int32]
	RichFloorName   AddrTo[StrWithPrefix16]
	UnlockTime      AddrTo[RemoteTime]
	UnlockText      AddrTo[StrWithPrefix16]
	UnlockText_Time AddrTo[StrWithPrefix16]
	ShopList        AddrTo[StrWithPrefix16]
	IsResetable     int32
	RecommendBuff   AddrTo[[]RichFloorInfoMetaData_BuffInfo]
	OverallID       int32
	RpgCoinLimit    int32
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
	// Properties
	Type        int32
	ParaList    AddrTo[[]int32]
	UseHintText AddrTo[StrWithPrefix16]
}

type RichmanMetaData struct {
	// Properties
	DiceID int32
}

type RichMapInfoMetaData struct {
	// Properties
	RichAreaID   int32
	InfoTitle    AddrTo[StrWithPrefix16]
	InfoText     AddrTo[StrWithPrefix16]
	InfoIconPath AddrTo[StrWithPrefix16]
}

type RichMonsterInfoMetaData struct {
	// Properties
	StageID       int32
	MonsterName   AddrTo[StrWithPrefix16]
	MonsterDesc   AddrTo[StrWithPrefix16]
	RecommendBuff AddrTo[[]RichChallengeBossMetaData_MonopolyBuff]
}

type RichShopMetaData struct {
	// Properties
	ShopID        int32
	Item          int32
	PurchaseLimit int32
	CostItem      int32
	CostNum       int32
}

type RogueBuffPoolMetaData struct {
	// Properties
	IsExtra    bool
	BuffWeight AddrTo[[]RogueBuffPoolMetaData_Pool]
}

type RogueBuffPoolMetaData_Pool struct {
	// Fields
	ID     int32
	Weight int32
}

type RogueTowerDataMetaData struct {
	// Properties
	PoolIDList        AddrTo[[]uint32]
	ReduceMinScore    AddrTo[[]RogueTowerDataMetaData_MaterialReduceMinScore]
	SiteLockTipsParam int32
	BaseHeal          int32
	ExtraHeal         AddrTo[[]RogueTowerDataMetaData_ExtraHealth]
	RogueTowerCoin    int32
	BuffReset         AddrTo[[]RogueTowerDataMetaData_ResetMaterialCost]
	BuffResetCost     AddrTo[[]int32]
	BuffLevelupCost   AddrTo[[]int32]
	MaxSaveDataNum    int32
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
	// Properties
	EndSiteList AddrTo[[]uint32]
}

type RogueTowerStageMetaData struct {
	// Properties
	ScoreRate AddrTo[[]RogueTowerStageMetaData_ScorePercent]
}

type RogueTowerStageMetaData_ScorePercent struct {
	// Fields
	MinScore int32
	Percent  int32
}

type RoutineDataMetaData struct {
	// Properties
	RoutineValue     int32
	CountMax         int32
	RefreshText      AddrTo[Hash]
	MissionImagePath AddrTo[StrWithPrefix16]
}

type RoutineRewardMetaData struct {
	// Properties
	RoutineValue int32
	RewardId     int32
}

type RoutineScheduleMetaData struct {
	// Properties
	WeeklyRewardList AddrTo[[]int32]
	IntroImagePath   AddrTo[StrWithPrefix16]
}

type RpgAreaLineMetaData struct {
	// Properties
	LineNode AddrTo[StrWithPrefix16]
}

type RpgAreaMetaData struct {
	// Properties
	AreaLinkName      AddrTo[StrWithPrefix16]
	AreaType          uint8
	ZoneID            int32
	SiteList          AddrTo[[]int32]
	AreaRewardPreview int32
	ThemeID           int32
	AreaPanelNode     AddrTo[StrWithPrefix16]
	AreaTitle         AddrTo[Hash]
	AreaDesc          AddrTo[Hash]
	AreaPicPath       AddrTo[StrWithPrefix16]
	AreaShowPicPath   AddrTo[StrWithPrefix16]
	AreaPassedPicPath AddrTo[StrWithPrefix16]
	AreaDialogBGPath  AddrTo[StrWithPrefix16]
	AreaInfo          AddrTo[[]Hash]
	AreaLockInfo      AddrTo[Hash]
	AreaBuffList      AddrTo[[]int32]
	AreaMissionList   AddrTo[[]int32]
}

type RpgBuffDataMetaData struct {
	// Properties
	MaxLevel       int32
	BuffSuit       AddrTo[RpgBuffDataMetaData_SuitBuffData]
	WeaponTypeList AddrTo[[]int32]
	Name           AddrTo[Hash]
	Desc           AddrTo[Hash]
	Icon           AddrTo[StrWithPrefix16]
	EffectIcon     AddrTo[StrWithPrefix16]
	AbilityName    AddrTo[StrWithPrefix16]
	ConfigType     AddrTo[StrWithPrefix16]
	Hp             int32
	Atk            int32
	Def            int32
	AllDamageRatio int32
	ParamList      AddrTo[[]StrWithPrefix16]
	BGColor        AddrTo[StrWithPrefix16]
	EffectParam    int32
	CGID           int32
}

type RpgBuffDataMetaData_SuitBuffData struct {
	// Fields
	SuitBuffID int32
	SuitID     int32
	SuitNum    int32
}

type RpgBuffLimitMetaData struct {
	// Properties
	TaleID         int32
	UnlockFlagList AddrTo[StrWithPrefix16]
	CostContent    AddrTo[[]BeastStageDisplayMetaData_LevelDropDisplay]
}

type RpgBuffSuitClientInfoMetaData struct {
	// Properties
	RpgID          int32
	BGColorDisplay AddrTo[StrWithPrefix16]
	SuitID         int32
}

type RpgClientMainDataMetaData struct {
	// Properties
	TaleType             uint8
	BeginEvent           int32
	TaleCurrency         AddrTo[[]int32]
	BuffMaterialList     AddrTo[[]int32]
	GachaLink            AddrTo[StrWithPrefix16]
	TaleWebLink          AddrTo[StrWithPrefix16]
	TaleTutorial         AddrTo[[]GlobalWarScheduleMetaData_ConfigTutorial]
	TaleInfo             AddrTo[[]StrWithPrefix16]
	FuncLinkType         int32
	FuncParam            AddrTo[[]int32]
	ShopLink             int32
	MuseumID             int32
	GenCfgPath           AddrTo[StrWithPrefix16]
	ContextName          AddrTo[StrWithPrefix16]
	WeaponList           AddrTo[[]int32]
	ShowUnlockSiteList   AddrTo[[]int32]
	ShowPriority         int32
	BookPageShowPageList AddrTo[[]int32]
	MenuShowList         AddrTo[[]StrWithPrefix16]
	EntryShowType        int32
}

type RpgCollectionRewardMetaData struct {
	// Properties
	CollectionCount             int32
	PositionShowIcon            AddrTo[StrWithPrefix16]
	LimitTime                   AddrTo[StrWithPrefix16]
	ImagePath                   AddrTo[StrWithPrefix16]
	RewardID                    int32
	RewardProgressUIPrefabPath  AddrTo[StrWithPrefix16]
	RewardPointUIPrefabPath     AddrTo[StrWithPrefix16]
	RewardPointCurrentIconPath  AddrTo[StrWithPrefix16]
	RewardPointLockedIconPath   AddrTo[StrWithPrefix16]
	RewardPointFinishedIconPath AddrTo[StrWithPrefix16]
	IsSpecialReward             int32
}

type RpgConditionMetaData struct {
	// Properties
	Type            uint8
	Parameter       AddrTo[[]int32]
	StringParameter AddrTo[[]StrWithPrefix16]
}

type RpgDungeonAdventureQuestMetaData struct {
	// Properties
	QuestDetail     AddrTo[StrWithPrefix16]
	QuestDetailPic  AddrTo[StrWithPrefix16]
	RequiredTagList AddrTo[[]int32]
	RewardDisplay   AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
}

type RpgEventDialogMetaData struct {
	// Properties
	UnlockFlagList AddrTo[StrWithPrefix16]
}

type RpgEventTextMetaData struct {
	// Properties
	DetailDesc     AddrTo[StrWithPrefix16]
	EventName      AddrTo[StrWithPrefix16]
	EventDesc      AddrTo[StrWithPrefix16]
	EventImagePath AddrTo[StrWithPrefix16]
	RewardDesc     AddrTo[StrWithPrefix16]
}

type RpgFilesMetaData struct {
	// Properties
	TaleID        int32
	FileTitle     AddrTo[Hash]
	ImageShowType int32
	FileImages    AddrTo[[]StrWithPrefix16]
	FileText      AddrTo[Hash]
	FileLockText  AddrTo[Hash]
	FileLockHint  AddrTo[Hash]
	AreaType      AddrTo[StrWithPrefix16]
}

type RpgLevelProgressData struct {
	// Properties
	Progress int32
}

type RpgMaterialMetaData struct {
	// Properties
	TaleID         int32
	ShowType       uint8
	UnlockItemDesc AddrTo[StrWithPrefix16]
	Order          int32
	RpgBuffId      int32
}

type RpgMissionCategoryMetaData struct {
	// Properties
	MissionCategoryList AddrTo[[]int32]
}

type RpgMissionMetaData struct {
	// Properties
	TaleID                     int32
	ShowType                   int32
	CustomData                 AddrTo[StrWithPrefix16]
	MissionIcon                AddrTo[StrWithPrefix16]
	RpgTaleMissionBGDesc       AddrTo[Hash]
	RpgTaleMissionLocationDesc AddrTo[Hash]
	MissionConditionText       AddrTo[Hash]
	IsMainMission              bool
	IsBranchMission            bool
}

type RpgNavSiteMetaData struct {
	// Properties
	UICameraID       AddrTo[[]float32]
	NavRefreshSiteID int32
	NavNextSiteList  AddrTo[[]int32]
	FogofWar         AddrTo[[]int32]
}

type RpgOverallMetaData struct {
	// Properties
	IsClientSet bool
	Min         int32
	Max         int32
}

type RpgPlotConditionMetaData struct {
	// Properties
	TaleID    int32
	OverallID int32
	Value     int32
}

type RpgQAvatarBattleCollideMetaData struct {
	// Properties
	CollideSet AddrTo[[]int32]
}

type RpgQAvatarBattleDivisionMetaData struct {
	// Properties
	DivisionName          AddrTo[Hash]
	NeedBattleScore       uint32
	FirstDivisionUpReward int32
}

type RpgQAvatarBattleSceneMetaData struct {
	// Properties
	StagePath         AddrTo[StrWithPrefix16]
	ClientGridMap     AddrTo[StrWithPrefix16]
	SceneGroundHeight float32
}

type RpgQAvatarBattleSiteMetaData struct {
	// Properties
	BGM              AddrTo[StrWithPrefix16]
	BattleTimeMax    uint32
	WinCondition     int32
	FinishBattleKill int16
	DieEffect        AddrTo[StrWithPrefix16]
	GadgetBuffName   AddrTo[StrWithPrefix16]
	GadgetBuffIcon   AddrTo[StrWithPrefix16]
	BeatKillNum      int32
	BeatRemainTime   int32
}

type RpgRestaurantMetaData struct {
	// Properties
	CoinMaterialID     int32
	IngredientList     AddrTo[[]int32]
	RecipeList         AddrTo[[]int32]
	QuestList          AddrTo[[]int32]
	AccelerateTicketID int32
}

type RpgRoleMetaData struct {
	// Properties
	PairAvatarID         int32
	RoleType             int32
	Fake_Avatar_Registry AddrTo[StrWithPrefix16]
	Fake_Avatar_Config   AddrTo[StrWithPrefix16]
	Color                AddrTo[StrWithPrefix16]
	AtkIcon              AddrTo[StrWithPrefix16]
	AtkText              AddrTo[Hash]
	HpIcon               AddrTo[StrWithPrefix16]
	HpText               AddrTo[Hash]
	SsIcon               AddrTo[StrWithPrefix16]
	SsText               AddrTo[Hash]
	SsBase               AddrTo[Hash]
	ScIcon               AddrTo[StrWithPrefix16]
	ScText               AddrTo[Hash]
	ScBase               AddrTo[Hash]
	SurvivalNameText     AddrTo[Hash]
	SurvivalDescText     AddrTo[Hash]
	SurvivalWeaponName   AddrTo[Hash]
	SurvivalWeaponIcon   AddrTo[StrWithPrefix16]
	SurvivalWeaponDesc   AddrTo[Hash]
	QImgPath             AddrTo[StrWithPrefix16]
	QDrawPath            AddrTo[StrWithPrefix16]
	QPicPath             AddrTo[StrWithPrefix16]
	SurvivalGetText      AddrTo[StrWithPrefix16]
	QavatarProductTip    AddrTo[StrWithPrefix16]
}

type RpgScheduleMetaData struct {
	// Properties
	TaleIDList AddrTo[[]int32]
	BeginTime  AddrTo[StrWithPrefix16]
	EndTime    AddrTo[StrWithPrefix16]
}

type RpgShootPlayTriggerMetaData struct {
	// Properties
	EventType           int32
	StageID             int32
	Prefab              AddrTo[StrWithPrefix16]
	PrefabSpawn         AddrTo[StrWithPrefix16]
	InLevelTriggerType  int32
	InLevelTriggerParam AddrTo[[]float32]
}

type RpgSimpleBuffMetaData struct {
	// Properties
	Hp             int32
	Atk            uint16
	Def            uint16
	CostNum        uint16
	AllDamageRatio uint16
	BuffEffectDesc AddrTo[StrWithPrefix16]
	BuffDesc       AddrTo[StrWithPrefix16]
}

type RpgSiteLineMetaData struct {
	// Properties
	LineNode AddrTo[StrWithPrefix16]
}

type RpgSiteLockTipsMetaData struct {
	// Properties
	SiteLockTips AddrTo[Hash]
}

type RpgSiteMetaData struct {
	// Properties
	SiteType             uint8
	SiteNodeName         AddrTo[StrWithPrefix16]
	HidePic              AddrTo[StrWithPrefix16]
	ShowPic              AddrTo[StrWithPrefix16]
	UnlockPic            AddrTo[StrWithPrefix16]
	CoolDownPic          AddrTo[StrWithPrefix16]
	PassedPic            AddrTo[StrWithPrefix16]
	ClosePic             AddrTo[StrWithPrefix16]
	TipsPic              AddrTo[StrWithPrefix16]
	ShadowPic            AddrTo[StrWithPrefix16]
	StageID              int32
	EnterCoinType        int32
	CoinNum              int32
	SiteProgressConfig   AddrTo[[]int32]
	SiteTitle            AddrTo[Hash]
	SiteDes              AddrTo[Hash]
	SiteUnlockPanelTitle AddrTo[Hash]
	SiteUnlockPanelDes   AddrTo[Hash]
	SiteUnlockPanelTime  float32
	SpecialReward        int32
	SiteIndexDes         AddrTo[Hash]
	SiteLockTips         AddrTo[Hash]
	SiteRewardPreview    AddrTo[StrWithPrefix16]
	ContentID            int32
	LimitID              int32
}

type RpgSitePackMetaData struct {
	// Properties
	AreaID     int32
	UnlockTime AddrTo[RemoteTime]
	SiteList   AddrTo[[]int32]
	PackName   AddrTo[StrWithPrefix16]
	PackImage  AddrTo[StrWithPrefix16]
}

type RpgSiteProgressConfigMetaData struct {
	// Properties
	ProgressType      uint8
	ProgressMaxNum    int32
	ProgressEventList AddrTo[[]int32]
	ProgressIcon      AddrTo[StrWithPrefix16]
	ProgressDesc      AddrTo[StrWithPrefix16]
}

type RpgSkillDisplayMetaData struct {
	// Properties
	TypeName    AddrTo[StrWithPrefix16]
	PrefabPath  AddrTo[StrWithPrefix16]
	UniqueTag   int32
	UIPointName AddrTo[StrWithPrefix16]
	UILineName  AddrTo[[]StrWithPrefix16]
	UnlockDesc  AddrTo[Hash]
	SkillDesc   AddrTo[Hash]
}

type RpgStageDropMetaData struct {
	// Properties
	DailyDropLimit AddrTo[[]int32]
	MaxCurrencyNum int32
}

type RpgStageScoreMetaData struct {
	// Properties
	StageID        int32
	MinScore       int32
	CurrencyID     int32
	MaxCurrencyNum int32
	CurrencyNum    int32
}

type RpgStigmataBookMetaData struct {
	// Properties
	GroupID        uint32
	Order          uint8
	UnlockPic      AddrTo[StrWithPrefix16]
	UnlockItemDesc AddrTo[StrWithPrefix16]
	IsMain         bool
}

type RpgSuddenEventMetaData struct {
	// Properties
	EventName      AddrTo[StrWithPrefix16]
	EventDesc      AddrTo[StrWithPrefix16]
	EventImagePath AddrTo[StrWithPrefix16]
}

type RpgSurvivalRoleMetaData struct {
	// Properties
	RPGRoleID       int32
	QAvatarPVPID    int32
	QAvatarRoleName AddrTo[StrWithPrefix16]
	AttrIcon        AddrTo[StrWithPrefix16]
	AttrImg         AddrTo[StrWithPrefix16]
	WeaponId        int32
	Star            int32
	AtkBase         int32
	HpBase          int32
	Prop1           int32
	Prop1RPGRoleID  int32
	Prop1Param1     float32
	Prop1Param2     float32
	Prop2           int32
	Prop2RPGRoleID  int32
	Prop2Param1     float32
	Prop2Param2     float32
	Prop3           int32
	Prop3RPGRoleID  int32
	Prop3Param1     float32
	Prop3Param2     float32
	Tspeed          int32
	Mspeed          int32
	TaleID          int32
	RoleBaseScore   int32
}

type RpgSurvivalTraitMetaData struct {
	// Properties
	TitleText    AddrTo[Hash]
	TraitText    AddrTo[Hash]
	ItemIcon     AddrTo[StrWithPrefix16]
	AbilityName  AddrTo[StrWithPrefix16]
	PVPSkillName AddrTo[StrWithPrefix16]
}

type RpgSurvivalWeaponMetaData struct {
	// Properties
	WeaponType int32
	AttackPara int32
	LockRange  int32
	Range      int32
	AtkSpeed   int32
}

type RpgTaleMetaData struct {
	// Properties
	TicketIDList            AddrTo[[]int32]
	AreaList                AddrTo[[]int32]
	MinLevel                uint16
	MaxLevel                uint16
	CollectionRewardType    uint8
	CollectionRewardParam   AddrTo[[]int32]
	GenActivityRewardShowID uint16
	QAvatarPvpMaterialLimit AddrTo[RpgTaleMetaData_DailyMaterialLimit]
	VirtualGroupID          int32
	StageList               AddrTo[[]int32]
	TileTowerList           AddrTo[[]int32]
	RichID                  int32
	DailyDropLimit          AddrTo[StrWithPrefix16]
	LinkedActivityType      int32
	LinkedActivityID        uint32
}

type RpgTaleMetaData_DailyMaterialLimit struct {
	// Fields
	LimitNum   int32
	MaterialID int32
}

type RpgTaleRefreshMetaData struct {
	// Properties
	CostHcoinNum int32
	CostMaterial AddrTo[[]MechMetaData_DisjoinOutputItem]
	RefreshTimes int32
}

type RpgTaleStageEnterTimesMetaData struct {
	// Properties
	TaleID     int32
	StageList  AddrTo[[]int32]
	EnterTimes int32
}

type RpgTicketMetaData struct {
	// Properties
	MaterialID int32
	MaxNum     int32
}

type RpgTowerV2MetaData struct {
	// Properties
	RewardID           int32
	Score              int32
	DisplayText        AddrTo[Hash]
	MonsterDisplayList AddrTo[[]int32]
	RandomBuffList     AddrTo[[]int32]
	FixedBuffList      AddrTo[[]int32]
	RandomType         uint8
	SaveFloor          int16
}

type RpgUIConfigMetaData struct {
	// Properties
	UIManager              AddrTo[StrWithPrefix16]
	AlbumPage              AddrTo[StrWithPrefix16]
	BookPage               AddrTo[StrWithPrefix16]
	PreparePage            AddrTo[StrWithPrefix16]
	AvatarUnlockDialog     AddrTo[StrWithPrefix16]
	ChapterUnlockDialog    AddrTo[StrWithPrefix16]
	RewardDialog           AddrTo[StrWithPrefix16]
	SubStageUnlockDialog   AddrTo[StrWithPrefix16]
	CollectionRewardDialog AddrTo[StrWithPrefix16]
	RpgSiteSelectDialog    AddrTo[StrWithPrefix16]
	RpgFileInfoDialog      AddrTo[StrWithPrefix16]
	MissionGotDialog       AddrTo[StrWithPrefix16]
	EntrancePrefab         AddrTo[StrWithPrefix16]
	CoinPrefab             AddrTo[StrWithPrefix16]
	BuffPrefab             AddrTo[StrWithPrefix16]
	MatrixPage             AddrTo[StrWithPrefix16]
	StigmataPage           AddrTo[StrWithPrefix16]
}

type RpgVirtualAvatarLimitMetaData struct {
	// Properties
	ConditionList AddrTo[StrWithPrefix16]
	ConditionTips AddrTo[Hash]
}

type RpgZoneMetaData struct {
	// Properties
	ZoneNode      AddrTo[StrWithPrefix16]
	ZoneName      AddrTo[Hash]
	UnlockTextMap AddrTo[Hash]
}

type SanctuaryBuildingActivityMetaData struct {
	// Properties
	MaxSanctuaryLevel int32
	ShopType          int32
	CurrencyID        int32
	GetTimeCD         int32
	BgPath            AddrTo[StrWithPrefix16]
	AvatarBgPath      AddrTo[StrWithPrefix16]
}

type SanctuaryBuildingLevelMetaData struct {
	// Properties
	LevelExp       int32
	BuildingIcon   AddrTo[StrWithPrefix16]
	ProductOutput  int32
	ProductStorage int32
	DescTextmapID  AddrTo[Hash]
}

type SanctuaryPlayerGroupMetaData struct {
	// Properties
	MinPlayerLevel int32
	MaxPlayerLevel int32
}

type SanctuaryStageConfigMetaData struct {
	// Properties
	StageRarityIcon       AddrTo[StrWithPrefix16]
	StageRarityFrameColor AddrTo[StrWithPrefix16]
	StageTypeIcon         AddrTo[StrWithPrefix16]
	StageBGIcon           AddrTo[StrWithPrefix16]
	RewardExpMin          int32
	RewardExpMax          int32
}

type ScDLCAchievementMetaData struct {
	// Properties
	Title            AddrTo[Hash]
	Desc             AddrTo[Hash]
	Icon             AddrTo[StrWithPrefix16]
	SequenceID       int32
	IsShowProgress   bool
	AchieveMissionID int32
	IsHide           bool
}

type ScDLCAvatarLevelMetaData struct {
	// Properties
	LevelUpCost     AddrTo[[]ScDLCTalentLevelMetaData_LevelUpMatRequire]
	HPAdd           int32
	ATKAdd          int32
	DEFAdd          int32
	LevelReward     int32
	FeverLevelLimit int32
}

type ScDLCChallengeMetaData struct {
	// Properties
	UnlockTime AddrTo[StrWithPrefix16]
}

type ScDLCChallengeRewardMetaData struct {
	// Properties
	StarNum  int32
	RewardID int32
}

type ScDLCExplorePointMaterialMetaData struct {
	// Properties
	MaterialByEvent AddrTo[StrWithPrefix16]
}

type ScDLCFeverAbilityMetaData struct {
	// Properties
	AbilityGroupID   int32
	Ability          AddrTo[StrWithPrefix16]
	Icon             AddrTo[StrWithPrefix16]
	Param            float32
	Desc             AddrTo[Hash]
	CurrentLevelDesc AddrTo[Hash]
	NextLevelDesc    AddrTo[Hash]
	RemainDesc       AddrTo[Hash]
	UnlockLevel      int32
	ShowLevel        int32
	Title            AddrTo[Hash]
	ShortName        AddrTo[Hash]
}

type ScDLCFurnitureMetaData struct {
	// Properties
	PrefabPath    AddrTo[StrWithPrefix16]
	ConditionList AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	NPCHeadMark   AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
}

type ScDLCImageReplaceMetaData struct {
	// Properties
	TargetImg AddrTo[StrWithPrefix16]
	Condition int32
}

type ScDLCInteractStateMetaData struct {
	// Properties
	IsShow                  bool
	Interactive             bool
	InteractName            AddrTo[StrWithPrefix16]
	InteractIcon            AddrTo[StrWithPrefix16]
	InteractRadius          float32
	InteractAngle           float32
	NPCID                   int32
	InteractEnterActionList AddrTo[[]int32]
	InteractExitActionList  AddrTo[[]int32]
}

type ScDLCLevelMetaData struct {
	// Properties
	Exp                        uint16
	NextLevelDesc              AddrTo[Hash]
	FeverReward                int32
	LevelReward                int32
	HighLight                  bool
	OwHardlv                   AddrTo[StrWithPrefix16]
	MaterialDropLimitAddedList AddrTo[[]DLCLevelMetaData_ItemLimit]
}

type ScDLCMapProgressMetaData struct {
	// Properties
	Desc          AddrTo[StrWithPrefix16]
	AreaID        int32
	EventList     AddrTo[[]int32]
	ShowCondition int32
}

type ScDLCMissionRewardMetaData struct {
	// Properties
	ID            uint16
	TypeID        int32
	TypeName      AddrTo[Hash]
	ShowCondition int32
}

type ScDLCMPStageDisplayMetaData struct {
	// Properties
	FirstRewardDisplay int32
	RewardDisplay      int32
}

type ScDLCNPCMetaData struct {
	// Properties
	PrefabPath    AddrTo[StrWithPrefix16]
	PrefabScale   float32
	ConditionList AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
	NPCHeadMark   AddrTo[[]ChapterOWBuildingMetaData_ConditionStatePair]
}

type ScDLCNPCPositionMetaData struct {
	// Properties
	NPCPosList  AddrTo[[]ChapterOWNPCPositionMetaData_Positions]
	ConditionID int32
}

type ScDLCOpenFunctionMetaData struct {
	// Properties
	ConditionID    int32
	BtnPath        AddrTo[StrWithPrefix16]
	ShowUnlockTips bool
	Title          AddrTo[Hash]
	Desc           AddrTo[Hash]
	TitleLabel     AddrTo[Hash]
	IconPath       AddrTo[StrWithPrefix16]
}

type ScDLCShowMapConditionMetaData struct {
	// Properties
	ConditionID int32
	MaskPath    AddrTo[StrWithPrefix16]
}

type ScDLCSkillChipMetaData struct {
	// Properties
	Icon                 AddrTo[StrWithPrefix16]
	Desc                 AddrTo[Hash]
	Cost                 int32
	Ability              AddrTo[StrWithPrefix16]
	DisplayID            int32
	MatchFeverAbilityIDs AddrTo[[]int32]
	Title                AddrTo[Hash]
}

type ScDLCSpecialExplorePointMetaData struct {
	// Properties
	ConditionID        int32
	MapID              int32
	AreaID             int32
	TraceStoryID       int32
	AllowTransfer      bool
	RequiredAvatarList AddrTo[[]int32]
	RequiredAvatarHint AddrTo[Hash]
}

type ScDLCSpecialItemMetaData struct {
	// Properties
	RemindType uint8
}

type ScDLCStoryMetaData struct {
	// Properties
	Action uint8
}

type ScDLCTeachStageMetaData struct {
	// Properties
	ShowCondition     int32
	AvatarList        AddrTo[[]int32]
	AvatarListType    AddrTo[[]int32]
	SupportAvatar     int32
	SupportAvatarType int32
}

type ScDLCTowerFloorMetaData struct {
	// Properties
	RecommendLevel int32
	WarningLevel   int32
	HardLevel      int32
	SpawnType      int32
	SpawnTypeDesc  AddrTo[StrWithPrefix16]
	Parameter      AddrTo[[]int32]
	WaveListPool   AddrTo[[]int32]
	TimeBonus      int32
	IsCheckPoint   int32
	Reward         int32
	StageSection   AddrTo[StrWithPrefix16]
	HardLevelGroup int32
	MaxCoin        int32
	MaxScore       int32
	UnlockStoryID  int32
	MaxTime        int32
}

type ScDLCTowerScoreRewardMetaData struct {
	// Properties
	Reward int32
}

type ScratchTicketDataMetaData struct {
	// Properties
	TicketGoodsIDList   AddrTo[[]int32]
	PlateLevelNum       int32
	TicketID            int32
	ResetItemID         int32
	ResetItemNumList    AddrTo[[]int32]
	MaxRoundNum         int32
	TicketOutlineImg    AddrTo[StrWithPrefix16]
	TicketGoodsImgPath1 AddrTo[StrWithPrefix16]
	TicketGoodsImgPath2 AddrTo[StrWithPrefix16]
	TicketGoodsImgPath3 AddrTo[StrWithPrefix16]
}

type ScratchTicketItemDataMetaData struct {
	// Properties
	ItemID     int32
	ItemLv     int32
	ItemNum    int32
	IsCoreItem bool
	Rarity     int32
}

type ScratchTicketPlateDataMetaData struct {
	// Properties
	IsCoreGacha     bool
	CardBackImgPath AddrTo[StrWithPrefix16]
}

type SecondAnniversaryMemoryMetaData struct {
	// Properties
	MemoryName  AddrTo[Hash]
	MemoryText  AddrTo[Hash]
	MemoryImage AddrTo[StrWithPrefix16]
}

type SecondAnniversaryMissionMetaData struct {
	// Properties
	MissionImage AddrTo[StrWithPrefix16]
}

type SeriesMetaData struct {
	// Properties
	Title         AddrTo[Hash]
	EnableScore   int32
	SweepMaxTimes uint8
}

type ServantDialogMetaData struct {
	// Properties
	Condition int32
	Weight    int32
	Content   AddrTo[Hash]
	AudioID   AddrTo[StrWithPrefix16]
	Emotion   AddrTo[StrWithPrefix16]
	Position  int32
	Time      float32
}

type ServantTouchMetaData struct {
	// Properties
	BodyStateName            AddrTo[StrWithPrefix16]
	FaceStateNameList        AddrTo[[]StrWithPrefix16]
	FaceEffectName           AddrTo[StrWithPrefix16]
	SpecialTimes             int32
	SpecialBodyStateName     AddrTo[StrWithPrefix16]
	SpecialFaceStateNameList AddrTo[[]StrWithPrefix16]
	SpecialFaceEffectName    AddrTo[StrWithPrefix16]
}

type SettingAudioVolumeMetaData struct {
	// Properties
	Volume float32
}

type SettingGraphicsDeviceLimitMetaData struct {
	// Properties
	ParamBlackList AddrTo[[]uint8]
}

type SettingGraphicsItemLineMetaData struct {
	// Properties
	GraphicsSettingItemText AddrTo[Hash]
	GraphicsSettingItemName AddrTo[StrWithPrefix16]
	ItemType                uint8
	OccupiedCols            uint8
	ChildrenColCount        uint8
	ChildrenLineList        AddrTo[[]int32]
}

type SettingGraphicsTitleLineMetaData struct {
	// Properties
	SettingGraphicsItemLineText AddrTo[Hash]
	SettingGraphicsItemLineList AddrTo[[]int32]
	ItemColCount                uint8
}

type ShareChannelContentMetaData struct {
	// Properties
	Title   AddrTo[Hash]
	URL     AddrTo[Hash]
	Text    AddrTo[Hash]
	Topic   AddrTo[Hash]
	Section AddrTo[Hash]
	Image   AddrTo[StrWithPrefix16]
}

type ShareChannelMetaData struct {
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
	ShareBtnIcon    AddrTo[StrWithPrefix16]
	AllowWebLink    bool
}

type SharePakDataMetaData struct {
	// Properties
	SubPakId int32
}

type ShareShowDetailConfigMetaData struct {
	// Properties
	Paths AddrTo[[]StrWithPrefix16]
}

type ShareSwitchMetaData struct {
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

type ShigureKiraActivityMetaData struct {
	// Properties
	DropLimitScheduleID      int32
	ProgressAdditionOneClick int32
	TotalProgress            int32
	ShopMaterialID           int32
	ProgressMaterialID       int32
	QuestAreaIDList          AddrTo[[]int32]
	ProgressIDList           AddrTo[[]int32]
	StrengthenIDRange        AddrTo[ShigureKiraActivityMetaData_AbilityIDRange]
}

type ShigureKiraActivityMetaData_AbilityIDRange struct {
	// Fields
	Max int32
	Min int32
}

type ShigureKiraFanchantMemMetaData struct {
	// Properties
	ProgressID           int32
	QavatarID            int32
	QavatarName          AddrTo[Hash]
	FanchantSkillName    AddrTo[Hash]
	FanchantSkillDetail  AddrTo[Hash]
	RecruitConditionTips AddrTo[Hash]
	RewardID             int32
	UnlockConditionTips  AddrTo[Hash]
	QAvatarIconPath      AddrTo[StrWithPrefix16]
	QAvatarTips          AddrTo[Hash]
	BuffList             AddrTo[[]int32]
	AreaID               int32
}

type ShigureKiraProgressPointMetaData struct {
	// Properties
	SiteID          int32
	SlotType        uint8
	Progress        int32
	CostMaterialNum int32
	FanchantID      int32
	UnlockTips      AddrTo[Hash]
}

type ShigureKiraSiteMetaData struct {
	// Properties
	MinDropNum    int32
	MaxDropNum    int32
	QuestSiteType uint8
}

type ShigureKiraStrengthenMetaData struct {
	// Properties
	ProgressAddition      int32
	NextStrengthenID      int32
	StrengthenATK         int32
	StrengthenAddedDamage int32
}

type ShopAboutToOnlineItem struct {
	// Properties
	BeginDate AddrTo[StrWithPrefix16]
	EndDate   AddrTo[StrWithPrefix16]
}

type ShopDiscountMetaData struct {
	// Properties
	Discountlist AddrTo[StrWithPrefix16]
}

type ShopGoodsClassifyMetaData struct {
	// Properties
	GoodsClassTitle AddrTo[Hash]
}

type ShopGoodsDisplayDataMetaData struct {
	// Properties
	GoodsID            uint32
	ContentDisplay     AddrTo[Hash]
	ShowMaterialNumber uint32
	DisplayMaterialID  int32
	DisplayMissionID1  int32
	DisplayMissionID2  int32
}

type ShopGoodsJumpMetaData struct {
	// Properties
	LinkType     int32
	LinkParams   AddrTo[[]int32]
	LinkParamStr AddrTo[StrWithPrefix16]
	TextMap      AddrTo[Hash]
}

type ShopGoodsMetaData struct {
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
	GoodsClassIdList  AddrTo[[]uint16]
	GoodsSortId       uint32
	IsAutoOpen        uint8
	BuyMultipleMax    int32
	GlobalMaxBuyTimes int32
	ChooseDialogType  uint8
}

type ShopGoodsPriceRateMetaData struct {
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
	// Properties
	GiftPackLimitDetailTextmap AddrTo[Hash]
	QuotaDescTextmap           AddrTo[Hash]
}

type ShopGoodsScheduleMetaData struct {
	// Properties
	GoodsID   int32
	ValidTime uint32
}

type ShopGoodsTagMetaData struct {
	// Properties
	LabelImage AddrTo[StrWithPrefix16]
	LabelName  AddrTo[StrWithPrefix16]
}

type ShopTabsMetaData struct {
	// Properties
	ShopTabTitle AddrTo[Hash]
	ShopTypeList AddrTo[[]uint32]
	TabIconPath  AddrTo[StrWithPrefix16]
}

type SignInRewardMetaData struct {
	// Properties
	RewardItemID int32
	Display      int32
}

type SimulateRankRewardMetaData struct {
	// Properties
	RewardId int32
}

type SinDemonExBossLevelMetaData struct {
	// Properties
	LightColor    AddrTo[StrWithPrefix16]
	LevelIconPath AddrTo[StrWithPrefix16]
	LevelName     AddrTo[Hash]
}

type SinDemonExConfigMetaData struct {
	// Properties
	LevelLimitMin       int32
	LevelLimitMax       int32
	ChallengeTimes      int32
	RankName            AddrTo[Hash]
	RankIcon            AddrTo[StrWithPrefix16]
	ExbossScoreLine     int32
	ActivityEntranceTag AddrTo[Hash]
	ActivityTagTitle    AddrTo[Hash]
	ActivityDescTitle   AddrTo[Hash]
	ActivityDesc        AddrTo[Hash]
	CanReChallenge      bool
	IsSweepOn           bool
}

type SinDemonExMonsterMetaData struct {
	// Properties
	BossGroupId            int32
	BossName               AddrTo[StrWithPrefix16]
	BossPrefabPath         AddrTo[StrWithPrefix16]
	MonsterId              int32
	HardLevel              int32
	HardLevelGroup         int32
	MonsterHp              int32
	MonsterLevel           uint8
	MonsterBaseScore       int32
	SceneName              AddrTo[StrWithPrefix16]
	BossAttribute          uint8
	BossSkillTipsList      AddrTo[[]int32]
	DefaultShowSkillDetail bool
	BossDesc               AddrTo[Hash]
	ImagePath              AddrTo[StrWithPrefix16]
	RestrictList           AddrTo[[]uint8]
	EventMark              AddrTo[Hash]
	TimesScore             int32
	CornerMarkPath         AddrTo[StrWithPrefix16]
	UpTagList              AddrTo[[]ElfMetaData_Tag]
	DownTagList            AddrTo[[]ElfMetaData_Tag]
	ExtraTimeScore         int32
	ConfigID               int32
}

type SinDemonExMonsterScheduleMetaData struct {
	// Properties
	BossIdList             AddrTo[[]int32]
	IsTimesScore           bool
	BossMonsterWeatherList AddrTo[[]SinDemonExMonsterScheduleMetaData_BossMonsterWeatherData]
	BossScoreRewardList    AddrTo[[]SinDemonExMonsterScheduleMetaData_BossMonsterRewardData]
	BossTurboModeList      AddrTo[[]SinDemonExMonsterScheduleMetaData_BossMonsterTurboData]
	IsActivity             bool
	UpBossIdList           AddrTo[[]int32]
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
	// Properties
	ConfigId     int32
	Rank         int32
	RewardId     int32
	ScheduleId   int32
	DivisionName AddrTo[StrWithPrefix16]
}

type SinDemonExScoreRewardMetaData struct {
	// Properties
	ConfigId   int32
	Score      int32
	RewardId   int32
	ScheduleId int32
	IsEvent    int32
	TimesScore int32
}

type SinDemonExSeasonBossMetaData struct {
	// Properties
	ScheduleIDList    AddrTo[[]int32]
	UpBossList        AddrTo[[]int32]
	NormalBossList    AddrTo[[]int32]
	SeasonPageImgPath AddrTo[StrWithPrefix16]
}

type SinDemonExSkillTipsMetaData struct {
	// Properties
	SkillNameText      AddrTo[Hash]
	SkillNameBriefText AddrTo[Hash]
	SkillTipText       AddrTo[Hash]
	SkillTipBriefText  AddrTo[Hash]
	ShowDetail         bool
	ShowCG             bool
	DetailPicPath      AddrTo[StrWithPrefix16]
	CgID               int32
	BgColor            AddrTo[StrWithPrefix16]
}

type SingleRaidActivityMetaData struct {
	// Properties
	StepStageEnterTimes int32
	RewardMissionList   AddrTo[[]int32]
	RewardTitle         AddrTo[Hash]
	RewardShow          int32
}

type SingleRaidStepStageMetaData struct {
	// Properties
	StepStagePoint int32
	Difficulty     int32
}

type SingleWantedActivityMetaData struct {
	// Properties
	TicketRefreshNumMax       int32
	MpAssistMaterialLimitList AddrTo[[]SingleWantedActivityMetaData_DropItem]
	ResourceMaterialID        uint32
	TicketMaterialID          uint32
	ShopLink                  int32
	StigmataConfigIDList      AddrTo[[]uint32]
}

type SingleWantedActivityMetaData_DropItem struct {
	// Fields
	ItemID int32
	Num    int32
}

type SingleWantedEquipShowConfigMetaData struct {
	// Properties
	EquipmentID        int32
	BeginTime          AddrTo[RemoteTime]
	EndTime            AddrTo[RemoteTime]
	OtherStigmataInfo  AddrTo[SingleWantedEquipShowConfigMetaData_StigmataInfo]
	OtherStigmataInfo1 AddrTo[SingleWantedEquipShowConfigMetaData_StigmataInfo]
	OtherStigmataInfo2 AddrTo[SingleWantedEquipShowConfigMetaData_StigmataInfo]
}

type SingleWantedEquipShowConfigMetaData_StigmataInfo struct {
	// Fields
	Index      int32
	StigmataID int32
}

type SingleWantedMaterialTierMetaData struct {
	// Properties
	TagDisplay AddrTo[StrWithPrefix16]
}

type SingleWantedProgressMetaData struct {
	// Properties
	MPSubTabName         AddrTo[Hash]
	FirstCostMaterialNum int32
	CostMaterialNum      int32
	ExtraRewardList      AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	ExtraDropCostList    AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	MpAssistMaterialList AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	FirstDropDisplayList AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	DropDisplayList      AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	WebLink              AddrTo[StrWithPrefix16]
}

type SingleWantedScheduelMetaData struct {
	// Properties
	MinLevel               int32
	SingleWantedActivityID uint32
}

type SingleWantedStageGroupMetaData struct {
	// Properties
	StageIDList       AddrTo[[]int32]
	MPStageIDList     AddrTo[[]int32]
	StageGroupThemeID uint32
	StageGroupDesc    AddrTo[Hash]
}

type SingleWantedThemeConfigMetaData struct {
	// Properties
	ThemeName          AddrTo[Hash]
	ThemeBriefPicPath  AddrTo[StrWithPrefix16]
	ThemeDetailPicPath AddrTo[StrWithPrefix16]
	MPListDataConfigID int32
}

type SkillTutorialStepData struct {
	// Properties
	TutorialID   int32
	SkillStepStr AddrTo[StrWithPrefix16]
	StepIndex    int32
	Destroy      int32
	DestroyDelay float32
}

type SkillVideoMetaData struct {
	// Properties
	AvatarID        int32
	SkillID         int32
	SubSkillID      int32
	SkillType       int32
	TutorialVideoID int32
	TutoriallevelID int32
	TutorialDesc    AddrTo[Hash]
}

type SLGBattlePointMetaData struct {
	// Properties
	Name                       AddrTo[StrWithPrefix16]
	Desc                       AddrTo[StrWithPrefix16]
	Stage                      int32
	PointLevel                 int32
	IsEmpty                    bool
	CanAttack                  bool
	Building                   int32
	NearPointList              AddrTo[[]int32]
	MinOccupyThreshold         int32
	FirstOccupyLoyalty         int32
	PointLoyalty               int32
	CanQuickPass               bool
	QuickPassCostList          AddrTo[[]ChapterOWBuildingLevelMetaData_CostMaterialData]
	BattleCostList             AddrTo[[]ChapterOWBuildingLevelMetaData_CostMaterialData]
	QuickPassAPCost            int32
	BattleAPCost               int32
	QuickPassReward            int32
	QuickPassScore             int32
	BuildingOverrideName       AddrTo[StrWithPrefix16]
	BuildingOverrideDesc       AddrTo[StrWithPrefix16]
	QuickPassSystemStaminaCost int32
}

type SLGBigBossPointMetaData struct {
	// Properties
	BossStage int32
	Loyalty   int32
	BeginTime AddrTo[RemoteTime]
	EndTime   AddrTo[RemoteTime]
}

type SLGBroadCastMetaData struct {
	// Properties
	ConfigID  int32
	Type      int32
	Parameter AddrTo[StrWithPrefix16]
	Desc      AddrTo[Hash]
	ShowTime  float32
	IsShow    bool
}

type SLGBuffMetaData struct {
	// Properties
	Icon              AddrTo[StrWithPrefix16]
	Desc              AddrTo[Hash]
	BuffRatioList     AddrTo[[]int32]
	BuffRatioDescList AddrTo[[]Hash]
}

type SLGBuildingMetaData struct {
	// Properties
	Type         int32
	Parameter    int32
	Model        AddrTo[StrWithPrefix16]
	SmallIcon    AddrTo[StrWithPrefix16]
	Name         AddrTo[StrWithPrefix16]
	BuildingDesc AddrTo[StrWithPrefix16]
}

type SLGCountrySetMetaData struct {
	// Properties
	Color              AddrTo[StrWithPrefix16]
	Name               AddrTo[Hash]
	PanelPic           AddrTo[StrWithPrefix16]
	BarColor           AddrTo[StrWithPrefix16]
	CharacterSmallIcon AddrTo[StrWithPrefix16]
	CountryModel       AddrTo[StrWithPrefix16]
	ScoreBarColor      AddrTo[StrWithPrefix16]
}

type SLGRankRewardMetaData struct {
	// Properties
	Name   AddrTo[StrWithPrefix16]
	Desc   AddrTo[StrWithPrefix16]
	Reward int32
}

type SLGScheduleMetaData struct {
	// Properties
	MaxAP               int32
	SingleBattleList    AddrTo[[]int32]
	MissionList         AddrTo[[]int32]
	MinLevel            int32
	MaxLevel            int32
	APMaterial          int32
	APRecover           int32
	MainRefreshCD       int32
	StoryList           AddrTo[[]int32]
	Shop                int32
	ShopMaterial        int32
	BuffShopMaterial    int32
	BeginEntryPerformID int32
	EndEntryMission     int32
	EndEntryPerformID   int32
	JsonConfig          AddrTo[StrWithPrefix16]
	PhotoOpenTime       AddrTo[RemoteTime]
	TaleTutorial        AddrTo[[]GlobalWarScheduleMetaData_ConfigTutorial]
	QuickPassTicket     int32
	ShopShowID          int32
}

type SLGScoreRewardMetaData struct {
	// Properties
	ScheduleID int32
	Desc       AddrTo[StrWithPrefix16]
	Type       int32
	Score      int32
	Reward     int32
}

type SLGShopBuffMetaData struct {
	// Properties
	MaterialID    int32
	MaterialNum   int32
	AbilityName   AddrTo[StrWithPrefix16]
	ShopBuffTitle AddrTo[StrWithPrefix16]
	ShopBuffDesc  AddrTo[StrWithPrefix16]
}

type SLGSingleBattleMetaData struct {
	// Properties
	MapID               uint16
	BeginTime           AddrTo[RemoteTime]
	EndTime             AddrTo[RemoteTime]
	BattleStartTime     AddrTo[RemoteTime]
	OpenTime            AddrTo[RemoteTimeSpan]
	CloseTime           AddrTo[RemoteTimeSpan]
	BattleReportCD      int32
	CountryList         AddrTo[[]SLGSingleBattleMetaData_CountryPoint]
	CenterPoint         int32
	PointScoreAccountCD int32
	BattleFog           bool
	FogDistance         int32
	BattlePreviewPic    AddrTo[StrWithPrefix16]
}

type SLGSingleBattleMetaData_CountryPoint struct {
	// Fields
	CountryID int32
	PointID   int32
}

type SLGSmallBossPointMetaData struct {
	// Properties
	BossStage          int32
	BossHP             int32
	FirstOccupyLoyalty int32
}

type SLGStageMetaData struct {
	// Properties
	Reward AddrTo[[]SLGStageMetaData_ScoreRewardItem]
}

type SLGStageMetaData_ScoreRewardItem struct {
	// Fields
	RewardID int32
	Score    int32
}

type SLGStoryMetaData struct {
	// Properties
	PhotoID         int32
	UnlockMissionID int32
	BeginTime       AddrTo[RemoteTime]
	EndTime         AddrTo[RemoteTime]
	IsShow          bool
	Desc            AddrTo[Hash]
}

type SLGVoteMetaData struct {
	// Properties
	Pic                AddrTo[StrWithPrefix16]
	StoryID            AddrTo[[]int32]
	DanmakuSlotID      int32
	DanmakuColor       AddrTo[StrWithPrefix16]
	Desc               AddrTo[StrWithPrefix16]
	CharacterSmallIcon AddrTo[StrWithPrefix16]
	Name               AddrTo[StrWithPrefix16]
}

type SLGVoteScheduleMetaData struct {
	// Properties
	VoteIDList    AddrTo[[]int32]
	DanmakuID     int32
	VoteMaterial  int32
	VoteBeginTime AddrTo[RemoteTime]
	VoteEndTime   AddrTo[RemoteTime]
}

type SlotMachineBoxitemMetaData struct {
	// Properties
	RewardID          int32
	CombinationIDList AddrTo[[]int32]
}

type SlotMachineCombinationMetaData struct {
	// Properties
	Combination        AddrTo[[]SlotMachineCombinationMetaData_CombElement]
	CombinationNumList AddrTo[[]SlotMachineCombinationMetaData_CombTier]
	EffectLevel        int32
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
	// Properties
	IconImgPath AddrTo[StrWithPrefix16]
}

type SlotMachineProgressRewardMetaData struct {
	// Properties
	Progress int32
	RewardID int32
}

type SlotMachineScheduleMetaData struct {
	// Properties
	NormalMaterialID  int32
	NormalMaterialNum int32
	IconIDList        AddrTo[[]int32]
	SuperIconIDList   AddrTo[[]int32]
}

type SpareShareConfigMetaData struct {
	// Properties
	DeviceType     uint8
	ShareChannelID AddrTo[StrWithPrefix16]
	Type           uint8
}

type SpecialAffixDataMetaData struct {
	// Properties
	AffixTreeNodeID  int32
	NodeNumRequire   int32
	TriggerMaterial  int32
	RefineCost       AddrTo[[]AffixTreeNodeMetaData_MaterialNeed]
	PreviewDesc      AddrTo[Hash]
	RefineButtonDesc AddrTo[Hash]
}

type SpecialConfigMetaData struct {
	// Properties
	ReplaceType  int32
	ThemeResPath AddrTo[StrWithPrefix16]
}

type StageDamageStatisticsMetaData struct {
	// Properties
	ShowElf bool
}

type StageDetailAbilityMetaData struct {
	// Properties
	DisplayTitle  AddrTo[Hash]
	Threat        int32
	SkillTypeList AddrTo[[]int32]
	SkillInfo     AddrTo[Hash]
	DisplayGuide  AddrTo[Hash]
}

type StageDetailAccumuatorGaugeMetaData struct {
	// Properties
	Desc  AddrTo[Hash]
	TagID uint16
}

type StageDetailEffectMetaData struct {
	// Properties
	EffectText AddrTo[Hash]
	EffectIcon AddrTo[StrWithPrefix16]
}

type StageDetailLayoutMetaData struct {
	// Properties
	GroupDisplaySequence1 AddrTo[[]int32]
	GroupDisplaySequence2 AddrTo[[]int32]
}

type StageDetailMetaData struct {
	// Properties
	DisplayMission         AddrTo[Hash]
	EffectList             AddrTo[[]int32]
	EnemyList              AddrTo[[]int32]
	DisplayGuide           AddrTo[Hash]
	RecommendAvatar        AddrTo[[]int32]
	SummerRanchRecommended AddrTo[[]int32]
	UpTagList              AddrTo[[]ElfMetaData_Tag]
	DownTagList            AddrTo[[]ElfMetaData_Tag]
	StageTagCheckType      uint8
	StageTagCheckNum       uint8
	StageDetailLayoutID    int32
	StageRecommendType     uint8
}

type StageDetailMonsterMetaData struct {
	// Properties
	Nature            int32
	DisplayName       AddrTo[Hash]
	DisplayDescID     int32
	BossType          int32
	AbilitiesList     AddrTo[[]int32]
	IconPath          AddrTo[StrWithPrefix16]
	PrefabPath        AddrTo[StrWithPrefix16]
	ControllerPath    AddrTo[StrWithPrefix16]
	ModelScale        float32
	ModelPosition     AddrTo[[]float32]
	ModelRotation     AddrTo[[]float32]
	ResistList        AddrTo[[]int32]
	AccumuatorGaugeID uint16
	RecommendTag      AddrTo[[]uint16]
	UnRecommendTag    AddrTo[[]uint16]
	PageLayout        AddrTo[[]uint8]
}

type StageDetailMonsterResistMetaData struct {
	// Properties
	IconPath      AddrTo[StrWithPrefix16]
	ResistName    AddrTo[Hash]
	StarNum       uint8
	PositiveState uint8
}

type StageDetailNatureMetaData struct {
	// Properties
	IconPath AddrTo[StrWithPrefix16]
}

type StageDetailSkillTypeMetaData struct {
	// Properties
	DisplayName AddrTo[Hash]
	IconPath    AddrTo[StrWithPrefix16]
}

type StageDialogMetaData struct {
	// Properties
	Name     AddrTo[Hash]
	Content  AddrTo[Hash]
	AudioID  AddrTo[StrWithPrefix16]
	AvatarID int32
	Emotion  AddrTo[StrWithPrefix16]
	Position int32
	Time     float32
}

type StageDifficultySelectMetaData struct {
	// Properties
	DifficultyList AddrTo[[]uint8]
}

type StageDropItemDataMetaData struct {
	// Properties
	ItemID  int32
	ItemNum int32
}

type StageEnhanceMetaData struct {
	// Properties
	Leveltype       int32
	TypePara        AddrTo[StrWithPrefix16]
	AvatarList      AddrTo[StrWithPrefix16]
	AvatarBuffType  int32
	AvatarBuffRatio int32
	HpRatio         AddrTo[StrWithPrefix16]
	AtkRatio        AddrTo[StrWithPrefix16]
	DefRatio        AddrTo[StrWithPrefix16]
	HardLevelPara   AddrTo[StrWithPrefix16]
	CountDown       int32
}

type StageLiveRecommendMetaData struct {
	// Properties
	RecommandTypeParam AddrTo[[]int32]
	MaxRecNum          int32
	MaxRecDetailNum    int32
	PriorityShowType   int32
}

type StageMaxScoreMetaData struct {
	// Properties
	MaxScore int32
}

type StageMessageMetaData struct {
	// Properties
	Value int32
}

type StageMultiTeamMetaData struct {
	// Properties
	A_Position     AddrTo[[]int32]
	A_MonsterTypes AddrTo[[]int32]
	B_Position     AddrTo[[]int32]
	B_MonsterTypes AddrTo[[]int32]
	C_Position     AddrTo[[]int32]
	C_MonsterTypes AddrTo[[]int32]
}

type StageNeedMuteChallengePopupMetaData struct {
	// Properties
	NeedMuteChallengePopup bool
}

type StagePhotoMetaData struct {
	// Properties
	PhotoID int32
}

type StageQAvatarMetaData struct {
	// Properties
	BossHP              int32
	RhythmLevel         int32
	RecommendQAvatarNum int32
}

type StageRandomEffectMetaData struct {
	// Properties
	RandomEffectType int32
	IsCarryOn        bool
}

type StageRankConfigMetaData struct {
	// Properties
	MinLevel uint32
	MaxLevel uint32
}

type StageRankGroupConfigMetaData struct {
	// Properties
	RankIDList AddrTo[[]uint32]
}

type StageRechallengeMetaData struct {
	// Properties
	Rechallenge int32
}

type StageResistanceMetaData struct {
	// Properties
	ResistanceNeed float32
	LevelType      uint8
	AvatarType     uint8
}

type StageScoreRewardDisplayMetaData struct {
	// Properties
	RewardMaterial int32
	BestTime       uint16
	MaxScore       int32
}

type StageScoreRewardMetaData struct {
	// Properties
	RewardMaterialList  AddrTo[AffixTreeNodeMetaData_MaterialNeed]
	LevelProgressReward uint16
	RemainedTimeReward  uint16
}

type StageStoryMapMetaData struct {
	// Properties
	QuestID AddrTo[[]int32]
}

type StageStoryMetaData struct {
	// Properties
	Name           AddrTo[Hash]
	Sub            int32
	Description    AddrTo[Hash]
	Target         AddrTo[Hash]
	MaxCount       int32
	LuaPath        AddrTo[StrWithPrefix16]
	TargetLocation AddrTo[StrWithPrefix16]
}

type StageTypeMetaData struct {
	// Properties
	StageDetailLayoutID   int32
	ApplicationFocusPause uint8
}

type StepMissionCompensateMetaData struct {
	// Properties
	UnlockLevel            uint8
	MainLineStepIDList     AddrTo[[]int32]
	NewChallengeStepIDList AddrTo[[]int32]
	OldChallengeStepIDList AddrTo[[]int32]
}

type StigmataAffixMetaData struct {
	// Properties
	NameMono   AddrTo[Hash]
	NameDual   AddrTo[Hash]
	PropID     int32
	PropParam1 float32
	PropParam2 float32
	PropParam3 float32
	UIType     int32
	UIValue    float32
	UINature   int32
	UIClass    int32
}

type StigmataBriefMetaData struct {
	// Properties
	SetID      int32
	Height     AddrTo[Hash]
	Weight     AddrTo[Hash]
	HomePlace  AddrTo[Hash]
	ActiveSite AddrTo[Hash]
	Property   AddrTo[Hash]
	Sector     AddrTo[Hash]
	Story01    AddrTo[Hash]
	Story02    AddrTo[Hash]
	Story03    AddrTo[Hash]
	Story04    AddrTo[Hash]
	Story05    AddrTo[Hash]
}

type StigmataMemoirMetaData struct {
	// Properties
	SetID             int32
	OnePieceSet       int32
	Unlock            int32
	UiOrder           int32
	ThemeGroup        int32
	StigmataQNameText AddrTo[Hash]
	StigmataQIcon     AddrTo[StrWithPrefix16]
	MaxProcess        int32
	RewardGroupList   AddrTo[[]int32]
}

type StigmataPositionMetaData struct {
	// Properties
	RootWidth      float32
	RootHeight     float32
	BackEnable     bool
	BackPosX       float32
	BackPosY       float32
	BackWidth      float32
	BackHeight     float32
	BackPath       AddrTo[StrWithPrefix16]
	ImageWidth     float32
	ImageHeight    float32
	ImagePath      AddrTo[StrWithPrefix16]
	FrontEnable    bool
	FrontPosX      float32
	FrontPosY      float32
	FrontWidth     float32
	FrontHeight    float32
	FrontRotationX float32
	FrontRotationY float32
	FrontRotationZ float32
	FrontPath      AddrTo[StrWithPrefix16]
}

type StigmataRuneAffixMetaData struct {
	// Properties
	Level          int32
	PropID         int32
	ValueMin       float32
	ValueMax       float32
	ValueStep      float32
	BaseValue      float32
	AddPerPercent  float32
	NameMono       AddrTo[Hash]
	NameDual       AddrTo[Hash]
	PropParam2     float32
	PropParam3     float32
	UIType         int32
	UINature       int32
	UIClass        AddrTo[[]int32]
	AffixSkillType int32
}

type StigmataTagMetaData struct {
	// Properties
	SpecificAvatarBonus        AddrTo[[]StigmataTagMetaData_RatioItem]
	SpecificAvatarBonusSupport AddrTo[[]StigmataTagMetaData_RatioItem]
}

type StigmataTagMetaData_RatioItem struct {
	// Fields
	Pred  int32
	Ratio float32
}

type StigmataThemeMetaData struct {
	// Properties
	ThemeNameText AddrTo[Hash]
}

type StigmataVirtualSetMetaData struct {
	// Properties
	StigmataMainID   AddrTo[[]int32]
	SetModify        AddrTo[[]StigmataTagMetaData_RatioItem]
	SetModifySupport AddrTo[[]StigmataTagMetaData_RatioItem]
}

type StorySweepGroupInfo struct {
	// Properties
	Material  int32
	SweepName AddrTo[Hash]
	SweepIcon AddrTo[StrWithPrefix16]
}

type StorySweepMetaData struct {
	// Properties
	StageID      int32
	GroupType    int32
	GroupID      int32
	TimeCost     int32
	RestrictList AddrTo[[]int32]
	SelectOrder  int32
}

type SubMallTypeDataMetaData struct {
	// Properties
	FatherID                  int32
	Name                      AddrTo[Hash]
	SortID                    int32
	AddOnEventMissionList     AddrTo[[]int32]
	AddOnEventMissionIconPath AddrTo[StrWithPrefix16]
	BeginTime                 AddrTo[StrWithPrefix16]
	EndTime                   AddrTo[StrWithPrefix16]
}

type SubPakDataMetaData struct {
	// Properties
	DownloadType uint8
	GroupID      int32
}

type SubPakGroupMetaData struct {
	// Properties
	GroupName        AddrTo[StrWithPrefix16]
	GroupDesc        AddrTo[StrWithPrefix16]
	DownloadPriority float32
}

type SubsBenefitDataMetaData struct {
	// Properties
	RewardID int32
}

type SubsDataMetaData struct {
	// Properties
	EntryName   AddrTo[StrWithPrefix16]
	EntryImg    AddrTo[StrWithPrefix16]
	EntryDesc   AddrTo[StrWithPrefix16]
	DetailTitle AddrTo[StrWithPrefix16]
	DetailImg   AddrTo[StrWithPrefix16]
	BenefitList AddrTo[[]int32]
}

type SubsidiaryEventMetaData struct {
	// Properties
	AvatarOffset           AddrTo[ElfMetaData_VectorFloat3]
	Path                   AddrTo[[]SubsidiaryEventMetaData_DressPath]
	SubsidiaryMotion       AddrTo[SubsidiaryEventMetaData_OverrideClip]
	Face                   AddrTo[StrWithPrefix16]
	FaceDelayTime          float32
	AnimatorControllerPath AddrTo[StrWithPrefix16]
}

type SubsidiaryEventMetaData_OverrideClip struct {
	// Objects
	ClipPath StrWithPrefix16
	Name     StrWithPrefix16
}

type SubsidiaryEventMetaData_DressPath struct {
	// Fields
	DressID int32

	// Objects
	Path StrWithPrefix16
}

type SubsWareDataMetaData struct {
	// Properties
	SubsID  int32
	SubsDur int32
	Sort    int32
	Promo   AddrTo[StrWithPrefix16]
	Renewal AddrTo[StrWithPrefix16]
	BGImg   AddrTo[StrWithPrefix16]
	Name    AddrTo[StrWithPrefix16]
	Desc    AddrTo[StrWithPrefix16]
}

type SupportActivityScheduleMetaData struct {
	// Properties
	MinLevel                    int32
	MaxLevel                    int32
	EndTime                     AddrTo[StrWithPrefix16]
	SupportConfigID             int32
	SupportMissionConfigID      int32
	GlobalSupportRewardConfigID int32
	WebLink                     AddrTo[StrWithPrefix16]
	LinkShopGoodsID             int32
	AvatarID                    int32
}

type SupportAvatarInitialMetaData struct {
	// Properties
	InitialSupportAvatarIDs AddrTo[[]uint16]
}

type SupportAvatarLevelMetaData struct {
	// Properties
	BuffID            uint32
	SkillFunctionName AddrTo[StrWithPrefix16]
	AbilityName       AddrTo[StrWithPrefix16]
	ParamList         AddrTo[[]StrWithPrefix16]
	CD                int32
	SkillIcon         AddrTo[StrWithPrefix16]
	SkillName         AddrTo[Hash]
	SkillDesc         AddrTo[Hash]
}

type SupportAvatarMetaData struct {
	// Properties
	AvatarID       uint16
	DefaultDressID int32
	ActivityType   uint16
	ActivityID     uint32
	SkillMaxLevel  uint16
	MaterialID     uint32
}

type SupportConfigMetaData struct {
	// Properties
	BeginTime                        AddrTo[RemoteTime]
	EndTime                          AddrTo[RemoteTime]
	CountDownTime                    AddrTo[RemoteTime]
	AvatarSupportPrefabPath          AddrTo[StrWithPrefix16]
	AvatarSupportSharePagePrefabPath AddrTo[StrWithPrefix16]
	AvatarSupportMissionPrefabPath   AddrTo[StrWithPrefix16]
	MaxShowNum                       int32
	AvatarSupportContextType         int32
	UiFormat                         uint8
	ShareRewardID                    int32
}

type SupportMissionConfigMetaData struct {
	// Properties
	MissionID       int32
	TitleTextMap    AddrTo[Hash]
	StoryTextMap    AddrTo[Hash]
	CgVideoID       int32
	ButtonDisappear bool
	StarIntro       AddrTo[StrWithPrefix16]
}

type SurvivalWeaponShowMetaData struct {
	// Properties
	IconPath             AddrTo[StrWithPrefix16]
	WeaponAssetPath      AddrTo[StrWithPrefix16]
	WeaponAttachRoot     AddrTo[StrWithPrefix16]
	WeaponAttachPosition AddrTo[[]float32]
	WeaponAttachRotation AddrTo[[]float32]
	WeaponAnimation      AddrTo[StrWithPrefix16]
}

type SusannaTrialAbilityPicMetaData struct {
	// Properties
	Entry       int32
	SiteName    AddrTo[Hash]
	RpgBuffList AddrTo[[]SusannaTrialAbilityPicMetaData_RelatedRpgBuff]
	MaxNum      int32
}

type SusannaTrialAbilityPicMetaData_RelatedRpgBuff struct {
	// Fields
	Product    int32
	RpgSkillID int32
}

type SusannaTrialActivityMetaData struct {
	// Properties
	AreaList                    AddrTo[[]int32]
	SkillMaterialList           AddrTo[[]int32]
	TrainingPageDefaultCurrList AddrTo[[]int32]
}

type SusannaTrialAreaMetaData struct {
	// Properties
	ProgressMaterialId int32
	HiddenStageNum     uint8
	AreaTopCurrency    AddrTo[[]int32]
	AreaOutput         int32
	MissionSkillID     AddrTo[[]int32]
	SkillUISIte        AddrTo[StrWithPrefix16]
	TutorIcon          AddrTo[StrWithPrefix16]
	HardLevelHint      AddrTo[[]SusannaTrialAreaMetaData_HardLevelSkill]
	HintTextMap        AddrTo[Hash]
	PlotID             int32
	TimeLimit          AddrTo[RemoteTime]
	ProcessLimit       AddrTo[[]SusannaTrialAreaMetaData_ProgressLimit]
	TimeLimitHint      AddrTo[Hash]
	ProcessLimitHint   AddrTo[Hash]
}

type SusannaTrialAreaMetaData_HardLevelSkill struct {
	// Fields
	Level      int32
	RpgSkillID int32
}

type SusannaTrialAreaMetaData_ProgressLimit struct {
	// Fields
	AreaID int32
	Number int32
}

type SusannaTrialMainStageTimeMetaData struct {
	// Properties
	NeedShowTime      bool
	Time              AddrTo[RemoteTime]
	TextMapID         AddrTo[Hash]
	NeedSpHInt        bool
	NeedSpHIntTextMap AddrTo[Hash]
}

type SusannaTrialOverallMetaData struct {
	// Properties
	OverallID int32
	Value     int32
}

type SusannaTrialPhotoDescMetaData struct {
	// Properties
	TextMapID AddrTo[Hash]
}

type SusannaTrialPlotMetaData struct {
	// Properties
	SiteID     int32
	SiteDesc   AddrTo[Hash]
	OptionList AddrTo[[]int32]
	CostItemID AddrTo[[]int32]
	RewardID   int32
}

type SusannaTrialPopTextMetaData struct {
	// Properties
	Type             uint8
	IsHeadIconShow   bool
	Path             AddrTo[StrWithPrefix16]
	Content          AddrTo[StrWithPrefix16]
	IsHeartImageShow bool
}

type SusannaTrialProgressLineMetaData struct {
	// Properties
	AreaID           int32
	RequiredProgress uint8
	RewardID         int32
	TokenNum         int32
	PlotID           int32
	IsPlotReward     bool
}

type SusannaTrialSiteTypeMetaData struct {
	// Properties
	IsAlwaysShow  bool
	IsOnlyOneTime bool
}

type SusannaTrialSkillSIteMetaData struct {
	// Properties
	UISiteID uint8
}

type SusannaTrialStageMetaData struct {
	// Properties
	AreaID        uint32
	StageType     uint8
	AreaMapName   AddrTo[StrWithPrefix16]
	StageBGM      AddrTo[StrWithPrefix16]
	SpawnPoint    AddrTo[StrWithPrefix16]
	MonsterPoint  AddrTo[[]StrWithPrefix16]
	MonsterID     AddrTo[[]int32]
	MonsterNumber AddrTo[[]int32]
	HardLevel     int32
	SkillList     AddrTo[[]int32]
	QAvatarID     int32
	CGID          int32
	TextID        AddrTo[Hash]
}

type TeamAssaultBossMetaData struct {
	// Properties
	BaseScore       int32
	MaxScore        int32
	HardLevel       int32
	HpRatio         int32
	ScoreRatio      int32
	RestrictList    AddrTo[[]int32]
	TimeLimit       int32
	MonsterDetailID int32
	HardShow        int32
	BossNameText    AddrTo[Hash]
}

type TeamAssaultMetaData struct {
	// Properties
	ActivityEnterTimes int32
	BossList           AddrTo[[]int32]
	StageID            int32
	UiIconPath         AddrTo[StrWithPrefix16]
	UiBgPath           AddrTo[StrWithPrefix16]
}

type TeamBuffMetaData struct {
	// Properties
	Name         AddrTo[Hash]
	Type         int32
	Desc         AddrTo[Hash]
	Model        AddrTo[StrWithPrefix16]
	IconPath     AddrTo[StrWithPrefix16]
	MaxSuperpose int32
	ParamBase_1  float32
	ParamAdd_1   float32
	ParamBase_2  float32
	ParamAdd_2   float32
	ParamBase_3  float32
	ParamAdd_3   float32
}

type TextIDReplaceMetaData struct {
	// Properties
	ReplaceTextID  AddrTo[Hash]
	ValidMissionID int32
}

type TextMapMetaData struct {
	// Properties
	Text AddrTo[StrWithPrefix16]
}

type ThemeActivityMetaData struct {
	// Properties
	DailyUpCount       int32
	ActivityBonusScore int32
	ActivityBonusRatio float32
	ActivityImg        AddrTo[StrWithPrefix16]
	ActivityTitle      AddrTo[Hash]
	ActivityDes        AddrTo[Hash]
	BuffList           AddrTo[[]int32]
	ActivityTips       AddrTo[[]Hash]
}

type ThemeAvatarMetaData struct {
	// Properties
	AvatarIDList          AddrTo[[]int32]
	ClientAvatarBonusType int32
	AvatarBonusScore      int32
	AvatarBonusRatio      float32
	AvatarTitle           AddrTo[Hash]
	AvatarDes             AddrTo[Hash]
	AvatarDisplayList     AddrTo[[]int32]
	BuffList              AddrTo[[]int32]
	AvatarTips            AddrTo[[]Hash]
}

type ThemeBonusDataMetaData struct {
	// Properties
	BonusThemeType      int32
	DisplayDroplist     AddrTo[[]SingleWantedActivityMetaData_DropItem]
	DisplayDropHideType int32
	BonusThemeMax       int32
	BonusThemeSingleMax int32
	BonusList           AddrTo[[]ThemeBonusDataMetaData_BonusItem]
	BonusDescTitle      AddrTo[Hash]
	BonusDescText       AddrTo[Hash]
	BonusTips           AddrTo[[]Hash]
}

type ThemeBonusDataMetaData_BonusItem struct {
	// Fields
	ID    int32
	Type  int32
	Value int32
}

type ThemeBuffData struct {
	// Properties
	BuffType      int32
	BuffNameStrId AddrTo[StrWithPrefix16]
	BuffNameText  AddrTo[Hash]
	BuffDescStrId AddrTo[Hash]
	BuffIcon      AddrTo[StrWithPrefix16]
}

type ThemeBuffMetaData struct {
	// Properties
	IsScoreBuff      int32
	BuffName         AddrTo[Hash]
	BuffDes          AddrTo[Hash]
	BuffIcon         AddrTo[StrWithPrefix16]
	AbilityName      AddrTo[StrWithPrefix16]
	ParamEmptyString AddrTo[StrWithPrefix16]
	ParamString      AddrTo[StrWithPrefix16]
	ParamBase_1      float32
	ParamAdd_1       float32
	ParamBase_2      float32
	ParamAdd_2       float32
	ParamBase_3      float32
	ParamAdd_3       float32
	InLevelType      int32
}

type ThemeEquipMetaData struct {
	// Properties
	WeaponMainIDList   AddrTo[[]int32]
	StigmataMainIDList AddrTo[[]int32]
	EquipTitle         AddrTo[Hash]
	EquipDes           AddrTo[Hash]
	EquipDisplayList   AddrTo[[]int32]
	BuffList           AddrTo[[]int32]
	EquipTips          AddrTo[[]Hash]
}

type ThemeEventImgMetaData struct {
	// Properties
	CanSwitch bool
	BeginTime AddrTo[RemoteTime]
	EndTime   AddrTo[RemoteTime]
	ImgPath   AddrTo[[]StrWithPrefix16]
	GoToLink  AddrTo[[]StrWithPrefix16]
}

type ThemeEventMetaData struct {
	// Properties
	BGPrefabPath     AddrTo[StrWithPrefix16]
	UIPrefabPath     AddrTo[StrWithPrefix16]
	AvatarPrefabPath AddrTo[StrWithPrefix16]
	UseNewPanel      bool
	MissionIDList    AddrTo[[]int32]
}

type ThemeScheduleMetaData struct {
	// Properties
	ThemeType     int32
	ThemeTypePara AddrTo[[]int32]
	BeginTime     AddrTo[StrWithPrefix16]
	EndTime       AddrTo[StrWithPrefix16]
	BonusThemeID  int32
	SeasonData    int32
	ActivityData  int32
	AvatarData    int32
	EquipData     int32
}

type ThemeSeasonMetaData struct {
	// Properties
	ThemeStageList AddrTo[[]int32]
	ThemeTitle     AddrTo[Hash]
	ThemeDesc      AddrTo[Hash]
	ThemeBGPath    AddrTo[StrWithPrefix16]
}

type ThemeSpecialConfigMetaData struct {
	// Properties
	ReplaceType    uint8
	DefaultResPath AddrTo[StrWithPrefix16]
	ThemeResPath   AddrTo[StrWithPrefix16]
}

type ThemeStageMetaData struct {
	// Properties
	ThemeStageName   AddrTo[Hash]
	ThemeStageBGPath AddrTo[StrWithPrefix16]
	ThemeStageDesc   AddrTo[Hash]
}

type ThemeStigmataGroup struct {
	// Properties
	StigmataGroupName     AddrTo[Hash]
	StigmataGroupShowList AddrTo[[]int32]
	StigmataIDList        AddrTo[[]int32]
}

type ThemeTowerScheduleMetaData struct {
	// Properties
	MinPlayerLevel      uint8
	BeginTime           AddrTo[RemoteTime]
	EndTime             AddrTo[RemoteTime]
	Title               AddrTo[Hash]
	SectionIDList       AddrTo[[]int32]
	MissionIDList       AddrTo[[]int32]
	ShowMaterialIDList  AddrTo[[]int32]
	PictuorialID        int32
	RewardLineID        int32
	IntroductionWebLink AddrTo[StrWithPrefix16]
}

type ThemeTowerScoreDropMetaData struct {
	// Properties
	DropMaterialID  int32
	DropMaterialNum int32
}

type ThemeTowerSectionLevelMetaData struct {
	// Properties
	CostMaterialNum int32
}

type ThemeTowerSectionMetaData struct {
	// Properties
	TowerSectionTitle        AddrTo[Hash]
	OpenTime                 AddrTo[RemoteTime]
	TowerSectionPicPath      AddrTo[StrWithPrefix16]
	SectionLevelConfigID     uint8
	TowerStageIDList         AddrTo[[]int32]
	ShowMaterialIDList       AddrTo[[]int32]
	SectionLevelUpMaterialID int32
	StageResource            AddrTo[StrWithPrefix16]
	MaxDropMaterialNum       int32
}

type ThemeTowerSignMetaData struct {
	// Properties
	SectionSignID      int32
	SectionStageSignID int32
}

type ThemeTowerStageBuffMetaData struct {
	// Properties
	Type        uint8
	Rarity      uint8
	IconPath    AddrTo[StrWithPrefix16]
	Name        AddrTo[Hash]
	Desc        AddrTo[Hash]
	LevelUPDesc AddrTo[[]Hash]
	AbilityName AddrTo[StrWithPrefix16]
	ParamList   AddrTo[[]StrWithPrefix16]
}

type ThemeTowerStageFloorMetaData struct {
}

type ThemeTowerStageMetaData struct {
	// Properties
	StepConfigID            int32
	FloorConfigID           int32
	FirstWinRewardID        int32
	TowerStageIconPath      AddrTo[StrWithPrefix16]
	TowerStageBossID        int32
	BossModelPosition       AddrTo[[]float32]
	BossModelRotation       AddrTo[[]float32]
	BossModelScale          float32
	WeatherBuffID           int32
	SignID                  int32
	SpecialBuffIDList       AddrTo[[]int32]
	CommonBuffIDList        AddrTo[[]int32]
	StageBuffRandom         AddrTo[[]ThemeTowerStageMetaData_RandomWeight]
	UniqueStageBuffMinTimes int32
	ActiveSkillBuffID       int32
	MonsterIDList           AddrTo[[]int32]
	PictuorialID            int32
	IntroductionWebLink     AddrTo[StrWithPrefix16]
}

type ThemeTowerStageMetaData_RandomWeight struct {
	// Fields
	BuffID int32
	Weight int32
}

type ThemeTowerStageStepMetaData struct {
	// Properties
	ScoreDropConfigID      int32
	ScoreRatio             float32
	HardLevel              int32
	SelectTitle            AddrTo[Hash]
	RecommendrSectionLevel uint8
	PredictRewardRange     AddrTo[ShigureKiraActivityMetaData_AbilityIDRange]
	PredictScoreRange      AddrTo[ShigureKiraActivityMetaData_AbilityIDRange]
}

type ThemeUIConfigMetaData struct {
	// Properties
	MainPage          AddrTo[StrWithPrefix16]
	PlayerStatusPanel AddrTo[StrWithPrefix16]
	ChatDialogs       AddrTo[StrWithPrefix16]
}

type ThemeWeatherAtmosphereMetaData struct {
	// Properties
	Path AddrTo[StrWithPrefix16]
}

type ThemeWeatherConfigMetaData struct {
	// Properties
	AtmosphereConfigID    uint8
	ReplacePrefabConfigID uint8
	DisplayName           AddrTo[Hash]
	IconPath              AddrTo[StrWithPrefix16]
}

type ThemeWeatherGroupConfigMetaData struct {
	// Properties
	WeatherIDList        AddrTo[[]uint8]
	AllLevelsClearedRate AddrTo[[]int32]
	ChooseRate           AddrTo[[]int32]
}

type ThemeWeatherReplacePrefabMetaData struct {
	// Properties
	ReplacePrefabPath        AddrTo[StrWithPrefix16]
	TimeCondition            AddrTo[[]ThemeWeatherReplacePrefabMetaData_DayTimeInterval]
	RealTimeWeatherCondition AddrTo[[]uint8]
}

type ThemeWeatherReplacePrefabMetaData_DayTimeInterval struct {
	// Fields
	EndDayTime   uint8
	StartDayTime uint8
}

type ThreadCollectionFileDataMetaData struct {
	// Properties
	PhotoId            uint32
	HidePhotoID        uint32
	HideWhenLocked     bool
	HideInArchive      bool
	CollectionSiteType int32
	Name               AddrTo[StrWithPrefix16]
	Desc               AddrTo[StrWithPrefix16]
	PlotID             int32
	LuaPath            AddrTo[StrWithPrefix16]
}

type TileEntityDataMetaData struct {
	// Properties
	ModelID          int32
	JsonFile         AddrTo[StrWithPrefix16]
	MonopolyJsonFile AddrTo[StrWithPrefix16]
}

type TileMapInfoMetaData struct {
	// Properties
	BgMod           AddrTo[StrWithPrefix16]
	AudioBank       AddrTo[[]StrWithPrefix16]
	FloorNameText   AddrTo[Hash]
	FloorPanelTitle AddrTo[Hash]
	FloorPanelDesc  AddrTo[Hash]
	AudioNameEnter  AddrTo[StrWithPrefix16]
	AudioNameExit   AddrTo[StrWithPrefix16]
}

type TileMapMetaData struct {
	// Properties
	LuaFile            AddrTo[StrWithPrefix16]
	ValueTriggerIDList AddrTo[[]int32]
	Entrance           int32
	ForwardFork        AddrTo[StrWithPrefix16]
}

type TileModelMetaData struct {
	// Properties
	ModelPath  AddrTo[StrWithPrefix16]
	ModelScale AddrTo[[]float32]
	ModelType  uint8
}

type TileTerrainDataMetaData struct {
	// Properties
	OnEnterValueModIDList  AddrTo[[]int32]
	OnExitValueModIDList   AddrTo[[]int32]
	OnPassbyValueModIDList AddrTo[[]int32]
}

type TileValueKeyMetaData struct {
	// Properties
	RpgID         int32
	RpgOverallKey int32
}

type TileValueModifierMetaData struct {
	// Properties
	Key                 int32
	Operation           uint16
	Param               int32
	EnablePlayerTrigger bool
	EntityIDList        AddrTo[[]int32]
}

type TileValueTriggerMetaData struct {
	// Properties
	Key                      int32
	ConditionType            uint16
	Param                    int32
	AutoDisablePlayerControl bool
}

type TimePocketMoneyDataMetaData struct {
	// Properties
	RewardId  int32
	TotalTime int32
}

type TouchBuffMetaData struct {
	// Properties
	Effect     AddrTo[StrWithPrefix16]
	BuffDetail AddrTo[Hash]
	Param1     float32
	Param2     float32
	Param3     float32
	Param1Add  float32
	Param2Add  float32
	Param3Add  float32
}

type TouchLevelMetaData struct {
	// Properties
	TouchExp             int32
	Prop                 float32
	Rate                 float32
	BattleGain           int32
	AddGoodfeelPerMinute int32
}

type TouchMetaData struct {
	// Properties
	Point int32
	Buff  AddrTo[[]int32]
}

type TowerBuffConfigMetaData struct {
	// Properties
	TowerBuffIcon AddrTo[StrWithPrefix16]
	TowerBuffDes  AddrTo[Hash]
}

type TowerGrowBuffConfigMetaData struct {
	// Properties
	HP             int32
	ATK            int32
	DEF            int32
	AllDamageRatio int32
}

type TowerRaidBuffMetaData struct {
	// Properties
	BuffIDGroup int32
	BuffTag     AddrTo[StrWithPrefix16]
	StagePos    int32
	Textmap     AddrTo[Hash]
	BuffIcon    AddrTo[StrWithPrefix16]
}

type TowerRaidConfigMetaData struct {
	// Properties
	MinDifficulty    int32
	MaxDifficulty    int32
	TowerRaidGroupID int32
}

type TowerRaidDifficultyMetaData struct {
	// Properties
	AddHardLevel           int32
	AvatarPlayTimes        int32
	UpAvatarPlayTimes      int32
	BuffSwitch             bool
	Reward_1               int32
	Reward_2               int32
	Reward_3               int32
	FinalDisplayDropList   AddrTo[StrWithPrefix16]
	DifficultyDesc         AddrTo[Hash]
	DifficultyName         AddrTo[Hash]
	RewardDesc_1           AddrTo[Hash]
	RewardDesc_2           AddrTo[Hash]
	RewardDesc_3           AddrTo[Hash]
	StageDropDisplayList_1 AddrTo[StrWithPrefix16]
}

type TowerRaidScheduleMetaData struct {
	// Properties
	EndTime        AddrTo[StrWithPrefix16]
	ThemeAvatar    int32
	BuffGroupID    AddrTo[[]int32]
	DailyPlayTimes int32
}

type TowerRaidStageMetaData struct {
	// Properties
	StageIDList          AddrTo[[]int32]
	StageNameList        AddrTo[[]int32]
	RewardCheckPointList AddrTo[[]int32]
	MaxStageCount        int32
	StageBGPrefabPath    AddrTo[StrWithPrefix16]
	StageDecorationPath  AddrTo[StrWithPrefix16]
}

type TowerStageConfigMetaData struct {
	// Properties
	Stageid                    int32
	PreStage                   int32
	Reward                     int32
	RewardShow                 int32
	RecommendGrowBuffLevel     int32
	CgExtraKey                 AddrTo[StrWithPrefix16]
	TowerBuffConfig            AddrTo[[]int32]
	DLCChallengeSpawnPointName AddrTo[StrWithPrefix16]
	StageUIPath                AddrTo[StrWithPrefix16]
}

type TradingCardActivityScheduleMetaData struct {
	// Properties
	OpenTime           AddrTo[RemoteTime]
	CloseTime          AddrTo[RemoteTime]
	NeedMaterialIDList AddrTo[[]ChapterOWTalentLevelMetaData_CostMaterial]
	Reward             int32
	CanTradeNum        int32
}

type TrainingLevelAchieveMetaData struct {
	// Properties
	AvatarID int32
}

type TrainingLevelData struct {
	// Properties
	AvatarID           int32
	AvatarLevel        int32
	AvatarSubskillList AddrTo[[]int32]
	HintMessage        AddrTo[Hash]
}

type TutorialData struct {
	// Properties
	Index                int32
	TriggerMinLevel      int32
	TriggerMaxLevel      int32
	TriggerMissionID     int32
	TriggerOnDoing       bool
	TriggerOnFinish      bool
	TriggerOnClose       bool
	TriggerUIContextName AddrTo[StrWithPrefix16]
	TriggerSpecial       int32
	TriggerSpecialParam  AddrTo[StrWithPrefix16]
	StartStepID          int32
	Reward               int32
	SkipGroup            int32
	SkipFinishMissionID  int32
}

type TutorialGraphicMetaData struct {
	// Properties
	ImagePath    AddrTo[StrWithPrefix16]
	TextTitle    AddrTo[Hash]
	TextContent  AddrTo[Hash]
	TextContent2 AddrTo[Hash]
}

type TutorialRestrictMetaData struct {
	// Properties
	ExcludeTutorialList AddrTo[[]int32]
}

type TutorialStepData struct {
	// Properties
	StepType               uint8
	RootPath               AddrTo[StrWithPrefix16]
	TargetUIPath           AddrTo[StrWithPrefix16]
	TargetButtonPath       AddrTo[StrWithPrefix16]
	HandIconPosType        uint8
	PlayEffect             bool
	ScrollUIPath           AddrTo[StrWithPrefix16]
	OverrideLink           int32
	LinkParams             AddrTo[[]int32]
	GuideDesc              AddrTo[Hash]
	BubblePosType          uint8
	NextStepID             int32
	DelayTime              float32
	FinishOnStart          AddrTo[[]int32]
	FinishOnEnd            AddrTo[[]int32]
	FacePath               AddrTo[StrWithPrefix16]
	TargetCamera           AddrTo[StrWithPrefix16]
	TutorialSoundBankName  AddrTo[StrWithPrefix16]
	TutorialVoiceEvent     AddrTo[StrWithPrefix16]
	AnimationTutorialPath  AddrTo[StrWithPrefix16]
	DisableHighlightInvoke uint8
	TargetUIManager        AddrTo[StrWithPrefix16]
	TargetCamPosition      AddrTo[[]float32]
	ScrollUIPathEnd        AddrTo[StrWithPrefix16]
}

type TutorialStepScrollerData struct {
	// Properties
	DelayTime         float32
	TargetIndex       int32
	TargetSubPathList AddrTo[[]StrWithPrefix16]
	ComparePath       AddrTo[StrWithPrefix16]
	CompareType       AddrTo[StrWithPrefix16]
	CompareField      AddrTo[StrWithPrefix16]
	CompareValue      AddrTo[StrWithPrefix16]
	UseTextID         bool
}

type TutorialVideoMetaData struct {
	// Properties
	GroupID   int32
	StepId    int32
	CgID      int32
	TextTitle AddrTo[Hash]
	ImagePath AddrTo[StrWithPrefix16]
	TextMAPID AddrTo[Hash]
}

type TutorialWeeklyScheduleMetaData struct {
	// Properties
	Day        int32
	Begin      AddrTo[StrWithPrefix16]
	End        AddrTo[StrWithPrefix16]
	TutorialID int32
}

type TVTAvatarConfigMetaData struct {
	// Properties
	OffsetX float32
	OffsetY float32
	Scale   float32
}

type TVTCardLevelMetaData struct {
	// Properties
	AffectRange        uint8
	CardCost           uint16
	Description        AddrTo[Hash]
	ShortDescription   AddrTo[Hash]
	EffectFunctionName AddrTo[StrWithPrefix16]
	ParamList          AddrTo[[]StrWithPrefix16]
	UniqueMaterialList AddrTo[[]StrWithPrefix16]
	CommonMaterialList AddrTo[[]StrWithPrefix16]
}

type TVTCardMetaData struct {
	// Properties
	CardType   uint8
	MaxLevel   uint8
	Rarity     uint8
	CardName   AddrTo[Hash]
	IconPath   AddrTo[StrWithPrefix16]
	UnlockRank uint8
}

type TVTCardSlotMetaData struct {
	// Properties
	// IsUnlocked bool
	// UnlockConditionHint AddrTo[StrWithPrefix16]
	UnlockCondition AddrTo[StrWithPrefix16]
	// UnlockTime AddrTo[StrWithPrefix16]
	UnlockDivisionLv uint16
}

type TVTConstValuesMetaData struct {
	// Properties
	TypeName           AddrTo[Hash]
	TypeInfo           AddrTo[Hash]
	TypeTutorialIDList AddrTo[[]int32]
}

type TVTDivisionMetaData struct {
	// Properties
	ChildDivisionName       AddrTo[Hash]
	DivisionName            AddrTo[Hash]
	DivisionIconPath        AddrTo[StrWithPrefix16]
	DivisionNeedStar        uint32
	FirstDivisionUpReward   int32
	DivisionProtectMaterial int32
	DivisionProtectUseID    uint8
	SeasonOffReward         int32
}

type TVTDivisionProtectUseMetaData struct {
	// Properties
	DivisionProtectCost      uint16
	DivisionProtectMaxNum    uint32
	DivisionProtectInitValue uint16
}

type TVTInLevelInteractionMetaData struct {
	// Properties
	UIPositionID    int32
	Type            int32
	ShowDescription AddrTo[StrWithPrefix16]
	Detail          AddrTo[StrWithPrefix16]
}

type TVTMaterialTransformMetaData struct {
}

type TVTMessageConfigMetaData struct {
	// Properties
	MessagePara uint32
}

type TVTMessageMetaData struct {
	// Properties
	MessageType          uint8
	MessagePara          uint32
	StageID              int32
	IsBattleWinCondition bool
	BossName             AddrTo[Hash]
	BossIcon             AddrTo[StrWithPrefix16]
}

type TVTMissionMetaData struct {
	// Properties
	MissionType  uint8
	NeedDivision uint32
	IconPath     AddrTo[StrWithPrefix16]
}

type TVTRobotDataMetaData struct {
	// Properties
	RobotName    AddrTo[Hash]
	HeadID       int32
	StageID      int32
	AvatadIDList AddrTo[[]int32]
	ElfID        int32
	CardIDList   AddrTo[[]int32]
}

type TVTScheduleMetaData struct {
	// Properties
	ScheduleMissionList AddrTo[[]int32]
}

type TVTSeasonMetaData struct {
	// Properties
	SeasonName           AddrTo[Hash]
	TutorialList         AddrTo[[]int32]
	WebLink              AddrTo[StrWithPrefix16]
	ShopId               uint32
	RewardShowID         int32
	CurrencyID           AddrTo[[]int32]
	CardSlotList         AddrTo[[]int32]
	CardSystemUnlockType uint16
}

type TVTStageMetaData struct {
	// Properties
	MaxEnergyDrop      int32
	MessageIDList      AddrTo[[]int32]
	TVTTutorialIDList  AddrTo[[]int32]
	TVTTutorialStageID int32
}

type TVTTutorialMetaData struct {
	// Properties
	PicPath    AddrTo[StrWithPrefix16]
	DescTitle  AddrTo[StrWithPrefix16]
	DescDetail AddrTo[StrWithPrefix16]
}

type UIOverrideItemMetaData struct {
	// Properties
	OverrideType  uint8
	NodePath      AddrTo[StrWithPrefix16]
	OverrideParam AddrTo[StrWithPrefix16]
}

type UIOverrideScheduleMetaData struct {
	// Properties
	UIType         uint16
	BeginTime      AddrTo[RemoteTime]
	EndTime        AddrTo[RemoteTime]
	UIOverrideList AddrTo[[]uint32]
}

type UltraEndlessBaseRewardMetaData struct {
	// Properties
	GroupLevel uint8
	MinScore   int32
	RewardID   int32
}

type UltraEndlessBattleConfigMetaData struct {
	// Properties
	SiteMapName           AddrTo[StrWithPrefix16]
	MissionIDList         AddrTo[[]int32]
	Weeklyreport_BossInfo int32
	Weeklyreport_BuffInfo int32
}

type UltraEndlessBuffMetaData struct {
	// Properties
	BuffName     AddrTo[StrWithPrefix16]
	IconPath     AddrTo[StrWithPrefix16]
	BuffNameText AddrTo[Hash]
	BuffDescText AddrTo[Hash]
}

type UltraEndlessDisplayMetaData struct {
	// Properties
	UltraEndlessTitle AddrTo[StrWithPrefix16]
	GroupIcon         AddrTo[StrWithPrefix16]
	GroupSmallIcon    AddrTo[StrWithPrefix16]
	GroupIconColor    AddrTo[StrWithPrefix16]
}

type UltraEndlessFloorMetaData struct {
	// Properties
	NeedScore uint16
	MaxScore  uint16
}

type UltraEndlessGroupMetaData struct {
	// Properties
	GroupID    uint8
	GroupName  AddrTo[StrWithPrefix16]
	NeedCupNum uint16
}

type UltraEndlessRewardMetaData struct {
	// Properties
	EndScheduleId uint32
	SettleReward  uint32
}

type UltraEndlessScheduleMetaData struct {
	// Properties
	StartTime        AddrTo[RemoteTime]
	BattleConfigList AddrTo[[]UltraEndlessScheduleMetaData_GroupLevelBattleConfig]
}

type UltraEndlessScheduleMetaData_GroupLevelBattleConfig struct {
	// Fields
	GroupLevel     int32
	BattleConfigID uint32
}

type UltraEndlessSettleCupMetaData struct {
	// Properties
	OnePlayerNoAttendCup int16
	CupChangeList        AddrTo[[]int16]
}

type UltraEndlessSiteMetaData struct {
	// Properties
	StageID         int32
	BuffID          int32
	SiteNodeName    AddrTo[StrWithPrefix16]
	PreSiteList     AddrTo[[]int32]
	SiteName        AddrTo[StrWithPrefix16]
	LevelEndWebLink AddrTo[StrWithPrefix16]
	HardLevelGroup  int32
	WebLink         AddrTo[StrWithPrefix16]
}

type UltraEndlessStageMetaData struct {
	// Properties
	TutorialStageID int32
}

type UltraEndlessTopRankShowMetaData struct {
	// Properties
	RankMin  uint16
	RankMax  uint16
	RewardId int32
}

type UniqueGachaPoolMetaData struct {
}

type UniqueGachaReplaceItemMetaData struct {
	// Properties
	CandidatePoolList AddrTo[[]UniqueGachaReplaceItemMetaData_CandidatePool]
}

type UniqueGachaReplaceItemMetaData_CandidatePool struct {
	// Fields
	SelectNum uint8
	PoolID    uint16
}

type UniqueMonsterMetaData struct {
	// Properties
	Name            AddrTo[Hash]
	MonsterName     AddrTo[StrWithPrefix16]
	TypeName        AddrTo[StrWithPrefix16]
	HPRatio         float32
	AttackRatio     float32
	DefenseRatio    float32
	MoveSpeedRatio  float32
	ATKRatios       AddrTo[[]float32]
	ConfigType      AddrTo[StrWithPrefix16]
	AIName          AddrTo[StrWithPrefix16]
	AttackCDNames   AddrTo[[]StrWithPrefix16]
	AttackCDs       AddrTo[[]float32]
	Abilities       AddrTo[StrWithPrefix16]
	HPSegmentNum    int32
	Scale           AddrTo[[]float32]
	BossRank        int32
	HandBookId      int32
	MonsterRank     int32
	StageDetailInfo AddrTo[StrWithPrefix16]
}

type UnlockUIData struct {
	// Properties
	UnlockByMission  int32
	OnDoing          int32
	OnFinish         int32
	OnClose          int32
	UnlockByTutorial int32
}

type VibrateConfigsMetaData struct {
	// Properties
	ConfigPath AddrTo[StrWithPrefix16]
}

type VirtualAvatarMetaData struct {
	// Properties
	AvatarType             int32
	AvatarId               int32
	PvpAvatarId            int32
	ArtifactLevel          int32
	BaseRarity             int32
	BaseLevel              int32
	DefaultWeapon          int32
	DefaultDress           int32
	PreId                  int32
	EvoMaterialList        AddrTo[[]AvatarArtifactLevelMetaData_UpgradeMaterialItem]
	UnlockMaterialList     AddrTo[[]BeastStageDisplayMetaData_LevelDropDisplay]
	UnlockStigmataSlotList AddrTo[StrWithPrefix16]
}

type VirtualBuffMetaData struct {
	// Properties
	BuffId   AddrTo[StrWithPrefix16]
	BuffIcon AddrTo[StrWithPrefix16]
}

type VirtualCustomLevelDataMetaData struct {
	// Properties
	CostMaterialID  int32
	Cost            int32
	VirtualItemList AddrTo[[]StrWithPrefix16]
	UnlockStageID   int32
	PreCustomLevel  AddrTo[StrWithPrefix16]
}

type VirtualCustomMetaData struct {
	// Properties
	CustomName            AddrTo[Hash]
	IconPath_1            AddrTo[StrWithPrefix16]
	IconPath_2            AddrTo[StrWithPrefix16]
	IconPath_3            AddrTo[StrWithPrefix16]
	IconPath_4            AddrTo[StrWithPrefix16]
	VirtualItemNameList   AddrTo[[]Hash]
	VirtualItemDescList   AddrTo[[]Hash]
	VirtualItemLockedList AddrTo[[]Hash]
}

type VirtualGachaChestMetaData struct {
	// Properties
	GroupID            int32
	GachaPoolList      AddrTo[[]VirtualGachaChestMetaData_GachaPool]
	SelectChestImgPath AddrTo[StrWithPrefix16]
	ChestName          AddrTo[Hash]
	IsOpenAfterStage   bool
}

type VirtualGachaChestMetaData_GachaPool struct {
	// Fields
	GachaPoolID int32
	MaterialID  int32
}

type VirtualGachaItemMetaData struct {
	// Properties
	GachaPoolId int32
	DropId      int32
}

type VirtualGroupConfigMetaData struct {
	// Properties
	GroupTypeParam       int32
	DefaultVirtualAvatar AddrTo[[]int32]
	MaxTeamMemberCount   int32
}

type VirtualRoleMetaData struct {
	// Properties
	RoleID                int32
	VirtualGroupID        int32
	RoleSectionID         int32
	UnlockVirtualAvatarId int32
	UnlockTips            AddrTo[Hash]
}

type VirtualStigmataDetailMetaData struct {
	// Properties
	Star               uint8
	Rarity             int32
	DisplayTitle       AddrTo[StrWithPrefix16]
	BagIconPath        AddrTo[StrWithPrefix16]
	BagDetailIconPath  AddrTo[StrWithPrefix16]
	SlotIconPath       AddrTo[StrWithPrefix16]
	PveHp              int32
	PveAttack          int32
	PveAbilityName     AddrTo[StrWithPrefix16]
	PveAbilityLevel    int32
	PvpHp              int32
	PvpAttack          int32
	PveBuffDes         AddrTo[Hash]
	PvpBuffDes         AddrTo[Hash]
	PveBuffShortDes    AddrTo[Hash]
	PvpBuffShortDes    AddrTo[Hash]
	CommonBuffShortDes AddrTo[Hash]
	PveAbilityParam    AddrTo[[]float32]
}

type VirtualStigmataMetaData struct {
	// Properties
	StigmataId        int32
	BaseRarity        int32
	BaseLevel         int32
	IsAllowMultiDress bool
	PreId             int32
	Type              uint8
}

type VirtualTrainStageDropMetaData struct {
	// Properties
	DropTimes int32
}

type WarehouseRequireData struct {
	// Properties
	BeginTime    AddrTo[StrWithPrefix16]
	EndTime      AddrTo[StrWithPrefix16]
	ItemID       int32
	ItemNum      int32
	CostSC       int32
	CostItem     int32
	CostItemNum  int32
	RequireLevel int32
	CD           int32
	Reward       int32
	Obsolete     int32
}

type WaveRushBuffDataMetaData struct {
	// Properties
	BuffName     AddrTo[StrWithPrefix16]
	Type         uint8
	ActivateCore AddrTo[[]WaveRushBuffDataMetaData_SectTuple]
	BuffSect     uint8
	Name         AddrTo[Hash]
	Icon         AddrTo[StrWithPrefix16]
	IsShowDialog bool
}

type WaveRushBuffDataMetaData_SectTuple struct {
	// Fields
	Type uint8
	Num  int32
}

type WaveRushBuffLevelGroupMetaData struct {
	// Properties
	ShopIndex              int32
	LevelUpCostMaterialNum int32
	ParamList              AddrTo[[]float32]
	BasePropertyList       AddrTo[[]int32]
	Description            AddrTo[Hash]
	SimpleDescription      AddrTo[Hash]
	LimitConditionType     uint8
	LimitConditionParam    AddrTo[[]StrWithPrefix16]
	LockedTips             AddrTo[Hash]
}

type WaveRushFuncOpenMetaData struct {
	// Properties
	ShowConditionType    uint8
	ShowConditionParam   AddrTo[[]StrWithPrefix16]
	UnlockConditionType  uint8
	UnlockConditionParam AddrTo[[]StrWithPrefix16]
	LockedTips           AddrTo[Hash]
}

type WaveRushScheduleMetaData struct {
	// Properties
	BeginTime          AddrTo[StrWithPrefix16]
	EndTime            AddrTo[StrWithPrefix16]
	MinLevel           int32
	UIConfigID         int32
	StageGroupConfigID int32
	BuffGroupConfigID  int32
	CoreBuff           AddrTo[[]int32]
	MaxCoreBuffNum     uint8
	SlotNumPreSiteList AddrTo[[]WaveRushScheduleMetaData_LimitGroup]
	RewardLineID       int32
	Mission            AddrTo[[]int32]
	ActivityMaterialID int32
	MaterialID         AddrTo[[]int32]
	EntryShowType      uint8
	TutorialStage      int32
}

type WaveRushScheduleMetaData_LimitGroup struct {
	// Fields
	BuffLimit uint8
	SiteID    int32
}

type WaveRushStageGroupMetaData struct {
	// Properties
	SiteID             int32
	SiteName           AddrTo[Hash]
	Area               int32
	Type               uint8
	SiteType           uint8
	PreSiteID          int32
	BeginTimeOffSet    int32
	UnlockBuffIDList   AddrTo[[]int32]
	SiteIndex          int32
	ShowDialog         uint8
	DialogDesc         AddrTo[[]Hash]
	IsNeedRank         bool
	IsNeedRankEntrance bool
	IsShowSiteReward   bool
}

type WaveRushStageMetaData struct {
	// Properties
	Type              uint8
	MaxScore          int32
	IsInRewardLine    bool
	StageTime         int32
	WeatherIDList     AddrTo[[]int32]
	SimpleStageDetail AddrTo[StrWithPrefix16]
	IsOnlySimple      bool
	StageScene        AddrTo[StrWithPrefix16]
	WaveIDList        AddrTo[[]int32]
	HardlevelGroupID  int32
	Hardlevel         int32
}

type WaveRushStageScoreDropMetaData struct {
	// Properties
	DropMaterialNum int32
}

type WaveRushStageWaveMetaData struct {
	// Properties
	UIDList          AddrTo[StrWithPrefix16]
	WaveBossUID      AddrTo[[]int32]
	FeverAdd         AddrTo[StrWithPrefix16]
	TimeAdd          AddrTo[StrWithPrefix16]
	ScoreAdd         AddrTo[StrWithPrefix16]
	WaveDistribution uint8
}

type WaveRushUIConfigMetaData struct {
	// Properties
	UIManager         AddrTo[StrWithPrefix16]
	AreaMapList       AddrTo[[]WaveRushUIConfigMetaData_AreaTuple]
	WebLink           AddrTo[StrWithPrefix16]
	Tutorial          AddrTo[[]GlobalWarScheduleMetaData_ConfigTutorial]
	MissionMaxShowNum int32
}

type WaveRushUIConfigMetaData_AreaTuple struct {
	// Fields
	ID int32

	// Objects
	PrefabPath StrWithPrefix16
}

type WeaponBaseTypeMetaData struct {
	// Properties
	CapsulesList AddrTo[[]int32]
	TypeIcon     AddrTo[StrWithPrefix16]
	TypeName     AddrTo[Hash]
	TypeNameEn   AddrTo[Hash]
}

type WeaponCustomDisplayMetaData struct {
	// Properties
	BodyMod AddrTo[StrWithPrefix16]
}

type WeaponLowResOverrideMetaData struct {
	// Properties
	BodyMod AddrTo[StrWithPrefix16]
}

type WeaponMainIDMetaData struct {
	// Properties
	ReforgeTargetIDList AddrTo[[]int32]
}

type WeaponShowOverrideMetaData struct {
	// Properties
	BodyMod          AddrTo[StrWithPrefix16]
	IconPath         AddrTo[StrWithPrefix16]
	ForbidAutoRotate bool
}

type WeaponTagMetaData struct {
	// Properties
	SpecificAvatarBonus        AddrTo[[]StigmataTagMetaData_RatioItem]
	SpecificAvatarBonusSupport AddrTo[[]StigmataTagMetaData_RatioItem]
}

type WebLinkAvatarMetaData struct {
	// Properties
	WebLinks  AddrTo[StrWithPrefix16]
	IsWebView bool
}

type WebLinkElfMetaData struct {
	// Properties
	WebLinks  AddrTo[StrWithPrefix16]
	IsWebView bool
}

type WebLinkEquipMetaData struct {
	// Properties
	WebLinks  AddrTo[StrWithPrefix16]
	IsWebView bool
}

type WebLinkExBossMetaData struct {
	// Properties
	WebLinks  AddrTo[StrWithPrefix16]
	IsWebView bool
}

type WeekDayActivityMetaData struct {
	// Properties
	SeriesID                    int32
	ActivityType                uint8
	ModeSwitch                  uint8
	RefIDSwitch                 AddrTo[[]int32]
	EnableCornerName            bool
	TextMapCornerName           AddrTo[Hash]
	Title                       AddrTo[Hash]
	SubTitle                    AddrTo[Hash]
	Desc                        AddrTo[Hash]
	DescLock                    AddrTo[Hash]
	SmallImgPath                AddrTo[StrWithPrefix16]
	BgImgPath                   AddrTo[StrWithPrefix16]
	EnterImgPath                AddrTo[StrWithPrefix16]
	ShowEnterImg                bool
	ShowRedHint                 bool
	LevelPanelPath              AddrTo[StrWithPrefix16]
	LevelIDList                 AddrTo[[]int32]
	DisplayDropList             AddrTo[[]int32]
	StageEventList              AddrTo[[]int32]
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
	ActivityLinkContent         AddrTo[StrWithPrefix16]
	IsAutoEnterMode             bool
	OpenScoinCost               uint16
	SeasonID                    uint8
	AvatarSampleSwitch          bool
	RemoveTimeoutDisplay        bool
	TeamListID                  uint16
}

type WeekDayActivityScheduleMetaData struct {
	// Properties
	ActivityId      int32
	BeginDate       AddrTo[StrWithPrefix16]
	VersionTagTitle AddrTo[Hash]
}

type WeeklyReportMedalMetaData struct {
	// Properties
	MedalList AddrTo[[]uint32]
}

type WeeklyReportMetaData struct {
	// Properties
	TabName   AddrTo[StrWithPrefix16]
	Priority  int32
	Levelmax  int32
	Levelmin  int32
	BeginTime AddrTo[StrWithPrefix16]
}

type WeeklyReportSubPageDataMetaData struct {
	// Properties
	WebLink         AddrTo[StrWithPrefix16]
	Ishave_extraweb AddrTo[[]int32]
	Levelmax        int32
	Levelmin        int32
}

type WelfareMissionMetaData struct {
	// Properties
	Comment     AddrTo[Hash]
	MissionIcon AddrTo[StrWithPrefix16]
	PlotID      int32
	DisplayType int32
}

type WikiActivityCatalogMetaData struct {
	// Properties
	ActivityID  uint8
	Title       AddrTo[StrWithPrefix16]
	Description AddrTo[StrWithPrefix16]
	Author      AddrTo[StrWithPrefix16]
	AuthorPic   AddrTo[StrWithPrefix16]
}

type WikiActivityChapterMetaData struct {
	// Properties
	ActivityID  uint8
	Title       AddrTo[StrWithPrefix16]
	Description AddrTo[StrWithPrefix16]
	Cover       AddrTo[StrWithPrefix16]
	Author      AddrTo[StrWithPrefix16]
	AuthorQIcon AddrTo[StrWithPrefix16]
}

type WikiActivityEventMetaData struct {
	// Properties
	ChapterID uint16
	Title     AddrTo[StrWithPrefix16]
	Plot      uint16
}

type WikiActivityMetaData struct {
	// Properties
	Name        AddrTo[StrWithPrefix16]
	Time        AddrTo[StrWithPrefix16]
	Year        AddrTo[StrWithPrefix16]
	Season      AddrTo[StrWithPrefix16]
	SeasonPic   AddrTo[StrWithPrefix16]
	SeasonColor AddrTo[StrWithPrefix16]
	Description AddrTo[StrWithPrefix16]
	Cover       AddrTo[StrWithPrefix16]
	Thumbnails  AddrTo[StrWithPrefix16]
}

type WikiCollectionRankMetaData struct {
	// Properties
	RankLevel  int32
	RankName   AddrTo[Hash]
	RankNameEn AddrTo[Hash]
	RankIcon   AddrTo[StrWithPrefix16]
	RankScore  int32
	Reward     int32
}

type WikiTypeMetaData struct {
	// Properties
	WikiTypeName          AddrTo[Hash]
	WikiTypeNameEn        AddrTo[Hash]
	WikiTypeColor         AddrTo[StrWithPrefix16]
	WikiTypeIconTypeColor AddrTo[StrWithPrefix16]
	WikiTypeLineColor     AddrTo[StrWithPrefix16]
	WikiTypeIcon          AddrTo[StrWithPrefix16]
	WikiTypeNameIcon      AddrTo[StrWithPrefix16]
	WikiTypeShareIcon     AddrTo[StrWithPrefix16]
	WikiTypeImg           AddrTo[StrWithPrefix16]
	Lock                  int32
	UnLockLevel           int32
	CrownCoef             int32
	WikiCollectionCoef    int32
}

type WorldMapActivityInfoMetaData struct {
	// Properties
	IconPath    AddrTo[StrWithPrefix16]
	Desc        AddrTo[Hash]
	LinkContent AddrTo[StrWithPrefix16]
}

type WorldMapCoverInfoMetaData struct {
	// Properties
	EntryID       int32
	ContentIDList AddrTo[[]int32]
	Title         AddrTo[Hash]
	Desc          AddrTo[Hash]
}

type WorldMapEntryMetaData struct {
	// Properties
	EntryID                int32
	FeatureType            AddrTo[[]int32]
	ContentIDList          AddrTo[[]int32]
	EntryTitle             AddrTo[Hash]
	EntryImagePath         AddrTo[StrWithPrefix16]
	EntryTitleImgPath      AddrTo[StrWithPrefix16]
	EntryCG                int32
	UnlockLevel            int32
	UnlockHintLevel        int32
	ForceShowContent       int32
	InfoID                 int32
	ShopType               int32
	EntryDesc              AddrTo[Hash]
	EntryDetailImagePath   AddrTo[StrWithPrefix16]
	EntryRewardID          int32
	VersionTag             bool
	VersionTagMinLv        int32
	VersionTagMaxLv        int32
	IsShowActivitySchedule bool
}

type WorldMapScheduleMetaData struct {
	// Properties
	WorldMapID  int32
	AdvanceNote AddrTo[Hash]
}

type ActChallengeMetaDataReader_StructKey struct {
	// Fields
	ActId      int32
	Difficulty int32
}

type ActivityChargeLevelMetaDataReader_StructKey struct {
	// Fields
	AcitivityID int32
	ChargeID    uint8
}

type ActivityLoginBonusMetaDataReader_StructKey struct {
	// Fields
	ID   int32
	Days int32
}

type ActivityPanelScoreDataReader_StructKey struct {
	// Fields
	PanelID   int32
	MissionID int32
}

type ActivityPanelScoreRewardDataReader_StructKey struct {
	// Fields
	PanelID  int32
	Progross int32
}

type ActRitaRankingMetaDataReader_StructKey struct {
	// Fields
	LetterRankID int32
	RankLevel    int32
}

type AddParamGroupMetaDataReader_StructKey struct {
	// Fields
	SkillID int32
	GroupID uint8
}

type AiCyberDailyStageDropLimitMetaDataReader_StructKey struct {
	// Fields
	ActivityScheduleID uint32
	Priority           uint8
}

type AiCyberDailyStageScoreDropMetaDataReader_StructKey struct {
	// Fields
	StageGroupID uint32
	MinScore     int32
}

type AiCyberStoryDataMetaDataReader_StructKey struct {
	// Fields
	StageGroupID int32
	Index        uint8
}

type ArmadaBossActivityScoreMetaDataReader_StructKey struct {
	// Fields
	ActivityID    int32
	ScoreType     int32
	ActivityScore int32
}

type ArmadaRecruitTipsMetaDataReader_StructKey struct {
	// Fields
	ArmadaLevel int32
	TYPE        int32
}

type ArmadaReunionLevelMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	Level      int32
}

type ArmadaReunionMissionMetaDataReader_StructKey struct {
	// Fields
	ScheduleID             int32
	ArmadaReunionMissionID int32
}

type AvatarForgeWhiteListMetaDataReader_StructKey struct {
	// Fields
	AvatarID     int32
	LevelgroupID int32
}

type AvatarGoodFeelMetaDataReader_StructKey struct {
	// Fields
	AvatarId      int32
	GoodFeelLevel int32
}

type AvatarMissionActivityRewardMetaDataReader_StructKey struct {
	// Fields
	ID              int32
	AccumulatedDays int32
}

type AvatarStarMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	Star     int32
	SubStar  int32
}

type AvatarStarTypeMetaDataReader_StructKey struct {
	// Fields
	Star             uint8
	SubStar          uint8
	AvatarType       uint8
	AvatarStarUpType uint8
}

type AvatarStarUpSubSkillDataReader_StructKey struct {
	// Fields
	SkillID int32
	Star    uint8
	SubStar uint8
}

type AvatarSubSkillLevelMetaDataReader_StructKey struct {
	// Fields
	ItemType   int32
	SkillLevel int32
}

type BattlePassSeasonLevelConfigMetaDataReader_StructKey struct {
	// Fields
	SeasonLevelComfigID uint8
	SeasonLevel         uint8
}

type BeastStageDisplayMetaDataReader_StructKey struct {
	// Fields
	BeastStageID int32
	Level        int32
}

type BossChallengeRewardMetaDataReader_StructKey struct {
	// Fields
	RewardConfigID int32
	Index          uint16
}

type ChapterActivityPhaseMetaDataReader_StructKey struct {
	// Fields
	ChapterID int32
	PhaseNum  int32
}

type ChapterActivityRewardLineMetaDataReader_StructKey struct {
	// Fields
	ConfigID    int32
	CountNumber int32
}

type ChapterBWFogConfigMetaDataReader_StructKey struct {
	// Fields
	FogNodeID uint8
	WorldType uint8
}

type ChapterBWWorldBuffSetMetaDataReader_StructKey struct {
	// Fields
	StageID int32
	WaveID  int32
}

type ChapterBWWorldRuneLevelMetaDataReader_StructKey struct {
	// Fields
	LevelConfigID int32
	Level         int32
}

type ChapterBWWorldSageTowerCfgMetaDataReader_StructKey struct {
	// Fields
	SiteType int32
	Param    int32
}

type ChapterBWWorldTowerStageMetaDataReader_StructKey struct {
	// Fields
	StageID int32
	Floor   int32
}

type ChapterChallengeMetaDataReader_StructKey struct {
	// Fields
	ChapterID    int32
	ChallengeNum int32
}

type ChapterCultivateAttrMetaDataReader_StructKey struct {
	// Fields
	CultivateID    uint32
	CultivateLevel uint32
}

type ChapterCultivateBuffMetaDataReader_StructKey struct {
	// Fields
	ID    uint32
	Level uint32
}

type ChapterCultivateJumpMetaDataReader_StructKey struct {
	// Fields
	CultivateLevel uint32
	ChapterID      int32
}

type ChapterCultivateLevelDataMetaDataReader_StructKey struct {
	// Fields
	CultivateUseID uint32
	Level          uint32
}

type ChapterCultivateMonsterMetaDataReader_StructKey struct {
	// Fields
	SkillID    uint32
	SkillLevel uint32
}

type ChapterOWBuildingLevelMetaDataReader_StructKey struct {
	// Fields
	MainID uint32
	Level  uint32
}

type ChapterOWBuildingStateMetaDataReader_StructKey struct {
	// Fields
	MainID    int32
	StateType uint8
}

type ChapterOWChallengeDataMetaDataReader_StructKey struct {
	// Fields
	ConfigID       uint32
	ChallengeIndex uint32
}

type ChapterOWFameMetaDataReader_StructKey struct {
	// Fields
	FameLevel uint8
	MapID     int32
}

type ChapterOWHeroCardLevelMetaDataReader_StructKey struct {
	// Fields
	CardID int32
	Level  uint8
}

type ChapterOWHeroDisplayMetaDataReader_StructKey struct {
	// Fields
	ID         uint32
	HeroStatus uint8
}

type ChapterOWHeroLevelMetaDataReader_StructKey struct {
	// Fields
	HeroID int32
	Level  uint8
}

type ChapterOWMoonTowerFloorMetaDataReader_StructKey struct {
	// Fields
	FloorConfigID uint16
	FloorOrder    uint16
}

type ChapterOWMoonTowerScoreDropMetaDataReader_StructKey struct {
	// Fields
	ScoreDropConfigID uint16
	Score             uint32
}

type ChapterOWRewardMetaDataReader_StructKey struct {
	// Fields
	MapID       int32
	CountNumber int32
}

type ChapterOWTalentLevelMetaDataReader_StructKey struct {
	// Fields
	TalentID    int32
	TalentLevel uint8
}

type ChapterOWTerminalLevelMetaDataReader_StructKey struct {
	// Fields
	MapID int32
	Level uint8
}

type ChatLobbyNPCSystemFunctionMetaDataReader_StructKey struct {
	// Fields
	NPCID       int32
	UIPostionID int32
	Type        int32
}

type CollectionTypeMetaDataReader_StructKey struct {
	// Fields
	SysType        uint8
	CollectionType uint8
}

type CompensationOfDormitoryMetaDataReader_StructKey struct {
	// Fields
	FacilityType int32
	Level        int32
}

type CompensationOfElfBreakMetaDataReader_StructKey struct {
	// Fields
	ElfID     int32
	BreakStep int32
}

type CompensationOfElfSlotUnlockMetaDataReader_StructKey struct {
	// Fields
	ElfID   int32
	SlotNum int32
}

type CompensationOfExtendGradeMetaDataReader_StructKey struct {
	// Fields
	CabinType   int32
	ExtendGrade int32
}

type CompensationOfIslandMetaDataReader_StructKey struct {
	// Fields
	CabinType int32
	Level     int32
}

type CoupleTowerRewardMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	MinFloor   int32
}

type DailyMissionMaterialIconMetaDataReader_StructKey struct {
	// Fields
	MaterialID   int32
	MaterialType int16
}

type DailyRechargeRewardGroupMetaDataReader_StructKey struct {
	// Fields
	GroupID  int32
	Progress int32
}

type DiceyDungeonBubbleMetaDataReader_StructKey struct {
	// Fields
	RoleID   int32
	BubbleID uint8
}

type DiceyDungeonOrnamentMetaDataReader_StructKey struct {
	// Fields
	OrnamentID int32
	Level      int32
}

type DiceyDungeonRoleMetaDataReader_StructKey struct {
	// Fields
	RoleID int32
	Level  int32
}

type DiceyDungeonTutorialDataMetaDataReader_StructKey struct {
	// Fields
	DungeonRoomID int32
	TriggerType   uint8
}

type DiceyDungeonWeaponMetaDataReader_StructKey struct {
	// Fields
	WeaponID int32
	Level    int32
}

type DLCReviveCostMetaDataReader_StructKey struct {
	// Fields
	ReviveTimes int32
	ReviveType  int32
}

type DLCSupportDataMetaDataReader_StructKey struct {
	// Fields
	SupportType  uint8
	SupportParam int32
}

type DLCSupportRewardMetaDataReader_StructKey struct {
	// Fields
	NPCID        uint8
	SupportLevel uint8
}

type DLCTalentLevelMetaDataReader_StructKey struct {
	// Fields
	DLCTalentID int32
	TalentLevel int32
}

type DLCTowerFloorMetaDataReader_StructKey struct {
	// Fields
	Floor      int32
	ConfigType int32
}

type DormitoryDecorationMetaDataReader_StructKey struct {
	// Fields
	DecorationLevel int32
	TargetHouse     int32
}

type ElfSkillTreeMetaDataReader_StructKey struct {
	// Fields
	ElfSkillID int32
	ElfSkillLv int32
}

type ElfStarMetaDataReader_StructKey struct {
	// Fields
	ElfID int32
	Star  int32
}

type ElfStoryActMetaDataReader_StructKey struct {
	// Fields
	StoryId     int32
	TriggerTime int32
}

type EquipTypeDetailMetaDataReader_StructKey struct {
	// Fields
	EquipmentType int32
	SortType      int32
}

type ExaminationRewardMetaDataReader_StructKey struct {
	// Fields
	ConfigID                 int32
	ExaminationQuestionIndex int32
}

type ExBossMonsterLevelConfigMetaDataReader_StructKey struct {
	// Fields
	ID           int32
	MonsterLevel uint8
}

type ExBossMonsterScoreRewardMetaDataReader_StructKey struct {
	// Fields
	RewardConfigID int32
	ScoreLineNode  int32
}

type ExtraStoryChallengeModeMetaDataReader_StructKey struct {
	// Fields
	ChapterId  int32
	Difficulty int32
}

type FeatureSubPakConfigMetaDataReader_StructKey struct {
	// Objects
	ContextName StrWithPrefix16
	Tag         StrWithPrefix16
}

type FoundationRewardMetaDataReader_StructKey struct {
	// Objects
	Product_name StrWithPrefix16
	Level        int32
}

type GalAvatarStandByMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	DressID  int32
}

type GardenCropDataMetaDataReader_StructKey struct {
	// Fields
	CropID   int32
	GardenID int32
}

type GardenWeatherDataMetaDataReader_StructKey struct {
	// Fields
	WeatherID int32
	GardenID  int32
}

type GeneralActivityRankDataMetaDataReader_StructKey struct {
	// Fields
	RankData int32
	Rank     int32
}

type GerenalRulesEffectMetaDataReader_StructKey struct {
	// Fields
	RuleID int32
	Value  int32
}

type GlobalExploreBuffDataMetaDataReader_StructKey struct {
	// Fields
	AreaBuffID int32
	Level      int32
}

type GlobalExploreMapPathMetaDataReader_StructKey struct {
	// Fields
	PosX int32
	PosY int32
}

type GlobalWarSpecialBuffMetaDataReader_StructKey struct {
	// Fields
	SpecialBuffType  uint16
	SpecialBuffLevel uint16
}

type GoBackFundRewardMetaDataReader_StructKey struct {
	// Fields
	FundID int32
	Sort   int32
}

type GoBackGrowUpActivityMetaDataReader_StructKey struct {
	// Fields
	GrowUpConfigID uint8
	Rank           uint8
}

type GoBackMissionDayMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	TabId      int32
}

type GoBackWebMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	TabID      int32
}

type GodWarBuffMetaDataReader_StructKey struct {
	// Fields
	BuffID    uint32
	BuffLevel uint32
}

type GodWarChallengeRewardMetaDataReader_StructKey struct {
	// Fields
	TaleID int32
	Step   int32
}

type GodWarMainAvatarMetaDataReader_StructKey struct {
	// Fields
	GodWarID     uint16
	MainAvatarID uint32
}

type GodWarMaterialMetaDataReader_StructKey struct {
	// Fields
	GodWarID   int32
	MaterialID int32
}

type GodWarPunishRewardMetaDataReader_StructKey struct {
	// Fields
	PunishLevel uint16
	PunishScore uint32
}

type GodWarPunishStepMetaDataReader_StructKey struct {
	// Fields
	PunishStepConfigID uint32
	PunishStep         uint8
}

type GodWarRelationDataMetaDataReader_StructKey struct {
	// Fields
	GodWarID uint16
	RoleID   uint32
	Level    uint16
}

type GodWarSupportAvatarMetaDataReader_StructKey struct {
	// Fields
	GodWarID        uint16
	SupportAvatarID uint32
}

type GodWarSupportLevelMetaDataReader_StructKey struct {
	// Fields
	GodWarID uint16
	Level    uint32
}

type GodWarTalentDataMetaDataReader_StructKey struct {
	// Fields
	TalentID uint32
	GodWarID uint16
}

type GodWarTalentLevelDataMetaDataReader_StructKey struct {
	// Fields
	TalentID    uint32
	TalentLevel uint16
}

type GodWarUseAvatarMetaDataReader_StructKey struct {
	// Fields
	ChapterID int32
	Step      int32
}

type GrandKeyBuffActiveInfoMetaDataReader_StructKey struct {
	// Fields
	UnlockGrandKeyLevel int32
	Order               int32
}

type GrandKeyLevelReader_StructKey struct {
	// Fields
	GrandKeyID int32
	Level      int32
}

type GratuityStageClassMetaDataReader_StructKey struct {
	// Fields
	StageClassID       int32
	GratuityScheduleID int32
}

type GreedyEndlessBattleConfigMetaDataReader_StructKey struct {
	// Fields
	BattleConfig int32
	GroupLevel   int32
	Floor        int32
}

type GreedyEndlessFloorConfigMetaDataReader_StructKey struct {
	// Fields
	FloorConfigID int32
	GroupLevel    int32
	Floor         int32
}

type GreedyEndlessRankRewardMetaDataReader_StructKey struct {
	// Fields
	GroupLevel  int32
	RankPercent int32
}

type GreedyEndlessSettleConfigMetaDataReader_StructKey struct {
	// Fields
	SettleRewardConfigID int32
	GroupLevel           int32
}

type HoukaiTownPathMetaDataReader_StructKey struct {
	// Fields
	TowerID     int32
	PrePosition int32
}

type HoukaiTownStrengthLevelMetaDataReader_StructKey struct {
	// Fields
	LevelType uint8
	Strength  int32
}

type HybridSiteCameraMetaDataReader_StructKey struct {
	// Fields
	HybridSiteID uint32
	ChapterID    uint32
}

type InControlUIButtonConfigReader_StructKey struct {
	// Objects
	ContextName      StrWithPrefix16
	PlayerActionName StrWithPrefix16
}

type InviterActivityRewardConfigMetaDataReader_StructKey struct {
	// Fields
	ConfigID   int32
	InviteeNum int32
}

type JigsawGroupMetaDataReader_StructKey struct {
	// Fields
	JigsawID int32
	GroupID  int32
}

type KingdomsWarWaveMetaDataReader_StructKey struct {
	// Fields
	GroupID int32
	Wave    int32
}

type LoginWishActivityRewardMetaDataReader_StructKey struct {
	// Fields
	ActivityId uint16
	LoginDay   uint8
}

type MapLevelMetaDataReader_StructKey struct {
	// Fields
	Level int32
	MapId int32
}

type MassiveWarDamageRewardMetaDataReader_StructKey struct {
	// Fields
	ScheduleID uint8
	DamageRank uint8
}

type MassiveWarRankRewardMetaDataReader_StructKey struct {
	// Fields
	ScheduleID uint8
	RankId     uint8
}

type MassiveWarRewardMetaDataReader_StructKey struct {
	// Fields
	PlayerGroupID uint16
	ScoreLevel    uint16
}

type MatchThreeRoleMetaDataReader_StructKey struct {
	// Fields
	RoleID uint8
	Level  int32
}

type MatchThreeRoomTutorialMetaDataReader_StructKey struct {
	// Fields
	RoomID int32
	Round  uint8
}

type MatchThreeSkillGemMetaDataReader_StructKey struct {
	// Fields
	SkillGemID uint8
	Level      uint8
}

type MatchThreeTowerMetaDataReader_StructKey struct {
	// Fields
	RoomID int32
	WaveID uint8
}

type MemoirsSpecialDataGroupMetaDataReader_StructKey struct {
	// Fields
	GroupID uint16
	ID      uint16
}

type MiniMonopolyGridMetaDataReader_StructKey struct {
	// Fields
	MapId  uint16
	GridId uint8
}

type MirageAvatarMetaDataReader_StructKey struct {
	// Fields
	StageID  int32
	AvatarID int32
}

type MirageChallengeLevelMetaDataReader_StructKey struct {
	// Fields
	StageID           int32
	StageDifficultyID int32
}

type MirageTalentLevelMetaDataReader_StructKey struct {
	// Fields
	TalentID int32
	Level    uint8
}

type MissionPanelConversationMetaDataReader_StructKey struct {
	// Fields
	ActivityID     int32
	ConversationID uint16
}

type MonopolyLevelInfoMetaDataReader_StructKey struct {
	// Fields
	BuffLevelID uint16
	MonopolyID  uint16
}

type MonsterCardLevelLimitMetaDataReader_StructKey struct {
	// Fields
	ActivityID uint32
	Level      uint8
}

type MonsterCardLevelMetaDataReader_StructKey struct {
	// Fields
	UniqueID uint32
	Level    uint8
}

type MonsterCardStarMetaDataReader_StructKey struct {
	// Fields
	UniqueID uint32
	Star     uint8
}

type MonsterConfigMetaDataReader_StructKey struct {
	// Objects
	MonsterName StrWithPrefix16
	TypeName    StrWithPrefix16
}

type MonsterResistanceMetaDataReader_StructKey struct {
	// Fields
	UniqueMonsterID uint32

	// Objects
	MonsterName StrWithPrefix16
	TypeName    StrWithPrefix16
}

type MPStagePlayerLevelDropDataReader_StructKey struct {
	// Fields
	StageID        int32
	MinPlayerLevel int32
}

type MPTeamSkillMetaDataReader_StructKey struct {
	// Fields
	AvatarId        int32
	MPTeamSkillType int32
}

type NatureCounterMetaDataReader_StructKey struct {
	// Fields
	PlayerNatureType int32
	MonsterNature    int32
}

type NetworkErrCodeMetaDataReader_StructKey struct {
	// Objects
	ErrType StrWithPrefix16
	RetCode StrWithPrefix16
}

type NewbieActivityPanelMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	ActivityID int32
}

type NewbieActivityStageMissionMetaDataReader_StructKey struct {
	// Fields
	ID         int32
	StageLevel int32
}

type NewbieAvatarGuideDataReader_StructKey struct {
	// Fields
	AvatarID         int32
	IsAvatarArtifact bool
}

type NewbieBattleBuffReader_StructKey struct {
	// Fields
	MinLevel int32
	MaxLevel int32
}

type NewbieGrowUpMetaDataReader_StructKey struct {
	// Fields
	ScheduleID uint32
	Rank       uint8
}

type NewbieLevelRushRewardMetaDataReader_StructKey struct {
	// Fields
	RewardConfigID int32
	UnlockLevel    uint8
}

type NinjaSlotLevelMetaDataReader_StructKey struct {
	// Fields
	SlotID uint8
	Level  int32
}

type NPCLevelMetaDataReader_StructKey struct {
	// Fields
	HardLevelGroup int32
	HardLevel      int32
}

type OpenWorldContentMetaDataReader_StructKey struct {
	// Fields
	ID    int32
	MapID int32
}

type OpenWorldFunctionMetaDataReader_StructKey struct {
	// Fields
	Function int32
	MapId    int32
}

type OpenWorldQuestDeleteRuleReader_StructKey struct {
	// Fields
	FinishWay  int32
	IsActivate int32
}

type OpenWorldQuestJudgeDataReader_StructKey struct {
	// Fields
	JudgeLv int32
	QuestLv int32
	MapId   int32
}

type OpenWorldQuestMapLevelMetaDataReader_StructKey struct {
	// Fields
	QuestLevel int32
	MapId      int32
}

type OpenWorldQuestRarityMetaDataReader_StructKey struct {
	// Fields
	QuestLevel int32
	MapId      int32
}

type OpenWorldQuestThemeScheduleDataReader_StructKey struct {
	// Fields
	Step   int32
	AreaID int32
}

type OpenWorldRewardUpListMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	QuestLevel int32
}

type OWActivityLevelMetaDataReader_StructKey struct {
	// Fields
	Level        int32
	OWActivityID int32
}

type OWActivityScheduleMetaDataReader_StructKey struct {
	// Fields
	OWActivityID int32
	MapID        int32
}

type OWAvatarActivityLevelMetaDataReader_StructKey struct {
	// Fields
	ActivityLevel uint8
	ActivityID    uint16
}

type OWAvatarCultivateLevelMetaDataReader_StructKey struct {
	// Fields
	CultivateID int32
	Level       uint8
}

type OWAvatarUnlockMetaDataReader_StructKey struct {
	// Fields
	AvatarID   uint16
	ActivityID uint16
}

type OWEndlessBattleConfigReader_StructKey struct {
	// Fields
	BattleConfigID    int32
	EndlessConfigType uint8
}

type OWEndlessGroupMetaReader_StructKey struct {
	// Fields
	GroupLevel int32
	Type       uint8
}

type OWEndlessPlayerBaseRewardMetaDataReader_StructKey struct {
	// Fields
	GroupLevel int32
	Type       uint8
	MinScore   int32
}

type OWEndlessPlayerGroupMetaReader_StructKey struct {
	// Fields
	PlayerGroup int32
	Type        uint8
}

type OWEndlessRewardConfigMetaDataReader_StructKey struct {
	// Fields
	ConfigID   uint32
	GroupLevel uint32
}

type OWEndlessSingleFloorDataMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	StageID    int32
	Floor      uint8
}

type OWEndlessSingleMonsterGroupMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	BattleID   int32
}

type OWHuntActivityHunterMetaDataReader_StructKey struct {
	// Fields
	HunterID  int32
	Hardlevel uint8
}

type OWHuntActivityMonsterMetaDataReader_StructKey struct {
	// Fields
	MonsterShowID int32
	Hardlevel     uint8
}

type OWHuntActivitySHLevelMetaDataReader_StructKey struct {
	// Fields
	StrongholdLevel     uint8
	StrongholdLevelType uint8
}

type PayInfoMetaDataReader_StructKey struct {
	// Fields
	UserType uint8
	DevType  uint8
}

type PhoneEntranceAcountOverrideMetaDataReader_StructKey struct {
	// Fields
	EntranceId  int32
	AccountType int32
}

type PictureStepMetaDataReader_StructKey struct {
	// Fields
	StepID   int32
	ChoiceID int32
}

type PicTutorialMetaDataReader_StructKey struct {
	// Fields
	ActivityID int32
	ID         int32
}

type PlatformShooterQavatarMetaDataReader_StructKey struct {
	// Fields
	QavatarId int16
	Level     int16
}

type PlatformShooterWeaponMetaDataReader_StructKey struct {
	// Fields
	WeaponId int16
	Level    int16
}

type PVZQavatarMetaDataReader_StructKey struct {
	// Fields
	QavatarID uint32
	Level     uint32
}

type PVZTileMetaDataReader_StructKey struct {
	// Fields
	TowerID int32
	FloorID int32
}

type QAvatarBattleKillingStreakMetaDataReader_StructKey struct {
	// Fields
	Num  uint16
	Type int32
}

type QAvatarMissionMetaDataReader_StructKey struct {
	// Fields
	MissionID int32
	AvatarID  int32
}

type RandomDialogActionMetaDataReader_StructKey struct {
	// Fields
	PlotId   int32
	DialogId int32
}

type RandomEffectLevelMetaDataReader_StructKey struct {
	// Fields
	EffectID    int32
	EffectLevel int32
}

type RecommendEquipMetaDataReader_StructKey struct {
	// Fields
	ID      int32
	EquipID int32
}

type ReplayLobbySubPageMetaDataReader_StructKey struct {
	// Fields
	ID       int32
	FatherID int32
}

type ReunionCookBookMetaDataReader_StructKey struct {
	// Fields
	CookGroupID int32
	CookID      int32
}

type ReunionCookRewardMetaDataReader_StructKey struct {
	// Fields
	ScoreRewardGroup int32
	ID               uint8
}

type RichBuildingMetaDataReader_StructKey struct {
	// Fields
	BuildingType  int32
	BuildingLevel int32
}

type RichFloorInfoMetaDataReader_StructKey struct {
	// Fields
	Area    int32
	FloorID int32
}

type RpgAreaLineMetaDataReader_StructKey struct {
	// Fields
	PreAreaID int32
	AreaID    int32
}

type RpgBuffSuitClientInfoMetaDataReader_StructKey struct {
	// Fields
	BuffDataID int32
	Level      int32
}

type RpgCollectionRewardDataReader_StructKey struct {
	// Fields
	TaleID     int32
	PositionID int32
}

type RpgSimpleBuffMetaDataReader_StructKey struct {
	// Fields
	BuffLevel      uint16
	BuffMaterialID int32
}

type RpgSiteLineMetaDataReader_StructKey struct {
	// Fields
	PreSiteID int32
	SiteID    int32
}

type RpgStageDropMetaDataReader_StructKey struct {
	// Fields
	StageID        int32
	DropMaterialID int32
}

type RpgTowerV2MetaDataReader_StructKey struct {
	// Fields
	Stage int32
	Floor int16
}

type SanctuaryBuildingLevelMetaDataReader_StructKey struct {
	// Fields
	SanctuaryLevel      int32
	SanctuaryActivityID int32
}

type SanctuaryPlayerGroupMetaDataReader_StructKey struct {
	// Fields
	PlayerGroupID int32
	ActivityID    int32
}

type ScDLCAvatarLevelMetaDataReader_StructKey struct {
	// Fields
	AvatarID int32
	Level    int32
}

type ScDLCFeverAbilityMetaDataReader_StructKey struct {
	// Fields
	ID           int32
	AbilityLevel int32
}

type ScDLCStoryMetaDataReader_StructKey struct {
	// Fields
	ID    uint32
	State uint8
}

type ScratchTicketItemDataMetaDataReader_StructKey struct {
	// Fields
	PlateID       int32
	PlateDetailID int32
}

type ServantDialogMetaDataReader_StructKey struct {
	// Fields
	DialogID int32
	MapId    int32
}

type ServantTouchMetaDataReader_StructKey struct {
	// Fields
	Id    int32
	MapID int32
}

type SettingAudioVolumeMetaDataReader_StructKey struct {
	// Objects
	AudioSettingOption StrWithPrefix16
	AudioSettingType   uint8
	ParamIndex         uint8
}

type SettingGraphicsDeviceLimitMetaDataReader_StructKey struct {
	// Objects
	Device                  StrWithPrefix16
	GraphicsSettingItemName StrWithPrefix16
}

type SettingGraphicsItemLineMetaDataReader_StructKey struct {
	// Fields
	ID        int32
	Hierarchy uint8
}

type ShareSwitchMetaDataReader_StructKey struct {
	// Fields
	ChannelId  int32
	DeviceType int16
}

type ShopAboutToOnlineItemReader_StructKey struct {
	// Fields
	ShopID int32
	ItemID int32
}

type SignInRewardMetaDataReader_StructKey struct {
	// Fields
	Month int32
	Day   int32
}

type SimulateRankRewardMetaDataReader_StructKey struct {
	// Fields
	ActivityId  int32
	RankPercent int32
}

type SinDemonExMonsterScheduleMetaDataReader_StructKey struct {
	// Fields
	ID       int32
	ConfigId int32
}

type SingleWantedProgressMetaDataReader_StructKey struct {
	// Fields
	ThemeID  uint32
	Progress uint32
}

type SLGBattlePointMetaDataReader_StructKey struct {
	// Fields
	PointID int32
	MapID   int32
}

type SLGBigBossPointMetaDataReader_StructKey struct {
	// Fields
	BossScheduleID int32
	BossID         int32
}

type SLGRankRewardMetaDataReader_StructKey struct {
	// Fields
	ScheduleID int32
	Rank       int32
}

type SlotMachineProgressRewardMetaDataReader_StructKey struct {
	// Fields
	PanelID          int32
	ProgressRewardID int32
}

type SpecialConfigMetaDataReader_StructKey struct {
	// Fields
	ThemeID int32

	// Objects
	DefaultResPath StrWithPrefix16
}

type StageScoreRewardMetaDataReader_StructKey struct {
	// Fields
	StageID  uint32
	MinScore uint32
}

type SupportAvatarInitialMetaDataReader_StructKey struct {
	// Fields
	ActivityType uint16
	ActivityID   uint32
}

type SupportAvatarLevelMetaDataReader_StructKey struct {
	// Fields
	SupportAvatarID    uint16
	SupportAvatarLevel uint16
}

type SupportMissionConfigMetaDataReader_StructKey struct {
	// Fields
	SupportMissionConfigID int32
	Index                  int32
}

type ThemeTowerScoreDropMetaDataReader_StructKey struct {
	// Fields
	ScoreDropConfigID int32
	Score             int32
}

type ThemeTowerSectionLevelMetaDataReader_StructKey struct {
	// Fields
	SectionLevelConfigID int32
	Level                uint8
}

type ThemeTowerStageBuffMetaDataReader_StructKey struct {
	// Fields
	BuffID int32
	Level  uint8
}

type ThemeTowerStageFloorMetaDataReader_StructKey struct {
	// Fields
	FloorConfigID int32
	FloorOrder    uint8
}

type ThemeTowerStageStepMetaDataReader_StructKey struct {
	// Fields
	StepConfigID int32
	StepLevel    uint8
}

type TileValueKeyMetaDataReader_StructKey struct {
	// Fields
	MapID int32
	Key   int32
}

type TowerGrowBuffConfigMetaDataReader_StructKey struct {
	// Fields
	BuffLevelID int32
	ActivityID  int32
}

type TutorialGraphicMetaDataReader_StructKey struct {
	// Objects
	TutorialKey StrWithPrefix16
	StepId      uint8
}

type TVTCardLevelMetaDataReader_StructKey struct {
	// Fields
	ID    int32
	Level uint8
}

type TVTDivisionMetaDataReader_StructKey struct {
	// Fields
	DivisionID uint32
	SeasonID   uint32
}

type TVTMaterialTransformMetaDataReader_StructKey struct {
	// Fields
	MaterialID      int32
	TransformCardID uint16
}

type TVTMessageConfigMetaDataReader_StructKey struct {
	// Fields
	MessageID      uint32
	HardLevelGroup uint32
	HardLevel      uint32
}

type UltraEndlessFloorMetaDataReader_StructKey struct {
	// Fields
	FloorID uint8
	StageID int32
}

type UltraEndlessRewardMetaDataReader_StructKey struct {
	// Fields
	BeginScheduleId uint32
	GroupLevel      uint8
}

type UltraEndlessSettleCupMetaDataReader_StructKey struct {
	// Fields
	GroupLevel       uint8
	PlayerNumInGroup uint8
}

type UniqueGachaPoolMetaDataReader_StructKey struct {
	// Fields
	PoolId uint16
	ItemId uint32
}

type VirtualCustomLevelDataMetaDataReader_StructKey struct {
	// Fields
	CustomID int32
	Level    int32
}

type VirtualTrainStageDropMetaDataReader_StructKey struct {
	// Fields
	StageId  int32
	EndFloor int16
}

type WaveRushBuffLevelGroupMetaDataReader_StructKey struct {
	// Fields
	BuffGroupConfigID int32
	BuffID            int32
	Level             int32
}

type WaveRushStageGroupMetaDataReader_StructKey struct {
	// Fields
	StageGroupConfigID int32
	StageID            int32
}

type WaveRushStageScoreDropMetaDataReader_StructKey struct {
	// Fields
	StageID      int32
	NeedMinScore int32
}

type WeaponCustomDisplayMetaDataReader_StructKey struct {
	// Fields
	ID           int32
	AvatarIDList int32
}

type WeaponLowResOverrideMetaDataReader_StructKey struct {
	// Fields
	AvatarID uint16
	WeaponID int32
}

type WeeklyReportSubPageDataMetaDataReader_StructKey struct {
	// Fields
	ID       int32
	FatherID int32
}

type EditorUniqueMonsterMetaData struct {
	// Properties
	Info           AddrTo[StrWithPrefix16]
	Group          uint32
	Category       AddrTo[StrWithPrefix16]
	Name           AddrTo[Hash]
	MonsterName    AddrTo[StrWithPrefix16]
	TypeName       AddrTo[StrWithPrefix16]
	HPRatio        float32
	AttackRatio    float32
	DefenseRatio   float32
	MoveSpeedRatio float32
	ATKRatios      AddrTo[[]float32]
	ConfigType     AddrTo[StrWithPrefix16]
	AIName         AddrTo[StrWithPrefix16]
	AttackCDNames  AddrTo[[]StrWithPrefix16]
	AttackCDs      AddrTo[[]float32]
	Abilities      AddrTo[StrWithPrefix16]
	HPSegmentNum   int32
	Scale          AddrTo[[]float32]
	BossRank       int32
	HandBookId     int32
}

type MonsterActivityMetaData struct {
	// Properties
	ActivityName AddrTo[StrWithPrefix16]
}

type MonsterProtoTypeMetaData struct {
	// Properties
	Info           AddrTo[StrWithPrefix16]
	Details        AddrTo[StrWithPrefix16]
	Category       AddrTo[StrWithPrefix16]
	Name           AddrTo[Hash]
	MonsterName    AddrTo[StrWithPrefix16]
	TypeName       AddrTo[StrWithPrefix16]
	HPRatio        float32
	AttackRatio    float32
	DefenseRatio   float32
	MoveSpeedRatio float32
	ATKRatios      AddrTo[[]float32]
	ConfigType     AddrTo[StrWithPrefix16]
	AIName         AddrTo[StrWithPrefix16]
	AttackCDNames  AddrTo[[]StrWithPrefix16]
	AttackCDs      AddrTo[[]float32]
	Abilities      AddrTo[StrWithPrefix16]
	HPSegmentNum   int32
	Scale          AddrTo[[]float32]
	BossRank       int32
	HandBookId     int32
}

type PropListMetaData struct {
	// Properties
	Info     AddrTo[StrWithPrefix16]
	Details  AddrTo[StrWithPrefix16]
	Category AddrTo[StrWithPrefix16]
}

type MultiPlatforms_UserInfoMetaData struct {
	UserType    byte
	AccountType byte
	UserLabel   StrWithPrefix16
}

type MultiPlatforms_UserInfoMetaDataReader_StructKey struct {
	AccountType byte
	UserType    byte
}

type UnknownAsset1MetaData struct {
	Field1 bool
	ID     int32
	Field2 []RandomDialogCGRawMetaData_CGRawPos
}
