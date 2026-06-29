import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable } from 'rxjs';

import {
  CreatePropertyRequest,
  PropertyFormValue,
  UpdatePropertyRequest,
} from '../model/property.model';
import { PropertyWriteApiService } from './property-write-api.service';
import { Property } from '@core/property/property.model';

@Injectable()
export class PropertyWriteService {
  private readonly propertyApi = inject(PropertyWriteApiService);

  private readonly isLoadingWritable = signal(false);
  private readonly errorWritable = signal<string | null>(null);

  readonly isLoadingSignal = computed(() => this.isLoadingWritable());
  readonly errorSignal = computed(() => this.errorWritable());

  save$(property: PropertyFormValue): Observable<Property> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: CreatePropertyRequest = {
      ...property,
    };

    return this.propertyApi
      .create$(request)
      .pipe(finalize(() => this.isLoadingWritable.set(false)));
  }

  update$(id: number, property: PropertyFormValue): Observable<Property> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: UpdatePropertyRequest = {
      ...property,
    };

    return this.propertyApi
      .update$(id, request)
      .pipe(finalize(() => this.isLoadingWritable.set(false)));
  }
}
