import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable } from 'rxjs';

import { RentPayment } from '@core/rent-payment/rent-payment.model';

import {
  CreateRentPaymentRequest,
  RentPaymentFormValue,
  UpdateRentPaymentRequest,
} from '../model/rent-payment-model';
import { RentPaymentWriteApiService } from './rent-payment-write-api.service';

@Injectable()
export class RentPaymentWriteService {
  private readonly api = inject(RentPaymentWriteApiService);

  private readonly isLoadingWritable = signal(false);
  private readonly errorWritable = signal<string | null>(null);

  readonly isLoadingSignal = computed(() => this.isLoadingWritable());
  readonly errorSignal = computed(() => this.errorWritable());

  save$(rentPayment: RentPaymentFormValue): Observable<RentPayment> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: CreateRentPaymentRequest = {
      ...rentPayment,
    };

    return this.api.create$(request).pipe(finalize(() => this.isLoadingWritable.set(false)));
  }

  update$(id: number, rentPayment: RentPaymentFormValue): Observable<RentPayment> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: UpdateRentPaymentRequest = {
      ...rentPayment,
    };

    return this.api.update$(id, request).pipe(finalize(() => this.isLoadingWritable.set(false)));
  }
}
