import { HttpErrorResponse } from '@angular/common/http';
import { computed, inject, Injectable, signal } from '@angular/core';
import { catchError, EMPTY, map, Observable, take, tap } from 'rxjs';

import { Country, State } from '@core/location/location.model';
import { LocationApiService } from '@core/location/location-api.service';

@Injectable({ providedIn: 'root' })
export class LocationService {
  private readonly api = inject(LocationApiService);

  private readonly countriesWritable = signal<Country[]>([]);
  private readonly statesWritable = signal<State[]>([]);

  readonly countriesSignal = computed(() => this.countriesWritable());
  readonly statesSignal = computed(() => this.statesWritable());

  listCountries$(): Observable<Country[]> {
    return this.api.listCountries$().pipe(
      take(1),
      tap((apiResponse) => {
        this.countriesWritable.set(apiResponse.data);
      }),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  listStates$(countryId: number): Observable<State[]> {
    return this.api.listStates$(countryId).pipe(
      take(1),
      tap((apiResponse) => {
        this.statesWritable.set(apiResponse.data);
      }),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  getCountryById(id: number): Country | null {
    return this.countriesSignal().find((country) => country.id === id) ?? null;
  }

  getStateById(id: number): State | null {
    return this.statesSignal().find((state) => state.id === id) ?? null;
  }
}
