package types

type MessageMap struct {
	Name    string            `types:"map_name"`
	Entries map[string]string `types:"entries"`
}

type MapsDocument struct {
	Maps []Map `types:"maps"`
}

type GeneralResponse struct {
	HttpCode int
	Code     int
	Message  string
}

type Maps struct {
	Map []Map `types:"maps"`
}

type Map struct {
	MapName string `types:"map_name"`
	Entry   Entry  `types:"entry"`
}

type Entry struct {
	Code    string `types:"code"`
	Message string `types:"code"`
}
