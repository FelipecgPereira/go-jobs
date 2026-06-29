import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiService } from '../services/api.service';

@Component({
  selector: 'app-b2b',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './b2b.component.html',
  styleUrl: './b2b.component.css',
})
export class B2bComponent implements OnInit {
  form: FormGroup;
  summary: number | null = null;
  constructor(private api: ApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      customerId: [0, [Validators.required, Validators.min(1)]],
      projectId: [0],
      status: ['pending', Validators.required],
      start: ['', Validators.required],
      end: ['', Validators.required],
    });
  }

  ngOnInit(): void {}

  submit(): void {
    if (this.form.invalid) return;
    const value = this.form.value;
    this.api.createB2b({ customerId: value.customerId, projectId: value.projectId || undefined, status: value.status }).subscribe({
      next: () => alert('B2B criado com sucesso!'),
      error: () => alert('Falha ao criar B2B.'),
    });
  }

  loadSummary(): void {
    const value = this.form.value;
    this.api.getB2bSummary(value.status, value.start, value.end).subscribe({
      next: (response) => (this.summary = response.total),
      error: () => (this.summary = null),
    });
  }
}
