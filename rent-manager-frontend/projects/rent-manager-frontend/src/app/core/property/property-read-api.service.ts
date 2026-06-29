import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { Property, PropertySummary } from '@core/property/property.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class PropertyReadApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/properties');
  }

  list$(page: number, limit: number): Observable<Property[]> {
    const params = new HttpParams().set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<Property[]>>(this.baseUrl, { params })
      .pipe(map((response) => response.data));
  }

  listSummaries$(page: number, limit: number): Observable<ApiResponse<PropertySummary[]>> {
    const params = new HttpParams().set('page', page).set('limit', limit);

    return this.http.get<ApiResponse<PropertySummary[]>>(`${this.baseUrl}/summaries`, { params });
  }

  search$(value: string, page: number, limit: number): Observable<Property[]> {
    const params = new HttpParams().set('value', value).set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<Property[]>>(`${this.baseUrl}/search`, { params })
      .pipe(map((response) => response.data));
  }

  getById$(id: number): Observable<Property> {
    return this.http
      .get<ApiResponse<Property>>(`${this.baseUrl}/${id}`)
      .pipe(map((response) => response.data));
  }
}
