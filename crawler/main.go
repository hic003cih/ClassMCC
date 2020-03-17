package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}

/*
func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun"
		ParserFunc: parser.ParseCityList
	})
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error: status code", resp.StatusCode)

		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Println("%s\n", all)

	printCityList(all)
}
*/
/* func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
} */
/*
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
		// fmt.Printf("%s\n", m)
		// for _, subMatch := range m {
		//	fmt.Printf("%s", subMatch)
		//}
		fmt.Println()
	}

	fmt.Printf("Matches found: %d\n", len(matches))
}
*/
