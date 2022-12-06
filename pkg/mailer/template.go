package pkg

func templateGender(title, body string, templateId int) string {
	var template string

	switch templateId {
	case 1:
		template =
			` <div style="background-color: #f5f5f5; padding: 10px">
      <div
        style="
          padding: 1rem;
          background-color: white;
          width: 100%;
          max-width: 500px;
          margin: auto;
        "
      >
        <div style="margin: 0 0 50px 0">
          <img
            src="https://res.cloudinary.com/myapp12091999/image/upload/v1668392660/xp2mdx78d0ti7vgylzdi.png"
            width="100px"
          />
        </div>
        <div style="margin: 0 0 50px 0">Hi,` + title + `</div>
        <div>
          Thank you for creating a Poly Career account. For your security,
          please verify your account by clicking the button below.
        </div>
        <div style="margin: 50px 0;text-align:center;">
          <a
            style="background-color: #6667ab; color: white;padding:10px;text-decoration:none;"
            href="` + body + `"
            >Verify my account</a
          >
        </div>
        <div style="margin: 50px 0">
          Questions? Need help? Please visit Poly Career Support. Happy
        </div>
        <div style="margin: 50px 0">
          Thanks,<br />
          Poly Career
        </div>
      </div>
    </div>
    `
	case 2:
		template =
			` <div style="background-color: #f5f5f5; padding: 10px">
      <div
        style="
          padding: 1rem;
          background-color: white;
          width: 100%;
          max-width: 500px;
          margin: auto;
        "
      >
        <div style="margin: 0 0 50px 0">
          <img
            src="https://res.cloudinary.com/myapp12091999/image/upload/v1668392660/xp2mdx78d0ti7vgylzdi.png"
            width="100px"
          />
        </div>
        <div style="margin: 0 0 50px 0">
        You are requesting to change the email to: ` + title + `. For your security,
          please verify your account by clicking the button below.
        </div>
     
        <div style="margin: 50px 0;text-align:center;">
          <a
            style="background-color: #6667ab; color: white;padding:10px;text-decoration:none;"
            href="` + body + `"
            >Confirm change email</a
          >
        </div>
        <div style="margin: 50px 0">
          Questions? Need help? Please visit Poly Career Support. Happy
        </div>
        <div style="margin: 50px 0">
          Thanks,<br />
          Poly Career
        </div>
      </div>
    </div>
    `
	}

	return template
}
