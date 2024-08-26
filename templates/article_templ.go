// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Article() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Personal Blog</title><style>\n\t\t\t* {\n\t\t\t\tfont-family: sans-serif;\n\t\t\t}\n\t\t\tbody {\n\t\t\t\tmargin: 20px 100px;\n\t\t\t}\n\t\t\th1 {\n\t\t\t\tfont-size: 60px;\n\t\t\t\tmargin: 10px 0;\n\t\t\t}\n\t\t\th3 {\n\t\t\t\tfont-size: 20px;\n\t\t\t\tmargin: 15px 0;\n\t\t\t\tcolor: #aaa;\n\t\t\t}\n\t\t\tp {\n\t\t\t\tfont-size: 25px;\n\t\t\t}\n\t\t</style></head><body><h1>title</h1><h3>August 7, 2024</h3><p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Labore et.</p></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
