import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { TranslatePipe } from '@ngx-translate/core';
import { MatTooltipModule } from '@angular/material/tooltip';

import { RentPaymentScheduleItem } from '@core/rent-payment/rent-payment.model';
import {CurrencyPipe} from "@angular/common";

@Component({
  selector: 'rm-rent-payment-schedule',
  standalone: true,
  imports: [MatTableModule, MatButtonModule, MatIconModule, MatTooltipModule, TranslatePipe, CurrencyPipe],
  templateUrl: './rent-payment-schedule.component.html',
  styleUrl: './rent-payment-schedule.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentPaymentScheduleComponent {
  readonly items = input<RentPaymentScheduleItem[]>([]);

  readonly registerPayment = output<RentPaymentScheduleItem>();
  readonly editPayment = output<RentPaymentScheduleItem>();

  readonly displayedColumns = [
    'period',
    'due_date',
    'suggested_total_amount',
    'paid_amount',
    'status',
    'actions',
  ];
}
