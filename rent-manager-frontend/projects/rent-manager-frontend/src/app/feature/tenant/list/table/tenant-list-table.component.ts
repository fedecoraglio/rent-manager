import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { TranslatePipe } from '@ngx-translate/core';
import { Tenant } from '@core/tenant/tenant-core.model';

@Component({
  selector: 'rm-tenant-list-table',
  standalone: true,
  imports: [MatTableModule, MatIconModule, MatButtonModule, TranslatePipe],
  templateUrl: './tenant-list-table.component.html',
  styleUrl: './tenant-list-table.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantListTableComponent {
  readonly tenants = input<Tenant[]>([]);

  readonly view = output<number>();
  readonly edit = output<number>();

  readonly displayedColumns = [
    'id',
    'name',
    'document_number',
    'email',
    'phone',
    'city',
    'actions',
  ];
}
