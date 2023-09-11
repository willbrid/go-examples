package main

import (
	"platform/authorization"
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sportsstore/admin"
	"sportsstore/admin/auth"
	"sportsstore/models/repo"
	"sportsstore/store"
	"sportsstore/store/cart"
	"sync"
)

func registerServices() {
	services.RegisterDefaultServices()
	// repo.RegisterMemoryRepoService()
	repo.RegisterSqlRepositoryService()
	sessions.RegisterSessionService()
	cart.RegisterCartService()
	authorization.RegisterDefaultSignInService()
	authorization.RegisterDefaultUserService()
	auth.RegisterUserStoreService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},
		authorization.NewAuthComponent(
			"admin",
			authorization.NewRoleCondition("Administrator"),
			admin.AdminHandler{},
			admin.ProductsHandler{},
			admin.CategoriesHandler{},
			admin.OrdersHandler{},
			admin.DatabaseHandler{},
			admin.SignOutHandler{},
		).AddFallback("/admin/section/", "^/admin[/]?$"),
		handling.NewRouter(
			handling.HandlerEntry{"", store.ProductHandler{}},
			handling.HandlerEntry{"", store.CategoryHandler{}},
			handling.HandlerEntry{"", store.CartHandler{}},
			handling.HandlerEntry{"", store.OrderHandler{}},
			handling.HandlerEntry{"", admin.AuthenticationHandler{}},
		).AddMethodAlias("/", store.ProductHandler.GetProducts, 0, 1).
			AddMethodAlias("/products[/]?[A-z0-9]*?", store.ProductHandler.GetProducts, 0, 1),
	)
}

func main() {
	registerServices()
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
