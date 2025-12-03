export interface User {
  id: number;
  username: string;
  email: string;
}

export interface Task {
  id: number;
  user_id: number;
  title: string;
  completed: boolean;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface CreateTaskRequest {
  title: string;
  completed: boolean;
}





