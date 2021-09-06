package main

import (
	"fmt"
	"testing"
)

// Так себе тесты, не успел дописать полностью
// Чисто, чтобы проверить работоспособность
var strArr = [][]string{}

func TestReadPoints(t *testing.T) {
	var filepath = "example/input.txt"
	strArr, _ = ReadPoints(filepath)
}

func TestCreatePointsStruct(t *testing.T) {
	p, _ := CreatePointsStruct(strArr)
	fmt.Println(len(p))
}
