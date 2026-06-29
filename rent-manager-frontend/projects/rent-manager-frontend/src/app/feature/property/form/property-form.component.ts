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
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { TranslatePipe } from '@ngx-translate/core';

import { LocationService } from '@core/location/location.service';
import { PropertyFormValue } from '../model/property.model';
import { PropertyFormField } from './field/property-form-field';
import { PropertyCatalogService } from '@core/property-catalog/property-catalog.service';
import { OwnerReadService } from '@core/owner/owner-read.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Property } from '@core/property/property.model';

@Component({
  selector: 'rm-property-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    TranslatePipe,
  ],
  templateUrl: './property-form.component.html',
  styleUrl: './property-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PropertyFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();
  private readonly locationService = inject(LocationService);
  private readonly propertyCatalogService = inject(PropertyCatalogService);
  private readonly ownerService = inject(OwnerReadService);

  readonly countriesSignal = this.locationService.countriesSignal;
  readonly propertyTypesSignal = this.propertyCatalogService.propertyTypesSignal;
  readonly propertyStatusesSignal = this.propertyCatalogService.propertyStatusesSignal;
  readonly ownersSignal = this.ownerService.ownersSignal;
  readonly statesSignal = this.locationService.statesSignal;
  readonly fields = PropertyFormField;

  readonly formGroup = new FormGroup({
    [PropertyFormField.OwnerId]: new FormControl(0, Validators.required),
    [PropertyFormField.TypeId]: new FormControl(0, Validators.required),
    [PropertyFormField.StatusId]: new FormControl(0, Validators.required),
    [PropertyFormField.CountryId]: new FormControl(1, Validators.required),
    [PropertyFormField.StateId]: new FormControl(0, Validators.required),
    [PropertyFormField.Code]: new FormControl('', Validators.required),
    [PropertyFormField.Title]: new FormControl('', Validators.required),
    [PropertyFormField.Description]: new FormControl(''),
    [PropertyFormField.Street]: new FormControl('', Validators.required),
    [PropertyFormField.StreetNumber]: new FormControl(''),
    [PropertyFormField.Floor]: new FormControl(''),
    [PropertyFormField.Apartment]: new FormControl(''),
    [PropertyFormField.City]: new FormControl('', Validators.required),
    [PropertyFormField.PostalCode]: new FormControl(''),
  });

  readonly property$ = new ReplaySubject<Property | null>(1);

  @Input() isReadonly = false;

  @Input() set property(property: Property | null) {
    this.property$.next(property);
  }

  constructor() {
    this.ownerService.list$(1, 100).pipe(takeUntilDestroyed()).subscribe();
  }

  get ownerIdField() {
    return this.formGroup.get(PropertyFormField.OwnerId);
  }

  get typeIdField() {
    return this.formGroup.get(PropertyFormField.TypeId);
  }

  get statusIdField() {
    return this.formGroup.get(PropertyFormField.StatusId);
  }

  get countryIdField() {
    return this.formGroup.get(PropertyFormField.CountryId);
  }

  get stateIdField() {
    return this.formGroup.get(PropertyFormField.StateId);
  }

  get codeField() {
    return this.formGroup.get(PropertyFormField.Code);
  }

  get titleField() {
    return this.formGroup.get(PropertyFormField.Title);
  }

  get streetField() {
    return this.formGroup.get(PropertyFormField.Street);
  }

  get cityField() {
    return this.formGroup.get(PropertyFormField.City);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): PropertyFormValue {
    return {
      ...(this.formGroup.getRawValue() as PropertyFormValue),
    };
  }

  ngAfterViewInit(): void {
    this.property$
      .pipe(filter(Boolean), takeUntil(this.destroy$))
      .subscribe((property: Property) => {
        this.formGroup.patchValue({
          owner_id: property.owner_id,
          type_id: property.type_id,
          status_id: property.status_id,
          country_id: property.country_id,
          state_id: property.state_id,
          code: property.code,
          title: property.title,
          description: property.description,
          street: property.street,
          street_number: property.street_number,
          floor: property.floor,
          apartment: property.apartment,
          city: property.city,
          postal_code: property.postal_code,
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
