// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package ticket

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"strconv"
)

func Ticket(ticket models.Ticket, userType string, mode string) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto mt-5\"><div class=\"grid grid-cols-5\"><div class=\"col-span-1 grid grid-cols-5\"><div class=\"flex flex-col col-span-4\"><button class=\"uk-button uk-width-1-1 bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50\">Save</button> <button class=\"uk-button uk-width-1-1 bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Save and Exit</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if mode == "create" {
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"uk-button uk-width-1-1 bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Close Ticket</button> <button class=\"uk-button uk-width-1-1 bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Reopen</button> <button class=\"uk-button uk-width-1-1 bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Assign to Me</button> <button class=\"uk-button uk-width-1-1 bg-zinc-50 text-zinc-900 border border-zinc-50 hover:bg-zinc-900 hover:text-zinc-50 mt-3\">Assign Ticket</button> <button class=\"uk-button uk-width-1-1 bg-zinc-900 text-red-600 border border-red-600 hover:bg-red-600 hover:text-zinc-900 mt-3\">Delete Ticket</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"uk-divider-vertical min-h-full\"></div></div><div class=\"col-span-4\"><form class=\"flex flex-col\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if mode != "create" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid grid-cols-4\"><div class=\"col-span-3\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if ticket.Type == "Request" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-xl\">REQ/")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(ticket.Ticket_Number))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 34, Col: 72}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-xl\">INC/")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 string
					templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(ticket.Ticket_Number))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 36, Col: 72}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><span class=\"border border-zinc-50 rounded-lg bg-zinc-50 text-zinc-900\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Status)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 40, Col: 96}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div><hr class=\"mt-5 mb-3\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid grid-cols-4\"><div><label>Ticket type:</label> <select class=\"uk-select text-zinc-50\"><option value=\"none\">Select...</option> <option>Request</option> <option>Incident</option></select></div><div></div><div><label>Customer contact:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\"></div></div><div class=\"grid grid-cols-4 mt-3\"><div><label>Category:</label> <select name=\"category\" class=\"uk-select text-zinc-50\" hx-get=\"/getSubcategories\" hx-target=\"#subcategories\" hx-swap=\"outerHTML\" hx-trigger=\"change\"><option value=\"none\">Select...</option> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, category := range models.GetAllCategories() {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(category.Category_ID.String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 66, Col: 55}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(category.Category_Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 66, Col: 82}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><div></div><div><label>Subcategory:</label> <select class=\"uk-select text-zinc-50\" id=\"subcategories\" disabled><option value=\"none\">Please select category</option></select></div></div><div class=\"mt-3\"><label>Title:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\"></div><div class=\"mt-3\"><label>Description:</label> <textarea class=\"uk-textarea font-mono text-zinc-50 text-base\" rows=\"10\"></textarea></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if mode != "create" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"mt-5\"><div class=\"grid grid-cols-4 mt-3\"><div class=\"col-span-1\"><label>Last Updated:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Updated_at.Format("Mon 02/01 15:04 2006"))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 91, Col: 133}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"grid grid-cols-4 mt-3\"><div><label>Assigned Team:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"Assigned Team\"></div><div></div><div><label>Assigned Analyst:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"Assigned Analyst\"></div></div><div class=\"grid grid-cols-4\"><div><label>Opened by:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"Opened by\"></div><div></div><div><label>Opened:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Opened_Date.Format("Mon 02/01 15:04 2006"))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 113, Col: 134}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"grid grid-cols-4\"><div><label>Closed by:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"Closed by\"></div><div></div><div><label>Closed:</label> <input class=\"uk-input text-zinc-50 text-base\" type=\"text\" disabled value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(ticket.Closed_Date.Time.Format("Mon 02/01 15:04 2006"))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/ticket/ticket_form.templ`, Line: 124, Col: 139}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Navbar(userType).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
