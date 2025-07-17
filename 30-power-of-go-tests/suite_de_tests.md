# Suite de tests

### Aucun test

Même si les tests sont considérés comme une bonne pratique, beaucoup de projets en production n’en ont pas du tout, surtout dans les startups où la vitesse de développement est prioritaire sur la qualité ou la maintenabilité.

**Conséquence** : au fil du temps, le code devient si difficile à modifier sans rien casser que personne n’ose le toucher.

**Solution** : Commencez à écrire des tests petit à petit, sans attendre une autorisation formelle.
- Si vous modifiez du code non testé, ajoutez un test autour.
- Pas besoin d’un changement global d’équipe ou d’approbation : c’est une bonne habitude, adoptez-la.

Une fois que vous aurez travaillé sur un système avec des tests automatisés complets, je ne pense pas que vous voudrez vous en passer. Vous aurez cette incroyable liberté de modifier le code, de le refactoriser et de publier régulièrement des versions pour les utilisateurs finaux.

### Code legacy

Ajouter des tests à une base de code non testée peut sembler décourageant. 

**Solution** : 

- tester uniquement ce que vous modifiez.
Par exemple, si vous devez modifier une ligne dans une grosse fonction, n'essayez pas de tester toute la fonction. À la place, extraire un petit bloc autour de la ligne concernée, le placer dans une fonction séparée, puis écrire un test pour cette nouvelle petite fonction.

- Répétez cela chaque fois que vous touchez du code. Cette méthode concentre naturellement les tests sur les zones les plus actives et sensibles du projet.

C’est une stratégie efficace pour améliorer progressivement la testabilité du code sans devoir tout tester d’un coup. Mais attention : si le code est trop intriqué pour être testé tel quel, il faudra peut-être le refactoriser, ce qui est risqué sans tests. C’est le cercle vicieux classique du code legacy.

### Tests insuffisants

Lorsqu’un système contient des tests insuffisants ou de mauvaise qualité, il est utile d'appliquer la stratégie « tester ce que l'on touche » pour améliorer progressivement la couverture.
Mais il faut aussi s’attaquer aux causes profondes :

- Certains développeurs évitent les tests par manque de temps ou d’expérience. Or, avec la pratique, les tests deviennent plus rapides à écrire.

- La confiance se développe avec l’habitude. Une bonne approche consiste à décider en équipe d’écrire des tests pour chaque tâche pendant une période d’essai, avec entraide en cas de blocage.

Un autre problème fréquent : écrire les tests en dernier, ce qui conduit souvent à des tests peu efficaces.

- Écrire les tests en premier (TDD) présente de nombreux avantages, notamment en facilitant le développement et en améliorant la conception.

- Cela dit, l’ordre n’est pas une question morale : si écrire les tests en dernier aboutit à un code bien conçu et bien testé, c’est tout aussi valable.

- Mais si les tests sont faibles ou absents, alors changer de méthode (comme adopter le TDD) peut être bénéfique.

### Révision de code inefficace

La revue de code est souvent délicate, car les auteurs peuvent être sur la défensive. Pour éviter une critique inefficace et superficielle (souvent focalisée sur le style), il est plus utile de centrer la revue sur les tests.

Liste de contrôle en 10 points pour une revue centrée sur les tests :

- **Comprendre le but** : Connaît-on la raison métier ou technique du changement ?
- **Présence de tests pertinents** : Y a-t-il un test clair pour chaque objectif ou exigence du changement ?
- **Reproductibilité** : Les tests passent-ils sur une nouvelle machine sans dépendance cachée ?
- **Détection de bugs** : Les tests détectent-ils des bugs simulés (tests de mutation) ?
- **Clarté des échecs** : Les messages d'erreur sont-ils informatifs sans avoir à lire le code de test ?
- **Robustesse aux entrées** : Le code est-il testé avec des entrées variées (fuzzing, cas limites, entrées vides) ?
- **Couverture de code** : Le nouveau code est-il complètement testé ? Le reste est-il justifié ?
- **Code minimaliste** : Le code ne fait-il que ce qui est nécessaire pour passer les tests (pas de surconception) ?
- **Tests ciblés** : Les tests évitent-ils de couvrir plus que le code concerné ?
- **Lisibilité des tests** : Les tests expriment-ils clairement le comportement attendu, même pour un non-technique ?

### Tests optimistes

Même des tests avec une bonne couverture peuvent être trop optimistes, c’est-à-dire qu’ils cherchent uniquement à confirmer le bon fonctionnement du code, sans vérifier que les préconditions sont bien remplies ou que les erreurs sont détectées.

**Exemple typique** :

