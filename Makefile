.PHONY: scaffold

scaffold:
	@if [ -z "$(SERVICE_NAME)" ]; then \
		echo "Erro: SERVICE_NAME não fornecido"; \
		echo "Uso: make scaffold SERVICE_NAME=<nome-do-serviço>"; \
		exit 1; \
	fi
	@echo "Criando estrutura de pastas para: $(SERVICE_NAME)"
	@mkdir -p "operator-services/$(SERVICE_NAME)/cmd"
	@mkdir -p "operator-services/$(SERVICE_NAME)/config"
	@mkdir -p "operator-services/$(SERVICE_NAME)/docker"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/application/useCase"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/core/model"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/domain/entity"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/infrastructure/persistence/database"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/infrastructure/persistence/storage"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/infrastructure/services"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/infrastructure/web"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/module/api"
	@mkdir -p "operator-services/$(SERVICE_NAME)/internal/pkg/validation"
	@mkdir -p "operator-services/$(SERVICE_NAME)/terraform"
	@touch "operator-services/$(SERVICE_NAME)/cmd/main.go"
	@touch "operator-services/$(SERVICE_NAME)/config/config.go"
	@touch "operator-services/$(SERVICE_NAME)/app.env"
	@touch "operator-services/$(SERVICE_NAME)/docker-compose.yaml"
	@touch "operator-services/$(SERVICE_NAME)/Dockerfile"
	@touch "operator-services/$(SERVICE_NAME)/internal/application/useCase/usecase.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/infrastructure/persistence/database/adapter.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/infrastructure/persistence/database/port.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/infrastructure/persistence/storage/adapter.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/infrastructure/persistence/storage/port.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/infrastructure/web/adapter.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/infrastructure/web/port.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/module/api/api.go"
	@touch "operator-services/$(SERVICE_NAME)/internal/pkg/validation/validator.go"
	@touch "operator-services/$(SERVICE_NAME)/terraform/main.tf"
	@touch "operator-services/$(SERVICE_NAME)/terraform/ecr.tf"
	@echo "✓ Estrutura criada com sucesso em: operator-services/$(SERVICE_NAME)"
	@echo ""
	@echo "Estrutura criada:"
	@tree "operator-services/$(SERVICE_NAME)" 2>/dev/null || find "operator-services/$(SERVICE_NAME)" -print | sed -e 's;[^/]*/;|____;g;s;____|; |;g'
