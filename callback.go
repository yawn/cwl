package cwl

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	callbackStream  = os.Stdout
	callbackEncoder = json.NewEncoder(callbackStream)
)

type Callback func(*Event)

func CallbackTabs(e *Event) {

	fmt.Fprintf(callbackStream, "%v\t%s\t%s\t%s\t%s\n",
		e.Timestamp,
		e.Region,
		e.Group,
		e.Stream,
		e.Message,
	)

}

func CallbackJSON(e *Event) {
	callbackEncoder.Encode(e)
}
