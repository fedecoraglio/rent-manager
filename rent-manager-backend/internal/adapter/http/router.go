package http

import (
	"log/slog"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	"rent-manager-backend/internal/adapter/config"
	"rent-manager-backend/internal/core/port"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.HTTP,
	userHandler *UserHandler,
	authHandler *AuthHandler,
	roleHandler *RoleHandler,
	ownerHandler *OwnerHandler,
	locationHandler *LocationHandler,
	tenantHandler *TenantHandler,
	propertyHandler *PropertyHandler,
	propertyCatalogHandler *PropertyCatalogHandler,
	rentalContractHandler *RentalContractHandler,
	contractCatalogHandler *ContractCatalogHandler,
	rentPaymentHandler *RentPaymentHandler,
	inflationIndexHandler *InflationIndexHandler,
	tokenProvider port.TokenProvider,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList
	ginConfig.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}

	ginConfig.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
	}
	router := gin.New()
	router.Use(
		sloggin.New(slog.Default()),
		gin.Recovery(),
		cors.New(ginConfig),
	)

	authMiddleware := AuthMiddleware(tokenProvider)

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		users := v1.Group("/users")
		users.Use(authMiddleware)
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
			users.GET("/search", userHandler.SearchUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
		roles := v1.Group("/roles")
		roles.Use(authMiddleware)
		{
			roles.GET("", roleHandler.ListRoles)
		}
		owners := v1.Group("/owners")
		owners.Use(authMiddleware)
		{
			owners.POST("", ownerHandler.CreateOwner)
			owners.GET("", ownerHandler.ListOwners)
			owners.GET("/search", ownerHandler.SearchOwners)
			owners.GET("/:id", ownerHandler.GetOwnerByID)
			owners.PUT("/:id", ownerHandler.UpdateOwner)
		}
		locations := v1.Group("/locations")
		locations.Use(authMiddleware)
		{
			locations.GET("/countries", locationHandler.ListCountries)
			locations.GET("/countries/:countryId/states", locationHandler.ListStatesByCountry)
		}
		tenants := v1.Group("/tenants")
		tenants.Use(authMiddleware)
		{
			tenants.POST("", tenantHandler.CreateTenant)
			tenants.GET("", tenantHandler.ListTenants)
			tenants.GET("/search", tenantHandler.SearchTenants)
			tenants.GET("/:id", tenantHandler.GetTenantByID)
			tenants.PUT("/:id", tenantHandler.UpdateTenant)
		}
		properties := v1.Group("/properties")
		properties.Use(authMiddleware)
		{
			properties.POST("", propertyHandler.CreateProperty)
			properties.GET("", propertyHandler.ListProperties)
			properties.GET("/search", propertyHandler.SearchProperties)
			properties.GET("summaries", propertyHandler.ListPropertiesSummary)
			properties.GET("/:id", propertyHandler.GetPropertyByID)
			properties.PUT("/:id", propertyHandler.UpdateProperty)
		}
		propertyCatalogs := v1.Group("/property-catalogs")
		propertyCatalogs.Use(authMiddleware)
		{
			propertyCatalogs.GET("/types", propertyCatalogHandler.ListPropertyTypes)
			propertyCatalogs.GET("/statuses", propertyCatalogHandler.ListPropertyStatuses)
		}
		rentalContracts := v1.Group("/rental-contracts")
		rentalContracts.Use(authMiddleware)
		{
			rentalContracts.POST("", rentalContractHandler.CreateRentalContract)
			rentalContracts.GET("", rentalContractHandler.ListRentalContracts)
			rentalContracts.GET("/:rentalContractId", rentalContractHandler.GetRentalContractByID)
			rentalContracts.PUT("/:rentalContractId", rentalContractHandler.UpdateRentalContract)
			rentalContracts.GET("/:rentalContractId/payment-schedules", rentPaymentHandler.GetRentPaymentSchedule)
			rentalContracts.GET("/:rentalContractId/payment-suggestions", rentPaymentHandler.GetRentPaymentSuggestion)
			rentalContracts.GET("/:rentalContractId/summary", rentPaymentHandler.GetRentalContractSummary)
		}
		contractCatalogs := v1.Group("/contract-catalogs")
		contractCatalogs.Use(authMiddleware)
		{
			contractCatalogs.GET("/statuses", contractCatalogHandler.ListContractStatuses)
			contractCatalogs.GET("/interest-calculation-types", contractCatalogHandler.ListInterestCalculationTypes)
			contractCatalogs.GET("/rent-adjustment-types", contractCatalogHandler.ListRentAdjustmentTypes)
		}
		rentPayments := v1.Group("/rent-payments")
		rentPayments.Use(authMiddleware)
		{
			rentPayments.POST("", rentPaymentHandler.CreateRentPayment)
			rentPayments.GET("", rentPaymentHandler.ListRentPayments)
			rentPayments.GET("/:id", rentPaymentHandler.GetRentPaymentByID)
			rentPayments.PUT("/:id", rentPaymentHandler.UpdateRentPayment)
		}
		inflationIndexes := v1.Group("/inflation-indexes")
		{
			inflationIndexes.POST("", inflationIndexHandler.Create)
			inflationIndexes.PUT("/:id", inflationIndexHandler.Update)
			inflationIndexes.GET("/:id", inflationIndexHandler.GetByID)
			inflationIndexes.GET("", inflationIndexHandler.List)
		}
	}

	return &Router{router}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
