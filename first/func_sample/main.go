package func_sample

import "fmt"

func errorResult(isError bool) (string, error) {
	if isError {
		return "", fmt.Errorf("error")
	}
	return "non error", nil
}

func twoResults() (string, string) {
	return "one", "two"
}

func oneResult() string {
	return "one"
}

func Run() {
	fmt.Println(oneResult())
	fmt.Println(twoResults())

	// when error
	result, err := errorResult(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	// when non error
	result, err = errorResult(false)
	fmt.Println(result)
}