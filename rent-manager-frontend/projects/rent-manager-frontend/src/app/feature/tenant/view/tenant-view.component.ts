import { ChangeDetectionStrategy, Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { TenantFormComponent } from '../form/tenant-form.component';
import { TenantReadService } from '@core/tenant/tenant-read.service';

@Component({
  selector: 'rm-tenant-view',
  standalone: true,
  imports: [TenantFormComponent, MatButtonModule, MatIconModule, TranslatePipe],
  templateUrl: './tenant-view.component.html',
  styleUrl: './tenant-view.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantViewComponent implements OnInit {
  private readonly tenantReadService = inject(TenantReadService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly selectedTenantSignal = this.tenantReadService.selectedTenantSignal;
  readonly errorSignal = this.tenantReadService.errorSignal;

  readonly goToList$ = new Subject<void>();
  readonly goToEdit$ = new Subject<void>();

  private tenantId!: number;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/tenants']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToEdit$
      .pipe(
        filter(() => !!this.tenantId),
        switchMap(() => from(this.router.navigate(['/app/tenants/edit', this.tenantId]))),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  ngOnInit(): void {
    this.tenantId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.tenantId) {
      this.goToList$.next();
      return;
    }

    this.tenantReadService.get$(this.tenantId).subscribe();
  }
}
