package main

import (
	"advent2023/utils"
	"strconv"
	"strings"
)

type Almanac struct {
	maps []FilterSet
}

type FilterSet struct {
	filters []Filter
}

type Filter struct {
	src int
	dst int
	length int 
}

func BuildAlmanac(file_path string) Almanac {
	almanac := Almanac{maps: []FilterSet{}}

	lines, _ := utils.IterLines(file_path)
	filter_set := FilterSet{filters: []Filter{}}
	for currentLine := range lines {
		if currentLine == "" {
			if len(filter_set.filters) > 0 {
				almanac.maps = append(almanac.maps, filter_set)
				filter_set = FilterSet{filters: []Filter{}}
			}
		} else if !strings.Contains(currentLine, ":") {
			split := strings.Split(currentLine, " ")
			dst_start, _ := strconv.Atoi(split[0])
			src_start, _ := strconv.Atoi(split[1])
			range_len, _ := strconv.Atoi(split[2])
			filter_set.filters = append(filter_set.filters, Filter{src:src_start, dst: dst_start, length: range_len})
		}
	}

	return almanac
}

func FilterNumber(input int, almanac Almanac) int {
	output := input
	for _, filter_set := range almanac.maps {
		for _, filter := range filter_set.filters {
			if output >= filter.src && output < filter.src + filter.length {
				output = output - filter.src + filter.dst
				break
			}
		}
	}
	return output
}
