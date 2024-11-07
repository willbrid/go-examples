# Singleton

Le **Singleton** est un patron de conception de création qui s’assure de l’existence d’un seul objet de son genre et fournit un unique point d’accès vers cet objet.

**Exemple :** En général, une instance de singleton est créée lors de la première initialisation de la struct. De ce faire, nous prenons l'exemple de la définition de la méthode **getInstance** dans la struct. Cette méthode sera responsable de la création et du renvoi de l’instance du singleton. Une fois créée, cette même instance sera retournée chaque fois que **getInstance** est appelée.