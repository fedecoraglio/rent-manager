import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { Router, RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { forkJoin, Subject, switchMap } from 'rxjs';
import { TranslatePipe, TranslateService } from '@ngx-translate/core';

import { AuthService } from '@core/auth/auth.service';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { AppNotificationComponent } from '@ui/notification/app-notification.component';
import { LoadingService } from '@core/loading/loading.service';
import { RoleService } from '@core/role/role.service';
import { LocationService } from '@core/location/location.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { PropertyCatalogService } from '@core/property-catalog/property-catalog.service';
import { ContractCatalogService } from '@core/contract-catalog/contract-catalog.service';

interface MenuItem {
  label: string;
  icon: string;
  route: string;
}

@Component({
  selector: 'rm-main-layout',
  standalone: true,
  imports: [
    RouterOutlet,
    RouterLink,
    RouterLinkActive,
    MatButtonModule,
    MatIconModule,
    MatProgressBarModule,
    AppNotificationComponent,
    TranslatePipe,
  ],
  templateUrl: './main-layout.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MainLayoutComponent {
  private readonly authService = inject(AuthService);
  private readonly locationService = inject(LocationService);
  private readonly roleService = inject(RoleService);
  private readonly propertyCatalogService = inject(PropertyCatalogService);
  private readonly contractCatalogService = inject(ContractCatalogService);
  private readonly translateService = inject(TranslateService);
  private readonly router = inject(Router);
  private readonly loadInitData$ = new Subject<void>();

  readonly loadingService = inject(LoadingService);

  readonly menuItems: MenuItem[] = [
    {
      label: this.translateService.instant('menu.dashboard'),
      icon: 'dashboard',
      route: '/app/dashboard',
    },
    {
      label: this.translateService.instant('menu.users'),
      icon: 'group',
      route: '/app/users',
    },
    {
      label: this.translateService.instant('menu.owners'),
      icon: 'people',
      route: '/app/owners',
    },
    {
      label: this.translateService.instant('menu.tenants'),
      icon: 'person_pin',
      route: '/app/tenants',
    },
    {
      label: this.translateService.instant('menu.inflationIndexes'),
      icon: 'trending_up',
      route: '/app/inflation-indexes',
    },
    {
      label: this.translateService.instant('menu.properties'),
      icon: 'home_work',
      route: '/app/properties',
    },
    {
      label: this.translateService.instant('menu.contracts'),
      icon: 'home_work',
      route: '/app/rent-contracts',
    },
  ];

  constructor() {
    this.loadInitData$
      .pipe(
        switchMap(() =>
          forkJoin([
            this.roleService.listRoles$(),
            this.locationService.listCountries$(),
            this.locationService.listStates$(1),
            this.propertyCatalogService.listPropertyTypes$(),
            this.propertyCatalogService.listPropertyStatuses$(),
            this.contractCatalogService.listContractStatuses$(),
            this.contractCatalogService.listInterestCalculationTypes$(),
            this.contractCatalogService.listRentAdjustmentTypes$(),
          ]),
        ),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.loadInitData$.next();
  }

  logout(): void {
    this.authService.logout();
    void this.router.navigate(['/login']);
  }
}
