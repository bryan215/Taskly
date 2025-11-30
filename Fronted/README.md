# Frontend - Gestor de Tareas

Frontend básico para interactuar con la API de tareas.

## Características

- ✅ Crear nuevas tareas
- 🔍 Buscar tarea por ID
- 📋 Ver todas las tareas
- ✓ Completar/Descompletar tareas
- 🗑 Eliminar tareas

## Cómo usar

1. Asegúrate de que el backend esté corriendo en `http://localhost:8080`
2. Abre `index.html` en tu navegador
3. ¡Listo! Ya puedes gestionar tus tareas

## Archivos

- `index.html` - Estructura HTML
- `style.css` - Estilos CSS
- `app.js` - Lógica JavaScript y llamadas a la API

## Endpoints utilizados

- `POST /api/v1/tasks` - Crear tarea
- `GET /api/v1/tasks/:id` - Obtener tarea por ID
- `GET /api/v1/tasks` - Obtener todas las tareas
- `DELETE /api/v1/tasks/:id` - Eliminar tarea
- `PATCH /api/v1/task/completed` - Marcar como completada/no completada

