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

- **internal/** : applications privées et bibliothèques de code. C'est le code dont nous ne souhaitons pas voir importé dans d'autres applications ou bibliothèques. Pour une architecture hexagonale, on retrouve les couches de notre hexagone.

- **pkg/** : l'on place le code qui peut être réutilisé par les applications externes. D'autres projets peuvent importer ces bibliothèques.
Il peut stocker la couche d'infrastructure, comme la base de données (par exemple pour mysql : dossier pkg/mysql et 
fichier pkg/mysql/mysql.go), comme un système de cache (par exemple pour redis : dossier pkg/redis et 
fichier pkg/redis/redis.go), comme un système de gestion de fil d'attente (par exemple pour rabbitmq : dossier pkg/rabbitmq 
et fichier pkg/rabbitmq/rabbitmq.go)

### Organisation complète

```
my-project/
├── cmd/
│   └── main.go
├── internal/
│   └── app/
|   └── delivery/
|       └── http/
|           └── v1/
|               └── handler.go
|           └── middleware.go
|           └── handler.go
|       └── grpc/
|       └── workers/
│   └── domain/
|   └── dto/
|   └── usecase/
|   └── repository/
|   └── gateway/
|       └── microservice/
|       └── grpc/
|       └── messaging/
|   └── pkg/
├── docs/
├── examples/
├── pkg/
│   └── auth/
|   └── database/gateway
|   └── cache/
|   └── queue/
|   └── storage/
├── web/
├── config/
├── deployments/
│   └── docker-compose.yaml
├── .env.example
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

- **internal/app/** : orchestration de l’application (bootstrap, gestion des modules). Peut contenir des services applicatifs globaux.

- **internal/delivery/** : les adaptateurs entrants.

--- **internal/delivery/http** : gestion des routes HTTP, contrôleurs, middlewares <br>
--- **internal/delivery/grpc** : serveurs gRPC, implémentations des services proto <br>
--- **internal/delivery/grpc** : consumers de file (Kafka, RabbitMQ, etc.)

- **config/** : templates de fichiers de configuration ou configurations par défaut ou modules de configuration pour l'application

- **delivery/** : il contient les différentes points d'entrées (Handlers). Les Handlers sont regroupés par domaine d'application.
Pour chaque groupe, sa propre structure de routeur est créée, dont les méthodes traitent les chemins.
La structure de la logique métier est injectée dans la structure du routeur, qui sera appelée par les Handlers.

- **internal/domain/** : il contient les entités métiers

- **internal/dto/** : c'est similaire avec les entités mais les différences sont : ne pas définir de **dto** dans **repository/**, les entités **dto** sont utiles pour obtenir le corps de notre requête, les paramètres de notre requête ou répondre avec nos données de notre API REST.

- **internal/usecase/** (ou **internal/service/**) : il contient toute la logique métier, comme la création d'un produit, la création d'un utilisateur, la validation des produits, la validation de la quantité de produits, l'exécution d'une transaction...

- **internal/repository/** : il contient toutes les requêtes de base de données.

- **internal/gateway** : les adaptateurs sortants.

--- **internal/gateway/microservice/** : il contient toutes les requêtes http vers les microservices externes. <br>
--- **internal/gateway/grpc/** : il contient toutes les requêtes grpc vers les microservices externes. <br>
--- **internal/gateway/messaging/** : il contient toutes les requêtes en consommation vers les brokers.

- **internal/pkg/** : il contient les packages propre à notre application (par exemple les packages **util** et **helper**)

- **pkg/database/** : initialisation DB, migrations

- **pkg/cache/** : il contient toutes les requêtes vers un serveur de cache (redis, mémoire).

- **pkg/queue/** : il contient toutes les requêtes vers un serveur de fil d'attente.

- **pkg/storage/** : initialisation s3, filesystem

- **web/** : les composants spécifiques aux applications web : assets statiques, templates serveurs et SPAs.

- **examples/** : exemples d’utilisation du projet, mini-programmes de démo.

- **docs/** : documentation technique de l'application

- **deployments/** : templates et configurations pour les IaaS, PaaS, système et l'orchestration de conteneurs (docker-compose, kubernetes/helm, mesos, terraform, bosh)

- **.env.example** : fichier exemple de configuration des variables d'environnement de l'application

<br>

**Framework et bibliothèque**

- Gin (Web Framework) : https://github.com/gin-gonic/gin
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Feature complete Kafka library in pure Go : https://github.com/twmb/franz-go
- Lodash-style Go library : https://github.com/samber/lo
- Go implementation of JSON Web Tokens : https://github.com/golang-jwt/jwt
- Go library to parse environment variables into structs : https://github.com/caarlos0/env
- Go library which loads env vars from a .env file : https://github.com/joho/godotenv
- Go a fast Go linters runner : https://github.com/golangci/golangci-lint

**Références :**

- [https://github.com/evrone/go-clean-template](https://github.com/evrone/go-clean-template)
- [https://github.com/Creatly/creatly-backend](https://github.com/Creatly/creatly-backend)
- [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [https://github.com/nurcahyaari/golang-starter](https://github.com/nurcahyaari/golang-starter)
- [https://github.com/khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture)
- [https://github.com/projectcontour/contour](https://github.com/projectcontour/contour)