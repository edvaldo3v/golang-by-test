package hello

const prefixHelloLanguage = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return prefixHelloLanguage + name
}
