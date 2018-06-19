package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:8080")
	Scheme("http")
})

var BasicAuth = BasicAuthSecurity("basic_auth")

var _ = Resource("operands", func() {
	Security(BasicAuth)
	Action("get", func() {
		Routing(GET("/"))
		Payload(payload)
		Response(OK, "text/plain")
	})
})

var payload = Type("payload", func() {
	Attribute("sample", String)
})
