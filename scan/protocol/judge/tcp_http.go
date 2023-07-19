package judge

import (
	"bytes"
	"ciscn/config"
	Conn "ciscn/scan/protocol/conn"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func TcpHTTP(result map[string]interface{}, Args map[string]interface{}) (map[string]interface{}, bool) {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)

	//var httpServerDetect common.HttpServerDetect
	//fmt.Println(len(embed.HttpServers))
	//fmt.Println(string(buff))
	//for _, server := range embed.HttpServers {
	//	pattern := common.FixRegexpPattenFormat(server.Pattern)
	//
	//	//fmt.Println(pattern)
	//	re := regexp.MustCompile(pattern)
	//	match := re.FindStringSubmatch(string(buff))
	//	if len(match) > 1 {
	//		serverVersion := server.VersionInfo
	//
	//		httpServerDetect.IsDetected = true
	//		serverName := fmt.Sprintf("%s/%s", serverVersion.Vendorproductname, serverVersion.Version)
	//		httpServerDetect.ServiceApp = append(httpServerDetect.ServiceApp, serverName)
	//
	//		if serverVersion.Devicetype != "" {
	//			httpServerDetect.Deviceinfo = append(httpServerDetect.Deviceinfo, serverVersion.Devicetype)
	//		}
	//	}
	//
	//}
	//fmt.Println("done")
	//
	//return httpServerDetect

	ok, err := regexp.Match(`^HTTP/\d.\d \d*`, buff)
	if err != nil {
		return nil, false
	}
	if ok {
		result["protocol"] = "http"
		httpResult, httpErr := httpIdentifyResult(result, Args)
		if httpErr != nil {
			result["banner.string"] = "None"
			return nil, true
		}
		result["banner.string"] = httpResult["http.title"].(string)
		u, err := url.Parse(httpResult["http.target"].(string))
		if err != nil {
			result["path"] = ""
		} else {
			result["path"] = u.Path
		}
		r := httpResult["http.result"].(string)
		c := fmt.Sprintf("[%s]", httpResult["http.code"].(string))
		if len(r) != 0 {
			result["identify.bool"] = true
			result["identify.string"] = r
			result["note"] = httpResult["http.target"].(string)
			return result, true
		} else {
			result["identify.bool"] = true
			result["identify.string"] = c
			result["note"] = httpResult["http.target"].(string)
			return result, true
		}
	}
	return result, false
}

func httpIdentifyResult(result map[string]interface{}, Args map[string]interface{}) (map[string]interface{}, error) {
	timeout := Args["FlagTimeout"].(int)
	var targetUrl string
	if Args["FlagUrl"].(string) != "" {
		targetUrl = Args["FlagUrl"].(string)
	} else {
		host := result["host"].(string)
		port := strconv.Itoa(result["port"].(int))
		add := net.JoinHostPort(host, port)
		if result["type"].(string) == "tcp" {
			if port == "80" {
				targetUrl = "http://" + host
			} else {
				targetUrl = "http://" + add
			}
		}
		if result["type"].(string) == "tls" {
			if port == "443" {
				targetUrl = "https://" + host
			} else {
				targetUrl = "https://" + add
			}
		}
	}

	var httpType string
	var httpCode string
	var httpResult string
	var httpUrl string
	var httpTitle string
	r, err := identify(targetUrl, timeout)
	if err != nil {
		return nil, err
	}
	for _, results := range r {
		httpType = results.Type
		httpCode = results.RespCode
		httpResult = results.Result
		httpUrl = results.Url
		httpTitle = results.Title
	}
	res := map[string]interface{}{
		"http.type":   httpType,
		"http.code":   httpCode,
		"http.result": httpResult,
		"http.target": httpUrl,
		"http.title":  httpTitle,
	}
	return res, nil
}

type RespLab struct {
	Url            string
	RespBody       string
	RespHeader     string
	RespStatusCode string
	RespTitle      string
	faviconMd5     string
}

func getFaviconMd5(Url string, timeout int) string {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	Url = Url + "/favicon.ico"
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return ""
	}
	for key, value := range config.DefaultHeader {
		req.Header.Set(key, value)
	}
	//req.Header.Set("Accept-Language", "zh,zh-TW;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6")
	//req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	//req.Header.Set("Cookie", "rememberMe=int")
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	hash := md5.Sum(bodyBytes)
	md5 := fmt.Sprintf("%x", hash)
	return md5
}

