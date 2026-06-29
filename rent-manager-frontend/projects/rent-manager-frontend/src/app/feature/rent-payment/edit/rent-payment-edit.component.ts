import { ChangeDetectionStrategy, Component, inject, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { RentPaymentService } from '@core/rent-payment/rent-payment.service';
import { RentPaymentFormComponent } from '../form/rent-payment-form.component';
import { RentPaymentWriteService } from '../service/rent-payment-write.service';

@Component({
  selector: 'rm-rent-payment-edit',
  standalone: true,
  imports: [MatButtonModule, TranslatePipe, RentPaymentFormComponent],
  templateUrl: './rent-payment-edit.component.html',
  styleUrl: './rent-payment-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentPaymentEditComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  private readonly rentPaymentService = inject(RentPaymentService);
  private readonly rentPaymentWriteService = inject(RentPaymentWriteService);

  @ViewChild(RentPaymentFormComponent)
  rentPaymentForm!: RentPaymentFormComponent;

  readonly rentPaymentSignal = this.rentPaymentService.selectedRentPaymentSignal;
  readonly isLoadingSignal = this.rentPaymentWriteService.isLoadingSignal;

  private rentPaymentId!: number;

  ngOnInit(): void {
    this.rentPaymentId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.rentPaymentId) {
      this.cancel();
      return;
    }

    this.rentPaymentService.get$(this.rentPaymentId).subscribe();
  }

  save(): void {
    if (!this.rentPaymentForm.isValid) {
      return;
    }

    this.rentPaymentWriteService.update$(this.rentPaymentId, this.rentPaymentForm.value).subscribe({
      next: () => this.cancel(),
    });
  }

  cancel(): void {
    void this.router.navigate(['/app/rental-contracts']);
  }
}
