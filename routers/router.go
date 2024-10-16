// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"customer_management_service/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.SetStaticPath("/uploads", "uploads")

	ns := beego.NewNamespace("/v1",

		// beego.NSNamespace("/user_types",
		// 	beego.NSInclude(
		// 		&controllers.UserTypesController{},
		// 	),
		// ),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/customer-categories",
			beego.NSInclude(
				&controllers.Customer_categoriesController{},
			),
		),

		beego.NSNamespace("/customers",
			beego.NSInclude(
				&controllers.CustomersController{},
			),
		),

		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&controllers.AccountsController{},
			),
		),

		beego.NSNamespace("/newsletter",
			beego.NSInclude(
				&controllers.Newsletter_customersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
