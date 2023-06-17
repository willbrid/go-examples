package main

import (
	"html/template"
	"os"
	"strings"
	templateText "text/template"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, &Kayak)
}

func Exec1(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func Exec2(t *template.Template) error {
	productMap := map[string]Product{}
	for _, p := range Products {
		productMap[p.Name] = p
	}
	return t.Execute(os.Stdout, &productMap)
}

func Exec3(t *templateText.Template) error {
	productMap := map[string]Product{}
	for _, p := range Products {
		productMap[p.Name] = p
	}
	return t.Execute(os.Stdout, &productMap)
}

// La fonction GetCategories reçoit une tranche Product et renvoie l'ensemble des valeurs Category uniques.
func GetCategories(products []Product) (categories []string) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}

	return
}

/*
*

	Le package html/template définit un ensemble d'alias de type chaîne qui sont utilisés pour indiquer que le résultat d'une fonction nécessite
	un traitement spécial :
	- template.CSS : ce type désigne le contenu CSS.
	- template.HTML : ce type désigne un fragment de HTML.
	- template.HTMLAttr : ce type indique une valeur qui sera utilisée comme valeur pour un attribut HTML.
	- template.JS : ce type désigne un fragment de code JavaScript.
	- template.JSStr : ce type désigne une valeur destinée à apparaître entre guillemets dans une expression JavaScript.
	- template.Srcset : ce type indique une valeur qui peut être utilisée dans l'attribut srcset d'un élément img.
	- template.URL : ce type désigne une URL.

*
*/
func GetCategoriesHtml(products []Product) (categories []template.HTML) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			categories = append(categories, "<b>p.Category</b>")
		}
	}

	return
}

