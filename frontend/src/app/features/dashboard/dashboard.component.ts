import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-dashboard',
  template: `
    <mat-card>
      <mat-card-title>Projects</mat-card-title>
      <mat-list>
        <mat-list-item *ngFor="let project of projects">{{ project.name }}</mat-list-item>
      </mat-list>
    </mat-card>
  `
})
export class DashboardComponent implements OnInit {
  projects: Array<{ id: string; name: string }> = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.http
      .get<Array<{ id: string; name: string }>>(`${environment.apiBaseUrl}/projects`)
      .subscribe((response) => (this.projects = response));
  }
}
