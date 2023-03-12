import { BrowserModule } 
 from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } 
  from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } 
  from '@angular/common/http';
import { DataListComponent } from './data-list/data-list.component';
import { CreateDataComponent } from './create-data/create-data.component';
import { UpdateDataComponent } from './update-data/update-data.component';
import { DataDetailsComponent } from './data-details/data-details.component';
@NgModule({
  declarations: [
    AppComponent,
    DataListComponent,
    CreateDataComponent,
    UpdateDataComponent,
    DataDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
