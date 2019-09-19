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
		"Salut, yang",
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
