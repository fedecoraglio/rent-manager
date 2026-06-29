import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { Country, State } from '@core/location/location.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class LocationApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/locations');
  }

  listCountries$(): Observable<ApiResponse<Country[]>> {
    return this.http
      .get<ApiResponse<Country[]>>(`${this.baseUrl}/countries`)
      .pipe(map((response) => response));
  }

  listStates$(countryId: number): Observable<ApiResponse<State[]>> {
    return this.http
      .get<ApiResponse<State[]>>(`${this.baseUrl}/countries/${countryId}/states`)
      .pipe(map((response) => response));
  }
}
