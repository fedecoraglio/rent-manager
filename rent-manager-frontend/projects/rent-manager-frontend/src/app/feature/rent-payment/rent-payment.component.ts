import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'rm-rent-payment',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './rent-payment.component.html',
  styleUrl: './rent-payment.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentPaymentComponent {}
