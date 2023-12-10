package inputs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	currentYear    = 2023
	aocURLtemplate = "https://adventofcode.com/%d/day/%d/input"
)

func Input(day int) []byte {
	d, err := fromCache(currentYear, day)
	if err != nil {
		o, err := online(currentYear, day)
		if err != nil {
			panic(fmt.Errorf("online(): %v", err))
		}
		if err := cache(currentYear, day, o); err != nil {
			panic(fmt.Errorf("cache(): %v", err))
		}
		return o
	}
	return d
}

func fromCache(year, day int) ([]byte, error) {
	return os.ReadFile(filename(year, day))
}

func cache(year, day int, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(filename(year, day)), 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("mkdir: %v", err)
	}
	return os.WriteFile(filename(year, day), data, 0600)
}

func filename(year, day int) string {
	h, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("get home dir: %v", err))
	}
	return filepath.Join(h, fmt.Sprintf("AoC%d", year), fmt.Sprintf("%d.txt", day))
}

var client http.Client

func online(year, day int) ([]byte, error) {
	fmt.Printf("--- Fetching inputs for day %d (%d) from adventofcode.com ---\n", day, year)
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, errors.New("env var AOC_SESSION not set - should contain AoC session from cookies")
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(aocURLtemplate, year, day), nil)
	if err != nil {
		return nil, fmt.Errorf("new req: %v", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client: %v", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body read: %v", err)
	}
	if got, want := resp.StatusCode, http.StatusOK; got != want {
		return nil, fmt.Errorf("aoc responsecode = %d, want %d", got, want)
	}
	if l := len(data); l == 0 {
		return nil, fmt.Errorf("response length = %d, want non-zero", l)
	}
	return data, nil
}

func Lines(i []byte) []string {
	var o []string
	for _, l := range bytes.Split(i[:len(i)-1], []byte("\n")) {
		o = append(o, string(l))
	}
	return o
}
