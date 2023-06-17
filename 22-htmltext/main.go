package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, &Kayak)
}

func Exec1(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
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
		// Cette fonction allt3.Lookup renvoie un *Template pour le modèle chargé spécifié.
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
		// Cette fonction allt3.Lookup renvoie un *Template pour le modèle chargé spécifié.
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
		// Cette fonction allt3.Lookup renvoie un *Template pour le modèle chargé spécifié.
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
		// Cette fonction allt3.Lookup renvoie un *Template pour le modèle chargé spécifié.
		selectedTemplated5 := allt7.Lookup("template-action4.html")
		err8 := Exec1(selectedTemplated5)
		os.Stdout.WriteString("\n")
		if err8 != nil {
			Printfln("Error: %v", err8.Error())
		}
	}
}
