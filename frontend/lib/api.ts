import type { Task, CreateTaskRequest } from '@/types';
import { cookieUtils } from './cookies';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;
    const token = cookieUtils.getToken();
    
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
      ...options.headers,
    };

    // Agregar token si existe
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }

    const config: RequestInit = {
      headers,
      ...options,
    };

    const response = await fetch(url, config);

    // Manejar error 401 (no autorizado) - redirigir al login
    if (response.status === 401) {
      cookieUtils.clearAuth();
      if (typeof window !== 'undefined') {
        window.location.href = '/login';
      }
      const error = await response.json().catch(() => ({}));
      throw new Error(error.error || 'Sesión expirada. Por favor inicia sesión nuevamente.');
    }

    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      throw new Error(error.error || `HTTP error! status: ${response.status}`);
    }

    return response.json();
  }

  // Auth
  async login(credentials: { username: string; password: string }) {
    return this.request<{ token: string }>(
      '/users/login',
      {
        method: 'POST',
        body: JSON.stringify(credentials),
      }
    );
  }

  async register(userData: { username: string; email: string; password: string }) {
    return this.request<{ message: string }>(
      '/users/register',
      {
        method: 'POST',
        body: JSON.stringify(userData),
      }
    );
  }

  // Tasks
  async getMyTasks() {
    return this.request<{ tasks: Task[] }>('/tasks');
  }

  async createTask(taskData: CreateTaskRequest) {
    return this.request<Task>('/tasks', {
      method: 'POST',
      body: JSON.stringify(taskData),
    });
  }

  async getAllTasks() {
    return this.request<{ tasks: Task[] }>('/tasks');
  }

  async getTaskById(id: number) {
    return this.request<Task>(`/tasks/${id}`);
  }

  async updateTaskStatus(id: number, status: boolean) {
    return this.request<Task>(`/tasks/${id}/completed`, {
      method: 'PATCH',
      body: JSON.stringify({ status }),
    });
  }

  async deleteTask(id: number) {
    return this.request<{ message: string }>(`/tasks/${id}`, {
      method: 'DELETE',
    });
  }
}

export const apiClient = new ApiClient(API_BASE_URL);

