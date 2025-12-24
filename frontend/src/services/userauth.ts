import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UserAuth {
  constructor(private http: HttpClient) {}

  url = 'http://localhost:1234/'
  private token: string | null = null;

  register(username: string, password: string) {
    return this.http.post(this.url +"register", { username, password });
  }

  login(username: string, password: string) {
    return this.http.post<{ token: string }>(this.url +"login", { username, password }).pipe(
      tap(response => {
        this.token = response.token;
        localStorage.setItem('jwtToken', response.token);
      })
    );
  }
  getToken(): string | null {
    return this.token || localStorage.getItem('jwtToken');
  }
}
