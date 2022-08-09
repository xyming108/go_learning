package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func MacIntToStr(macInt int64) string {

	s := strings.ToUpper(strconv.FormatInt(macInt, 16))

	n := 12 - len(s)
	for n > 0 {
		s = "0" + s
		n--
	}

	return s
}

func MacStrToInt(macStr string) (int64, error) {
	macStr = strings.Replace(macStr, ":", "", -1)
	macStr = strings.Replace(macStr, "-", "", -1)

	if len(macStr) != 12 {
		return 0, errors.New("MAC字符串要为12字节")
	}

	return strconv.ParseInt(macStr, 16, 64)
}

func main() {
	toInt, err := MacStrToInt("E8EECC2F868C")
	if err != nil {
		return
	}
	toInt1, err := MacStrToInt("E8EECC2FD4AB")
	if err != nil {
		return
	}
	log.Println("=========0:", toInt)
	log.Println("=========1:", toInt1)
}
