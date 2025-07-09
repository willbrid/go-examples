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
