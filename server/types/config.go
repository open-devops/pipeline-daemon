package types

type ConfigItem struct {
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ConfigItems []ConfigItem
