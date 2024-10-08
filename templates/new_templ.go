// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func New() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Personal Blog</title><style>\n\t\t\t* {\n\t\t\t\tfont-family: sans-serif;\n\t\t\t}\n\t\t\tbody {\n\t\t\t\tmargin: 20px 100px;\n\t\t\t}\n\t\t\th1 {\n\t\t\t\tfont-size: 50px;\n\t\t\t}\n\t\t\tform {\n\t\t\t\tdisplay: flex;\n\t\t\t\tflex-direction: column;\n\t\t\t}\n\t\t\tinput {\n\t\t\t\tfont-size: 20px;\n\t\t\t\tmargin: 20px 0;\n\t\t\t\twidth: 60%;\n\t\t\t}\n\t\t\ttextarea {\n\t\t\t\tfont-size: 20px;\n\t\t\t\tmargin: 20px 0;\n\t\t\t\twidth: 60%;\n\t\t\t}\n\t\t</style></head><body><h1>New Article</h1><form action=\"\" method=\"post\"><input type=\"text\" id=\"title\" name=\"title\" placeholder=\"Article Title\" required> <input type=\"date\" id=\"publish_date\" name=\"publish_date\" placeholder=\"Publishing Date\" required> <textarea id=\"content\" name=\"content\" rows=\"10\" cols=\"50\" placeholder=\"Content\" required></textarea> <input type=\"submit\" value=\"Publish\"></form></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
