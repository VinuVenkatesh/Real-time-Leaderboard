import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class UserAuth {
  constructor(private http: HttpClient) {}
  url = 'http://localhost:1234/'

  register(username: string, password: string) {
    return this.http.post(this.url +"register", { username, password });
  }

  login(username: string, password: string) {
    return this.http.post(this.url +"login", { username, password });
  }
}
