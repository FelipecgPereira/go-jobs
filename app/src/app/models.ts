export interface AuthResponse {
  message: string;
  token?: string;
}

export interface UserPayload {
  name: string;
  email: string;
  password: string;
}

export interface Customer {
  id?: number;
  name: string;
  email: string;
  phone: string;
  userID?: number;
}

export interface Project {
  id?: number;
  name: string;
  price: number;
  startDate: string;
  endDate: string;
  userId?: number;
}

export interface B2bPayload {
  customerId: number;
  projectId?: number;
  status: string;
}

export interface B2bSummary {
  total: number;
}
