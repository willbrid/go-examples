# Principes SOLID

### Introduction

Les principes **SOLID** sont un ensemble de cinq principes de conception logicielle orientée objet, introduits par **Robert C. Martin**, également connu sous le nom de **Uncle Bob**. Ces principes ont pour but de rendre le code plus maintenable, extensible et compréhensible. L'idée principale est de faciliter l'évolutivité d'un système tout en réduisant la complexité liée à l'ajout de nouvelles fonctionnalités. Les principes SOLID sont une base importante pour les développeurs cherchant à concevoir des architectures de code propres et modulaires.

### Pourquoi les principes SOLID sont importants ?

L'application des principes SOLID permet d'écrire un code plus modulable, réutilisable et facile à tester. Cela permet aussi de limiter les erreurs lors des modifications ou extensions de fonctionnalités dans un projet.

En Go, bien que le langage ne soit pas orienté objet au sens traditionnel, les principes SOLID peuvent être appliqués à travers l'utilisation des **interfaces**, des **structs**, et des **paquets bien définis**, ce qui améliore la maintenabilité et la flexibilité du code.

### Détails sur les principes SOLID

##### Single Responsibility Principle (SRP)

Le principe de responsabilité unique stipule qu'une classe ou une fonction doit avoir une seule raison de changer, c'est-à-dire qu'elle doit être responsable d'une seule fonctionnalité ou tâche.

**Exemple :** Supposons que nous construisons une application pour gérer des comptes utilisateurs.

- sans appliquer le principe **SRP**, le struct **User** a 2 responsabilités : gérer les données utilisateur et les enregistrer dans la base de données.

```
type User struct {
    Name string
    Email string
}

func (u *User) GetName() {
    return u.Name
}

func (u *User) GetEmail() {
    return u.Email
}

func (u *User) Save() error {
    return nil
}
```

- en appliquant le principe **SRP**, nous devons séparer ces responsabilités.

```
type User struct {
    Name string
    Email string
}

func (u *User) GetName() {
    return u.Name
}

func (u *User) GetEmail() {
    return u.Email
}

type UserRepository struct {}

func (r *UserRepository) save (u *User) error {
    return nil
}
```

A présent le struct **User** est uniquement responsable de la gestion des données des utilisateurs tandis que le struct **UserRepository** est responsable des opérations en base de données.

##### Open/Closed Principle (OCP)

Le principe **ouvert/fermé** stipule que les entités logicielles (classes, modules, fonctions) doivent être ouvertes à l'extension mais fermées à la modification. Cela signifie que nous devons être en mesure d'ajouter de nouvelles fonctionnalités sans modifier le code existant.

**Exemple :** Supposons que nous ayons une fonction simple qui calcule l'aire d'un rectangle.

```
type Rectangle struct {
    Width  float64
    Height float64
}

func Area(rectangle *Rectangle) float64 {
    return rectangle.Width * rectangle.Height
}
```

Suppons que nous souhaitons ajouter la prise en charge du calcul de l'aire d'un cercle.

- sans appliquer le principe **OCP**, on aura :

```
type Rectangle struct {
    Width  float64
    Height float64
}

type Circle struct {
    Radius float64
}

func Area(shape interface{}) float64 {
    switch s := shape.(type) {
    case *Rectangle:
        return s.Width * s.Height
    case *Circle:
        return math.Pi * math.Pow(s.Radius, 2)
    default:
        return 0
    }
}
```

- en appliquant le principe **OCP**, nous pouvons définir une interface et l'implémenter pour chaque type de surface

```
type Shape interface {
    Area() float64
}

func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (c *Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}
```

Dans cet exemple, pour ajouter un nouveau type de surface (par exemple **triangle**), nous pouvons simplement définir un nouveau struct **Triangle** et implémenter la méthode **Area**.

##### Liskov Substitution Principle (LSP)

Le principe de substitution de Liskov stipule que les objets d'une classe dérivée doivent pouvoir remplacer les objets de la classe de base sans altérer le comportement du programme.

**Exemple :** supposons un système gérant les oiseaux.

- sans appliquer le principe **LSP**, on a le code suivant :

```
type Bird interface {
    Fly() string
}

type Pigeon struct{}

func (p *Pigeon) Fly() string {
    return "Pigeon is flying."
}

type Penguin struct{}

func (p *Penguin) Fly() string {
    return "Penguin is flying."
}
```

