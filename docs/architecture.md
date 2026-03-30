# EntraX Enterprise Web Application Blueprint

## 1) Scalable system architecture diagram

```mermaid
flowchart LR
  U[Enterprise User] --> AAD[Microsoft Entra ID\n(Azure AD)]
  U --> FE[Angular SPA\nApp Service / AKS Ingress]
  FE -->|MSAL Token| API[Go REST API\nApp Service / AKS]
  API --> DB[(Azure PostgreSQL / Azure SQL)]
  API --> KV[Azure Key Vault]
  API --> MON[Azure Monitor + App Insights]
  FE --> MON
  API --> REDIS[(Azure Cache for Redis)]
  AKS[AKS / App Service Plan] --> API
  AKS --> FE
```

## 2) Recommended folder structure

```text
entrax/
├── backend/
│   ├── cmd/api/main.go
│   ├── internal/
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── models/
│   │   └── services/
│   ├── pkg/logger/
│   ├── Dockerfile
│   └── go.mod
├── frontend/
│   ├── src/app/
│   │   ├── core/{guards,interceptors,services}
│   │   ├── features/{auth,dashboard}
│   │   └── shared/
│   ├── src/environments/
│   └── Dockerfile
├── deploy/
│   └── docker-compose.yml
└── docs/
    └── architecture.md
```

## 3) Sample REST API endpoints

- `GET /health` → liveness/readiness endpoint.
- `GET /api/v1/projects` → list projects (AAD token required).
- `POST /api/v1/projects` → create project (AAD token required).

### Sample request

```http
POST /api/v1/projects
Authorization: Bearer <access_token>
Content-Type: application/json

{
  "name": "Platform Uplift",
  "description": "Refactor and harden core services"
}
```

## 4) Authentication flow (Microsoft Entra ID)

1. User accesses Angular SPA.
2. Angular MSAL library triggers Entra ID login redirect.
3. Entra ID returns ID token + access token to SPA.
4. SPA sends `Authorization: Bearer <token>` to Go API.
5. API validates token issuer/audience/signature (JWKS in production).
6. API authorizes user/roles and executes request.

## 5) Azure service configuration

### Compute

- **Option A: Azure App Service**
  - Separate Web App for API and Frontend containers.
  - Use deployment slots for blue/green rollouts.
- **Option B: AKS**
  - Ingress Controller + cert-manager.
  - HPA for API pods, node autoscaling.

### Data

- Azure Database for PostgreSQL Flexible Server **or** Azure SQL Database.
- Private endpoint + firewall restrictions.
- Automatic backups and PITR enabled.

### Secrets

- Store DB credentials, JWT config, API secrets in Azure Key Vault.
- Use managed identity from App Service/AKS workload identity.

### Observability

- Azure Monitor + Application Insights.
- Structured logs from Go (JSON) and browser telemetry from Angular.
- Alerts on 5xx rates, p95 latency, and DB connection saturation.

## 6) Docker + deployment guide

### Local validation

```bash
cd deploy
docker compose up --build
```

### Azure CLI deployment (App Service)

```bash
# Resource group
az group create --name rg-entrax-prod --location eastus

# Container registry
az acr create --name acretraxprod --resource-group rg-entrax-prod --sku Standard
az acr login --name acretraxprod

# Build & push images
az acr build --registry acretraxprod --image entrax-api:v1 ./backend
az acr build --registry acretraxprod --image entrax-web:v1 ./frontend

# App Service Plan + Web Apps
az appservice plan create --name asp-entrax --resource-group rg-entrax-prod --is-linux --sku P1v3
az webapp create --resource-group rg-entrax-prod --plan asp-entrax --name entrax-api-prod --deployment-container-image-name acretraxprod.azurecr.io/entrax-api:v1
az webapp create --resource-group rg-entrax-prod --plan asp-entrax --name entrax-web-prod --deployment-container-image-name acretraxprod.azurecr.io/entrax-web:v1
```

### Azure CLI deployment (AKS)

```bash
az aks create --resource-group rg-entrax-prod --name aks-entrax --node-count 3 --enable-managed-identity --generate-ssh-keys
az aks get-credentials --resource-group rg-entrax-prod --name aks-entrax
kubectl apply -f k8s/
```

## 7) Security, scalability, and maintainability best practices

- Enforce HTTPS/TLS 1.2+, HSTS, and secure headers.
- Validate and sanitize all API input; centralize error responses.
- Use least-privilege RBAC and role claims from Entra ID.
- Keep secrets out of code; rotate keys via Key Vault.
- Apply clean architecture boundaries and dependency inversion.
- Add CI/CD quality gates: lint, unit/integration tests, SAST, IaC scan.
- Design for horizontal scale: stateless API, connection pooling, caching.