func defaultRequests(Url string, timeout int) ([]RespLab, error) {
	var redirectUrl string
	var respTitle string
	var responseHeader string
	var responseBody string
	var responseStatusCode string
	var res []string

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return nil, err
	}
	// set requests header
	for key, value := range config.DefaultHeader {
		req.Header.Set(key, value)
	}
	resp, err := Conn.ConnHttp(req, timeout)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// get response status code
	var statusCode = resp.StatusCode
	responseStatusCode = strconv.Itoa(statusCode)
	// -------------------------------------------------------------------------------
	// When the http request is 302 or other 30x,
	// Need to intercept the request and get the return status code for display,
	// Send the request again according to the redirect url
	// In the custom request, the return status code is not checked
	// -------------------------------------------------------------------------------
	if len(regexp.MustCompile("30").FindAllStringIndex(responseStatusCode, -1)) == 1 {
		redirectPath := resp.Header.Get("Location")
		if len(regexp.MustCompile("http").FindAllStringIndex(redirectPath, -1)) == 1 {
			redirectUrl = redirectPath
		} else {
			if Url[len(Url)-1:] == "/" {
				redirectUrl = Url + redirectPath
			}
			redirectUrl = Url + "/" + redirectPath
		}
		req, err := http.NewRequest("GET", redirectUrl, nil)
		if err != nil {
			return nil, err
		}
		for key, value := range config.DefaultHeader {
			req.Header.Set(key, value)
		}
		resp, err := Conn.ConnHttp(req, timeout)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		// Solve the problem of two 30x jumps
		var twoStatusCode = resp.StatusCode
		responseStatusCodeTwo := strconv.Itoa(twoStatusCode)
		if len(regexp.MustCompile("30").FindAllStringIndex(responseStatusCodeTwo, -1)) == 1 {
			redirectPath := resp.Header.Get("Location")
			if len(regexp.MustCompile("http").FindAllStringIndex(redirectPath, -1)) == 1 {
				redirectUrl = redirectPath
			} else {
				redirectUrl = Url + redirectPath
			}
			req, err := http.NewRequest("GET", redirectUrl, nil)
			if err != nil {
				return nil, err
			}
			for key, value := range config.DefaultHeader {
				req.Header.Set(key, value)
			}
			resp, err := Conn.ConnHttp(req, timeout)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			// get response body for string
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			responseBody = string(bodyBytes)
			// Solve the problem of garbled body codes with unmatched numbers
			if !utf8.Valid(bodyBytes) {
				data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bodyBytes)
				responseBody = string(data)
			}
			// Get Response title
			grepTitle := regexp.MustCompile("<title>(.*)</title>")
			if len(grepTitle.FindStringSubmatch(responseBody)) != 0 {
				respTitle = grepTitle.FindStringSubmatch(responseBody)[1]
			} else {
				respTitle = "None"
			}
			// get response header for string
			for name, values := range resp.Header {
				for _, value := range values {
					res = append(res, fmt.Sprintf("%s: %s", name, value))
				}
			}
			for _, re := range res {
				responseHeader += re + "\n"
			}
			faviconMd5 := getFaviconMd5(Url, timeout)
			RespData := []RespLab{
				{redirectUrl, responseBody, responseHeader, responseStatusCode, respTitle, faviconMd5},
			}
			return RespData, nil
		}
		// get response body for string
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		responseBody = string(bodyBytes)
		// Solve the problem of garbled body codes with unmatched numbers
		if !utf8.Valid(bodyBytes) {
			data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bodyBytes)
			responseBody = string(data)
		}
		// Get Response title
		grepTitle := regexp.MustCompile("<title>(.*)</title>")
		if len(grepTitle.FindStringSubmatch(responseBody)) != 0 {
			respTitle = grepTitle.FindStringSubmatch(responseBody)[1]
		} else {
			respTitle = "None"
		}
		// get response header for string
		for name, values := range resp.Header {
			for _, value := range values {
				res = append(res, fmt.Sprintf("%s: %s", name, value))
			}
		}
		for _, re := range res {
			responseHeader += re + "\n"
		}
		faviconMd5 := getFaviconMd5(Url, timeout)
		RespData := []RespLab{
			{redirectUrl, responseBody, responseHeader, responseStatusCode, respTitle, faviconMd5},
		}

		return RespData, nil
	}
	// get response body for string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)
	// Solve the problem of garbled body codes with unmatched numbers
	if !utf8.Valid(bodyBytes) {
		data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bodyBytes)
		responseBody = string(data)
	}
	// Get Response title
	grepTitle := regexp.MustCompile("<title>(.*)</title>")
	if len(grepTitle.FindStringSubmatch(responseBody)) != 0 {
		respTitle = grepTitle.FindStringSubmatch(responseBody)[1]
	} else {
		respTitle = "None"
	}
	// get response header for string
	for name, values := range resp.Header {
		for _, value := range values {
			res = append(res, fmt.Sprintf("%s: %s", name, value))
		}
	}
	for _, re := range res {
		responseHeader += re + "\n"
	}
	faviconMd5 := getFaviconMd5(Url, timeout)
	RespData := []RespLab{
		{Url, responseBody, responseHeader, responseStatusCode, respTitle, faviconMd5},
	}
	return RespData, nil
}

