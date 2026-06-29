import {
  ChangeDetectionStrategy,
  Component,
  computed,
  inject,
  OnDestroy,
  OnInit,
  signal,
} from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { BehaviorSubject, filter, from, map, Subject, switchMap, takeUntil, tap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { RentContractReadService } from '@core/rent-contract/rent-contract-read.service';
import { RentPaymentService } from '@core/rent-payment/rent-payment.service';
import { RentPaymentScheduleItem } from '@core/rent-payment/rent-payment.model';
import { RentPaymentScheduleComponent } from '@ui/rent-payment-schedule/rent-payment-schedule.component';
import { PropertyReadService } from '@core/property/property-read.service';
import { ContractCatalogService } from '@core/contract-catalog/contract-catalog.service';
import { RentContract } from '@core/rent-contract/rent-contract.model';

@Component({
  selector: 'rm-rent-contract-view',
  standalone: true,
  imports: [MatButtonModule, MatIconModule, TranslatePipe, RentPaymentScheduleComponent],
  templateUrl: './rent-contract-view.component.html',
  styleUrl: './rent-contract-view.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractViewComponent implements OnInit, OnDestroy {
  private readonly destroy$ = new Subject<void>();
  private readonly rentContractReadService = inject(RentContractReadService);
  private readonly rentPaymentService = inject(RentPaymentService);
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly contractCatalogService = inject(ContractCatalogService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly dailyInterestWritable = signal(0);
  private readonly interestTypeNameWritable = signal('');
  private readonly adjustmentTypeNameWritable = signal('');
  private readonly contractStatusWritable = signal('');
  private readonly rentContract$ = new BehaviorSubject<RentContract | null>(null);

  private rentContractId!: number;

  readonly dailyInterestSignal = computed(() => this.dailyInterestWritable());
  readonly adjustmentTypeNameSignal = computed(() => this.adjustmentTypeNameWritable());
  readonly interestTypeNameSignal = computed(() => this.interestTypeNameWritable());
  readonly contractStatusSignal = computed(() => this.contractStatusWritable());
  readonly goToEdit$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly propertySignal = this.propertyReadService.selectedPropertySignal;
  readonly rentContractSignal = this.rentContractReadService.selectedRentContractSignal;
  readonly paymentScheduleSignal = this.rentPaymentService.scheduleSignal;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/rent-contracts']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToEdit$
      .pipe(
        filter(() => !!this.rentContractId),
        switchMap(() =>
          from(this.router.navigate(['/app/rent-contracts/edit', this.rentContractId])),
        ),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  ngOnInit(): void {
    this.rentContractId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.rentContractId) {
      this.goToList$.next();
      return;
    }

    this.rentContract$
      .pipe(
        filter(Boolean),
        tap((rentContract) => {
          this.dailyInterestWritable.set(rentContract.daily_interest_percentage);
          const interestType = this.contractCatalogService
            .interestCalculationTypesSignal()
            .find((it) => it.id === rentContract?.interest_calculation_type_id);
          if (interestType) {
            this.interestTypeNameWritable.set(interestType.name);
          }
          const status = this.contractCatalogService
            .contractStatusesSignal()
            .find((st) => st.id === rentContract.status_id);
          if (status) {
            this.contractStatusWritable.set(status.name);
          }
          const adjustment = this.contractCatalogService
            .rentAdjustmentTypesSignal()
            .find((at) => at.id === rentContract.adjustment_type_id);
          if (adjustment) {
            this.adjustmentTypeNameWritable.set(adjustment.name);
          }
        }),
        takeUntil(this.destroy$),
      )
      .subscribe();

    this.rentContractReadService
      .get$(this.rentContractId)
      .pipe(takeUntil(this.destroy$))
      .subscribe();
    this.rentContractReadService
      .get$(this.rentContractId)
      .pipe(
        switchMap((contract) =>
          this.rentPaymentService.getSchedule$(this.rentContractId).pipe(
            map(() => contract),
            takeUntil(this.destroy$),
          ),
        ),
        switchMap((contract) =>
          this.propertyReadService.get$(contract.property_id).pipe(
            map(() => contract),
            takeUntil(this.destroy$),
          ),
        ),
        takeUntil(this.destroy$),
      )
      .subscribe((rentContract) => this.rentContract$.next(rentContract));
  }

  registerPayment(item: RentPaymentScheduleItem): void {
    void this.router.navigate(['/app/rent-payments/create'], {
      queryParams: {
        rentalContractId: item.rental_contract_id,
        period: item.period,
      },
    });
  }

  editPayment(): void {
    this.goToEdit$.next();
  }
}
