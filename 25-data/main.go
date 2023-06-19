package main

/**
Accédons à https://www.sqlite.org/download.html, recherchons la section des fichiers binaires précompilés pour notre système d'exploitation et
téléchargeons le package d'outils. Décompressons l'archive zip et copions le fichier sqlite3 dans le dossier data.

La bibliothèque standard Go comprend des fonctionnalités permettant de travailler avec des bases de données de manière simple et cohérente,
mais s'appuie sur des packages de pilotes de base de données pour implémenter ces fonctionnalités pour chaque moteur de base de données
ou serveur spécifique.

Pour sqlite, on a la commande ci-après à exécuter à la racine de notre projet : go get modernc.org/sqlite
Création de base de données sqlite :
> sqlite3 products.db ".read products.sql"
Pour confirmer que la base de données a été créée et remplie avec des données :
> sqlite3 products.db "select * from PRODUCTS"

La plupart des serveurs de base de données sont configurés séparément afin que le pilote de base de données ouvre une connexion à un processus distinct.
SQLite est une base de données intégrée et est inclus dans le package du pilote, ce qui signifie qu'aucune configuration supplémentaire n'est requise.

La bibliothèque standard fournit le package database/sql pour travailler avec des bases de données :
- Drivers() : cette fonction renvoie une tranche de chaînes, chacune contenant le nom d'un pilote de base de données.
- Open(driver, connectionStr) : cette fonction ouvre une base de données à l'aide du pilote et de la chaîne de connexion spécifiés.
  Les résultats sont un pointeur vers une classe de base de données, qui est utilisée pour interagir avec la base de données et une erreur
  qui indique des problèmes d'ouverture de la base de données.

Bien que ce soit une bonne idée d'appeler la méthode db.Close, nous n'avons besoin de le faire que lorsque nous avons complètement terminé avec
la base de données. Une seule base de données peut être utilisée pour effectuer des requêtes répétées sur la même base de données,
et les connexions à la base de données seront gérées automatiquement en arrière-plan. Il n'est pas nécessaire d'appeler la méthode sql.Open
pour obtenir une nouvelle base de données pour chaque requête, puis d'utiliser db.Close pour la fermer une fois la requête terminée.
**/

func main() {
	Printfln("-------Data-------")
	listDrivers()
	db, err := openDatabase()
	if err == nil {
		// db.Close() : cette fonction ferme la base de données et empêche d'effectuer d'autres opérations.
		db.Close()
	} else {
		panic(err)
	}
}
