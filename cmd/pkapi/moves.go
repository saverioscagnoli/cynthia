package pkapi

type Move struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	Accuracy         *int             `json:"accuracy,omitempty"`
	Power            *int             `json:"power,omitempty"`
	BasePP           int              `json:"base_pp"`
	Priority         int              `json:"priority"`
	EffectChance     *int             `json:"effect_chance,omitempty"`
	DamageClass      *MoveDamageClass `json:"damage_class,omitempty"`
	Type             Type             `json:"type"`
	MinHits          *int             `json:"min_hits,omitempty"`
	MaxHits          *int             `json:"max_hits,omitempty"`
	MinTurns         *int             `json:"min_turns,omitempty"`
	MaxTurns         *int             `json:"max_turns,omitempty"`
	Drain            *int             `json:"drain,omitempty"`
	Healing          *int             `json:"healing,omitempty"`
	DrainOrRecoil    *int             `json:"drain_or_recoil"`
	CritRateBonus    int              `json:"crit_rate_bonus"`
	FlinchChance     int              `json:"flinch_chance"`
	StatChangeChance int              `json:"stat_change_chance"`
}