- On appelle Create(Alice) puis on vérifie si Alice existe.
- Si Create ne fait rien, le test passe quand même si Alice existait déjà (résidu d’un test précédent).
- Le test manque donc une vérification préalable de l’état initial (s’assurer qu’Alice n’existait pas avant).
- Si la base de données n’est pas réinitialisée entre les tests, cela masque des bugs.

Un bon test doit vérifier l’avant et l’après, pas juste le résultat attendu. Sinon, on peut croire à tort que tout fonctionne.

### Des tests pointilleux

Bien que la plupart des problèmes viennent d’un manque de tests, il arrive que certains développeurs aillent trop loin en testant plus que nécessaire.

**Risques du surtesting** :

- **Tests inutiles sur des comportements évidents ou difficilement ignorables**.
- **Vérifications excessives** : comparaison de toute une structure ou d’un fichier, alors qu’un petit sous-ensemble est pertinent.
- **Fragilité des tests** : tester des champs ou des messages d’erreur non essentiels rend les tests sensibles aux changements sans importance.
- **Tests paresseux** : comparer la sortie à un fichier standard complet, alors qu’une propriété clé ou une invariance suffirait.

**Bonne pratique** : Se concentrer sur le comportement utile à vérifier, éviter les détails non pertinents et préférer les propriétés générales (comme l’existence d’une erreur, ou des invariants) plutôt que des valeurs exactes.

### Comparaisons trop précises

Lorsque nous testons des fonctions retournant des valeurs à virgule flottante (float64), une comparaison stricte peut échouer à tort à cause des imprécisions inhérentes aux calculs flottants.

- **Pourquoi ?**

Deux valeurs à virgule flottante quasi identiques peuvent ne pas être égales en **Go**, car leur représentation en mémoire est intrinsèquement imprécise. En plus les nombres à virgule flottante impliquent une sorte de compression avec perte.

- **Solution**

Ne comparer pas deux flottants avec **==** mais comparer leur écart absolu avec un seuil de tolérance.

```
if !cmp.Equal(want, got, cmpopts.EquateApprox(0, 0.00001)) {
  t.Errorf("not close enough: want %.5f, got %.5f", want, got)
}
```

### Trop de tests

Chaque ligne de code — y compris les tests — représente une dette technique potentielle. Si elle n'est plus nécessaire, elle devient un fardeau inutile qu'il faut envisager de supprimer.

Les suites de tests doivent être entretenues, nettoyées et simplifiées régulièrement, tout comme on taille un jardin pour le maintenir sain.

Lorsque des fonctionnalités deviennent obsolètes, leurs tests associés et même leurs implémentations doivent être supprimés.

De nouveaux tests peuvent remplacer d'anciens tests, surtout dans une approche de développement itératif. Il faut alors penser à supprimer les tests redondants.

Des tests spécifiques peuvent être remplacés par des tests plus globaux une fois un composant stabilisé. Les tests agissent alors comme un échafaudage temporaire, utile pendant la construction.

Trop généraliser les tests (ex. : un seul test pour tout le système) nuit à la localisation des bugs. Il faut un bon équilibre entre couverture large et précision.

Les tests doivent refléter la structure modulaire du système. Si un composant ne peut être testé isolément, il n’est probablement pas bien conçu.

Une duplication partielle des tests est acceptable, surtout pour les comportements critiques. Mais il faut éviter les redondances inutiles par un nettoyage régulier.

### Frameworks de test

Le style de test simple utilisant uniquement le package standard de Go (testing) et éventuellement go-cmp est suffisant, clair et facile à comprendre. Il est préférable de l'utiliser plutôt que d'introduire des frameworks tiers comme **testify** ou **ginkgo/gomega**.

- **Pourquoi éviter les frameworks tiers ?**

--- Ils n'apportent rien d'essentiel que Go ne fournit pas déjà. <br>
--- Ils complexifient les tests en ajoutant une API à apprendre. <br>
--- Ils introduisent une dépendance inutile au projet. <br>
--- Ils ralentissent l'intégration des nouveaux développeurs. <br>
--- Ils éloignent les tests du style Go idiomatique.

- Sur les assertions (ex : assert.Equal(want, got))

--- Ce style masque les erreurs : on affirme un comportement sans réfléchir à l’origine des échecs. <br>
--- Il n'encourage pas à enrichir les messages d’erreur, ce qui nuit au débogage. <br>
--- Il ressemble moins au vrai code Go, contrairement aux tests écrits avec le package **testing**.

Utiliser **testing** (et **go-cmp**) permet d’écrire des tests : **compréhensibles**, **idiomatic Go**, **sans surcharge cognitive** .

