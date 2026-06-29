import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { LoginRequest, LoginResponse } from './auth.models';
import { ApiResponse } from '@core/api/api.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class AuthApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/auth/login');
  }

  login(request: LoginRequest): Observable<LoginResponse> {
    return this.http
      .post<ApiResponse<LoginResponse>>(`${this.baseUrl}`, request)
      .pipe(map((response) => response.data));
  }
}
