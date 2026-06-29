import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { Owner } from '@core/owner/owner.model';

@Component({
  selector: 'rm-owner-list-table',
  standalone: true,
  imports: [MatTableModule, MatIconModule, MatButtonModule, TranslatePipe],
  templateUrl: './owner-list-table.component.html',
  styleUrl: './owner-list-table.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerListTableComponent {
  readonly owners = input<Owner[]>([]);
  readonly view = output<number>();
  readonly edit = output<number>();

  readonly displayedColumns = ['id', 'name', 'email', 'document_number', 'phone', 'actions'];
}
