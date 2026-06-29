package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"rent-manager-backend/internal/adapter/config"
	"rent-manager-backend/internal/adapter/http"
	"rent-manager-backend/internal/adapter/security"
	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/adapter/store/mysql/repository"
	inflationIndexPolicy "rent-manager-backend/internal/core/policy/inflation_index"
	ownerPolicy "rent-manager-backend/internal/core/policy/owner"
	propertyPolicy "rent-manager-backend/internal/core/policy/property"
	rentContractPolicy "rent-manager-backend/internal/core/policy/rent_contract"
	policy "rent-manager-backend/internal/core/policy/rent_payment"
	tenantPolicy "rent-manager-backend/internal/core/policy/tenant"
	userPolicy "rent-manager-backend/internal/core/policy/user"
	"rent-manager-backend/internal/core/service"
	contractCatalogUseCase "rent-manager-backend/internal/core/usecase/contract_catalog"
	inflationIndexUseCase "rent-manager-backend/internal/core/usecase/inflation_index"
	locationUseCase "rent-manager-backend/internal/core/usecase/location"
	ownerUseCase "rent-manager-backend/internal/core/usecase/owner"
	propertyUseCase "rent-manager-backend/internal/core/usecase/property"
	propertyCatalogUseCase "rent-manager-backend/internal/core/usecase/property_catalog"
	rentPaymentUseCase "rent-manager-backend/internal/core/usecase/rent_payment"
	rentalContractUseCase "rent-manager-backend/internal/core/usecase/rental_contract"
	roleUsecase "rent-manager-backend/internal/core/usecase/role"
	tenantUseCase "rent-manager-backend/internal/core/usecase/tenant"
	userUseCase "rent-manager-backend/internal/core/usecase/user"
	"time"
)

