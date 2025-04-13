[![Go Reference](https://pkg.go.dev/badge/github.com/lucky-patty/promptgo.svg)](https://pkg.go.dev/github.com/lucky-patty/promptgo)
# Promptgo 
This is prompt-ui for Golang. 
I will add more feature and test in the future including fixing the original bug in the package.

**Note:** I am very new in this area and push this project to gain more understanding of Golang CLI

# Usage
```
choice, err := promptgo.Select("Choose environment:", []string{"local", "production"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("You selected:", choice)
```
**Note:** I will add more example
