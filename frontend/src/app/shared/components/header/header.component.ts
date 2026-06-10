import { Component } from '@angular/core';

@Component({
  selector: 'app-header',
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css',
})
export class Header {
  isLight = false;

  constructor() {
    this.isLight = localStorage.getItem('theme') === 'light';
    document.documentElement.classList.toggle('light', this.isLight);
  }

  toggleTheme() {
    this.isLight = !this.isLight;
    document.documentElement.classList.toggle('light', this.isLight);
    localStorage.setItem('theme', this.isLight ? 'light' : 'dark');
  }
}
