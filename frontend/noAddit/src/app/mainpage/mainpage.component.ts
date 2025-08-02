import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-mainpage',
  standalone: false,
  templateUrl: './mainpage.component.html',
  styleUrl: './mainpage.component.scss'
})
export class MainpageComponent implements OnInit {
  userConfig: any;

  constructor(private router: Router) {}

  ngOnInit() {
    const savedConfig = localStorage.getItem('userConfig');
    if (!savedConfig) {
      console.log('No saved localstorage values, rerouting to config...')
      this.router.navigate(['/config']);
      return;
    }
    this.userConfig = JSON.parse(savedConfig);
  }
}