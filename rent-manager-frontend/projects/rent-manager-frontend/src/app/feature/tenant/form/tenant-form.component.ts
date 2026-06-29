import {
  AfterViewInit,
  ChangeDetectionStrategy,
  Component,
  inject,
  Input,
  OnDestroy,
} from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { filter, ReplaySubject, Subject, takeUntil } from 'rxjs';
import { TranslatePipe } from '@ngx-translate/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';

import { LocationService } from '@core/location/location.service';
import { TenantFormValue } from '../model/tenant.model';
import { TenantFormField } from './field/tenant-form-field';
import { Tenant } from '@core/tenant/tenant-core.model';

@Component({
  selector: 'rm-tenant-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    TranslatePipe,
  ],
  templateUrl: './tenant-form.component.html',
  styleUrl: './tenant-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();
  private readonly locationService = inject(LocationService);

  readonly countriesSignal = this.locationService.countriesSignal;
  readonly statesSignal = this.locationService.statesSignal;
  readonly fields = TenantFormField;

  readonly formGroup = new FormGroup({
    [TenantFormField.CountryId]: new FormControl(1, Validators.required),
    [TenantFormField.StateId]: new FormControl(0, Validators.required),
    [TenantFormField.Name]: new FormControl('', Validators.required),
    [TenantFormField.Email]: new FormControl('', [Validators.email]),
    [TenantFormField.DocumentNumber]: new FormControl('', Validators.required),
    [TenantFormField.Phone]: new FormControl(''),
    [TenantFormField.City]: new FormControl(''),
    [TenantFormField.Street]: new FormControl(''),
    [TenantFormField.StreetNumber]: new FormControl(''),
    [TenantFormField.Floor]: new FormControl(''),
    [TenantFormField.Apartment]: new FormControl(''),
    [TenantFormField.PostalCode]: new FormControl(''),
  });

  readonly tenant$ = new ReplaySubject<Tenant | null>(1);

  @Input() isReadonly = false;

  @Input() set tenant(tenant: Tenant | null) {
    this.tenant$.next(tenant);
  }

  get nameField() {
    return this.formGroup.get(TenantFormField.Name);
  }

  get emailField() {
    return this.formGroup.get(TenantFormField.Email);
  }

  get documentNumberField() {
    return this.formGroup.get(TenantFormField.DocumentNumber);
  }

  get countryIdField() {
    return this.formGroup.get(TenantFormField.CountryId);
  }

  get stateIdField() {
    return this.formGroup.get(TenantFormField.StateId);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): TenantFormValue {
    return {
      ...(this.formGroup.getRawValue() as TenantFormValue),
    };
  }

  ngAfterViewInit(): void {
    this.tenant$.pipe(filter(Boolean), takeUntil(this.destroy$)).subscribe((tenant: Tenant) => {
      this.formGroup.patchValue({
        country_id: tenant.country_id,
        state_id: tenant.state_id,
        name: tenant.name,
        email: tenant.email,
        document_number: tenant.document_number,
        phone: tenant.phone,
        city: tenant.city,
        street: tenant.street,
        street_number: tenant.street_number,
        floor: tenant.floor,
        apartment: tenant.apartment,
        postal_code: tenant.postal_code,
      });

      if (this.isReadonly) {
        this.formGroup.disable();
      }

      this.formGroup.updateValueAndValidity();
    });
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
