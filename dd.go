package split_github_io

import (
	"fmt"
	"net/http"
)

func main(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "fdsfds")
}