Cependant, les pingouins ne peuvent pas voler, donc l'implémentation de Penguin viole le principe **LSP**.

- en appliquant le principe **LSP**, on aura :

```
type Bird interface {
    MakeSound() string
}

type FlyingBird interface {
    Bird
    Fly() string
}

type Pigeon struct{}

func (p *Pigeon) MakeSound() string {
    return "Pigeon is making sound."
}

func (p *Pigeon) Fly() string {
    return "Pigeon is flying."
}

type Penguin struct{}

func (p *Penguin) MakeSound() string {
    return "Penguin is making sound."
}
```

Désormais, le type **Penguin** implémente correctement l'interface **Bird** sans violer le principe **LSP**, et nous avons introduit une nouvelle interface **FlyingBird** pour les oiseaux qui peuvent voler.

##### Interface Segregation Principle (ISP)

Le principe de ségrégation des interfaces stipule qu'une classe ne doit pas être obligée d'implémenter des interfaces dont elle n'a pas besoin. Il est préférable de diviser les interfaces en plusieurs interfaces spécifiques plutôt qu'une seule interface générale.

**Exemple :** Imaginons un système de travail où certains employés doivent développer du code, d'autres doivent tester du code. 

- sans appliquer le principe **ISP**, on a le code suivant :

```
type Engineer interface {
    WriteCode()
    TestCode()
}

type SoftwareEngineer struct {}

func (s SoftwareEngineer) WriteCode() {
    fmt.Println("Writing code...")
}

func (s SoftwareEngineer) TestCode() {
    fmt.Println("Testing code...")
}

type QAEngineer struct {}

func (q QAEngineer) WriteCode() {
    fmt.Println("Writing code...")
}

func (q QAEngineer) TestCode() {
    fmt.Println("Testing code...")
}
```

Ici, **SoftwareEngineer** implémente l'interface **Engineer** et partant implémente une méthode inutile (**TestCode**). De même **QAEngineer** implémente l'interface **Engineer** et partant implémente une méthode inutile (**WriteCode**).

- en appliquant le principe **ISP**, on aura :

```
type DeveloperEngineer interface {
    WriteCode()
}

type TesterEngineer interface {
    TestCode()
}

type SoftwareEngineer struct {}

func (s SoftwareEngineer) WriteCode() {
    fmt.Println("Writing code...")
}

type QAEngineer struct {}

func (q QAEngineer) TestCode() {
    fmt.Println("Testing code...")
}
```

Ici, **SoftwareEngineer** implémente l'interface **DeveloperEngineer**, et **QAEngineer** implémente l'interface **TesterEngineer**. Aucun n'est forcé d'implémenter des méthodes inutiles.

##### Dependency Inversion Principle (DIP)

Le **principe d'inversion de dépendance** stipule que les modules de haut niveau ne doivent pas dépendre des modules de bas niveau, mais plutôt des abstractions. Cela permet de découpler les différents composants du système.

**Exemple :** suppons un projet où il y'a un package service permettant de sauvegarder les données en base de données un package mysql.

- sans appliquer le principe **DIP**, on a le code suivant :

```
type MySQLDatabase struct{}

func (m *MySQLDatabase) SaveData(data string) {
  fmt.Println("Saving data in MySQL")
}

type DataService struct {
  mySQLDatabase *MySQLDatabase
}

func (d *DataService) Save(data string) {
  d.mySQLDatabase.SaveData(data)
}
```

Dans ce cas, le **DataService** dépend directement du **MySQLDatabase**, ce qui rend difficile : le passage à une autre base de données différente de mysql ou le test du **DataService** de manière isolée.

- en appliquant le principe **DIP**, on aura :

```
type Database interface {
    SaveData(data string)
}

type MySQLDatabase struct {}

func (m *MySQLDatabase) SaveData(data string) {
    fmt.Println("Saving data in MySQL")
}

type DataService struct {
    db Database
}

func (s *DataService) Save(data string) {
    s.db.SaveData(data)
}
```

**DataService** dépend de l'interface abstraite **Database**, ce qui permet de changer facilement la base de données utilisée (par exemple, passer de MySQL à PostgreSQL) sans modifier le code du service.

<br>

**Référence :** [https://medium.com/@vishal/understanding-solid-principles-in-golang-a-guide-with-examples-f887172782a3](https://medium.com/@vishal/understanding-solid-principles-in-golang-a-guide-with-examples-f887172782a3)