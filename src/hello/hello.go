package main

const spanish  = "Spanish"
const englishHelloPrefix = "Hello，"
const helloEnd  = "World"
const spanishHelloPrefix   = "Hola，"
const frenchHelloPrefix  = "Bonjour le "
const french = "French"

func Hello(str string,language string) string{
	// 编写一个测试
	// 让编译通过
	// 运行测试，查看失败原因并检查错误消息是很有意义的
	// 编写足够的代码以使测试通过
	// 重构

	if str== "" {
		str = helloEnd
	}
	return greetingPrefix(language) + str
	//编写一个失败的测试，并查看失败信息，可以看到我们已经为需求写了一个 相关 的测试，并且看到它产生了一个 易于理解的失败描述
	//编写最少量的代码以使其通过，因此我们知道我们有可工作软件
	//然后 重构，支持我们测试的安全性，以确保我们拥有易于使用的精心制作的代码
}

func greetingPrefix(language string)(prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix


}