
# push-swap
Si vous travaillez sur un projet Git avec plusieurs branches vous pouvez suivre ces étapes :

1. **Assurez-vous d'être sur la bonne branche :**
   Avant de commencer à travailler, assurez-vous que vous êtes sur la branche "racabab" en utilisant la commande suivante :
   ```bash
   git checkout nom_branch
   ```

2. **Récupérez les dernières modifications depuis la branche distante :**
   Avant de commencer à travailler, récupérez les dernières modifications depuis la branche distante pour vous assurer que vous travaillez sur la version la plus récente :
   ```bash
   git pull origin nom_branch
   ```

3. **Effectuez vos modifications :**
   Faites les modifications nécessaires dans vos fichiers.

4. **Ajoutez les modifications pour le commit :**
   Ajoutez les fichiers modifiés à l'index en utilisant la commande :
   ```bash
   git add nom_du_fichier_modifié
   ```

   Vous pouvez également ajouter tous les fichiers modifiés en une seule fois avec :
   ```bash
   git add .
   ```

5. **Effectuez le commit :**
   Une fois que vous avez ajouté les modifications à l'index, effectuez le commit avec un message descriptif :
   ```bash
   git commit -m "Votre message de commit ici"
   ```

6. **Poussez les modifications vers la branche distante :**
   Enfin, poussez vos modifications vers la branche "racabab" sur le dépôt distant :
   ```bash
   git push origin nom_branch
   ```
