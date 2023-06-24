package main

import "database/sql"

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

Les méthodes de la classe sql.Rows :
- Next() : cette méthode passe à la ligne de résultat suivante. Le résultat est un booléen, qui est vrai lorsqu'il y a des données à lire et
           faux lorsque la fin des données a été atteinte, à quel point la méthode Close est automatiquement appelée.
- NextResultSet() : cette méthode passe au jeu de résultats suivant lorsqu'il existe plusieurs jeux de résultats dans la même réponse de base de données.
                    La méthode renvoie true s'il existe un autre ensemble de lignes à traiter.
- Scan(...targets) : cette méthode affecte les valeurs SQL de la ligne actuelle aux variables spécifiées. Les valeurs sont attribuées via des pointeurs et
                     la méthode renvoie une erreur qui indique quand les valeurs ne peuvent pas être analysées.
- Close() : cette méthode empêche une énumération supplémentaire des résultats et est utilisée lorsque toutes les données ne sont pas requises.
            Il n'est pas nécessaire d'appeler cette méthode si la méthode Next est utilisée pour avancer jusqu'à ce qu'elle renvoie false.

Les méthodes de la classe sql.Row :
- Scan(...targets) : cette méthode affecte les valeurs SQL de la ligne actuelle aux variables spécifiées.
  Les valeurs sont attribuées via des pointeurs et la méthode renvoie une erreur qui indique quand les valeurs ne peuvent pas être analysées
  ou s'il n'y a pas de lignes dans le résultat. S'il y a plusieurs lignes dans la réponse, toutes sauf la première ligne seront ignorées.
- Err() : cette méthode renvoie une erreur qui indique des problèmes d'exécution de la requête.

La méthode Exec(query, ...args) est utilisée pour exécuter des instructions qui ne produisent pas de lignes. Le résultat de la méthode Exec
est une valeur Result, qui définit une erreur qui indique des problèmes d'exécution de l'instruction et les méthodes suivantes :
- RowsAffected() : cette méthode renvoie le nombre de lignes qui ont été affectées par l'instruction, exprimée sous la forme d'un int64.
                   Cette méthode renvoie également une erreur, qui est utilisée lorsqu'il y a des problèmes d'analyse de la réponse ou
				   lorsque la base de données ne prend pas en charge cette fonctionnalité.
- LastInsertId() : cette méthode renvoie un int64 qui représente la valeur générée par la base de données lors de l'exécution de l'instruction,
                   qui est généralement une clé générée automatiquement. Cette méthode renvoie également une erreur, qui est utilisée lorsque
				   la valeur renvoyée par la base de données ne peut pas être analysée dans un Go int.

Prepare(query) : cette méthode de la classe db crée une instruction préparée pour la requête spécifiée.
                 Les résultats sont une structure Stmt et une erreur indiquant des problèmes lors de la préparation de l'instruction.
Les méthodes de la classe Stmt :
- Query(...vals) : cette méthode exécute l'instruction préparée, avec les valeurs facultatives.
                   Les résultats sont une classe sql.Rows et une erreur. Cette méthode est équivalente à la méthode DB.Query.
- QueryRow(...vals) : cette méthode exécute l'instruction préparée, avec les valeurs facultatives.
                      Les résultats sont une classe sql.Row et une erreur. Cette méthode est équivalente à la méthode DB.QueryRow.
- Exec(...vals) : cette méthode exécute l'instruction préparée, avec les valeurs facultatives.
                  Les résultats sont un résultat et une erreur. Cette méthode est équivalente à la méthode DB.Exec.
- Close() : cette méthode ferme l'instruction. Les instructions ne peuvent pas être exécutées après leur fermeture.

Begin() : cette méthode de la classe db démarre une nouvelle transaction. Les résultats sont un pointeur vers une valeur Tx et une erreur indiquant
des problèmes lors de la création de la transaction.
-> Les méthodes de la classe sql.Tx :
- Query(query, ...args) : cette méthode est équivalente à la méthode DB.Query, mais la requête est exécutée dans le cadre de la transaction.
- QueryRow(query, ...args) : cette méthode est équivalente à la méthode DB.QueryRow, mais la requête est exécutée dans le cadre de la transaction.
- Exec(query, ...args) : cette méthode est équivalente à la méthode DB.Exec, mais la requête/instruction est exécutée dans le cadre de la transaction.
- Prepare(query) : cette méthode est équivalente à la méthode DB.Query, mais l'instruction préparée qu'elle crée est exécutée dans le cadre de la transaction.
- Stmt(statement) : cette méthode accepte une instruction préparée créée en dehors de la portée de la transaction et en renvoie une qui est exécutée
                    dans la portée de la transaction.
