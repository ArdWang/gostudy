package main
// gopm get -g -v golang.org/x/text
import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	//"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	 resp, err := http.Get("https://www.zhenai.com/zhenghun")

	 if err != nil {
	 	panic(err)
	 }

	 defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
	 	fmt.Println("Error : status code",resp.StatusCode)
		 return
	 }

	 //all, err := ioutil.ReadAll(resp.Body)

	 e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)

	 if err != nil {
	 	panic(err)
	 }

	 fmt.Printf("%s\n", all)

}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _,_ := charset.DetermineEncoding(bytes,"")

	return e

	//charset.DetermineEncoding(bytes)
}



//备份代码
//transform.NewReader(res)
/*


 utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

  all, err := ioutil.ReadAll(utf8Reader)

  if err != nil {
  	panic(err)
  }

  fmt.Printf("%s\n", all)

*/

