# Architecture d'un projet Go

Structurer un projet Go de manière efficace est essentiel pour maintenir un code propre, réutilisable et évolutif.

### Organisation de base d'un projet Go

```
my-project/
├── cmd/
│   └── main.go
├── internal/
├── pkg/
│   └── util/
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

- **cmd/** : les applications principales de ce projet. Le nom de répertoire de chaque application doit correspondre au nom de l'exécutable que nous souhaitons avoir.

- **internal/** : applications privées et bibliothèques de code. C'est le code dont nous ne souhaitons pas voir importé dans d'autres applications ou bibliothèques.

- **pkg/** : l'on place le code qui peut être réutilisé par les applications externes. D'autres projets peuvent importer ces bibliothèques.

### Organisation de base avec les répertoires d'application de services ou d'application web

```
my-project/
├── cmd/
│   └── main.go
├── internal/
├── pkg/
│   └── util/
├── api/
│   └── openapi-spec/
│   └── handler/
│   └── router/
├── web/
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

- **api/** : spécifications OpenAPI/Swagger, fichiers de schémas JSON, fichiers de définitions de protocoles.

- **web/** : les composants spécifiques aux applications web : assets statiques, templates serveurs et SPAs.

### Organisation de base avec les répertoires d'application de services ou d'application web et les répertoires communs aux applications

```
my-project/
├── cmd/
│   └── main.go
├── internal/
│   └── middleware/
│   └── auth/
│   └── util/
├── docs/
├── pkg/
│   └── util/
├── api/
│   └── openapi-spec/
│   └── handler/
│   └── router/
├── config/
├── deployments/
│   └── docker-compose.yaml
├── .env.example
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

- **config/** : templates de fichiers de configuration ou configurations par défaut ou modules de configuration pour l'application

- **/deployments** : templates et configurations pour les IaaS, PaaS, système et l'orchestration de conteneurs (docker-compose, kubernetes/helm, mesos, terraform, bosh)

- **/docs** : documentation de l'application

- **.env.example** : fichier exemple de configuration des variables d'environnement de l'application

### Organisation de base avec modèle d'hexagonal (ports & adapters) 

```
my-project/
├── cmd/
│   └── main.go
├── internal/
│   └── middleware/
│   └── auth/
│   └── util/
├── docs/
├── pkg/
│   └── util/
├── api/
│   └── openapi-spec/
│   └── handler/
│   └── router/
├── domain/
├── dto/
├── usecase/
├── repository/
├── infrastructure/
├── config/
├── deployments/
│   └── docker-compose.yaml
├── .env.example
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

- **domain/** : il contient les entités métiers

- **dto/** : c'est similaire avec les entités mais les différences sont : ne pas définir de **dto** dans **repository/**, les entités **dto** sont utiles pour obtenir le corps de notre requête, les paramètres de notre requête ou répondre avec nos données de notre API REST.

- **usecase/** (ou **service/**) : il contient toute la logique métier, comme la création d'un produit, la création d'un utilisateur, la validation des produits, la validation de la quantité de produits, l'exécution d'une transaction...

- **repository/** : il contient toutes les requêtes de base de données.

- **infrastructure/** : il stocke la couche d'infrastructure, comme la base de données (par exemple pour mysql : dossier infrastructure/db et fichier infrastructure/db/mysql.go), comme un système de cache (par exemple pour redis : dossier infrastructure/cache et fichier infrastructure/cache/redis.go), comme un système de gestion de fil d'attente (par exemple pour rabbitmq : dossier infrastructure/queue et fichier infrastructure/queue/rabbitmq.go).

<br>

**Référence :**

- [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [https://github.com/nurcahyaari/golang-starter](https://github.com/nurcahyaari/golang-starter)