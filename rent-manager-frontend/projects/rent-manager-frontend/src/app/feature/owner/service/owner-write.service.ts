import { inject, Injectable, signal } from '@angular/core';
import { finalize, Observable } from 'rxjs';

import { OwnerWriteApiService } from './owner-write-api.service';
import { Owner } from '@core/owner/owner.model';
import {
  CreateOwnerRequest,
  OwnerFormValue,
  UpdateOwnerRequest,
} from '@feature/owner/model/owner.model';

@Injectable({ providedIn: 'root' })
export class OwnerWriteService {
  private readonly ownerApiService = inject(OwnerWriteApiService);

  readonly isLoadingSignal = signal(false);
  readonly errorSignal = signal<string | null>(null);

  save$(item: OwnerFormValue): Observable<Owner> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    const request: CreateOwnerRequest = {
      name: item.name,
      email: item.email,
      document_number: item.document_number,
      phone: item.phone,
    };

    return this.ownerApiService
      .create(request)
      .pipe(finalize(() => this.isLoadingSignal.set(false)));
  }

  update$(id: number, item: OwnerFormValue): Observable<Owner> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    const request: UpdateOwnerRequest = {
      name: item.name,
      email: item.email,
      document_number: item.document_number,
      phone: item.phone,
    };

    return this.ownerApiService
      .update(id, request)
      .pipe(finalize(() => this.isLoadingSignal.set(false)));
  }
}
