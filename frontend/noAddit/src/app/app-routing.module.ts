import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ConfigComponent } from './config/config.component';
import { MainpageComponent } from './mainpage/mainpage.component';

const routes: Routes = [
  {path:'config', component: ConfigComponent},
  {path:'home', component: MainpageComponent},
  {path:'', redirectTo: '/home', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
