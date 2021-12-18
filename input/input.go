package input

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

func ReadConfigFile(path string) (map[string]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, xerrors.Errorf("file open error: %w", err)
	}
	defer fp.Close()

	buf := make([]byte, fileSize())
	_, err = fp.Read(buf)
	if err != nil {
		return nil, xerrors.Errorf("file read error: %w", err)
	}

	checkbuf := make([]byte, 10)
	n, err := fp.Read(checkbuf)
	if err != nil {
		return nil, xerrors.Errorf("file read error: %w", err)
	}
	if n != 0 {
		return nil, xerrors.Errorf("file size over")
	}

	values := make(map[string]string, 0)
	if err = json.Unmarshal(buf, &values); err != nil {
		return nil, xerrors.Errorf("config file Marshal error: %w", err)
	}

	return values, nil
}

func fileSize() int {
	defaultSize := 20000
	sizeString := os.Getenv("CONFIG_FILE_SIZE")
	if sizeString == "" {
		return defaultSize
	}

	sizeInt, err := strconv.Atoi(sizeString)
	if err == nil {
		return sizeInt
	}

	sizeInt = withUnitToInt(sizeString)
	if sizeInt <= 0 {
		return defaultSize
	}

	return sizeInt
}

func withUnitToInt(withUnit string) int {
	var err error
	num := 0
	suf := ""

	withUnit = strings.ReplaceAll(withUnit, " ", "")
	withUnit = strings.ReplaceAll(withUnit, "　", "")
	withUnit = strings.ReplaceAll(withUnit, "\t", "")
	withUnit = strings.ReplaceAll(withUnit, "\n", "")

	for i := 0; i < 4; i++ {
		num, err = strconv.Atoi(withUnit[:len(withUnit)-i-1])
		if err != nil {
			continue
		}
		suf = withUnit[len(withUnit)-i-1:]
		break
	}

	switch suf {
	case "GiB", "gib", "GIB":
		g := 1 << 30
		return num * g
	case "MiB", "mib", "MIB":
		m := 1 << 20
		return num * m
	case "KiB", "kib", "KIB":
		k := 1 << 10
		return num * k
	case "GB", "gb":
		g := 1000000000
		return num * g
	case "MB", "mb":
		m := 1000000
		return num * m
	case "KB", "kb", "kB":
		k := 1000
		return num * k
	case "b", "B":
		return num
	}
	return 0
}
