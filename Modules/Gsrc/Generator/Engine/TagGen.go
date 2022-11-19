package Engine

import "fmt"

func Generate_Box(title, data string) string {
	var TemplateBox = `
	<div class="box">
		<div class="right-side">
			<div class="box-topic"><a href="/Server_Html/DNS">%s</div>
			<hr><br></a>
			<div class="number">%s</div>
		</div>
	</div>
	`
	TemplateBox = fmt.Sprintf(TemplateBox, title, data)
	return TemplateBox
}

func Generate_Div(divname string) string { return fmt.Sprintf(`<div class="%s">`, divname) }
func Generate_Sec(secname string) string { return fmt.Sprintf(`<section class="%s">`, secname) }
func Generate_CH(secdata ...string) string {
	return fmt.Sprintf(`<div class="codeheader" id="%s">%s</div>`, secdata[0], secdata[1])
}
func Generate_CS(secdata ...string) string {
	return fmt.Sprintf(`<pre class="%s"> %s </pre>`, secdata[0], secdata[1])
}
