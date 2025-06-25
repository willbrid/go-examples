# Dépendance

Le plus grand défi pour les développeurs de tests, est probablement le code qui comporte des dépendances. Une dépendance, c’est ce dont un composant a besoin pour fonctionner. Certaines dépendances sont parfaitement raisonnables et l'on peut facilement les construire et les inclure dans un test. Cependant un trop grand nombre de dépendances peut être source de problèmes de conception car cela rend le code : **difficile à comprendre**, **difficile à tester**, **fragile à modifier**.

### N'écrivez simplement pas de fonctions non testables

- La méthode **rouge-vert-refactorisation** (**test → code → amélioration**) aide, mais ne garantit pas une bonne conception.
- Il faut éviter d’écrire des fonctions non testables, mais ce n’est pas toujours possible, surtout avec du code existant.
- Les tests fragiles (qui échouent souvent sans raison claire) révèlent souvent des composants trop dépendants les uns des autres.

Une solution pour résoudre un problème de test causé par des dépendances c'est de **repenser la conception**. <br>
Par exemple, si deux composants échouent ensemble lorsqu’un seul change, ils sont sûrement trop liés. Dans ce cas, mieux vaut les fusionner : ils forment en réalité un seul composant, même si le code les sépare.

### Réduire la portée de la dépendance

Parfois, on ne peut pas fusionner deux composants dépendants, à cause de contraintes métier, techniques, ou parce que la dépendance est externe (comme un serveur email) : dans ce cas, l'idée est de **réduire ou supprimer la dépendance**.

**Exemple** : une méthode qui envoie un e-mail selon que le compte de l'utilisateur est proche de l'expiration.

- Tester si un e-mail a bien été envoyé est compliqué (infrastructure, serveur mail, DNS, erreurs possibles, etc.).
- Mais en réalité, ce qu'on veut vraiment tester, c’est la logique de décision (par exemple : "le compte va-t-il bientôt expirer ?").

**Solution** : 

- Séparer la logique de décision de l’action "**envoyer un e-mail**".
- Ainsi, on peut tester la logique sans serveur mail, juste avec des données simulées.
- L’envoi d’e-mail pourra être testé à part, plus simplement.

### Méfiez-vous de l’injection de dépendance

Un gros problème dans le code apparaît quand une fonction utilise une dépendance de manière implicite, par exemple une base de données cachée (via une variable globale ou une configuration invisible). Cela rend le code difficile à tester.

**Exemple** : Une fonction **CreateUserInDB** qui ne prend pas en paramètre la base de données. Elle l'utilise en interne sans qu’on sache comment. Mauvais pour les tests, car on ne peut pas facilement contrôler ou simuler cette dépendance.

**Solution simple** : On peut refactoriser pour que la base de données soit passée explicitement en paramètre. C’est mieux, mais on a encore le problème de devoir mettre en place une vraie base (ce qui rend les tests fragiles).

**Idée clé** : On ne veut pas vraiment tester que l’utilisateur est écrit en base, mais plutôt que la logique qui décide d’écrire est correcte.

L'objectif de l'injection de dépendances est de faciliter l'écriture de composants nécessitant de nombreuses dépendances.

### Éviter les dommages causés par les tests

**Problème** : Une fonction qui interroge une API externe est difficile à tester
- Le réseau est lent ou instable
- L’API peut être en panne
- Les tests deviennent coûteux ou peu fiables
- Et surtout : on ne veut pas tester une API externe, mais notre propre code.

**Tentative de solution** : Injecter l’URL de l’API en paramètre pour pouvoir utiliser un faux serveur en test
- Cela fonctionne, mais alourdit inutilement l’API de notre fonction : deux paramètres au lieu d’un
- Et l’utilisateur réel doit aussi fournir ce deuxième paramètre, inutile pour lui
- Cela dégrade la conception pour des raisons uniquement liées aux tests.

**Meilleure approche** :
- Ne pas perturber l’API publique juste pour faciliter les tests
- Isoler ou encapsuler la dépendance externe (l’appel à l’API)
- Utiliser un faux (un composant simulé) pour tester uniquement la logique propre à notre application, sans interagir réellement avec l’API.

