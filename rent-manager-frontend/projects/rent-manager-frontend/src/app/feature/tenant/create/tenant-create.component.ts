import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { TenantFormComponent } from '../form/tenant-form.component';
import { TenantWriteService } from '../service/tenant-write.service';

@Component({
  selector: 'rm-tenant-create',
  standalone: true,
  imports: [TenantFormComponent, MatButtonModule, TranslatePipe],
  templateUrl: './tenant-create.component.html',
  styleUrl: './tenant-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantCreateComponent {
  private readonly tenantWriteService = inject(TenantWriteService);
  private readonly router = inject(Router);

  readonly create$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly isLoadingSignal = this.tenantWriteService.isLoadingSignal;

  @ViewChild('tenantForm') tenantForm!: TenantFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/tenants']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.create$
      .pipe(
        filter(() => this.tenantForm.isValid),
        switchMap(() => this.tenantWriteService.save$(this.tenantForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
