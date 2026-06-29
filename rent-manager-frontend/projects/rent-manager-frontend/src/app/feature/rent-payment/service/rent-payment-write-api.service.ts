import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { RentPayment } from '@core/rent-payment/rent-payment.model';

import { CreateRentPaymentRequest, UpdateRentPaymentRequest } from '../model/rent-payment-model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable()
export class RentPaymentWriteApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/rent-payments');
  }

  create$(request: CreateRentPaymentRequest): Observable<RentPayment> {
    return this.http
      .post<ApiResponse<RentPayment>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  update$(id: number, request: UpdateRentPaymentRequest): Observable<RentPayment> {
    return this.http
      .put<ApiResponse<RentPayment>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }
}
