import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';
import { forkJoin } from 'rxjs';
import { ApiService } from '../services/api.service';
import { Customer, Project } from '../models';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
})
export class DashboardComponent implements OnInit {
  customers: Customer[] = [];
  projects: Project[] = [];
  loading = false;

  constructor(private api: ApiService, private router: Router) {}

  ngOnInit(): void {
    this.loadData();
  }

  loadData(): void {
    this.loading = true;

    forkJoin({
      customers: this.api.getCustomers(),
      projects: this.api.getProjects(),
    }).subscribe({
      next: ({ customers, projects }) => {
        this.customers = customers;
        this.projects = projects;
      },
      error: () => {
        this.customers = [];
        this.projects = [];
      },
      complete: () => {
        this.loading = false;
      },
    });
  }

  logout(): void {
    this.api.logout();
    this.router.navigate(['/login']);
  }
}
