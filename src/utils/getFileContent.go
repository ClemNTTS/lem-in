package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFileText(file string) []string {
	text, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	contentText := string(text)
	tabValues := strings.Split(contentText, "\r\n")
	return tabValues
}

func ScrapValues(vals []string) (int, string, []string, string, []string) {
	start, end := "", ""
	nbAnt, _ := strconv.Atoi(vals[0])
	var tabRooms []string
	var tabLinks []string
	vals = vals[1:]
	for ind, value := range vals {
		if value == ""{
			continue
		}

		if (value[0] == 'L' || value[0] == '#') && value[1] != '#' {
			continue
		}
		if !strings.Contains(value, "-") {
			if value == "##start" && ind < len(vals) {
				start = string(vals[ind+1])
				ind++
				continue
			} else if value == "##end" && ind < len(vals) {
				end = string(vals[ind+1])
				ind++
				continue
			}
			tabRooms = append(tabRooms, value)
		} else {
			tabLinks = append(tabLinks, value)
		}
	}
	return nbAnt, start, tabRooms, end, tabLinks
}

func SplitTabRoom(tabRooms []string) [][]string {
	var tabAllRooms [][]string

	for _, value := range tabRooms {
		tabtemp := strings.Split(value, " ")
		if len(tabtemp) == 3{
			tabAllRooms = append(tabAllRooms, tabtemp)
		}
	}
	return tabAllRooms
}

func Affich(tab []string) {
	for _, value := range tab {
		fmt.Println(value)
	}
}
