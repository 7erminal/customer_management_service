package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_categoriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "AddCustomer",
            Router: `/add-customer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Newsletter_customersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Role_permissionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "GetAllBranches",
            Router: `/branches`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "SignUp2",
            Router: `/2/sign-up`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "VerifyUsername",
            Router: `/get-user-by-username/:username`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "InviteUser",
            Router: `/invite-user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "SignUp",
            Router: `/sign-up`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "VerifyInvite",
            Router: `/verify-invite`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "VerifyUser",
            Router: `/verify-user/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
