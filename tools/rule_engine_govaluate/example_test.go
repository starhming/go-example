package rule_engine_govaluate

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Knetic/govaluate"
)

func TestDemo1(t *testing.T) {
	expr, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		log.Fatal("syntax error:", err)
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		log.Fatal("evaluate error:", err)
	}
	fmt.Println(result)
}

func TestDemo2(t *testing.T) {
	expr, _ := govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{})
	parameters["foo"] = -1
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	parameters = make(map[string]interface{})
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)

	// expr, _ = govaluate.NewEvaluableExpression("in_cap_fps_arr >  0")

	expr, _ = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	parameters = make(map[string]interface{})
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}

func TestDemo3(t *testing.T) {
	expr, _ := govaluate.NewEvaluableExpression("a + b + c")

	for i := 0; i < 100; i++ {
		parameters := make(map[string]interface{})
		parameters["a"] = 1
		parameters["b"] = 2
		parameters["c"] = 3
		go func(i int, parameters map[string]interface{}) {
			result, _ := expr.Evaluate(parameters)
			fmt.Println(i, ":", result)
		}(i, parameters)
	}

	time.Sleep(time.Minute)
}

func TestDemo4(t *testing.T) {
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return length, nil
		},
	}

	exprString := "strlen('teststring')"
	expr, _ := govaluate.NewEvaluableExpressionWithFunctions(exprString, functions)
	result, _ := expr.Evaluate(nil)
	fmt.Println(result)
}

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u User) Fullname() string {
	return u.FirstName + " " + u.LastName
}

func TestAccessors(t *testing.T) {
	u := User{FirstName: "li", LastName: "dajun", Age: 18}
	parameters := make(map[string]interface{})
	parameters["u"] = u

	expr, _ := govaluate.NewEvaluableExpression("u.Fullname()")
	result, _ := expr.Evaluate(parameters)
	fmt.Println("user", result)

	expr, _ = govaluate.NewEvaluableExpression("u.Age > 18")
	result, _ = expr.Evaluate(parameters)
	fmt.Println("age > 18?", result)
}
