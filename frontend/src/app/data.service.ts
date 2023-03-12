import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  private baseUrl = 'http://localhost:8080/api/gas';

  constructor(private http: HttpClient) { }

  getDatasList(): Observable<any> {
    return this.http.get(`${this.baseUrl}`);
  }

  // createUser
  createData(data: Object): Observable<Object> {
    return this.http.post(`${this.baseUrl}`, data); 
  }

  // deleteUser
  deleteData(id: string): Observable<any> {
    return this.http.delete(`${this.baseUrl}/${id}`, 
                    { responseType: 'text' });
  }

  // updateUser
  updateData(id: string, value: any): 
    Observable<Object> {
  return this.http.put(`${this.baseUrl}/${id}`, 
                   value);
  }

  // getUser
  getData(id: string): Observable<any> {
    return this.http.get(`${this.baseUrl}/${id}`);
  }

}
