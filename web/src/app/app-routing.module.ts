import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ListComponent} from './components/resources/list/list.component';


const routes: Routes = [
  {path: '**', redirectTo: 'resources/list'},
  {path: 'resources/list', component: ListComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
