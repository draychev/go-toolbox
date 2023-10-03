package login

import (
	"fmt"
	"testing"
)

const loginURL = "https://example.com/"

var l = NewLoginFacility(loginURL)

type emailSender struct {
	nm   Name
	addr EmailAddress
	body EmailBody
}

func (s *emailSender) SendEmail(nm Name, addr EmailAddress, body EmailBody) {
	s.nm = nm
	s.addr = addr
	s.body = body
}

func TestSendEmailLink(t *testing.T) {
	es := emailSender{}
	l.MakeTokenAndEmailLink(Name("a"), EmailAddress("b"), &es)

	{
		expected := Name("a")
		if es.nm != expected {
			t.Errorf("Name is expected to be %+v; it is actually %+v", expected, es.nm)
		}
	}

	{
		expected := EmailAddress("b")
		if es.addr != expected {
			t.Errorf("Email is expected to be %+v; it is actually %+v", expected, es.addr)
		}
	}

	{
		expected := EmailBody(`<a href="https://example.com/#token=some-token">https://example.com/#token=some-token</a>`)
		if es.body != expected {
			t.Errorf("Email body is expected to be '%+v'; it is actually '%+v'", expected, es.body)
		}
	}
}

func TestGetToken(t *testing.T) {
	{
		actualToken := getToken("https://example.com/#token=xyz")
		expected := Token("xyz")
		if actualToken != expected {
			t.Errorf("getToken() was expected to return %s; we got %s", expected, actualToken)
		}
	}

	{
		actualToken := getToken("https://example.com/?token=xyz")
		expected := Token("")
		if actualToken != expected {
			t.Errorf("getToken() was expected to return %s; we got %s", expected, actualToken)
		}
	}
}

func TestValidateLoginLink(t *testing.T) {
	{
		actualToken := l.ValidateLoginLink(Link("https://example.com/#token=xyz"))
		expected := Token("xyz")
		if actualToken != expected {
			t.Errorf("ValidateLoginLink() was expected to return %s; we got %s", expected, actualToken)
		}
	}
}

func TestRenderLoginPage(t *testing.T) {
	// Expected rendered page
	expectedResult := fmt.Sprintf(`<section>
<form action="%s" method="post">
<h2>Login</h2>
<label for="name">Let's be friends! What's your name?</label>
<input type="text" id="name" name="name" placeholder="Your Name" required="">
<label for="email">To what email should I send a login link?</label>
<input type="email" id="email" name="email" placeholder="your@email.com" required="">
<button type="submit">enter</button>
</form>
</section>`, loginURL)
	// Call the function with the test URL
	result := l.RenderLoginPage()

	// Compare the result with the expected output
	if result != expectedResult {
		t.Errorf("RenderLoginPage() returned unexpected result, got: %s, want: %s", result, expectedResult)
	}
}
