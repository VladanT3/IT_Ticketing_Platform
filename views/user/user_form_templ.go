// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
import "github.com/VladanT3/IT_Ticketing_Platform/models"

func UserForm(user_type string, analyst models.Analyst, view_type string, old_analyst models.Analyst, errs [4]bool) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto mt-5\"><form hx-put=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("/user/update/" + analyst.Analyst_ID.String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 9, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"grid grid-cols-5\"><input type=\"hidden\" name=\"view_type\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(view_type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 10, Col: 59}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <input type=\"hidden\" name=\"analyst_id\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(analyst.Analyst_ID.String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 11, Col: 78}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"col-span-1 grid grid-cols-5\"><div class=\"col-span-4 flex flex-col\"><button type=\"submit\" class=\"uk-button bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50\">Save</button> <button type=\"submit\" class=\"uk-button bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Save and Exit</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if view_type == "User View" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/user/view\" class=\"uk-button bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Exit</a> <button type=\"button\" class=\"uk-button bg-zinc-900 text-red-600 border border-red-600 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Delete User</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if view_type == "Team View" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/user/team/view\" class=\"uk-button bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Exit</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"uk-divider-vertical min-h-full\"></div></div><div class=\"col-span-4 grid grid-cols-3\"><div></div><div class=\"flex flex-col\"><div class=\"flex flex-row\"><div class=\"flex flex-col\" style=\"width: 45%;\"><label>First name:</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !errs[0] {
				if old_analyst.First_Name != "" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"first_name\" value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 string
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(old_analyst.First_Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 33, Col: 77}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"first_name\" class=\"uk-input text-base text-zinc-50\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"first_name\" class=\"uk-input text-base text-zinc-50 border border-red-600\"> <label class=\"text-red-600\">Please input a first name!</label>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div style=\"width: 10%;\"></div><div class=\"flex flex-col\" style=\"width: 45%;\"><label>Last name:</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !errs[1] {
				if old_analyst.Last_Name != "" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"last_name\" value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(old_analyst.Last_Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 47, Col: 75}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"last_name\" class=\"uk-input text-base text-zinc-50\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"last_name\" class=\"uk-input text-base text-zinc-50 border border-red-600\"> <label class=\"text-red-600\">Please input a last name!</label>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><label class=\"mt-3\">Email:</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !errs[2] {
				if old_analyst.Email != "" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"email\" value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(old_analyst.Email)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 60, Col: 65}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\"> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"email\" class=\"uk-input text-base text-zinc-50\"> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"email\" class=\"uk-input text-base text-zinc-50 border border-red-600\"> <label class=\"text-red-600\">Please input an email!</label> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"mt-3\">Phone number:</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !errs[3] {
				if old_analyst.Phone_Number != "" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"phone_number\" value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(old_analyst.Phone_Number)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 71, Col: 79}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"uk-input text-base text-zinc-50\"> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"phone_number\" class=\"uk-input text-base text-zinc-50\"> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"phone_number\" class=\"uk-input text-base text-zinc-50 border border-red-600\"> <label class=\"text-red-600\">Please input a phone number!</label> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"mt-3\">Team:</label> <select class=\"uk-select text-zinc-50\" name=\"team\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, team := range models.GetAllTeams() {
				if team.Team_ID == analyst.Team_ID.UUID {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option selected value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var10 string
					templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(team.Team_ID.String())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 83, Col: 55}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var11 string
					templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(team.Team_Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 83, Col: 74}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var12 string
					templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(team.Team_ID.String())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 85, Col: 46}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var13 string
					templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(team.Team_Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/user/user_form.templ`, Line: 85, Col: 65}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><div></div></div></form></div>")
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
