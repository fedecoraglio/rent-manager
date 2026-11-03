import { ChangeDetectionStrategy, Component, inject, signal, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { TranslatePipe } from '@ngx-translate/core';
import { catchError, EMPTY, finalize, from, startWith, Subject, switchMap, tap } from 'rxjs';

import { RentPaymentApiService } from '@core/rent-payment/rent-payment-api.service';
import { RentPaymentSuggestion } from '@core/rent-payment/rent-payment.model';
import { RentPaymentFormComponent } from '../form/rent-payment-form.component';
import { RentPaymentWriteService } from '../service/rent-payment-write.service';

@Component({
  selector: 'rm-rent-payment-create',
  standalone: true,
  imports: [MatButtonModule, TranslatePipe, RentPaymentFormComponent, MatIconModule],
  templateUrl: './rent-payment-create.component.html',
  styleUrl: './rent-payment-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentPaymentCreateComponent {
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly rentPaymentApi = inject(RentPaymentApiService);
  private readonly paymentDates$ = new Subject<string>();
  private readonly rentPaymentWriteService = inject(RentPaymentWriteService);
  private readonly doSave$ = new Subject<void>();
  private readonly doCancel$ = new Subject<void>();

  @ViewChild(RentPaymentFormComponent)
  rentPaymentForm!: RentPaymentFormComponent;

  readonly suggestionSignal = signal<RentPaymentSuggestion | null>(null);
  readonly isRefreshingSignal = signal(false);
  readonly suggestionErrorSignal = signal(false);
  readonly suggestionReadySignal = signal(false);

  readonly isLoadingSignal = this.rentPaymentWriteService.isLoadingSignal;

  constructor() {
    const rentalContractId = Number(this.route.snapshot.queryParamMap.get('rentalContractId'));
    const period = this.route.snapshot.queryParamMap.get('period') ?? '';
    const today = new Date();
    const paymentDate = [
      today.getFullYear(),
      String(today.getMonth() + 1).padStart(2, '0'),
      String(today.getDate()).padStart(2, '0'),
    ].join('-');
    this.paymentDates$
      .pipe(
        startWith(paymentDate),
        switchMap((date) => {
          this.suggestionReadySignal.set(false);
          this.suggestionErrorSignal.set(false);
          if (!date) return EMPTY;
          this.isRefreshingSignal.set(true);
          return this.rentPaymentApi.getSuggestion$(rentalContractId, period, date).pipe(
            tap((response) => {
              this.suggestionSignal.set(response.data);
              this.suggestionReadySignal.set(true);
            }),
            catchError(() => {
              this.suggestionErrorSignal.set(true);
              return EMPTY;
            }),
            finalize(() => this.isRefreshingSignal.set(false)),
          );
        }),
        takeUntilDestroyed(),
      )
      .subscribe();
    this.doSave$
      .pipe(
        switchMap(() => this.rentPaymentWriteService.save$(this.rentPaymentForm.value)),
        switchMap(() => from(this.router.navigate(['/app/rent-contracts/view', rentalContractId]))),
        takeUntilDestroyed(),
      )
      .subscribe();
    this.doCancel$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/rent-contracts/view', rentalContractId]))),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  refreshSuggestion(paymentDate: string): void {
    this.paymentDates$.next(paymentDate);
  }

  save(): void {
    if (
      !this.suggestionReadySignal() ||
      this.isRefreshingSignal() ||
      !this.rentPaymentForm.isValid
    ) {
      return;
    }
    this.doSave$.next();
  }

  cancel(): void {
    this.doCancel$.next();
  }
}
