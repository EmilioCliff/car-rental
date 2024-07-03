// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"strings"

	"github.com/EmilioCliff/car-rental/db"
)

func safariFormScript(id int) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_safariFormScript_d1c3`,
		Function: `function __templ_safariFormScript_d1c3(id){window.location.href=` + "`" + `/safari/${id}/contact-form` + "`" + `
	// // console.log(id)
	// fetch("/safari/contact", {
    //     method: "POST",
    //     body: JSON.stringify({
    //         id: id,
    //     }),
    //     headers: {
    //         "Content-Type": "application/json; charset=UTF-8"
    //     }
    // })
    // .catch((error) => {
    //     console.log(error)
    // })
}`,
		Call:       templ.SafeScript(`__templ_safariFormScript_d1c3`, id),
		CallInline: templ.SafeScriptInline(`__templ_safariFormScript_d1c3`, id),
	}
}

func SafariNumBody(safari db.Safari) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"service-header-bg\" id=\"home\"><p class=\"animate-pop-in\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d DAYS %s", safari.Days, safari.Title))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 29, Col: 83}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></section><p class=\"safari-description safari-num-intro animate-slide-in\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(safari.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 32, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><section class=\"safari-day-breakdown\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, day := range safari.Breakdown {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"safari-day\"><p class=\"safari-day-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("Day %d: %s", day.Day, day.Title))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 38, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"safari-card-item\"><img class=\"safari-day-card-img grid\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(string(templ.URL(day.Image)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 43, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\"></div><div class=\"safari-card-item\"><p class=\"safari-day-card-des\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(day.Activity)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 49, Col: 20}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, safariFormScript(int(safari.ID)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"btn safari-num-btn\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.ComponentScript = safariFormScript(int(safari.ID))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Book The Safari</div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</section><section class=\"safari-day-accom\"><svg class=\"section-animal-logo\" fill=\"var(--title-color)\" version=\"1.1\" id=\"Capa_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 505.087 505.087\" xml:space=\"preserve\"><path d=\"M488.909,418.392c-34.84-27.219-3.266-93.632-3.266-93.632c63.147-161.134-46.108-242.307-64.788-252.887\r\n\t\t\t\tc-38.033-21.541-128.497-42.648-174.166,15.319c-12.916,16.394-8.599,36.754-2.654,50.984\r\n\t\t\t\tc15.302,36.618,56.436,69.544,91.693,73.388c10.054,1.123,17.377-1.182,22.336-6.958c15.761-18.347,7.451-66.618,2.909-83.672\r\n\t\t\t\tc-1.233-4.644,1.531-9.416,6.184-10.658c1.341-0.356,2.664-0.3,3.938-0.041c3.142,0.637,5.842,2.909,6.72,6.225\r\n\t\t\t\tc1.905,7.179,17.956,70.998-6.532,99.501c-5.401,6.286-15.004,13.286-30.97,13.286c-2.05,0-4.21-0.119-6.473-0.366\r\n\t\t\t\tc-41.593-4.542-88.095-41.432-105.88-83.995c-4.475-10.708-6.707-21.065-6.948-30.835c-99.728,0-213.575,55.596-220.991,155.597\r\n\t\t\t\tc-0.238,3.208,1.735,6.167,4.789,7.179c3.054,1.021,6.413-0.169,8.14-2.882c6.124-9.616,13.75-20.027,21.325-26.585\r\n\t\t\t\tc1.087,15.749,5.871,30.796,14.136,44.532c16.552,27.508,25.39,59.098,25.39,91.198l0.158,59.253\r\n\t\t\t\tc0.843,7.261,17.995,13.065,39.181,13.065c21.194,0,38.346-5.804,39.19-13.065l0.158-38.455c0-10.096,4.21-19.742,11.61-26.616\r\n\t\t\t\tc7.4-6.873,17.316-9.628,27.413-9.628c6.263,0,19.079,0.009,19.828,0c9.688-0.101,19.019,4.337,25.899,11.159\r\n\t\t\t\tc6.89,6.822,10.761,16.11,10.761,25.799l0.158,37.74c0.843,7.261,17.994,13.065,39.181,13.065c21.194,0,38.346-5.804,39.19-13.065\r\n\t\t\t\tl0.158-33.215c0-98.599,72.254-142.598,80.839-147.814c42.606-25.886,23.68,52.557,23.68,52.557\r\n\t\t\t\tc-28.307,88.188,37.017,128.472,37.017,128.472l0.293-0.339c0.204,0.208,0.262,0.544,0.498,0.722\r\n\t\t\t\tc4.184,3.155,11.799,0.128,16.995-6.779c4.792-6.34,5.668-13.564,2.569-17.171L488.909,418.392z\"></path></svg><p class=\"section-subtitle safari-section-sub\">Neccessities</p><p class=\"section-title safari-section-tit\">SAFARI PRICES</p><div class=\"price-table-cont\"><h1>SAFARI PRICE GUIDE BY SEASON, NUMBER OF PERSON,LODGE?CAMP CATEGORY</h1><h2>Safari price (in USD) is per person in a private 4x4 pop-top Tour Jeep. <span>(Drag to view safari seasons)</span></h2><hr><div class=\"table-wrapper-car\"><div class=\"price-grab\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for season, value := range safari.Prices {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<table class=\"safari-price-table\"><thead><tr><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strings.ToUpper(string(season[0])) + season[1:])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 103, Col: 59}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" Season <br>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			switch season {
			case "shoulder":
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("(15th Nov-31stMar)")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			case "peak":
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("(1st Jul-15th Oct)&(23rd Dec-2nd Jan)")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			case "low":
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("(1st Apr - 30th June)")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>Economy</td><td>Comfort</td><td>Luxury</td></tr></thead> <tbody><tr><td>1 Person in a private tour</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(value.One.Economy)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 121, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(value.One.Comfort)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 122, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(value.One.Luxury)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 123, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr><tr><td>2-3 Person in a private tour</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(value.Two.Economy)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 127, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(value.Two.Comfort)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 128, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(value.Two.Luxury)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 129, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr><tr><td>4-5 Person in a private tour</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(value.Three.Economy)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 133, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(value.Three.Comfort)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 134, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 string
			templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(value.Three.Luxury)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 135, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr><tr><td>6-7 Person in a private tour</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(value.Four.Economy)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 139, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(value.Four.Comfort)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 140, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(value.Four.Luxury)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 141, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr><tr><td>Single Room Extra</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var21 string
			templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(value.Extra.Economy)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 145, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var22 string
			templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(value.Extra.Comfort)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 146, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var23 string
			templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(value.Extra.Luxury)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 147, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr><tr><td>Child below 3 Years <br>Child between 3-12 Years</td><td>FREE <br>75% of adult rate</td><td>FREE <br>75% of adult rate</td><td>FREE <br>75% of adult rate</td></tr><tr><td><span>NOTE:</span>Visitors staying in safari lodges and tented camps shall pay an extra USD 100 per person per day from 1stJuly - 31st December</td><td></td><td></td><td></td></tr></tbody></table>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"table-pips\"><div class=\"table-pip active\"></div><div class=\"table-pip\"></div><div class=\"table-pip\"></div></div></div></section><section class=\"safari-accomodation-sec\"><div class=\"accomo-sect-1\"><h1>ACCOMODATION</h1><h2>Safari accommodation options for each location</h2><div class=\"underline\"><div class=\"above-underline\"></div></div><div class=\"accomodation-options\"><div class=\"accomodation-option option-active\" data-accomodate-value=\"economy\"><p>ECONOMY LEVEL</p></div><div class=\"accomodation-option\" data-accomodate-value=\"comfort\"><p>COMFORT LEVEL</p></div><div class=\"accomodation-option\" data-accomodate-value=\"luxury\"><p>LUXURY LEVEL</p></div></div><div class=\"accomodation-list\" data-safari-id=\"1\"><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, value := range safari.Accomodation.Economy {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var24 string
			templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(value)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 323, Col: 18}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div><div class=\"tour-breakdown\"><h1>TOUR PRICE <span>INCLUDES</span></h1><div class=\"underline\"><div class=\"above-underline\"></div></div><div class=\"tour-col-2\"><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 0; i <= len(safari.TourIncludes)/2; i++ {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><i class=\"ri-edit-circle-fill\"></i><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var25 string
			templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(safari.TourIncludes[i])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 342, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := (len(safari.TourIncludes) / 2) + 1; i < len(safari.TourIncludes); i++ {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><i class=\"ri-edit-circle-fill\"></i><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var26 string
			templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs(safari.TourIncludes[i])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/safari-num.templ`, Line: 371, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Here</p></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div></section><section class=\"safari-not-included\"><h3>Not included In the Tour Price:</h3><ul><li>Bottled & Soft drinks</li><li>Items and services of a personal nature</li><li>Tips and Gratuities</li></ul></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func SafariNumPage(body templ.Component) templ.Component {
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
		templ_7745c5c3_Var27 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var27 == nil {
			templ_7745c5c3_Var27 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = BaseTemplate([]string{"general", "nav", "intro", "service", "footer", "safaris", "safaris-num"}, []string{"q&a", "nav", "safari-num"}, body).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
