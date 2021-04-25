package examples

import (
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

/*
 * Let's supose that we receive the query ?all[]=1&all[2]&all[3]&b=1 in the following HttpHandler func.
 *
 * The url.ParseQuery will parse all params to a map[string][]string, typed as url.Values and
 * so all the keys in the query will be treated as an slice
 *
 */
func ExampleGetAllQueryParametersHttpHandler(_ http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	values, _ := url.ParseQuery(req.URL.RawQuery)

	for key, value := range values {
		// key is string
		// value is a slice of strings
		log.Println(key, value)

		// so, in order to get the first item (whatever it is slice/array in query string, you have to access as slice/array)
		log.Println(value[0])
	}

	// will output on the terminal:
	// all[] [1, 2, 3]
	// b [1]
}
