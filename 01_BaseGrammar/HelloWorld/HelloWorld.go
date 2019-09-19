package HelloWorld

const (
	helloPrefix        = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

//Switch的使用
func LanguageSwitch(name, language string) string {
	//条件语句的使用
	if name == "" {
		name = "Word!"
	}

	prefix := helloPrefix

	//switch的使用
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}
