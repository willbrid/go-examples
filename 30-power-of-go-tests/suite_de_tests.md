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