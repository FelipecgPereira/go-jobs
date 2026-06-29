import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiService } from '../services/api.service';
import { Project } from '../models';

@Component({
  selector: 'app-projects',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './projects.component.html',
  styleUrl: './projects.component.css',
})
export class ProjectsComponent implements OnInit {
  projects: Project[] = [];
  form: FormGroup;
  editingId: number | null = null;

  constructor(private api: ApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      name: ['', Validators.required],
      price: [0, [Validators.required, Validators.min(0)]],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
    });
  }

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    this.api.getProjects().subscribe({ next: (projects) => (this.projects = projects) });
  }

  submit(): void {
    if (this.form.invalid) return;

    const payload = this.form.value as Project;

    if (this.editingId) {
      this.api.updateProject(this.editingId, payload).subscribe({ next: () => { this.reset(); this.load(); } });
      return;
    }

    this.api.createProject(payload).subscribe({ next: () => { this.reset(); this.load(); } });
  }

  edit(project: Project): void {
    this.editingId = project.id ?? null;
    this.form.patchValue({ ...project, startDate: project.startDate?.slice(0, 10), endDate: project.endDate?.slice(0, 10) });
  }

  reset(): void {
    this.form.reset();
    this.editingId = null;
  }
}
