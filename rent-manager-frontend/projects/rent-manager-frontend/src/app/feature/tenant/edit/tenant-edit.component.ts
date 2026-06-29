import { ChangeDetectionStrategy, Component, inject, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { TenantFormComponent } from '../form/tenant-form.component';
import { TenantWriteService } from '../service/tenant-write.service';
import { TenantReadService } from '@core/tenant/tenant-read.service';

@Component({
  selector: 'rm-tenant-edit',
  standalone: true,
  imports: [TenantFormComponent, MatButtonModule, MatIconModule, TranslatePipe],
  templateUrl: './tenant-edit.component.html',
  styleUrl: './tenant-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantEditComponent implements OnInit {
  private readonly tenantWriteService = inject(TenantWriteService);
  private readonly tenantReadService = inject(TenantReadService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly selectedTenantSignal = this.tenantReadService.selectedTenantSignal;
  readonly isLoadingSignal = this.tenantWriteService.isLoadingSignal;
  readonly errorSignal = this.tenantWriteService.errorSignal;

  readonly goToList$ = new Subject<void>();
  readonly update$ = new Subject<void>();

  private tenantId!: number;

  @ViewChild('tenantForm') tenantForm!: TenantFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/tenants']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.update$
      .pipe(
        filter(() => this.tenantForm.isValid),
        switchMap(() => this.tenantWriteService.update$(this.tenantId, this.tenantForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
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
