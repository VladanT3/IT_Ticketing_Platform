// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
import "github.com/VladanT3/IT_Ticketing_Platform/models"

func UserView(user_type string, current_user models.Analyst, view_type string) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto mt-5\"><div class=\"flex flex-col\"><h1 class=\"text-2xl mb-3\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(view_type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_view.templ`, Line: 10, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><hr></div><div class=\"grid grid-cols-5 mt-3\"><div class=\"col-span-1 grid grid-cols-5\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if view_type == "Team View" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"col-span-4\" hx-post=\"/users/filter\" hx-trigger=\"keyup changed delay:500ms from:#search\" hx-target=\"#users\" hx-swap=\"innerHTML\"><input type=\"hidden\" name=\"view_type\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(view_type)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_view.templ`, Line: 17, Col: 62}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"flex flex-col mt-3\"><label>Search:</label> <input type=\"text\" name=\"search\" id=\"search\" class=\"uk-input text-zinc-50 text-base\"></div></form>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if view_type == "User View" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"col-span-4 flex flex-col\" hx-post=\"/users/filter\" hx-trigger=\"keyup changed delay:500ms from:#search\" hx-target=\"#users\" hx-swap=\"innerHTML\"><input type=\"hidden\" name=\"view_type\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(view_type)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_view.templ`, Line: 26, Col: 62}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"flex flex-col mt-3\"><label>Search:</label> <input type=\"text\" name=\"search\" id=\"search\" class=\"uk-input text-zinc-50 text-base\"></div><label class=\"mt-3\">Select user type:</label> <label><input type=\"radio\" name=\"user_type\" value=\"All\" class=\"uk-checkbox border rounded-full border-zinc-50\" checked> All</label> <label><input type=\"radio\" name=\"user_type\" value=\"Managers\" class=\"uk-checkbox border rounded-full border-zinc-50\"> Managers</label> <label><input type=\"radio\" name=\"user_type\" value=\"Adminstrators\" class=\"uk-checkbox border rounded-full border-zinc-50\"> Administrators</label></form>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"uk-divider-vertical min-h-full\"></div></div><div class=\"col-span-4 flex flex-col\" id=\"users\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if view_type == "User View" {
				templ_7745c5c3_Err = Users(models.GetAllAnalysts()).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if view_type == "Team View" {
				templ_7745c5c3_Err = Users(models.GetTeamsAnalysts(current_user.Team_ID.UUID.String())).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
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