func main() {
	// Load environment variables
	configApp, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	slog.Info("Starting the application", "app", configApp.App.Name, "env", configApp.App.Env)

	// Init database
	ctx := context.Background()
	db, err := mysql.New(ctx, configApp.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", configApp.DB.Connection)

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")

	// Security
	passwordHasher := security.NewBcryptPasswordHasher(12)
	tokenProvider := security.NewJWTTokenProvider(
		configApp.Auth.JWTSecret,
		time.Hour*24,
	)

	// Rent
	rentalContractRepository := repository.NewRentalContractRepository(db)
	inflationIndexRepository := repository.NewInflationIndexRepository(db)
	inflationAdjustmentCalculator := service.NewIPCInflationAdjustmentCalculator(inflationIndexRepository)
	// Rent Payment
	rentPaymentCalculator := service.NewRentPaymentCalculator(inflationAdjustmentCalculator)
	rentPaymentRepository := repository.NewRentPaymentRepository(db)
	rentPaymentCreatePolicy := policy.NewRentPaymentCreatePolicy(rentPaymentRepository, rentalContractRepository)
	rentPaymentUpdatePolicy := policy.NewRentPaymentUpdatePolicy(rentPaymentRepository, rentalContractRepository)
	getRentalContractSummaryUseCase := rentPaymentUseCase.NewGetRentalContractSummaryUseCase(
		rentalContractRepository,
		rentPaymentRepository,
		rentPaymentCalculator,
	)
	getRentPaymentScheduleUseCase := rentPaymentUseCase.NewGetRentPaymentScheduleUseCase(
		rentalContractRepository,
		rentPaymentRepository,
		rentPaymentCalculator,
	)
	createRentPaymentUseCase := rentPaymentUseCase.NewCreateRentPaymentUseCase(
		rentPaymentRepository,
		rentPaymentCreatePolicy,
	)
	updateRentPaymentUseCase := rentPaymentUseCase.NewUpdateRentPaymentUseCase(
		rentPaymentRepository,
		rentPaymentUpdatePolicy,
	)
	getRentPaymentByIDUseCase := rentPaymentUseCase.NewGetRentPaymentByIDUseCase(rentPaymentRepository)
	listRentPaymentsUseCase := rentPaymentUseCase.NewListRentPaymentsUseCase(rentPaymentRepository)
	getRentPaymentSuggestionUseCase := rentPaymentUseCase.NewGetRentPaymentSuggestionUseCase(
		rentalContractRepository,
		rentPaymentRepository,
		rentPaymentCalculator,
	)

	// Role
	roleRepository := repository.NewRoleRepository(db)
	listRoleUseCase := roleUsecase.NewListRolesUseCase(roleRepository)
	roleHandler := http.NewRoleHandler(listRoleUseCase)

	// User
	userRepository := repository.NewUserRepository(db)
	userCreatePolicy := userPolicy.NewUserCreatePolicy(userRepository)
	userUpdatePolicy := userPolicy.NewUserUpdatePolicy(userRepository)
	createUserUseCase := userUseCase.NewCreateUserUseCase(userRepository, userCreatePolicy, passwordHasher)
	getUserByIDUseCase := userUseCase.NewGetUserByIDUseCase(userRepository)
	getUserByEmailUseCase := userUseCase.NewGetUserByEmailUseCase(userRepository)
	listUsersUseCase := userUseCase.NewListUsersUseCase(userRepository)
	searchUsersUseCase := userUseCase.NewSearchUsersUseCase(userRepository)
	updateUserUseCase := userUseCase.NewUpdateUserUseCase(userRepository, userUpdatePolicy)
	deleteUserUseCase := userUseCase.NewDeleteUserUseCase(userRepository)
	loginUseCase := userUseCase.NewLoginUseCase(
		userRepository,
		passwordHasher,
		tokenProvider,
	)
	userHandler := http.NewUserHandler(
		createUserUseCase,
		getUserByIDUseCase,
		getUserByEmailUseCase,
		listUsersUseCase,
		searchUsersUseCase,
		updateUserUseCase,
		deleteUserUseCase,
	)
	authHandler := http.NewAuthHandler(loginUseCase)

	// Owner
	ownerRepository := repository.NewOwnerRepository(db)
	ownerCreatePolicy := ownerPolicy.NewOwnerCreatePolicy(ownerRepository)
	ownerUpdatePolicy := ownerPolicy.NewOwnerUpdatePolicy(ownerRepository)
	createOwnerUseCase := ownerUseCase.NewCreateOwnerUseCase(ownerRepository, ownerCreatePolicy)
	getOwnerByIDUseCase := ownerUseCase.NewGetOwnerByIDUseCase(ownerRepository)
	getOwnerByDocumentNumberUseCase := ownerUseCase.NewGetOwnerByDocumentNumberUseCase(ownerRepository)
	listOwnersUseCase := ownerUseCase.NewListOwnersUseCase(ownerRepository)
	searchOwnersUseCase := ownerUseCase.NewSearchOwnersUseCase(ownerRepository)
	updateOwnerUseCase := ownerUseCase.NewUpdateOwnerUseCase(ownerRepository, ownerUpdatePolicy)
	ownerHandler := http.NewOwnerHandler(
		createOwnerUseCase,
		getOwnerByIDUseCase,
		getOwnerByDocumentNumberUseCase,
		listOwnersUseCase,
		searchOwnersUseCase,
		updateOwnerUseCase,
	)

	// Location
	locationRepository := repository.NewLocationRepository(db)
	listCountriesUseCase := locationUseCase.NewListCountriesUseCase(locationRepository)
	listStatesByCountryUseCase := locationUseCase.NewListStatesByCountryUseCase(locationRepository)
	locationHandler := http.NewLocationHandler(listCountriesUseCase, listStatesByCountryUseCase)

	// Tenant
	tenantRepository := repository.NewTenantRepository(db)
	tenantCreatePolicy := tenantPolicy.NewTenantCreatePolicy(tenantRepository)
	tenantUpdatePolicy := tenantPolicy.NewTenantUpdatePolicy(tenantRepository)
	createTenantUseCase := tenantUseCase.NewCreateTenantUseCase(tenantRepository, tenantCreatePolicy)
	getTenantByIDUseCase := tenantUseCase.NewGetTenantByIDUseCase(tenantRepository)
	getTenantByDocumentNumberUseCase := tenantUseCase.NewGetTenantByDocumentNumberUseCase(tenantRepository)
	listTenantsUseCase := tenantUseCase.NewListTenantsUseCase(tenantRepository)
	searchTenantsUseCase := tenantUseCase.NewSearchTenantsUseCase(tenantRepository)
	updateTenantUseCase := tenantUseCase.NewUpdateTenantUseCase(tenantRepository, tenantUpdatePolicy)
	tenantHandler := http.NewTenantHandler(
		createTenantUseCase,
		getTenantByIDUseCase,
		getTenantByDocumentNumberUseCase,
		listTenantsUseCase,
		searchTenantsUseCase,
		updateTenantUseCase,
	)
	// Property
	propertyRepository := repository.NewPropertyRepository(db)
	propertyCreatePolicy := propertyPolicy.NewPropertyCreatePolicy(propertyRepository)
	propertyUpdatePolicy := propertyPolicy.NewPropertyUpdatePolicy(propertyRepository)
	createPropertyUseCase := propertyUseCase.NewCreatePropertyUseCase(propertyRepository, propertyCreatePolicy)
	getPropertyByIDUseCase := propertyUseCase.NewGetPropertyByIDUseCase(propertyRepository)
	getPropertyByCodeUseCase := propertyUseCase.NewGetPropertyByCodeUseCase(propertyRepository)
	listPropertiesUseCase := propertyUseCase.NewListPropertiesUseCase(propertyRepository)
	listPropertiesSummariesUseCase := propertyUseCase.NewListPropertiesSummariesUseCase(
		rentalContractRepository,
		getRentalContractSummaryUseCase,
	)
	searchPropertiesUseCase := propertyUseCase.NewSearchPropertiesUseCase(propertyRepository)
	updatePropertyUseCase := propertyUseCase.NewUpdatePropertyUseCase(propertyRepository, propertyUpdatePolicy)
	propertyHandler := http.NewPropertyHandler(
		createPropertyUseCase,
		getPropertyByIDUseCase,
		getPropertyByCodeUseCase,
		listPropertiesUseCase,
		listPropertiesSummariesUseCase,
		searchPropertiesUseCase,
		updatePropertyUseCase,
	)
	// Property Catalog
	propertyCatalogRepository := repository.NewPropertyCatalogRepository(db)
	listPropertyTypesUseCase := propertyCatalogUseCase.NewListPropertyTypesUseCase(propertyCatalogRepository)
	listPropertyStatusesUseCase := propertyCatalogUseCase.NewListPropertyStatusesUseCase(propertyCatalogRepository)
	propertyCatalogHandler := http.NewPropertyCatalogHandler(
		listPropertyTypesUseCase,
		listPropertyStatusesUseCase,
	)
	// Contract Catalog
	contractCatalogRepository := repository.NewContractCatalogRepository(db)
	listContractStatusesUseCase := contractCatalogUseCase.NewListContractStatusesUseCase(contractCatalogRepository)
	listInterestCalculationTypesUseCase := contractCatalogUseCase.NewListInterestCalculationTypesUseCase(
		contractCatalogRepository,
	)
	listRentAdjustmentTypesUseCase := contractCatalogUseCase.NewListRentAdjustmentTypesUseCase(
		contractCatalogRepository,
	)
	contractCatalogHandler := http.NewContractCatalogHandler(
		listContractStatusesUseCase,
		listInterestCalculationTypesUseCase,
		listRentAdjustmentTypesUseCase,
	)
	// Rental Contract
	rentalContractCreatePolicy := rentContractPolicy.NewRentalContractCreatePolicy(rentalContractRepository)
	rentalContractUpdatePolicy := rentContractPolicy.NewRentalContractUpdatePolicy(rentalContractRepository)
	createRentalContractUseCase := rentalContractUseCase.NewCreateRentalContractUseCase(
		rentalContractRepository,
		rentalContractCreatePolicy,
	)
	updateRentalContractUseCase := rentalContractUseCase.NewUpdateRentalContractUseCase(
		rentalContractRepository,
		rentalContractUpdatePolicy,
	)
	getRentalContractByIDUseCase := rentalContractUseCase.NewGetRentalContractByIDUseCase(rentalContractRepository)
	listRentalContractsUseCase := rentalContractUseCase.NewListRentalContractsUseCase(
		rentalContractRepository,
	)
	rentalContractHandler := http.NewRentalContractHandler(
		createRentalContractUseCase,
		updateRentalContractUseCase,
		getRentalContractByIDUseCase,
		listRentalContractsUseCase,
	)
	rentPaymentHandler := http.NewRentPaymentHandler(
		getRentPaymentScheduleUseCase,
		createRentPaymentUseCase,
		updateRentPaymentUseCase,
		getRentPaymentByIDUseCase,
		listRentPaymentsUseCase,
		getRentPaymentSuggestionUseCase,
		getRentalContractSummaryUseCase,
	)
	inflationIndexCreatePolicy := inflationIndexPolicy.NewInflationIndexCreatePolicy(inflationIndexRepository)
	inflationIndexUpdatePolicy := inflationIndexPolicy.NewInflationIndexUpdatePolicy(inflationIndexRepository)
	createInflationIndexUseCase := inflationIndexUseCase.NewCreateInflationIndexUseCase(inflationIndexRepository,
		inflationIndexCreatePolicy,
	)
	updateInflationIndexUseCase := inflationIndexUseCase.NewUpdateInflationIndexUseCase(
		inflationIndexRepository,
		inflationIndexUpdatePolicy,
	)
	getInflationIndexByIDUseCase := inflationIndexUseCase.NewGetInflationIndexByIDUseCase(inflationIndexRepository)
	listInflationIndexesUseCase := inflationIndexUseCase.NewListInflationIndexesUseCase(inflationIndexRepository)
	inflationIndexHandler := http.NewInflationIndexHandler(
		createInflationIndexUseCase,
		updateInflationIndexUseCase,
		getInflationIndexByIDUseCase,
		listInflationIndexesUseCase,
	)

	// Init router
	router, err := http.NewRouter(
		configApp.HTTP,
		userHandler,
		authHandler,
		roleHandler,
		ownerHandler,
		locationHandler,
		tenantHandler,
		propertyHandler,
		propertyCatalogHandler,
		rentalContractHandler,
		contractCatalogHandler,
		rentPaymentHandler,
		inflationIndexHandler,
		tokenProvider,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", configApp.HTTP.URL, configApp.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
