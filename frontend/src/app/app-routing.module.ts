import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DataListComponent } from './data-list/data-list.component';
import { CreateDataComponent } from './create-data/create-data.component';
import { UpdateDataComponent } from './update-data/update-data.component';
import { DataDetailsComponent } from './data-details/data-details.component';

const routes: Routes = [
   { path: '', redirectTo: 'datas', pathMatch: 'full' },
   { path: 'datas', component: DataListComponent },
   { path: 'add', component: CreateDataComponent },
   { path: 'update/:id', component: UpdateDataComponent },
   { path: 'details/:id', component: DataDetailsComponent },
 ];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }