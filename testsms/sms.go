package sms

/*
func sendsms() {
	// Kavenegar API setup
	api := kavenegar.New("426B73485959394B674E6861747950314B736659774E666234793148597A5170326D56654A6D6F374F32343D")
	sender := "10008663"

	// Fetch current date in YYYYMMDD format
	currentDate := time.Now().Format("20060102")
	date, err := strconv.ParseInt(currentDate, 10, 64)
	if err != nil {
		log.Fatal("Invalid date format: ", err)
	}

	// Get list of absent users until 9 AM
	absentUsers, err := db.GetAbsentUsersUntil9AM(context.Background(), date)
	if err != nil {
		log.Fatal("Error fetching absent users: ", err)
	}

	// Loop through absent users and send SMS
	for _, user := range absentUsers {
		receptor := []string{user.PhoneNumber}
		message := fmt.Sprintf("Hello %s %s, you are marked as absent. Please enter the attendance system.", user.FirstName, user.LastName)

		if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
			switch err := err.(type) {
			case *kavenegar.APIError:
				fmt.Println("API Error: ", err.Error())
			case *kavenegar.HTTPError:
				fmt.Println("HTTP Error: ", err.Error())
			default:
				fmt.Println("Other Error: ", err.Error())
			}
		} else {
			for _, r := range res {
				fmt.Println("MessageID =", r.MessageID)
				fmt.Println("Status    =", r.Status)
			}
		}
	}
}
*/
