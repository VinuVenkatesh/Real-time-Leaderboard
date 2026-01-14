import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivityResponse } from '../models/activity';

@Injectable({
  providedIn: 'root',
})
export class Activities {
  constructor(private http: HttpClient) {}

  url = 'http://localhost:1234/'

  getAllActivities() {
    const token = localStorage.getItem('jwtToken');
    const headers = { 'Authorization': token || '' };
    return this.http.get<ActivityResponse>(this.url + "listactivities", { headers });
  }
  
}
