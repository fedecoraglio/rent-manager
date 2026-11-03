# Releases y despliegue

Argo CD sigue `develop`, en `platform/kubernetes`, con sincronización automática.
El workflow **Release Images** se ejecuta al subir un tag `vMAJOR.MINOR.PATCH`
(por ejemplo, `v1.0.0`) cuyo commit pertenezca al historial de `develop`.
Un push normal a `develop` ya no publica imágenes.

## Publicar una versión

Primero subir los cambios, incluido el workflow, a `develop`. Desde un checkout
limpio y actualizado, elegir una versión nueva (no reutilizar ni mover tags):

```powershell
git switch develop
git pull --ff-only origin develop
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

Reemplazar `v1.0.0` por la versión que corresponda. No es necesario crear además
una GitHub Release. Publicar una versión por vez y esperar a que termine.

El workflow construye backend y frontend desde el mismo tag y publica ambas
imágenes en GHCR. Solo después de que ambas compilaciones terminen correctamente,
actualiza los dos manifiestos en un commit sobre `develop`. Cada referencia
incluye el tag de release y el digest exacto de la imagen publicada.
Argo detecta ese commit y reemplaza los pods automáticamente. No hace falta
`rollout restart`. Los manifiestos conservan `:develop` hasta la primera release
exitosa, para no apuntar a una imagen que todavía no existe.

Después del workflow, actualizar el checkout local para incorporar el commit
automático antes de seguir trabajando:

```powershell
git pull --ff-only origin develop
kubectl -n argocd get application rent-manager
kubectl -n rent-manager get deployments -o wide
```

## Permisos y fallos

El workflow usa `GITHUB_TOKEN` con `contents: write` y `packages: write`.
Las reglas de protección de `develop` deben permitir el commit del workflow.
Si la rama exige PR o prohíbe ese push, el workflow fallará en el último paso;
hay que adaptar la publicación a un PR, sin desactivar las protecciones.
No requiere credenciales del cluster en GitHub.

Si falla una compilación, los manifiestos no se actualizan. Si falla el push por
un conflicto, revisar el fallo antes de reintentar. Un workflow exitoso publica
la versión deseada; comprobar también que Argo esté `Synced` y `Healthy`.

Para volver a una versión anterior, recuperar las dos referencias de imagen
(tag y digest) del commit de despliegue anterior y subir ese cambio a `develop`.
Argo aplicará esas referencias. No mover un tag existente.
