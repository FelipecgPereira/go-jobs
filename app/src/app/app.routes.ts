import { Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { CustomersComponent } from './customers/customers.component';
import { ProjectsComponent } from './projects/projects.component';
import { B2bComponent } from './b2b/b2b.component';
import { authGuard } from './auth.guard';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', component: AuthComponent },
  {
    path: 'dashboard',
    component: DashboardComponent,
    canActivate: [authGuard],
  },
  {
    path: 'customers',
    component: CustomersComponent,
    canActivate: [authGuard],
  },
  {
    path: 'projects',
    component: ProjectsComponent,
    canActivate: [authGuard],
  },
  {
    path: 'b2b',
    component: B2bComponent,
    canActivate: [authGuard],
  },
];
