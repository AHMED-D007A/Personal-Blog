// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Home() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Personal Blog</title><style>\n\t\t\t* {\n\t\t\t\tfont-family: sans-serif;\n\t\t\t}\n\t\t\tbody {\n\t\t\t\tmargin: 20px 100px;\n\t\t\t}\n\t\t\th1 {\n\t\t\t\tfont-size: 50px;\n\t\t\t}\n\t\t\tul {\n\t\t\t\tpadding: 0;\n\t\t\t\tmargin: 0;\n\t\t\t}\n\t\t\ta {\n\t\t\t\ttext-decoration: none;\n\t\t\t}\n\t\t\tli {\n\t\t\t\tlist-style: none;\n\t\t\t\tdisplay: flex;\n\t\t\t\tjustify-content: space-between;\n\t\t\t\tpadding: 0px 50px;\n\t\t\t\tmargin: 20px 0;\n\t\t\t\tbackground-color: #eee;\n\t\t\t}\n\t\t\t.title {\n\t\t\t\tfont-size: 25px;\n\t\t\t\tfont-weight: 600;\n\t\t\t\tcolor: #333;\n\t\t\t}\n\t\t\t.date {\n\t\t\t\tfont-size: 25px;\n\t\t\t\tfont-weight: 500;\n\t\t\t\tcolor: #aaa;\n\t\t\t}\n\t\t</style></head><body><h1>Personal Blog</h1><div><ul><a href=\"\"><li><p class=\"title\">title</p><p class=\"date\">August 7, 2024</p></li></a></ul></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
