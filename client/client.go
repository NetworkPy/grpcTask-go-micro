package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/NetworkPy/grpcTask/internal/point"
	micro "github.com/asim/go-micro/v3"
)

func main() {
	pathInput := "example/input.txt"
	pathOut := "example/output.txt"
	strArr, err := ReadPoints(pathInput)
	if err != nil {
		log.Fatal(err)
	}

	points, err := CreatePointsStruct(strArr)
	if err != nil {
		log.Fatal(err)
	}

	service := micro.NewService()

	service.Init()
	client := point.NewPointserviceService("pointsservice", service.Client())

	res, err := client.CreateGoodPoints(context.Background(), &point.PointsReq{
		Points: points,
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err = WritePoints(res.Points, pathOut); err != nil {
		log.Fatalln(err)
	}
}

// ReadPoints...
func ReadPoints(path string) ([][]string, error) {
	file, err := os.Open(path)

	scanner := bufio.NewScanner(file) // Построчно сканируем файл (вдруг там будет 10к точек)
	if err != nil {
		return nil, err
	}

	result := [][]string{}
	for scanner.Scan() {
		textStr := scanner.Text()
		testArr := strings.Split(textStr, ";")
		result = append(result, testArr)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	return result, nil
}

//  CreatePointsStruct...
func CreatePointsStruct(strArr [][]string) ([]*point.Point, error) {
	ps := make([]*point.Point, len(strArr)) // Массив для структур
	for idx, str := range strArr {

		// Приведение типов
		x, err := strconv.Atoi(str[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(str[1])
		if err != nil {
			return nil, err
		}
		// Создаем новую структуру для точки
		p := new(point.Point)
		p.X, p.Y = int64(x), int64(y)
		ps[idx] = p
	}
	return ps, nil
}

// WritePoints...
func WritePoints(p []*point.Point, filepath string) error {
	var file *os.File
	var strBuilder strings.Builder // Очень быстрый способ объединения строк
	if _, err := os.Stat(filepath); err == nil {
		file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		defer file.Close()
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) { // Проверка на существование файла
		file, err = os.Create(filepath)
		defer file.Close()
		if err != nil {
			return err
		}
	}
	for _, point := range p {
		x := strconv.FormatInt(point.X, 10)
		y := strconv.FormatInt(point.Y, 10)

		strBuilder.WriteString(x)
		strBuilder.WriteString(";")
		strBuilder.WriteString(y)
		strBuilder.WriteString("\n")
	}

	// Перезапись файла
	_, err := file.WriteString(strBuilder.String())
	if err != nil {
		return err
	}
	return nil
}
