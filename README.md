# db-lab3

Este proyecto contiene el backend del laboratorio 3 de base de datos.

## 📁 Estructura del proyecto

Este repositorio representa solo el backend. Para funcionar correctamente, debe estar en el mismo nivel de carpeta que el frontend (`db-lab3-front`), ya que el archivo `docker-compose.yml` está dentro del frontend y desde ahí se levantan ambos servicios.

```
.
├── db-lab3-front/      # Frontend con docker-compose.yml
└── db-lab3/            # Este repositorio (backend)
```

## 🚀 Instrucciones para levantar el sistema completo

```bash
# Clonar ambos repositorios uno junto al otro
git clone https://github.com/vicperezch/db-lab3-front.git
git clone https://github.com/vicperezch/db-lab3.git

# Ir a la carpeta del frontend
cd db-lab3-front

# Levantar todo el sistema (frontend y backend)
docker-compose up --build
```

Cuando todo esté levantado, abre tu navegador y accede a:

```
http://localhost:3000
```

> El backend corre dentro del contenedor y se comunica automáticamente con el frontend.

Para detener los servicios:

```bash
docker-compose down
```

## ✅ Notas

- No necesitas ejecutar nada manualmente dentro de `db-lab3`.
- Asegúrate de que los puertos 3000 (frontend) y 8080 (backend) estén libres.
- Toda la configuración de red y volúmenes está en el `docker-compose.yml` del frontend.

---

Backend del laboratorio 3 – se integra automáticamente con `db-lab3-front` usando Docker.
