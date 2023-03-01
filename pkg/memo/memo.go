package memo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadMemo(title string) (result []byte) {
	memoPath := os.Getenv("MEMOPATH")
	data, err := os.Open(filepath.Join(memoPath, title))
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	d, _ := ioutil.ReadAll(data)
	return d
}

func CreateMemo(title string) (result string) {
	memoPath := os.Getenv("MEMOPATH")
	filename := fmt.Sprintf("%s%s", title, ".txt")
	f, e := os.Create(filepath.Join(memoPath, filename))
	if e != nil {
		panic(e)
	}
	defer f.Close()
	return filename
}

func FilterMemoEntries(entries []string, keyword string) (result []string) {
	for _, e := range entries {
		if strings.Contains(e, keyword) {
			result = append(result, e)
		}
	}
	return result
}

func ReadMemoEntries() (result []string) {
	memoPath := os.Getenv("MEMOPATH")
	entries, err := os.ReadDir(memoPath)

	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		result = append(result, e.Name())
	}
	return result
}
