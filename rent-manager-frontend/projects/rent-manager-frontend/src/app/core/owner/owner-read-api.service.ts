import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { Owner } from '@core/owner/owner.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class OwnerReadApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/owners');
  }

  search(value: string, page: number, limit: number): Observable<Owner[]> {
    const params = new HttpParams().set('value', value).set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<Owner[]>>(`${this.baseUrl}/search`, { params })
      .pipe(map((response) => response.data));
  }

  list(page: number, limit: number): Observable<Owner[]> {
    const params = new HttpParams().set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<Owner[]>>(this.baseUrl, { params })
      .pipe(map((response) => response.data));
  }

  getById(id: number): Observable<Owner> {
    return this.http
      .get<ApiResponse<Owner>>(`${this.baseUrl}/${id}`)
      .pipe(map((response) => response.data));
  }
}
