# gotolang
Création d'un langage de programmation interprété en golang

## Exemple

- Téléchargez l'executable correspondant à votre OS dans la dernière release
- Mettez le chemin de l'executable dans votre variable d'environnement %PATH%
- Créez un fichier `test.gtl` n'importe où sur votre système
- Remplissez le avec le code suivant :
```go
const _var = "test";
const _var2 = 10;
   const _var3 = 15;

print(_var);
print(_var2);
print(_var3);
```
- lancez la commande :
```bash
gotolang /path/to/test.gtl
```
