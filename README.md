# gotolang
Création d'un langage de programmation interprété en golang

## Exemple

- Téléchargez l'executable correspondant à votre OS dans la dernière release
- Mettez le chemin de l'executable dans votre variable d'environnement %PATH%
- Créez un fichier `test.gtl` n'importe où sur votre système
- Remplissez le avec le code suivant :
```go
const var = "test";
const var2 = 10;
   const var3 = 15;

print(var);
print(var2);
print(var3);

func t(toto: string, test: int) => null;

// func toto() {
//  return null;
// };

// test de commentaire

// test de commentaires
// sur plusieurs lignes
```
- lancez la commande :
```bash
gotolang /path/to/test.gtl
```