func main() {
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}

	/**
	Cette fonction template.ParseFiles(...files) charge un ou plusieurs fichiers, qui sont spécifiés par leur nom. Le résultat est un modèle
	qui peut être utilisé pour générer du contenu et une erreur qui signale des problèmes de chargement des modèles.
	**/
	t1, err1 := template.ParseFiles("templates/template.html")
	if err1 == nil {
		/**
		Cette fonction t1.Execute exécute le modèle en utilisant les données spécifiées et écrit la sortie dans le writer spécifié.
		**/
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error : %v", err1.Error())
	}

	// Chargement de deux modèles par appel à la méthode template.ParseFiles pour chaque modèle
	t2, err2 := template.ParseFiles("templates/template.html")
	t3, err3 := template.ParseFiles("templates/extras.html")
	if err2 == nil && err3 == nil {
		t2.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		t3.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error : %v %v", err2.Error(), err3.Error())
	}

	/**
	Lorsque plusieurs fichiers sont chargés avec ParseFiles, le résultat est une valeur Template sur laquelle la méthode ExecuteTemplate
	peut être appelée pour exécuter un modèle spécifié. Le nom de fichier est utilisé comme nom de modèle, ce qui signifie que les modèles
	de cet exemple sont nommés template.html et extras.html.
	**/
	allt1, allErr1 := template.ParseFiles("templates/template.html", "templates/extras.html")
	if allErr1 == nil {
		/**
		Cette fonction allt1.ExecuteTemplate exécute le modèle avec le nom et les données spécifiés et écrit la sortie dans le Writer spécifié.
		**/
		allt1.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allt1.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error : %v", allErr1.Error())
	}

	/**
	Cette fonction template.ParseGlob charge un ou plusieurs fichiers, qui sont sélectionnés avec un motif. Le résultat est un modèle qui peut être utilisé
	pour générer du contenu et une erreur qui signale des problèmes de chargement des modèles.
	Le modèle transmis à la fonction ParseGlob sélectionne tous les fichiers avec l'extension de fichier html dans le dossier des modèles.
	**/
	allt2, allErr2 := template.ParseGlob("templates/*.html")
	if allErr2 == nil {
		// Cette fonction allt2.Templates renvoie une tranche contenant des pointeurs vers les valeurs de modèle qui ont été chargées.
		for _, t := range allt2.Templates() {
			// Cette méthode t.Name renvoie le nom du modèle.
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error : %v", allErr2.Error())
	}

	allt3, allErr3 := template.ParseGlob("templates/*.html")
	if allErr3 == nil {
		// Cette fonction allt3.Lookup renvoie un *Template pour le modèle chargé spécifié.
		selectedTemplated1 := allt3.Lookup("template.html")
		err4 := Exec(selectedTemplated1)
		os.Stdout.WriteString("\n")
		if err4 != nil {
			Printfln("Error: %v", err4.Error())
		}
	}

	allt4, allErr4 := template.ParseGlob("templates/*.html")
	if allErr4 == nil {
		selectedTemplated2 := allt4.Lookup("template-action1.html")
		err5 := Exec(selectedTemplated2)
		os.Stdout.WriteString("\n")
		if err5 != nil {
			Printfln("Error: %v", err5.Error())
		}
	}

	// Action trimming space
	allt5, allErr5 := template.ParseGlob("templates/*.html")
	if allErr5 == nil {
		selectedTemplated3 := allt5.Lookup("template-action2.html")
		err6 := Exec(selectedTemplated3)
		os.Stdout.WriteString("\n")
		if err6 != nil {
			Printfln("Error: %v", err6.Error())
		}
	}

	// Action range
	allt6, allErr6 := template.ParseGlob("templates/*.html")
	if allErr6 == nil {
		selectedTemplated4 := allt6.Lookup("template-action3.html")
		err7 := Exec1(selectedTemplated4)
		os.Stdout.WriteString("\n")
		if err7 != nil {
			Printfln("Error: %v", err7.Error())
		}
	}

	// Utilisation des fonctions intégrées
	allt7, allErr7 := template.ParseGlob("templates/*.html")
	if allErr7 == nil {
		selectedTemplated5 := allt7.Lookup("template-action4.html")
		err8 := Exec1(selectedTemplated5)
		os.Stdout.WriteString("\n")
		if err8 != nil {
			Printfln("Error: %v", err8.Error())
		}
	}

	// Exécution conditionnelle du contenu du modèle
	allt8, allErr8 := template.ParseGlob("templates/*.html")
	if allErr8 == nil {
		selectedTemplated6 := allt8.Lookup("template-action5.html")
		err9 := Exec1(selectedTemplated6)
		os.Stdout.WriteString("\n")
		if err9 != nil {
			Printfln("Error: %v", err9.Error())
		}
	}

	// Création de modèles imbriqués nommés
	allt9, allErr9 := template.ParseGlob("templates/*.html")
	if allErr9 == nil {
		selectedTemplated7 := allt9.Lookup("template-action6.html")
		err10 := Exec1(selectedTemplated7)
		os.Stdout.WriteString("\n")
		if err10 != nil {
			Printfln("Error: %v", err10.Error())
		}
	}

	// N'importe lequel des modèles imbriqués nommés peut être exécuté directement
	allt10, allErr10 := template.ParseGlob("templates/*.html")
	if allErr10 == nil {
		selectedTemplated8 := allt10.Lookup("mainTemplate") // Appel d'un modèle imbriqué
		err11 := Exec1(selectedTemplated8)
		os.Stdout.WriteString("\n")
		if err11 != nil {
			Printfln("Error: %v", err11.Error())
		}
	}

	// Définition des blocs de modèle
	allt11, allErr11 := template.ParseFiles("templates/template-action8.html", "templates/list.html")
	if allErr11 == nil {
		selectedTemplated9 := allt11.Lookup("mainTemplate") // Appel d'un modèle imbriqué
		err12 := Exec1(selectedTemplated9)
		os.Stdout.WriteString("\n")
		if err12 != nil {
			Printfln("Error: %v", err12.Error())
		}
	}

	// Définition des fonctions de modèle
	alltNew := template.New("allTemplates")
	/**
	Pour configurer la fonction GetCategories afin qu'elle puisse être utilisée par un modèle, la méthode Funcs est appelée,
	en passant un map de noms aux fonctions.

	La map spécifie que la fonction GetCategories sera appelée en utilisant le nom getCats. La méthode Funcs doit être appelée avant
	l'analyse des fichiers de modèle, ce qui signifie la création d'un modèle à l'aide de la fonction New, qui permet ensuite d'enregistrer
	les fonctions personnalisées avant l'appel de la méthode ParseFiles ou ParseGlob
	**/
	alltNew.Funcs(map[string]interface{}{
		"getCats": GetCategories,
	})
	allt12, allErr12 := alltNew.ParseFiles("templates/template-action9.html")
	if allErr12 == nil {
		selectedTemplated10 := allt12.Lookup("mainTemplate") // Appel d'un modèle imbriqué
		err13 := Exec1(selectedTemplated10)
		os.Stdout.WriteString("\n")
		if err13 != nil {
			Printfln("Error: %v", err13.Error())
		}
	}

	// Désactivation de l'encodage des résultats de la fonction
	// Les fonctions de modèle peuvent également être utilisées pour donner accès aux fonctionnalités fournies par la bibliothèque standard.
	// Définition des variables de modèle
	alltNew1 := template.New("allTemplates1")
	alltNew1.Funcs(map[string]interface{}{
		"getCatsHtml": GetCategoriesHtml,
		"getCats":     GetCategories,
		// Le nouveau mappage donne accès à la fonction ToLower, qui transforme les chaînes en minuscules
		"lower": strings.ToLower,
	})
	allt13, allErr13 := alltNew1.ParseFiles("templates/template-action10.html")
	if allErr13 == nil {
		selectedTemplated11 := allt13.Lookup("mainTemplate") // Appel d'un modèle imbriqué
		err14 := Exec1(selectedTemplated11)
		os.Stdout.WriteString("\n")
		if err14 != nil {
			Printfln("Error: %v", err14.Error())
		}
	}

	// Utilisation de variables de modèle dans les actions range
	// Les variables peuvent également être utilisées avec l'action range, qui permet d'utiliser des cartes dans des modèles.
	allt14, allErr14 := template.ParseFiles("templates/template-action11.html")
	if allErr14 == nil {
		selectedTemplated12 := allt14.Lookup("mainTemplate") // Appel d'un modèle imbriqué
		err15 := Exec2(selectedTemplated12)
		os.Stdout.WriteString("\n")
		if err15 != nil {
			Printfln("Error: %v", err15.Error())
		}
	}

	/**
	Création de modèles de texte
	Mis à part la modification de l'instruction d'importation du package templateText "text/template" et la sélection de fichiers avec l'extension txt,
	le processus de chargement et d'exécution du modèle de texte est le même que celui du modèle de template.
	**/
	alltNew2 := templateText.New("allTemplates2")
	alltNew2.Funcs(map[string]interface{}{
		"getCats": GetCategories,
		"lower":   strings.ToLower,
	})
	allt15, allErr15 := alltNew2.ParseGlob("templates/*.txt")
	if allErr15 == nil {
		selectedTemplated11 := allt15.Lookup("mainTemplate") // Appel d'un modèle imbriqué
		err16 := Exec3(selectedTemplated11)
		os.Stdout.WriteString("\n")
		if err16 != nil {
			Printfln("Error: %v", err16.Error())
		}
	}
}
