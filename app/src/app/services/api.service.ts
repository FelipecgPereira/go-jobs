import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthResponse, B2bPayload, B2bSummary, Customer, Project, UserPayload } from '../models';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private readonly apiUrl = 'http://localhost:3000';

  constructor(private http: HttpClient) {}

  isAuthenticated(): boolean {
    return !!localStorage.getItem('gojobs_token');
  }

  logout(): void {
    localStorage.removeItem('gojobs_token');
  }

  login(credentials: { email: string; password: string }): Observable<AuthResponse & { token?: string }> {
    return this.http.post<AuthResponse & { token?: string }>(`${this.apiUrl}/auth`, credentials);
  }

  register(user: UserPayload): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.apiUrl}/signup`, user);
  }

  getCustomers(): Observable<Customer[]> {
    return this.http.get<any[]>(`${this.apiUrl}/customer`, { headers: this.getHeaders() }).pipe(
      map((items) =>
        items.map((item) => ({
          id: item.Id ?? item.id,
          name: item.Name ?? item.name,
          email: item.Email ?? item.email,
          phone: item.Phone ?? item.phone,
          userID: item.UserID ?? item.userID,
        })),
      ),
    );
  }

  createCustomer(customer: Customer): Observable<any> {
    return this.http.post(`${this.apiUrl}/customer`, customer, { headers: this.getHeaders() });
  }

  updateCustomer(id: number, customer: Customer): Observable<any> {
    return this.http.put(`${this.apiUrl}/customer/${id}`, customer, { headers: this.getHeaders() });
  }

  getProjects(): Observable<Project[]> {
    return this.http.get<any[]>(`${this.apiUrl}/project`, { headers: this.getHeaders() }).pipe(
      map((items) =>
        items.map((item) => ({
          id: item.Id ?? item.id,
          name: item.Name ?? item.name,
          price: item.Price ?? item.price,
          startDate: item.StartDate ?? item.startDate,
          endDate: item.EndDate ?? item.endDate,
          userId: item.UserId ?? item.userId,
        })),
      ),
    );
  }

  createProject(project: Project): Observable<any> {
    return this.http.post(`${this.apiUrl}/project/`, project, { headers: this.getHeaders() });
  }

  updateProject(id: number, project: Project): Observable<any> {
    return this.http.put(`${this.apiUrl}/project/${id}`, project, { headers: this.getHeaders() });
  }

  createB2b(payload: B2bPayload): Observable<any> {
    return this.http.post(`${this.apiUrl}/b2b/`, payload, { headers: this.getHeaders() });
  }

  updateB2b(id: number, payload: B2bPayload): Observable<any> {
    return this.http.put(`${this.apiUrl}/b2b/${id}`, payload, { headers: this.getHeaders() });
  }

  getB2bSummary(status: string, start: string, end: string): Observable<B2bSummary> {
    const params = new HttpParams().set('status', status).set('start', start).set('end', end);
    return this.http.get<B2bSummary>(`${this.apiUrl}/b2b/sum`, { headers: this.getHeaders(), params });
  }

  private getHeaders(): HttpHeaders {
    const token = localStorage.getItem('gojobs_token') ?? '';
    return new HttpHeaders({
      'Content-Type': 'application/json',
      Authorization: token,
    });
  }
}
