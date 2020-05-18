package HelloWorld

var HelloCases = []struct {
	name     string
	language string
	want     string
}{
	{
		"wei",
		"",
		"Hello, wei",
	},
	{
		"yang",
		"french",
		"Bonjour, yang",
	},
	{
		"li",
		"spanish",
		"Hola, li",
	},
	{
		"",
		"",
		"Hello, Word!",
	},
}
