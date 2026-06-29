import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { Role } from '@core/role/role.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class RoleApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/roles');
  }

  listRoles$(page = 1, limit = 10): Observable<ApiResponse<Role[]>> {
    const params = new HttpParams().set('page', page).set('limit', limit);
    return this.http
      .get<ApiResponse<Role[]>>(this.baseUrl, { params })
      .pipe(map((response) => response));
  }
}
