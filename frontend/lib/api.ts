import type { Task, CreateTaskRequest } from '@/types';

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
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    // Log de peticiones en desarrollo
    if (process.env.NODE_ENV === 'development') {
      console.log(`[API] ${options.method || 'GET'} ${url}`, options.body ? JSON.parse(options.body as string) : '');
    }

    const response = await fetch(url, config);

    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      if (process.env.NODE_ENV === 'development') {
        console.error(`[API] Error ${response.status} ${url}:`, error);
      }
      throw new Error(error.error || `HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    if (process.env.NODE_ENV === 'development') {
      console.log(`[API] Response ${url}:`, data);
    }
    return data;
  }

  // Auth
  async login(credentials: { username: string; password: string }) {
    return this.request<{ message: string; user: { id: number; username: string; email: string } }>(
      '/users/login',
      {
        method: 'POST',
        body: JSON.stringify(credentials),
      }
    );
  }

  async register(userData: { username: string; email: string; password: string }) {
    return this.request<{ id: number; username: string; email: string }>(
      '/users/register',
      {
        method: 'POST',
        body: JSON.stringify(userData),
      }
    );
  }

  // Tasks
  async getTasksByUserId(userId: number) {
    return this.request<{ tasks: Task[] }>(`/users/${userId}/tasks`);
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