- Commit() : cette méthode valide les modifications en attente dans la base de données, renvoyant une erreur qui indique des problèmes
             d'application des modifications.
- Rollback() : cette méthode interrompt les transactions afin que les modifications en attente soient ignorées.
               Cette méthode renvoie une erreur qui indique des problèmes lors de l'abandon de la transaction.

La réflexion est une fonctionnalité qui permet d'inspecter et d'utiliser les types et les valeurs lors de l'exécution.
Il existe des méthodes définies par la classe sql.Rows qui sont utiles lors de l'utilisation de la réflexion pour traiter une réponse de base de données :
- Columns() : cette méthode renvoie une tranche de chaînes contenant les noms des colonnes de résultats et une erreur,
              qui est utilisée lorsque les résultats ont été fermés.
- ColumnTypes() : cette méthode renvoie une tranche *ColumnType, qui décrit les types de données des colonnes de résultats.
                  Cette méthode renvoie également une erreur, qui est utilisée lorsque les résultats ont été fermés.
				  La méthode ColumnTypes renvoie une tranche de pointeurs vers la classe ColumnType, qui définit les méthodes :
	- Name() : cette méthode renvoie le nom de la colonne tel qu'il est spécifié dans les résultats, exprimé sous forme de chaîne.
	- DatabaseTypeName() : cette méthode renvoie le nom du type de colonne dans la base de données, exprimé sous forme de chaîne.
	- Nullable() : cette méthode renvoie deux résultats booléens. Le premier résultat est vrai si le type de base de données peut être nul.
	               Le deuxième résultat est vrai si le pilote prend en charge les valeurs nullables.
	- DecimalSize() : cette méthode renvoie des détails sur la taille des valeurs décimales. Les résultats sont un int64 qui spécifie la précision,
	                  un int64 qui spécifie l'échelle et un booléen qui est vrai pour les types décimaux et faux pour les autres types.
	- Length() : cette méthode renvoie la longueur des types de base de données qui peuvent avoir des longueurs variables. Les résultats sont un
	             int64 qui spécifie la longueur et un booléen qui est vrai pour les types qui définissent une longueur et faux pour les autres types.
	- ScanType() : cette méthode renvoie un reflect.Type qui indique le type Go qui sera utilisé lors de l'analyse de cette colonne avec
	               la méthode Rows.Scan.
**/

type Product0 struct {
	Id       int
	Name     string
	Category string
	Price    float64
}

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id   int
	Name string
	Category
	Price float64
}

