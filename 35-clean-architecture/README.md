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
Il peut stocker la couche d'infrastructure, comme la base de données (par exemple pour mysql : dossier pkg/mysql et 
fichier pkg/mysql/mysql.go), comme un système de cache (par exemple pour redis : dossier pkg/redis et 
fichier pkg/redis/redis.go), comme un système de gestion de fil d'attente (par exemple pour rabbitmq : dossier pkg/rabbitmq 
et fichier pkg/rabbitmq/rabbitmq.go)

### Organisation de base avec les répertoires d'application web et les répertoires communs aux applications

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

- **config/** : templates de fichiers de configuration ou configurations par défaut ou modules de configuration pour l'application

- **/deployments** : templates et configurations pour les IaaS, PaaS, système et l'orchestration de conteneurs (docker-compose, kubernetes/helm, mesos, terraform, bosh)

- **/docs** : documentation de l'application

- **.env.example** : fichier exemple de configuration des variables d'environnement de l'application

- **web/** : les composants spécifiques aux applications web : assets statiques, templates serveurs et SPAs.

### Organisation complète

```
my-project/
├── cmd/
│   └── main.go
├── internal/
│   └── middleware/
│   └── auth/
|   └── delivery/
|       └── http/
|       └── grpc/
|       └── workers/
│   └── domain/
|   └── dto/
|   └── usecase/
|   └── repository/
|   └── microservice/
|   └── cache/
|   └── queue/
├── docs/
├── pkg/
│   └── util/
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

- **delivery/** : il contient les différentes points d'entrées (Handlers). Les Handlers sont regroupés par domaine d'application.
Pour chaque groupe, sa propre structure de routeur est créée, dont les méthodes traitent les chemins.
La structure de la logique métier est injectée dans la structure du routeur, qui sera appelée par les Handlers.

- **domain/** : il contient les entités métiers

- **dto/** : c'est similaire avec les entités mais les différences sont : ne pas définir de **dto** dans **repository/**, les entités **dto** 
sont utiles pour obtenir le corps de notre requête, les paramètres de notre requête ou répondre avec nos données de notre API REST.

- **usecase/** (ou **service/**) : il contient toute la logique métier, comme la création d'un produit, la création d'un utilisateur, 
la validation des produits, la validation de la quantité de produits, l'exécution d'une transaction...

- **repository/** : il contient toutes les requêtes de base de données.

- **microservice/** : il contient toutes les requêtes vers les microservices externes.

- **cache/** : il contient toutes les requêtes vers un serveur de cache.

- **queue/** : il contient toutes les requêtes vers un serveur de fil d'attente.

<br>

**Framework et bibliothèque**

- Gin (Web Framework) : https://github.com/gin-gonic/gin
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Confluent Kafka Golang : https://github.com/confluentinc/confluent-kafka-go
- Lodash-style Go library : https://github.com/samber/lo
- Go implementation of JSON Web Tokens : https://github.com/golang-jwt/jwt
- Go library to parse environment variables into structs : https://github.com/caarlos0/env
- Go library which loads env vars from a .env file : https://github.com/joho/godotenv

**Références :**

- [https://github.com/evrone/go-clean-template](https://github.com/evrone/go-clean-template)
- [https://github.com/Creatly/creatly-backend](https://github.com/Creatly/creatly-backend)
- [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [https://github.com/nurcahyaari/golang-starter](https://github.com/nurcahyaari/golang-starter)
- [https://github.com/khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture)
- [https://github.com/Creatly/creatly-backend](https://github.com/Creatly/creatly-backend)