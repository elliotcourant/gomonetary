package monetary

import (
	"os"
	"path/filepath"
	"strings"
)

type LocaleName string

type LocaleNames []LocaleName

func GetSupportedLocales() LocaleNames {
	files := make([]LocaleName, 0)
	// Locales are stored in a folder in the file system on unix and linux. I don't know yet
	// where or how its stored on windows.
	// But essentially we can walk the directory it's in and check to see what locales
	// are present with LC_MONETARY and then use those.
	if err := filepath.Walk(localePath, func(path string, info os.FileInfo, err error) error {
		// The monetary suffix is what tells us which "file" stores the monetary info.
		// There are others that have LC_{info} categories. We can check right away to
		// make sure that it is the right type.
		if !strings.HasSuffix(path, localeMonetarySuffix) {
			return nil
		}
		name := path
		// Since this path has the monetary category suffix we need to strip out the rest
		// of the path.
		name = strings.TrimPrefix(name, localePath)
		name = strings.TrimSuffix(name, localeMonetarySuffix)
		files = append(files, LocaleName(name))
		return nil
	}); err != nil {
		// If there is an error trying to walk the directory, then we just want to panic.
		panic(err)
	}
	return files
}
