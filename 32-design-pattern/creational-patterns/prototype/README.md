# Prototype

Le **prototype** est un patron de conception de création qui permet de cloner des objets - même complexes - sans se coupler à leur classe.

Toutes les classes prototype devraient avoir une interface commune rendant possible la copie des objets, même sans connaître leur classe concrète. Les objets prototype peuvent créer des copies complètes puisqu’ils peuvent accéder aux attributs privés des autres objets de la même classe.

Ce modèle est particulièrement utile lorsque la création de nouveaux objets est coûteuse, ou lorsque nous voulons éviter de modifier l'objet original.

**Exemple :** Prenons l'exemple basé sur le système de fichiers d’un système d’exploitation (SE). Le système de fichier du SE est récursif : les dossiers contiennent des fichiers et des dossiers, qui peuvent eux-mêmes inclure des fichiers et des dossiers, et ainsi de suite.
Chaque fichier/dossier est représenté par une interface **inode**. L’interface inode possède également la fonction **clone**.
Les structures **file** et **folder** sont des structures à cloner et elles implémentent les fonctions **print** et **clone**, car elles sont du type **inode**.