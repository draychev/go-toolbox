// login is a package to make it easy
package login

/** Derek Sivers has a very simple login page
* Steps:
*  1. Enter your Name and Email
*  2. Check your Inbox for an email with a link - click it
*  3. You are logged in
**/

import (
	"fmt"
	"strings"
)

type Name string
type EmailAddress string
type Link string
type EmailBody string
type Token string

type LoginFacility interface {
	RenderLoginPage() string
	MakeTokenAndEmailLink(Name, EmailAddress, EmailDispatch)
	ValidateLoginLink(Link) Token
}

type EmailDispatch interface {
	SendEmail(Name, EmailAddress, EmailBody)
}

type login struct {
	loginURL         string
	token            string
	linkURLWithToken string
}

func NewLoginFacility(loginURL string) LoginFacility {
	// TODO initialize
	return &login{
		loginURL: loginURL,
	}
}

func (l *login) composeEmail() EmailBody {
	// TODO actually compose the email
	return EmailBody(fmt.Sprintf(`<a href="%s">%s</a>`, l.linkURLWithToken, l.linkURLWithToken))
}

func (l *login) MakeTokenAndEmailLink(nm Name, addr EmailAddress, dispatch EmailDispatch) {
	// TODO: actually create a new JWT (maybe)
	l.token = "some-token"
	l.linkURLWithToken = fmt.Sprintf("%s#token=%s", l.loginURL, l.token)
	emailBody := l.composeEmail()
	dispatch.SendEmail(nm, addr, emailBody)
}

func getLast(str string, sep string) string {
	chunks := strings.Split(str, sep)
	if len(chunks) > 1 {
		return chunks[len(chunks)-1]
	}
	return ""
}

func getToken(l Link) Token {
	str := string(l)
	splitOnHash := getLast(str, "#")
	splitOnEqual := getLast(splitOnHash, "=")
	return Token(splitOnEqual)
}

func (l *login) ValidateLoginLink(lnk Link) Token {
	// TODO this needs more
	return getToken(lnk)
}

func (l *login) RenderLoginPage() string {
	return fmt.Sprintf(`<section>
<form action="%s" method="post">
<h2>Login</h2>
<label for="name">Let's be friends! What's your name?</label>
<input type="text" id="name" name="name" placeholder="Your Name" required="">
<label for="email">To what email should I send a login link?</label>
<input type="email" id="email" name="email" placeholder="your@email.com" required="">
<button type="submit">enter</button>
</form>
</section>`, l.loginURL)
}
