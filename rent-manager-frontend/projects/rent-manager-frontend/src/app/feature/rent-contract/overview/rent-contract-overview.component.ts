import { ChangeDetectionStrategy, Component, OnDestroy, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subject, switchMap, takeUntil, tap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { PropertyReadService } from '@core/property/property-read.service';
import { RentContractReadService } from '@core/rent-contract/rent-contract-read.service';
import { RentPaymentService } from '@core/rent-payment/rent-payment.service';
import { RentPaymentScheduleItem } from '@core/rent-payment/rent-payment.model';
import { RentPaymentScheduleComponent } from '@ui/rent-payment-schedule/rent-payment-schedule.component';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { TranslatePipe } from '@ngx-translate/core';

@Component({
  selector: 'rm-rent-contract-overview',
  standalone: true,
  imports: [MatButtonModule, MatIconModule, RentPaymentScheduleComponent, TranslatePipe],
  templateUrl: './rent-contract-overview.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractOverviewComponent implements OnInit, OnDestroy {
  private readonly destroy$ = new Subject<void>();
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly rentContractReadService = inject(RentContractReadService);
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly rentPaymentService = inject(RentPaymentService);
  private readonly goToDashboard$ = new Subject<void>();

  private rentalContractId!: number;

  readonly rentContractSignal = this.rentContractReadService.selectedRentContractSignal;
  readonly propertySignal = this.propertyReadService.selectedPropertySignal;
  readonly paymentScheduleSignal = this.rentPaymentService.scheduleSignal;

  constructor() {
    this.goToDashboard$
      .pipe(
        switchMap(() => this.router.navigate(['/app/dashboard'])),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  ngOnInit(): void {
    this.rentalContractId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.rentalContractId) {
      void this.router.navigate(['/app/dashboard']);
      return;
    }

    this.rentContractReadService
      .get$(this.rentalContractId)
      .pipe(
        tap(() => {
          this.rentPaymentService
            .getSchedule$(this.rentalContractId)
            .pipe(takeUntil(this.destroy$))
            .subscribe();
        }),
        switchMap((contract) => this.propertyReadService.get$(contract.property_id)),
        takeUntil(this.destroy$),
      )
      .subscribe();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  goBack(): void {
    this.goToDashboard$.next();
  }

  registerPayment(item: RentPaymentScheduleItem): void {
    void this.router.navigate(['/app/rent-payments/create'], {
      queryParams: {
        rentalContractId: item.rental_contract_id,
        period: item.period,
      },
    });
  }
}
