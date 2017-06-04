package cafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func RequestCafe() {
	file, err := ioutil.ReadFile("json/meshi.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var me []menus
	jsonerr := json.Unmarshal(file, &me)
	if err != nil {
		fmt.Println("Format Error: ", jsonerr)
	}

	fmt.Println(me[1].Fish)
	fmt.Println(me[2].Fish)
	fmt.Println(me[3].Fish)
	fmt.Println(me[4].Fish)
	fmt.Println(me[0].Fish)
}
