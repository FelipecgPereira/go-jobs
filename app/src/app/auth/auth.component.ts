import { Component } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../services/api.service';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './auth.component.html',
  styleUrl: './auth.component.css',
})
export class AuthComponent {
  authForm: FormGroup;
  isLoginMode = true;
  message = '';

  constructor(
    private fb: FormBuilder,
    private api: ApiService,
    private router: Router,
  ) {
    this.authForm = this.fb.group({
      name: [''],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required]],
    });
  }

  toggleMode(): void {
    this.isLoginMode = !this.isLoginMode;
    this.message = '';
    this.authForm.reset({ name: '', email: '', password: '' });
  }

  submit(): void {
    if (this.authForm.invalid) {
      this.message = 'Preencha todos os campos corretamente.';
      return;
    }

    const payload = this.authForm.value;

    if (this.isLoginMode) {
      this.api.login({ email: payload.email, password: payload.password }).subscribe({
        next: (response) => {
          localStorage.setItem('gojobs_token', response.token ?? '');
          this.router.navigate(['/dashboard']);
        },
        error: () => {
          this.message = 'Falha no login. Verifique seu e-mail e senha.';
        },
      });
      return;
    }

    this.api.register({ name: payload.name, email: payload.email, password: payload.password }).subscribe({
      next: () => {
        this.message = 'Cadastro realizado com sucesso. Faça login para continuar.';
        this.isLoginMode = true;
        this.authForm.reset({ name: '', email: '', password: '' });
      },
      error: () => {
        this.message = 'Não foi possível criar o usuário.';
      },
    });
  }
}