### Tests instables

Les tests instables sont des tests qui peuvent réussir ou échouer de manière aléatoire, même si le système est correct. Cela réduit la fiabilité de la suite de tests.

**Quelques exemples de causes fréquentes des tests instables :**

- Problèmes de timing

--- Les temps d’attente fixes sont risqués et doivent être évités. <br>
--- Préférer une attente conditionnelle minimale. <br>
--- Lorsqu’on teste des durées, utiliser les plus courtes possibles (ex. : 1 ms au lieu de 1 s).

- Dépendances au temps réel

--- Des tests peuvent échouer selon l’heure du jour. <br>
--- Solution : injecter une fausse fonction Now() pour contrôler le temps dans les tests.

- Ordre non déterministe

--- Exemple : les maps en Go, dont l’ordre d’itération n’est pas garanti. <br>
--- Solution : ne pas supposer d’ordre; comparer de manière adaptée.

**Pourquoi c’est un problème sérieux ?**

- Les tests instables affaiblissent la valeur des tests.
- Ils deviennent rapidement ignorés ou désactivés par les développeurs.
- Si vous entendez « ce test échoue parfois », il est déjà inutile.

**Que faire ?**

- Corriger la cause si possible.
- Sinon, supprimer le test : un test instable est pire que pas de test.

**Test instable** : échoue aléatoirement, sans lien clair avec le code. <br>
**Test fragile** : échoue après un changement sans rapport direct. Souvent dû à un couplage excessif dans le code testé.

### Environnement partagé

Un environnement partagé (ex. : base de données commune entre plusieurs tests) est une source fréquente de tests instables, car :

- L’état peut être modifié par d’autres tests avant ou pendant l’exécution.
- Le test ne maîtrise plus l’état initial, ce qui rend les résultats non fiables.

**Solution idéale** : chaque test doit avoir son propre environnement isolé, avec un état initial connu et indépendant.

**Si l’isolation complète n’est pas possible :**

- Travailler sans modifier l’environnement partagé
- Utiliser des ressources dédiées par test : <br>
--- Créer une base de données propre par test, ou au moins une table unique. <br>
--- Utiliser une transaction et faire un rollback à la fin pour restaurer l’état initial.
- Ne jamais supprimer ce qu’on n’a pas créé.
- Ne jamais automatiser la suppression globale d’une base de données : un bug peut effacer la base de **production** par erreur.

### Tests en échec

Quand des tests échouent en permanence sans être corrigés, cela détruit la confiance qu’on peut avoir dans la suite de tests. Résultat :

- Les développeurs commencent à ignorer les tests.
- La suite devient inutile, même si d’autres tests sont valides.

**Principe fondamental : tolérance zéro**

- Aucun test ne doit échouer.
- La correction d’un test en échec doit devenir la priorité absolue, avant tout autre changement.

**Pourquoi cette rigueur ?**

- Les tests doivent garantir que le produit est toujours livrable.
- Corriger un bug maintenant coûte moins cher que plus tard.
- Tolérer les erreurs crée une dette technique risquée.

**Et si on a déjà trop d’erreurs ?**

- Fermer tous les vieux bugs et supprimer tous les tests cassés.
- Les vrais problèmes réapparaîtront naturellement via les utilisateurs ou la pratique.

### Tests lents

Une suite de tests trop lente devient vite inutilisable, car elle ralentit les retours et finit par être ignorée.

**Combien de temps est "trop long" ?**

- Objectif idéal : retour sous 5 minutes.
- Limite psychologique (appelée **temps de Beck**) : 10 minutes. Au-delà, tout le monde s’accorde à dire qu’il faut agir.

**7 stratégies pour accélérer les tests :**

- Tests parallèles

Refactorisez pour supprimer les dépendances partagées, et permettre l’exécution en parallèle.

- Évitez les E/S inutiles

Utilisez des systèmes de fichiers en mémoire (comme fstest.MapFS) et des io.Reader/Writer plutôt que des fichiers disques.

- Évitez les appels réseau

Remplacez les API distantes par des faux locaux pour gagner en rapidité.

- Partage intelligent des fixtures

Réutilisez les initialisations coûteuses entre tests, sans compromettre la stabilité.

- Pas de veilles fixes (sleep)

Préférez des attentes conditionnelles (« attendre le succès ») pour réduire les délais inutiles.

- Mieux vaut du matériel rapide que du code lent

Si l’optimisation ne suffit pas, exécutez les tests sur des machines très puissantes (ex : cloud à 256 cœurs).

- Extraire les tests lents

Si certains tests restent lents, déplacez-les dans une suite à part, exécutée la nuit par exemple.
