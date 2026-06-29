import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { PropertyFormComponent } from '../form/property-form.component';
import { PropertyWriteService } from '../service/property-write.service';
import { PropertyReadService } from '@core/property/property-read.service';

@Component({
  selector: 'rm-property-create',
  standalone: true,
  imports: [PropertyFormComponent, MatButtonModule, TranslatePipe],
  templateUrl: './property-create.component.html',
  styleUrl: './property-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PropertyCreateComponent {
  private readonly propertyWriteService = inject(PropertyWriteService);
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly router = inject(Router);

  readonly create$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly isLoadingSignal = this.propertyReadService.isLoadingSignal;

  @ViewChild('propertyForm') propertyForm!: PropertyFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/properties']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.create$
      .pipe(
        filter(() => this.propertyForm.isValid),
        switchMap(() => this.propertyWriteService.save$(this.propertyForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
