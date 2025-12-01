# Taskly Frontend

Frontend moderno construido con Next.js 14, TypeScript y Tailwind CSS.

## Características

- ✅ Autenticación con login/registro
- ✅ Gestión de cookies para sesión de usuario
- ✅ Visualización de tareas del usuario
- ✅ Crear, editar y eliminar tareas
- ✅ UI moderna y responsive
- ✅ TypeScript para type safety
- ✅ Arquitectura limpia y escalable

## Instalación

```bash
npm install
```

## Configuración

Crea un archivo `.env.local` en la raíz del proyecto:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## Desarrollo

```bash
npm run dev
```

Abre [http://localhost:3000](http://localhost:3000) en tu navegador.

## Estructura del Proyecto

```
frontend/
├── app/              # App Router de Next.js
│   ├── login/        # Página de login
│   ├── register/     # Página de registro
│   └── tasks/        # Página de tareas
├── lib/              # Utilidades y servicios
│   ├── api.ts        # Cliente API
│   └── cookies.ts    # Utilidades de cookies
└── types/            # Definiciones TypeScript
    └── index.ts      # Interfaces y tipos
```

## Tecnologías

- **Next.js 14** - Framework React con App Router
- **TypeScript** - Type safety
- **Tailwind CSS** - Estilos utility-first
- **js-cookie** - Gestión de cookies del lado del cliente
