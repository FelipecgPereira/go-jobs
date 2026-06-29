import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiService } from '../services/api.service';
import { Customer } from '../models';

@Component({
  selector: 'app-customers',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './customers.component.html',
  styleUrl: './customers.component.css',
})
export class CustomersComponent implements OnInit {
  customers: Customer[] = [];
  form: FormGroup;
  editingId: number | null = null;

  constructor(private api: ApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: [''],
    });
  }

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    this.api.getCustomers().subscribe({ next: (customers) => (this.customers = customers) });
  }

  submit(): void {
    if (this.form.invalid) return;

    const payload = this.form.value as Customer;

    if (this.editingId) {
      this.api.updateCustomer(this.editingId, payload).subscribe({ next: () => { this.reset(); this.load(); } });
      return;
    }

    this.api.createCustomer(payload).subscribe({ next: () => { this.reset(); this.load(); } });
  }

  edit(customer: Customer): void {
    this.editingId = customer.id ?? null;
    this.form.patchValue(customer);
  }

  reset(): void {
    this.form.reset();
    this.editingId = null;
  }
}
