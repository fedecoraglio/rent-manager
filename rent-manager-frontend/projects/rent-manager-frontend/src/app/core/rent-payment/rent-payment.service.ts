import { HttpErrorResponse } from '@angular/common/http';
import { computed, inject, Injectable, signal } from '@angular/core';
import { catchError, EMPTY, map, Observable, take, tap } from 'rxjs';

import { RentPaymentApiService } from './rent-payment-api.service';
import {
  RentalContractSummary,
  RentPayment,
  RentPaymentScheduleItem,
  RentPaymentSuggestion,
} from './rent-payment.model';

@Injectable({ providedIn: 'root' })
export class RentPaymentService {
  private readonly api = inject(RentPaymentApiService);

  private readonly rentPaymentsWritable = signal<RentPayment[]>([]);
  private readonly selectedRentPaymentWritable = signal<RentPayment | null>(null);
  private readonly scheduleWritable = signal<RentPaymentScheduleItem[]>([]);
  private readonly suggestionWritable = signal<RentPaymentSuggestion | null>(null);
  private readonly summaryWritable = signal<RentalContractSummary | null>(null);

  readonly rentPaymentsSignal = computed(() => this.rentPaymentsWritable());
  readonly selectedRentPaymentSignal = computed(() => this.selectedRentPaymentWritable());
  readonly scheduleSignal = computed(() => this.scheduleWritable());
  readonly suggestionSignal = computed(() => this.suggestionWritable());
  readonly summarySignal = computed(() => this.summaryWritable());

  list$(rentalContractId = 0, page = 1, limit = 10): Observable<RentPayment[]> {
    return this.api.list$(rentalContractId, page, limit).pipe(
      take(1),
      tap((apiResponse) => this.rentPaymentsWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  get$(id: number): Observable<RentPayment> {
    return this.api.getById$(id).pipe(
      take(1),
      tap((apiResponse) => this.selectedRentPaymentWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  getSchedule$(rentalContractId: number): Observable<RentPaymentScheduleItem[]> {
    return this.api.getSchedule$(rentalContractId).pipe(
      take(1),
      tap((apiResponse) => this.scheduleWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  getSuggestion$(
    rentalContractId: number,
    period: string,
    paymentDate: string,
  ): Observable<RentPaymentSuggestion> {
    return this.api.getSuggestion$(rentalContractId, period, paymentDate).pipe(
      take(1),
      tap((apiResponse) => this.suggestionWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  getSummary$(rentalContractId: number): Observable<RentalContractSummary> {
    return this.api.getSummary$(rentalContractId).pipe(
      take(1),
      tap((apiResponse) => this.summaryWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }
}