func customRequests(Url string, timeout int, Method string, Path string, Header []string, Body string) ([]RespLab, error) {
	var respTitle string
	// Splicing Custom Path
	u, err := url.Parse(Url)
	u.Path = path.Join(u.Path, Path)
	Url = u.String()
	if strings.HasSuffix(Path, "/") {
		Url = Url + "/"
	}
	// Send Http requests
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	bodyByte := bytes.NewBuffer([]byte(Body))
	req, err := http.NewRequest(Method, Url, bodyByte)
	if err != nil {
		return nil, err
	}

	// Set Requests Headers
	for _, header := range Header {
		grepKey := regexp.MustCompile("(.*): ")
		var headerKey = grepKey.FindStringSubmatch(header)[1]
		grepValue := regexp.MustCompile(": (.*)")
		var headerValue = grepValue.FindStringSubmatch(header)[1]
		req.Header.Set(headerKey, headerValue)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Get Response Body for string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var responseBody = string(bodyBytes)
	// Solve the problem of garbled body codes with unmatched numbers
	if !utf8.Valid(bodyBytes) {
		data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bodyBytes)
		responseBody = string(data)
	}
	// Get Response title
	grepTitle := regexp.MustCompile("<title>(.*)</title>")
	if len(grepTitle.FindStringSubmatch(responseBody)) != 0 {
		respTitle = grepTitle.FindStringSubmatch(responseBody)[1]
	} else {
		respTitle = "None"
	}
	// Get Response Header for string
	var res []string
	for name, values := range resp.Header {
		for _, value := range values {
			res = append(res, fmt.Sprintf("%s: %s", name, value))
		}
	}
	var responseHeader string
	for _, re := range res {
		responseHeader += re + "\n"
	}
	// get response status code
	var statusCode = resp.StatusCode
	responseStatusCode := strconv.Itoa(statusCode)
	RespData := []RespLab{
		{Url, responseBody, responseHeader, responseStatusCode, respTitle, ""},
	}
	return RespData, nil
}

type IdentifyResult struct {
	Type     string
	RespCode string
	Result   string
	Url      string
	Title    string
}

