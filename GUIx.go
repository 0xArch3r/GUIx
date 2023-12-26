package GUIx

type StyleOption struct {
	Key   string
	Value string
}

func WithColor(color string) StyleOption {
	return StyleOption{
		Key:   "color",
		Value: color,
	}
}
