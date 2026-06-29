import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';

import {
  RentalContractSummary,
  RentPayment,
  RentPaymentScheduleItem,
  RentPaymentSuggestion,
} from './rent-payment.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class RentPaymentApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get rentPaymentsUrl(): string {
    return this.apiUrl.build('/rent-payments');
  }

  private get rentalContractsUrl(): string {
    return this.apiUrl.build('/rental-contracts');
  }

  list$(
    rentalContractId: number,
    page: number,
    limit: number,
  ): Observable<ApiResponse<RentPayment[]>> {
    let params = new HttpParams().set('page', page).set('limit', limit);

    if (rentalContractId > 0) {
      params = params.set('rental_contract_id', rentalContractId);
    }

    return this.http
      .get<ApiResponse<RentPayment[]>>(this.rentPaymentsUrl, { params })
      .pipe(map((response) => response));
  }

  getById$(id: number): Observable<ApiResponse<RentPayment>> {
    return this.http
      .get<ApiResponse<RentPayment>>(`${this.rentPaymentsUrl}/${id}`)
      .pipe(map((response) => response));
  }

  getSchedule$(rentalContractId: number): Observable<ApiResponse<RentPaymentScheduleItem[]>> {
    return this.http
      .get<
        ApiResponse<RentPaymentScheduleItem[]>
      >(`${this.rentalContractsUrl}/${rentalContractId}/payment-schedules`)
      .pipe(map((response) => response));
  }

  getSuggestion$(
    rentalContractId: number,
    period: string,
    paymentDate: string,
  ): Observable<ApiResponse<RentPaymentSuggestion>> {
    const params = new HttpParams().set('period', period).set('payment_date', paymentDate);

    return this.http
      .get<
        ApiResponse<RentPaymentSuggestion>
      >(`${this.rentalContractsUrl}/${rentalContractId}/payment-suggestions`, { params })
      .pipe(map((response) => response));
  }

  getSummary$(rentalContractId: number): Observable<ApiResponse<RentalContractSummary>> {
    return this.http
      .get<
        ApiResponse<RentalContractSummary>
      >(`${this.rentalContractsUrl}/${rentalContractId}/summary`)
      .pipe(map((response) => response));
  }
}
