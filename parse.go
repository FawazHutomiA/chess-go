package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ParseMove: "b2 b3", "b2,b3", "1,2 2,2"
func ParseMove(s string) (Pos, Pos, error) {
	s = strings.TrimSpace(s)
	parts := regexp.MustCompile(`[, ]+`).Split(s, -1)
	if len(parts) == 2 {
		from, err := parseCoord(parts[0])
		if err != nil {
			return Pos{}, Pos{}, fmt.Errorf("invalid from: %w", err)
		}
		to, err := parseCoord(parts[1])
		if err != nil {
			return Pos{}, Pos{}, fmt.Errorf("invalid to: %w", err)
		}
		return from, to, nil
	}
	return Pos{}, Pos{}, errors.New("cannot parse move, expected two coordinates")
}

func parseCoord(s string) (Pos, error) {
	s = strings.TrimSpace(s)
	if matched, _ := regexp.MatchString(`^[a-h][1-8]$`, s); matched {
		col := int(s[0] - 'a')
		rank := int(s[1] - '0')
		row := 8 - rank
		return Pos{R: row, C: col}, nil
	}
	if strings.Contains(s, ",") {
		parts := strings.Split(s, ",")
		if len(parts) != 2 {
			return Pos{}, errors.New("invalid numeric coord")
		}
		r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil || r < 1 || r > 8 {
			return Pos{}, errors.New("row out of range")
		}
		c, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil || c < 1 || c > 8 {
			return Pos{}, errors.New("col out of range")
		}
		return Pos{R: r - 1, C: c - 1}, nil
	}
	return Pos{}, errors.New("unknown coord format")
}
