package main

import (
	"fmt"
	"strconv"
)

type ItemData struct {
	Type            string
	Title           string
	Brand           string
	OperatingSystem string
}

type Inventory struct {
	Item []ItemData
}

// Why this isn't constant because Go does not support struct constants
// https://stackoverflow.com/questions/43368604/constant-struct-in-go/43368686
var inventory = []Inventory{
	{
		[]ItemData{
			{
				Title: "Jurrasic Park",
				Type:  "Book",
			},
			{
				Brand: "Mercedes Benz",
				Type:  "Car",
			},
			{
				Type:            "Smartphone",
				Brand:           "Samsung",
				OperatingSystem: "Android",
			},
			{
				Type:  "Car",
				Brand: "Ferari",
			},
			{
				Type:  "Book",
				Title: "Harry Potter and The Chamber of Secret",
			},
		},
	},
	{
		[]ItemData{
			{
				Type:  "Car",
				Brand: "Tesla",
			},
			{
				Type:            "Smartphone",
				Brand:           "Apple",
				OperatingSystem: "iOS",
			},
			{
				Type:            "Smartphone",
				Brand:           "Xiaomi",
				OperatingSystem: "Android",
			},
			{
				Type:  "Book",
				Title: "Learning Data Mining with Python",
			},
		},
	},
}

func main() {
	result := GetResult()
	fmt.Print(result)
}

func GetResult() string {
	var tmpType []string
	var tmpCount int = 0
	var result string = ""
	for _, items := range inventory {
		CategorizeType(items.Item, &tmpType)
	}

	for i, tipe := range tmpType {
		index := strconv.Itoa(i + 1)
		result = result + index + ". " + tipe + "\n"
		for _, items := range inventory {
			for _, item := range items.Item {
				if tipe == item.Type {
					tmpCount++

					value, err := ExtractValue(item, item.Type)
					if err != nil {
						panic(err)
					}

					count := strconv.Itoa(tmpCount)
					result = result + index + "." + count + " " + value + "\n"
				}
			}
		}
	}

	return result
}

func CategorizeType(items []ItemData, tmpType *[]string) {
	for _, item := range items {
		if !Find(*tmpType, item.Type) {
			*tmpType = append(*tmpType, item.Type)
		}
	}
}

func ExtractValue(item ItemData, val string) (string, error) {
	if val == "Book" {
		val = item.Title
		return val, nil
	}

	if val == "Car" {
		val = item.Brand
		return val, nil
	}

	if val == "Smartphone" {
		val = item.Brand + " (" + item.OperatingSystem + ")"
		return val, nil
	}

	return val, fmt.Errorf("Type not detected")
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
