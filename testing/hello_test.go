package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_FailNow_로그출력_X_다음코드_실행_X(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		t.FailNow()
	}
	fmt.Println("다음 코드 실행 안함")
}

func Test_Fatal_로그출력_O_다음코드_실행_X(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		t.Fatal("fatal은 로그를 출력할 수 있음") //나머지는 FailNow()와 같음
	}
	fmt.Println("다음 코드 실행 안함")
}

func Test_Fail_로그출력_X_다음코드_실행_O(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		t.Fail()
	}
	fmt.Println("다음 코드 실행함")
}

func Test_Error_로그출력_O_다음코드_실행_O(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		t.Error("Error는 로그를 출력할 수 있음") //나머지는 Fail() 메서드와 같음
	}
	fmt.Println("다음 코드 실행함")
}

func Test_Log_로그출력_O_다음코드_실행_O(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		t.Log("로그를 출력할 수 있음")
	}
	fmt.Println("다음 코드 실행함")
}

func Test_Failed_로그출력_X_다음코드_실행_O(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		t.Failed() //실패하더라도 리포트하지 않음
	}
	fmt.Println("다음 코드 실행함")
}

//simple test
func TestGreet(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek"
	if actual != expected {
		//t.Errorf("expected %s, but was %s", expected, actual)
		t.Fail()
	}
	fmt.Println("end")
}

func TestGreet_(t *testing.T) {
	actual := greet("")
	expected := "hello, there!"
	if actual != expected {
		t.Errorf("expected %s, but was %s", expected, actual)
	}
}

//sub-tests
func TestGreet2(t *testing.T) {
	t.Run("test blank value", func(te *testing.T) {
		actual := greet("")
		expected := "hello, there!"
		if actual != expected {
			te.Errorf("expected %s, but was %s", expected, actual)
		}
	})

	t.Run("test valid value", func(te *testing.T) {
		actual := greet("abhishek")
		expected := "hello, abhishek!"
		if actual != expected {
			te.Errorf("expected %s, but was %s", expected, actual)
		}
	})
}

//table driven tests
func TestGreet3(t *testing.T) {
	type testCase struct {
		name             string
		input            string
		expectedGreeting string
	}

	testCases := []testCase{
		{name: "test blank value", input: "", expectedGreeting: "hello, there!"},
		{name: "test valid value", input: "abhishek", expectedGreeting: "hello, abhishek!"},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(te *testing.T) {
			actual := greet(test.input)
			expected := test.expectedGreeting
			if actual != expected {
				te.Errorf("expected %s, but was %s", expected, actual)
			}
		})
	}
}

//parallel tests
func TestGreet4(t *testing.T) {
	type testCase struct {
		name             string
		input            string
		expectedGreeting string
	}

	testCases := []testCase{
		{name: "test blank value", input: "", expectedGreeting: "hello, there!"},
		{name: "test valid value", input: "abhishek", expectedGreeting: "hello, abhishek!"},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(te *testing.T) {
			te.Parallel()
			//simulating time taking test: total time for test execution little > 3s which means tests ran parallely

			time.Sleep(3 * time.Second)
			actual := greet(test.input)
			expected := test.expectedGreeting
			if actual != expected {
				te.Errorf("expected %s, but was %s", expected, actual)
			}
		})
	}
}
