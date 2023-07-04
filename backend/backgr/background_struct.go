package backgr

type BackgroundJsonStruct struct {
	Backgrounds []BackgroundJsonData `json:"backgrounds"`
}

type Type struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type CharacterTrait struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type Ideal struct {
	Dice    string `json:"dice"`
	Good    string `json:"Good"`
	Lawful  string `json:"Lawful"`
	Chaotic string `json:"chaotic"`
	Evil    string `json:"Evil"`
	Neutral string `json:"Neutral"`
	Any     string `json:"Any"`
}
type Affection struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type Weakness struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}

type BackgroundJsonData struct {
	BackgroundName   string         `json:"background_name"`
	BackgroundNameRu string         `json:"background_name_ru"`
	Description      string         `json:"description"`
	Personalization  string         `json:"personalization"`
	SkillMastery     []string       `json:"skill_mastery"`
	Type             Type           `json:"type,omitempty"`
	CharacterTrait   CharacterTrait `json:"character_trait"`
	Ideal            Ideal          `json:"ideal"`
	Affection        Affection      `json:"affection"`
	Weakness         Weakness       `json:"weakness"`
}

type Background struct {
	BackgroundName   string   `json:"background_name"`
	BackgroundNameRu string   `json:"background_name_ru"`
	Description      string   `json:"description"`
	Advice           string   `json:"advice"`
	Personalization  string   `json:"personalization"`
	SkillMastery     []string `json:"skill_mastery"`
	Type             string   `json:"type"`
	CharacterTrait   string   `json:"character_trait"`
	Ideal            string   `json:"ideal"`
	Affection        string   `json:"affection"`
	Weakness         string   `json:"weakness"`
}
