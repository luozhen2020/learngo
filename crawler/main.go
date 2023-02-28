package main

import (
	"xyy/learngo/crawler/engine"
	"xyy/learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        `https://www.zhenai.com/zhenghun`,
		ParserFunc: parser.ParseCityList,
	})
}

/*const testProfileUrl = `<a href="http://album.zhenai.com/u/1932439817" target="_blank">温馨昊歌</a>`

func main() {
	resp, err := http.Get(`http://album.zhenai.com/u/1932439817`)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code", resp.StatusCode)
		return
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
*/
