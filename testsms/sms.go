package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.New("426B73485959394B674E6861747950314B736659774E666234793148597A5170326D56654A6D6F374F32343D")
	sender := "10008663"
	receptor := []string{"09371327163"}
	message := "Hello Go!"
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}
}
