// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/VladanT3/IT_Ticketing_Platform/views/layouts"

func ChangePasswordForm(user_type string, old_password string, old_repeat_password string, diff_pass_err bool) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto mt-5 grid grid-cols-3\"><div></div><div><form class=\"flex flex-col border rounded-lg border-zinc-50 p-5\" action=\"/user/password/change\" method=\"post\"><legend class=\"uk-legend text-2xl self-center\">Change password</legend> <label>New password:</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !diff_pass_err {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input required type=\"password\" name=\"password\" id=\"password\" autofocus value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(old_password)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/change_password_form.templ`, Line: 14, Col: 98}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input required type=\"password\" name=\"password\" id=\"password\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(old_password)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/change_password_form.templ`, Line: 16, Col: 88}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label><input class=\"uk-checkbox border rounded-full border-zinc-50\" type=\"checkbox\" id=\"show_password\"> Show password</label> <label class=\"mt-3\">Repeat password:</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !diff_pass_err {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input required type=\"password\" name=\"repeat_password\" id=\"repeat_password\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(old_repeat_password)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/change_password_form.templ`, Line: 21, Col: 109}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input autofocus required type=\"password\" name=\"repeat_password\" id=\"repeat_password\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(old_repeat_password)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/change_password_form.templ`, Line: 23, Col: 119}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50 border border-red-600\"> <label class=\"text-red-600\">Passwords need to be the same!</label> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label><input class=\"uk-checkbox border rounded-full border-zinc-50\" type=\"checkbox\" id=\"show_repeat_password\"> Show password</label> <button type=\"submit\" class=\"uk-button bg-zinc-50 text-zinc-900 border border-zinc-50 self-center mt-3 hover:bg-zinc-900 hover:text-zinc-50\">Change</button></form></div><div></div></div><script>\n\t\t\tconst show_password = document.getElementById('show_password');\n\t\t\tconst pass_input = document.getElementById('password');\n\t\t\tconst show_repeat_password = document.getElementById('show_repeat_password');\n\t\t\tconst repeat_pass_input = document.getElementById('repeat_password');\n\n\t\t\tshow_password.addEventListener(\"change\", function() {\n\t\t\t\tif (show_password.checked)\n\t\t\t\t\tpass_input.type = 'text';\n\t\t\t\telse\n\t\t\t\t\tpass_input.type = 'password';\n\t\t\t});\n\t\t\tshow_repeat_password.addEventListener(\"change\", function() {\n\t\t\t\tif (show_repeat_password.checked)\n\t\t\t\t\trepeat_pass_input.type = 'text';\n\t\t\t\telse\n\t\t\t\t\trepeat_pass_input.type = 'password';\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Navbar(user_type).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
