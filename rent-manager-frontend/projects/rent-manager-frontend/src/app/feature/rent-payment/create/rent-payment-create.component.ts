import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { TranslatePipe } from '@ngx-translate/core';
import { from, Subject, switchMap } from 'rxjs';

import { RentPaymentService } from '@core/rent-payment/rent-payment.service';
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
  private readonly rentPaymentService = inject(RentPaymentService);
  private readonly rentPaymentWriteService = inject(RentPaymentWriteService);
  private readonly doSave$ = new Subject<void>();
  private readonly doCancel$ = new Subject<void>();

  @ViewChild(RentPaymentFormComponent)
  rentPaymentForm!: RentPaymentFormComponent;

  readonly suggestionSignal = this.rentPaymentService.suggestionSignal;

  readonly isLoadingSignal = this.rentPaymentWriteService.isLoadingSignal;

  constructor() {
    const rentalContractId = Number(this.route.snapshot.queryParamMap.get('rentalContractId'));
    const period = this.route.snapshot.queryParamMap.get('period') ?? '';
    const paymentDate = new Date().toISOString().split('T')[0];
    this.rentPaymentService
      .getSuggestion$(rentalContractId, period, paymentDate)
      .pipe(takeUntilDestroyed())
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

  save(): void {
    if (!this.rentPaymentForm.isValid) {
      return;
    }
    this.doSave$.next();
  }

  cancel(): void {
    this.doCancel$.next();
  }
}