### Décomposer le comportement en sous-composants

**Problème** : Dans les cas d'exemple vus plus haut : « envoyer un e-mail » ou « créer un utilisateur », on ne peut pas toujours éliminer une dépendance externe (comme un serveur SMTP ou une base de données), mais on peut réduire sa portée pour améliorer la testabilité.

**Exemple** : Nous prenons l'exemple d'une fonction **GetForecast** qui fait appel à une API externe. Nous pourrions la décomposer en plusieurs étapes indépendantes
- FormatURL : construire l’URL de la requête avec la localisation.
- Appel HTTP : envoyer la requête à l’API.
- ParseResponse : analyser et formater la réponse.

**Solution** : Isoler les étapes testables
- **FormatURL** peut être testée comme une simple fonction de transformation de chaînes.
- **ParseResponse** peut être testée en lui injectant un exemple de réponse JSON.
- Cela permet de tester l’essentiel de la logique métier sans dépendance réseau.

Même si la dépendance (ici, l'API externe) ne peut pas être éliminée, il est souvent possible de découper la logique métier en sous-composants testables indépendamment.

### Réassembler les morceaux testés

Réécrire **GetForecast** en s’appuyant sur les fonctions **FormatURL** et **ParseResponse**, testées séparément. <br>
- Le code d’appel réseau peut rester simple et être testé manuellement si nécessaire.
- L’important est de tester en profondeur les parties critiques : **construction de l’URL**, **analyse de la réponse**.

Pour tester des comportements difficiles :
- Fragmenter le code en petits morceaux indépendants.
- Tester chaque fragment de manière autonome, sans dépendances réseau ou système.
- Assembler les fragments dans la version finale du programme, ce qui est souvent simple à valider.

Décomposer un code dépendant en fonctions testables permet d'assurer fiabilité et simplicité, tout en évitant les tests lents et fragiles.

### Extraire et isoler la logique clé

Supposons que nous souhaitons tester une logique de routage des requêtes à leur gestionnaire qui définit une api précise. La façon la plus naturelle de tester un tel comportement aurait été de lancer une instance locale du serveur et de lui envoyer des requêtes de différents types, en vérifiant que le gestionnaire approprié est invoqué. Ce qui serait difficile, puisque :
- les gestionnaires provoquent des effets secondaires lourds
- la construction des requêtes est fastidieuse et spécifique à chaque type
- les réponses sont trop similaires pour identifier quel gestionnaire a été réellement appelé

Une solution de test serait : que chaque type de requête est dirigé vers le bon gestionnaire :
- en utilisant une table de correspondance entre types de requêtes et gestionnaires
- en testant directement le contenu de cette table, plutôt que le système complet

Même si le système global est difficile à tester, on peut isoler et tester la **logique critique** (ici, le routage) de manière simple et fiable, en la dissociant des **dépendances lourdes**.

### Isoler en utilisant un adaptateur

**Problème** : Les programmes interagissent souvent avec des dépendances externes (base de données, API, exécutable), ce qui complique :
- la conception du code
- la testabilité, notamment à cause de l’instabilité ou de l’indisponibilité de ces dépendances.

**Solution** : le modèle d’adaptateur <br>
Un adaptateur (ou ambassadeur) est un composant qui centralise la gestion d’une dépendance externe :
- Il traduit les requêtes internes en un format que comprend l’API externe (ex. : FormatURL)
- Il interprète les réponses de l’API pour les convertir en structures internes (ex. : ParseResponse).

**Avantages** :

- Isolation du code dépendant de l’API : le reste du système n’a pas à connaître les détails de l’API externe
- Facilite les tests : on peut tester le reste du code sans dépendre de l’API distante
- Améliore la conception : l’interaction avec la dépendance est clairement encapsulée.

L'utilisation d'un adaptateur permet de rendre le système plus modulaire, plus robuste, et plus testable, tout en réduisant l'impact des dépendances externes.

### Faux, mocks, stubs, doubles et spies

Dans le contexte des tests logiciels, un faux (ou test double) est un objet ou une structure utilisée uniquement à des fins de test, en remplacement d’une dépendance externe ou d’une opération coûteuse.

**Types de faux évoqués**

- **mock**

Un **mock** est un faux qui a des attentes préprogrammées quant à son utilisation et qui vérifie ces attentes lors d'un test.

- **Stub** 

--- Ne fait rien ou renvoie des réponses fixes <br>
--- Utilisé uniquement pour satisfaire la syntaxe <br>
--- Considéré comme une mauvaise conception si le composant testé dépend de lui sans réelle utilité.

- **Faux utile (ex. : MapStore)**

--- Composant fonctionnel mais allégé (par exemple, une base en mémoire au lieu d’une base persistante) <br>
--- Accélère les tests tout en restant fidèle au comportement réel.

- **Espion (spy)**

--- Faux qui enregistre les interactions (ex. : un bytes.Buffer utilisé pour vérifier ce qui a été écrit) <br>
--- Permet d'inspecter les effets secondaires pour s'assurer qu’ils sont corrects.

### Ne soyez pas tenté d’écrire des mocks

Un simulacre (**mock**) est un faux (test double) programmable :
- Il échoue automatiquement le test si certaines attentes prédéfinies ne sont pas respectées (ex. : méthode non appelée ou appelée avec de mauvais paramètres).
- Contrairement aux espions, qui enregistrent passivement les appels pour inspection après coup, un mock vérifie activement pendant l’exécution que les appels sont corrects.

**Inconvénients des mocks**

- Complexité excessive : Les mocks peuvent devenir aussi complexes que les composants qu’ils remplacent
- Fragilité des tests : Les mocks vérifient des détails d’implémentation (séquences d’appels), pas de comportement observable. Cela rend les tests sensibles aux refactorings internes
- Couplage fort à la conception : Les mocks présupposent une structure rigide (souvent orientée objet) et freinent l'évolution naturelle du design
- Pollution d'interface : Créer des interfaces uniquement pour permettre le mocking rend les tests plus fragiles et le code plus verbeux.

**Quand utiliser un mock ?**

- En dernier recours, lorsqu’aucune autre technique (faux simple, espion, découplage) ne permet de tester efficacement un comportement
- Exemple toléré : sqlmock pour simuler une base SQL difficile à tester autrement, mais à utiliser avec parcimonie.

Les mocks ne sont pas intrinsèquement mauvais, mais doivent être utilisés avec modération. <br>
Leur usage excessif est souvent un symptôme de mauvaise conception ou d’un recours excessif à l’orienté objet. <br>
La priorité reste de créer des tests simples, fiables et découplés, en favorisant d’abord des structures plus légères comme les espions ou les faux en mémoire.

### Transformez le temps en données

Pour écrire les tests sur les fonctions qui manipulent le temps, il vaut mieux transformer le temps en une donnée. **Go** facilite cette approche car il fournit le type **time.Time**, mais il faut être prudent : le temps n’est pas juste un nombre. En interne, Go utilise deux horloges :

- une horloge murale (modifiable par le système ou l’utilisateur),
- une horloge monotone (qui ne peut qu’avancer).

Cela signifie que deux **time.Time** ne sont pas forcément identiques même si leurs dates semblent l’être, à cause de la partie monotone incluse. Donc, au lieu de comparer deux **time.Time** avec **==**, il vaut mieux utiliser la méthode : **.Equal()** si les deux temps sont "exactement" les mêmes, ou la méthode **.Sub()** pour mesurer une différence (**time.Duration**) et vérifier si elle est "assez petite".

Exemple

```
func TestOneHourAgo_ReturnsExpectedTime(t *testing.T) {
    t.Parallel()
    now := time.Now()
    want := now.Add(-time.Hour)
    got := past.OneHourAgo()
    delta := want.Sub(got).Abs()
    
    if delta > 10*time.Microsecond {
        t.Errorf("want %v, got %v", want, got)
    }
}
```