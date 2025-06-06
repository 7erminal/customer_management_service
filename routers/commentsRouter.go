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

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_emergency_contactsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Customer_guarantorsController"],
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

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "GetAllByBranch",
            Router: `/branch/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "GetCustomerCount",
            Router: `/count/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "UpdateCustomerLastTxnDate",
            Router: `/last-txn/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "UpdateCustomerImage",
            Router: `/update-customer-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:Identification_typesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
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

    beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetOneByName",
            Router: `/role/:role`,
            AllowHTTPMethods: []string{"get"},
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

    beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "UploadShopLogo",
            Router: `/upload-shop-logo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UserExtraDetailsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
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
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetUsersUnderBranch",
            Router: `/branch/:branch_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "UpdateUserBranch",
            Router: `/branch/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetUserCount",
            Router: `/count/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Deactivate",
            Router: `/deactivate/:id`,
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
            Method: "GetUserInvites",
            Router: `/get-user-invites`,
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
            Method: "GetInvite",
            Router: `/invite/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "RevokeUserInvites",
            Router: `/revoke-invite`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "UpdateUserRole",
            Router: `/role/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetUsersWithRole",
            Router: `/role/:role_id`,
            AllowHTTPMethods: []string{"get"},
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
            Method: "UpdateUserImage",
            Router: `/update-user-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "UpdateInviteToken",
            Router: `/update-user-invite`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "UpdateUserInvite",
            Router: `/update-user-invite/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"] = append(beego.GlobalControllerRouter["customer_management_service/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetUserInvite",
            Router: `/user-invite/:token`,
            AllowHTTPMethods: []string{"get"},
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
