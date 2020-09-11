# bullet
[![Go Report Card](https://goreportcard.com/badge/github.com/pczajkowski/bullet)](https://goreportcard.com/report/github.com/pczajkowski/bullet)

**Initialize client**

	b := bullet.NewBullet("YOUR_TOKEN")

**Send note**

	err := b.SendNote("testTitle", "testText")
	if err != nil {
		log.Fatal(err)
	}

**Send link**

	err := b.SendLink("testTitle", "testText", "https://abc.com")
	if err != nil {
		log.Fatal(err)
	}

**Send file**

	err := b.SendFile("testFile", "Something", "./test.txt")
	if err != nil {
		log.Fatal(err)
	}


