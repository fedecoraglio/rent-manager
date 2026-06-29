import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { InflationIndex } from '@core/inflation-index/inflation-index.model';

@Component({
  selector: 'rm-inflation-index-list-table',
  standalone: true,
  imports: [MatTableModule, MatIconModule, MatButtonModule, TranslatePipe],
  templateUrl: './inflation-index-list-table.component.html',
  styleUrl: './inflation-index-list-table.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InflationIndexListTableComponent {
  readonly inflationIndexes = input<InflationIndex[]>([]);
  readonly view = output<number>();
  readonly edit = output<number>();

  readonly displayedColumns = ['period', 'percentage', 'source', 'notes', 'actions'];
}
