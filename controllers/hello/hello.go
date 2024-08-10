package hello

import (
	"encoding/json"
	"fmt"
	"hello/services/hello"
	"hello/utils"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request)  {
	times, _ := strconv.Atoi(utils.GetPathVar("times", r))

	fmt.Println("Request :: ", times)

	json.NewEncoder(w).Encode(hello.GenerateTimesLongMap(times))
}