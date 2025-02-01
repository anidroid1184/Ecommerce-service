# Ecommerce App

Backend: Go (Gin + SQLite/GORM).  
Frontend: React (Vite + TailwindCSS).  

## Características
- **Backend**: Crear y listar productos mediante `/api/products`.  
- **Frontend**: Agregar y ver productos.  

## Requisitos
- Go (v1.20 o superior).  
- Node.js (v18 o superior).  

## Configuración
### Backend
1. `cd backend`  
2. `go mod tidy`  
3. `go run main.go` → Se ejecuta en `http://localhost:8080`.  

### Frontend
1. `cd frontend`  
2. `npm install`  
3. `npm run dev` → Se ejecuta en `http://localhost:5173`.  

## Uso
1. Abre `http://localhost:5173`.  
2. Agrega productos usando el formulario.  
3. Visualiza los productos en la lista.  

---

**Nota**: Asegúrate de que tanto el backend como el frontend estén en ejecución.
