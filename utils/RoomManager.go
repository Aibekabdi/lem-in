package lemin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func TakeInfo(info []string) ([]string, [][]string, Colony) {
	var r []string
	var t [][]string
	var colony Colony
	if len(info) >= 1 {
		for i := 0; i < len(info); i++ {
			if len(info[i]) > 0 {
				if i == 0 {
					nb, _ := strconv.Atoi(info[0])
					colony.Ants = nb
				} else if (info[i] != "##start" && info[i] != "##end") && (info[i][0] == '#' || info[i][0] == 'L') {
					continue
				} else if len(strings.Split(info[i], " ")) == 3 {
					r = append(r, strings.Split(info[i], " ")[0])
				} else if len(strings.Split(info[i], "-")) == 2 {
					t = append(t, strings.Split(info[i], "-"))
				}
				if i+1 < len(info) {
					if info[i] == "##start" {
						colony.Start = strings.Split(info[i+1], " ")[0]
					} else if info[i] == "##end" {
						colony.End = strings.Split(info[i+1], " ")[0]
					}
				}
			}
		}
	}
	return r, t, colony
}

func Reader(arg string) ([]string, error) {
	file, err := os.Open("example/" + arg)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var info []string
	for scanner.Scan() {
		info = append(info, scanner.Text())
	}
	return info, nil
}
