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

func Generate_Div(divname string) string     { return fmt.Sprintf(`<div class="%s">`, divname) }
func Generate_Sec(secname string) string     { return fmt.Sprintf(`<section class="%s">`, secname) }
func Generate_NavBegin() string              { return "<nav>" }
func Generate_NavEnd() string                { return "</nav>" }
func Generate_Table(tablename string) string { return fmt.Sprintf(`<table class="%s"><br>`, tablename) }
func Generate_HRBR() string                  { return "<hr><br>" }
func Generate_THEAD() string                 { return "<thead><tr>" }
func Generate_THEADEND() string              { return "</thead><tbody>" }
func Generate_TBLEEND() string               { return "</tbody></table>" }
func Generate_H1START() string               { return "<h1>" }
func Generate_H1END() string                 { return "</h1>" }
func Generate_TH() string                    { return "<th>" }
func Generate_THE() string                   { return "</th>" }
func Generate_TD() string                    { return "<td>" }
func Generate_TDE() string                   { return "</td>" }
func Generate_STG() string                   { return "<style>" }
func Generate_STGE() string                  { return "</style>" }
func Generate_CH(secdata ...string) string {
	return fmt.Sprintf(`<div class="codeheader" id="%s">%s</div>`, secdata[0], secdata[1])
}
func Generate_CS(secdata ...string) string {
	return fmt.Sprintf(`<pre class="%s"> %s </pre>`, secdata[0], secdata[1])
}

func EmbednewStyle(stylename string, values ...string) string {
	var temple string
	temple += Generate_STG()
	temple += "." + stylename + "{"
	for i := range values {
		temple += values[i]
	}
	temple += "}"
	temple += Generate_STGE()
	return temple
}
