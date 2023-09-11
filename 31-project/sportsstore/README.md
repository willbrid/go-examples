# Déploiement

### Création d'un certificat auto-signé

- Création d'une clé privée et d'une demande de signature de certificat

```
openssl req -newkey rsa:2048 -keyout certificate.key -out certificate.csr
```

- Création d'un certificat auto-signé

```
openssl x509 -signkey certificate.key -in certificate.csr -req -days 365 -out certificate.crt
```

- Création d'un certificat signé par une autorité de certification avec notre propre autorité de certification

```
openssl req -x509 -sha256 -days 1825 -newkey rsa:2048 -keyout rootCA.key -out rootCA.crt
```

```
vi certificate.ext
```

```
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
subjectAltName = @alt_names
[alt_names]
DNS.1 = localhost
```

```
openssl x509 -req -CA rootCA.crt -CAkey rootCA.key -in certificate.csr -out certificate.crt -days 365 -CAcreateserial -extfile certificate.ext
```

### Configuration du fichier config.json

```
{
    "logging" : {
        "level": "information"
    },
    "files": {
        "path": "files"
    },
    "templates": {
        "path": "templates/*.html",
        "reload": false
    },
    "sessions": {
        "key": "MY_SESSION_KEY",
        "cyclekey": false
    },
    "sql": {
        "connection_str": "store.db",
        "always_reset": false,
        "commands": {
            "Init": "sql/init_db.sql",
            "Seed": "sql/seed_db.sql",
            "GetProduct": "sql/get_product.sql",
            "GetProducts": "sql/get_products.sql",
            "GetCategories": "sql/get_categories.sql",
            "GetPage": "sql/get_product_page.sql",
            "GetPageCount": "sql/get_page_count.sql",
            "GetCategoryPage": "sql/get_category_product_page.sql",
            "GetCategoryPageCount": "sql/get_category_product_page_count.sql",
            "GetOrder": "sql/get_order.sql",
            "GetOrderLines": "sql/get_order_lines.sql",
            "GetOrders": "sql/get_orders.sql",
            "GetOrdersLines": "sql/get_orders_lines.sql",
            "SaveOrder": "sql/save_order.sql",
            "SaveOrderLine": "sql/save_order_line.sql",
            "SaveProduct": "sql/save_product.sql",
            "UpdateProduct": "sql/update_product.sql",
            "SaveCategory": "sql/save_category.sql",
            "UpdateCategory": "sql/update_category.sql",
            "UpdateOrder": "sql/update_order.sql"
        }
    },
    "authorization": {
        "failUrl": "/signin"
    },
    "http": {
        "enableHttp": false,
        "enableHttps": true,
        "httpsPort": 5500,
        "httpsCert": "certificate.crt",
        "httpsKey": "certificate.key"
    }
}
```

### Construction du binaire de l'application pour la plateforme linux

```
go build
```

Cette commande créera un fichier binaire **sportsstore**.

### Création de l'image docker

```
vi Dockerfile
```

```
FROM alpine:latest

COPY sportsstore /app/
COPY templates /app/templates
COPY sql/* /app/sql/
COPY files/* /app/files/
COPY config.json /app/
COPY certificate.* /app/

EXPOSE 5500

WORKDIR /app
ENTRYPOINT ["./sportsstore"]
```

```
docker build -t gosportsstore:0.0.1 .
```