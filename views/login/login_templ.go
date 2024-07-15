// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package login

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/VladanT3/IT_Ticketing_Platform/views/layouts"

func Login(emailErr string, passErr string, email string, pass string) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto\"><div class=\"grid grid-rows-3\"><div></div><div class=\"grid grid-cols-5\"><div></div><div></div><div><form method=\"POST\" action=\"/overview\" class=\"dark:bg-white-800 flex flex-col border rounded-lg border-zinc-600 border-dashed p-5\"><legend class=\"uk-legend text-zinc-300 text-2xl self-center font-mono\">Login</legend> <label class=\"text-zinc-300 font-mono\">Email</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if emailErr == "" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"uk-input dark:bg-zinc-600 text-base font-mono\" type=\"email\" name=\"email\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(email)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/login/login.templ`, Line: 18, Col: 108}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if emailErr != "" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"uk-input dark:bg-zinc-600 text-base border border-red-600 font-mono\" type=\"email\" name=\"email\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(email)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/login/login.templ`, Line: 20, Col: 130}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <label class=\"text-red-600 font-mono\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(emailErr)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/login/login.templ`, Line: 21, Col: 56}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"text-zinc-300 mt-3 font-mono\">Password</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if passErr == "" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"uk-input dark:bg-zinc-600 text-base font-mono\" id=\"passInput\" type=\"password\" name=\"password\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(pass)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/login/login.templ`, Line: 25, Col: 128}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if passErr != "" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"uk-input dark:bg-zinc-600 text-base border border-red-600 font-mono\" id=\"passInput\" type=\"password\" name=\"password\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(pass)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/login/login.templ`, Line: 27, Col: 150}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <label class=\"text-red-600 font-mono\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(passErr)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/login/login.templ`, Line: 28, Col: 55}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"text-zinc-300 font-mono\"><input class=\"uk-checkbox\" type=\"checkbox\" id=\"showPassword\"> Show password</label> <button class=\"uk-button uk-button-ghost bg-cyan-500 hover:bg-cyan-400 text-zinc-800 self-center mt-5 font-mono\">Log In</button></form></div><div></div><div></div></div><div></div></div></div><script>\n\t\t\tconst showPassword = document.getElementById('showPassword');\n\t\t\tconst passInput = document.getElementById('passInput');\n\n\t\t\tshowPassword.addEventListener(\"change\", function () {\n\t\t\t\tif (showPassword.checked)\n\t\t\t\t\tpassInput.type = 'text';\n\t\t\t\telse\n\t\t\t\t\tpassInput.type = 'password';\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
