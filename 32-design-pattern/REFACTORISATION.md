# Refactorisation

La refactorisation est un processus systématique d'amélioration du code existant sans ajouter de nouvelles fonctionnalités, visant à transformer un code complexe ou désordonné en un code propre et une conception plus simple.

### Code propre

L'objectif principal de la refactorisation est de lutter contre la dette technique. Il transforme un code complexe ou désordonné en code propre et en conception simple.

Un code propre est un code facile à lire, à comprendre et à maintenir. Un code propre rend le développement logiciel prévisible et augmente la qualité du produit final.

**Quelques caractéristiques d'un code propre** :

- un code propre est évident pour les autres programmeurs
- un code propre ne contient pas de doublons
- un code propre contient un nombre minimal de classes et d’autres éléments mobiles
- un code propre passe tous les tests (le code est sale lorsque seulement **95%** des tests ont réussi)
- un code propre est plus facile et moins cher à maintenir

### Dette technique

**Les causes de la dette technique** :

- **Pression commerciale** : parfois, les circonstances commerciales peuvent nous obliger à déployer des fonctionnalités avant qu'elles ne soient complètement terminées.

- **manque de compréhension des conséquences de la dette technique** : parfois, notre employeur peut ne pas comprendre que la dette technique présente un « intérêt » dans la mesure où elle ralentit le rythme de développement à mesure que la dette s’accumule. Cela peut rendre trop difficile de consacrer le temps de l’équipe au refactoring car la direction n’en voit pas l’intérêt.

- **ne pas lutter contre la stricte cohérence des composants** : cela se produit lorsque le projet ressemble à un monolithe plutôt qu’au produit de modules individuels. Dans ce cas, toute modification apportée à une partie du projet affectera les autres. Le développement de l’équipe est rendu plus difficile car il est difficile d’isoler le travail de chaque membre.

- **manque de tests** : certains changements sont implémentés et déployés directement dans la production sans aucun test préalable. Les conséquences peuvent être catastrophiques. Par exemple, un correctif apparemment innocent peut envoyer un étrange e-mail de test à des milliers de clients ou, pire encore, vider ou corrompre une base de données entière.

- **manque de documentation** : cela ralentit l'intégration de nouvelles personnes dans le projet et peut interrompre le développement si des personnes clés quittent le projet.

- **manque d’interaction entre les membres de l’équipe** : si la base de connaissances n’est pas distribuée dans toute l’entreprise, les gens finiront par travailler avec une compréhension obsolète des processus et des informations sur le projet. Cette situation peut être exacerbée lorsque les développeurs juniors sont mal formés par leurs mentors.

- **développement simultané à long terme dans plusieurs branches** : cela peut conduire à l'accumulation d'une dette technique, qui augmente ensuite lorsque les modifications sont fusionnées. Plus les modifications sont effectuées de manière isolée, plus la dette technique totale est importante.

- **refactorisation retardée** : les exigences du projet évoluent constamment et il peut arriver qu'à un moment donné, certaines parties du code soient obsolètes, deviennent lourdes et doivent être repensées pour répondre aux nouvelles exigences. <br>
D'un autre côté, les programmeurs du projet écrivent chaque jour un nouveau code qui fonctionne avec les parties obsolètes. Par conséquent, plus la refactorisation est retardée, plus le code dépendant devra être retravaillé à l'avenir.

- **manque de surveillance de la conformité** : cela se produit lorsque tous ceux qui travaillent sur le projet écrivent du code comme ils l'entendent.

- **Incompétence** : c’est lorsque le développeur ne sait tout simplement pas comment écrire un code décent.


### Quand refactoriser

- **Règle de trois**

--- lorsque nous faisons quelque chose pour la première fois, faisons-le. <br>
--- lorsque nous faisons quelque chose de similaire pour la deuxième fois, craignons de devoir le répéter, mais faisons quand même la même chose. <br>
--- lorsque nous faisons quelque chose pour la troisième fois, commençons à refactoriser.

- **Lors de l'ajout d'une fonctionnalité**

--- La refactorsation nous aide à comprendre le code des autres. Si nous devons gérer le code sale de quelqu'un d'autre, essayons de le refactoriser en premier. Un code propre est beaucoup plus facile à comprendre. Nous l'améliorerons non seulement pour nous-même, mais aussi pour ceux qui l'utiliseront après nous. <br>
--- La refactorsation facilite l'ajout de nouvelles fonctionnalités. Il est beaucoup plus facile d'apporter des modifications à un code propre.

- **Lors de la correction d'un bug**

--- Les bugs dans le code se trouvent dans les endroits les plus sombres et les plus sales du code. Nettoyons notre code et les erreurs se découvriront pratiquement d'elles-mêmes. <br>
--- Les managers apprécient la refactorisation proactive car il élimine le besoin de tâches de refactorisation spéciale par la suite.

- **Lors d'une revue de code**

--- La revue de code peut être la dernière chance de mettre de l'ordre dans le code avant qu'il ne soit disponible au public. <br>
--- Il est préférable d'effectuer ces revues en binôme avec un auteur. De cette façon, nous pourrons résoudre rapidement les problèmes simples et évaluer le temps nécessaire pour résoudre les problèmes plus difficiles.


### Comment refactoriser

La refactorisation doit être effectuée sous la forme d'une série de petits changements, chacun d'entre eux améliorant légèrement le code existant tout en laissant le programme en état de fonctionnement.

- Le code devrait être plus propre.
- De nouvelles fonctionnalités ne doivent pas être créées pendant la réfactorisation.
- Tous les tests existants doivent réussir après la refactorisation (un excellent moyen d'éviter des échecs de tests lors de la réfactorisation est d'écrire des tests de style **BDD**).


**Référence** : [Refactoring Guru](https://refactoring.guru/refactoring)