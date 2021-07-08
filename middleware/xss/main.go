package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	p := bluemonday.NewPolicy()

	//p := bluemonday.StrictPolicy()

	// Require URLs to be parseable by net/url.Parse and either:
	//   mailto: http:// or https://
	p.AllowStandardURLs()

	// We only allow <p> and <a href="">
	p.AllowAttrs("href").OnElements("a")
	p.AllowElements("p")

	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com">Google</a>
	fmt.Println(html)
}
