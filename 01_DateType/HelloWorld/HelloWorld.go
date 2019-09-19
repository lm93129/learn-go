package HelloWorld

const helloPrefix = "hello,"

func HelloName(name string) string {
	if name == "" {
		name = "Word!"
	}
	return helloPrefix + name
}

func HelloWorld() string {
	return "Hello World!"
}
