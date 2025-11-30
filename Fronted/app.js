const API_BASE_URL = 'http://localhost:8080/api/v1';

// Mostrar mensaje
function showMessage(text, type = 'success') {
    const messageEl = document.getElementById('message');
    messageEl.textContent = text;
    messageEl.className = `message ${type} show`;
    
    setTimeout(() => {
        messageEl.classList.remove('show');
    }, 3000);
}

// Crear tarea
document.getElementById('createTaskForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const title = document.getElementById('taskTitle').value;
    const completed = document.getElementById('taskCompleted').checked;
    
    try {
        const response = await fetch(`${API_BASE_URL}/tasks`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                completed: completed
            })
        });
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Error al crear la tarea');
        }
        
        const task = await response.json();
        showMessage(`Tarea "${task.title}" creada exitosamente!`, 'success');
        document.getElementById('createTaskForm').reset();
        loadAllTasks();
    } catch (error) {
        showMessage(error.message, 'error');
    }
});

// Buscar tarea por ID
document.getElementById('getTaskForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const id = document.getElementById('taskId').value;
    const resultDiv = document.getElementById('taskResult');
    
    try {
        const response = await fetch(`${API_BASE_URL}/tasks/${id}`);
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Tarea no encontrada');
        }
        
        const task = await response.json();
        resultDiv.innerHTML = `
            <div class="task-item ${task.completed ? 'completed' : ''}">
                <div class="task-info">
                    <div class="task-title">${task.title}</div>
                    <div class="task-id">ID: ${task.id}</div>
                    <span class="task-status ${task.completed ? 'completed' : 'pending'}">
                        ${task.completed ? '✓ Completada' : '⏳ Pendiente'}
                    </span>
                </div>
            </div>
        `;
        showMessage('Tarea encontrada!', 'success');
    } catch (error) {
        resultDiv.innerHTML = `<p style="color: #dc3545;">${error.message}</p>`;
        showMessage(error.message, 'error');
    }
});

// Cargar todas las tareas
async function loadAllTasks() {
    const tasksList = document.getElementById('tasksList');
    tasksList.innerHTML = '<p>Cargando...</p>';
    
    try {
        const response = await fetch(`${API_BASE_URL}/tasks`);
        
        if (!response.ok) {
            throw new Error('Error al cargar las tareas');
        }
        
        const data = await response.json();
        const tasks = data.tasks || [];
        
        if (tasks.length === 0) {
            tasksList.innerHTML = '<div class="empty-state">No hay tareas disponibles</div>';
            return;
        }
        
        tasksList.innerHTML = tasks.map(task => `
            <div class="task-item ${task.completed ? 'completed' : ''}">
                <div class="task-info">
                    <div class="task-title">${task.title}</div>
                    <div class="task-id">ID: ${task.id}</div>
                    <span class="task-status ${task.completed ? 'completed' : 'pending'}">
                        ${task.completed ? '✓ Completada' : '⏳ Pendiente'}
                    </span>
                </div>
                <div class="task-actions">
                    <button class="btn-small btn-success" onclick="toggleTask(${task.id}, ${!task.completed})">
                        ${task.completed ? '↩ Desmarcar' : '✓ Completar'}
                    </button>
                    <button class="btn-small btn-danger" onclick="deleteTask(${task.id})">
                        🗑 Eliminar
                    </button>
                </div>
            </div>
        `).join('');
    } catch (error) {
        tasksList.innerHTML = `<p style="color: #dc3545;">${error.message}</p>`;
        showMessage(error.message, 'error');
    }
}

// Toggle completar/descompletar tarea
async function toggleTask(id, status) {
    try {
        const response = await fetch(`${API_BASE_URL}/tasks/${id}/completed`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                status: status
            })
        });
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Error al actualizar la tarea');
        }
        
        showMessage(`Tarea ${status ? 'completada' : 'desmarcada'} exitosamente!`, 'success');
        loadAllTasks();
    } catch (error) {
        showMessage(error.message, 'error');
    }
}

// Eliminar tarea
async function deleteTask(id) {
    if (!confirm('¿Estás seguro de que quieres eliminar esta tarea?')) {
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Error al eliminar la tarea');
        }
        
        const data = await response.json();
        showMessage(data.message || 'Tarea eliminada exitosamente!', 'success');
        loadAllTasks();
    } catch (error) {
        showMessage(error.message, 'error');
    }
}

// Cargar tareas al iniciar
loadAllTasks();

