import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable } from 'rxjs';

import { RentContract } from '@core/rent-contract/rent-contract.model';
import {
  CreateRentContractRequest,
  RentContractFormValue,
  UpdateRentContractRequest,
} from '../model/rent-contract.model';
import { RentContractWriteApiService } from '@feature/rent-contract/service/rent-contract-write-api.service';

@Injectable()
export class RentContractWriteService {
  private readonly api = inject(RentContractWriteApiService);

  private readonly isLoadingWritable = signal(false);
  private readonly errorWritable = signal<string | null>(null);

  readonly isLoadingSignal = computed(() => this.isLoadingWritable());
  readonly errorSignal = computed(() => this.errorWritable());

  save$(rentContract: RentContractFormValue): Observable<RentContract> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: CreateRentContractRequest = {
      ...rentContract,
    };

    return this.api.create$(request).pipe(finalize(() => this.isLoadingWritable.set(false)));
  }

  update$(id: number, rentContract: RentContractFormValue): Observable<RentContract> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: UpdateRentContractRequest = {
      ...rentContract,
    };

    return this.api.update$(id, request).pipe(finalize(() => this.isLoadingWritable.set(false)));
  }
}
