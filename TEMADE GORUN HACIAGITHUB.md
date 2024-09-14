
 2000  go run ./cmd/api
 2001  git remote -v
 2002  git remote rm origin
 2003  git remote add origin git@github.diego-all:diego-all/products-API.git 
 2004  git remote -v

 2023  git tag -a v0.1 -m "initial version" f824e
 2024  git tag
 2025  git show-ref --tags
 2026  git status
 2027  git pull origin feature/product
 2028  git push origin --tags




github.com/diego-all/products-API



go install github.com/diego-all/products-API@latest

root@pho3nix:/home/diegoall/MAESTRIA_ING# go install github.com/diego-all/products-API@latest
go: github.com/diego-all/products-API@latest: module github.com/diego-all/products-API@latest found (v0.0.0-20240608063122-f824ec00d01a), but does not contain package github.com/diego-all/products-API


No funciona al parecer es como para descargar paquete e importarlo o usarlo.



DESPUES DE CREAR EL TAG SIGUE IGUAL


root@pho3nix:/home/diegoall/MAESTRIA_ING/PRUEBA# go run github.com/diego-all/products-API@latest
go: github.com/diego-all/products-API@latest: module github.com/diego-all/products-API@latest found (v0.0.0-20240608063122-f824ec00d01a), but does not contain package github.com/diego-all/products-API


segun goxygen



