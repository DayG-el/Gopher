package main

var globalBlock int = 100 // gloabl variable block
func main() {
	print(globalBlock)
	{
		var localBlock = 200 // local variable block
		print(localBlock)
		{
			var localBlock = 1250 // this is a local variable
			localBlock = localBlock * 9
			println(localBlock)
			print(globalBlock)
		}

	}
	calc(globalBlock)
	print(globalBlock)
}
func calc(block int) {
	globalBlock = globalBlock + 10
}
