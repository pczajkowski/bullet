# bullet
[![Go Report Card](https://goreportcard.com/badge/github.com/pczajkowski/bullet)](https://goreportcard.com/report/github.com/pczajkowski/bullet)
[![DeepSource](https://deepsource.io/gh/pczajkowski/bullet.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/pczajkowski/bullet/?ref=repository-badge)

**Initialize client**

	b := bullet.NewBullet("YOUR_TOKEN")

**Send note, use empty string in place of deviceID to send to all devices**

	err := b.SendNote("testTitle", "testText", "deviceID")
	if err != nil {
		log.Fatal(err)
	}

**Send link, use empty string in place of deviceID to send to all devices**

	err := b.SendLink("testTitle", "testText", "https://abc.com", "deviceID")
	if err != nil {
		log.Fatal(err)
	}

**Send file, use empty string in place of deviceID to send to all devices**

	err := b.SendFile("testFile", "Something", "./test.txt", "deviceID")
	if err != nil {
		log.Fatal(err)
	}

**List devices**

	devices, err := b.ListDevices()
	if err != nil {
		t.Error(err)
	}

**List all active pushes (modifiedAfter can be nil, limit <= 0 gives default of 500). Cursor is returned with every call (in Devices struct) and is used for pagination, so if you want to see next page of results in your next call provide cursor received from the last call.**

	devices, err := b.ListPushes(true, nil, 0, "")
	if err != nil {
		t.Error(err)
	}
