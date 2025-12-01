const API_BASE_URL = 'http://localhost:8080/api/v1';
let simulatedUserId = null;

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

// Crear usuario
const createUserForm = document.getElementById('createUserForm');
if (createUserForm) {
    createUserForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const username = document.getElementById('userUsername').value.trim();
        const email = document.getElementById('userEmail').value.trim();
        const password = document.getElementById('userPassword').value;

        try {
            const response = await fetch(`${API_BASE_URL}/users/register`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, email, password })
            });

            if (!response.ok) {
                const error = await response.json();
                throw new Error(error.error || 'Error al crear usuario');
            }

            const user = await response.json();
            showMessage(`Usuario ${user.username} creado con éxito`, 'success');
            createUserForm.reset();
        } catch (error) {
            showMessage(error.message, 'error');
        }
    });
}

// Login de prueba
const loginForm = document.getElementById('loginForm');
if (loginForm) {
    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const username = document.getElementById('loginUsername').value.trim();
        const password = document.getElementById('loginPassword').value;
        const loginResult = document.getElementById('loginResult');

        try {
            const response = await fetch(`${API_BASE_URL}/users/login`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, password })
            });

            if (!response.ok) {
                const error = await response.json().catch(() => ({}));
                throw new Error(error.error || 'Credenciales inválidas');
            }

            const data = await response.json();
            showMessage(`Login exitoso. Hola ${data.user.username}`, 'success');

            if (loginResult) {
                loginResult.innerHTML = `
                    <p><strong>ID:</strong> ${data.user.id}</p>
                    <p><strong>Usuario:</strong> ${data.user.username}</p>
                    <p><strong>Email:</strong> ${data.user.email}</p>
                `;
            }

            // sincroniza el selector de usuario simulado
            const userIdInput = document.getElementById('simulatedUserId');
            if (userIdInput && data.user?.id) {
                simulatedUserId = data.user.id;
                userIdInput.value = simulatedUserId;
                loadUserTasks();
            }
        } catch (error) {
            if (loginResult) {
                loginResult.innerHTML = `<p style="color:#dc3545;">${error.message}</p>`;
            }
            showMessage(error.message, 'error');
        }
    });
}

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

async function fetchTasks() {
    const response = await fetch(`${API_BASE_URL}/tasks`);
    if (!response.ok) {
        let errorMessage = 'Error al cargar las tareas';
        try {
            const errorData = await response.json();
            errorMessage = errorData.error || errorMessage;
        } catch (_) {
            // ignorado
        }
        throw new Error(errorMessage);
    }

    const data = await response.json();
    return data.tasks || [];
}

function renderTasks(containerId, tasks, options = {}) {
    const container = document.getElementById(containerId);
    if (!container) return;

    const { showUser = false, showActions = true } = options;

    if (!tasks.length) {
        container.innerHTML = '<div class="empty-state">No hay tareas disponibles</div>';
        return;
    }

    container.innerHTML = tasks.map(task => `
        <div class="task-item ${task.completed ? 'completed' : ''}">
            <div class="task-info">
                <div class="task-title">${task.title}</div>
                <div class="task-id">ID: ${task.id}</div>
                ${showUser ? `
                <div class="task-user">
                    <span>👤 Usuario ID: ${task.user_id ?? 'N/A'}</span>
                    ${task.user && task.user.username ? `<span class="task-username">@${task.user.username}</span>` : ''}
                </div>` : ''}
                <span class="task-status ${task.completed ? 'completed' : 'pending'}">
                    ${task.completed ? '✓ Completada' : '⏳ Pendiente'}
                </span>
            </div>
            ${showActions ? `
            <div class="task-actions">
                <button class="btn-small btn-success" onclick="toggleTask(${task.id}, ${!task.completed})">
                    ${task.completed ? '↩ Desmarcar' : '✓ Completar'}
                </button>
                <button class="btn-small btn-danger" onclick="deleteTask(${task.id})">
                    🗑 Eliminar
                </button>
            </div>` : ''}
        </div>
    `).join('');
}

// Cargar todas las tareas (usuario + admin)
async function loadAllTasks() {
    const tasksList = document.getElementById('tasksList');
    const adminTasksList = document.getElementById('adminTasksList');

    tasksList.innerHTML = '<p>Cargando...</p>';
    if (adminTasksList) {
        adminTasksList.innerHTML = '<p>Cargando...</p>';
    }
    
    try {
        const tasks = await fetchTasks();
        renderTasks('tasksList', tasks, { showUser: false, showActions: true });
        if (adminTasksList) {
            renderTasks('adminTasksList', tasks, { showUser: true, showActions: false });
        }
    } catch (error) {
        tasksList.innerHTML = `<p style="color: #dc3545;">${error.message}</p>`;
        if (adminTasksList) {
            adminTasksList.innerHTML = `<p style="color: #dc3545;">${error.message}</p>`;
        }
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
setupSimulatedLogin();

function setupSimulatedLogin() {
    const form = document.getElementById('simulatedLoginForm');
    const input = document.getElementById('simulatedUserId');
    if (!form || !input) {
        return;
    }

    simulatedUserId = Number(input.value) || 1;
    input.value = simulatedUserId;

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        const value = Number(input.value);
        if (!value || value < 1) {
            showMessage('Ingresa un ID de usuario válido', 'error');
            return;
        }
        simulatedUserId = value;
        showMessage(`Usuario simulado: #${simulatedUserId}`, 'success');
        loadUserTasks();
    });

    loadUserTasks();
}

async function loadUserTasks() {
    const list = document.getElementById('userTasksList');
    if (!list) return;

    if (!simulatedUserId) {
        list.innerHTML = '<div class="empty-state">Selecciona un usuario para ver sus tareas.</div>';
        return;
    }

    list.innerHTML = '<p>Cargando...</p>';

    try {
        const response = await fetch(`${API_BASE_URL}/users/${simulatedUserId}/tasks`);

        if (!response.ok) {
            let errorMsg = 'Error al cargar las tareas del usuario';
            try {
                const error = await response.json();
                errorMsg = error.error || errorMsg;
            } catch (_) {
                // ignorado
            }
            throw new Error(errorMsg);
        }

        const data = await response.json();
        const tasks = data.tasks || [];

        renderTasks('userTasksList', tasks, { showUser: false, showActions: false });
        if (!tasks.length) {
            document.getElementById('userTasksList').innerHTML = `<div class="empty-state">El usuario #${simulatedUserId} no tiene tareas.</div>`;
        }
    } catch (error) {
        list.innerHTML = `<p style="color: #dc3545;">${error.message}</p>`;
        showMessage(error.message, 'error');
    }
}

