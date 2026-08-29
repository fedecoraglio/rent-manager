import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { TranslatePipe } from '@ngx-translate/core';

import { RentContract } from '@core/rent-contract/rent-contract.model';
import {CurrencyPipe} from "@angular/common";

@Component({
  selector: 'rm-rent-contract-list-table',
  standalone: true,
  imports: [MatTableModule, MatIconModule, MatButtonModule, TranslatePipe, CurrencyPipe],
  templateUrl: './rent-contract-list-table.component.html',
  styleUrl: './rent-contract-list-table.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractListTableComponent {
  readonly rentContracts = input<RentContract[]>([]);

  readonly view = output<number>();
  readonly edit = output<number>();

  readonly displayedColumns = [
    'property',
    'tenant',
    'start_date',
    'end_date',
    'monthly_amount',
    'actions',
  ];
}