func identify(url string, timeout int) ([]IdentifyResult, error) {
	var DefaultFavicon string
	var CustomFavicon string
	var DefaultTarget string
	var CustomTarget string
	var Favicon string
	var RequestRule string
	var RespTitle string
	var RespBody string
	var RespHeader string
	var RespCode string
	var DefaultRespTitle string
	var DefaultRespBody string
	var DefaultRespHeader string
	var DefaultRespCode string
	var CustomRespTitle string
	var CustomRespBody string
	var CustomRespHeader string
	var CustomRespCode string

	R, err := defaultRequests(url, timeout)
	if err != nil {
		return nil, err
	}
	for _, resp := range R {
		DefaultRespBody = resp.RespBody
		DefaultRespHeader = resp.RespHeader
		DefaultRespCode = resp.RespStatusCode
		DefaultRespTitle = resp.RespTitle
		DefaultTarget = resp.Url
		DefaultFavicon = resp.faviconMd5
	}
	// start identify
	var succes_type string
	var identify_result string
	type Identify_Result struct {
		Name string
		Rank int
		Type string
	}
	var IdentifyData []Identify_Result
	for _, rule := range config.RuleData1 {
		if rule.Http.ReqMethod != "" {
			r, err := customRequests(url, timeout, rule.Http.ReqMethod, rule.Http.ReqPath, rule.Http.ReqHeader, rule.Http.ReqBody)
			if err != nil {
				return nil, err
			}

			for _, resp := range r {
				CustomRespBody = resp.RespBody
				CustomRespHeader = resp.RespHeader
				CustomRespCode = resp.RespStatusCode
				CustomRespTitle = resp.RespTitle
				CustomTarget = resp.Url
				CustomFavicon = resp.faviconMd5
			}
			url = CustomTarget
			Favicon = CustomFavicon
			RespBody = CustomRespBody
			RespHeader = CustomRespHeader
			RespCode = CustomRespCode
			RespTitle = CustomRespTitle
			// If the http request fails, then RespBody and RespHeader are both null
			// At this time, it is considered that the url does not exist
			if RespBody == RespHeader {
				continue
			}
			if rule.Mode == "" {
				if len(regexp.MustCompile("header").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkHeader(url, RespHeader, rule.Rule.InHeader, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "CustomRequest"
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("body").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkBody(url, RespBody, rule.Rule.InBody, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("ico").FindAllStringIndex(rule.Type, -1)) == 1 {
					if checkFavicon(Favicon, rule.Rule.InIcoMd5) == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name, Rank: rule.Rank, Type: rule.Type})
						succes_type = rule.Type
						continue
					}
				}
			}
			if rule.Mode == "or" {
				if len(regexp.MustCompile("header").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkHeader(url, RespHeader, rule.Rule.InHeader, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("body").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkBody(url, RespBody, rule.Rule.InBody, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("ico").FindAllStringIndex(rule.Type, -1)) == 1 {
					if checkFavicon(Favicon, rule.Rule.InIcoMd5) == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name, Rank: rule.Rank, Type: rule.Type})
						succes_type = rule.Type
						continue
					}
				}
			}
		} else { // Default Request Result
			url = DefaultTarget
			Favicon = DefaultFavicon
			RespBody = DefaultRespBody
			RespHeader = DefaultRespHeader
			RespCode = DefaultRespCode
			RespTitle = DefaultRespTitle
			// If the http request fails, then RespBody and RespHeader are both null
			// At this time, it is considered that the url does not exist
			if RespBody == RespHeader {
				continue
			}
			if rule.Mode == "" {
				if len(regexp.MustCompile("header").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkHeader(url, RespHeader, rule.Rule.InHeader, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "DefaultRequest"
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("body").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkBody(url, RespBody, rule.Rule.InBody, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "DefaultRequest"
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("ico").FindAllStringIndex(rule.Type, -1)) == 1 {
					if checkFavicon(Favicon, rule.Rule.InIcoMd5) == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "DefaultRequest"
						succes_type = rule.Type
						continue
					}
				}
			}
			if rule.Mode == "or" {
				if len(regexp.MustCompile("header").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkHeader(url, RespHeader, rule.Rule.InHeader, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "DefaultRequest"
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("body").FindAllStringIndex(rule.Type, -1)) == 1 {
					version, ok := checkBody(url, RespBody, rule.Rule.InBody, rule.Name, RespTitle, RespCode)
					if ok == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name + "/" + version, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "DefaultRequest"
						succes_type = rule.Type
						continue
					}
				}
				if len(regexp.MustCompile("ico").FindAllStringIndex(rule.Type, -1)) == 1 {
					if checkFavicon(Favicon, rule.Rule.InIcoMd5) == true {
						IdentifyData = append(IdentifyData, Identify_Result{Name: rule.Name, Rank: rule.Rank, Type: rule.Type})
						RequestRule = "DefaultRequest"
						succes_type = rule.Type
						continue
					}
				}
			}
		}
	}
	if RequestRule == "DefaultRequest" {
		RespBody = DefaultRespBody
		RespHeader = DefaultRespHeader
		RespCode = DefaultRespCode
		RespTitle = DefaultRespTitle
		url = DefaultTarget
	} else if RequestRule == "CustomRequest" {
		url = CustomTarget
		RespBody = CustomRespBody
		RespHeader = CustomRespHeader
		RespCode = CustomRespCode
		RespTitle = CustomRespTitle
	}
	for _, rs := range IdentifyData {
		switch rs.Rank {
		case 1:
			identify_result += "[" + (rs.Name) + "]"
		case 2:
			identify_result += "[" + (rs.Name) + "]"
		case 3:
			identify_result += "[" + (rs.Name) + "]"
		}
	}
	r := strings.ReplaceAll(identify_result, "][", "] [")
	res := []IdentifyResult{{succes_type, RespCode, r, url, RespTitle}}
	return res, nil
}

func checkHeader(url, responseHeader, ruleHeader, name, title, RespCode string) (string, bool) {
	grep := regexp.MustCompile("(?i)" + ruleHeader)
	match := grep.FindStringSubmatch(responseHeader)
	if len(match) > 1 {
		version := "N"
		if len(match) > 2 {
			version = match[2]
		}
		if version == "" {
			version = "N"
		}
		return version, true
	} else {
		return "N", false
	}
}

func checkBody(url, responseBody, ruleBody, name, title, RespCode string) (string, bool) {
	grep := regexp.MustCompile("(?i)" + ruleBody)
	if len(grep.FindStringSubmatch(responseBody)) != 0 {
		//fmt.Print("[body] ")
		return "N", true
	} else {
		return "N", false
	}
}

func checkFavicon(Favicon, ruleFaviconMd5 string) bool {
	grep := regexp.MustCompile("(?i)" + ruleFaviconMd5)
	if len(grep.FindStringSubmatch(Favicon)) != 0 {
		return true
	} else {
		return false
	}
}
