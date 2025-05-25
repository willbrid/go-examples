# Langage testscript

Les scripts de test utilisés avec testscript ne sont ni des scripts shell ni du code Go, mais sont écrits dans un langage spécifique (DSL) conçu pour tester des programmes en ligne de commande. Ce DSL n’est pas un langage généraliste, mais un outil simple et élégant pour écrire des tests automatisés.

### Négation d'une assertion avec le préfixe !

Nous pouvons également affirmer que la sortie du programme ne doit pas correspondre, en préfixant la ligne stdout par le caractère **!**.

```
exec echo 'hello world'
! stdout 'hello world\n'
```

Nous pouvons également utiliser l'opérateur ! pour nier d'autres assertions, par exemple **exec**.
Un **exec** seul affirme que le programme donné réussit, c'est-à-dire qu'il renvoie un état de sortie **nul** (**zéro**). L'effet de l'opérateur **! exec** est d'affirmer que le programme échoue, c'est-à-dire que son état de sortie n'est pas **nul** (**différent de zéro**).

Pourquoi vouloir affirmer que l'exécution d'une commande donnée échoue ? C'est le comportement attendu d'un outil en ligne de commande lorsque l'utilisateur spécifie un indicateur ou un argument non valide, par exemple. Dans ce cas, il est courant d'afficher un message d'erreur et de quitter l'application avec un code de sortie différent de zéro. L'assertion correspondante pour l'erreur standard correspondante est nommée **stderr**.

```
! exec cat doesntexist
stderr 'cat: doesntexist: No such file or directory'
```

 Dans ce cas, pour éviter que **cat** affiche quoi que ce soit sur sa sortie standard en cas d'erreur, on peut utiliser **! stdout** :

 ```
! exec cat doesntexist
stderr 'cat: doesntexist: No such file or directory'
! stdout .
 ```