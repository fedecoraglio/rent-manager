import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { CreateUserRequest, UpdateUserRequest, User } from '../model/user.model';
import { ApiResponse } from '@core/api/api.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable()
export class UserApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/users');
  }

  create(request: CreateUserRequest): Observable<User> {
    return this.http
      .post<ApiResponse<User>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  list(page: number, limit: number): Observable<User[]> {
    const params = new HttpParams().set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<User[]>>(this.baseUrl, { params })
      .pipe(map((response) => response.data));
  }

  search(value: string, page: number, limit: number): Observable<User[]> {
    const params = new HttpParams().set('value', value).set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<User[]>>(`${this.baseUrl}/search`, { params })
      .pipe(map((response) => response.data));
  }

  getById(id: number): Observable<User> {
    return this.http
      .get<ApiResponse<User>>(`${this.baseUrl}/${id}`)
      .pipe(map((response) => response.data));
  }

  update(id: number, request: UpdateUserRequest): Observable<User> {
    return this.http
      .put<ApiResponse<User>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}/${id}`);
  }
}
