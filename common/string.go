package common

import (
	"regexp"
	"strings"
)

func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func FixRegexpPattenFormat(nmapServiceProbes string) string {
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, "${backquote}", "`")
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `GET / HTTP/1.0\r\n\r\n`,
		`GET / HTTP/1.0\r\nHost: baidu.com\r\nUser-Agent: Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)\r\nAccept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\r\nAccept: */*\r\n\r\n`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `\1`, `$1`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?=\\)`, `(?:\\)`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?=[\w._-]{5,15}\r?\n$)`, `(?:[\w._-]{5,15}\r?\n$)`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?:[^\r\n]*r\n(?!\r\n))*?`, `(?:[^\r\n]+\r\n)*?`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?:[^\r\n]*\r\n(?!\r\n))*?`, `(?:[^\r\n]+\r\n)*?`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?:[^\r\n]+\r\n(?!\r\n))*?`, `(?:[^\r\n]+\r\n)*?`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!2526)`, ``)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!400)`, ``)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!\0\0)`, ``)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!/head>)`, ``)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!HTTP|RTSP|SIP)`, ``)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!.*[sS][sS][hH]).*`, `.*`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!\xff)`, `.`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?!x)`, `[^x]`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?<=.)`, `(?:.)`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `(?<=\?)`, `(?:\?)`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `\x20\x02\x00.`, `\x20\x02..`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `match rtmp`, `# match rtmp`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `nmap`, `pamn`)
	nmapServiceProbes = strings.ReplaceAll(nmapServiceProbes, `Nmap`, `pamn`)
	return nmapServiceProbes
}

func MergeDuplicates(input []string) []string {
	seen := make(map[string]bool)
	merged := []string{}

	for _, str := range input {
		if seen[str] {
			continue
		}

		merged = append(merged, str)

		seen[str] = true
	}

	return merged
}
func ExtractStrings(input string) []string {
	pattern := `\[(.*?)\]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)
	var result []string

	for _, match := range matches {
		if len(match) > 1 {
			result = append(result, match[1])
		}
	}

	return result
}

func SplitAddressPort(input string) (string, string) {
	parts := strings.Split(input, ":")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return input, ""
}
