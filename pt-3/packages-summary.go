package main

/*
			This is all about packages, there are two types:
			a. Executable & b. Non-executable
		a. Executable Packages: they are used for running applications, strictly non-importable, package must contain the name
		main, must contain a call to the function main
	   b. Non-executable packages also called library packages, third-party packages. These packages are not executed directly.
		  You cannot do "go run filename.go" on them. They do not require a call to the main function. They are used for code re-usabilitty,
			they are importable and can take on any name.
*/
