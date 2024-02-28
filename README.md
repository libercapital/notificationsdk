# Welcome to notification-sdk 👋

[![pipeline status](https://gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk/badges/main/pipeline.svg)](https://gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk/-/commits/main) [![coverage report](https://gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk/badges/main/coverage.svg)](https://gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk/-/commits/main)

## Getting started

Responsible for encapsulating calls to the Notification service

## Install

```bash
go get github.com/libercapital/notificationsdk
```

## How to use

```golang
client := notificationsdk.NewClient(notificationsdk.Config{URL: "https://api.xpto.com/"})

client.SendEmail(ctx, "access-token", notificationsdk.EmailRequest{
		To: notificationsdk.Address{
			Name:  "to.Name",
			Email: "to.Address",
		},
		From: notificationsdk.Address{
			Name:  "XPTO",
			Email: "john@email.com",
		},
		Template: "sendgrid-template",
		MetaData: map[string]interface{}{},
		Attachments: []notificationsdk.Attachment{
			{
				FileName:    "Arquivo.pdf",
				FileType:    "application/pdf",
				FileContent: "ASNFZ4mrze8BI0VniavN7w==",
			},
		},
	})
```

## Author

## Contributors

👤 **Eduardo Mello**

- Github: [@EduardoRMello](https://github.com/EduardoRMello)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
