package frontmatter

import (
	"regexp"

	"gopkg.in/yaml.v2"
)

const fmRegex = `(?ms)^\s*---.*---$`

// Unmarshal parses the data containing YAML frontmatter and stores YAML encoded data in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) (rest []byte, err error) {
	regex := regexp.MustCompile(fmRegex)
	fmLoc := regex.FindIndex(data)
	if fmLoc == nil {
		return data, nil
	}
	if err = yaml.Unmarshal(data[fmLoc[0]:fmLoc[1]], v); err != nil {
		return []byte{}, err
	}
	if len(data) <= fmLoc[1]+1 {
		return
	}
	return data[fmLoc[1]+1:], nil
}

// Marshal returns the frontmatter encoding of v.
func Marshal(v interface{}) (data []byte, err error) {
	if data, err = yaml.Marshal(&v); err != nil {
		return data, err
	}
	data = append([]byte("---\n"), data...)
	data = append(data, []byte("---\n")...)
	return data, nil
}
