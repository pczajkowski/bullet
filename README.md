# bullet
[![Go Report Card](https://goreportcard.com/badge/github.com/pczajkowski/bullet)](https://goreportcard.com/report/github.com/pczajkowski/bullet)

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