func queryDatabase(db *sql.DB) {
	/**
	Cette méthode db.Query(query, ...args) exécute la requête spécifiée, en utilisant les arguments facultatifs.
	Les résultats sont une classe sql.Rows, qui contient les résultats de la requête, et une erreur qui indique des problèmes d'exécution de la requête.
	**/
	rows, err := db.Query("SELECT * from Products")
	if err == nil {
		for rows.Next() {
			var id, category int
			var name string
			var price float64
			scanErr := rows.Scan(&id, &name, &category, &price)
			if scanErr == nil {
				Printfln("Row: %v %v %v %v", id, name, category, price)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error : %v", err.Error())
	}
}

func queryDatabase1(db *sql.DB) []Product0 {
	products := []Product0{}
	rows, err := db.Query("SELECT * from Products")
	if err == nil {
		p := Product0{}
		for rows.Next() {
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error : %v", err.Error())
	}

	return products
}

func queryDatabase2(db *sql.DB) []Product {
	products := []Product{}
	rows, err := db.Query(`
		SELECT Products.Id, Products.Name, Products.Price,
		Categories.Id as Cat_Id, Categories.Name as CatName
		FROM Products, Categories WHERE Products.Category = Categories.Id
	`)
	if err == nil {
		p := Product{}
		for rows.Next() {
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error : %v", err.Error())
	}

	return products
}

func queryDatabase3(db *sql.DB, categoryName string) []Product {
	products := []Product{}
	/**
		Les arguments facultatifs de la méthode Query sont des valeurs pour les espaces réservés dans la chaîne de requête,
		ce qui permet d'utiliser une seule chaîne pour différentes requêtes.
	    **/
	rows, err := db.Query(`
		SELECT Products.Id, Products.Name, Products.Price,
		Categories.Id as Cat_Id, Categories.Name as CatName
		FROM Products, Categories WHERE Products.Category = Categories.Id
		AND Categories.Name = ?`, categoryName)
	if err == nil {
		p := Product{}
		for rows.Next() {
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error : %v", err.Error())
	}

	return products
}

func queryDatabase4(db *sql.DB, id int) (p Product) {
	// La méthode QueryRow(query, ...args) exécute une requête censée renvoyer une seule ligne (classe sql.Row), ce qui évite d'avoir à énumérer les résultats
	row := db.QueryRow(`
		SELECT Products.Id, Products.Name, Products.Price,
		Categories.Id as Cat_Id, Categories.Name as CatName
		FROM Products, Categories WHERE Products.Category = Categories.Id
		AND Products.Id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}

	return
}

func queryDatabase5(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`SELECT Products.Id, Products.Name, Products.Price,
			Categories.Id as "Category.Id", Categories.Name as "Category.Name"
			FROM Products, Categories
			WHERE Products.Category = Categories.Id`)
	if err != nil {
		return
	} else {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err)
		}
	}
	return
}

func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`INSERT INTO Products (Name, Category, Price) VALUES(?,?,?)`, p.Name, p.Category.Id, p.Price)
	if err == nil {
		id, err = res.LastInsertId()
		if err != nil {
			Printfln("Result error: %v", err.Error())
		} else {
			Printfln("Last insert ID : %v", id)
		}
	} else {
		Printfln("Exec error: %v", err.Error())
	}

	return
}

func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) {
	result, err := InsertNewCategoryPrepare(db).Exec(name)
	if err == nil {
		newID, _ := result.LastInsertId()
		for _, id := range productIDs {
			ChangeProductCategoryPrepare(db).Exec(int(newID), id)
		}
	} else {
		Printfln("Prepared statement error: %v", err)
	}
}

func insertAndUseCategoryWithTransaction(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if err == nil {
		catResult, err := tx.Stmt(InsertNewCategoryPrepare(db)).Exec(name)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(ChangeProductCategoryPrepare(db))
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if err != nil || updatedFailed {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return
}

func main() {
	Printfln("-------Data-------")
	listDrivers()
	db, err := openDatabase()
	if err == nil {
		queryDatabase(db)
		// db.Close() : cette fonction ferme la base de données et empêche d'effectuer d'autres opérations.
		db.Close()
	} else {
		panic(err)
	}

	db1, err1 := openDatabase()
	if err1 == nil {
		products := queryDatabase1(db1)
		for _, p := range products {
			Printfln("#%v: %v", p.Id, p)
		}
		db1.Close()
	} else {
		panic(err1)
	}

	db2, err2 := openDatabase()
	if err2 == nil {
		products := queryDatabase2(db2)
		for _, p := range products {
			Printfln("#%v: %v", p.Id, p)
		}
		db2.Close()
	} else {
		panic(err2)
	}

	db3, err3 := openDatabase()
	if err3 == nil {
		for _, cat := range []string{"Soccer", "Watersports"} {
			Printfln("--- %v Results ---", cat)
			products := queryDatabase3(db3, cat)
			for i, p := range products {
				Printfln("#%v: %v %v %v", i, p.Name, p.Category.Name, p.Price)
			}
		}
		db3.Close()
	} else {
		panic(err3)
	}

	db4, err4 := openDatabase()
	if err4 == nil {
		for _, id := range []int{1, 3, 10} {
			p := queryDatabase4(db4, id)
			Printfln("Product: %v", p)
		}
		db4.Close()
	} else {
		panic(err4)
	}

	db5, err5 := openDatabase()
	if err5 == nil {
		newProduct := Product{Name: "Stadium", Category: Category{Id: 2}, Price: 79500}
		newId := insertRow(db5, &newProduct)
		p := queryDatabase4(db5, int(newId))
		Printfln("New Product: %v", p)
		db5.Close()
	} else {
		panic(err5)
	}

	db6, err6 := openDatabase()
	if err6 == nil {
		insertAndUseCategory(db6, "Misc Products", 2)
		p := queryDatabase4(db6, 2)
		Printfln("Product: %v", p)
		db6.Close()
	} else {
		panic(err6)
	}

	db7, err7 := openDatabase()
	if err7 == nil {
		insertAndUseCategoryWithTransaction(db7, "Category_1", 2)
		p := queryDatabase4(db7, 2)
		Printfln("Product: %v", p)
		insertAndUseCategoryWithTransaction(db7, "Category_2", 100)
		db7.Close()
	} else {
		panic(err7)
	}

	db8, err8 := openDatabase()
	if err8 == nil {
		products, _ := queryDatabase5(db8)
		for _, p := range products {
			Printfln("Product: %v", p)
		}
		db8.Close()
	} else {
		panic(err8)
	}
}
