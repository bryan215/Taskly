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
  message: string;
  token: string;
  user: {
    id: number;
    username: string;
    email: string;
  };
}

export interface CreateTaskRequest {
  title: string;
  completed: boolean;
  user_id: number;
}





